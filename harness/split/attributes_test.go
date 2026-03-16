package split_test

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/harness/harness-go-sdk/harness/split"
	"github.com/stretchr/testify/require"
)

func TestAttributesService_List(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		require.Equal(t, "/internal/api/v2/schema/ws/ws-1/trafficTypes/tt-1", r.URL.Path)
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		_, _ = w.Write([]byte(`[{"id":"attr-1","displayName":"Region","dataType":"string","isSearchable":true}]`))
	}))
	defer server.Close()

	cfg := split.NewDefaultConfiguration()
	cfg.BasePath = server.URL
	client := split.NewAPIClient(cfg)

	list, err := client.Attributes.List("ws-1", "tt-1", nil)
	require.NoError(t, err)
	require.Len(t, list, 1)
	require.Equal(t, "attr-1", list[0].ID)
	require.Equal(t, "Region", list[0].DisplayName)
	require.True(t, list[0].IsSearchable)
}

func TestAttributesService_FindByID(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		_, _ = w.Write([]byte(`[{"id":"attr-1","displayName":"Region"}]`))
	}))
	defer server.Close()

	cfg := split.NewDefaultConfiguration()
	cfg.BasePath = server.URL
	client := split.NewAPIClient(cfg)

	attr, err := client.Attributes.FindByID("ws-1", "tt-1", "attr-1", nil)
	require.NoError(t, err)
	require.NotNil(t, attr)
	require.Equal(t, "attr-1", attr.ID)
}

func TestAttributesService_Create(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		require.Equal(t, http.MethodPost, r.Method)
		require.Equal(t, "/internal/api/v2/schema/ws/ws-1/trafficTypes/tt-1", r.URL.Path)
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusCreated)
		_, _ = w.Write([]byte(`{"id":"attr-new","displayName":"Country","dataType":"string"}`))
	}))
	defer server.Close()

	cfg := split.NewDefaultConfiguration()
	cfg.BasePath = server.URL
	client := split.NewAPIClient(cfg)

	attr, err := client.Attributes.Create("ws-1", "tt-1", split.AttributeRequest{
		DisplayName: "Country",
		DataType:    "string",
	})
	require.NoError(t, err)
	require.NotNil(t, attr)
	require.Equal(t, "attr-new", attr.ID)
	require.Equal(t, "Country", attr.DisplayName)
}

func TestAttributesService_Delete(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		require.Equal(t, http.MethodDelete, r.Method)
		require.Contains(t, r.URL.Path, "/schema/ws/ws-1/trafficTypes/tt-1/")
		w.WriteHeader(http.StatusNoContent)
	}))
	defer server.Close()

	cfg := split.NewDefaultConfiguration()
	cfg.BasePath = server.URL
	client := split.NewAPIClient(cfg)

	err := client.Attributes.Delete("ws-1", "tt-1", "attr-1")
	require.NoError(t, err)
}
