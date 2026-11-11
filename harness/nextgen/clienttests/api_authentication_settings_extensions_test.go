package clienttests

// Offline tests for the hand-written authentication settings extensions
// (PL-72067): executeAuthenticationSettingsWriteRequest, executeSamlRequest
// and executeUpdateLoginSettingsRequest. They cover the response-decode
// branches (<300, 400/404, 500) and the SAML multipart/form wire format,
// including two regressions fixed in this PR:
//   - the 500 branch must attach the decoded ModelError to GenericSwaggerError
//     on decode success (previously only set on decode failure),
//   - the SAML booleans authorizationEnabled/jitEnabled must always be sent
//     (false when unset), because the backend treats a missing value as false
//     and nulls the stored groupMembershipAttr.

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

	"github.com/harness/harness-go-sdk/harness/nextgen"
	"github.com/hashicorp/go-retryablehttp"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

// newPassthroughClient builds a client whose retryablehttp ErrorHandler returns
// the response instead of draining it: with the default policy a retried 500
// surfaces as a "giving up" error and the helpers' error-decode branches never
// run. Pass-through makes the response reachable so the decode logic itself
// can be exercised.
func newPassthroughClient(baseURL string) *nextgen.APIClient {
	httpClient := retryablehttp.NewClient()
	httpClient.RetryMax = 0
	httpClient.Logger = nil
	httpClient.ErrorHandler = func(resp *http.Response, _ error, _ int) (*http.Response, error) {
		return resp, nil
	}
	cfg := nextgen.NewConfiguration()
	cfg.BasePath = baseURL
	cfg.HTTPClient = httpClient
	return nextgen.NewAPIClient(cfg)
}

// serve starts a test server returning the given status/body for every request
// and captures the last request for assertions.
func serve(t *testing.T, status int, contentType, body string) (*httptest.Server, *capturedRequest) {
	t.Helper()
	var lastReq capturedRequest
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		lastReq = capture(r)
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
			wantModel: nextgen.Failure{Status: "FAILURE", Code: "INVALID_REQUEST", Message: "bad domain"},
		},
		"Given a 404 - then the Failure model is attached to the error": {
			status:    http.StatusNotFound,
			body:      `{"status":"FAILURE","code":"RESOURCE_NOT_FOUND","message":"not found"}`,
			wantErr:   true,
			wantModel: nextgen.Failure{Status: "FAILURE", Code: "RESOURCE_NOT_FOUND", Message: "not found"},
		},
		"Given a 500 with a valid error body - then the ModelError model is attached to the error": {
			status:    http.StatusInternalServerError,
			body:      `{"status":"ERROR","code":"UNKNOWN_ERROR","message":"boom"}`,
			wantErr:   true,
			wantModel: nextgen.ModelError{Status: "ERROR", Code: "UNKNOWN_ERROR", Message: "boom"},
		},
		"Given a 500 with an undecodable body - then no model is attached": {
			status:  http.StatusInternalServerError,
			body:    `<html>proxy error</html>`,
			wantErr: true,
		},
	}

	for name, tc := range testCases {
		t.Run(name, func(t *testing.T) {
			server, lastReq := serve(t, tc.status, "application/json", tc.body)

			out, resp, err := newPassthroughClient(server.URL).AuthenticationSettingsApi.SetWhitelistedDomains(
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
			swaggerErr, ok := err.(nextgen.GenericSwaggerError)
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
	server, _ := serve(t, http.StatusOK, "application/json", body)

	out, _, err := newClient(server.URL).AuthenticationSettingsApi.GetAuthenticationSettingsV2(context.Background(), "acc")

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
			"passwordStrengthPolicy":{"enabled":true}}}
	]}}`
	server, _ := serve(t, http.StatusOK, "application/json", body)

	out, _, err := newClient(server.URL).AuthenticationSettingsApi.GetLoginSettings(context.Background(), "acc")

	require.NoError(t, err)
	require.NotNil(t, out.Resource)
	assert.Equal(t, "ls-1", out.Resource.Uuid)
	require.NotNil(t, out.Resource.PasswordStrengthPolicy)
	assert.True(t, out.Resource.PasswordStrengthPolicy.Enabled)
}

// capturedRequest holds a request captured inside the handler (the body is
// closed by the server once the handler returns, so it must be read there).
type capturedRequest struct {
	method      string
	path        string
	query       url.Values
	contentType string
	body        []byte
}

func capture(r *http.Request) capturedRequest {
	body, _ := io.ReadAll(r.Body)
	return capturedRequest{
		method:      r.Method,
		path:        r.URL.Path,
		query:       r.URL.Query(),
		contentType: r.Header.Get("Content-Type"),
		body:        body,
	}
}

// parseMultipart parses a captured multipart body into form values and file parts.
func parseMultipart(t *testing.T, req capturedRequest) (values map[string]string, fileParts map[string]string) {
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
	var lastReq capturedRequest
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		lastReq = capture(r)
		w.Header().Set("Content-Type", "application/json")
		_, _ = w.Write([]byte(`{"resource":{"uuid":"saml-1","displayName":"Test"}}`))
	}))
	defer server.Close()

	// Booleans left nil: the SDK must send explicit "false" (UI parity) rather
	// than omitting the fields, which the backend also treats as false but
	// only after a fragile null -> "null" -> parseBoolean chain.
	out, _, err := newClient(server.URL).AuthenticationSettingsApi.CreateSamlSettings(
		context.Background(), "acc", []byte("<xml/>"), &nextgen.SamlSettingsRequest{DisplayName: "Test"})

	require.NoError(t, err)
	assert.Equal(t, "saml-1", out.Resource.Uuid)

	assert.Equal(t, http.MethodPost, lastReq.method)
	assert.Equal(t, "/ng/api/authentication-settings/saml-metadata", lastReq.path)
	assert.Equal(t, "acc", lastReq.query.Get("accountId"))
	assert.True(t, strings.HasPrefix(lastReq.contentType, "multipart/form-data"))

	values, files := parseMultipart(t, lastReq)
	assert.Equal(t, "<xml/>", files["file"])
	assert.Equal(t, "Test", values["displayName"])
	assert.Equal(t, "false", values["authorizationEnabled"])
	assert.Equal(t, "false", values["jitEnabled"])
}

func TestSamlRequest_CreateRequiresMetadata(t *testing.T) {
	_, resp, err := newClient("http://localhost:1").AuthenticationSettingsApi.CreateSamlSettings(
		context.Background(), "acc", nil, nil)

	require.Error(t, err)
	assert.Nil(t, resp)
	assert.Contains(t, err.Error(), "metadata must not be empty")
}

func TestSamlRequest_UpdateWithoutMetadataSendsFormFieldsOnly(t *testing.T) {
	var lastReq capturedRequest
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		lastReq = capture(r)
		w.Header().Set("Content-Type", "application/json")
		_, _ = w.Write([]byte(`{"resource":{"accountId":"acc"}}`))
	}))
	defer server.Close()

	enabled := true
	_, _, err := newClient(server.URL).AuthenticationSettingsApi.UpdateSamlSettings(
		context.Background(), "acc", "saml-1", nil,
		&nextgen.SamlSettingsRequest{DisplayName: "Renamed", AuthorizationEnabled: &enabled, JitEnabled: &enabled})

	require.NoError(t, err)

	assert.Equal(t, http.MethodPut, lastReq.method)
	assert.Equal(t, "/ng/api/authentication-settings/saml-metadata-upload/saml-1", lastReq.path)
	assert.Equal(t, "acc", lastReq.query.Get("accountIdentifier"))

	values, files := parseMultipart(t, lastReq)
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
			wantModel: nextgen.Failure{Status: "FAILURE", Code: "INVALID_REQUEST", Message: "No Saml Metadata found for account acc and saml sso id x"},
		},
		"Given a 500 with a valid error body - then the ModelError model is attached": {
			status:    http.StatusInternalServerError,
			body:      `{"status":"ERROR","code":"UNKNOWN_ERROR","message":"boom"}`,
			wantErr:   true,
			wantModel: nextgen.ModelError{Status: "ERROR", Code: "UNKNOWN_ERROR", Message: "boom"},
		},
		"Given a 500 with an undecodable body - then no model is attached": {
			status:  http.StatusInternalServerError,
			body:    `garbage`,
			wantErr: true,
		},
	}

	for name, tc := range testCases {
		t.Run(name, func(t *testing.T) {
			server, _ := serve(t, tc.status, "application/json", tc.body)

			_, resp, err := newPassthroughClient(server.URL).AuthenticationSettingsApi.DeleteSamlSettings(context.Background(), "acc", "x")

			require.NotNil(t, resp)
			require.Error(t, err)
			swaggerErr, ok := err.(nextgen.GenericSwaggerError)
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
	var lastReq capturedRequest
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		lastReq = capture(r)
		w.Header().Set("Content-Type", "application/json")
		_, _ = w.Write([]byte(`{"resource":true}`))
	}))
	defer server.Close()

	out, _, err := newClient(server.URL).AuthenticationSettingsApi.SetSamlAuthenticationEnabled(
		context.Background(), "acc", "saml-1", true)

	require.NoError(t, err)
	assert.True(t, out.Resource)
	assert.Equal(t, http.MethodPut, lastReq.method)
	assert.Equal(t, "/ng/api/authentication-settings/saml-metadata-upload/saml-1/authentication", lastReq.path)
	assert.Equal(t, url.Values{"accountIdentifier": {"acc"}, "enable": {"true"}}, lastReq.query)
}

func TestUpdateLoginSettings_WireFormatAnd500Model(t *testing.T) {
	t.Run("Given a full login settings body - then it is PUT to the login settings path", func(t *testing.T) {
		var lastReq capturedRequest
		server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			lastReq = capture(r)
			w.Header().Set("Content-Type", "application/json")
			_, _ = w.Write([]byte(`{"resource":{"uuid":"ls-1"}}`))
		}))
		defer server.Close()

		settings := nextgen.LoginSettings{
			Uuid:      "ls-1",
			AccountId: "acc",
			UserLockoutPolicy: &nextgen.UserLockoutPolicy{
				EnableLockoutPolicy:                 true,
				NumberOfFailedAttemptsBeforeLockout: 5,
				LockOutPeriod:                       24,
				NotifyUser:                          true,
				UserGroupsToNotify:                  []map[string]interface{}{{"name": "admins"}},
			},
		}
		out, _, err := newClient(server.URL).AuthenticationSettingsApi.UpdateLoginSettings(context.Background(), settings, "ls-1", "acc")

		require.NoError(t, err)
		assert.Equal(t, "ls-1", out.Resource.Uuid)
		assert.Equal(t, http.MethodPut, lastReq.method)
		assert.Equal(t, "/ng/api/authentication-settings/login-settings/ls-1", lastReq.path)
		assert.Equal(t, "acc", lastReq.query.Get("accountIdentifier"))
		// userGroupsToNotify must survive the round trip untouched
		assert.Contains(t, string(lastReq.body), `"userGroupsToNotify":[{"name":"admins"}]`)
	})

	t.Run("Given a 500 with a valid error body - then the ModelError model is attached", func(t *testing.T) {
		server, _ := serve(t, http.StatusInternalServerError, "application/json",
			`{"status":"ERROR","code":"UNKNOWN_ERROR","message":"boom"}`)

		_, _, err := newPassthroughClient(server.URL).AuthenticationSettingsApi.UpdateLoginSettings(
			context.Background(), nextgen.LoginSettings{}, "ls-1", "acc")

		require.Error(t, err)
		swaggerErr, ok := err.(nextgen.GenericSwaggerError)
		require.True(t, ok, "expected GenericSwaggerError, got %T", err)
		assert.Equal(t, nextgen.ModelError{Status: "ERROR", Code: "UNKNOWN_ERROR", Message: "boom"}, swaggerErr.Model())
	})
}
