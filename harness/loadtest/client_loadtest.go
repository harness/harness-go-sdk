package loadtest

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"net/url"

	"github.com/antihax/optional"
)

// LoadTestClient is the load-test API used by Terraform. Validation stays here.
// The HTTP path is the generated operation, and the provider sets the gateway prefix.
type LoadTestClient struct {
	client *APIClient
}

func NewLoadTestClient(client *APIClient) *LoadTestClient {
	return &LoadTestClient{client: client}
}

func (c *LoadTestClient) Create(ctx context.Context, scope Scope, req LoadTestCreateRequest) (*LoadTest, *http.Response, error) {
	if err := validateLoadTestCreate(req); err != nil {
		return nil, nil, err
	}
	api, err := c.api()
	if err != nil {
		return nil, nil, err
	}
	body, err := convert[InternalApiCreateLoadTestRequest](req)
	if err != nil {
		return nil, nil, err
	}
	out, resp, err := api.LoadTestsApi.CreateLoadTest(ctx, body, scope.AccountID, &LoadTestsApiCreateLoadTestOpts{
		OrganizationIdentifier: optString(scope.OrgID),
		ProjectIdentifier:      optString(scope.ProjectID),
		ParentUniqueId:         optString(scope.ParentUniqueID),
	})
	if err != nil {
		return nil, resp, err
	}
	got, convErr := convertPtr[LoadTest](out)
	if convErr != nil {
		return nil, resp, convErr
	}
	return got, resp, nil
}

func (c *LoadTestClient) Get(ctx context.Context, scope Scope, identity string) (*LoadTest, *http.Response, error) {
	if identity == "" {
		return nil, nil, fmt.Errorf("load test identity is required")
	}
	api, err := c.api()
	if err != nil {
		return nil, nil, err
	}
	out, resp, err := api.LoadTestsApi.GetLoadTest(ctx, pathIdentity(identity), scope.AccountID, &LoadTestsApiGetLoadTestOpts{
		OrganizationIdentifier: optString(scope.OrgID),
		ProjectIdentifier:      optString(scope.ProjectID),
		ParentUniqueId:         optString(scope.ParentUniqueID),
	})
	if err != nil {
		return nil, resp, err
	}
	got, convErr := convertPtr[LoadTest](out)
	if convErr != nil {
		return nil, resp, convErr
	}
	return got, resp, nil
}

func (c *LoadTestClient) List(ctx context.Context, req ListLoadTestsRequest) (*LoadTestListResponse, *http.Response, error) {
	api, err := c.api()
	if err != nil {
		return nil, nil, err
	}
	opts := &LoadTestsApiListLoadTestsOpts{
		OrganizationIdentifier: optString(req.Scope.OrgID),
		ProjectIdentifier:      optString(req.Scope.ProjectID),
		ParentUniqueId:         optString(req.Scope.ParentUniqueID),
		Search:                 optString(req.Search),
		SortField:              optString(req.SortField),
		ToolType:               optString(req.ToolType),
		EnvironmentIdentifier:  optString(req.EnvironmentIdentifier),
	}
	if req.Page != 0 {
		opts.Page = optional.NewInt32(int32(req.Page))
	}
	if req.Limit != 0 {
		opts.Limit = optional.NewInt32(int32(req.Limit))
	}
	if req.SortAscending != nil {
		opts.SortAscending = optional.NewBool(*req.SortAscending)
	}
	if len(req.Tags) > 0 {
		opts.Tags = optional.NewInterface(req.Tags)
	}
	out, resp, err := api.LoadTestsApi.ListLoadTests(ctx, req.Scope.AccountID, opts)
	if err != nil {
		return nil, resp, err
	}
	got, convErr := convertPtr[LoadTestListResponse](out)
	if convErr != nil {
		return nil, resp, convErr
	}
	return got, resp, nil
}

func (c *LoadTestClient) GetVariables(ctx context.Context, scope Scope, identity string) (*LoadTestVariablesResponse, *http.Response, error) {
	if identity == "" {
		return nil, nil, fmt.Errorf("load test identity is required")
	}
	api, err := c.api()
	if err != nil {
		return nil, nil, err
	}
	out, resp, err := api.LoadTestsApi.GetVariables(ctx, pathIdentity(identity), scope.AccountID, &LoadTestsApiGetVariablesOpts{
		OrganizationIdentifier: optString(scope.OrgID),
		ProjectIdentifier:      optString(scope.ProjectID),
		ParentUniqueId:         optString(scope.ParentUniqueID),
	})
	if err != nil {
		return nil, resp, err
	}
	got, convErr := convertPtr[LoadTestVariablesResponse](out)
	if convErr != nil {
		return nil, resp, convErr
	}
	return got, resp, nil
}

func (c *LoadTestClient) Update(ctx context.Context, scope Scope, identity string, req LoadTestUpdateRequest) (*LoadTest, *http.Response, error) {
	if identity == "" {
		return nil, nil, fmt.Errorf("load test identity is required")
	}
	api, err := c.api()
	if err != nil {
		return nil, nil, err
	}
	body, err := convert[InternalApiUpdateLoadTestRequest](req)
	if err != nil {
		return nil, nil, err
	}
	out, resp, err := api.LoadTestsApi.UpdateLoadTest(ctx, body, scope.AccountID, pathIdentity(identity), &LoadTestsApiUpdateLoadTestOpts{
		OrganizationIdentifier: optString(scope.OrgID),
		ProjectIdentifier:      optString(scope.ProjectID),
		ParentUniqueId:         optString(scope.ParentUniqueID),
	})
	if err != nil {
		return nil, resp, err
	}
	got, convErr := convertPtr[LoadTest](out)
	if convErr != nil {
		return nil, resp, convErr
	}
	return got, resp, nil
}

func (c *LoadTestClient) Delete(ctx context.Context, scope Scope, identity string) (*LoadTestSuccessResponse, *http.Response, error) {
	if identity == "" {
		return nil, nil, fmt.Errorf("load test identity is required")
	}
	api, err := c.api()
	if err != nil {
		return nil, nil, err
	}
	out, resp, err := api.LoadTestsApi.DeleteLoadTest(ctx, pathIdentity(identity), scope.AccountID, &LoadTestsApiDeleteLoadTestOpts{
		OrganizationIdentifier: optString(scope.OrgID),
		ProjectIdentifier:      optString(scope.ProjectID),
		ParentUniqueId:         optString(scope.ParentUniqueID),
	})
	if err != nil {
		return nil, resp, err
	}
	got, convErr := convertPtr[LoadTestSuccessResponse](out)
	if convErr != nil {
		return nil, resp, convErr
	}
	return got, resp, nil
}

func (c *LoadTestClient) CreateRun(ctx context.Context, scope Scope, loadTestIdentity string, req LoadTestRunCreateRequest) (*LoadTestRun, *http.Response, error) {
	if loadTestIdentity == "" {
		return nil, nil, fmt.Errorf("load test identity is required")
	}
	if req.Identity == "" {
		return nil, nil, fmt.Errorf("run identity is required")
	}
	api, err := c.api()
	if err != nil {
		return nil, nil, err
	}
	body, err := convert[InternalApiCreateLoadTestRunRequest](req)
	if err != nil {
		return nil, nil, err
	}
	out, resp, err := api.RunsApi.CreateLoadTestRun(ctx, body, scope.AccountID, pathIdentity(loadTestIdentity), &RunsApiCreateLoadTestRunOpts{
		OrganizationIdentifier: optString(scope.OrgID),
		ProjectIdentifier:      optString(scope.ProjectID),
		ParentUniqueId:         optString(scope.ParentUniqueID),
	})
	if err != nil {
		return nil, resp, err
	}
	got, convErr := convertPtr[LoadTestRun](out)
	if convErr != nil {
		return nil, resp, convErr
	}
	return got, resp, nil
}

func (c *LoadTestClient) GetRun(ctx context.Context, scope Scope, identity string) (*LoadTestRun, *http.Response, error) {
	if identity == "" {
		return nil, nil, fmt.Errorf("run identity is required")
	}
	api, err := c.api()
	if err != nil {
		return nil, nil, err
	}
	out, resp, err := api.RunsApi.GetLoadTestRun(ctx, pathIdentity(identity), scope.AccountID, &RunsApiGetLoadTestRunOpts{
		OrganizationIdentifier: optString(scope.OrgID),
		ProjectIdentifier:      optString(scope.ProjectID),
		ParentUniqueId:         optString(scope.ParentUniqueID),
	})
	if err != nil {
		return nil, resp, err
	}
	got, convErr := convertPtr[LoadTestRun](out)
	if convErr != nil {
		return nil, resp, convErr
	}
	return got, resp, nil
}

func (c *LoadTestClient) ListRuns(ctx context.Context, loadTestIdentity string, req ListLoadTestRunsRequest) (*LoadTestRunListResponse, *http.Response, error) {
	if loadTestIdentity == "" {
		return nil, nil, fmt.Errorf("load test identity is required")
	}
	api, err := c.api()
	if err != nil {
		return nil, nil, err
	}
	opts := &RunsApiListLoadTestRunsOpts{
		OrganizationIdentifier: optString(req.Scope.OrgID),
		ProjectIdentifier:      optString(req.Scope.ProjectID),
		ParentUniqueId:         optString(req.Scope.ParentUniqueID),
		Search:                 optString(req.Search),
		SortField:              optString(req.SortField),
		Status:                 optString(req.Status),
	}
	if req.Page != 0 {
		opts.Page = optional.NewInt32(int32(req.Page))
	}
	if req.Limit != 0 {
		opts.Limit = optional.NewInt32(int32(req.Limit))
	}
	if req.SortAscending != nil {
		opts.SortAscending = optional.NewBool(*req.SortAscending)
	}
	out, resp, err := api.RunsApi.ListLoadTestRuns(ctx, pathIdentity(loadTestIdentity), req.Scope.AccountID, opts)
	if err != nil {
		return nil, resp, err
	}
	got, convErr := convertPtr[LoadTestRunListResponse](out)
	if convErr != nil {
		return nil, resp, convErr
	}
	return got, resp, nil
}

func (c *LoadTestClient) StopRun(ctx context.Context, scope Scope, loadTestIdentity, identity string) (*LoadTestSuccessResponse, *http.Response, error) {
	if identity == "" {
		return nil, nil, fmt.Errorf("run identity is required")
	}
	if loadTestIdentity == "" {
		return nil, nil, fmt.Errorf("load test identity is required")
	}
	api, err := c.api()
	if err != nil {
		return nil, nil, err
	}
	out, resp, err := api.RunsApi.StopLoadTestRun(ctx, pathIdentity(identity), scope.AccountID, &RunsApiStopLoadTestRunOpts{
		LoadTestIdentity:       optString(loadTestIdentity),
		OrganizationIdentifier: optString(scope.OrgID),
		ProjectIdentifier:      optString(scope.ProjectID),
		ParentUniqueId:         optString(scope.ParentUniqueID),
	})
	if err != nil {
		return nil, resp, err
	}
	got, convErr := convertPtr[LoadTestSuccessResponse](out)
	if convErr != nil {
		return nil, resp, convErr
	}
	return got, resp, nil
}

func (c *LoadTestClient) api() (*APIClient, error) {
	if c == nil || c.client == nil || c.client.cfg == nil || c.client.cfg.BasePath == "" {
		return nil, fmt.Errorf("load test API client is not configured")
	}
	return c.client, nil
}

func validateLoadTestCreate(req LoadTestCreateRequest) error {
	if req.Identity == "" || req.Name == "" || req.ToolType == "" {
		return fmt.Errorf("identity, name, and toolType are required")
	}
	switch req.ToolType {
	case LoadTestToolLocust, LoadTestToolK6:
	case LoadTestToolJMeter:
		return fmt.Errorf("toolType JMeter cannot be created with this client: Load Test Manager rewrites the saved plan. Get still reads an existing JMeter test")
	default:
		return fmt.Errorf("toolType %q is not supported for create; use Locust or K6", req.ToolType)
	}
	if req.ManagedInfra {
		if req.InfraIdentifier != "" || req.EnvironmentIdentifier != "" {
			return fmt.Errorf("infraIdentifier and environmentIdentifier must be omitted when managedInfra is true")
		}
		return nil
	}
	if req.InfraIdentifier == "" || req.EnvironmentIdentifier == "" {
		return fmt.Errorf("infraIdentifier and environmentIdentifier are required unless managedInfra is true")
	}
	return nil
}

func pathIdentity(identity string) string {
	return url.PathEscape(identity)
}

func optString(value string) optional.String {
	if value == "" {
		return optional.String{}
	}
	return optional.NewString(value)
}

func convert[T any](in any) (T, error) {
	var out T
	raw, err := json.Marshal(in)
	if err != nil {
		return out, err
	}
	if err := json.Unmarshal(raw, &out); err != nil {
		return out, err
	}
	return out, nil
}

func convertPtr[T any](in any) (*T, error) {
	got, err := convert[T](in)
	if err != nil {
		return nil, err
	}
	return &got, nil
}
