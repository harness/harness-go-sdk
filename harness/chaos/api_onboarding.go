package chaos

import (
	"context"
	"fmt"
	"net/http"
	"net/url"
	"regexp"
	"sort"
	"strings"
	"time"
)

// onboardingPollInterval is how often a service onboarding waits for discovery to finish.
var onboardingPollInterval = 5 * time.Second

// OnboardingClient calls the v2 onboarding API. Construct it with NewOnboardingClient.
// It is not a field on the generated APIClient, so swagger regeneration does not have to know about it.
type OnboardingClient struct {
	client *APIClient
	// PollInterval is how often FinishServiceOnboarding checks discovery. Zero uses 5 seconds.
	PollInterval time.Duration
}

// NewOnboardingClient returns a client for service onboarding.
func NewOnboardingClient(client *APIClient) *OnboardingClient {
	return &OnboardingClient{client: client}
}

// ServiceOnboardingRequest starts a manual onboarding session whose method is Service.
// EnvironmentID and InfrastructureID are the existing environment and chaos infrastructure.
// The create API also registers that chaos infrastructure. Set ServiceID or ServiceName
// for one discovered service. ServiceID is the Id shown under the service name, or the
// discovery record id. ServiceName is the service name shown in the list. Name and
// Identity override the values taken from that discovered service. Probes are existing
// project probe ids to associate.
type ServiceOnboardingRequest struct {
	EnvironmentID    string
	InfrastructureID string
	ServiceID        string
	ServiceName      string
	Name             string
	Identity         string
	Probes           []ServiceProbe
}

// ServiceProbe associates one project probe with the onboarded chaos service.
// Inputs are optional runtime values keyed by input name.
type ServiceProbe struct {
	ProbeID string
	Inputs  map[string]string
}

// ServiceOnboarding is the session the create and get calls return.
type ServiceOnboarding struct {
	OnboardingID      string `json:"onboardingID"`
	EnvironmentRef    string `json:"environmentRef"`
	InfrastructureRef string `json:"infrastructureRef"`
	Mode              string `json:"mode"`
	ManualMethod      string `json:"manualMethod"`
	Status            string `json:"status"`
	DiscoveryStatus   string `json:"-"`
	AgentIdentity     string `json:"-"`
	ServiceStatus     string `json:"-"`
	InstallCommand    string `json:"-"`
	ChaosServiceID    string `json:"-"`
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
		Status string `json:"status"`
	} `json:"onboardingChaosServiceStatus"`
}

type confirmDiscoveryResponse struct {
	AgentCreated       bool `json:"agentCreated"`
	AgentLinked        bool `json:"agentLinked"`
	DiscoveryCompleted bool `json:"discoveryCompleted"`
}

// OpenServiceOnboarding registers a Manual session with method Service and returns its
// infrastructure manifest. It does not look up an existing session; the caller does that.
func (c *OnboardingClient) OpenServiceOnboarding(ctx context.Context, scope Scope, req ServiceOnboardingRequest) (*ServiceOnboarding, *http.Response, error) {
	if req.EnvironmentID == "" || req.InfrastructureID == "" {
		return nil, nil, fmt.Errorf("environment id and infrastructure id are required")
	}
	opened, resp, err := c.createSession(ctx, scope, req)
	if err != nil {
		return opened, resp, err
	}
	manifest, manifestErr := c.InfrastructureManifest(ctx, scope, req.InfrastructureID, req.EnvironmentID)
	if manifestErr != nil {
		return opened, resp, manifestErr
	}
	opened.InstallCommand = manifest
	return opened, resp, nil
}

// FinishServiceOnboarding confirms discovery and creates one chaos service for the
// discovered service id. Probes on the request are associated with that service.
// A session that already has ChaosServiceID is returned unchanged.
func (c *OnboardingClient) FinishServiceOnboarding(ctx context.Context, scope Scope, current *ServiceOnboarding, req ServiceOnboardingRequest) (*ServiceOnboarding, *http.Response, error) {
	if current == nil || current.OnboardingID == "" {
		return current, nil, fmt.Errorf("onboarding id is required")
	}
	if req.ServiceID == "" && req.ServiceName == "" {
		return current, nil, fmt.Errorf("discovered service id or name is required")
	}
	if req.ServiceID != "" && req.ServiceName != "" {
		return current, nil, fmt.Errorf("set discovered service id or name, not both")
	}
	if current.ChaosServiceID != "" {
		return current, nil, nil
	}
	ready, waitResp, waitErr := c.waitForDiscovery(ctx, scope, current)
	if waitErr != nil {
		return current, waitResp, waitErr
	}
	ready.InstallCommand = current.InstallCommand
	created, createResp, createErr := c.createOnboardedService(ctx, scope, ready, req)
	if createErr != nil {
		return ready, createResp, createErr
	}
	ready.ChaosServiceID = created.Identity
	return ready, waitResp, nil
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
	}
}

func onboardingPath(id string) string {
	return "/rest/v2/onboarding/" + url.PathEscape(id)
}

// ChaosService is one chaos service created from a discovered service.
type ChaosService struct {
	Identity           string
	Name               string
	AgentID            string
	EnvironmentID      string
	InfrastructureID   string
	ExternalServiceID  string
	InfrastructureType string
	Probes             []ServiceProbe
}

// ChaosServiceRequest is the body for creating or updating one chaos service.
// Name and AgentID can be left empty on update; they are read from the current service.
type ChaosServiceRequest struct {
	Identity          string
	Name              string
	AgentID           string
	EnvironmentID     string
	InfrastructureID  string
	ExternalServiceID string
	Probes            []ServiceProbe
}

// CreateChaosService creates one chaos service on an existing onboarding session.
// It does not register another session.
func (c *OnboardingClient) CreateChaosService(ctx context.Context, scope Scope, onboardingID string, req ServiceOnboardingRequest) (*ChaosService, *http.Response, error) {
	if onboardingID == "" {
		return nil, nil, fmt.Errorf("onboarding id is required")
	}
	if req.ServiceID == "" && req.ServiceName == "" {
		return nil, nil, fmt.Errorf("discovered service id or name is required")
	}
	if req.ServiceID != "" && req.ServiceName != "" {
		return nil, nil, fmt.Errorf("set discovered service id or name, not both")
	}
	session, resp, err := c.Get(ctx, scope, onboardingID)
	if err != nil {
		return nil, resp, err
	}
	if session.DiscoveryStatus != "COMPLETED" {
		return nil, resp, fmt.Errorf("discovery for onboarding %s is %s", onboardingID, session.DiscoveryStatus)
	}
	if req.InfrastructureID != "" && session.InfrastructureRef != "" && req.InfrastructureID != session.InfrastructureRef {
		return nil, resp, fmt.Errorf("onboarding %s is for infrastructure %s", onboardingID, session.InfrastructureRef)
	}
	return c.createOnboardedService(ctx, scope, session, req)
}

func (c *OnboardingClient) createOnboardedService(ctx context.Context, scope Scope, session *ServiceOnboarding, req ServiceOnboardingRequest) (*ChaosService, *http.Response, error) {
	discoveryAgent := session.AgentIdentity
	if discoveryAgent == "" {
		discoveryAgent = session.InfrastructureRef
	}
	listed, err := c.listDiscoveredServices(ctx, scope, discoveryAgent, session.EnvironmentRef)
	if err != nil {
		return nil, nil, err
	}
	found, err := matchDiscoveredService(listed, req)
	if err != nil {
		return nil, nil, err
	}
	name := req.Name
	if name == "" {
		name = found.Name
	}
	if name == "" {
		name = req.ServiceID
	}
	if name == "" {
		name = req.ServiceName
	}
	identity := req.Identity
	if identity == "" {
		namespace := ""
		switch {
		case found.Spec.Kubernetes != nil && found.Spec.Kubernetes.Namespace != "":
			namespace = found.Spec.Kubernetes.Namespace
		case found.Spec.KubernetesWorkload != nil && found.Spec.KubernetesWorkload.Namespace != "":
			namespace = found.Spec.KubernetesWorkload.Namespace
		case found.Spec.KubernetesService != nil:
			namespace = found.Spec.KubernetesService.Namespace
		}
		identity = chaosServiceIdentity(name, found.Type, namespace, discoveryAgent)
	}
	// A Kubernetes service uses the infrastructure id for both fields, matching the dialog.
	// Linux and Windows services use the agent id the dialog uses for that infrastructure.
	infrastructureID := session.InfrastructureRef
	infrastructureType := "KubernetesV2"
	agentID := infrastructureID
	switch found.Type {
	case "LinuxVM", "LinuxVMProcess":
		infrastructureType = "Linux"
		agentID = "custom-discovery-agent"
	case "WindowsVM", "WindowsVMProcess":
		infrastructureType = "Windows"
		agentID = "custom-discovery-agent"
	}
	body := chaosServiceBody{
		Identity:           identity,
		Name:               name,
		AgentID:            agentID,
		EnvironmentID:      session.EnvironmentRef,
		InfrastructureID:   infrastructureID,
		InfrastructureType: infrastructureType,
		ExternalServiceID:  found.ID,
		OnboardingID:       session.OnboardingID,
		Probes:             encodeProbes(req.Probes),
	}
	return c.writeChaosService(ctx, scope, http.MethodPost, "/v3/chaos-services", body)
}

// GetChaosService reads one chaos service by identity.
func (c *OnboardingClient) GetChaosService(ctx context.Context, scope Scope, identity string) (*ChaosService, *http.Response, error) {
	if identity == "" {
		return nil, nil, fmt.Errorf("chaos service identity is required")
	}
	api, err := c.api()
	if err != nil {
		return nil, nil, err
	}
	var raw chaosServiceResponse
	resp, err := api.doChaosJSON(ctx, http.MethodGet, "/v3/chaos-services/"+url.PathEscape(identity), scopeQuery(scope), nil, &raw)
	if err != nil {
		return nil, resp, err
	}
	got := raw.service()
	return &got, resp, nil
}

// UpdateChaosService replaces the probe associations on an existing chaos service.
func (c *OnboardingClient) UpdateChaosService(ctx context.Context, scope Scope, req ChaosServiceRequest) (*ChaosService, *http.Response, error) {
	if req.Identity == "" {
		return nil, nil, fmt.Errorf("chaos service identity is required")
	}
	// The update API leaves infrastructureType unchanged when it is omitted, and
	// replaces it when it is sent. Always send the type already stored so a probe
	// update cannot clear Linux or Windows.
	current, resp, err := c.GetChaosService(ctx, scope, req.Identity)
	if err != nil {
		return nil, resp, err
	}
	if req.Name == "" {
		req.Name = current.Name
	}
	if req.AgentID == "" {
		req.AgentID = current.AgentID
	}
	if req.ExternalServiceID == "" {
		req.ExternalServiceID = current.ExternalServiceID
	}
	if req.EnvironmentID == "" {
		req.EnvironmentID = current.EnvironmentID
	}
	if req.InfrastructureID == "" {
		req.InfrastructureID = current.InfrastructureID
	}
	body := chaosServiceBody{
		Identity:           req.Identity,
		Name:               req.Name,
		AgentID:            req.AgentID,
		EnvironmentID:      req.EnvironmentID,
		InfrastructureID:   req.InfrastructureID,
		InfrastructureType: current.InfrastructureType,
		ExternalServiceID:  req.ExternalServiceID,
		Probes:             encodeProbes(req.Probes),
	}
	return c.writeChaosService(ctx, scope, http.MethodPut, "/v3/chaos-services/"+url.PathEscape(req.Identity), body)
}

// DeleteChaosService removes one chaos service. A missing service is not an error.
func (c *OnboardingClient) DeleteChaosService(ctx context.Context, scope Scope, identity string) (*http.Response, error) {
	if identity == "" {
		return nil, nil
	}
	api, err := c.api()
	if err != nil {
		return nil, err
	}
	resp, err := api.doChaosJSON(ctx, http.MethodDelete, "/v3/chaos-services/"+url.PathEscape(identity), scopeQuery(scope), nil, nil)
	if err != nil && ((resp != nil && resp.StatusCode == http.StatusNotFound) || OnboardingGone(err)) {
		return resp, nil
	}
	return resp, err
}

func (c *OnboardingClient) writeChaosService(ctx context.Context, scope Scope, method, path string, body chaosServiceBody) (*ChaosService, *http.Response, error) {
	api, err := c.api()
	if err != nil {
		return nil, nil, err
	}
	var raw chaosServiceResponse
	resp, err := api.doChaosJSON(ctx, method, path, scopeQuery(scope), body, &raw)
	if err != nil {
		return nil, resp, err
	}
	got := raw.service()
	if got.Identity == "" {
		got.Identity = body.Identity
	}
	return &got, resp, nil
}

func (c *OnboardingClient) listDiscoveredServices(ctx context.Context, scope Scope, agentID, environmentID string) ([]discoveredService, error) {
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
		Items []discoveredService `json:"items"`
	}
	path := base + "/" + url.PathEscape(agentID) + "/discoveredservices"
	if _, err = api.doAbsoluteJSON(ctx, http.MethodGet, path, q, nil, &body); err != nil {
		return nil, fmt.Errorf("list discovered services for agent %s: %w", agentID, err)
	}
	return body.Items, nil
}

// matchDiscoveredService accepts the Id shown in the UI, the discovery record id,
// or the service name. The create API is called with the record id.
func matchDiscoveredService(listed []discoveredService, req ServiceOnboardingRequest) (discoveredService, error) {
	if req.ServiceName != "" {
		var matches []discoveredService
		for _, item := range listed {
			if !item.Removed && item.Name == req.ServiceName {
				matches = append(matches, item)
			}
		}
		switch len(matches) {
		case 1:
			return matches[0], nil
		case 0:
			return discoveredService{}, fmt.Errorf("discovered service %s was not found", req.ServiceName)
		default:
			labels := make([]string, 0, len(matches))
			for _, item := range matches {
				if item.Identity != "" {
					labels = append(labels, item.Identity)
				} else {
					labels = append(labels, item.ID)
				}
			}
			return discoveredService{}, fmt.Errorf("discovered service name %q matches %s; set discovered_service_id", req.ServiceName, strings.Join(labels, ", "))
		}
	}

	var found discoveredService
	var recordIDMatch *discoveredService
	for i := range listed {
		item := listed[i]
		if item.Removed {
			continue
		}
		if item.Identity != "" && item.Identity == req.ServiceID {
			return item, nil
		}
		if item.ID == req.ServiceID {
			recordIDMatch = &listed[i]
		}
	}
	if recordIDMatch != nil {
		found = *recordIDMatch
	}
	if found.ID == "" {
		return discoveredService{}, fmt.Errorf("discovered service %s was not found", req.ServiceID)
	}
	return found, nil
}

type discoveredService struct {
	ID            string `json:"id"`
	Identity      string `json:"identity"`
	Name          string `json:"name"`
	Type          string `json:"type"`
	AgentIdentity string `json:"agentIdentity"`
	Removed       bool   `json:"removed"`
	Spec          struct {
		Kubernetes         *discoveredNamespace `json:"kubernetes"`
		KubernetesWorkload *discoveredNamespace `json:"kubernetesWorkload"`
		KubernetesService  *discoveredNamespace `json:"kubernetesService"`
	} `json:"spec"`
}

type discoveredNamespace struct {
	Namespace string `json:"namespace"`
}

type chaosServiceBody struct {
	Identity           string              `json:"identity"`
	Name               string              `json:"name"`
	AgentID            string              `json:"agentId"`
	EnvironmentID      string              `json:"environmentId"`
	InfrastructureID   string              `json:"infrastructureId"`
	InfrastructureType string              `json:"infrastructureType,omitempty"`
	ExternalServiceID  string              `json:"externalServiceId"`
	OnboardingID       string              `json:"onboardingId,omitempty"`
	Probes             []chaosServiceProbe `json:"probes"`
}

type chaosServiceProbe struct {
	ProbeID string            `json:"probeId"`
	Inputs  []chaosProbeInput `json:"inputs,omitempty"`
}

type chaosProbeInput struct {
	Name  string `json:"name"`
	Value string `json:"value"`
}

type chaosServiceResponse struct {
	Identity           string              `json:"identity"`
	Name               string              `json:"name"`
	AgentID            string              `json:"agentId"`
	EnvironmentID      string              `json:"environmentId"`
	InfrastructureID   string              `json:"infrastructureId"`
	InfrastructureType string              `json:"infrastructureType"`
	ExternalServiceID  string              `json:"externalServiceId"`
	Probes             []chaosServiceProbe `json:"probes"`
}

func (raw chaosServiceResponse) service() ChaosService {
	probes := make([]ServiceProbe, 0, len(raw.Probes))
	for _, probe := range raw.Probes {
		item := ServiceProbe{ProbeID: probe.ProbeID}
		if len(probe.Inputs) > 0 {
			item.Inputs = map[string]string{}
			for _, input := range probe.Inputs {
				item.Inputs[input.Name] = input.Value
			}
		}
		probes = append(probes, item)
	}
	return ChaosService{
		Identity:           raw.Identity,
		Name:               raw.Name,
		AgentID:            raw.AgentID,
		EnvironmentID:      raw.EnvironmentID,
		InfrastructureID:   raw.InfrastructureID,
		InfrastructureType: raw.InfrastructureType,
		ExternalServiceID:  raw.ExternalServiceID,
		Probes:             probes,
	}
}

func encodeProbes(probes []ServiceProbe) []chaosServiceProbe {
	out := make([]chaosServiceProbe, 0, len(probes))
	for _, probe := range probes {
		item := chaosServiceProbe{ProbeID: probe.ProbeID}
		if len(probe.Inputs) > 0 {
			names := make([]string, 0, len(probe.Inputs))
			for name := range probe.Inputs {
				names = append(names, name)
			}
			sort.Strings(names)
			item.Inputs = make([]chaosProbeInput, 0, len(names))
			for _, name := range names {
				item.Inputs = append(item.Inputs, chaosProbeInput{Name: name, Value: probe.Inputs[name]})
			}
		}
		out = append(out, item)
	}
	return out
}

// chaosServiceIdentity matches the onboarding API: trim, lowercase, and remove
// spaces from the name. The namespace is included only for a Kubernetes service
// that has one. The agent id is the discovery agent.
func chaosServiceIdentity(name, serviceType, namespace, agentID string) string {
	name = strings.ToLower(strings.ReplaceAll(strings.TrimSpace(name), " ", ""))
	kubernetes := serviceType == "Kubernetes" || serviceType == "KubernetesService" || serviceType == "KubernetesWorkload" || serviceType == "KubernetesNode"
	if kubernetes && namespace != "" {
		return fmt.Sprintf("%s-%s-%s", name, namespace, agentID)
	}
	return fmt.Sprintf("%s-%s", name, agentID)
}
