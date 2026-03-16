package split

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestRuleBasedSegmentsService_List(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		require.Equal(t, http.MethodGet, r.Method)
		require.Equal(t, "/internal/api/v2/rule-based-segments/ws/ws-1", r.URL.Path)
		w.WriteHeader(http.StatusOK)
		_ = json.NewEncoder(w).Encode([]RuleBasedSegmentMetadata{
			{Name: "rbs-1", TrafficType: &TrafficType{ID: "tt-1", Name: "user"}},
		})
	}))
	defer server.Close()

	cfg := NewDefaultConfiguration()
	cfg.BasePath = server.URL
	client := NewAPIClient(cfg)

	list, err := client.RuleBasedSegments.List("ws-1")
	require.NoError(t, err)
	require.Len(t, list, 1)
	require.Equal(t, "rbs-1", list[0].Name)
}

func TestRuleBasedSegmentsService_Get(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		require.Equal(t, http.MethodGet, r.Method)
		require.Contains(t, r.URL.Path, "/rule-based-segments/ws/ws-1/")
		w.WriteHeader(http.StatusOK)
		_ = json.NewEncoder(w).Encode(RuleBasedSegmentDefinition{
			Title:   "Example",
			Comment: "Test",
			Rules:   []RuleBasedSegmentRule{},
		})
	}))
	defer server.Close()

	cfg := NewDefaultConfiguration()
	cfg.BasePath = server.URL
	client := NewAPIClient(cfg)

	def, err := client.RuleBasedSegments.Get("ws-1", "my_rbs")
	require.NoError(t, err)
	require.Equal(t, "Example", def.Title)
	require.Equal(t, "Test", def.Comment)
}

func TestRuleBasedSegmentsService_Create(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		require.Equal(t, http.MethodPost, r.Method)
		require.Equal(t, "/internal/api/v2/rule-based-segments/ws/ws-1/trafficTypes/tt-1", r.URL.Path)
		var req RuleBasedSegmentCreateRequest
		_ = json.NewDecoder(r.Body).Decode(&req)
		require.Equal(t, "new_rbs", req.Name)
		w.WriteHeader(http.StatusCreated)
		_ = json.NewEncoder(w).Encode(RuleBasedSegmentMetadata{Name: "new_rbs", TrafficType: &TrafficType{ID: "tt-1", Name: "user"}})
	}))
	defer server.Close()

	cfg := NewDefaultConfiguration()
	cfg.BasePath = server.URL
	client := NewAPIClient(cfg)

	meta, err := client.RuleBasedSegments.Create("ws-1", "tt-1", RuleBasedSegmentCreateRequest{Name: "new_rbs"})
	require.NoError(t, err)
	require.Equal(t, "new_rbs", meta.Name)
}
