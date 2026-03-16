package split_test

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/harness/harness-go-sdk/harness/split"
	"github.com/stretchr/testify/require"
)

func TestSplitsService_List(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		require.Equal(t, "/internal/api/v2/splits/ws/ws-1", r.URL.Path)
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		_, _ = w.Write([]byte(`{"objects":[{"id":"split-1","name":"feature_x","description":"Feature X"}],"totalCount":1}`))
	}))
	defer server.Close()

	cfg := split.NewDefaultConfiguration()
	cfg.BasePath = server.URL
	client := split.NewAPIClient(cfg)

	result, err := client.Splits.List("ws-1")
	require.NoError(t, err)
	require.NotNil(t, result)
	require.Len(t, result.Objects, 1)
	require.Equal(t, "split-1", result.Objects[0].ID)
	require.Equal(t, "feature_x", result.Objects[0].Name)
}

func TestSplitsService_Get(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		require.Equal(t, "/internal/api/v2/splits/ws/ws-1/feature_x", r.URL.Path)
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		_, _ = w.Write([]byte(`{"id":"split-1","name":"feature_x","description":"Toggle"}`))
	}))
	defer server.Close()

	cfg := split.NewDefaultConfiguration()
	cfg.BasePath = server.URL
	client := split.NewAPIClient(cfg)

	s, err := client.Splits.Get("ws-1", "feature_x")
	require.NoError(t, err)
	require.NotNil(t, s)
	require.Equal(t, "split-1", s.ID)
	require.Equal(t, "feature_x", s.Name)
}

func TestSplitsService_Create(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		require.Equal(t, http.MethodPost, r.Method)
		require.Equal(t, "/internal/api/v2/splits/ws/ws-1/trafficTypes/tt-1", r.URL.Path)
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusCreated)
		_, _ = w.Write([]byte(`{"id":"split-new","name":"new_flag","description":"New"}`))
	}))
	defer server.Close()

	cfg := split.NewDefaultConfiguration()
	cfg.BasePath = server.URL
	client := split.NewAPIClient(cfg)

	s, err := client.Splits.Create("ws-1", "tt-1", split.SplitCreateRequest{Name: "new_flag", Description: "New"})
	require.NoError(t, err)
	require.NotNil(t, s)
	require.Equal(t, "split-new", s.ID)
	require.Equal(t, "new_flag", s.Name)
}

func TestSplitsService_Delete(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		require.Equal(t, http.MethodDelete, r.Method)
		require.Contains(t, r.URL.Path, "/splits/ws/ws-1/")
		require.Contains(t, r.URL.Path, "feature_x")
		w.WriteHeader(http.StatusNoContent)
	}))
	defer server.Close()

	cfg := split.NewDefaultConfiguration()
	cfg.BasePath = server.URL
	client := split.NewAPIClient(cfg)

	err := client.Splits.Delete("ws-1", "feature_x")
	require.NoError(t, err)
}

func TestSplitsService_ListDefinitions(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		require.Equal(t, "/internal/api/v2/splits/ws/ws-1/environments/env-1", r.URL.Path)
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		_, _ = w.Write([]byte(`{"objects":[{"id":"def-1","name":"feature_x","defaultTreatment":"on","trafficAllocation":100}],"totalCount":1}`))
	}))
	defer server.Close()

	cfg := split.NewDefaultConfiguration()
	cfg.BasePath = server.URL
	client := split.NewAPIClient(cfg)

	result, err := client.Splits.ListDefinitions("ws-1", "env-1")
	require.NoError(t, err)
	require.NotNil(t, result)
	require.Len(t, result.Objects, 1)
	require.Equal(t, "feature_x", result.Objects[0].Name)
	require.Equal(t, "on", result.Objects[0].DefaultTreatment)
}

func TestSplitsService_GetDefinition(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		require.Equal(t, "/internal/api/v2/splits/ws/ws-1/feature_x/environments/env-1", r.URL.Path)
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		_, _ = w.Write([]byte(`{"id":"def-1","name":"feature_x","defaultTreatment":"off","trafficAllocation":100}`))
	}))
	defer server.Close()

	cfg := split.NewDefaultConfiguration()
	cfg.BasePath = server.URL
	client := split.NewAPIClient(cfg)

	def, err := client.Splits.GetDefinition("ws-1", "feature_x", "env-1")
	require.NoError(t, err)
	require.NotNil(t, def)
	require.Equal(t, "feature_x", def.Name)
	require.Equal(t, "off", def.DefaultTreatment)
}
