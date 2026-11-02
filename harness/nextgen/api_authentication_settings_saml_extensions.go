/*
 * Harness NextGen Software Delivery Platform API Reference
 *
 * This file contains hand-written extensions to the generated AuthenticationSettingsApi
 * for endpoints that are not present (or are generated incorrectly) in the published
 * OpenAPI spec:
 *   - the generated UploadSamlMetaData/UpdateSamlMetaData functions send the metadata
 *     file as a URL-encoded form value instead of a multipart file part,
 *   - the multi-config endpoints keyed by samlSSOId are not generated at all.
 * When the spec is regenerated and includes these operations, this file should be
 * dropped in favor of the generated code.
 */
package nextgen

import (
	"context"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
	"strconv"
	"strings"
)

// SamlSettingsRequest carries the optional form fields of the SAML metadata
// upload/update endpoints. Pointer fields are only sent when non-nil.
type SamlSettingsRequest struct {
	// Display Name of the SAML configuration.
	DisplayName string
	// Group membership attribute in the SAML assertion.
	GroupMembershipAttr string
	// Specify whether or not to enable authorization.
	AuthorizationEnabled *bool
	// Logout URL of the SAML identity provider.
	LogoutUrl string
	// SAML metadata entity identifier.
	EntityIdentifier string
	// SAML provider type.
	SamlProviderType string
	// Optional SAML clientId for Azure SSO.
	ClientId string
	// Optional Harness secret reference string for the Azure SSO client secret.
	ClientSecret string
	// Friendly name of the app on the SAML SSO provider end in Harness.
	FriendlySamlName string
	// Enable Just in time user provisioning.
	JitEnabled *bool
	// Optional key to match in the SAML assertion for Just in time user provisioning.
	JitValidationKey string
	// Optional value to match in the SAML assertion for Just in time user provisioning.
	JitValidationValue string
}

// executeSamlRequest performs an HTTP call against the SAML endpoints of
// /ng/api/authentication-settings and decodes the response into out.
// When metadata is non-nil, the request is sent as multipart/form-data with the
// metadata as the "file" part and the request fields as form fields.
func (a *AuthenticationSettingsApiService) executeSamlRequest(ctx context.Context, method string, path string, accountQueryParam string, accountId string, metadata []byte, req *SamlSettingsRequest, out interface{}) (*http.Response, error) {
	localVarPath := a.client.cfg.BasePath + path

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	localVarQueryParams.Add(accountQueryParam, parameterToString(accountId, ""))

	var localVarFileName string
	var localVarFileBytes []byte

	if metadata != nil {
		localVarHttpContentType := selectHeaderContentType([]string{"multipart/form-data"})
		if localVarHttpContentType != "" {
			localVarHeaderParams["Content-Type"] = localVarHttpContentType
		}
		localVarFileName = "metadata.xml"
		localVarFileBytes = metadata

		if req != nil {
			if req.DisplayName != "" {
				localVarFormParams.Add("displayName", req.DisplayName)
			}
			if req.GroupMembershipAttr != "" {
				localVarFormParams.Add("groupMembershipAttr", req.GroupMembershipAttr)
			}
			if req.AuthorizationEnabled != nil {
				localVarFormParams.Add("authorizationEnabled", strconv.FormatBool(*req.AuthorizationEnabled))
			}
			if req.LogoutUrl != "" {
				localVarFormParams.Add("logoutUrl", req.LogoutUrl)
			}
			if req.EntityIdentifier != "" {
				localVarFormParams.Add("entityIdentifier", req.EntityIdentifier)
			}
			if req.SamlProviderType != "" {
				localVarFormParams.Add("samlProviderType", req.SamlProviderType)
			}
			if req.ClientId != "" {
				localVarFormParams.Add("clientId", req.ClientId)
			}
			if req.ClientSecret != "" {
				localVarFormParams.Add("clientSecret", req.ClientSecret)
			}
			if req.FriendlySamlName != "" {
				localVarFormParams.Add("friendlySamlName", req.FriendlySamlName)
			}
			if req.JitEnabled != nil {
				localVarFormParams.Add("jitEnabled", strconv.FormatBool(*req.JitEnabled))
			}
			if req.JitValidationKey != "" {
				localVarFormParams.Add("jitValidationKey", req.JitValidationKey)
			}
			if req.JitValidationValue != "" {
				localVarFormParams.Add("jitValidationValue", req.JitValidationValue)
			}
		}
	}

	// to determine the Accept header
	localVarHttpHeaderAccepts := []string{"application/json", "application/yaml"}

	// set Accept header
	localVarHttpHeaderAccept := selectHeaderAccept(localVarHttpHeaderAccepts)
	if localVarHttpHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHttpHeaderAccept
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
	r, err := a.client.prepareRequest(ctx, localVarPath, strings.ToUpper(method), nil, localVarHeaderParams, localVarQueryParams, localVarFormParams, localVarFileName, localVarFileBytes)
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
AuthenticationSettingsApiService Get authentication settings version 2 for an account
  - @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
  - @param accountIdentifier Account Identifier for the Entity.

@return RestResponseAuthenticationSettingsV2Response
*/
func (a *AuthenticationSettingsApiService) GetAuthenticationSettingsV2(ctx context.Context, accountIdentifier string) (RestResponseAuthenticationSettingsV2Response, *http.Response, error) {
	var localVarReturnValue RestResponseAuthenticationSettingsV2Response
	httpResp, err := a.executeSamlRequest(ctx, http.MethodGet, "/ng/api/authentication-settings/v2", "accountIdentifier", accountIdentifier, nil, nil, &localVarReturnValue)
	return localVarReturnValue, httpResp, err
}

/*
AuthenticationSettingsApiService Create a SAML SSO configuration for an account
  - @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
  - @param accountId Account Identifier for the Entity.
  - @param metadata SAML IdP metadata XML, sent as the multipart "file" part.
  - @param req Optional form fields of the SAML configuration.

@return RestResponseSamlSettingsDto
*/
func (a *AuthenticationSettingsApiService) CreateSamlSettings(ctx context.Context, accountId string, metadata []byte, req *SamlSettingsRequest) (RestResponseSamlSettingsDto, *http.Response, error) {
	var localVarReturnValue RestResponseSamlSettingsDto
	if len(metadata) == 0 {
		return localVarReturnValue, nil, fmt.Errorf("metadata must not be empty")
	}
	httpResp, err := a.executeSamlRequest(ctx, http.MethodPost, "/ng/api/authentication-settings/saml-metadata", "accountId", accountId, metadata, req, &localVarReturnValue)
	return localVarReturnValue, httpResp, err
}

/*
AuthenticationSettingsApiService Update the SAML SSO configuration with the given samlSSOId
  - @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
  - @param accountId Account Identifier for the Entity.
  - @param samlSSOId Identifier of the SAML SSO configuration to update.
  - @param metadata SAML IdP metadata XML, sent as the multipart "file" part.
  - @param req Optional form fields of the SAML configuration.

@return RestResponseSsoConfig
*/
func (a *AuthenticationSettingsApiService) UpdateSamlSettings(ctx context.Context, accountId string, samlSSOId string, metadata []byte, req *SamlSettingsRequest) (RestResponseSsoConfig, *http.Response, error) {
	var localVarReturnValue RestResponseSsoConfig
	if len(metadata) == 0 {
		return localVarReturnValue, nil, fmt.Errorf("metadata must not be empty")
	}
	httpResp, err := a.executeSamlRequest(ctx, http.MethodPut, "/ng/api/authentication-settings/saml-metadata-upload/"+samlSSOId, "accountId", accountId, metadata, req, &localVarReturnValue)
	return localVarReturnValue, httpResp, err
}

/*
AuthenticationSettingsApiService Delete the SAML SSO configuration with the given samlSSOId
  - @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
  - @param accountIdentifier Account Identifier for the Entity.
  - @param samlSSOId Identifier of the SAML SSO configuration to delete.

@return RestResponseSsoConfig
*/
func (a *AuthenticationSettingsApiService) DeleteSamlSettings(ctx context.Context, accountIdentifier string, samlSSOId string) (RestResponseSsoConfig, *http.Response, error) {
	var localVarReturnValue RestResponseSsoConfig
	httpResp, err := a.executeSamlRequest(ctx, http.MethodDelete, "/ng/api/authentication-settings/saml-metadata/"+samlSSOId+"/delete", "accountIdentifier", accountIdentifier, nil, nil, &localVarReturnValue)
	return localVarReturnValue, httpResp, err
}

/*
AuthenticationSettingsApiService Enable or disable login for the SAML SSO configuration with the given samlSSOId
  - @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
  - @param accountIdentifier Account Identifier for the Entity.
  - @param samlSSOId Identifier of the SAML SSO configuration.
  - @param enable Whether login via this SAML configuration is enabled.

@return RestResponseBoolean
*/
func (a *AuthenticationSettingsApiService) SetSamlAuthenticationEnabled(ctx context.Context, accountIdentifier string, samlSSOId string, enable bool) (RestResponseBoolean, *http.Response, error) {
	var localVarReturnValue RestResponseBoolean
	path := "/ng/api/authentication-settings/saml-metadata-upload/" + samlSSOId + "/authentication?enable=" + strconv.FormatBool(enable)
	httpResp, err := a.executeSamlRequest(ctx, http.MethodPut, path, "accountIdentifier", accountIdentifier, nil, nil, &localVarReturnValue)
	return localVarReturnValue, httpResp, err
}
