package split_test

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/harness/harness-go-sdk/harness/split"
	"github.com/stretchr/testify/require"
)

func TestTrafficTypesService_List(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		require.Equal(t, "/internal/api/v2/trafficTypes/ws/ws-123", r.URL.Path)
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		_, _ = w.Write([]byte(`[{"id":"tt-1","name":"user","displayAttributeId":""},{"id":"tt-2","name":"account","displayAttributeId":"attr"}]`))
	}))
	defer server.Close()

	cfg := split.NewDefaultConfiguration()
	cfg.BasePath = server.URL
	client := split.NewAPIClient(cfg)

	list, err := client.TrafficTypes.List("ws-123")
	require.NoError(t, err)
	require.Len(t, list, 2)
	require.Equal(t, "tt-1", list[0].ID)
	require.Equal(t, "user", list[0].Name)
	require.Equal(t, "tt-2", list[1].ID)
	require.Equal(t, "account", list[1].Name)
}

func TestTrafficTypesService_FindByName(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		_, _ = w.Write([]byte(`[{"id":"tt-1","name":"user","displayAttributeId":""}]`))
	}))
	defer server.Close()

	cfg := split.NewDefaultConfiguration()
	cfg.BasePath = server.URL
	client := split.NewAPIClient(cfg)

	tt, err := client.TrafficTypes.FindByName("ws-123", "user")
	require.NoError(t, err)
	require.NotNil(t, tt)
	require.Equal(t, "tt-1", tt.ID)
	require.Equal(t, "user", tt.Name)

	tt2, err := client.TrafficTypes.FindByName("ws-123", "nonexistent")
	require.NoError(t, err)
	require.Nil(t, tt2)
}

func TestTrafficTypesService_FindByID(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		_, _ = w.Write([]byte(`[{"id":"tt-1","name":"user","displayAttributeId":""}]`))
	}))
	defer server.Close()

	cfg := split.NewDefaultConfiguration()
	cfg.BasePath = server.URL
	client := split.NewAPIClient(cfg)

	tt, err := client.TrafficTypes.FindByID("ws-123", "tt-1")
	require.NoError(t, err)
	require.NotNil(t, tt)
	require.Equal(t, "tt-1", tt.ID)
}

func TestTrafficTypesService_Delete(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		require.Equal(t, http.MethodDelete, r.Method)
		// Delete uses /trafficTypes/{id} only (no workspace in path), per Split API / terraform-provider-split
		require.Equal(t, "/internal/api/v2/trafficTypes/tt-1", r.URL.Path)
		w.WriteHeader(http.StatusNoContent)
	}))
	defer server.Close()

	cfg := split.NewDefaultConfiguration()
	cfg.BasePath = server.URL
	client := split.NewAPIClient(cfg)

	err := client.TrafficTypes.Delete("ws-123", "tt-1")
	require.NoError(t, err)
}
