package split_test

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/harness/harness-go-sdk/harness/split"
	"github.com/stretchr/testify/require"
)

func TestFlagSetsService_List(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		require.Equal(t, "/api/v3/flag-sets", r.URL.Path)
		require.Equal(t, "ws-1", r.URL.Query().Get("workspace_id"))
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		_, _ = w.Write([]byte(`{"objects":[{"id":"fs-1","name":"My Flags","description":"Group"},{"id":"fs-2","name":"Other"}],"nextMarker":null}`))
	}))
	defer server.Close()

	cfg := split.NewDefaultConfiguration()
	cfg.BasePath = server.URL
	client := split.NewAPIClient(cfg)

	list, err := client.FlagSets.List("ws-1")
	require.NoError(t, err)
	require.Len(t, list, 2)
	require.Equal(t, "fs-1", list[0].ID)
	require.Equal(t, "My Flags", list[0].Name)
	require.Equal(t, "fs-2", list[1].ID)
}

func TestFlagSetsService_FindByID(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		require.Equal(t, "/api/v3/flag-sets/fs-1", r.URL.Path)
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		_, _ = w.Write([]byte(`{"id":"fs-1","name":"My Flags","description":"Desc"}`))
	}))
	defer server.Close()

	cfg := split.NewDefaultConfiguration()
	cfg.BasePath = server.URL
	client := split.NewAPIClient(cfg)

	fs, err := client.FlagSets.FindByID("fs-1")
	require.NoError(t, err)
	require.NotNil(t, fs)
	require.Equal(t, "fs-1", fs.ID)
	require.Equal(t, "My Flags", fs.Name)
}

func TestFlagSetsService_FindByName(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		_, _ = w.Write([]byte(`{"objects":[{"id":"fs-1","name":"Target"}],"nextMarker":null}`))
	}))
	defer server.Close()

	cfg := split.NewDefaultConfiguration()
	cfg.BasePath = server.URL
	client := split.NewAPIClient(cfg)

	fs, err := client.FlagSets.FindByName("ws-1", "Target")
	require.NoError(t, err)
	require.NotNil(t, fs)
	require.Equal(t, "fs-1", fs.ID)
	require.Equal(t, "Target", fs.Name)
}

func TestFlagSetsService_Create(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		require.Equal(t, http.MethodPost, r.Method)
		require.Equal(t, "/api/v3/flag-sets", r.URL.Path)
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusCreated)
		_, _ = w.Write([]byte(`{"id":"fs-new","name":"New Set","description":""}`))
	}))
	defer server.Close()

	cfg := split.NewDefaultConfiguration()
	cfg.BasePath = server.URL
	client := split.NewAPIClient(cfg)

	fs, err := client.FlagSets.Create(split.FlagSetRequest{
		Name:        "New Set",
		Description: "",
		Workspace:   &split.WorkspaceIDRef{Type: "workspace", ID: "ws-1"},
	})
	require.NoError(t, err)
	require.NotNil(t, fs)
	require.Equal(t, "fs-new", fs.ID)
	require.Equal(t, "New Set", fs.Name)
}

func TestFlagSetsService_Delete(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		require.Equal(t, http.MethodDelete, r.Method)
		require.Equal(t, "/api/v3/flag-sets/fs-1", r.URL.Path)
		w.WriteHeader(http.StatusNoContent)
	}))
	defer server.Close()

	cfg := split.NewDefaultConfiguration()
	cfg.BasePath = server.URL
	client := split.NewAPIClient(cfg)

	err := client.FlagSets.Delete("fs-1")
	require.NoError(t, err)
}
