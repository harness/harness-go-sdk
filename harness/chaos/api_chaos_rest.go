package chaos

import (
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"strings"
)

// scopeQuery builds the account/org/project query string used by chaos manager v3 APIs.
func scopeQuery(scope Scope) url.Values {
	q := url.Values{}
	if scope.AccountID != "" {
		q.Set("accountIdentifier", scope.AccountID)
	}
	if scope.OrgID != "" {
		q.Set("organizationIdentifier", scope.OrgID)
	}
	if scope.ProjectID != "" {
		q.Set("projectIdentifier", scope.ProjectID)
	}
	if scope.ParentUniqueID != "" {
		q.Set("parentUniqueId", scope.ParentUniqueID)
	}
	if scope.CorrelationID != "" {
		q.Set("correlationID", scope.CorrelationID)
	}
	return q
}

// doChaosJSON performs a JSON request against the chaos manager base path.
// Generated swagger clients keep using their own methods; this helper exists
// only for hand-written clients so those generated files stay untouched.
func (c *APIClient) doChaosJSON(ctx context.Context, method, path string, query url.Values, body interface{}, out interface{}) (*http.Response, error) {
	if c == nil || c.cfg == nil {
		return nil, fmt.Errorf("chaos API client is not configured")
	}
	return c.doAbsoluteJSON(ctx, method, c.cfg.BasePath+path, query, body, out)
}

// doAbsoluteJSON is doChaosJSON for a full URL. Service discovery is a sibling
// of the chaos manager base path, so infrastructure scans call it directly.
func (c *APIClient) doAbsoluteJSON(ctx context.Context, method, rawURL string, query url.Values, body interface{}, out interface{}) (*http.Response, error) {
	if c == nil || c.cfg == nil {
		return nil, fmt.Errorf("chaos API client is not configured")
	}

	headers := map[string]string{
		"Accept": "application/json",
	}
	if body != nil {
		headers["Content-Type"] = "application/json"
	}
	if ctx != nil {
		if auth, ok := ctx.Value(ContextAPIKey).(APIKey); ok {
			key := auth.Key
			if auth.Prefix != "" {
				key = auth.Prefix + " " + auth.Key
			}
			if key != "" {
				headers["X-API-Key"] = key
			}
		}
	}

	req, err := c.prepareRequest(ctx, rawURL, strings.ToUpper(method), body, headers, query, url.Values{}, "", nil)
	if err != nil {
		return nil, err
	}

	resp, err := c.callAPI(req)
	if err != nil || resp == nil {
		return resp, err
	}
	defer resp.Body.Close()

	respBody, err := io.ReadAll(resp.Body)
	if err != nil {
		return resp, err
	}

	if resp.StatusCode >= 300 {
		return resp, swaggerError(resp, respBody)
	}

	if out == nil || len(respBody) == 0 {
		return resp, nil
	}
	if err := decodeChaosJSON(out, respBody, resp.Header.Get("Content-Type")); err != nil {
		return resp, err
	}
	return resp, nil
}

// decodeChaosJSON accepts application/json, an empty content type, and text/plain.
// The generated decode helper rejects anything that is not exactly a json or xml
// content type, which drops valid chaos-manager bodies behind some gateways.
func decodeChaosJSON(out interface{}, body []byte, contentType string) error {
	ct := strings.ToLower(contentType)
	if strings.Contains(ct, "xml") || strings.Contains(ct, "html") {
		return fmt.Errorf("unexpected content type %q", contentType)
	}
	if err := json.Unmarshal(body, out); err != nil {
		return fmt.Errorf("decode response: %w", err)
	}
	return nil
}

func swaggerError(resp *http.Response, body []byte) GenericSwaggerError {
	msg := resp.Status
	var apiErr apiErrorBody
	if err := json.Unmarshal(body, &apiErr); err == nil {
		if apiErr.Message != "" {
			msg = apiErr.Message
		} else if apiErr.Description != "" {
			msg = apiErr.Description
		}
	}
	return GenericSwaggerError{
		body:  body,
		error: msg,
		model: apiErr,
	}
}

type apiErrorBody struct {
	Code          interface{} `json:"code,omitempty"`
	Description   string      `json:"description,omitempty"`
	Message       string      `json:"message,omitempty"`
	CorrelationID string      `json:"correlationID,omitempty"`
}
