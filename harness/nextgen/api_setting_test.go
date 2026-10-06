package nextgen

import (
	"context"
	"encoding/json"
	"io"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/antihax/optional"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestSettingApiGetSettingValue(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		assert.Equal(t, http.MethodGet, r.Method)
		assert.Equal(t, "/ng/api/settings/default_repo_for_git_experience", r.URL.Path)
		assert.Equal(t, "acc", r.URL.Query().Get("accountIdentifier"))
		assert.Equal(t, "org", r.URL.Query().Get("orgIdentifier"))
		assert.Equal(t, "proj", r.URL.Query().Get("projectIdentifier"))
		w.Header().Set("Content-Type", "application/json")
		_, _ = w.Write([]byte(`{"status":"SUCCESS","data":{"valueType":"String","value":"my-repo"}}`))
	}))
	defer server.Close()

	resp, _, err := newTestClient(server.URL).SettingApi.GetSettingValue(context.Background(), "default_repo_for_git_experience", "acc", &SettingApiGetSettingValueOpts{
		OrgIdentifier:     optional.NewString("org"),
		ProjectIdentifier: optional.NewString("proj"),
	})
	require.NoError(t, err)
	assert.Equal(t, "String", resp.Data.ValueType)
	assert.Equal(t, "my-repo", resp.Data.Value)
}

func TestSettingApiGetSettingsList(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		assert.Equal(t, http.MethodGet, r.Method)
		assert.Equal(t, "/ng/api/settings", r.URL.Path)
		assert.Equal(t, "GIT_EXPERIENCE", r.URL.Query().Get("category"))
		assert.Equal(t, "git_experience", r.URL.Query().Get("group"))
		assert.Equal(t, "true", r.URL.Query().Get("includeParentScopes"))
		w.Header().Set("Content-Type", "application/json")
		_, _ = w.Write([]byte(`{"status":"SUCCESS","data":[{"setting":{"identifier":"default_repo_for_git_experience","category":"GIT_EXPERIENCE","valueType":"String","allowOverrides":false,"value":"my-repo","settingSource":"ORG"},"lastModifiedAt":1000}]}`))
	}))
	defer server.Close()

	resp, _, err := newTestClient(server.URL).SettingApi.GetSettingsList(context.Background(), "acc", "GIT_EXPERIENCE", &SettingApiGetSettingsListOpts{
		OrgIdentifier:       optional.NewString("org"),
		Group:               optional.NewString("git_experience"),
		IncludeParentScopes: optional.NewBool(true),
	})
	require.NoError(t, err)
	require.Len(t, resp.Data, 1)
	assert.Equal(t, "default_repo_for_git_experience", resp.Data[0].Setting.Identifier)
	assert.Equal(t, "ORG", resp.Data[0].Setting.SettingSource)
	assert.False(t, resp.Data[0].Setting.AllowOverrides)
	assert.Equal(t, int64(1000), resp.Data[0].LastModifiedAt)
}

func TestSettingApiUpdateSettingValue(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		assert.Equal(t, http.MethodPut, r.Method)
		assert.Equal(t, "/ng/api/settings", r.URL.Path)
		body, _ := io.ReadAll(r.Body)
		// allowOverrides is required by the API, so false must still be sent.
		assert.JSONEq(t, `[{"identifier":"default_repo_for_git_experience","value":"my-repo","allowOverrides":false,"updateType":"UPDATE"}]`, string(body))
		w.Header().Set("Content-Type", "application/json")
		_, _ = w.Write([]byte(`{"status":"SUCCESS","data":[{"identifier":"default_repo_for_git_experience","updateStatus":true,"setting":{"identifier":"default_repo_for_git_experience","value":"my-repo"}}]}`))
	}))
	defer server.Close()

	resp, _, err := newTestClient(server.URL).SettingApi.UpdateSettingValue(context.Background(), []SettingRequestDto{{
		Identifier:     "default_repo_for_git_experience",
		Value:          "my-repo",
		AllowOverrides: false,
		UpdateType:     "UPDATE",
	}}, "acc", nil)
	require.NoError(t, err)
	require.Len(t, resp.Data, 1)
	assert.True(t, resp.Data[0].UpdateStatus)
}

func TestSettingApiErrorResponse(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusBadRequest)
		_ = json.NewEncoder(w).Encode(Failure{Status: "FAILURE", Code: "INVALID_REQUEST", Message: "Invalid category"})
	}))
	defer server.Close()

	_, httpResp, err := newTestClient(server.URL).SettingApi.GetSettingsList(context.Background(), "acc", "BAD", nil)
	require.Error(t, err)
	assert.Equal(t, http.StatusBadRequest, httpResp.StatusCode)
	swaggerErr, ok := err.(GenericSwaggerError)
	require.True(t, ok)
	assert.Equal(t, "Invalid category", swaggerErr.Model().(Failure).Message)
}
