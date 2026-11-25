package nextgen

import (
	"bytes"
	"context"
	"io"
	"net/http"
	"os"
	"testing"

	"github.com/stretchr/testify/require"
)

/*
IaCM Workspace-Level Drift Config — Acceptance Tests (TestAccIacmWorkspaceDriftConfig_*)

# Required env vars:
	export HARNESS_ACCOUNT_ID='<account-id>'
	export HARNESS_PLATFORM_API_KEY='<pat-or-sa-key>'
	export IACM_WORKSPACE_DRIFT_ORG_ID='<org-id>'
	export IACM_WORKSPACE_DRIFT_PROJECT_ID='<project-id>'
	export IACM_WORKSPACE_DRIFT_WORKSPACE_ID='<workspace-id>'
	export IACM_WORKSPACE_DRIFT_PIPELINE_ID='<pipeline-id>'

	# Set to "true" to leave the drift config in place after tests (for manual UI verification)
	export IS_SKIP_DELETION_TEST="true"

	# Set to "true" to log full request and response bodies for each API call
	export RUN_TESTS_VERBOSE=true

	# Verify env vars are set (without revealing values):

	for v in HARNESS_ACCOUNT_ID HARNESS_PLATFORM_API_KEY IACM_WORKSPACE_DRIFT_ORG_ID IACM_WORKSPACE_DRIFT_PROJECT_ID IACM_WORKSPACE_DRIFT_WORKSPACE_ID IACM_WORKSPACE_DRIFT_PIPELINE_ID; do
	  if [ -n "${!v}" ]; then echo "$v=(set)"; else echo "$v=(MISSING)"; fi
	done

# Run all tests:
	go test -v -run "TestAccIacmWorkspaceDriftConfig" -count=1 -timeout 300s ./harness/nextgen/ \
	  -args -account $HARNESS_ACCOUNT_ID -api-key $HARNESS_PLATFORM_API_KEY

# Run a single test:
	go test -v -run "TestAccIacmWorkspaceDriftConfig_CreateAndGet" -count=1 -timeout 300s ./harness/nextgen/ \
	  -args -account $HARNESS_ACCOUNT_ID -api-key $HARNESS_PLATFORM_API_KEY

Tests:
	TestAccIacmWorkspaceDriftConfig_CreateAndGet
	  Create workspace drift config and read it back.
	  Verifies that UpsertIacmWorkspaceDriftConfig (PUT) correctly enables drift and all
	  fields (pipeline_id, provisioner, run_every, run_every_unit) round-trip through the API.
	  PUT returns 200 OK (synchronous — no polling needed).

	TestAccIacmWorkspaceDriftConfig_Update
	  Update run_every from 1 to 6 hours and verify the change is reflected.
	  Confirms that PUT uses upsert semantics and the backend applies the new schedule.

	TestAccIacmWorkspaceDriftConfig_Disable
	  Create config, delete it, and verify enabled=false.
	  Confirms that DisableIacmWorkspaceDriftConfig (DELETE) correctly disables drift.
	  DELETE returns 200 OK (synchronous).

	TestAccIacmWorkspaceDriftConfig_ErrorResponse
	  GET with a non-existent org/project/workspace to trigger a client error (4xx).
	  Verifies that decodeIacmError returns a GenericSwaggerError with a non-empty body.
*/

type iacmWorkspaceDriftTestEnv struct {
	accountID   string
	apiKey      string
	orgID       string
	projectID   string
	workspaceID string
	pipelineID  string
}

func getIacmWorkspaceDriftTestEnv(t *testing.T) iacmWorkspaceDriftTestEnv {
	t.Helper()

	env := iacmWorkspaceDriftTestEnv{
		accountID:   os.Getenv("HARNESS_ACCOUNT_ID"),
		apiKey:      os.Getenv("HARNESS_PLATFORM_API_KEY"),
		orgID:       os.Getenv("IACM_WORKSPACE_DRIFT_ORG_ID"),
		projectID:   os.Getenv("IACM_WORKSPACE_DRIFT_PROJECT_ID"),
		workspaceID: os.Getenv("IACM_WORKSPACE_DRIFT_WORKSPACE_ID"),
		pipelineID:  os.Getenv("IACM_WORKSPACE_DRIFT_PIPELINE_ID"),
	}

	if env.accountID == "" {
		t.Skip("HARNESS_ACCOUNT_ID must be set for acceptance tests")
	}

	missing := []string{}
	if env.apiKey == "" {
		missing = append(missing, "HARNESS_PLATFORM_API_KEY")
	}
	if env.orgID == "" {
		missing = append(missing, "IACM_WORKSPACE_DRIFT_ORG_ID")
	}
	if env.projectID == "" {
		missing = append(missing, "IACM_WORKSPACE_DRIFT_PROJECT_ID")
	}
	if env.workspaceID == "" {
		missing = append(missing, "IACM_WORKSPACE_DRIFT_WORKSPACE_ID")
	}
	if env.pipelineID == "" {
		missing = append(missing, "IACM_WORKSPACE_DRIFT_PIPELINE_ID")
	}
	if len(missing) > 0 {
		t.Skipf("missing required env vars: %v", missing)
	}

	return env
}

// iacmWorkspaceDriftLoggingRoundTripper logs both the request body and response body for each call.
type iacmWorkspaceDriftLoggingRoundTripper struct {
	t       *testing.T
	wrapped http.RoundTripper
}

func (l *iacmWorkspaceDriftLoggingRoundTripper) RoundTrip(req *http.Request) (*http.Response, error) {
	if req.Body != nil {
		reqBody, err := io.ReadAll(req.Body)
		if err != nil {
			l.t.Logf(">>> %s %s (failed to read request body: %v)", req.Method, req.URL.RequestURI(), err)
		} else {
			req.Body = io.NopCloser(bytes.NewReader(reqBody))
			l.t.Logf(">>> %s %s body=%s", req.Method, req.URL.RequestURI(), string(reqBody))
		}
	} else {
		l.t.Logf(">>> %s %s", req.Method, req.URL.RequestURI())
	}

	resp, err := l.wrapped.RoundTrip(req)
	if err != nil || resp == nil {
		return resp, err
	}

	respBody, err := io.ReadAll(resp.Body)
	if err != nil {
		l.t.Logf("<<< status=%d (failed to read response body: %v)", resp.StatusCode, err)
	} else {
		resp.Body = io.NopCloser(bytes.NewReader(respBody))
		l.t.Logf("<<< status=%d body=%s", resp.StatusCode, string(respBody))
	}
	return resp, nil
}

func newIacmWorkspaceDriftClient(t *testing.T, env iacmWorkspaceDriftTestEnv) (*APIClient, context.Context) {
	t.Helper()

	cfg := NewConfiguration()
	cfg.AccountId = env.accountID
	cfg.ApiKey = env.apiKey
	if os.Getenv("RUN_TESTS_VERBOSE") == "true" {
		transport := cfg.HTTPClient.HTTPClient.Transport
		if transport == nil {
			transport = http.DefaultTransport
		}
		cfg.HTTPClient.HTTPClient.Transport = &iacmWorkspaceDriftLoggingRoundTripper{
			t:       t,
			wrapped: transport,
		}
	}

	client := NewAPIClient(cfg)
	ctx := context.WithValue(context.Background(), ContextAPIKey, APIKey{Key: env.apiKey})
	return client, ctx
}

// TestAccIacmWorkspaceDriftConfig_CreateAndGet asserts that upsert (PUT) creates a
// new workspace drift config and that a subsequent GET returns the same values.
// PUT returns 200 OK synchronously — no polling needed.
func TestAccIacmWorkspaceDriftConfig_CreateAndGet(t *testing.T) {
	env := getIacmWorkspaceDriftTestEnv(t)
	client, ctx := newIacmWorkspaceDriftClient(t, env)

	body := IacmWorkspaceDriftConfigRequestBody{
		PipelineId:   env.pipelineID,
		Provisioner:  "opentofu",
		RunEvery:     1,
		RunEveryUnit: "hours",
	}

	resp, err := client.IacmWorkspaceDriftConfigApi.UpsertIacmWorkspaceDriftConfig(
		ctx, env.accountID, env.orgID, env.projectID, env.workspaceID, body,
	)
	require.NoError(t, err, "UpsertIacmWorkspaceDriftConfig (create) failed")
	require.NotNil(t, resp)
	defer resp.Body.Close()
	require.Equal(t, http.StatusOK, resp.StatusCode)
	t.Logf("Created workspace drift config: org=%s project=%s workspace=%s pipeline=%s",
		env.orgID, env.projectID, env.workspaceID, env.pipelineID)

	result, _, err := client.IacmWorkspaceDriftConfigApi.GetIacmWorkspaceDriftConfig(
		ctx, env.accountID, env.orgID, env.projectID, env.workspaceID, "opentofu",
	)
	require.NoError(t, err, "GetIacmWorkspaceDriftConfig failed")
	require.True(t, result.Enabled, "expected enabled=true after upsert")
	require.Equal(t, env.pipelineID, result.PipelineId)
	require.Equal(t, int32(1), result.RunEvery)
	require.Equal(t, "hours", result.RunEveryUnit)
	t.Logf("GET workspace drift config: enabled=%v pipeline=%s run_every=%d %s trigger_id=%s",
		result.Enabled, result.PipelineId, result.RunEvery, result.RunEveryUnit, result.TriggerId)
}

// TestAccIacmWorkspaceDriftConfig_Update asserts that re-sending upsert with a
// different run_every value replaces the existing config (PUT upsert semantics).
// Creates with run_every=1, confirms, updates to run_every=6, confirms, restores to run_every=1.
func TestAccIacmWorkspaceDriftConfig_Update(t *testing.T) {
	env := getIacmWorkspaceDriftTestEnv(t)
	client, ctx := newIacmWorkspaceDriftClient(t, env)

	// Step 1: create with run_every=1 so we have a known starting state.
	createResp, err := client.IacmWorkspaceDriftConfigApi.UpsertIacmWorkspaceDriftConfig(
		ctx, env.accountID, env.orgID, env.projectID, env.workspaceID,
		IacmWorkspaceDriftConfigRequestBody{
			PipelineId:   env.pipelineID,
			Provisioner:  "opentofu",
			RunEvery:     1,
			RunEveryUnit: "hours",
		},
	)
	require.NoError(t, err, "UpsertIacmWorkspaceDriftConfig (create) failed")
	require.NotNil(t, createResp)
	defer createResp.Body.Close()
	require.Equal(t, http.StatusOK, createResp.StatusCode)

	// Step 2: verify GET confirms enabled=true and run_every=1.
	result, _, err := client.IacmWorkspaceDriftConfigApi.GetIacmWorkspaceDriftConfig(
		ctx, env.accountID, env.orgID, env.projectID, env.workspaceID, "opentofu",
	)
	require.NoError(t, err)
	require.True(t, result.Enabled)
	require.Equal(t, int32(1), result.RunEvery)
	t.Logf("Initial config confirmed: run_every=1")

	// Step 3: update to run_every=6.
	updateResp, err := client.IacmWorkspaceDriftConfigApi.UpsertIacmWorkspaceDriftConfig(
		ctx, env.accountID, env.orgID, env.projectID, env.workspaceID,
		IacmWorkspaceDriftConfigRequestBody{
			PipelineId:   env.pipelineID,
			Provisioner:  "opentofu",
			RunEvery:     6,
			RunEveryUnit: "hours",
		},
	)
	require.NoError(t, err, "UpsertIacmWorkspaceDriftConfig (update) failed")
	require.NotNil(t, updateResp)
	defer updateResp.Body.Close()
	require.Equal(t, http.StatusOK, updateResp.StatusCode)

	// Step 4: verify GET confirms run_every=6.
	result, _, err = client.IacmWorkspaceDriftConfigApi.GetIacmWorkspaceDriftConfig(
		ctx, env.accountID, env.orgID, env.projectID, env.workspaceID, "opentofu",
	)
	require.NoError(t, err)
	require.True(t, result.Enabled)
	require.Equal(t, int32(6), result.RunEvery)
	require.Equal(t, env.pipelineID, result.PipelineId)
	t.Logf("Updated config confirmed: run_every=%d pipeline=%s", result.RunEvery, result.PipelineId)

	// Step 5: restore run_every=1 so subsequent tests start from a known state.
	restoreResp, err := client.IacmWorkspaceDriftConfigApi.UpsertIacmWorkspaceDriftConfig(
		ctx, env.accountID, env.orgID, env.projectID, env.workspaceID,
		IacmWorkspaceDriftConfigRequestBody{
			PipelineId:   env.pipelineID,
			Provisioner:  "opentofu",
			RunEvery:     1,
			RunEveryUnit: "hours",
		},
	)
	require.NoError(t, err, "UpsertIacmWorkspaceDriftConfig (restore) failed")
	require.NotNil(t, restoreResp)
	defer restoreResp.Body.Close()
	t.Logf("Restored config to run_every=1")
}

// TestAccIacmWorkspaceDriftConfig_Disable asserts that DELETE disables workspace drift
// detection and that a subsequent GET reflects enabled=false.
// DELETE returns 200 OK synchronously — no polling needed.
func TestAccIacmWorkspaceDriftConfig_Disable(t *testing.T) {
	env := getIacmWorkspaceDriftTestEnv(t)

	if isSkipDeletion() {
		t.Skip("IS_SKIP_DELETION_TEST=true: skipping disable test")
	}

	client, ctx := newIacmWorkspaceDriftClient(t, env)

	// Ensure config exists before disabling.
	upsertResp, err := client.IacmWorkspaceDriftConfigApi.UpsertIacmWorkspaceDriftConfig(
		ctx, env.accountID, env.orgID, env.projectID, env.workspaceID,
		IacmWorkspaceDriftConfigRequestBody{
			PipelineId:   env.pipelineID,
			Provisioner:  "opentofu",
			RunEvery:     1,
			RunEveryUnit: "hours",
		},
	)
	require.NoError(t, err, "pre-condition upsert failed")
	require.NotNil(t, upsertResp)
	defer upsertResp.Body.Close()
	require.Equal(t, http.StatusOK, upsertResp.StatusCode)

	resp, err := client.IacmWorkspaceDriftConfigApi.DisableIacmWorkspaceDriftConfig(
		ctx, env.accountID, env.orgID, env.projectID, env.workspaceID, "opentofu",
	)
	require.NoError(t, err, "DisableIacmWorkspaceDriftConfig failed")
	require.NotNil(t, resp)
	defer resp.Body.Close()
	require.Equal(t, http.StatusOK, resp.StatusCode)
	t.Logf("Disabled workspace drift config: org=%s project=%s workspace=%s",
		env.orgID, env.projectID, env.workspaceID)

	result, _, err := client.IacmWorkspaceDriftConfigApi.GetIacmWorkspaceDriftConfig(
		ctx, env.accountID, env.orgID, env.projectID, env.workspaceID, "opentofu",
	)
	require.NoError(t, err, "GetIacmWorkspaceDriftConfig after disable failed")
	require.False(t, result.Enabled, "expected enabled=false after disable")
	t.Logf("Verified workspace drift config disabled: enabled=%v", result.Enabled)
}

// TestAccIacmWorkspaceDriftConfig_ErrorResponse asserts that a client error (4xx) response
// is decoded into a GenericSwaggerError with a non-empty body. Uses a non-existent
// org/project/workspace to trigger a 404.
func TestAccIacmWorkspaceDriftConfig_ErrorResponse(t *testing.T) {
	env := getIacmWorkspaceDriftTestEnv(t)
	client, ctx := newIacmWorkspaceDriftClient(t, env)

	_, resp, err := client.IacmWorkspaceDriftConfigApi.GetIacmWorkspaceDriftConfig(
		ctx, env.accountID, "nonexistent-org-hsf180", "nonexistent-project-hsf180", "nonexistent-workspace-hsf180", "opentofu",
	)

	require.Error(t, err, "expected an error for a non-existent org/project/workspace")
	require.NotNil(t, resp)
	require.GreaterOrEqual(t, resp.StatusCode, 400, "expected a client error (4xx), got %d", resp.StatusCode)
	require.Less(t, resp.StatusCode, 500, "expected a client error (4xx), not a server error (5xx), got %d", resp.StatusCode)

	swaggerErr, ok := err.(GenericSwaggerError)
	require.True(t, ok, "expected error to be a GenericSwaggerError, got %T", err)
	require.NotEmpty(t, swaggerErr.Body(), "expected non-empty error body from decodeIacmError")
	t.Logf("Error response: status=%d body=%s", resp.StatusCode, string(swaggerErr.Body()))
}
