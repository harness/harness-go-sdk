package nextgen

import (
	"context"
	"net/http"
)

type IacmExecutionConfigApiService service

const (
	iacmExCfgGetDefaultConfigUrl  = "/iacm-manager/execution-config/get-default-config"
	iacmExCfgGetCustomerConfigUrl = "/iacm-manager/execution-config/get-customer-config"
	iacmExCfgUpdateConfigUrl      = "/iacm-manager/execution-config/update-config"
	iacmExCfgResetConfigUrl       = "/iacm-manager/execution-config/reset-config"
)

func (a *IacmExecutionConfigApiService) GetDefaultConfig(ctx context.Context,
	infraType string) (ExecutionConfigResponse, error) {
	return runExecutionConfigRequest(a.client, ctx, http.MethodGet,
		iacmExCfgGetDefaultConfigUrl, executionConfigQueryParams(infraType, nil), []string{}, nil)
}

func (a *IacmExecutionConfigApiService) GetCustomerConfig(ctx context.Context,
	infraType string, overridesOnly bool) (ExecutionConfigResponse, error) {
	return runExecutionConfigRequest(a.client, ctx, http.MethodGet,
		iacmExCfgGetCustomerConfigUrl, executionConfigQueryParams(infraType, &overridesOnly), []string{}, nil)
}

func (a *IacmExecutionConfigApiService) UpdateConfig(ctx context.Context,
	infraType string, body []ExecutionConfigUpdate) (ExecutionConfigResponse, error) {
	return runExecutionConfigRequest(a.client, ctx, http.MethodPost,
		iacmExCfgUpdateConfigUrl, executionConfigQueryParams(infraType, nil), []string{"application/json"}, &body)
}

func (a *IacmExecutionConfigApiService) ResetConfig(ctx context.Context,
	infraType string, body []ExecutionConfigUpdate) (ExecutionConfigResponse, error) {
	return runExecutionConfigRequest(a.client, ctx, http.MethodPost,
		iacmExCfgResetConfigUrl, executionConfigQueryParams(infraType, nil), []string{"application/json"}, &body)
}
