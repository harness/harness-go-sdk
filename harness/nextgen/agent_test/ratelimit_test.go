package agent_test

import (
	"context"
	"fmt"
	"net/http"
	"net/http/httptest"
	"sync/atomic"
	"testing"
	"time"

	"github.com/harness/harness-go-sdk/harness/nextgen"
	"github.com/hashicorp/go-retryablehttp"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestRateLimitCheckRetry_429(t *testing.T) {
	resp := &http.Response{StatusCode: http.StatusTooManyRequests}
	shouldRetry, err := nextgen.RateLimitCheckRetry(context.Background(), resp, nil)
	assert.True(t, shouldRetry)
	assert.NoError(t, err)
}

func TestRateLimitCheckRetry_200(t *testing.T) {
	resp := &http.Response{StatusCode: http.StatusOK}
	shouldRetry, err := nextgen.RateLimitCheckRetry(context.Background(), resp, nil)
	assert.False(t, shouldRetry)
	assert.NoError(t, err)
}

func TestRateLimitCheckRetry_500(t *testing.T) {
	resp := &http.Response{StatusCode: http.StatusInternalServerError}
	shouldRetry, err := nextgen.RateLimitCheckRetry(context.Background(), resp, nil)
	assert.True(t, shouldRetry)
	assert.NoError(t, err)
}

func TestRateLimitBackoff_429WithRetryAfter(t *testing.T) {
	resp := &http.Response{
		StatusCode: http.StatusTooManyRequests,
		Header:     http.Header{},
	}
	resp.Header.Set("Retry-After", "10")

	backoff := nextgen.RateLimitBackoff(1*time.Second, 30*time.Second, 1, resp)
	assert.Equal(t, 10*time.Second, backoff)
}

func TestRateLimitBackoff_429WithReset(t *testing.T) {
	future := time.Now().Add(15 * time.Second).Unix()
	resp := &http.Response{
		StatusCode: http.StatusTooManyRequests,
		Header:     http.Header{},
	}
	resp.Header.Set("X-RateLimit-Reset", fmt.Sprintf("%d", future))

	backoff := nextgen.RateLimitBackoff(1*time.Second, 30*time.Second, 1, resp)
	assert.InDelta(t, 15, backoff.Seconds(), 2)
}

func TestRateLimitBackoff_429NoHeaders(t *testing.T) {
	resp := &http.Response{
		StatusCode: http.StatusTooManyRequests,
		Header:     http.Header{},
	}

	backoff := nextgen.RateLimitBackoff(1*time.Second, 30*time.Second, 1, resp)
	assert.Equal(t, 5*time.Second, backoff)
}

func TestRateLimitBackoff_NonRateLimitUsesDefault(t *testing.T) {
	resp := &http.Response{
		StatusCode: http.StatusInternalServerError,
		Header:     http.Header{},
	}

	backoff := nextgen.RateLimitBackoff(1*time.Second, 30*time.Second, 0, resp)
	assert.Equal(t, 1*time.Second, backoff)
}

func TestRateLimitRetry_Integration(t *testing.T) {
	var attempts int32
	srv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		n := atomic.AddInt32(&attempts, 1)
		if n <= 2 {
			w.Header().Set("Retry-After", "1")
			w.WriteHeader(http.StatusTooManyRequests)
			w.Write([]byte(`{"message":"rate limited"}`))
			return
		}
		w.WriteHeader(http.StatusOK)
		w.Write([]byte(`{"status":"SUCCESS"}`))
	}))
	defer srv.Close()

	client := retryablehttp.NewClient()
	client.RetryMax = 5
	client.CheckRetry = nextgen.RateLimitCheckRetry
	client.Backoff = nextgen.RateLimitBackoff
	client.RetryWaitMin = 100 * time.Millisecond
	client.RetryWaitMax = 2 * time.Second
	client.Logger = nil

	req, err := retryablehttp.NewRequest("GET", srv.URL+"/test", nil)
	require.NoError(t, err)

	resp, err := client.Do(req)
	require.NoError(t, err)
	defer resp.Body.Close()

	assert.Equal(t, http.StatusOK, resp.StatusCode)
	assert.Equal(t, int32(3), atomic.LoadInt32(&attempts))
}
