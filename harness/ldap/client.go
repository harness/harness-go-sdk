package ldap

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

// APIClient manages communication with the v1 LDAP settings API.
type APIClient struct {
	cfg *Configuration

	LdapSettingsApi *LdapSettingsApiService
}

func NewAPIClient(cfg *Configuration) *APIClient {
	if cfg.HTTPClient == nil {
		cfg.HTTPClient = NewConfiguration().HTTPClient
	}
	c := &APIClient{cfg: cfg}
	c.LdapSettingsApi = &LdapSettingsApiService{client: c}
	return c
}

// GenericError is the error returned for non-2xx responses. It carries the raw
// response body so callers can inspect the backend's ResponseMessage.
type GenericError struct {
	body []byte
	err  string
}

func (e GenericError) Error() string {
	return e.err
}

func (e GenericError) Body() []byte {
	return e.body
}

// Message extracts the backend's ResponseMessage.message, falling back to the
// raw body when it is not JSON of that shape.
func (e GenericError) Message() string {
	var msg ResponseMessage
	if err := json.Unmarshal(e.body, &msg); err == nil && msg.Message != "" {
		return msg.Message
	}
	return string(e.body)
}

// do executes one request against the API, authenticating with the x-api-key
// and Harness-Account headers, and decodes a <300 JSON response into result
// (when non-nil).
func (c *APIClient) do(ctx context.Context, method, path string, body interface{}, result interface{}) (*http.Response, error) {
	var reqBody io.Reader
	if body != nil {
		b, err := json.Marshal(body)
		if err != nil {
			return nil, err
		}
		reqBody = bytes.NewReader(b)
	}

	req, err := http.NewRequestWithContext(ctx, method, c.cfg.BasePath+path, reqBody)
	if err != nil {
		return nil, err
	}
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Accept", "application/json")
	req.Header.Set("x-api-key", c.cfg.ApiKey)
	req.Header.Set("Harness-Account", c.cfg.AccountId)
	if c.cfg.UserAgent != "" {
		req.Header.Set("User-Agent", c.cfg.UserAgent)
	}
	for k, v := range c.cfg.DefaultHeader {
		req.Header.Set(k, v)
	}

	if c.cfg.DebugLogging && c.cfg.Logger != nil {
		c.cfg.Logger.Debugf("%s %s", method, req.URL.String())
	}

	resp, err := c.cfg.HTTPClient.StandardClient().Do(req)
	if err != nil {
		return resp, err
	}
	defer resp.Body.Close()

	respBody, err := io.ReadAll(resp.Body)
	if err != nil {
		return resp, err
	}

	if resp.StatusCode >= 300 {
		return resp, GenericError{
			body: respBody,
			err:  fmt.Sprintf("%s %s: %s", method, path, http.StatusText(resp.StatusCode)),
		}
	}

	if result != nil && len(respBody) > 0 {
		if err := json.Unmarshal(respBody, result); err != nil {
			return resp, err
		}
	}
	return resp, nil
}
