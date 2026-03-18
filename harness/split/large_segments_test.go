package split

import (
	"encoding/json"
	"io"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestBuildCSVFromKeys(t *testing.T) {
	out := BuildCSVFromKeys([]string{"id1", "id2", "id3"})
	require.Equal(t, "id1\nid2\nid3", string(out))

	// Keys with comma are quoted
	out2 := BuildCSVFromKeys([]string{"a,b", "c"})
	require.Contains(t, string(out2), "a,b")
	require.Contains(t, string(out2), "c")
}

func TestLargeSegmentsService_OpenChangeRequestForLargeSegmentUpload(t *testing.T) {
	var reqBody LargeSegmentChangeRequestRequest
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		require.Equal(t, http.MethodPost, r.Method)
		require.Equal(t, "/internal/api/v2/changeRequests/ws/ws-1/environments/env-1", r.URL.Path)
		_ = json.NewDecoder(r.Body).Decode(&reqBody)
		resp := LargeSegmentChangeRequestResponse{
			ID:   "cr-123",
			Status: "AWAITING",
			TransactionMetadata: &ChangeRequestTxMetadata{
				Method: "PUT",
				URL:    "https://upload.example.com/presigned",
				Headers: map[string][]string{"Host": {"upload.example.com"}},
			},
		}
		w.WriteHeader(http.StatusOK)
		_ = json.NewEncoder(w).Encode(resp)
	}))
	defer server.Close()

	cfg := NewDefaultConfiguration()
	cfg.BasePath = server.URL
	client := NewAPIClient(cfg)

	creds, err := client.LargeSegments.OpenChangeRequestForLargeSegmentUpload("ws-1", "env-1", "my_large_segment", "title", "comment", nil)
	require.NoError(t, err)
	require.Equal(t, "cr-123", creds.ChangeRequestID)
	require.Equal(t, "https://upload.example.com/presigned", creds.UploadURL)
	require.Equal(t, "PUT", creds.Method)
	require.Equal(t, "upload.example.com", creds.Host)
	require.Equal(t, "my_large_segment", reqBody.LargeSegment.Name)
	require.Equal(t, "UPLOAD", reqBody.OperationType)
}

func TestLargeSegmentsService_UploadKeysToLargeSegment(t *testing.T) {
	var putBody []byte
	uploadServer := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		require.Equal(t, http.MethodPut, r.Method)
		require.Equal(t, "text/csv", r.Header.Get("Content-Type"))
		putBody, _ = io.ReadAll(r.Body)
		w.WriteHeader(http.StatusOK)
	}))
	defer uploadServer.Close()

	cfg := NewDefaultConfiguration()
	cfg.BasePath = "http://api.example.com"
	client := NewAPIClient(cfg)

	creds := &LargeSegmentUploadCredentials{
		UploadURL: uploadServer.URL,
		Method:    http.MethodPut,
	}
	err := client.LargeSegments.UploadKeysToLargeSegment(creds, []byte("k1\nk2\n"))
	require.NoError(t, err)
	require.Equal(t, []byte("k1\nk2\n"), putBody)
}

func TestLargeSegmentsService_List(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		require.Equal(t, http.MethodGet, r.Method)
		require.Equal(t, "/internal/api/v2/large-segments/ws/ws-1", r.URL.Path)
		w.WriteHeader(http.StatusOK)
		_ = json.NewEncoder(w).Encode([]LargeSegmentMetadata{
			{ID: "ls-1", Name: "large_1", TrafficType: &TrafficType{ID: "tt-1", Name: "user"}},
		})
	}))
	defer server.Close()

	cfg := NewDefaultConfiguration()
	cfg.BasePath = server.URL
	client := NewAPIClient(cfg)

	list, err := client.LargeSegments.List("ws-1")
	require.NoError(t, err)
	require.Len(t, list, 1)
	require.Equal(t, "large_1", list[0].Name)
}

func TestLargeSegmentsService_Get(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		require.Equal(t, http.MethodGet, r.Method)
		require.Contains(t, r.URL.Path, "/large-segments/ws/ws-1/")
		w.WriteHeader(http.StatusOK)
		_ = json.NewEncoder(w).Encode(LargeSegmentMetadata{
			ID: "ls-1", Name: "my_large", Description: "Example", TrafficType: &TrafficType{ID: "tt-1", Name: "user"},
		})
	}))
	defer server.Close()

	cfg := NewDefaultConfiguration()
	cfg.BasePath = server.URL
	client := NewAPIClient(cfg)

	meta, err := client.LargeSegments.Get("ws-1", "my_large")
	require.NoError(t, err)
	require.Equal(t, "my_large", meta.Name)
	require.Equal(t, "Example", meta.Description)
}

func TestLargeSegmentsService_Create(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		require.Equal(t, http.MethodPost, r.Method)
		require.Equal(t, "/internal/api/v2/large-segments/ws/ws-1/trafficTypes/tt-1", r.URL.Path)
		var req LargeSegmentCreateRequest
		_ = json.NewDecoder(r.Body).Decode(&req)
		require.Equal(t, "new_large", req.Name)
		w.WriteHeader(http.StatusCreated)
		_ = json.NewEncoder(w).Encode(LargeSegmentMetadata{ID: "ls-new", Name: "new_large", TrafficType: &TrafficType{ID: "tt-1", Name: "user"}})
	}))
	defer server.Close()

	cfg := NewDefaultConfiguration()
	cfg.BasePath = server.URL
	client := NewAPIClient(cfg)

	meta, err := client.LargeSegments.Create("ws-1", "tt-1", LargeSegmentCreateRequest{Name: "new_large", Description: "Created by test"})
	require.NoError(t, err)
	require.Equal(t, "new_large", meta.Name)
}
