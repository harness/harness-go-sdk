package chaos

import (
	"context"
	"fmt"
	"net/http"
	"net/url"
	"strconv"
)

// ScannedRiskClient reads risks produced by scans.
//
// Create, Update, and Delete are deliberately absent. When those operations
// are added to /v3/scanned-risks, implement them on this client and keep
// ScannedRisk as the request and response body so callers do not change shape.
type ScannedRiskClient struct {
	client *APIClient
}

// NewScannedRiskClient returns a read client for scanned risks.
func NewScannedRiskClient(client *APIClient) *ScannedRiskClient {
	return &ScannedRiskClient{client: client}
}

// Get returns one scanned risk by identity.
func (c *ScannedRiskClient) Get(ctx context.Context, scope Scope, identity string) (*ScannedRisk, *http.Response, error) {
	api, err := c.api()
	if err != nil {
		return nil, nil, err
	}
	var out ScannedRiskGetResponse
	resp, err := api.doChaosJSON(ctx, http.MethodGet, scannedRiskPath(identity), scopeQuery(scope), nil, &out)
	if err != nil {
		return nil, resp, err
	}
	return &out.ScannedRisk, resp, nil
}

// List returns scanned risks in the scope.
func (c *ScannedRiskClient) List(ctx context.Context, req ScannedRiskListRequest) (*ScannedRiskListResponse, *http.Response, error) {
	q := scopeQuery(req.Scope)
	addIntQuery(q, "page", req.Page)
	addIntQuery(q, "limit", req.Limit)
	addStringQuery(q, "search", req.Search)
	addStringQuery(q, "sortField", req.SortField)
	if req.SortAscending != nil {
		q.Set("sortAscending", strconv.FormatBool(*req.SortAscending))
	}
	addStringQuery(q, "severity", req.Severity)
	addStringQuery(q, "source", req.Source)
	addStringQuery(q, "validationType", req.ValidationType)
	addStringQuery(q, "serviceIdentity", req.ServiceIdentity)
	addStringQuery(q, "agentIdentity", req.AgentIdentity)
	addStringQuery(q, "environmentIdentity", req.EnvironmentIdentity)
	addStringQuery(q, "riskRuleId", req.RiskRuleID)
	addStringQuery(q, "tags", req.Tags)
	addInt64Query(q, "startTime", req.StartTime)
	addInt64Query(q, "endTime", req.EndTime)
	if req.IncludeAllScope != nil {
		q.Set("includeAllScope", strconv.FormatBool(*req.IncludeAllScope))
	}
	addStringQuery(q, "serviceFilterV2Kind", req.ServiceFilter.Kind)
	addStringQuery(q, "serviceFilterV2DiscoveredServiceAgentID", req.ServiceFilter.DiscoveredServiceAgentID)
	addStringQuery(q, "serviceFilterV2DiscoveredServiceEnvironmentID", req.ServiceFilter.DiscoveredServiceEnvironmentID)
	addStringQuery(q, "serviceFilterV2DiscoveredServiceExternalID", req.ServiceFilter.DiscoveredServiceExternalID)
	addStringQuery(q, "serviceFilterV2ChaosServiceIdentity", req.ServiceFilter.ChaosServiceIdentity)
	addStringQuery(q, "serviceFilterV2CDServiceIdentity", req.ServiceFilter.CDServiceIdentity)
	addStringQuery(q, "serviceFilterV2CDServiceEnvironmentID", req.ServiceFilter.CDServiceEnvironmentID)
	addStringQuery(q, "serviceFilterV2CDServiceInfrastructureID", req.ServiceFilter.CDServiceInfrastructureID)

	api, err := c.api()
	if err != nil {
		return nil, nil, err
	}
	var out ScannedRiskListResponse
	resp, err := api.doChaosJSON(ctx, http.MethodGet, "/v3/scanned-risks", q, nil, &out)
	if err != nil {
		return nil, resp, err
	}
	return &out, resp, nil
}

// ListOccurrences returns detections for one scanned risk.
func (c *ScannedRiskClient) ListOccurrences(ctx context.Context, identity string, req ScannedRiskOccurrencesRequest) (*ScannedRiskOccurrencesResponse, *http.Response, error) {
	q := scopeQuery(req.Scope)
	addIntQuery(q, "page", req.Page)
	addIntQuery(q, "limit", req.Limit)
	addStringQuery(q, "search", req.Search)
	addStringQuery(q, "sortField", req.SortField)
	if req.SortAscending != nil {
		q.Set("sortAscending", strconv.FormatBool(*req.SortAscending))
	}
	addStringQuery(q, "source", req.Source)
	addStringQuery(q, "scanType", req.ScanType)
	addStringQuery(q, "validationType", req.ValidationType)
	addInt64Query(q, "startTime", req.StartTime)
	addInt64Query(q, "endTime", req.EndTime)

	api, err := c.api()
	if err != nil {
		return nil, nil, err
	}
	var out ScannedRiskOccurrencesResponse
	path := scannedRiskPath(identity) + "/occurrences"
	resp, err := api.doChaosJSON(ctx, http.MethodGet, path, q, nil, &out)
	if err != nil {
		return nil, resp, err
	}
	return &out, resp, nil
}

func (c *ScannedRiskClient) api() (*APIClient, error) {
	if c == nil || c.client == nil {
		return nil, fmt.Errorf("scanned risk client is not configured")
	}
	return c.client, nil
}

func scannedRiskPath(identity string) string {
	return "/v3/scanned-risks/" + url.PathEscape(identity)
}
