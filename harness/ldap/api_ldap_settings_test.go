package ldap

import (
	"context"
	"encoding/json"
	"io"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/hashicorp/go-retryablehttp"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func newTestClient(baseURL string) *APIClient {
	httpClient := retryablehttp.NewClient()
	httpClient.RetryMax = 0
	httpClient.Logger = nil
	httpClient.ErrorHandler = func(resp *http.Response, _ error, _ int) (*http.Response, error) {
		return resp, nil
	}
	cfg := NewConfiguration()
	cfg.BasePath = baseURL
	cfg.AccountId = "acc"
	cfg.ApiKey = "test-key"
	cfg.HTTPClient = httpClient
	return NewAPIClient(cfg)
}

type capturedRequest struct {
	method      string
	path        string
	apiKey      string
	account     string
	body        []byte
	contentType string
}

func serve(t *testing.T, status int, body string) (*httptest.Server, *capturedRequest) {
	t.Helper()
	var lastReq capturedRequest
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		reqBody, _ := io.ReadAll(r.Body)
		lastReq = capturedRequest{
			method:      r.Method,
			path:        r.URL.Path,
			apiKey:      r.Header.Get("x-api-key"),
			account:     r.Header.Get("Harness-Account"),
			body:        reqBody,
			contentType: r.Header.Get("Content-Type"),
		}
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(status)
		_, _ = w.Write([]byte(body))
	}))
	t.Cleanup(server.Close)
	return server, &lastReq
}

func sampleSettings() LdapSettingsDTO {
	return LdapSettingsDTO{
		Identifier:     "corp_ldap",
		Name:           "Corp LDAP",
		Url:            "ldaps://ldap.example.com:636",
		CronExpression: "0 0 * * *",
		LdapConnectionSettings: &LdapConnectionSettingsDTO{
			Host:          "ldap.example.com",
			Port:          636,
			SslEnabled:    true,
			BindDN:        "cn=admin,dc=example,dc=com",
			SecretRefPath: "account.test-ldap-secret",
		},
		LdapUserSettings: []LdapUserSettingsDTO{{
			BaseDN:    "ou=users,dc=example,dc=com",
			EmailAttr: "mail",
			UidAttr:   "uid",
		}},
		LdapGroupSettings: []LdapGroupSettingsDTO{{
			BaseDN:   "ou=groups,dc=example,dc=com",
			NameAttr: "cn",
		}},
	}
}

const sampleResponseBody = `{"ldap_settings":{"account_identifier":"acc","identifier":"corp_ldap","name":"Corp LDAP","url":"ldaps://ldap.example.com:636","cron_expression":"0 0 * * *","ldap_connection_settings":{"host":"ldap.example.com","port":636,"ssl_enabled":true,"bind_d_n":"cn=admin,dc=example,dc=com","secret_ref_path":"account.test-ldap-secret"}},"created":1700000000000,"last_modified":1700000001000}`

func TestCreateLdapSettings_WireFormat(t *testing.T) {
	server, lastReq := serve(t, http.StatusOK, sampleResponseBody)

	out, resp, err := newTestClient(server.URL).LdapSettingsApi.CreateLdapSettings(context.Background(), sampleSettings())

	require.NoError(t, err)
	assert.Equal(t, http.StatusOK, resp.StatusCode)
	require.NotNil(t, out.LdapSettings)
	assert.Equal(t, "corp_ldap", out.LdapSettings.Identifier)
	require.NotNil(t, out.Created)
	assert.Equal(t, int64(1700000000000), *out.Created)

	assert.Equal(t, http.MethodPost, lastReq.method)
	assert.Equal(t, "/v1/ldap-settings", lastReq.path)
	assert.Equal(t, "test-key", lastReq.apiKey)
	assert.Equal(t, "acc", lastReq.account)
	assert.Equal(t, "application/json", lastReq.contentType)

	// The request body wraps the settings in ldap_settings and uses the
	// snake_case spellings of the v1 spec (bind_d_n, base_d_n).
	var reqBody map[string]map[string]interface{}
	require.NoError(t, json.Unmarshal(lastReq.body, &reqBody))
	settings := reqBody["ldap_settings"]
	require.NotNil(t, settings)
	assert.Equal(t, "corp_ldap", settings["identifier"])
	conn := settings["ldap_connection_settings"].(map[string]interface{})
	assert.Equal(t, "cn=admin,dc=example,dc=com", conn["bind_d_n"])
	assert.Equal(t, "account.test-ldap-secret", conn["secret_ref_path"])
	user := settings["ldap_user_settings"].([]interface{})[0].(map[string]interface{})
	assert.Equal(t, "ou=users,dc=example,dc=com", user["base_d_n"])
}

func TestGetLdapSettings_WireFormat(t *testing.T) {
	server, lastReq := serve(t, http.StatusOK, sampleResponseBody)

	out, _, err := newTestClient(server.URL).LdapSettingsApi.GetLdapSettings(context.Background())

	require.NoError(t, err)
	require.NotNil(t, out.LdapSettings)
	require.NotNil(t, out.LdapSettings.LdapConnectionSettings)
	assert.Equal(t, "ldap.example.com", out.LdapSettings.LdapConnectionSettings.Host)
	assert.Equal(t, int32(636), out.LdapSettings.LdapConnectionSettings.Port)

	assert.Equal(t, http.MethodGet, lastReq.method)
	assert.Equal(t, "/v1/ldap-settings", lastReq.path)
	assert.Empty(t, lastReq.body)
}

func TestGetLdapSettingsByIdentifier_WireFormat(t *testing.T) {
	server, lastReq := serve(t, http.StatusOK, sampleResponseBody)

	_, _, err := newTestClient(server.URL).LdapSettingsApi.GetLdapSettingsByIdentifier(context.Background(), "corp_ldap")

	require.NoError(t, err)
	assert.Equal(t, http.MethodGet, lastReq.method)
	assert.Equal(t, "/v1/ldap-settings/corp_ldap", lastReq.path)
}

func TestUpdateLdapSettings_WireFormat(t *testing.T) {
	server, lastReq := serve(t, http.StatusOK, sampleResponseBody)

	_, _, err := newTestClient(server.URL).LdapSettingsApi.UpdateLdapSettings(context.Background(), "corp_ldap", sampleSettings())

	require.NoError(t, err)
	assert.Equal(t, http.MethodPut, lastReq.method)
	assert.Equal(t, "/v1/ldap-settings/corp_ldap", lastReq.path)

	var reqBody map[string]map[string]interface{}
	require.NoError(t, json.Unmarshal(lastReq.body, &reqBody))
	assert.Equal(t, "Corp LDAP", reqBody["ldap_settings"]["name"])
}

func TestDeleteLdapSettings_WireFormat(t *testing.T) {
	server, lastReq := serve(t, http.StatusNoContent, "")

	resp, err := newTestClient(server.URL).LdapSettingsApi.DeleteLdapSettings(context.Background(), "corp_ldap")

	require.NoError(t, err)
	assert.Equal(t, http.StatusNoContent, resp.StatusCode)
	assert.Equal(t, http.MethodDelete, lastReq.method)
	assert.Equal(t, "/v1/ldap-settings/corp_ldap", lastReq.path)
}

func TestGetIterations_WireFormat(t *testing.T) {
	server, lastReq := serve(t, http.StatusOK, `{"iterations":[1700000000000,1700003600000]}`)

	out, _, err := newTestClient(server.URL).LdapSettingsApi.GetIterations(context.Background(), "0 * * * *")

	require.NoError(t, err)
	assert.Equal(t, []int64{1700000000000, 1700003600000}, out.Iterations)

	assert.Equal(t, http.MethodPost, lastReq.method)
	assert.Equal(t, "/v1/ldap-settings/iterations", lastReq.path)
	var reqBody map[string]string
	require.NoError(t, json.Unmarshal(lastReq.body, &reqBody))
	assert.Equal(t, "0 * * * *", reqBody["cron"])
}

func TestDo_ErrorPaths(t *testing.T) {
	testCases := map[string]struct {
		status      int
		body        string
		wantMessage string
	}{
		"Given a 400 with a ResponseMessage body - then Message unwraps it": {
			status:      http.StatusBadRequest,
			body:        `{"message":"Ldap settings already exist for this account."}`,
			wantMessage: "Ldap settings already exist for this account.",
		},
		"Given a 404 with a ResponseMessage body - then Message unwraps it": {
			status:      http.StatusNotFound,
			body:        `{"message":"Ldap settings not found for account."}`,
			wantMessage: "Ldap settings not found for account.",
		},
		"Given a 500 with a non-JSON body - then Message falls back to the raw body": {
			status:      http.StatusInternalServerError,
			body:        `<html>proxy error</html>`,
			wantMessage: `<html>proxy error</html>`,
		},
	}

	for name, tc := range testCases {
		t.Run(name, func(t *testing.T) {
			server, _ := serve(t, tc.status, tc.body)

			resp, err := newTestClient(server.URL).LdapSettingsApi.DeleteLdapSettings(context.Background(), "x")

			require.NotNil(t, resp)
			assert.Equal(t, tc.status, resp.StatusCode)
			require.Error(t, err)
			genErr, ok := err.(GenericError)
			require.True(t, ok, "expected GenericError, got %T", err)
			assert.Equal(t, tc.wantMessage, genErr.Message())
			assert.Equal(t, tc.body, string(genErr.Body()))
			assert.Contains(t, genErr.Error(), http.StatusText(tc.status))
		})
	}
}

func TestDo_SuccessDecodeError(t *testing.T) {
	server, _ := serve(t, http.StatusOK, `not-json`)

	_, _, err := newTestClient(server.URL).LdapSettingsApi.GetLdapSettings(context.Background())

	require.Error(t, err, "an undecodable 2xx body must surface as an error, not a zero-value result")
}
