package nextgen

import (
	"context"
	"io"
	"mime"
	"mime/multipart"
	"net/http"
	"net/http/httptest"
	"net/url"
	"strings"
	"testing"

	"github.com/hashicorp/go-retryablehttp"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func newAuthExtPassthroughClient(baseURL string) *APIClient {
	httpClient := retryablehttp.NewClient()
	httpClient.RetryMax = 0
	httpClient.Logger = nil
	httpClient.ErrorHandler = func(resp *http.Response, _ error, _ int) (*http.Response, error) {
		return resp, nil
	}
	cfg := NewConfiguration()
	cfg.BasePath = baseURL
	cfg.HTTPClient = httpClient
	return NewAPIClient(cfg)
}

func newAuthExtClient(baseURL string) *APIClient {
	httpClient := retryablehttp.NewClient()
	httpClient.RetryMax = 0
	httpClient.Logger = nil
	cfg := NewConfiguration()
	cfg.BasePath = baseURL
	cfg.HTTPClient = httpClient
	return NewAPIClient(cfg)
}

func authExtServe(t *testing.T, status int, contentType, body string) (*httptest.Server, *authExtCapturedRequest) {
	t.Helper()
	var lastReq authExtCapturedRequest
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		lastReq = authExtCapture(r)
		if contentType != "" {
			w.Header().Set("Content-Type", contentType)
		}
		w.WriteHeader(status)
		_, _ = w.Write([]byte(body))
	}))
	t.Cleanup(server.Close)
	return server, &lastReq
}

func TestAuthenticationSettingsWriteRequest_DecodeBranches(t *testing.T) {
	testCases := map[string]struct {
		status      int
		body        string
		wantErr     bool
		wantModel   interface{} // expected GenericSwaggerError.Model() value; nil means unset
		wantSuccess bool
	}{
		"Given a 200 with a valid body - then the response is decoded": {
			status:      http.StatusOK,
			body:        `{"resource":true}`,
			wantSuccess: true,
		},
		"Given a 200 with an undecodable body - then the decode error is returned": {
			status:  http.StatusOK,
			body:    `not-json`,
			wantErr: true,
		},
		"Given a 400 - then the Failure model is attached to the error": {
			status:    http.StatusBadRequest,
			body:      `{"status":"FAILURE","code":"INVALID_REQUEST","message":"bad domain"}`,
			wantErr:   true,
			wantModel: Failure{Status: "FAILURE", Code: "INVALID_REQUEST", Message: "bad domain"},
		},
		"Given a 404 - then the Failure model is attached to the error": {
			status:    http.StatusNotFound,
			body:      `{"status":"FAILURE","code":"RESOURCE_NOT_FOUND","message":"not found"}`,
			wantErr:   true,
			wantModel: Failure{Status: "FAILURE", Code: "RESOURCE_NOT_FOUND", Message: "not found"},
		},
		"Given a 500 with a valid error body - then the ModelError model is attached to the error": {
			status:    http.StatusInternalServerError,
			body:      `{"status":"ERROR","code":"UNKNOWN_ERROR","message":"boom"}`,
			wantErr:   true,
			wantModel: ModelError{Status: "ERROR", Code: "UNKNOWN_ERROR", Message: "boom"},
		},
		"Given a 500 with an undecodable body - then no model is attached": {
			status:  http.StatusInternalServerError,
			body:    `<html>proxy error</html>`,
			wantErr: true,
		},
	}

	for name, tc := range testCases {
		t.Run(name, func(t *testing.T) {
			server, lastReq := authExtServe(t, tc.status, "application/json", tc.body)

			out, resp, err := newAuthExtPassthroughClient(server.URL).AuthenticationSettingsApi.SetWhitelistedDomains(
				context.Background(), []string{"example.com"}, "acc")

			require.NotNil(t, resp)
			assert.Equal(t, tc.status, resp.StatusCode)
			// wire assertions: PUT with accountIdentifier query param and JSON body
			assert.Equal(t, http.MethodPut, lastReq.method)
			assert.Equal(t, "/ng/api/authentication-settings/whitelisted-domains", lastReq.path)
			assert.Equal(t, "acc", lastReq.query.Get("accountIdentifier"))

			if tc.wantSuccess {
				require.NoError(t, err)
				assert.True(t, out.Resource)
				return
			}

			require.Error(t, err)
			swaggerErr, ok := err.(GenericSwaggerError)
			if !ok {
				// plain decode error on the 2xx path
				assert.Equal(t, http.StatusOK, tc.status)
				return
			}
			if tc.wantModel == nil {
				assert.Nil(t, swaggerErr.Model())
			} else {
				assert.Equal(t, tc.wantModel, swaggerErr.Model())
			}
		})
	}
}

func TestGetAuthenticationSettingsV2_DecodesPolymorphicEntries(t *testing.T) {
	body := `{"resource":{
		"whitelistedDomains":["example.com"],
		"authenticationMechanism":"SAML",
		"twoFactorEnabled":true,
		"sessionTimeoutInMinutes":60,
		"ngAuthSettings":[
			{"settingsType":"USER_PASSWORD","loginSettings":{"uuid":"ls-1","accountId":"acc",
				"userLockoutPolicy":{"enableLockoutPolicy":true,"numberOfFailedAttemptsBeforeLockout":5,"lockOutPeriod":24,"notifyUser":true,"userGroupsToNotify":[{"name":"admins"}]},
				"passwordExpirationPolicy":{"enabled":false},
				"passwordStrengthPolicy":{"enabled":true}}},
			{"settingsType":"SAML","identifier":"saml-1","displayName":"Test","authenticationEnabled":true},
			{"settingsType":"OAUTH","allowedProviders":["GOOGLE"],"filter":"x"}
		]}}`
	server, _ := authExtServe(t, http.StatusOK, "application/json", body)

	out, _, err := newAuthExtClient(server.URL).AuthenticationSettingsApi.GetAuthenticationSettingsV2(context.Background(), "acc")

	require.NoError(t, err)
	require.NotNil(t, out.Resource)
	assert.Equal(t, []string{"example.com"}, out.Resource.WhitelistedDomains)
	assert.True(t, out.Resource.TwoFactorEnabled)
	assert.Equal(t, int32(60), out.Resource.SessionTimeoutInMinutes)
	require.Len(t, out.Resource.NgAuthSettings, 3)

	userPassword := out.Resource.NgAuthSettings[0]
	assert.Equal(t, "USER_PASSWORD", userPassword.SettingsType)
	require.NotNil(t, userPassword.LoginSettings)
	assert.Equal(t, "ls-1", userPassword.LoginSettings.Uuid)
	require.NotNil(t, userPassword.LoginSettings.UserLockoutPolicy)
	require.Len(t, userPassword.LoginSettings.UserLockoutPolicy.UserGroupsToNotify, 1)
	assert.Equal(t, "admins", userPassword.LoginSettings.UserLockoutPolicy.UserGroupsToNotify[0]["name"])

	saml := out.Resource.NgAuthSettings[1]
	assert.Equal(t, "SAML", saml.SettingsType)
	assert.Equal(t, "saml-1", saml.Identifier)
	assert.True(t, saml.AuthenticationEnabled)

	oauth := out.Resource.NgAuthSettings[2]
	assert.Equal(t, "OAUTH", oauth.SettingsType)
	assert.Equal(t, []string{"GOOGLE"}, oauth.AllowedProviders)
	assert.Equal(t, "x", oauth.Filter)
}

func TestGetLoginSettings_ReadsFromUserPasswordEntry(t *testing.T) {
	body := `{"resource":{"ngAuthSettings":[
		{"settingsType":"SAML","identifier":"saml-1"},
		{"settingsType":"USER_PASSWORD","loginSettings":{"uuid":"ls-1","accountId":"acc",
			"lastUpdatedBy":{"uuid":"u-1","name":"Test User","email":"test@harness.io","externalUserId":"ext-1"},
			"lastUpdatedAt":1795155972000,
			"passwordStrengthPolicy":{"enabled":true}}}
	]}}`
	server, _ := authExtServe(t, http.StatusOK, "application/json", body)

	out, _, err := newAuthExtClient(server.URL).AuthenticationSettingsApi.GetLoginSettings(context.Background(), "acc")

	require.NoError(t, err)
	require.NotNil(t, out.Resource)
	assert.Equal(t, "ls-1", out.Resource.Uuid)
	require.NotNil(t, out.Resource.PasswordStrengthPolicy)
	assert.True(t, out.Resource.PasswordStrengthPolicy.Enabled)
	require.NotNil(t, out.Resource.LastUpdatedBy)
	assert.Equal(t, "u-1", out.Resource.LastUpdatedBy.Uuid)
	assert.Equal(t, "Test User", out.Resource.LastUpdatedBy.Name)
	assert.Equal(t, "test@harness.io", out.Resource.LastUpdatedBy.Email)
	assert.Equal(t, "ext-1", out.Resource.LastUpdatedBy.ExternalUserId)
	assert.Equal(t, int64(1795155972000), out.Resource.LastUpdatedAt)
}

type authExtCapturedRequest struct {
	method      string
	path        string
	query       url.Values
	contentType string
	header      http.Header
	body        []byte
}

func authExtCapture(r *http.Request) authExtCapturedRequest {
	body, _ := io.ReadAll(r.Body)
	return authExtCapturedRequest{
		method:      r.Method,
		path:        r.URL.Path,
		query:       r.URL.Query(),
		contentType: r.Header.Get("Content-Type"),
		header:      r.Header,
		body:        body,
	}
}

func authExtParseMultipart(t *testing.T, req authExtCapturedRequest) (values map[string]string, fileParts map[string]string) {
	t.Helper()
	mediaType, params, err := mime.ParseMediaType(req.contentType)
	require.NoError(t, err)
	require.True(t, strings.HasPrefix(mediaType, "multipart/"), "expected multipart body, got %s", mediaType)
	reader := multipart.NewReader(strings.NewReader(string(req.body)), params["boundary"])
	values = map[string]string{}
	fileParts = map[string]string{}
	for {
		part, err := reader.NextPart()
		if err == io.EOF {
			break
		}
		require.NoError(t, err)
		data, err := io.ReadAll(part)
		require.NoError(t, err)
		if part.FileName() != "" {
			fileParts[part.FormName()] = string(data)
		} else {
			values[part.FormName()] = string(data)
		}
	}
	return values, fileParts
}

func TestSamlRequest_CreateSendsMultipartWithExplicitBooleans(t *testing.T) {
	var lastReq authExtCapturedRequest
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		lastReq = authExtCapture(r)
		w.Header().Set("Content-Type", "application/json")
		_, _ = w.Write([]byte(`{"resource":{"uuid":"saml-1","displayName":"Test"}}`))
	}))
	defer server.Close()

	// Booleans left nil: the SDK must send explicit "false" (UI parity) rather
	// than omitting the fields, which the backend also treats as false but
	// only after a fragile null -> "null" -> parseBoolean chain.
	out, _, err := newAuthExtClient(server.URL).AuthenticationSettingsApi.CreateSamlSettings(
		context.Background(), "acc", []byte("<xml/>"), &SamlSettingsRequest{DisplayName: "Test"})

	require.NoError(t, err)
	assert.Equal(t, "saml-1", out.Resource.Uuid)

	assert.Equal(t, http.MethodPost, lastReq.method)
	assert.Equal(t, "/ng/api/authentication-settings/saml-metadata", lastReq.path)
	assert.Equal(t, "acc", lastReq.query.Get("accountId"))
	assert.True(t, strings.HasPrefix(lastReq.contentType, "multipart/form-data"))

	values, files := authExtParseMultipart(t, lastReq)
	assert.Equal(t, "<xml/>", files["file"])
	assert.Equal(t, "Test", values["displayName"])
	assert.Equal(t, "false", values["authorizationEnabled"])
	assert.Equal(t, "false", values["jitEnabled"])
}

func TestSamlRequest_CreateRequiresMetadata(t *testing.T) {
	_, resp, err := newAuthExtClient("http://localhost:1").AuthenticationSettingsApi.CreateSamlSettings(
		context.Background(), "acc", nil, nil)

	require.Error(t, err)
	assert.Nil(t, resp)
	assert.Contains(t, err.Error(), "metadata must not be empty")
}

// The backend rebuilds the SAML configuration from the request on every call,
// so a nil request would silently reset authorizationEnabled/jitEnabled and
// clear the stored values. The SDK must reject it before any HTTP call.
func TestSamlRequest_CreateRequiresRequest(t *testing.T) {
	_, resp, err := newAuthExtClient("http://localhost:1").AuthenticationSettingsApi.CreateSamlSettings(
		context.Background(), "acc", []byte("<xml/>"), nil)

	require.Error(t, err)
	assert.Nil(t, resp)
	assert.Contains(t, err.Error(), "req must not be nil")
}

func TestSamlRequest_UpdateRequiresRequest(t *testing.T) {
	_, resp, err := newAuthExtClient("http://localhost:1").AuthenticationSettingsApi.UpdateSamlSettings(
		context.Background(), "acc", "saml-1", []byte("<xml/>"), nil)

	require.Error(t, err)
	assert.Nil(t, resp)
	assert.Contains(t, err.Error(), "req must not be nil")
}

func TestSamlRequest_UpdateWithoutMetadataSendsFormFieldsOnly(t *testing.T) {
	var lastReq authExtCapturedRequest
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		lastReq = authExtCapture(r)
		w.Header().Set("Content-Type", "application/json")
		_, _ = w.Write([]byte(`{"resource":{"accountId":"acc"}}`))
	}))
	defer server.Close()

	enabled := true
	_, _, err := newAuthExtClient(server.URL).AuthenticationSettingsApi.UpdateSamlSettings(
		context.Background(), "acc", "saml-1", nil,
		&SamlSettingsRequest{DisplayName: "Renamed", AuthorizationEnabled: &enabled, JitEnabled: &enabled})

	require.NoError(t, err)

	assert.Equal(t, http.MethodPut, lastReq.method)
	assert.Equal(t, "/ng/api/authentication-settings/saml-metadata-upload/saml-1", lastReq.path)
	assert.Equal(t, "acc", lastReq.query.Get("accountIdentifier"))

	values, files := authExtParseMultipart(t, lastReq)
	assert.Empty(t, files) // no file part: backend keeps the stored metadata
	assert.Equal(t, "Renamed", values["displayName"])
	assert.Equal(t, "true", values["authorizationEnabled"])
	assert.Equal(t, "true", values["jitEnabled"])
}

func TestSamlRequest_DecodeBranches(t *testing.T) {
	testCases := map[string]struct {
		status    int
		body      string
		wantErr   bool
		wantModel interface{}
	}{
		"Given a 400 - then the Failure model is attached": {
			status:    http.StatusBadRequest,
			body:      `{"status":"FAILURE","code":"INVALID_REQUEST","message":"No Saml Metadata found for account acc and saml sso id x"}`,
			wantErr:   true,
			wantModel: Failure{Status: "FAILURE", Code: "INVALID_REQUEST", Message: "No Saml Metadata found for account acc and saml sso id x"},
		},
		"Given a 500 with a valid error body - then the ModelError model is attached": {
			status:    http.StatusInternalServerError,
			body:      `{"status":"ERROR","code":"UNKNOWN_ERROR","message":"boom"}`,
			wantErr:   true,
			wantModel: ModelError{Status: "ERROR", Code: "UNKNOWN_ERROR", Message: "boom"},
		},
		"Given a 500 with an undecodable body - then no model is attached": {
			status:  http.StatusInternalServerError,
			body:    `garbage`,
			wantErr: true,
		},
	}

	for name, tc := range testCases {
		t.Run(name, func(t *testing.T) {
			server, _ := authExtServe(t, tc.status, "application/json", tc.body)

			_, resp, err := newAuthExtPassthroughClient(server.URL).AuthenticationSettingsApi.DeleteSamlSettings(context.Background(), "acc", "x")

			require.NotNil(t, resp)
			require.Error(t, err)
			swaggerErr, ok := err.(GenericSwaggerError)
			require.True(t, ok, "expected GenericSwaggerError, got %T", err)
			if tc.wantModel == nil {
				assert.Nil(t, swaggerErr.Model())
			} else {
				assert.Equal(t, tc.wantModel, swaggerErr.Model())
			}
		})
	}
}

func TestSetSamlAuthenticationEnabled_WireFormat(t *testing.T) {
	var lastReq authExtCapturedRequest
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		lastReq = authExtCapture(r)
		w.Header().Set("Content-Type", "application/json")
		_, _ = w.Write([]byte(`{"resource":true}`))
	}))
	defer server.Close()

	out, _, err := newAuthExtClient(server.URL).AuthenticationSettingsApi.SetSamlAuthenticationEnabled(
		context.Background(), "acc", "saml-1", true)

	require.NoError(t, err)
	assert.True(t, out.Resource)
	assert.Equal(t, http.MethodPut, lastReq.method)
	assert.Equal(t, "/ng/api/authentication-settings/saml-metadata-upload/saml-1/authentication", lastReq.path)
	assert.Equal(t, url.Values{"accountIdentifier": {"acc"}, "enable": {"true"}}, lastReq.query)
}

func TestSamlRequest_CreateSendsAllOptionalFormFields(t *testing.T) {
	var lastReq authExtCapturedRequest
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		lastReq = authExtCapture(r)
		w.Header().Set("Content-Type", "application/json")
		_, _ = w.Write([]byte(`{"resource":{"uuid":"saml-1"}}`))
	}))
	defer server.Close()

	disabled := false
	_, _, err := newAuthExtClient(server.URL).AuthenticationSettingsApi.CreateSamlSettings(
		context.Background(), "acc", []byte("<xml/>"), &SamlSettingsRequest{
			DisplayName:          "Test",
			GroupMembershipAttr:  "memberOf",
			AuthorizationEnabled: &disabled,
			LogoutUrl:            "https://idp.example.com/logout",
			EntityIdentifier:     "entity-1",
			SamlProviderType:     "AZURE",
			ClientId:             "client-1",
			ClientSecret:         "account.secret-ref",
			FriendlySamlName:     "friendly",
			JitEnabled:           &disabled,
			JitValidationKey:     "email",
			JitValidationValue:   "example.com",
		})

	require.NoError(t, err)

	values, files := authExtParseMultipart(t, lastReq)
	assert.Equal(t, "<xml/>", files["file"])
	assert.Equal(t, map[string]string{
		"displayName":          "Test",
		"groupMembershipAttr":  "memberOf",
		"authorizationEnabled": "false",
		"logoutUrl":            "https://idp.example.com/logout",
		"entityIdentifier":     "entity-1",
		"samlProviderType":     "AZURE",
		"clientId":             "client-1",
		"clientSecret":         "account.secret-ref",
		"friendlySamlName":     "friendly",
		"jitEnabled":           "false",
		"jitValidationKey":     "email",
		"jitValidationValue":   "example.com",
	}, values)
}

func TestSamlRequest_TransportError(t *testing.T) {
	enabled := true
	_, _, err := newAuthExtClient("http://localhost:1").AuthenticationSettingsApi.UpdateSamlSettings(
		context.Background(), "acc", "saml-1", []byte("<xml/>"),
		&SamlSettingsRequest{DisplayName: "Test", AuthorizationEnabled: &enabled, JitEnabled: &enabled})

	require.Error(t, err)
	_, ok := err.(GenericSwaggerError)
	assert.False(t, ok, "transport failures must not be wrapped in GenericSwaggerError, got %T", err)
}

func TestSamlRequest_ApiKeyWithPrefixSetsHeader(t *testing.T) {
	var lastReq authExtCapturedRequest
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		lastReq = authExtCapture(r)
		w.Header().Set("Content-Type", "application/json")
		_, _ = w.Write([]byte(`{"resource":true}`))
	}))
	defer server.Close()

	ctx := context.WithValue(context.Background(), ContextAPIKey, APIKey{Key: "k", Prefix: "Bearer"})
	_, _, err := newAuthExtClient(server.URL).AuthenticationSettingsApi.SetSamlAuthenticationEnabled(ctx, "acc", "saml-1", false)

	require.NoError(t, err)
	assert.Equal(t, "Bearer k", lastReq.header.Get("x-api-key"))
}

func TestSamlRequest_NonModelErrorStatusFallsThrough(t *testing.T) {
	server, _ := authExtServe(t, http.StatusForbidden, "application/json", `{"status":"FAILURE","code":"FORBIDDEN","message":"denied"}`)

	_, _, err := newAuthExtPassthroughClient(server.URL).AuthenticationSettingsApi.DeleteSamlSettings(context.Background(), "acc", "x")

	require.Error(t, err)
	swaggerErr, ok := err.(GenericSwaggerError)
	require.True(t, ok, "expected GenericSwaggerError, got %T", err)
	assert.Nil(t, swaggerErr.Model(), "only 400/404/500 decode into a typed model")
}

func TestSessionTimeoutAndPublicAccess_WireFormat(t *testing.T) {
	t.Run("Given session timeout settings - then they are PUT to the session timeout path", func(t *testing.T) {
		var lastReq authExtCapturedRequest
		server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			lastReq = authExtCapture(r)
			w.Header().Set("Content-Type", "application/json")
			_, _ = w.Write([]byte(`{"resource":true}`))
		}))
		defer server.Close()

		out, _, err := newAuthExtClient(server.URL).AuthenticationSettingsApi.SetSessionTimeoutAtAccountLevel(
			context.Background(), SessionTimeoutSettings{SessionTimeOutInMinutes: 60, AbsoluteSessionTimeOutInMinutes: 480}, "acc")

		require.NoError(t, err)
		assert.True(t, out.Resource)
		assert.Equal(t, http.MethodPut, lastReq.method)
		assert.Equal(t, "/ng/api/authentication-settings/session-timeout-account-level", lastReq.path)
		assert.Equal(t, "acc", lastReq.query.Get("accountIdentifier"))
		assert.JSONEq(t, `{"sessionTimeOutInMinutes":60,"absoluteSessionTimeOutInMinutes":480}`, string(lastReq.body))
	})

	t.Run("Given a public access flag - then it is PUT to the public access path", func(t *testing.T) {
		var lastReq authExtCapturedRequest
		server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			lastReq = authExtCapture(r)
			w.Header().Set("Content-Type", "application/json")
			_, _ = w.Write([]byte(`{"resource":true}`))
		}))
		defer server.Close()

		out, _, err := newAuthExtClient(server.URL).AuthenticationSettingsApi.SetPublicAccess(context.Background(), true, "acc")

		require.NoError(t, err)
		assert.True(t, out.Resource)
		assert.Equal(t, "/ng/api/authentication-settings/public-access", lastReq.path)
		assert.Equal(t, "true", strings.TrimSpace(string(lastReq.body)))
	})

	t.Run("Given an unreachable server - then the transport error is returned unwrapped", func(t *testing.T) {
		_, _, err := newAuthExtClient("http://localhost:1").AuthenticationSettingsApi.SetPublicAccess(context.Background(), true, "acc")

		require.Error(t, err)
		_, ok := err.(GenericSwaggerError)
		assert.False(t, ok, "transport failures must not be wrapped in GenericSwaggerError, got %T", err)
	})
}

func TestSetTwoFactorAuthAndOauthProviders_WireFormat(t *testing.T) {
	t.Run("Given two factor override settings - then they are PUT to the two factor path", func(t *testing.T) {
		var lastReq authExtCapturedRequest
		server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			lastReq = authExtCapture(r)
			w.Header().Set("Content-Type", "application/json")
			_, _ = w.Write([]byte(`{"resource":true}`))
		}))
		defer server.Close()

		out, _, err := newAuthExtClient(server.URL).AuthenticationSettingsApi.SetTwoFactorAuth(
			context.Background(), TwoFactorAdminOverrideSettings{AdminOverrideTwoFactorEnabled: true}, "acc")

		require.NoError(t, err)
		assert.True(t, out.Resource)
		assert.Equal(t, http.MethodPut, lastReq.method)
		assert.Equal(t, "/ng/api/authentication-settings/two-factor-admin-override-settings", lastReq.path)
		assert.Equal(t, "acc", lastReq.query.Get("accountIdentifier"))
		assert.JSONEq(t, `{"adminOverrideTwoFactorEnabled":true}`, string(lastReq.body))
	})

	t.Run("Given OAuth providers - then they are PUT to the oauth providers path", func(t *testing.T) {
		var lastReq authExtCapturedRequest
		server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			lastReq = authExtCapture(r)
			w.Header().Set("Content-Type", "application/json")
			_, _ = w.Write([]byte(`{"resource":true}`))
		}))
		defer server.Close()

		out, _, err := newAuthExtClient(server.URL).AuthenticationSettingsApi.SetOauthProviders(
			context.Background(), OAuthSettings{AllowedProviders: []string{"GOOGLE"}}, "acc")

		require.NoError(t, err)
		assert.True(t, out.Resource)
		assert.Equal(t, http.MethodPut, lastReq.method)
		assert.Equal(t, "/ng/api/authentication-settings/oauth/update-providers", lastReq.path)
		assert.Equal(t, "acc", lastReq.query.Get("accountIdentifier"))
		assert.JSONEq(t, `{"allowedProviders":["GOOGLE"]}`, string(lastReq.body))
	})

	t.Run("Given an API key with a prefix - then the prefixed key is sent", func(t *testing.T) {
		var lastReq authExtCapturedRequest
		server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			lastReq = authExtCapture(r)
			w.Header().Set("Content-Type", "application/json")
			_, _ = w.Write([]byte(`{"resource":true}`))
		}))
		defer server.Close()

		ctx := context.WithValue(context.Background(), ContextAPIKey, APIKey{Key: "k", Prefix: "Bearer"})
		_, _, err := newAuthExtClient(server.URL).AuthenticationSettingsApi.SetWhitelistedDomains(
			ctx, []string{"example.com"}, "acc")

		require.NoError(t, err)
		assert.Equal(t, "Bearer k", lastReq.header.Get("x-api-key"))
	})

	t.Run("Given a 403 - then no typed model is attached", func(t *testing.T) {
		server, _ := authExtServe(t, http.StatusForbidden, "application/json",
			`{"status":"FAILURE","code":"FORBIDDEN","message":"denied"}`)

		_, _, err := newAuthExtPassthroughClient(server.URL).AuthenticationSettingsApi.SetWhitelistedDomains(
			context.Background(), []string{"example.com"}, "acc")

		require.Error(t, err)
		swaggerErr, ok := err.(GenericSwaggerError)
		require.True(t, ok, "expected GenericSwaggerError, got %T", err)
		assert.Nil(t, swaggerErr.Model(), "only 400/404/500 decode into a typed model")
	})
}

func TestUpdateLoginSettings_WireFormatAnd500Model(t *testing.T) {
	t.Run("Given a full login settings body - then it is PUT to the login settings path", func(t *testing.T) {
		var lastReq authExtCapturedRequest
		server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			lastReq = authExtCapture(r)
			w.Header().Set("Content-Type", "application/json")
			_, _ = w.Write([]byte(`{"resource":{"uuid":"ls-1"}}`))
		}))
		defer server.Close()

		settings := LoginSettings{
			Uuid:      "ls-1",
			AccountId: "acc",
			UserLockoutPolicy: &UserLockoutPolicy{
				EnableLockoutPolicy:                 true,
				NumberOfFailedAttemptsBeforeLockout: 5,
				LockOutPeriod:                       24,
				NotifyUser:                          true,
				UserGroupsToNotify:                  []map[string]interface{}{{"name": "admins"}},
			},
		}
		out, _, err := newAuthExtClient(server.URL).AuthenticationSettingsApi.UpdateLoginSettings(context.Background(), settings, "ls-1", "acc")

		require.NoError(t, err)
		assert.Equal(t, "ls-1", out.Resource.Uuid)
		assert.Equal(t, http.MethodPut, lastReq.method)
		assert.Equal(t, "/ng/api/authentication-settings/login-settings/ls-1", lastReq.path)
		assert.Equal(t, "acc", lastReq.query.Get("accountIdentifier"))
		// userGroupsToNotify must survive the round trip untouched
		assert.Contains(t, string(lastReq.body), `"userGroupsToNotify":[{"name":"admins"}]`)
	})

	t.Run("Given a 500 with a valid error body - then the ModelError model is attached", func(t *testing.T) {
		server, _ := authExtServe(t, http.StatusInternalServerError, "application/json",
			`{"status":"ERROR","code":"UNKNOWN_ERROR","message":"boom"}`)

		_, _, err := newAuthExtPassthroughClient(server.URL).AuthenticationSettingsApi.UpdateLoginSettings(
			context.Background(), LoginSettings{}, "ls-1", "acc")

		require.Error(t, err)
		swaggerErr, ok := err.(GenericSwaggerError)
		require.True(t, ok, "expected GenericSwaggerError, got %T", err)
		assert.Equal(t, ModelError{Status: "ERROR", Code: "UNKNOWN_ERROR", Message: "boom"}, swaggerErr.Model())
	})

	t.Run("Given a 403 - then no typed model is attached", func(t *testing.T) {
		server, _ := authExtServe(t, http.StatusForbidden, "application/json",
			`{"status":"FAILURE","code":"FORBIDDEN","message":"denied"}`)

		_, _, err := newAuthExtPassthroughClient(server.URL).AuthenticationSettingsApi.UpdateLoginSettings(
			context.Background(), LoginSettings{}, "ls-1", "acc")

		require.Error(t, err)
		swaggerErr, ok := err.(GenericSwaggerError)
		require.True(t, ok, "expected GenericSwaggerError, got %T", err)
		assert.Nil(t, swaggerErr.Model(), "only 400/404/500 decode into a typed model")
	})

	t.Run("Given an unreachable server - then the transport error is returned unwrapped", func(t *testing.T) {
		_, _, err := newAuthExtClient("http://localhost:1").AuthenticationSettingsApi.UpdateLoginSettings(
			context.Background(), LoginSettings{}, "ls-1", "acc")

		require.Error(t, err)
		_, ok := err.(GenericSwaggerError)
		assert.False(t, ok, "transport failures must not be wrapped in GenericSwaggerError, got %T", err)
	})

	t.Run("Given a 2xx with an undecodable body - then the decode error is returned", func(t *testing.T) {
		server, _ := authExtServe(t, http.StatusOK, "application/json", `not-json`)

		_, _, err := newAuthExtClient(server.URL).AuthenticationSettingsApi.UpdateLoginSettings(
			context.Background(), LoginSettings{}, "ls-1", "acc")

		require.Error(t, err)
		_, ok := err.(GenericSwaggerError)
		assert.False(t, ok, "a 2xx decode failure must not be wrapped in GenericSwaggerError, got %T", err)
	})

	t.Run("Given a 400 and a 500 with undecodable bodies - then no typed model is attached", func(t *testing.T) {
		for _, status := range []int{http.StatusBadRequest, http.StatusInternalServerError} {
			server, _ := authExtServe(t, status, "application/json", `<html>proxy error</html>`)

			_, _, err := newAuthExtPassthroughClient(server.URL).AuthenticationSettingsApi.UpdateLoginSettings(
				context.Background(), LoginSettings{}, "ls-1", "acc")

			require.Error(t, err)
			swaggerErr, ok := err.(GenericSwaggerError)
			require.True(t, ok, "expected GenericSwaggerError, got %T", err)
			assert.Nil(t, swaggerErr.Model(), "an undecodable error body must not attach a model")
		}
	})
}

func TestGetLoginSettings_PropagatesV2Error(t *testing.T) {
	server, _ := authExtServe(t, http.StatusInternalServerError, "application/json",
		`{"status":"ERROR","code":"UNKNOWN_ERROR","message":"boom"}`)

	_, _, err := newAuthExtPassthroughClient(server.URL).AuthenticationSettingsApi.GetLoginSettings(context.Background(), "acc")

	require.Error(t, err)
	swaggerErr, ok := err.(GenericSwaggerError)
	require.True(t, ok, "expected GenericSwaggerError, got %T", err)
	assert.Equal(t, ModelError{Status: "ERROR", Code: "UNKNOWN_ERROR", Message: "boom"}, swaggerErr.Model())
}
