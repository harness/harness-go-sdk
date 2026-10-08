package nextgen

import (
	"fmt"
	"net/http"
	"strconv"
	"time"
)

// RateLimitError is returned when the server responds with 429.
// Agents can inspect RetryAfter to decide when to retry.
type RateLimitError struct {
	Limit      int
	Remaining  int
	ResetEpoch int64
	RetryAfter time.Duration
	Message    string
}

func (e *RateLimitError) Error() string {
	return fmt.Sprintf("rate limit exceeded: limit=%d, remaining=%d, retry_after=%s", e.Limit, e.Remaining, e.RetryAfter)
}

// NotFoundError is returned when the server responds with 404.
type NotFoundError struct {
	Resource string
	ID       string
	Message  string
}

func (e *NotFoundError) Error() string {
	if e.Resource != "" && e.ID != "" {
		return fmt.Sprintf("%s %q not found", e.Resource, e.ID)
	}
	if e.Message != "" {
		return e.Message
	}
	return "resource not found"
}

// RequestValidationError is returned when the server responds with 400.
type RequestValidationError struct {
	Field   string
	Message string
	Code    string
}

func (e *RequestValidationError) Error() string {
	if e.Field != "" {
		return fmt.Sprintf("validation error on field %q: %s", e.Field, e.Message)
	}
	return fmt.Sprintf("validation error: %s", e.Message)
}

// UnauthorizedError is returned when the server responds with 401 or 403.
type UnauthorizedError struct {
	StatusCode int
	Message    string
}

func (e *UnauthorizedError) Error() string {
	return fmt.Sprintf("unauthorized (HTTP %d): %s", e.StatusCode, e.Message)
}

// ServerError is returned when the server responds with 5xx after all retries.
type ServerError struct {
	StatusCode    int
	Message       string
	CorrelationID string
}

func (e *ServerError) Error() string {
	msg := fmt.Sprintf("server error (HTTP %d): %s", e.StatusCode, e.Message)
	if e.CorrelationID != "" {
		msg += fmt.Sprintf(" [correlation_id=%s]", e.CorrelationID)
	}
	return msg
}

// RateLimitInfo carries rate limit metadata from response headers.
type RateLimitInfo struct {
	Limit      int
	Remaining  int
	ResetEpoch int64
}

// ParseRateLimitInfo extracts rate limit metadata from response headers.
func ParseRateLimitInfo(resp *http.Response) *RateLimitInfo {
	if resp == nil {
		return nil
	}
	limit := resp.Header.Get("X-RateLimit-Limit")
	remaining := resp.Header.Get("X-RateLimit-Remaining")
	reset := resp.Header.Get("X-RateLimit-Reset")
	if limit == "" && remaining == "" && reset == "" {
		return nil
	}
	info := &RateLimitInfo{}
	if v, err := strconv.Atoi(limit); err == nil {
		info.Limit = v
	}
	if v, err := strconv.Atoi(remaining); err == nil {
		info.Remaining = v
	}
	if v, err := strconv.ParseInt(reset, 10, 64); err == nil {
		info.ResetEpoch = v
	}
	return info
}

// parseRateLimitError constructs a RateLimitError from a 429 response.
func parseRateLimitError(resp *http.Response) *RateLimitError {
	e := &RateLimitError{
		Message: "rate limit exceeded",
	}
	if v, err := strconv.Atoi(resp.Header.Get("X-RateLimit-Limit")); err == nil {
		e.Limit = v
	}
	if v, err := strconv.Atoi(resp.Header.Get("X-RateLimit-Remaining")); err == nil {
		e.Remaining = v
	}
	if v, err := strconv.ParseInt(resp.Header.Get("X-RateLimit-Reset"), 10, 64); err == nil {
		e.ResetEpoch = v
	}
	if retryAfter := resp.Header.Get("Retry-After"); retryAfter != "" {
		if sec, err := strconv.Atoi(retryAfter); err == nil {
			e.RetryAfter = time.Duration(sec) * time.Second
		} else if t, err := http.ParseTime(retryAfter); err == nil {
			e.RetryAfter = time.Until(t)
			if e.RetryAfter < 0 {
				e.RetryAfter = 0
			}
		}
	}
	if e.RetryAfter == 0 {
		e.RetryAfter = 60 * time.Second
	}
	return e
}
