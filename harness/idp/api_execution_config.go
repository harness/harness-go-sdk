package idp

import (
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"strings"
)

const (
	execCfgGetDefaultConfigUrl  = "/idp/execution-config/get-default-config"
	execCfgGetCustomerConfigUrl = "/idp/execution-config/get-customer-config"
	execCfgUpdateConfigUrl      = "/idp/execution-config/update-config"
	execCfgResetConfigUrl       = "/idp/execution-config/reset-config"
)

type ExecutionConfigApiService service

// ExecutionConfigData is a map of image field names to string values. It uses a
// custom JSON unmarshaler so that non-string API response values (e.g. booleans)
// are silently coerced to their string representation rather than causing an
// unmarshal error.
type ExecutionConfigData map[string]string

// ExecutionConfigResponse is the API response wrapper for the IDP execution config endpoint.
// Data is a dynamic map of image field names to image tags, e.g. {"registerCatalog": "harness/registercatalog:1.9.0"}.
type ExecutionConfigResponse struct {
	Status        string              `json:"status,omitempty"`
	Data          ExecutionConfigData `json:"data,omitempty"`
	MetaData      interface{}         `json:"metaData,omitempty"`
	CorrelationId string              `json:"correlationId,omitempty"`
}

// ExecutionConfigUpdate represents a single image field update for the update-config endpoint
type ExecutionConfigUpdate struct {
	Field string `json:"field"`
	Value string `json:"value,omitempty"`
}

func (d *ExecutionConfigData) UnmarshalJSON(b []byte) error {
	// API endpoints like reset-config may return "data": true instead of an object.
	// Treat any non-object value as empty data rather than an error.
	if len(b) == 0 || b[0] != '{' {
		*d = make(ExecutionConfigData)
		return nil
	}
	var raw map[string]json.RawMessage
	if err := json.Unmarshal(b, &raw); err != nil {
		return err
	}
	*d = make(ExecutionConfigData, len(raw))
	for k, v := range raw {
		var s string
		if err := json.Unmarshal(v, &s); err == nil {
			(*d)[k] = s
		} else {
			(*d)[k] = fmt.Sprintf("%s", v)
		}
	}
	return nil
}

func (a *ExecutionConfigApiService) GetDefaultConfig(ctx context.Context,
	infraType string) (ExecutionConfigResponse, error) {
	return runIdpExecutionConfigRequest(a.client, ctx, http.MethodGet,
		execCfgGetDefaultConfigUrl, execCfgQueryParams(infraType, nil), []string{}, nil)
}

func (a *ExecutionConfigApiService) GetCustomerConfig(ctx context.Context,
	infraType string, overridesOnly bool) (ExecutionConfigResponse, error) {
	return runIdpExecutionConfigRequest(a.client, ctx, http.MethodGet,
		execCfgGetCustomerConfigUrl, execCfgQueryParams(infraType, &overridesOnly), []string{}, nil)
}

func (a *ExecutionConfigApiService) UpdateConfig(ctx context.Context,
	infraType string, body []ExecutionConfigUpdate) (ExecutionConfigResponse, error) {
	return runIdpExecutionConfigRequest(a.client, ctx, http.MethodPost,
		execCfgUpdateConfigUrl, execCfgQueryParams(infraType, nil), []string{"application/json"}, &body)
}

func (a *ExecutionConfigApiService) ResetConfig(ctx context.Context,
	infraType string, body []ExecutionConfigUpdate) (ExecutionConfigResponse, error) {
	return runIdpExecutionConfigRequest(a.client, ctx, http.MethodPost,
		execCfgResetConfigUrl, execCfgQueryParams(infraType, nil), []string{"application/json"}, &body)
}

// idpGatewayBase strips the /v1 suffix from the IDP client BasePath so that
// execution-config endpoints (at /idp/...) are addressed correctly.
func idpGatewayBase(client *APIClient) string {
	return strings.TrimSuffix(client.cfg.BasePath, "/v1")
}

func runIdpExecutionConfigRequest(client *APIClient, ctx context.Context,
	httpMethod, path string, extraQueryParams url.Values,
	contentTypes []string, body interface{}) (ExecutionConfigResponse, error) {

	var (
		localVarPostBody  interface{}
		localVarFileName  string
		localVarFileBytes []byte
		localVarResult    ExecutionConfigResponse
	)

	localVarPath := idpGatewayBase(client) + path

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	localVarQueryParams.Add("accountIdentifier", parameterToString(client.AccountId, ""))
	for k, vs := range extraQueryParams {
		for _, v := range vs {
			localVarQueryParams.Add(k, v)
		}
	}

	localVarHttpContentType := selectHeaderContentType(contentTypes)
	if localVarHttpContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHttpContentType
	}

	localVarHttpHeaderAccept := selectHeaderAccept([]string{"application/json"})
	if localVarHttpHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHttpHeaderAccept
	}

	if ctx != nil {
		if auth, ok := ctx.Value(ContextAPIKey).(APIKey); ok {
			var key string
			if auth.Prefix != "" {
				key = auth.Prefix + " " + auth.Key
			} else {
				key = auth.Key
			}
			localVarHeaderParams["x-api-key"] = key
		}
	}

	if body != nil {
		localVarPostBody = body
	}

	r, err := client.prepareRequest(ctx, localVarPath, httpMethod, localVarPostBody, localVarHeaderParams,
		localVarQueryParams, localVarFormParams, localVarFileName, localVarFileBytes)
	if err != nil {
		return localVarResult, err
	}

	httpResp, err := client.callAPI(r)
	if err != nil {
		return localVarResult, err
	}
	if httpResp == nil {
		return localVarResult, nil
	}
	defer func() {
		_ = httpResp.Body.Close()
	}()

	respBody, err := io.ReadAll(httpResp.Body)
	if err != nil {
		return localVarResult, err
	}

	if httpResp.StatusCode >= 300 {
		return localVarResult, GenericSwaggerError{
			body:  respBody,
			error: httpResp.Status,
		}
	}

	err = client.decode(&localVarResult, respBody, httpResp.Header.Get("Content-Type"))
	return localVarResult, err
}

func execCfgQueryParams(infraType string, overridesOnly *bool) url.Values {
	params := url.Values{}
	params.Add("infra", parameterToString(infraType, ""))
	if overridesOnly != nil {
		params.Add("overridesOnly", parameterToString(*overridesOnly, ""))
	}
	return params
}
