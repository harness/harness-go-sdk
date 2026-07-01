package split

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestRuleBasedSegmentMetadata_unmarshalFMEListItem(t *testing.T) {
	raw := `{
		"id": "cf7f6903-d79b-11f0-b178-eaaf08872860",
		"workspaceIds": ["9306f600-c70b-11f0-9243-aa89d0eb095e"],
		"orgId": "4ec48f70-bbfd-11f0-88bf-02779340d5f4",
		"name": "qa_account_users",
		"status": "ACTIVE",
		"trafficTypeId": "7e042930-d79b-11f0-a575-5259b3e62aa3",
		"trafficTypeURN": null,
		"description": "",
		"creator": {"type": "User", "id": "0MR8DysRQ6WymBsEb_-t3A"},
		"creationTime": 1765572464962,
		"views": null,
		"favorite": null,
		"tags": null,
		"canEdit": null,
		"constraints": null
	}`
	var m RuleBasedSegmentMetadata
	require.NoError(t, json.Unmarshal([]byte(raw), &m))
	require.Equal(t, "cf7f6903-d79b-11f0-b178-eaaf08872860", m.ID)
	require.Equal(t, []string{"9306f600-c70b-11f0-9243-aa89d0eb095e"}, m.WorkspaceIDs)
	require.Equal(t, "4ec48f70-bbfd-11f0-88bf-02779340d5f4", m.OrgID)
	require.Equal(t, "qa_account_users", m.Name)
	require.Equal(t, "ACTIVE", m.Status)
	require.Equal(t, "7e042930-d79b-11f0-a575-5259b3e62aa3", m.TrafficTypeID)
	require.Nil(t, m.TrafficTypeURN)
	require.NotNil(t, m.Creator)
	require.Equal(t, "User", m.Creator.Type)
	require.Equal(t, "0MR8DysRQ6WymBsEb_-t3A", m.Creator.ID)
	require.Equal(t, int64(1765572464962), m.CreationTime)
}

func TestRuleBasedSegmentsService_List(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		require.Equal(t, http.MethodGet, r.Method)
		require.Equal(t, "/internal/api/v2/rule-based-segments/ws/ws-1", r.URL.Path)
		w.WriteHeader(http.StatusOK)
		_ = json.NewEncoder(w).Encode([]RuleBasedSegmentMetadata{
			{
				ID:            "rbs-id-1",
				Name:          "rbs-1",
				Status:        "ACTIVE",
				TrafficTypeID: "tt-1",
				TrafficType:   &TrafficType{ID: "tt-1", Name: "user"},
			},
		})
	}))
	defer server.Close()

	cfg := NewDefaultConfiguration()
	cfg.BasePath = server.URL
	client := NewAPIClient(cfg)

	list, err := client.RuleBasedSegments.List("ws-1")
	require.NoError(t, err)
	require.Len(t, list, 1)
	require.Equal(t, "rbs-id-1", list[0].ID)
	require.Equal(t, "rbs-1", list[0].Name)
	require.Equal(t, "tt-1", list[0].TrafficTypeID)
	require.NotNil(t, list[0].TrafficType)
	require.Equal(t, "user", list[0].TrafficType.Name)
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

func TestRuleBasedSegmentsService_ListInEnvironment(t *testing.T) {
	attr := "email"
	typ := "CONTAINS_STRING"
	str := "@beta.example.com"
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		require.Equal(t, http.MethodGet, r.Method)
		require.Equal(t, "/internal/api/v2/rule-based-segments/ws/ws-1/environments/env-1", r.URL.Path)
		w.WriteHeader(http.StatusOK)
		_ = json.NewEncoder(w).Encode(map[string]any{
			"objects": []map[string]any{
				{
					"id":          "seg_12345",
					"name":        "Beta Testers",
					"description": "Users in the beta testing group.",
					"title":       "Beta title",
					"rules": []map[string]any{
						{
							"condition": map[string]any{
								"combiner": "AND",
								"matchers": []map[string]any{
									{"attribute": attr, "type": typ, "string": str},
								},
							},
						},
					},
					"createdAt": "2025-09-25T18:03:00Z",
					"updatedAt": "2025-10-01T14:21:00Z",
				},
			},
		})
	}))
	defer server.Close()

	cfg := NewDefaultConfiguration()
	cfg.BasePath = server.URL
	client := NewAPIClient(cfg)

	list, err := client.RuleBasedSegments.ListInEnvironment("ws-1", "env-1")
	require.NoError(t, err)
	require.Len(t, list, 1)
	require.Equal(t, "seg_12345", list[0].ID)
	require.Equal(t, "Beta Testers", list[0].Name)
	require.Equal(t, "Users in the beta testing group.", list[0].Description)
	require.Equal(t, "Beta title", list[0].Title)
	require.Equal(t, "2025-09-25T18:03:00Z", list[0].CreatedAt)
	require.Equal(t, "2025-10-01T14:21:00Z", list[0].UpdatedAt)
	require.Len(t, list[0].Rules, 1)
	require.NotNil(t, list[0].Rules[0].Condition)
	require.Len(t, list[0].Rules[0].Condition.Matchers, 1)
	m := list[0].Rules[0].Condition.Matchers[0]
	require.NotNil(t, m.Attribute)
	require.Equal(t, "email", *m.Attribute)
	require.NotNil(t, m.String)
	require.Equal(t, "@beta.example.com", *m.String)
}

func TestRuleBasedSegmentsService_ListInEnvironment_bareArrayFallback(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
		_ = json.NewEncoder(w).Encode([]RuleBasedSegmentEnvironmentEntry{
			{Name: "only_name", RuleBasedSegmentDefinition: RuleBasedSegmentDefinition{Title: "t"}},
		})
	}))
	defer server.Close()

	cfg := NewDefaultConfiguration()
	cfg.BasePath = server.URL
	client := NewAPIClient(cfg)

	list, err := client.RuleBasedSegments.ListInEnvironment("ws-1", "env-1")
	require.NoError(t, err)
	require.Len(t, list, 1)
	require.Equal(t, "only_name", list[0].Name)
	require.Equal(t, "t", list[0].Title)
}
