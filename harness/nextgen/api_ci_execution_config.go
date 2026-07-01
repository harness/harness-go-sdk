package nextgen

import (
	"context"
	"net/http"
)

type CiExecutionConfigApiService service

const (
	ciExCfgGetDefaultConfigUrl  = "/ci/execution-config/get-default-config"
	ciExCfgGetCustomerConfigUrl = "/ci/execution-config/get-customer-config"
	ciExCfgUpdateConfigUrl      = "/ci/execution-config/update-config"
	ciExCfgResetConfigUrl       = "/ci/execution-config/reset-config"
)

func (a *CiExecutionConfigApiService) GetDefaultConfig(ctx context.Context,
	infraType string) (ExecutionConfigResponse, error) {
	return runExecutionConfigRequest(a.client, ctx, http.MethodGet,
		ciExCfgGetDefaultConfigUrl, executionConfigQueryParams(infraType, nil), []string{}, nil)
}

func (a *CiExecutionConfigApiService) GetCustomerConfig(ctx context.Context,
	infraType string, overridesOnly bool) (ExecutionConfigResponse, error) {
	return runExecutionConfigRequest(a.client, ctx, http.MethodGet,
		ciExCfgGetCustomerConfigUrl, executionConfigQueryParams(infraType, &overridesOnly), []string{}, nil)
}

func (a *CiExecutionConfigApiService) UpdateConfig(ctx context.Context,
	infraType string, body []ExecutionConfigUpdate) (ExecutionConfigResponse, error) {
	return runExecutionConfigRequest(a.client, ctx, http.MethodPost,
		ciExCfgUpdateConfigUrl, executionConfigQueryParams(infraType, nil),
		[]string{"application/json"}, &body)
}

func (a *CiExecutionConfigApiService) ResetConfig(ctx context.Context,
	infraType string, body []ExecutionConfigUpdate) (ExecutionConfigResponse, error) {
	return runExecutionConfigRequest(a.client, ctx, http.MethodPost,
		ciExCfgResetConfigUrl, executionConfigQueryParams(infraType, nil),
		[]string{"application/json"}, &body)
}
