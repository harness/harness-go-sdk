/*
 * Harness NextGen Software Delivery Platform API Reference
 *
 * This file contains hand-written extensions to the generated AuthenticationSettingsApi.
 * It covers endpoints that are not yet present in the published OpenAPI spec, and
 * re-implements generated operations whose spec declares a Content-Type of "*\/*",
 * which the generated client cannot serialize a JSON body for ("Invalid body type").
 * When the spec is regenerated with correct content types and includes these
 * operations, this file should be dropped in favor of the generated code.
 */
package nextgen

import (
	"context"
	"io/ioutil"
	"net/http"
	"net/url"
)

// executeAuthenticationSettingsWriteRequest performs an HTTP call with a JSON body
// against an /ng/api/authentication-settings endpoint and decodes the response
// into out.
func (a *AuthenticationSettingsApiService) executeAuthenticationSettingsWriteRequest(ctx context.Context, method string, path string, body interface{}, accountId string, out interface{}) (*http.Response, error) {
	localVarPath := a.client.cfg.BasePath + path

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	localVarQueryParams.Add("accountIdentifier", parameterToString(accountId, ""))

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
	r, err := a.client.prepareRequest(ctx, localVarPath, method, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, "", nil)
	if err != nil {
		return nil, err
	}

	localVarHttpResponse, err := a.client.callAPI(r)
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
		err = a.client.decode(&v, localVarBody, localVarHttpResponse.Header.Get("Content-Type"))
		if err != nil {
			newErr.error = err.Error()
			newErr.model = v
			return localVarHttpResponse, newErr
		}
	}
	return localVarHttpResponse, newErr
}

/*
AuthenticationSettingsApiService Set two factor auth at account level by accountIdentifier
Re-implements the generated SetTwoFactorAuthAtAccountLevel, which cannot serialize
its body because the spec declares a Content-Type of "*\/*".
  - @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
  - @param body Boolean that specify whether or not to override two factor enabled setting
  - @param accountIdentifier Account Identifier for the Entity.

@return RestResponseBoolean
*/
func (a *AuthenticationSettingsApiService) SetTwoFactorAuth(ctx context.Context, body TwoFactorAdminOverrideSettings, accountIdentifier string) (RestResponseBoolean, *http.Response, error) {
	var localVarReturnValue RestResponseBoolean
	httpResp, err := a.executeAuthenticationSettingsWriteRequest(ctx, http.MethodPut, "/ng/api/authentication-settings/two-factor-admin-override-settings", body, accountIdentifier, &localVarReturnValue)
	return localVarReturnValue, httpResp, err
}

/*
AuthenticationSettingsApiService Update the whitelisted domains of an account
Re-implements the generated UpdateWhitelistedDomains, which cannot serialize
its body because the spec declares a Content-Type of "*\/*".
  - @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
  - @param body The full list of whitelisted domains. IPs and IP ranges are rejected by the backend; only domain names are allowed.
  - @param accountIdentifier Account Identifier for the Entity.

@return RestResponseBoolean
*/
func (a *AuthenticationSettingsApiService) SetWhitelistedDomains(ctx context.Context, body []string, accountIdentifier string) (RestResponseBoolean, *http.Response, error) {
	var localVarReturnValue RestResponseBoolean
	httpResp, err := a.executeAuthenticationSettingsWriteRequest(ctx, http.MethodPut, "/ng/api/authentication-settings/whitelisted-domains", body, accountIdentifier, &localVarReturnValue)
	return localVarReturnValue, httpResp, err
}

/*
AuthenticationSettingsApiService Update the allowed OAuth providers of an account
Re-implements the generated UpdateOauthProviders, which cannot serialize
its body because the spec declares a Content-Type of "*\/*".
  - @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
  - @param body The OAuth settings, including the allowed providers.
  - @param accountIdentifier Account Identifier for the Entity.

@return RestResponseBoolean
*/
func (a *AuthenticationSettingsApiService) SetOauthProviders(ctx context.Context, body OAuthSettings, accountIdentifier string) (RestResponseBoolean, *http.Response, error) {
	var localVarReturnValue RestResponseBoolean
	httpResp, err := a.executeAuthenticationSettingsWriteRequest(ctx, http.MethodPut, "/ng/api/authentication-settings/oauth/update-providers", body, accountIdentifier, &localVarReturnValue)
	return localVarReturnValue, httpResp, err
}

/*
AuthenticationSettingsApiService Set session timeout at account level by accountIdentifier
Not present in the published OpenAPI spec.
  - @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
  - @param body Information about the session timeout for all users of this account in minutes.
  - @param accountIdentifier Account Identifier for the Entity.

@return RestResponseBoolean
*/
func (a *AuthenticationSettingsApiService) SetSessionTimeoutAtAccountLevel(ctx context.Context, body SessionTimeoutSettings, accountIdentifier string) (RestResponseBoolean, *http.Response, error) {
	var localVarReturnValue RestResponseBoolean
	httpResp, err := a.executeAuthenticationSettingsWriteRequest(ctx, http.MethodPut, "/ng/api/authentication-settings/session-timeout-account-level", body, accountIdentifier, &localVarReturnValue)
	return localVarReturnValue, httpResp, err
}

/*
AuthenticationSettingsApiService Enable/disable public access at account level by accountIdentifier
Not present in the published OpenAPI spec.
  - @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
  - @param body Boolean that specifies whether or not public access is enabled for the account.
  - @param accountIdentifier Account Identifier for the Entity.

@return RestResponseBoolean
*/
func (a *AuthenticationSettingsApiService) SetPublicAccess(ctx context.Context, body bool, accountIdentifier string) (RestResponseBoolean, *http.Response, error) {
	var localVarReturnValue RestResponseBoolean
	httpResp, err := a.executeAuthenticationSettingsWriteRequest(ctx, http.MethodPut, "/ng/api/authentication-settings/public-access", body, accountIdentifier, &localVarReturnValue)
	return localVarReturnValue, httpResp, err
}
