// Package clienttests holds offline tests for the nextgen client. It deliberately
// has no TestMain: the nextgen package's TestMain exits early unless live
// credentials are passed, which would silently skip these tests in CI.
package clienttests

import (
	"context"
	"net/http"
	"net/http/httptest"
	"net/url"
	"testing"

	"github.com/antihax/optional"
	"github.com/harness/harness-go-sdk/harness/nextgen"
	"github.com/hashicorp/go-retryablehttp"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

// CDS-136012: the gateway only learns the agent's scope from query params, and
// the server swagger omits them for this body-less POST. The client must send
// them, otherwise org/project scoped agents return 404.

func TestAgentServiceForServerRegenerateCredentialsWithScope_ScopeQueryParams(t *testing.T) {
	testCases := map[string]struct {
		opts     *nextgen.AgentsApiAgentServiceForServerRegenerateCredentialsWithScopeOpts
		expected url.Values
	}{
		"Given account scope with nil opts - then only account params are sent": {
			opts:     nil,
			expected: url.Values{"accountIdentifier": {"acc"}, "routingId": {"acc"}},
		},
		"Given org scope - then org param is sent": {
			opts: &nextgen.AgentsApiAgentServiceForServerRegenerateCredentialsWithScopeOpts{
				OrgIdentifier: optional.NewString("org1"),
			},
			expected: url.Values{"accountIdentifier": {"acc"}, "routingId": {"acc"}, "orgIdentifier": {"org1"}},
		},
		"Given project scope - then org and project params are sent": {
			opts: &nextgen.AgentsApiAgentServiceForServerRegenerateCredentialsWithScopeOpts{
				OrgIdentifier:     optional.NewString("org1"),
				ProjectIdentifier: optional.NewString("proj1"),
			},
			expected: url.Values{
				"accountIdentifier": {"acc"}, "routingId": {"acc"},
				"orgIdentifier": {"org1"}, "projectIdentifier": {"proj1"},
			},
		},
	}

	for name, tc := range testCases {
		t.Run(name, func(t *testing.T) {
			var gotMethod, gotPath string
			var gotQuery url.Values
			server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
				gotMethod, gotPath, gotQuery = r.Method, r.URL.Path, r.URL.Query()
				w.Header().Set("Content-Type", "application/json")
				_, _ = w.Write([]byte(`{"identifier":"agent1","privateKey":"pk"}`))
			}))
			defer server.Close()

			client := newClient(server.URL)
			agent, resp, err := client.AgentApi.AgentServiceForServerRegenerateCredentialsWithScope(context.Background(), "agent1", "acc", tc.opts)

			require.NoError(t, err)
			assert.Equal(t, http.StatusOK, resp.StatusCode)
			assert.Equal(t, "agent1", agent.Identifier)
			assert.Equal(t, http.MethodPost, gotMethod)
			assert.Equal(t, "/gitops/api/v1/agents/agent1/credentials", gotPath)
			assert.Equal(t, tc.expected, gotQuery)
		})
	}
}

func TestAgentServiceForServerRegenerateCredentialsWithScope_ErrorStatusReturned(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusNotFound)
	}))
	defer server.Close()

	_, resp, err := newClient(server.URL).AgentApi.AgentServiceForServerRegenerateCredentialsWithScope(context.Background(), "agent1", "acc", nil)

	require.Error(t, err)
	require.NotNil(t, resp)
	assert.Equal(t, http.StatusNotFound, resp.StatusCode)
}

// The original method is kept for backward compatibility and must keep its
// exact wire behavior: same path, no query params.
func TestAgentServiceForServerRegenerateCredentials_LegacySignatureUnchanged(t *testing.T) {
	var gotQuery url.Values
	var gotPath string
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		gotPath, gotQuery = r.URL.Path, r.URL.Query()
		w.Header().Set("Content-Type", "application/json")
		_, _ = w.Write([]byte(`{"identifier":"agent1"}`))
	}))
	defer server.Close()

	agent, _, err := newClient(server.URL).AgentApi.AgentServiceForServerRegenerateCredentials(context.Background(), "agent1")

	require.NoError(t, err)
	assert.Equal(t, "agent1", agent.Identifier)
	assert.Equal(t, "/gitops/api/v1/agents/agent1/credentials", gotPath)
	assert.Empty(t, gotQuery)
}

// newClient returns a client pointed at baseURL with retries disabled so error
// statuses are returned immediately.
func newClient(baseURL string) *nextgen.APIClient {
	httpClient := retryablehttp.NewClient()
	httpClient.RetryMax = 0
	httpClient.Logger = nil
	cfg := nextgen.NewConfiguration()
	cfg.BasePath = baseURL
	cfg.HTTPClient = httpClient
	return nextgen.NewAPIClient(cfg)
}
