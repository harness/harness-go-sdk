/*
 * Split API client - HTTP transport with rate-limit and retry.
 */
package split

import (
	"bytes"
	"io"
	"net/http"
	"net/url"
	"strconv"
	"strings"
	"time"
)

const (
	headerRateLimitResetOrg = "X-RateLimit-Reset-Seconds-Org"
	headerRateLimitResetIP  = "X-RateLimit-Reset-Seconds-IP"
	defaultResetSeconds     = 1
	max429Retries           = 3
	max5xxRetries           = 3
)

// rateLimitRetryTransport wraps an http.RoundTripper to add default headers, auth,
// and retry on 429 (rate limit) and optionally on 5xx.
type rateLimitRetryTransport struct {
	cfg     *Configuration
	wrapped http.RoundTripper
}

// newRateLimitRetryTransport returns a RoundTripper that applies cfg headers and auth,
// and retries on 429 (after waiting per Split headers) and on 5xx with backoff.
func newRateLimitRetryTransport(cfg *Configuration, wrapped http.RoundTripper) http.RoundTripper {
	if wrapped == nil {
		wrapped = http.DefaultTransport
	}
	return &rateLimitRetryTransport{cfg: cfg, wrapped: wrapped}
}

// RoundTrip implements http.RoundTripper.
func (t *rateLimitRetryTransport) RoundTrip(req *http.Request) (*http.Response, error) {
	// Resolve relative URL with BasePath
	if req.URL.Scheme == "" || req.URL.Host == "" {
		base := strings.TrimSuffix(t.cfg.BasePath, "/")
		path := req.URL.Path
		if path != "" && path[0] != '/' {
			path = "/" + path
		}
		full := base + path
		if req.URL.RawQuery != "" {
			full += "?" + req.URL.RawQuery
		}
		req.URL, _ = url.Parse(full)
	}

	// Buffer body for potential retries (body can only be read once)
	var bodyBytes []byte
	if req.Body != nil {
		var err error
		bodyBytes, err = io.ReadAll(req.Body)
		if err != nil {
			return nil, err
		}
		_ = req.Body.Close()
		req.Body = io.NopCloser(bytes.NewReader(bodyBytes))
	}

	// Apply headers and auth before first attempt
	t.applyHeaders(req)

	resp, err := t.doRoundTrip(req, bodyBytes)
	if err != nil {
		return nil, err
	}

	// 429: retry after reset seconds
	for i := 0; resp.StatusCode == http.StatusTooManyRequests && i < max429Retries; i++ {
		resetSec := t.parseResetSeconds(resp)
		_ = resp.Body.Close()
		time.Sleep(time.Duration(resetSec) * time.Second)
		req = t.cloneRequest(req, bodyBytes)
		t.applyHeaders(req)
		resp, err = t.doRoundTrip(req, bodyBytes)
		if err != nil {
			return nil, err
		}
	}

	// 5xx: optional retry with backoff
	for i := 0; resp.StatusCode >= 500 && resp.StatusCode < 600 && i < max5xxRetries; i++ {
		backoff := time.Duration(1<<uint(i)) * time.Second
		_ = resp.Body.Close()
		time.Sleep(backoff)
		req = t.cloneRequest(req, bodyBytes)
		t.applyHeaders(req)
		resp, err = t.doRoundTrip(req, bodyBytes)
		if err != nil {
			return nil, err
		}
	}

	return resp, nil
}

func (t *rateLimitRetryTransport) applyHeaders(req *http.Request) {
	if t.cfg.UserAgent != "" {
		req.Header.Set("User-Agent", t.cfg.UserAgent)
	}
	// Only set defaults if not already set by caller (e.g. CSV uploads use Content-Type: text/csv).
	if req.Header.Get("Content-Type") == "" {
		req.Header.Set("Content-Type", "application/json")
	}
	if req.Header.Get("Accept") == "" {
		req.Header.Set("Accept", "application/json")
	}
	if t.cfg.ApiKey != "" {
		req.Header.Set("x-api-key", t.cfg.ApiKey)
	}
	for k, v := range t.cfg.DefaultHeader {
		req.Header.Set(k, v)
	}
}

func (t *rateLimitRetryTransport) doRoundTrip(req *http.Request, bodyBytes []byte) (*http.Response, error) {
	// Ensure body is set for this attempt (may be retry)
	if len(bodyBytes) > 0 {
		req.Body = io.NopCloser(bytes.NewReader(bodyBytes))
	}
	return t.wrapped.RoundTrip(req)
}

func (t *rateLimitRetryTransport) cloneRequest(req *http.Request, bodyBytes []byte) *http.Request {
	req2 := req.Clone(req.Context())
	if len(bodyBytes) > 0 {
		req2.Body = io.NopCloser(bytes.NewReader(bodyBytes))
	}
	return req2
}

// parseResetSeconds returns the number of seconds to wait from 429 response headers.
// Uses max of X-RateLimit-Reset-Seconds-Org and X-RateLimit-Reset-Seconds-IP; default 1 if missing.
func (t *rateLimitRetryTransport) parseResetSeconds(resp *http.Response) int {
	org := parseResetHeader(resp.Header.Get(headerRateLimitResetOrg))
	ip := parseResetHeader(resp.Header.Get(headerRateLimitResetIP))
	if org > ip {
		return org
	}
	if ip > 0 {
		return ip
	}
	if org > 0 {
		return org
	}
	return defaultResetSeconds
}

func parseResetHeader(v string) int {
	v = strings.TrimSpace(v)
	if v == "" {
		return 0
	}
	n, _ := strconv.Atoi(v)
	if n < 0 {
		return 0
	}
	return n
}
