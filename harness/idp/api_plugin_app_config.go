package idp

import (
	"context"
	"fmt"
	"io"
	"net/http"
	"net/url"
)

const (
	pluginAppConfigSaveUrl   = "/v1/app-config"
	pluginInfoGetUrl         = "/v1/plugins-info/%s"
	pluginInfoListUrl        = "/v1/plugins-info"
	pluginToggleUrl          = "/v1/plugin-toggle/%s"
)

type PluginAppConfigApiService service

// SaveOrUpdate creates or updates a plugin app configuration.
func (a *PluginAppConfigApiService) SaveOrUpdate(ctx context.Context,
	body PluginAppConfigRequest) (PluginAppConfigResponse, error) {

	return runPluginAppConfigRequest[PluginAppConfigResponse](
		a.client, ctx, http.MethodPost,
		pluginAppConfigSaveUrl, nil,
		[]string{"application/json"}, &body)
}

// GetPluginInfo retrieves the plugin details and configuration for a given plugin ID.
func (a *PluginAppConfigApiService) GetPluginInfo(ctx context.Context,
	pluginId string) (PluginInfoResponse, error) {

	path := fmt.Sprintf(pluginInfoGetUrl, url.PathEscape(pluginId))
	return runPluginAppConfigRequest[PluginInfoResponse](
		a.client, ctx, http.MethodGet,
		path, nil,
		nil, nil)
}

// ListPlugins retrieves all available plugins for the account.
func (a *PluginAppConfigApiService) ListPlugins(ctx context.Context) ([]PluginListItem, error) {
	return runPluginAppConfigRequest[[]PluginListItem](
		a.client, ctx, http.MethodGet,
		pluginInfoListUrl, nil,
		nil, nil)
}

// Toggle enables or disables a plugin for the account.
func (a *PluginAppConfigApiService) Toggle(ctx context.Context,
	pluginId string, enabled bool) (PluginAppConfigResponse, error) {

	path := fmt.Sprintf(pluginToggleUrl, url.PathEscape(pluginId))
	queryParams := url.Values{}
	queryParams.Add("enabled", parameterToString(enabled, ""))
	return runPluginAppConfigRequest[PluginAppConfigResponse](
		a.client, ctx, http.MethodPost,
		path, queryParams,
		nil, nil)
}

func runPluginAppConfigRequest[T any](client *APIClient, ctx context.Context,
	httpMethod, path string, extraQueryParams url.Values,
	contentTypes []string, body any) (T, error) {

	var (
		localVarPostBody  any
		localVarFileName  string
		localVarFileBytes []byte
		localVarResult    T
	)

	localVarPath := idpGatewayBase(client) + path

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	localVarHeaderParams["Harness-Account"] = parameterToString(client.AccountId, "")

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
