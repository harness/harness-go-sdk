package nextgen

import (
	"bytes"
	"context"
	"io"
	"net/http"
	"os"
	"strings"
	"testing"

	"github.com/antihax/optional"
	"github.com/stretchr/testify/require"
)

/*
# To run Unit Mock Tests do the following

# Mock Unit tests:
#	- TestMockDelegateSetupResourceApi_GetLatestSupportedDelegateVersion

go test ./harness/nextgen -run TestMockDelegateSetupResourceApi -v -count=1


# To run Acceptance Tests do the following

# Acceptance tests - place real requests to server
#	- TestGetLatestSupportedDelegateVersion_AccountOnly
#	- TestGetLatestSupportedDelegateVersion_AccountAndOrg
#	- TestGetLatestSupportedDelegateVersion_AccountOrgAndProject

export HARNESS_ACCOUNT_ID='AM8HCbDiTXGQNrTIhNl7qQ'
export HARNESS_PLATFORM_ACCOUNT_LEVEL_API_KEY='xxxxxxx'
export HARNESS_PLATFORM_ORGANIZATION_LEVEL_API_KEY='xxxxxx'
export HARNESS_PLATFORM_PROJECT_LEVEL_API_KEY='xxxxxxxxxxxxxxxx'
export LATEST_DELEGATE_ORGANIZATION_ID='default'
export LATEST_DELEGATE_PROJECT_ID='senthilproj01'

go test ./harness/nextgen \
  -run TestGetLatestSupportedDelegateVersion \
  -v -count=1 \
  -args \
  -account $HARNESS_ACCOUNT_ID \
  -api-key $HARNESS_PLATFORM_ACCOUNT_LEVEL_API_KEY
*/

type roundTripperFunc func(*http.Request) (*http.Response, error)

func (f roundTripperFunc) RoundTrip(r *http.Request) (*http.Response, error) {
	return f(r)
}

func TestMockDelegateSetupResourceApi_GetLatestSupportedDelegateVersion(t *testing.T) {
	cfg := NewConfiguration()
	cfg.BasePath = "https://app.harness.io"

	var captured *http.Request
	cfg.HTTPClient.HTTPClient.Transport = roundTripperFunc(func(r *http.Request) (*http.Response, error) {
		captured = r
		return &http.Response{
			StatusCode: 200,
			Body:       io.NopCloser(strings.NewReader(`{"resource":{"latestSupportedVersion":"25.03.85100","latestSupportedMinimalVersion":"25.03.85100.minimal"}}`)),
			Header:     http.Header{"Content-Type": []string{"application/json"}},
		}, nil
	})

	client := NewAPIClient(cfg)
	ctx := context.WithValue(context.Background(), ContextAPIKey, APIKey{Key: "test-key"})

	opts := &DelegateSetupResourceApiLatestVersionOpts{
		OrgIdentifier:     optional.NewString("org1"),
		ProjectIdentifier: optional.NewString("proj1"),
	}

	resp, httpResp, err := client.DelegateSetupResourceApi.GetLatestSupportedDelegateVersion(ctx, "acc1", opts)
	require.NoError(t, err)
	require.NotNil(t, httpResp)
	require.Equal(t, 200, httpResp.StatusCode)
	require.Equal(t, "25.03.85100", resp.Resource.LatestSupportedVersion)

	require.NotNil(t, captured)
	require.Equal(t, http.MethodGet, captured.Method)
	require.Equal(t, "/ng/api/delegate-setup/latest-supported-version", captured.URL.Path)
	require.Equal(t, "acc1", captured.URL.Query().Get("accountIdentifier"))
	require.Equal(t, "org1", captured.URL.Query().Get("orgIdentifier"))
	require.Equal(t, "proj1", captured.URL.Query().Get("projectIdentifier"))
	require.Equal(t, "test-key", captured.Header.Get("x-api-key"))
}

// loggingRoundTripper wraps an existing transport to log the raw HTTP response body.
type loggingRoundTripper struct {
	t       *testing.T
	wrapped http.RoundTripper
}

func (l *loggingRoundTripper) RoundTrip(req *http.Request) (*http.Response, error) {
	resp, err := l.wrapped.RoundTrip(req)
	if err != nil || resp == nil {
		return resp, err
	}
	body, readErr := io.ReadAll(resp.Body)
	resp.Body.Close()
	if readErr == nil {
		l.t.Logf("raw response: status=%d body=%s", resp.StatusCode, string(body))
	}
	resp.Body = io.NopCloser(bytes.NewReader(body))
	return resp, nil
}

func newLatestVersionClientAndCtx(t *testing.T, apiKey string) (*APIClient, context.Context, string) {
	t.Helper()
	// NewConfiguration() automatically reads HARNESS_ACCOUNT_ID
	cfg := NewConfiguration()
	if cfg.AccountId == "" {
		t.Skip("HARNESS_ACCOUNT_ID must be set for acceptance tests")
	}
	cfg.ApiKey = apiKey
	// Wrap transport to log raw JSON response for debugging
	cfg.HTTPClient.HTTPClient.Transport = &loggingRoundTripper{
		t:       t,
		wrapped: cfg.HTTPClient.HTTPClient.Transport,
	}
	client := NewAPIClient(cfg)
	ctx := context.WithValue(context.Background(), ContextAPIKey, APIKey{Key: apiKey})
	return client, ctx, client.AccountId
}

// assertLatestVersionResponse checks that the response is a successful delegate version response.
func assertLatestVersionResponse(t *testing.T, resp RestResponseDelegateLatestVersionString, httpResp *http.Response, err error) {
	t.Helper()
	require.NoError(t, err, "API call returned an error")
	require.NotNil(t, httpResp, "HTTP response is nil")
	require.Equal(t, 200, httpResp.StatusCode, "unexpected HTTP status code")
	t.Logf("response body decoded: LatestSupportedVersion=%q LatestSupportedMinimalVersion=%q",
		resp.Resource.LatestSupportedVersion, resp.Resource.LatestSupportedMinimalVersion)
	require.NotEmpty(t, resp.Resource.LatestSupportedVersion,
		"resp.Resource.LatestSupportedVersion is empty — full resource: %+v", resp.Resource)
}

func TestGetLatestSupportedDelegateVersion_AccountOnly(t *testing.T) {
	apiKey := os.Getenv("HARNESS_PLATFORM_ACCOUNT_LEVEL_API_KEY")
	if apiKey == "" {
		t.Skip("HARNESS_PLATFORM_ACCOUNT_LEVEL_API_KEY must be set for this acceptance test")
	}
	client, ctx, accountID := newLatestVersionClientAndCtx(t, apiKey)

	resp, httpResp, err := client.DelegateSetupResourceApi.GetLatestSupportedDelegateVersion(ctx, accountID, nil)
	assertLatestVersionResponse(t, resp, httpResp, err)
}

func TestGetLatestSupportedDelegateVersion_AccountAndOrg(t *testing.T) {
	apiKey := os.Getenv("HARNESS_PLATFORM_ORGANIZATION_LEVEL_API_KEY")
	if apiKey == "" {
		t.Skip("HARNESS_PLATFORM_ORGANIZATION_LEVEL_API_KEY must be set for this acceptance test")
	}
	client, ctx, accountID := newLatestVersionClientAndCtx(t, apiKey)

	orgID := os.Getenv("LATEST_DELEGATE_ORGANIZATION_ID")
	if orgID == "" {
		t.Skip("LATEST_DELEGATE_ORGANIZATION_ID must be set for this acceptance test")
	}

	opts := &DelegateSetupResourceApiLatestVersionOpts{
		OrgIdentifier: optional.NewString(orgID),
	}

	resp, httpResp, err := client.DelegateSetupResourceApi.GetLatestSupportedDelegateVersion(ctx, accountID, opts)
	assertLatestVersionResponse(t, resp, httpResp, err)
}

func TestGetLatestSupportedDelegateVersion_AccountOrgAndProject(t *testing.T) {
	apiKey := os.Getenv("HARNESS_PLATFORM_PROJECT_LEVEL_API_KEY")
	if apiKey == "" {
		t.Skip("HARNESS_PLATFORM_PROJECT_LEVEL_API_KEY must be set for this acceptance test")
	}
	client, ctx, accountID := newLatestVersionClientAndCtx(t, apiKey)

	orgID := os.Getenv("LATEST_DELEGATE_ORGANIZATION_ID")
	projectID := os.Getenv("LATEST_DELEGATE_PROJECT_ID")
	if orgID == "" || projectID == "" {
		t.Skip("LATEST_DELEGATE_ORGANIZATION_ID and LATEST_DELEGATE_PROJECT_ID must be set for this acceptance test")
	}

	opts := &DelegateSetupResourceApiLatestVersionOpts{
		OrgIdentifier:     optional.NewString(orgID),
		ProjectIdentifier: optional.NewString(projectID),
	}

	resp, httpResp, err := client.DelegateSetupResourceApi.GetLatestSupportedDelegateVersion(ctx, accountID, opts)
	assertLatestVersionResponse(t, resp, httpResp, err)
}
