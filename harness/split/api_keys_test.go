package split_test

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/harness/harness-go-sdk/harness/split"
	"github.com/stretchr/testify/require"
)

func TestApiKeysService_Create(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		require.Equal(t, http.MethodPost, r.Method)
		require.Equal(t, "/internal/api/v2/apiKeys", r.URL.Path)
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusCreated)
		_, _ = w.Write([]byte(`{"id":"key-1","name":"My Key","type":"admin","apiKeyType":"admin","key":"secret-key-value","roles":["API_ALL_GRANTED"]}`))
	}))
	defer server.Close()

	cfg := split.NewDefaultConfiguration()
	cfg.BasePath = server.URL
	client := split.NewAPIClient(cfg)

	resp, err := client.ApiKeys.Create(split.KeyRequest{
		Name:  "My Key",
		KeyType: "admin",
		Roles: []string{"API_ALL_GRANTED"},
		Environments: []split.KeyEnvironmentRequest{{Type: "environment", Id: "env-1"}},
		Workspace:    &split.KeyWorkspaceRequest{Type: "workspace", Id: "ws-1"},
	})
	require.NoError(t, err)
	require.NotNil(t, resp)
	require.Equal(t, "key-1", resp.Id)
	require.Equal(t, "My Key", resp.Name)
	require.Equal(t, "secret-key-value", resp.Key)
}

func TestApiKeysService_Delete(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		require.Equal(t, http.MethodDelete, r.Method)
		require.Equal(t, "/internal/api/v2/apiKeys/key-1", r.URL.Path)
		w.WriteHeader(http.StatusNoContent)
	}))
	defer server.Close()

	cfg := split.NewDefaultConfiguration()
	cfg.BasePath = server.URL
	client := split.NewAPIClient(cfg)

	err := client.ApiKeys.Delete("key-1")
	require.NoError(t, err)
}
