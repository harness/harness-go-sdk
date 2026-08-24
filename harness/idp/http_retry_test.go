package idp

import (
	"bytes"
	"context"
	"io"
	"net/http"
	"net/http/httptest"
	"sync/atomic"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestGetCheckDoesNotRetryNotFoundMappedAs500(t *testing.T) {
	var hits atomic.Int32
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		hits.Add(1)
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusInternalServerError)
		_, _ = w.Write([]byte(`{"message":"Check details not found for checkId [readme]"}`))
	}))
	defer server.Close()

	cfg := NewConfiguration()
	cfg.BasePath = server.URL
	client := NewAPIClient(cfg)

	_, resp, err := client.ChecksApi.GetCheck(context.Background(), "readme", nil)
	require.Error(t, err)
	require.NotNil(t, resp)
	require.Equal(t, http.StatusInternalServerError, resp.StatusCode)
	require.Equal(t, int32(1), hits.Load())
}

func TestCreateCheckDoesNotRetryAlreadyCreatedMappedAs500(t *testing.T) {
	var hits atomic.Int32
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		hits.Add(1)
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusInternalServerError)
		_, _ = w.Write([]byte(`{"message":"Check [readme] already created for accountId [account-123]"}`))
	}))
	defer server.Close()

	cfg := NewConfiguration()
	cfg.BasePath = server.URL
	client := NewAPIClient(cfg)

	_, resp, err := client.ChecksApi.CreateCheck(context.Background(), CheckRequest{
		CheckDetails: &CheckDetails{Identifier: "readme", Name: "README"},
	}, nil)
	require.Error(t, err)
	require.NotNil(t, resp)
	require.Equal(t, http.StatusInternalServerError, resp.StatusCode)
	require.Equal(t, int32(1), hits.Load())
}

func TestCreateScorecardDoesNotRetryAlreadyExistsMappedAs500(t *testing.T) {
	var hits atomic.Int32
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		hits.Add(1)
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusInternalServerError)
		_, _ = w.Write([]byte(
			`{"message":"A scorecard with identifier 'tf_sanity_gold_standard' already exists. Please use a different identifier."}`,
		))
	}))
	defer server.Close()

	cfg := NewConfiguration()
	cfg.BasePath = server.URL
	client := NewAPIClient(cfg)

	_, resp, err := client.ScorecardsApi.CreateScorecard(context.Background(), ScorecardRequest{
		Scorecard: Scorecard{Identifier: "tf_sanity_gold_standard", Name: "Gold"},
	}, nil)
	require.EqualError(t, err,
		"500 Internal Server Error: A scorecard with identifier 'tf_sanity_gold_standard' already exists. Please use a different identifier.")
	require.NotNil(t, resp)
	require.Equal(t, http.StatusInternalServerError, resp.StatusCode)
	require.Equal(t, int32(1), hits.Load())
}

func TestUpdateScorecardDoesNotRetryMissingCheckMappedAs500(t *testing.T) {
	var hits atomic.Int32
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		hits.Add(1)
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusInternalServerError)
		_, _ = w.Write([]byte(`{"message":"Error while saving scorecard. Could not find check tf_sanity_readme_exists"}`))
	}))
	defer server.Close()

	cfg := NewConfiguration()
	cfg.BasePath = server.URL
	client := NewAPIClient(cfg)

	_, resp, err := client.ScorecardsApi.UpdateScorecard(context.Background(), ScorecardRequest{
		Scorecard: Scorecard{Identifier: "gold", Name: "Gold"},
		Checks:    []ScorecardCheck{{Identifier: "tf_sanity_readme_exists", Custom: true}},
	}, "gold", nil)
	require.EqualError(t, err, "500 Internal Server Error: Error while saving scorecard. Could not find check tf_sanity_readme_exists")
	require.NotNil(t, resp)
	require.Equal(t, http.StatusInternalServerError, resp.StatusCode)
	require.Equal(t, int32(1), hits.Load())
}

func TestIsNonRetryableIDPErrorPreservesBody(t *testing.T) {
	resp := &http.Response{
		StatusCode: http.StatusInternalServerError,
		Body:       io.NopCloser(bytes.NewReader([]byte(`{"message":"Scorecard details not found for scorecardId [gold]"}`))),
	}
	require.True(t, isNonRetryableIDPError(resp))

	body, err := io.ReadAll(resp.Body)
	require.NoError(t, err)
	require.Contains(t, string(body), "Scorecard details not found")
}

func TestDeleteCheckDoesNotRetryReferencedMappedAs500(t *testing.T) {
	var hits atomic.Int32
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		hits.Add(1)
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusInternalServerError)
		_, _ = w.Write([]byte(`{"message":"Could not delete the check [readme] as it is referenced by other scorecards"}`))
	}))
	defer server.Close()

	cfg := NewConfiguration()
	cfg.BasePath = server.URL
	client := NewAPIClient(cfg)

	resp, err := client.ChecksApi.DeleteCheck(context.Background(), "readme", nil)
	require.Error(t, err)
	require.NotNil(t, resp)
	require.Equal(t, http.StatusInternalServerError, resp.StatusCode)
	require.Equal(t, int32(1), hits.Load())
}
