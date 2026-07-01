package split_test

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/harness/harness-go-sdk/harness/split"
	"github.com/stretchr/testify/require"
)

func TestSegmentsService_List(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		require.Equal(t, "/internal/api/v2/segments/ws/ws-1", r.URL.Path)
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		_, _ = w.Write([]byte(`{"objects":[{"name":"segment-a","description":"Desc"}],"totalCount":1}`))
	}))
	defer server.Close()

	cfg := split.NewDefaultConfiguration()
	cfg.BasePath = server.URL
	client := split.NewAPIClient(cfg)

	result, err := client.Segments.List("ws-1")
	require.NoError(t, err)
	require.NotNil(t, result)
	require.Len(t, result.Objects, 1)
	require.Equal(t, "segment-a", result.Objects[0].Name)
	require.Equal(t, 1, result.TotalCount)
}

func TestSegmentsService_Get(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		require.Equal(t, "/internal/api/v2/segments/ws/ws-1/segment-a", r.URL.Path)
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		_, _ = w.Write([]byte(`{"name":"segment-a","description":"My segment"}`))
	}))
	defer server.Close()

	cfg := split.NewDefaultConfiguration()
	cfg.BasePath = server.URL
	client := split.NewAPIClient(cfg)

	seg, err := client.Segments.Get("ws-1", "segment-a")
	require.NoError(t, err)
	require.NotNil(t, seg)
	require.Equal(t, "segment-a", seg.Name)
	require.Equal(t, "My segment", seg.Description)
}

func TestSegmentsService_Create(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		require.Equal(t, http.MethodPost, r.Method)
		require.Equal(t, "/internal/api/v2/segments/ws/ws-1/trafficTypes/tt-1", r.URL.Path)
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusCreated)
		_, _ = w.Write([]byte(`{"name":"new-segment","description":"New"}`))
	}))
	defer server.Close()

	cfg := split.NewDefaultConfiguration()
	cfg.BasePath = server.URL
	client := split.NewAPIClient(cfg)

	seg, err := client.Segments.Create("ws-1", "tt-1", split.SegmentRequest{Name: "new-segment", Description: "New"})
	require.NoError(t, err)
	require.NotNil(t, seg)
	require.Equal(t, "new-segment", seg.Name)
}

func TestSegmentsService_Delete(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		require.Equal(t, http.MethodDelete, r.Method)
		require.Equal(t, "/internal/api/v2/segments/ws/ws-1/segment-a", r.URL.Path)
		w.WriteHeader(http.StatusNoContent)
	}))
	defer server.Close()

	cfg := split.NewDefaultConfiguration()
	cfg.BasePath = server.URL
	client := split.NewAPIClient(cfg)

	err := client.Segments.Delete("ws-1", "segment-a")
	require.NoError(t, err)
}

func TestSegmentsService_Activate(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		require.Equal(t, http.MethodPost, r.Method)
		require.Equal(t, "/internal/api/v2/segments/env-1/segment-a", r.URL.Path)
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		_, _ = w.Write([]byte(`{"name":"segment-a"}`))
	}))
	defer server.Close()

	cfg := split.NewDefaultConfiguration()
	cfg.BasePath = server.URL
	client := split.NewAPIClient(cfg)

	seg, err := client.Segments.Activate("env-1", "segment-a")
	require.NoError(t, err)
	require.NotNil(t, seg)
	require.Equal(t, "segment-a", seg.Name)
}

func TestSegmentsService_Deactivate(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		require.Equal(t, http.MethodDelete, r.Method)
		require.Equal(t, "/internal/api/v2/segments/env-1/segment-a", r.URL.Path)
		w.WriteHeader(http.StatusNoContent)
	}))
	defer server.Close()

	cfg := split.NewDefaultConfiguration()
	cfg.BasePath = server.URL
	client := split.NewAPIClient(cfg)

	err := client.Segments.Deactivate("env-1", "segment-a")
	require.NoError(t, err)
}
