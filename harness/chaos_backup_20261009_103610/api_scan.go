package chaos

import (
	"context"
	"fmt"
	"net/http"
	"net/url"
	"strconv"
)

// ScanClient calls the risk scan API (POST/GET/PUT/DELETE /v3/risk-scans).
// Construct it with NewScanClient. It is intentionally not wired into the
// generated APIClient struct, so swagger regeneration does not have to know about it.
type ScanClient struct {
	client *APIClient
}

// NewScanClient returns a client for pipeline and infrastructure scans.
func NewScanClient(client *APIClient) *ScanClient {
	return &ScanClient{client: client}
}

// Create starts a scan. The returned scan is usually still running; risks
// appear on a later GetReport once status is terminal.
func (c *ScanClient) Create(ctx context.Context, scope Scope, req ScanCreateRequest) (*Scan, *http.Response, error) {
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

// Get returns one scan by identity.
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

// Update changes name, description, and tags. Scan type and source cannot be updated.
func (c *ScanClient) Update(ctx context.Context, scope Scope, identity string, req ScanUpdateRequest) (*http.Response, error) {
	api, err := c.api()
	if err != nil {
		return nil, err
	}
	return api.doChaosJSON(ctx, http.MethodPut, scanPath(identity), scopeQuery(scope), &req, nil)
}

// Delete removes a scan.
func (c *ScanClient) Delete(ctx context.Context, scope Scope, identity string) (*http.Response, error) {
	api, err := c.api()
	if err != nil {
		return nil, err
	}
	return api.doChaosJSON(ctx, http.MethodDelete, scanPath(identity), scopeQuery(scope), nil, nil)
}

// GetReport returns the risks this scan produced, with their occurrences.
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

// List returns scans in the scope.
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

// Abort stops a running scan.
func (c *ScanClient) Abort(ctx context.Context, scope Scope, identity string) (*http.Response, error) {
	api, err := c.api()
	if err != nil {
		return nil, err
	}
	return api.doChaosJSON(ctx, http.MethodPost, scanPath(identity)+"/abort", scopeQuery(scope), nil, nil)
}

// Retry requeues a scan.
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
