package split_test

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/harness/harness-go-sdk/harness/split"
	"github.com/stretchr/testify/require"
)

func TestNewAPIClient(t *testing.T) {
	cfg := &split.Configuration{
		BasePath:      "https://api.split.io",
		DefaultHeader: make(map[string]string),
		UserAgent:     "harness-go-sdk-test",
	}
	client := split.NewAPIClient(cfg)
	require.NotNil(t, client)
	require.NotNil(t, client.HTTPClient)
	require.Equal(t, cfg.BasePath, client.BasePath)
	require.Equal(t, "https://api.split.io", client.BasePath)
}

func TestNewAPIClient_WithDefaultConfig(t *testing.T) {
	cfg := split.NewDefaultConfiguration()
	client := split.NewAPIClient(cfg)
	require.NotNil(t, client)
	require.NotNil(t, client.HTTPClient)
	require.Equal(t, split.DefaultSplitBasePath, client.BasePath)
}

func TestClient_429Retry(t *testing.T) {
	var requestCount int
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		requestCount++
		if requestCount == 1 {
			w.Header().Set("X-RateLimit-Reset-Seconds-Org", "0")
			w.Header().Set("X-RateLimit-Reset-Seconds-IP", "0")
			w.WriteHeader(http.StatusTooManyRequests)
			_, _ = w.Write([]byte(`{"error":"rate limited"}`))
			return
		}
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		_, _ = w.Write([]byte(`{}`))
	}))
	defer server.Close()

	cfg := split.NewDefaultConfiguration()
	cfg.BasePath = server.URL
	client := split.NewAPIClient(cfg)

	req, err := http.NewRequest(http.MethodGet, server.URL+"/test", nil)
	require.NoError(t, err)

	resp, err := client.Do(req)
	require.NoError(t, err)
	defer func() { _ = resp.Body.Close() }()

	require.Equal(t, http.StatusOK, resp.StatusCode, "client should succeed after retry")
	require.Equal(t, 2, requestCount, "server should receive 2 requests (first 429, then retry)")
}
