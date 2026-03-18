package split_test

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/harness/harness-go-sdk/harness/split"
	"github.com/stretchr/testify/require"
)

func TestWorkspacesService_List(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		require.Equal(t, "/internal/api/v2/workspaces", r.URL.Path)
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		_, _ = w.Write([]byte(`{"objects":[{"id":"ws-1","name":"Default","type":"workspace","organizationIdentifier":"org1","projectIdentifier":"proj1","requiresTitleAndComments":false}],"offset":0,"limit":20,"totalCount":1}`))
	}))
	defer server.Close()

	cfg := split.NewDefaultConfiguration()
	cfg.BasePath = server.URL
	client := split.NewAPIClient(cfg)

	ws, err := client.Workspaces.List(nil)
	require.NoError(t, err)
	require.Len(t, ws, 1)
	require.Equal(t, "ws-1", ws[0].ID)
	require.Equal(t, "Default", ws[0].Name)
	require.Equal(t, "org1", ws[0].OrganizationIdentifier)
	require.Equal(t, "proj1", ws[0].ProjectIdentifier)
}

func TestWorkspacesService_FindByOrganizationAndProject(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		require.Equal(t, "org1", r.URL.Query().Get("organizationIdentifier"))
		require.Equal(t, "proj1", r.URL.Query().Get("projectIdentifier"))
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		_, _ = w.Write([]byte(`{"objects":[{"id":"resolved-ws-id","name":"MyWorkspace","type":"workspace","organizationIdentifier":"org1","projectIdentifier":"proj1","requiresTitleAndComments":false}],"offset":0,"limit":20,"totalCount":1}`))
	}))
	defer server.Close()

	cfg := split.NewDefaultConfiguration()
	cfg.BasePath = server.URL
	client := split.NewAPIClient(cfg)

	ws, err := client.Workspaces.FindByOrganizationAndProject("org1", "proj1")
	require.NoError(t, err)
	require.Len(t, ws, 1)
	require.Equal(t, "resolved-ws-id", ws[0].ID)
}

func TestWorkspacesService_ResolveWorkspaceID(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		_, _ = w.Write([]byte(`{"objects":[{"id":"resolved-id","name":"P","type":"workspace","organizationIdentifier":"o","projectIdentifier":"p","requiresTitleAndComments":false}],"offset":0,"limit":20,"totalCount":1}`))
	}))
	defer server.Close()

	cfg := split.NewDefaultConfiguration()
	cfg.BasePath = server.URL
	client := split.NewAPIClient(cfg)

	id, err := client.Workspaces.ResolveWorkspaceID("o", "p")
	require.NoError(t, err)
	require.Equal(t, "resolved-id", id)
}

func TestWorkspacesService_ResolveWorkspaceID_NotFound(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		_, _ = w.Write([]byte(`{"objects":[],"offset":0,"limit":20,"totalCount":0}`))
	}))
	defer server.Close()

	cfg := split.NewDefaultConfiguration()
	cfg.BasePath = server.URL
	client := split.NewAPIClient(cfg)

	_, err := client.Workspaces.ResolveWorkspaceID("org", "proj")
	require.Error(t, err)
	require.Contains(t, err.Error(), "no workspace found")
}
