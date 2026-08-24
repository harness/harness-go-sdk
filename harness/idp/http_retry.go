package idp

import (
	"bytes"
	"context"
	"io"
	"net/http"
	"strings"

	"github.com/hashicorp/go-retryablehttp"
)

func applyIDPRetryPolicy(client *retryablehttp.Client) {
	if client == nil {
		return
	}
	next := client.CheckRetry
	if next == nil {
		next = retryablehttp.DefaultRetryPolicy
	}
	client.CheckRetry = func(ctx context.Context, resp *http.Response, err error) (bool, error) {
		if isNonRetryableIDPError(resp) {
			return false, nil
		}
		return next(ctx, resp, err)
	}
}

func isNonRetryableIDPError(resp *http.Response) bool {
	if resp == nil || resp.Body == nil || resp.StatusCode < 500 {
		return false
	}

	body, err := io.ReadAll(resp.Body)
	_ = resp.Body.Close()
	resp.Body = io.NopCloser(bytes.NewReader(body))
	if err != nil {
		return false
	}

	msg := strings.ToLower(string(body))
	return strings.Contains(msg, "not found") ||
		strings.Contains(msg, "could not find check") ||
		strings.Contains(msg, "invalidrequest") ||
		strings.Contains(msg, "already created") ||
		strings.Contains(msg, "already exists") ||
		strings.Contains(msg, "referenced by")
}
