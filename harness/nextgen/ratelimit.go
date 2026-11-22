package nextgen

import (
	"context"
	"net/http"
	"strconv"
	"time"

	"github.com/hashicorp/go-retryablehttp"
)

const (
	headerRateLimitLimit     = "X-RateLimit-Limit"
	headerRateLimitRemaining = "X-RateLimit-Remaining"
	headerRateLimitReset     = "X-RateLimit-Reset"
	headerRetryAfter         = "Retry-After"

	defaultRetryAfterSeconds = 5
	minRetryAfterSeconds     = 2
	maxRateLimitRetries      = 5
)

// RateLimitCheckRetry is a retryablehttp.CheckRetry that handles 429 responses
// by respecting Retry-After and X-RateLimit-Reset headers.
func RateLimitCheckRetry(ctx context.Context, resp *http.Response, err error) (bool, error) {
	if err != nil {
		return retryablehttp.DefaultRetryPolicy(ctx, resp, err)
	}
	if resp.StatusCode == http.StatusTooManyRequests {
		return true, nil
	}
	return retryablehttp.DefaultRetryPolicy(ctx, resp, err)
}

// RateLimitBackoff returns a retryablehttp.Backoff that respects rate limit headers.
func RateLimitBackoff(min, max time.Duration, attemptNum int, resp *http.Response) time.Duration {
	if resp != nil && resp.StatusCode == http.StatusTooManyRequests {
		wait := parseRetryWait(resp)
		if wait > 0 {
			return wait
		}
	}
	return retryablehttp.DefaultBackoff(min, max, attemptNum, resp)
}

func parseRetryWait(resp *http.Response) time.Duration {
	if s := resp.Header.Get(headerRetryAfter); s != "" {
		if sec, err := strconv.Atoi(s); err == nil && sec > 0 {
			if sec < minRetryAfterSeconds {
				sec = minRetryAfterSeconds
			}
			return time.Duration(sec) * time.Second
		}
		if t, err := http.ParseTime(s); err == nil {
			wait := time.Until(t)
			if wait < time.Duration(minRetryAfterSeconds)*time.Second {
				return time.Duration(minRetryAfterSeconds) * time.Second
			}
			return wait
		}
	}
	if s := resp.Header.Get(headerRateLimitReset); s != "" {
		if epoch, err := strconv.ParseInt(s, 10, 64); err == nil {
			wait := time.Until(time.Unix(epoch, 0))
			if wait > 0 {
				return wait
			}
		}
	}
	return time.Duration(defaultRetryAfterSeconds) * time.Second
}
