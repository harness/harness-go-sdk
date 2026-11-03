package chaos

import (
	"context"
	"fmt"
	"net/http"
	"net/url"
	"strconv"
	"strings"
)

type ScanClient struct {
	client *APIClient
}

func NewScanClient(client *APIClient) *ScanClient {
	return &ScanClient{client: client}
}

func (c *ScanClient) Create(ctx context.Context, scope Scope, req ScanCreateRequest) (*Scan, *http.Response, error) {
	method, err := APIDetectionMethod(req.DetectionMethod)
	if err != nil {
		return nil, nil, err
	}
	req.DetectionMethod = method
	api, err := c.api()
	if err != nil {
		return nil, nil, err
	}
	var out Scan
	resp, err := api.doChaosJSON(ctx, http.MethodPost, "/v3/risk-scans", scopeQuery(scope), &req, &out)
	if err != nil {
		return nil, resp, err
	}
	return &out, resp, nil
}

func (c *ScanClient) CreatePipelineScan(ctx context.Context, scope Scope, req PipelineScanRequest) (*Scan, *http.Response, error) {
	if req.PipelineIdentity == "" || req.PipelineExecutionID == "" {
		return nil, nil, fmt.Errorf("pipeline identity and pipeline execution id are required")
	}
	return c.Create(ctx, scope, ScanCreateRequest{
		Identity:        req.Identity,
		Name:            req.Name,
		Description:     req.Description,
		Tags:            req.Tags,
		ScanType:        ScanTypePipelineExecution,
		DetectionMethod: req.DetectionMethod,
		Source: ScanSource{Pipeline: &PipelineScan{
			PipelineIdentity:    req.PipelineIdentity,
			PipelineExecutionID: req.PipelineExecutionID,
			NotifyID:            req.NotifyID,
		}},
	})
}

func (c *ScanClient) CreateInfrastructureScan(ctx context.Context, scope Scope, req InfrastructureScanRequest) (*Scan, *http.Response, error) {
	source, err := c.infrastructureSource(ctx, scope, req)
	if err != nil {
		return nil, nil, err
	}
	return c.Create(ctx, scope, ScanCreateRequest{
		Identity:        req.Identity,
		Name:            req.Name,
		Description:     req.Description,
		Tags:            req.Tags,
		ScanType:        ScanTypeDiscoveryAgent,
		DetectionMethod: req.DetectionMethod,
		Source:          ScanSource{DiscoveryAgent: source},
	})
}

func (c *ScanClient) Get(ctx context.Context, scope Scope, identity string) (*Scan, *http.Response, error) {
	api, err := c.api()
	if err != nil {
		return nil, nil, err
	}
	var out Scan
	resp, err := api.doChaosJSON(ctx, http.MethodGet, scanPath(identity), scopeQuery(scope), nil, &out)
	if err != nil {
		return nil, resp, err
	}
	return &out, resp, nil
}

func (c *ScanClient) Update(ctx context.Context, scope Scope, identity string, req ScanUpdateRequest) (*http.Response, error) {
	api, err := c.api()
	if err != nil {
		return nil, err
	}
	return api.doChaosJSON(ctx, http.MethodPut, scanPath(identity), scopeQuery(scope), &req, nil)
}

func (c *ScanClient) Delete(ctx context.Context, scope Scope, identity string) (*http.Response, error) {
	api, err := c.api()
	if err != nil {
		return nil, err
	}
	return api.doChaosJSON(ctx, http.MethodDelete, scanPath(identity), scopeQuery(scope), nil, nil)
}

func (c *ScanClient) GetReport(ctx context.Context, scope Scope, identity string) (*ScanReport, *http.Response, error) {
	api, err := c.api()
	if err != nil {
		return nil, nil, err
	}
	var out ScanReport
	resp, err := api.doChaosJSON(ctx, http.MethodGet, scanPath(identity)+"/report", scopeQuery(scope), nil, &out)
	if err != nil {
		return nil, resp, err
	}
	return &out, resp, nil
}

func (c *ScanClient) List(ctx context.Context, req ScanListRequest) (*ScanListResponse, *http.Response, error) {
	q := scopeQuery(req.Scope)
	addIntQuery(q, "page", req.Page)
	addIntQuery(q, "limit", req.Limit)
	addStringQuery(q, "search", req.Search)
	addStringQuery(q, "sortField", req.SortField)
	if req.SortAscending != nil {
		q.Set("sortAscending", strconv.FormatBool(*req.SortAscending))
	}
	addStringQuery(q, "scanType", req.ScanType)
	addStringQuery(q, "status", req.Status)
	addStringQuery(q, "pipelineIdentity", req.PipelineIdentity)
	addStringQuery(q, "agentIdentity", req.AgentIdentity)
	addStringQuery(q, "environmentIdentity", req.EnvironmentIdentity)
	addStringQuery(q, "tags", req.Tags)
	addInt64Query(q, "startTime", req.StartTime)
	addInt64Query(q, "endTime", req.EndTime)

	api, err := c.api()
	if err != nil {
		return nil, nil, err
	}
	var out ScanListResponse
	resp, err := api.doChaosJSON(ctx, http.MethodGet, "/v3/risk-scans", q, nil, &out)
	if err != nil {
		return nil, resp, err
	}
	return &out, resp, nil
}

func (c *ScanClient) Abort(ctx context.Context, scope Scope, identity string) (*http.Response, error) {
	api, err := c.api()
	if err != nil {
		return nil, err
	}
	return api.doChaosJSON(ctx, http.MethodPost, scanPath(identity)+"/abort", scopeQuery(scope), nil, nil)
}

func (c *ScanClient) Retry(ctx context.Context, scope Scope, identity string) (*http.Response, error) {
	api, err := c.api()
	if err != nil {
		return nil, err
	}
	return api.doChaosJSON(ctx, http.MethodPost, scanPath(identity)+"/retry", scopeQuery(scope), nil, nil)
}

func (c *ScanClient) api() (*APIClient, error) {
	if c == nil || c.client == nil {
		return nil, fmt.Errorf("scan client is not configured")
	}
	return c.client, nil
}

func APIDetectionMethod(method string) (string, error) {
	switch {
	case method == "":
		return "", nil
	case strings.EqualFold(method, "ai"):
		return DetectionMethodAI, nil
	case strings.EqualFold(method, "ruleBased"), strings.EqualFold(method, DetectionMethodRiskAnalyzer):
		return DetectionMethodRiskAnalyzer, nil
	default:
		return "", fmt.Errorf("detection method %q must be ai or ruleBased", method)
	}
}

type discoveryAgent struct {
	Identity              string `json:"identity"`
	Name                  string `json:"name"`
	EnvironmentIdentifier string `json:"environmentIdentifier"`
	ParentUniqueID        string `json:"parentUniqueId"`
}

func (c *ScanClient) infrastructureSource(ctx context.Context, scope Scope, req InfrastructureScanRequest) (*DiscoveryAgentScan, error) {
	if req.InfrastructureID == "" || req.EnvironmentID == "" {
		return nil, fmt.Errorf("infrastructure id and environment id are required")
	}
	api, err := c.api()
	if err != nil {
		return nil, err
	}
	if api.ChaosSdkApi == nil {
		return nil, fmt.Errorf("chaos infrastructure client is not configured")
	}

	infra, _, infraErr := api.ChaosSdkApi.GetInfraV2(ctx, req.InfrastructureID, scope.AccountID, scope.OrgID, scope.ProjectID, req.EnvironmentID)
	if infraErr != nil {
		return nil, fmt.Errorf("get infrastructure %s in environment %s: %w", req.InfrastructureID, req.EnvironmentID, infraErr)
	}
	agentID := infra.DiscoveryAgentID
	if agentID == "" {
		// Onboarding names the discovery agent with the infrastructure id.
		agentID = req.InfrastructureID
	}
	agent, err := c.getDiscoveryAgent(ctx, scope, agentID, req.EnvironmentID)
	if err != nil {
		if infra.DiscoveryAgentID == "" {
			return nil, fmt.Errorf("infrastructure %s in environment %s has no discovery agent", req.InfrastructureID, req.EnvironmentID)
		}
		return nil, fmt.Errorf("infrastructure %s in environment %s: %w", req.InfrastructureID, req.EnvironmentID, err)
	}
	return discoveryScanSource(agent, req.EnvironmentID), nil
}

func (c *ScanClient) getDiscoveryAgent(ctx context.Context, scope Scope, agentID, environmentID string) (discoveryAgent, error) {
	api, err := c.api()
	if err != nil {
		return discoveryAgent{}, err
	}
	base, err := serviceDiscoveryAgentsURL(api.cfg.BasePath)
	if err != nil {
		return discoveryAgent{}, err
	}
	q := scopeQuery(scope)
	q.Set("environmentIdentifier", environmentID)
	var agent discoveryAgent
	if _, err = api.doAbsoluteJSON(ctx, http.MethodGet, base+"/"+url.PathEscape(agentID), q, nil, &agent); err != nil {
		return discoveryAgent{}, fmt.Errorf("get discovery agent %s: %w", agentID, err)
	}
	if agent.Identity == "" {
		return discoveryAgent{}, fmt.Errorf("discovery agent %s was not found", agentID)
	}
	return agent, nil
}

func discoveryScanSource(agent discoveryAgent, environmentID string) *DiscoveryAgentScan {
	environment := agent.EnvironmentIdentifier
	if environment == "" {
		environment = environmentID
	}
	return &DiscoveryAgentScan{
		AgentIdentity:             agent.Identity,
		EnvironmentIdentity:       environment,
		EnvironmentParentUniqueID: agent.ParentUniqueID,
	}
}

func serviceDiscoveryAgentsURL(basePath string) (string, error) {
	const suffix = "/chaos/manager/api"
	trimmed := strings.TrimRight(basePath, "/")
	if !strings.HasSuffix(trimmed, suffix) {
		return "", fmt.Errorf("chaos base path %q must end with %s", basePath, suffix)
	}
	return strings.TrimSuffix(trimmed, suffix) + "/servicediscovery/api/v1/agents", nil
}

func scanPath(identity string) string {
	return "/v3/risk-scans/" + url.PathEscape(identity)
}

func addStringQuery(q url.Values, key, value string) {
	if value != "" {
		q.Set(key, value)
	}
}

func addIntQuery(q url.Values, key string, value int) {
	if value != 0 {
		q.Set(key, strconv.Itoa(value))
	}
}

func addInt64Query(q url.Values, key string, value int64) {
	if value != 0 {
		q.Set(key, strconv.FormatInt(value, 10))
	}
}
