/*
 * Harness NextGen Software Delivery Platform API Reference
 *
 * This file contains hand-written extensions to the generated AuthenticationSettingsApi
 * for the login settings operations of AuthenticationSettingsResource
 * (/ng/api/authentication-settings). These operations are marked hidden in the
 * OpenAPI spec, so no generated code exists for them. When the spec is regenerated
 * and includes these operations, this file should be dropped in favor of the
 * generated code.
 */
package nextgen

import (
	"context"
	"io/ioutil"
	"net/http"
	"net/url"
)

// executeUpdateLoginSettingsRequest performs the PUT call against
// /ng/api/authentication-settings/login-settings/{loginSettingsId} and decodes
// the response into out.
func (a *AuthenticationSettingsApiService) executeUpdateLoginSettingsRequest(ctx context.Context, path string, body interface{}, accountId string, out interface{}) (*http.Response, error) {
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
	r, err := a.client.prepareRequest(ctx, localVarPath, http.MethodPut, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, "", nil)
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
			return localVarHttpResponse, newErr
		}
		newErr.model = v
		return localVarHttpResponse, newErr
	}
	return localVarHttpResponse, newErr
}

/*
AuthenticationSettingsApiService Get the login settings of an account
Returns the full login settings: password strength policy, user lockout policy
and password expiration policy, along with the server-generated uuid needed for
updates. The login settings document is read from the USER_PASSWORD entry of the
authentication settings response; no separate login settings endpoint is exposed.
  - @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
  - @param accountId Account Identifier for the Entity.

@return RestResponseLoginSettings
*/
func (a *AuthenticationSettingsApiService) GetLoginSettings(ctx context.Context, accountId string) (RestResponseLoginSettings, *http.Response, error) {
	var localVarReturnValue RestResponseLoginSettings

	// Read via the v2 API: the v1 list endpoint is rejected with a 400 on accounts
	// where the PL_ENABLE_MULTIPLE_IDP_SUPPORT feature flag is enabled, while v2
	// works regardless of the flag and returns the same response shape.
	authSettings, httpResp, err := a.GetAuthenticationSettingsV2(ctx, accountId)
	if err != nil {
		return localVarReturnValue, httpResp, err
	}

	if authSettings.Resource != nil {
		for _, settings := range authSettings.Resource.NgAuthSettings {
			if settings.SettingsType == "USER_PASSWORD" && settings.LoginSettings != nil {
				localVarReturnValue.Resource = settings.LoginSettings
				break
			}
		}
	}

	return localVarReturnValue, httpResp, nil
}

/*
AuthenticationSettingsApiService Update the login settings of an account
The API requires values for all fields of the login settings, not just the ones
being updated; read the current settings first and send the merged document.
  - @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
  - @param body The full updated login settings.
  - @param loginSettingsId Identifier of the login settings, as returned by GetLoginSettings.
  - @param accountId Account Identifier for the Entity.

@return RestResponseLoginSettings
*/
func (a *AuthenticationSettingsApiService) UpdateLoginSettings(ctx context.Context, body LoginSettings, loginSettingsId string, accountId string) (RestResponseLoginSettings, *http.Response, error) {
	var localVarReturnValue RestResponseLoginSettings
	httpResp, err := a.executeUpdateLoginSettingsRequest(ctx, "/ng/api/authentication-settings/login-settings/"+loginSettingsId, body, accountId, &localVarReturnValue)
	return localVarReturnValue, httpResp, err
}
