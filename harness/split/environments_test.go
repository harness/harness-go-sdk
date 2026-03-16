package split_test

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/harness/harness-go-sdk/harness/split"
	"github.com/stretchr/testify/require"
)

func TestEnvironmentsService_List(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		require.Equal(t, "/internal/api/v2/environments/ws/ws-1", r.URL.Path)
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		_, _ = w.Write([]byte(`[{"id":"env-1","name":"Production","production":true},{"id":"env-2","name":"Staging","production":false}]`))
	}))
	defer server.Close()

	cfg := split.NewDefaultConfiguration()
	cfg.BasePath = server.URL
	client := split.NewAPIClient(cfg)

	list, err := client.Environments.List("ws-1")
	require.NoError(t, err)
	require.Len(t, list, 2)
	require.Equal(t, "env-1", list[0].ID)
	require.Equal(t, "Production", list[0].Name)
	require.True(t, list[0].Production)
	require.Equal(t, "Staging", list[1].Name)
	require.False(t, list[1].Production)
}

func TestEnvironmentsService_FindByName(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		_, _ = w.Write([]byte(`[{"id":"env-1","name":"Production","production":true}]`))
	}))
	defer server.Close()

	cfg := split.NewDefaultConfiguration()
	cfg.BasePath = server.URL
	client := split.NewAPIClient(cfg)

	env, err := client.Environments.FindByName("ws-1", "Production")
	require.NoError(t, err)
	require.NotNil(t, env)
	require.Equal(t, "env-1", env.ID)
	require.Equal(t, "Production", env.Name)

	env2, err := client.Environments.FindByName("ws-1", "Nonexistent")
	require.NoError(t, err)
	require.Nil(t, env2)
}

func TestEnvironmentsService_Create(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		require.Equal(t, http.MethodPost, r.Method)
		require.Equal(t, "/internal/api/v2/environments/ws/ws-1", r.URL.Path)
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusCreated)
		_, _ = w.Write([]byte(`{"id":"env-new","name":"Dev","production":false}`))
	}))
	defer server.Close()

	cfg := split.NewDefaultConfiguration()
	cfg.BasePath = server.URL
	client := split.NewAPIClient(cfg)

	env, err := client.Environments.Create("ws-1", split.CreateEnvironmentRequest{Name: "Dev", Production: false})
	require.NoError(t, err)
	require.NotNil(t, env)
	require.Equal(t, "env-new", env.ID)
	require.Equal(t, "Dev", env.Name)
	require.False(t, env.Production)
}

func TestEnvironmentsService_ListSegmentsAll(t *testing.T) {
	callCount := 0
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		require.Contains(t, r.URL.Path, "/segments/ws/ws-1/environments/env-1")
		callCount++
		w.Header().Set("Content-Type", "application/json")
		// Page 1: offset=0, 2 items, total 3
		if callCount == 1 {
			w.WriteHeader(http.StatusOK)
			_, _ = w.Write([]byte(`{"objects":[{"id":"seg-a"},{"id":"seg-b"}],"totalCount":3}`))
			return
		}
		// Page 2: offset=2, 1 item
		w.WriteHeader(http.StatusOK)
		_, _ = w.Write([]byte(`{"objects":[{"id":"seg-c"}],"totalCount":3}`))
	}))
	defer server.Close()

	cfg := split.NewDefaultConfiguration()
	cfg.BasePath = server.URL
	client := split.NewAPIClient(cfg)

	all, err := client.Environments.ListSegmentsAll("ws-1", "env-1")
	require.NoError(t, err)
	require.Len(t, all, 3)
	require.Equal(t, []string{"seg-a", "seg-b", "seg-c"}, all)
	require.Equal(t, 2, callCount)
}

func TestEnvironmentsService_GetSegmentKeysAll(t *testing.T) {
	callCount := 0
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		require.Contains(t, r.URL.Path, "/segments/env-1/my_seg/keys")
		callCount++
		w.Header().Set("Content-Type", "application/json")
		if callCount == 1 {
			w.WriteHeader(http.StatusOK)
			_, _ = w.Write([]byte(`{"keys":[{"key":"k1"},{"key":"k2"}],"totalCount":3}`))
			return
		}
		w.WriteHeader(http.StatusOK)
		_, _ = w.Write([]byte(`{"keys":[{"key":"k3"}],"totalCount":3}`))
	}))
	defer server.Close()

	cfg := split.NewDefaultConfiguration()
	cfg.BasePath = server.URL
	client := split.NewAPIClient(cfg)

	all, err := client.Environments.GetSegmentKeysAll("ws-1", "env-1", "my_seg")
	require.NoError(t, err)
	require.Len(t, all, 3)
	require.Equal(t, []string{"k1", "k2", "k3"}, all)
	require.Equal(t, 2, callCount)
}
