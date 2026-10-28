package chaos

import (
	"context"
	"fmt"
	"net/http"
	"net/url"
	"regexp"
	"strconv"
	"strings"
	"time"
)

// onboardingPollInterval is how often a service onboarding waits for discovery to finish.
var onboardingPollInterval = 5 * time.Second

// OnboardingClient calls the v2 onboarding API. Construct it with NewOnboardingClient.
// It is not a field on the generated APIClient, so swagger regeneration does not have to know about it.
type OnboardingClient struct {
	client *APIClient
	// PollInterval is how often CreateServiceOnboarding checks discovery. Zero uses 5 seconds.
	PollInterval time.Duration
}

// NewOnboardingClient returns a client for service onboarding.
func NewOnboardingClient(client *APIClient) *OnboardingClient {
	return &OnboardingClient{client: client}
}

// ServiceOnboardingRequest starts a manual onboarding session whose method is Service.
// EnvironmentID and InfrastructureID are the existing environment and chaos infrastructure.
// The create API also registers that chaos infrastructure. ServiceIDs are discovered
// service ids. Leave ServiceIDs empty to onboard every service discovery finds.
type ServiceOnboardingRequest struct {
	EnvironmentID    string
	InfrastructureID string
	ServiceIDs       []string
}

// ServiceOnboarding is the session the create and get calls return.
type ServiceOnboarding struct {
	OnboardingID        string   `json:"onboardingID"`
	EnvironmentRef      string   `json:"environmentRef"`
	InfrastructureRef   string   `json:"infrastructureRef"`
	Mode                string   `json:"mode"`
	ManualMethod        string   `json:"manualMethod"`
	Status              string   `json:"status"`
	DiscoveryStatus     string   `json:"-"`
	AgentIdentity       string   `json:"-"`
	ServiceStatus       string   `json:"-"`
	OnboardedCount      int      `json:"-"`
	OnboardedServiceIDs []string `json:"-"`
	InstallCommand      string   `json:"-"`
}

type serviceOnboardingResponse struct {
	OnboardingID            string `json:"onboardingID"`
	EnvironmentRef          string `json:"environmentRef"`
	InfrastructureRef       string `json:"infrastructureRef"`
	Mode                    string `json:"mode"`
	ManualMethod            string `json:"manualMethod"`
	Status                  string `json:"status"`
	DiscoveredServiceStatus struct {
		Status          string `json:"status"`
		SDAgentIdentity string `json:"sDAgentIdentity"`
	} `json:"discoveredServiceStatus"`
	OnboardingChaosServiceStatus struct {
		Status            string `json:"status"`
		ChaosServiceCount int    `json:"chaosServiceCount"`
	} `json:"onboardingChaosServiceStatus"`
}

type onboardServicesRequest struct {
	ExternalServiceIDs []string `json:"externalServiceIds"`
}

// OnboardServicesResult is what the onboard-services call returns.
type OnboardServicesResult struct {
	OnboardedCount      int      `json:"onboardedCount"`
	OnboardedServiceIDs []string `json:"onboardedServiceIds"`
}

type confirmDiscoveryResponse struct {
	AgentCreated       bool `json:"agentCreated"`
	AgentLinked        bool `json:"agentLinked"`
	DiscoveryCompleted bool `json:"discoveryCompleted"`
}

// CreateServiceOnboarding starts a Manual session with manualMethod Service, or resumes one
// that already exists for the same infrastructure. It confirms discovery, onboards the
// services, and confirms that step. The session status stays PENDING because the risk-scan step is left queued.
func (c *OnboardingClient) CreateServiceOnboarding(ctx context.Context, scope Scope, req ServiceOnboardingRequest) (*ServiceOnboarding, *http.Response, error) {
	opened, resp, err := c.OpenServiceOnboarding(ctx, scope, req)
	if err != nil {
		return opened, resp, err
	}
	return c.FinishServiceOnboarding(ctx, scope, opened, req.ServiceIDs)
}

// OpenServiceOnboarding returns the session and its infrastructure manifest without waiting for discovery.
// A retry uses the session already created for this infrastructure instead of registering it again.
func (c *OnboardingClient) OpenServiceOnboarding(ctx context.Context, scope Scope, req ServiceOnboardingRequest) (*ServiceOnboarding, *http.Response, error) {
	if req.EnvironmentID == "" || req.InfrastructureID == "" {
		return nil, nil, fmt.Errorf("environment id and infrastructure id are required")
	}
	existing, resp, err := c.FindServiceOnboarding(ctx, scope, req.InfrastructureID)
	if err != nil {
		return nil, resp, err
	}
	if existing == nil {
		existing, resp, err = c.createSession(ctx, scope, req)
		if err != nil {
			return existing, resp, err
		}
	}
	manifest, manifestErr := c.InfrastructureManifest(ctx, scope, req.InfrastructureID, req.EnvironmentID)
	if manifestErr != nil {
		return existing, resp, manifestErr
	}
	existing.InstallCommand = manifest
	return existing, resp, nil
}

// FinishServiceOnboarding confirms discovery, onboards services, and confirms the service step.
func (c *OnboardingClient) FinishServiceOnboarding(ctx context.Context, scope Scope, current *ServiceOnboarding, serviceIDs []string) (*ServiceOnboarding, *http.Response, error) {
	if current == nil || current.OnboardingID == "" {
		return current, nil, fmt.Errorf("onboarding id is required")
	}
	if current.ServiceStatus == "COMPLETED" {
		return c.refreshOnboarded(ctx, scope, current)
	}
	ready, waitResp, waitErr := c.waitForDiscovery(ctx, scope, current)
	if waitErr != nil {
		ready = current
		if ready == nil {
			return current, waitResp, waitErr
		}
		ready.InstallCommand = current.InstallCommand
		return ready, waitResp, waitErr
	}
	ready.InstallCommand = current.InstallCommand
	out := ready
	if len(serviceIDs) == 0 {
		agentID := ready.AgentIdentity
		if agentID == "" {
			agentID = ready.InfrastructureRef
		}
		listed, listErr := c.DiscoveredServiceIDs(ctx, scope, agentID, environmentID(ready))
		if listErr != nil {
			return out, waitResp, listErr
		}
		if len(listed) == 0 {
			return out, waitResp, fmt.Errorf("onboarding %s discovery completed with no services", out.OnboardingID)
		}
		serviceIDs = listed
	}
	if out.ServiceStatus != "COMPLETED" {
		onboarded, onboardResp, onboardErr := c.OnboardServices(ctx, scope, out.OnboardingID, serviceIDs)
		if onboardErr != nil {
			return out, onboardResp, onboardErr
		}
		out.OnboardedCount = onboarded.OnboardedCount
		out.OnboardedServiceIDs = onboarded.OnboardedServiceIDs
		if confirmResp, confirmErr := c.ConfirmServiceOnboarding(ctx, scope, out.OnboardingID); confirmErr != nil {
			return out, confirmResp, confirmErr
		}
	}
	latest, getResp, getErr := c.refreshOnboarded(ctx, scope, out)
	if getErr != nil {
		return out, getResp, getErr
	}
	return latest, waitResp, nil
}

func (c *OnboardingClient) createSession(ctx context.Context, scope Scope, req ServiceOnboardingRequest) (*ServiceOnboarding, *http.Response, error) {
	api, err := c.api()
	if err != nil {
		return nil, nil, err
	}
	mode := MANUAL_V2OnboardingOnboardingMode
	method := SERVICE_V2OnboardingManualMethod
	body := V2OnboardingV2OnboardingRequest{
		EnvironmentRef:    req.EnvironmentID,
		InfrastructureRef: req.InfrastructureID,
		Mode:              &mode,
		ManualMethod:      &method,
	}
	var raw serviceOnboardingResponse
	resp, err := api.doChaosJSON(ctx, http.MethodPost, "/rest/v2/onboarding", scopeQuery(scope), &body, &raw)
	if err != nil {
		return nil, resp, err
	}
	return raw.session(), resp, nil
}

// FindServiceOnboarding returns the newest Service session for this infrastructure, or nil when there is none.
func (c *OnboardingClient) FindServiceOnboarding(ctx context.Context, scope Scope, infrastructureID string) (*ServiceOnboarding, *http.Response, error) {
	api, err := c.api()
	if err != nil {
		return nil, nil, err
	}
	q := scopeQuery(scope)
	// The list API applies search as a MongoDB regex. Escape it so an id
	// containing "." or "+" still matches that infrastructure exactly.
	q.Set("search", regexp.QuoteMeta(infrastructureID))
	q.Set("page", "0")
	q.Set("limit", "20")
	var body struct {
		Data []serviceOnboardingResponse `json:"data"`
	}
	resp, err := api.doChaosJSON(ctx, http.MethodGet, "/rest/v2/onboarding", q, nil, &body)
	if err != nil {
		return nil, resp, err
	}
	for _, raw := range body.Data {
		session := raw.session()
		if session.InfrastructureRef == infrastructureID && session.ManualMethod == "Service" {
			return session, resp, nil
		}
	}
	return nil, resp, nil
}

// Get returns one onboarding session.
func (c *OnboardingClient) Get(ctx context.Context, scope Scope, onboardingID string) (*ServiceOnboarding, *http.Response, error) {
	api, err := c.api()
	if err != nil {
		return nil, nil, err
	}
	var raw serviceOnboardingResponse
	resp, err := api.doChaosJSON(ctx, http.MethodGet, onboardingPath(onboardingID), scopeQuery(scope), nil, &raw)
	if err != nil {
		return nil, resp, err
	}
	return raw.session(), resp, nil
}

// OnboardServices registers the selected discovered services on a manual session.
// Discovery must already be completed. externalServiceIds are the discovered service external ids.
func (c *OnboardingClient) OnboardServices(ctx context.Context, scope Scope, onboardingID string, serviceIDs []string) (*OnboardServicesResult, *http.Response, error) {
	if onboardingID == "" {
		return nil, nil, fmt.Errorf("onboarding id is required")
	}
	if len(serviceIDs) == 0 {
		return nil, nil, fmt.Errorf("at least one service id is required")
	}
	api, err := c.api()
	if err != nil {
		return nil, nil, err
	}
	var out OnboardServicesResult
	path := "/rest/v2/onboarding-onboard-services/" + url.PathEscape(onboardingID)
	resp, err := api.doChaosJSON(ctx, http.MethodPost, path, scopeQuery(scope), &onboardServicesRequest{ExternalServiceIDs: serviceIDs}, &out)
	if err != nil {
		return nil, resp, err
	}
	return &out, resp, nil
}

// ConfirmServiceOnboarding marks the service onboarding step completed.
func (c *OnboardingClient) ConfirmServiceOnboarding(ctx context.Context, scope Scope, onboardingID string) (*http.Response, error) {
	if onboardingID == "" {
		return nil, fmt.Errorf("onboarding id is required")
	}
	api, err := c.api()
	if err != nil {
		return nil, err
	}
	path := "/rest/v2/onboarding-confirm-service-onboarding/" + url.PathEscape(onboardingID)
	return api.doChaosJSON(ctx, http.MethodPost, path, scopeQuery(scope), nil, nil)
}

// ConfirmDiscovery asks the manual session to mark discovery completed once the agent is installed.
func (c *OnboardingClient) ConfirmDiscovery(ctx context.Context, scope Scope, onboardingID string) (*confirmDiscoveryResponse, *http.Response, error) {
	api, err := c.api()
	if err != nil {
		return nil, nil, err
	}
	var out confirmDiscoveryResponse
	path := "/rest/v2/onboarding-confirm-discovery/" + url.PathEscape(onboardingID)
	resp, err := api.doChaosJSON(ctx, http.MethodPost, path, scopeQuery(scope), nil, &out)
	if err != nil {
		return nil, resp, err
	}
	return &out, resp, nil
}

// Delete marks the onboarding session and its network maps removed. The chaos infrastructure and discovery agent stay.
func (c *OnboardingClient) Delete(ctx context.Context, scope Scope, onboardingID string) (*http.Response, error) {
	if onboardingID == "" {
		return nil, fmt.Errorf("onboarding id is required")
	}
	api, err := c.api()
	if err != nil {
		return nil, err
	}
	return api.doChaosJSON(ctx, http.MethodDelete, onboardingPath(onboardingID), scopeQuery(scope), nil, nil)
}

// OnboardedServiceIDs reads the services this session onboarded.
func (c *OnboardingClient) OnboardedServiceIDs(ctx context.Context, scope Scope, onboardingID string) ([]string, error) {
	api, err := c.api()
	if err != nil {
		return nil, err
	}
	var ids []string
	for page := 0; ; page++ {
		q := scopeQuery(scope)
		q.Set("page", strconv.Itoa(page))
		q.Set("limit", "100")
		var body struct {
			Services []struct {
				ExternalServiceID string `json:"externalServiceId"`
			} `json:"services"`
			Pagination struct {
				TotalPages int `json:"totalPages"`
			} `json:"pagination"`
		}
		if _, err = api.doChaosJSON(ctx, http.MethodGet, onboardingPath(onboardingID)+"/services", q, nil, &body); err != nil {
			return nil, err
		}
		for _, service := range body.Services {
			if service.ExternalServiceID != "" {
				ids = append(ids, service.ExternalServiceID)
			}
		}
		if page+1 >= body.Pagination.TotalPages || len(body.Services) == 0 {
			break
		}
	}
	return ids, nil
}

// InfrastructureManifest reads the chaos infrastructure and returns the install manifest.
func (c *OnboardingClient) InfrastructureManifest(ctx context.Context, scope Scope, infrastructureID, environmentID string) (string, error) {
	api, err := c.api()
	if err != nil {
		return "", err
	}
	if api.ChaosSdkApi == nil {
		return "", fmt.Errorf("chaos infrastructure client is not configured")
	}
	infra, _, err := api.ChaosSdkApi.GetInfraV2(ctx, infrastructureID, scope.AccountID, scope.OrgID, scope.ProjectID, environmentID)
	if err != nil {
		return "", fmt.Errorf("get infrastructure %s: %w", infrastructureID, err)
	}
	body := map[string]string{
		"identity":       infrastructureID,
		"environmentID":  environmentID,
		"infraNamespace": infra.InfraNamespace,
		"serviceAccount": infra.ServiceAccount,
	}
	var out struct {
		Yaml string `json:"yaml"`
	}
	if _, err = api.doChaosJSON(ctx, http.MethodPost, "/rest/v2/infrastructure/yaml", scopeQuery(scope), body, &out); err != nil {
		return "", fmt.Errorf("get infrastructure manifest: %w", err)
	}
	if out.Yaml == "" {
		return "", fmt.Errorf("infrastructure %s returned an empty install manifest", infrastructureID)
	}
	return out.Yaml, nil
}

// OnboardingGone reports a get that failed because the session is no longer stored.
func OnboardingGone(err error) bool {
	if err == nil {
		return false
	}
	msg := err.Error()
	return strings.Contains(msg, "no documents") || strings.Contains(strings.ToLower(msg), "does not exist")
}

func (c *OnboardingClient) waitForDiscovery(ctx context.Context, scope Scope, current *ServiceOnboarding) (*ServiceOnboarding, *http.Response, error) {
	var resp *http.Response
	for {
		switch current.DiscoveryStatus {
		case "ERRORED", "SKIPPED":
			return current, resp, fmt.Errorf("discovery for onboarding %s is %s", current.OnboardingID, current.DiscoveryStatus)
		}
		confirmed, confirmResp, err := c.ConfirmDiscovery(ctx, scope, current.OnboardingID)
		resp = confirmResp
		if err != nil {
			return current, resp, err
		}
		if confirmed.DiscoveryCompleted || current.DiscoveryStatus == "COMPLETED" {
			next, getResp, getErr := c.Get(ctx, scope, current.OnboardingID)
			if getErr != nil {
				return current, getResp, getErr
			}
			return next, resp, nil
		}
		select {
		case <-ctx.Done():
			return current, resp, fmt.Errorf("discovery for onboarding %s is still %s: apply the infrastructure manifest and wait for the discovery agent to finish installing: %w", current.OnboardingID, current.DiscoveryStatus, ctx.Err())
		case <-time.After(c.interval()):
		}
		next, nextResp, err := c.Get(ctx, scope, current.OnboardingID)
		resp = nextResp
		if err != nil {
			return current, resp, err
		}
		current = next
		if current.DiscoveryStatus == "COMPLETED" {
			return current, resp, nil
		}
	}
}

func (c *OnboardingClient) refreshOnboarded(ctx context.Context, scope Scope, current *ServiceOnboarding) (*ServiceOnboarding, *http.Response, error) {
	latest, resp, err := c.Get(ctx, scope, current.OnboardingID)
	if err != nil {
		return current, resp, err
	}
	latest.InstallCommand = current.InstallCommand
	if ids, listErr := c.OnboardedServiceIDs(ctx, scope, current.OnboardingID); listErr == nil {
		latest.OnboardedServiceIDs = ids
		if len(ids) > 0 {
			latest.OnboardedCount = len(ids)
		}
	} else if len(current.OnboardedServiceIDs) > 0 {
		latest.OnboardedServiceIDs = current.OnboardedServiceIDs
	}
	if latest.OnboardedCount == 0 {
		latest.OnboardedCount = current.OnboardedCount
	}
	return latest, resp, nil
}

func environmentID(session *ServiceOnboarding) string {
	if session == nil {
		return ""
	}
	return session.EnvironmentRef
}

// DiscoveredServiceIDs lists the ids discovery recorded for an agent. Service onboarding sends these as externalServiceIds.
func (c *OnboardingClient) DiscoveredServiceIDs(ctx context.Context, scope Scope, agentID, environmentID string) ([]string, error) {
	api, err := c.api()
	if err != nil {
		return nil, err
	}
	base, err := serviceDiscoveryAgentsURL(api.cfg.BasePath)
	if err != nil {
		return nil, err
	}
	q := scopeQuery(scope)
	q.Set("environmentIdentifier", environmentID)
	q.Set("all", "true")
	q.Set("page", "0")
	q.Set("limit", "0")
	var body struct {
		Items []struct {
			ID      string `json:"id"`
			Removed bool   `json:"removed"`
		} `json:"items"`
	}
	path := base + "/" + url.PathEscape(agentID) + "/discoveredservices"
	if _, err = api.doAbsoluteJSON(ctx, http.MethodGet, path, q, nil, &body); err != nil {
		return nil, fmt.Errorf("list discovered services for agent %s: %w", agentID, err)
	}
	ids := make([]string, 0, len(body.Items))
	for _, item := range body.Items {
		if item.ID != "" && !item.Removed {
			ids = append(ids, item.ID)
		}
	}
	return ids, nil
}

func (c *OnboardingClient) interval() time.Duration {
	if c != nil && c.PollInterval > 0 {
		return c.PollInterval
	}
	return onboardingPollInterval
}

func (c *OnboardingClient) api() (*APIClient, error) {
	if c == nil || c.client == nil {
		return nil, fmt.Errorf("onboarding client is not configured")
	}
	return c.client, nil
}

func (raw serviceOnboardingResponse) session() *ServiceOnboarding {
	return &ServiceOnboarding{
		OnboardingID:      raw.OnboardingID,
		EnvironmentRef:    raw.EnvironmentRef,
		InfrastructureRef: raw.InfrastructureRef,
		Mode:              raw.Mode,
		ManualMethod:      raw.ManualMethod,
		Status:            raw.Status,
		DiscoveryStatus:   raw.DiscoveredServiceStatus.Status,
		AgentIdentity:     raw.DiscoveredServiceStatus.SDAgentIdentity,
		ServiceStatus:     raw.OnboardingChaosServiceStatus.Status,
		OnboardedCount:    raw.OnboardingChaosServiceStatus.ChaosServiceCount,
	}
}

func onboardingPath(id string) string {
	return "/rest/v2/onboarding/" + url.PathEscape(id)
}
