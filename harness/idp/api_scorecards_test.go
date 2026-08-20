package idp

import (
	"context"
	"encoding/json"
	"io"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/antihax/optional"
	"github.com/stretchr/testify/require"
)

func TestScorecardAndCheckCRUDIncludeHarnessAccount(t *testing.T) {
	const accountID = "account-123"

	tests := []struct {
		name       string
		method     string
		path       string
		statusCode int
		invoke     func(*APIClient) error
	}{
		{
			name:       "CreateCheck",
			method:     http.MethodPost,
			path:       "/v1/checks",
			statusCode: http.StatusCreated,
			invoke: func(client *APIClient) error {
				_, _, err := client.ChecksApi.CreateCheck(context.Background(), CheckRequest{
					CheckDetails: &CheckDetails{Identifier: "readme", Name: "README exists"},
				}, &ChecksApiCreateCheckOpts{HarnessAccount: optional.NewString(accountID)})
				return err
			},
		},
		{
			name:       "GetCheck",
			method:     http.MethodGet,
			path:       "/v1/checks/readme",
			statusCode: http.StatusOK,
			invoke: func(client *APIClient) error {
				_, _, err := client.ChecksApi.GetCheck(context.Background(), "readme", &ChecksApiGetCheckOpts{
					HarnessAccount: optional.NewString(accountID),
					Custom:         optional.NewBool(true),
				})
				return err
			},
		},
		{
			name:       "UpdateCheck",
			method:     http.MethodPut,
			path:       "/v1/checks/readme",
			statusCode: http.StatusOK,
			invoke: func(client *APIClient) error {
				_, _, err := client.ChecksApi.UpdateCheck(context.Background(), CheckRequest{
					CheckDetails: &CheckDetails{Identifier: "readme", Name: "README exists updated"},
				}, "readme", &ChecksApiUpdateCheckOpts{HarnessAccount: optional.NewString(accountID)})
				return err
			},
		},
		{
			name:       "DeleteCheck",
			method:     http.MethodDelete,
			path:       "/v1/checks/readme",
			statusCode: http.StatusNoContent,
			invoke: func(client *APIClient) error {
				_, err := client.ChecksApi.DeleteCheck(context.Background(), "readme", &ChecksApiDeleteCheckOpts{
					HarnessAccount: optional.NewString(accountID),
				})
				return err
			},
		},
		{
			name:       "CreateScorecard",
			method:     http.MethodPost,
			path:       "/v1/scorecards",
			statusCode: http.StatusCreated,
			invoke: func(client *APIClient) error {
				_, _, err := client.ScorecardsApi.CreateScorecard(context.Background(), ScorecardRequest{
					Scorecard: Scorecard{Identifier: "gold", Name: "Gold"},
					Checks:    []ScorecardCheck{{Identifier: "readme", Custom: true}},
				}, &ScorecardsApiCreateScorecardOpts{HarnessAccount: optional.NewString(accountID)})
				return err
			},
		},
		{
			name:       "GetScorecard",
			method:     http.MethodGet,
			path:       "/v1/scorecards/gold",
			statusCode: http.StatusOK,
			invoke: func(client *APIClient) error {
				_, _, err := client.ScorecardsApi.GetScorecard(context.Background(), "gold", &ScorecardsApiGetScorecardOpts{
					HarnessAccount: optional.NewString(accountID),
				})
				return err
			},
		},
		{
			name:       "UpdateScorecard",
			method:     http.MethodPut,
			path:       "/v1/scorecards/gold",
			statusCode: http.StatusOK,
			invoke: func(client *APIClient) error {
				_, _, err := client.ScorecardsApi.UpdateScorecard(context.Background(), ScorecardRequest{
					Scorecard: Scorecard{Identifier: "gold", Name: "Gold Updated"},
					Checks:    []ScorecardCheck{{Identifier: "readme", Custom: true}},
				}, "gold", &ScorecardsApiUpdateScorecardOpts{HarnessAccount: optional.NewString(accountID)})
				return err
			},
		},
		{
			name:       "DeleteScorecard",
			method:     http.MethodDelete,
			path:       "/v1/scorecards/gold",
			statusCode: http.StatusNoContent,
			invoke: func(client *APIClient) error {
				_, err := client.ScorecardsApi.DeleteScorecard(context.Background(), "gold", &ScorecardsApiDeleteScorecardOpts{
					HarnessAccount: optional.NewString(accountID),
				})
				return err
			},
		},
		{
			name:       "GetChecks",
			method:     http.MethodGet,
			path:       "/v1/checks",
			statusCode: http.StatusOK,
			invoke: func(client *APIClient) error {
				_, _, err := client.ChecksApi.GetChecks(context.Background(), &ChecksApiGetChecksOpts{
					HarnessAccount: optional.NewString(accountID),
					Custom:         optional.NewBool(true),
				})
				return err
			},
		},
		{
			name:       "GetScorecards",
			method:     http.MethodGet,
			path:       "/v1/scorecards",
			statusCode: http.StatusOK,
			invoke: func(client *APIClient) error {
				_, _, err := client.ScorecardsApi.GetScorecards(context.Background(), &ScorecardsApiGetScorecardsOpts{
					HarnessAccount: optional.NewString(accountID),
				})
				return err
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
				require.Equal(t, tt.method, r.Method)
				require.Equal(t, tt.path, r.URL.Path)
				require.Equal(t, accountID, r.Header.Get("Harness-Account"))
				require.Equal(t, accountID, r.URL.Query().Get("accountIdentifier"))

				w.Header().Set("Content-Type", "application/json")
				w.WriteHeader(tt.statusCode)
				switch r.URL.Path {
				case "/v1/checks/readme":
					if r.Method == http.MethodGet {
						_, _ = w.Write([]byte(`{"check_details":{"identifier":"readme","name":"README exists"}}`))
						return
					}
				case "/v1/scorecards/gold":
					if r.Method == http.MethodGet {
						_, _ = w.Write([]byte(`{"scorecard":{"identifier":"gold","name":"Gold"},"checks":[{"identifier":"readme","custom":true}]}`))
						return
					}
				case "/v1/checks":
					if r.Method == http.MethodGet {
						_, _ = w.Write([]byte(`[{"check":{"identifier":"readme","name":"README exists"}}]`))
						return
					}
				case "/v1/scorecards":
					if r.Method == http.MethodGet {
						_, _ = w.Write([]byte(`[{"scorecard":{"identifier":"gold","name":"Gold"}}]`))
						return
					}
				}
				_, _ = w.Write([]byte(`{"status":"SUCCESS"}`))
			}))
			defer server.Close()

			cfg := NewConfiguration()
			cfg.BasePath = server.URL
			client := NewAPIClient(cfg)

			require.NoError(t, tt.invoke(client))
		})
	}
}

func TestCreateScorecardSendsBody(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		body, err := io.ReadAll(r.Body)
		require.NoError(t, err)
		var req ScorecardRequest
		require.NoError(t, json.Unmarshal(body, &req))
		require.Equal(t, "gold", req.Scorecard.Identifier)
		require.Equal(t, "component", req.Scorecard.Filter.Kind)
		require.Len(t, req.Checks, 1)
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusCreated)
		_, _ = w.Write([]byte(`{"status":"SUCCESS"}`))
	}))
	defer server.Close()

	cfg := NewConfiguration()
	cfg.BasePath = server.URL
	client := NewAPIClient(cfg)

	_, _, err := client.ScorecardsApi.CreateScorecard(context.Background(), ScorecardRequest{
		Scorecard: Scorecard{
			Identifier: "gold",
			Name:       "Gold",
			Filter:     &ScorecardFilter{Kind: "component", Type_: "service"},
		},
		Checks: []ScorecardCheck{{Identifier: "readme", Custom: true, Weightage: 1}},
	}, nil)
	require.NoError(t, err)
}

func TestDeleteCheckSendsForceDelete(t *testing.T) {
	const accountID = "account-123"
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		require.Equal(t, http.MethodDelete, r.Method)
		require.Equal(t, "/v1/checks/readme", r.URL.Path)
		require.Equal(t, "true", r.URL.Query().Get("force_delete"))
		require.Equal(t, accountID, r.Header.Get("Harness-Account"))
		w.WriteHeader(http.StatusNoContent)
	}))
	defer server.Close()

	cfg := NewConfiguration()
	cfg.BasePath = server.URL
	client := NewAPIClient(cfg)

	_, err := client.ChecksApi.DeleteCheck(context.Background(), "readme", &ChecksApiDeleteCheckOpts{
		HarnessAccount: optional.NewString(accountID),
		ForceDelete:    optional.NewBool(true),
	})
	require.NoError(t, err)
}

func TestDeleteCheckDefaultsForceDeleteFalse(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		require.Equal(t, "false", r.URL.Query().Get("force_delete"))
		w.WriteHeader(http.StatusNoContent)
	}))
	defer server.Close()

	cfg := NewConfiguration()
	cfg.BasePath = server.URL
	client := NewAPIClient(cfg)

	_, err := client.ChecksApi.DeleteCheck(context.Background(), "readme", nil)
	require.NoError(t, err)
}
