package agent_test

import (
	"net/http"
	"testing"
	"time"

	"github.com/harness/harness-go-sdk/harness/nextgen"
	"github.com/stretchr/testify/assert"
)

func TestRateLimitError_Error(t *testing.T) {
	e := &nextgen.RateLimitError{Limit: 100, Remaining: 0, RetryAfter: 30 * time.Second}
	assert.Contains(t, e.Error(), "rate limit exceeded")
	assert.Contains(t, e.Error(), "limit=100")
	assert.Contains(t, e.Error(), "retry_after=30s")
}

func TestNotFoundError_Error(t *testing.T) {
	e := &nextgen.NotFoundError{Resource: "Pipeline", ID: "my-pipeline"}
	assert.Equal(t, `Pipeline "my-pipeline" not found`, e.Error())

	e2 := &nextgen.NotFoundError{Message: "custom message"}
	assert.Equal(t, "custom message", e2.Error())

	e3 := &nextgen.NotFoundError{}
	assert.Equal(t, "resource not found", e3.Error())
}

func TestRequestValidationError_Error(t *testing.T) {
	e := &nextgen.RequestValidationError{Field: "name", Message: "required"}
	assert.Contains(t, e.Error(), "name")
	assert.Contains(t, e.Error(), "required")

	e2 := &nextgen.RequestValidationError{Message: "invalid request"}
	assert.Contains(t, e2.Error(), "invalid request")
}

func TestUnauthorizedError_Error(t *testing.T) {
	e := &nextgen.UnauthorizedError{StatusCode: 401, Message: "invalid token"}
	assert.Contains(t, e.Error(), "401")
	assert.Contains(t, e.Error(), "invalid token")
}

func TestServerError_Error(t *testing.T) {
	e := &nextgen.ServerError{StatusCode: 503, Message: "service unavailable", CorrelationID: "abc-123"}
	assert.Contains(t, e.Error(), "503")
	assert.Contains(t, e.Error(), "abc-123")
}

func TestParseRateLimitInfo(t *testing.T) {
	tests := []struct {
		name     string
		headers  map[string]string
		expected *nextgen.RateLimitInfo
	}{
		{
			name:     "no headers",
			headers:  map[string]string{},
			expected: nil,
		},
		{
			name: "all headers present",
			headers: map[string]string{
				"X-RateLimit-Limit":     "100",
				"X-RateLimit-Remaining": "42",
				"X-RateLimit-Reset":     "1700000000",
			},
			expected: &nextgen.RateLimitInfo{Limit: 100, Remaining: 42, ResetEpoch: 1700000000},
		},
		{
			name: "partial headers",
			headers: map[string]string{
				"X-RateLimit-Limit": "200",
			},
			expected: &nextgen.RateLimitInfo{Limit: 200, Remaining: 0, ResetEpoch: 0},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			resp := &http.Response{Header: http.Header{}}
			for k, v := range tt.headers {
				resp.Header.Set(k, v)
			}
			info := nextgen.ParseRateLimitInfo(resp)
			if tt.expected == nil {
				assert.Nil(t, info)
			} else {
				assert.Equal(t, tt.expected.Limit, info.Limit)
				assert.Equal(t, tt.expected.Remaining, info.Remaining)
				assert.Equal(t, tt.expected.ResetEpoch, info.ResetEpoch)
			}
		})
	}
}

func TestParseRateLimitInfo_NilResponse(t *testing.T) {
	assert.Nil(t, nextgen.ParseRateLimitInfo(nil))
}
