/*
 * Harness NextGen Software Delivery Platform API Reference
 *
 * This file contains hand-written extensions to the generated AuthenticationSettingsApi
 * for endpoints that are not yet present in the published OpenAPI spec. When the spec
 * is regenerated and includes these operations, this file should be dropped in favor
 * of the generated code.
 */
package nextgen

import (
	"context"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
	"strings"
)

// executeLdapSettingsRequest performs an HTTP call against the
// /ng/api/authentication-settings/ldap/settings endpoint and decodes the response
// into out. body may be nil for requests without a payload.
func (a *AuthenticationSettingsApiService) executeLdapSettingsRequest(ctx context.Context, method string, body interface{}, accountId string, out interface{}) (*http.Response, error) {
	localVarPath := a.client.cfg.BasePath + "/ng/api/authentication-settings/ldap/settings"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	localVarQueryParams.Add("accountId", parameterToString(accountId, ""))

	// to determine the Content-Type header
	localVarHttpContentTypes := []string{"application/json"}

	// set Content-Type header
	localVarHttpContentType := selectHeaderContentType(localVarHttpContentTypes)
	if localVarHttpContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHttpContentType
	}

	// to determine the Accept header
	localVarHttpHeaderAccepts := []string{"application/json", "application/yaml"}

	// set Accept header
	localVarHttpHeaderAccept := selectHeaderAccept(localVarHttpHeaderAccepts)
	if localVarHttpHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHttpHeaderAccept
	}
	// body params
	var localVarPostBody interface{}
	if body != nil {
		localVarPostBody = &body
	}
	if ctx != nil {
		// API Key Authentication
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
	r, err := a.client.prepareRequest(ctx, localVarPath, strings.ToUpper(method), localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, "", nil)
	if err != nil {
		return nil, err
	}

	var localVarHttpResponse *http.Response
	if strings.EqualFold(method, http.MethodGet) {
		// The backend answers 500 (not 404) when the account has no LDAP
		// configuration. Skip the retrying client for reads so this expected
		// response does not trigger a long retry storm; callers map the 500
		// to "not found". Note: StandardClient() still retries (it only adapts
		// the interface), so use the inner raw *http.Client directly.
		localVarHttpResponse, err = a.client.cfg.HTTPClient.HTTPClient.Do(r.Request)
	} else {
		localVarHttpResponse, err = a.client.callAPI(r)
	}
	if err != nil || localVarHttpResponse == nil {
		return localVarHttpResponse, err
	}

	localVarBody, err := ioutil.ReadAll(localVarHttpResponse.Body)
	localVarHttpResponse.Body.Close()
	if err != nil {
		return localVarHttpResponse, err
	}

	if localVarHttpResponse.StatusCode < 300 {
		// If we succeed, return the data, otherwise pass on to decode error.
		if out != nil {
			if err := a.client.decode(out, localVarBody, localVarHttpResponse.Header.Get("Content-Type")); err != nil {
				return localVarHttpResponse, err
			}
		}
		return localVarHttpResponse, nil
	}

	newErr := GenericSwaggerError{
		body:  localVarBody,
		error: localVarHttpResponse.Status,
	}
	if localVarHttpResponse.StatusCode == 400 || localVarHttpResponse.StatusCode == 404 {
		var v Failure
		err = a.client.decode(&v, localVarBody, localVarHttpResponse.Header.Get("Content-Type"))
		if err != nil {
			newErr.error = err.Error()
			return localVarHttpResponse, newErr
		}
		newErr.model = v
		return localVarHttpResponse, newErr
	}
	if localVarHttpResponse.StatusCode == 500 {
		var v ModelError
		if err := a.client.decode(&v, localVarBody, localVarHttpResponse.Header.Get("Content-Type")); err != nil {
			newErr.error = err.Error()
			newErr.model = v
			return localVarHttpResponse, newErr
		}
	}
	return localVarHttpResponse, newErr
}

/*
AuthenticationSettingsApiService Return configured Ldap settings for the account
  - @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
  - @param accountId Account Identifier for the Entity.

@return RestResponseLdapSettings
*/
func (a *AuthenticationSettingsApiService) GetLdapSettings(ctx context.Context, accountId string) (RestResponseLdapSettings, *http.Response, error) {
	var localVarReturnValue RestResponseLdapSettings
	httpResp, err := a.executeLdapSettingsRequest(ctx, http.MethodGet, nil, accountId, &localVarReturnValue)
	return localVarReturnValue, httpResp, err
}

/*
AuthenticationSettingsApiService Create Ldap settings - with user queries, group queries
  - @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
  - @param body LDAP settings to create. Values for connection settings are needed, user and group settings can also be provided.
  - @param accountId Account Identifier for the Entity.

@return RestResponseLdapSettings
*/
func (a *AuthenticationSettingsApiService) CreateLdapSettings(ctx context.Context, body LdapSettings, accountId string) (RestResponseLdapSettings, *http.Response, error) {
	var localVarReturnValue RestResponseLdapSettings
	httpResp, err := a.executeLdapSettingsRequest(ctx, http.MethodPost, body, accountId, &localVarReturnValue)
	return localVarReturnValue, httpResp, err
}

/*
AuthenticationSettingsApiService Update Ldap settings - user queries, group queries
  - @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
  - @param body The updated LDAP settings. Values for all fields are needed, not just the fields being updated.
  - @param accountId Account Identifier for the Entity.

@return RestResponseLdapSettings
*/
func (a *AuthenticationSettingsApiService) UpdateLdapSettings(ctx context.Context, body LdapSettings, accountId string) (RestResponseLdapSettings, *http.Response, error) {
	var localVarReturnValue RestResponseLdapSettings
	httpResp, err := a.executeLdapSettingsRequest(ctx, http.MethodPut, body, accountId, &localVarReturnValue)
	return localVarReturnValue, httpResp, err
}

/*
AuthenticationSettingsApiService Delete Ldap settings
  - @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
  - @param accountId Account Identifier for the Entity.

@return RestResponseBoolean
*/
func (a *AuthenticationSettingsApiService) DeleteLdapSettings(ctx context.Context, accountId string) (RestResponseBoolean, *http.Response, error) {
	var localVarReturnValue RestResponseBoolean
	httpResp, err := a.executeLdapSettingsRequest(ctx, http.MethodDelete, nil, accountId, &localVarReturnValue)
	if err == nil && !localVarReturnValue.Resource {
		err = fmt.Errorf("failed to delete LDAP settings for account %s", accountId)
	}
	return localVarReturnValue, httpResp, err
}
