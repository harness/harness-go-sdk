package nextgen

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/hashicorp/go-retryablehttp"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

// newTestClient returns an APIClient whose BasePath points at the given test
// server and whose underlying HTTP client does not retry, so error responses
// return immediately.
func newTestClient(baseURL string) *APIClient {
	httpClient := retryablehttp.NewClient()
	httpClient.RetryMax = 0
	httpClient.Logger = nil

	cfg := &Configuration{
		BasePath:      baseURL,
		DefaultHeader: make(map[string]string),
		HTTPClient:    httpClient,
	}
	return NewAPIClient(cfg)
}

// TestWorkspaceTemplatesRemoveWorkspaceTemplate_Success asserts the client
// issues a DELETE to the correct scoped path with the Harness-Account header
// and treats a 204 No Content response as success (idempotent de-association).
func TestWorkspaceTemplatesRemoveWorkspaceTemplate_Success(t *testing.T) {
	var (
		gotMethod  string
		gotPath    string
		gotAccount string
	)

	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		gotMethod = r.Method
		gotPath = r.URL.Path
		gotAccount = r.Header.Get("Harness-Account")
		w.WriteHeader(http.StatusNoContent)
	}))
	defer server.Close()

	client := newTestClient(server.URL)

	resp, err := client.WorkspaceTemplatesApi.WorkspaceTemplatesRemoveWorkspaceTemplate(
		nil, "acc1", "org1", "proj1", "template1", "workspace1",
	)

	require.NoError(t, err)
	require.NotNil(t, resp)
	assert.Equal(t, http.StatusNoContent, resp.StatusCode)
	assert.Equal(t, http.MethodDelete, gotMethod)
	assert.Equal(t, "/iacm/api/orgs/org1/projects/proj1/workspace/templates/template1/workspace1", gotPath)
	assert.Equal(t, "acc1", gotAccount)
}

// TestWorkspaceTemplatesRemoveWorkspaceTemplate_ValidationErrors asserts that
// missing required parameters are rejected client-side before any request is
// sent.
func TestWorkspaceTemplatesRemoveWorkspaceTemplate_ValidationErrors(t *testing.T) {
	requestMade := false
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		requestMade = true
		w.WriteHeader(http.StatusNoContent)
	}))
	defer server.Close()

	client := newTestClient(server.URL)

	testCases := []struct {
		name        string
		account     string
		org         string
		project     string
		templateID  string
		workspaceID string
	}{
		{name: "missing account", account: "", org: "org1", project: "proj1", templateID: "t1", workspaceID: "w1"},
		{name: "missing org", account: "acc1", org: "", project: "proj1", templateID: "t1", workspaceID: "w1"},
		{name: "missing project", account: "acc1", org: "org1", project: "", templateID: "t1", workspaceID: "w1"},
		{name: "missing template id", account: "acc1", org: "org1", project: "proj1", templateID: "", workspaceID: "w1"},
		{name: "missing workspace id", account: "acc1", org: "org1", project: "proj1", templateID: "t1", workspaceID: ""},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			requestMade = false
			_, err := client.WorkspaceTemplatesApi.WorkspaceTemplatesRemoveWorkspaceTemplate(
				nil, tc.account, tc.org, tc.project, tc.templateID, tc.workspaceID,
			)
			require.Error(t, err)
			assert.False(t, requestMade, "no HTTP request should be sent when validation fails")
		})
	}
}

// TestWorkspaceTemplatesRemoveWorkspaceTemplate_ErrorResponse asserts that a
// non-2xx response is surfaced as an error while still returning the HTTP
// response for inspection.
func TestWorkspaceTemplatesRemoveWorkspaceTemplate_ErrorResponse(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusNotFound)
		_, _ = w.Write([]byte(`{"message":"workspace template not found"}`))
	}))
	defer server.Close()

	client := newTestClient(server.URL)

	resp, err := client.WorkspaceTemplatesApi.WorkspaceTemplatesRemoveWorkspaceTemplate(
		nil, "acc1", "org1", "proj1", "template1", "workspace1",
	)

	require.Error(t, err)
	require.NotNil(t, resp)
	assert.Equal(t, http.StatusNotFound, resp.StatusCode)
}
