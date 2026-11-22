package agent_test

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"net/http/httptest"
	"strconv"
	"testing"

	"github.com/harness/harness-go-sdk/harness/nextgen"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestListAllPipelines_MultiplePages(t *testing.T) {
	pages := [][]nextgen.PmsPipelineSummaryResponse{
		{{Identifier: "pipe-1", Name: "Pipeline 1"}, {Identifier: "pipe-2", Name: "Pipeline 2"}},
		{{Identifier: "pipe-3", Name: "Pipeline 3"}},
	}

	srv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		assert.Contains(t, r.URL.Path, "/pipeline/api/pipelines/list")
		pageStr := r.URL.Query().Get("page")
		page, _ := strconv.Atoi(pageStr)

		var content []nextgen.PmsPipelineSummaryResponse
		isLast := true
		if page < len(pages) {
			content = pages[page]
			isLast = page == len(pages)-1
		}

		resp := nextgen.ResponseDtoPagePmsPipelineSummaryResponse{
			Status: "SUCCESS",
			Data: &nextgen.PagePmsPipelineSummaryResponse{
				Content: content,
				Last:    isLast,
				Number:  int32(page),
				Size:    25,
				Empty:   len(content) == 0,
			},
		}
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(resp)
	}))
	defer srv.Close()

	cfg := nextgen.NewConfiguration()
	cfg.BasePath = srv.URL
	cfg.ApiKey = "test"
	client := nextgen.NewAPIClient(cfg)

	ctx := createAPIKeyContext("test")
	iter := nextgen.ListAllPipelines(ctx, client, "acc1", "org1", "proj1")

	var names []string
	for iter.Next() {
		names = append(names, iter.Value().Name)
	}
	require.NoError(t, iter.Err())
	assert.Equal(t, []string{"Pipeline 1", "Pipeline 2", "Pipeline 3"}, names)
}

func TestListAllPipelines_EmptyProject(t *testing.T) {
	srv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		resp := nextgen.ResponseDtoPagePmsPipelineSummaryResponse{
			Status: "SUCCESS",
			Data: &nextgen.PagePmsPipelineSummaryResponse{
				Content: nil,
				Last:    true,
				Empty:   true,
			},
		}
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(resp)
	}))
	defer srv.Close()

	cfg := nextgen.NewConfiguration()
	cfg.BasePath = srv.URL
	cfg.ApiKey = "test"
	client := nextgen.NewAPIClient(cfg)

	ctx := createAPIKeyContext("test")
	iter := nextgen.ListAllPipelines(ctx, client, "acc1", "org1", "proj1")

	results, err := iter.Collect()
	require.NoError(t, err)
	assert.Empty(t, results)
}

func TestListAllPipelines_APIError(t *testing.T) {
	srv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusForbidden)
		fmt.Fprintf(w, `{"code":"UNAUTHORIZED","message":"not allowed"}`)
	}))
	defer srv.Close()

	cfg := nextgen.NewConfiguration()
	cfg.BasePath = srv.URL
	cfg.ApiKey = "bad-key"
	client := nextgen.NewAPIClient(cfg)

	ctx := createAPIKeyContext("bad-key")
	iter := nextgen.ListAllPipelines(ctx, client, "acc1", "org1", "proj1")

	results, err := iter.Collect()
	assert.Error(t, err)
	assert.Empty(t, results)
}

func createAPIKeyContext(key string) context.Context {
	return context.WithValue(context.Background(), nextgen.ContextAPIKey, nextgen.APIKey{Key: key})
}
