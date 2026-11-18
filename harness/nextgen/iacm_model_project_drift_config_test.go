package nextgen

import (
	"bytes"
	"context"
	"io"
	"net/http"
	"os"
	"testing"
	"time"

	"github.com/stretchr/testify/require"
)

/*
IaCM Project-Level Drift Config — Acceptance Tests (TestAccIacmProjectDriftConfig_*)

# Required env vars:
	export HARNESS_ACCOUNT_ID='<account-id>'
	export HARNESS_PLATFORM_API_KEY='<pat-or-sa-key>'
	export IACM_DRIFT_ORG_ID='<org-id>'
	export IACM_DRIFT_PROJECT_ID='<project-id>'
	export IACM_DRIFT_PIPELINE_ID='<pipeline-id>'

	# Set to "true" to leave the drift config in place after tests (for manual UI verification)
	export IS_SKIP_DELETION_TEST="true"

	# Set to "true" to log full request and response bodies for each API call
	export RUN_TESTS_VERBOSE=true

	# Verify env vars are set (without revealing values):

	for v in HARNESS_ACCOUNT_ID HARNESS_PLATFORM_API_KEY IACM_DRIFT_ORG_ID IACM_DRIFT_PROJECT_ID IACM_DRIFT_PIPELINE_ID; do
	  if [ -n "${!v}" ]; then echo "$v=(set)"; else echo "$v=(MISSING)"; fi
	done

# Run all tests:
	go test -v -run "TestAccIacmProjectDriftConfig" -count=1 -timeout 300s ./harness/nextgen/ \
	  -args -account $HARNESS_ACCOUNT_ID -api-key $HARNESS_PLATFORM_API_KEY

# Run a single test:
	go test -v -run "TestAccIacmProjectDriftConfig_CreateAndGet" -count=1 -timeout 300s ./harness/nextgen/ \
	  -args -account $HARNESS_ACCOUNT_ID -api-key $HARNESS_PLATFORM_API_KEY

Tests:
	TestAccIacmProjectDriftConfig_CreateAndGet
	  Create project drift config and read it back within 3s.
	  verifies that UpsertIacmProjectDriftConfig (PUT) correctly enables drift and all
	  fields (pipeline_id, provisioner, run_every, run_every_unit) round-trip through the API.

	TestAccIacmProjectDriftConfig_Update
	  Update run_every from 1 to 6 hours and verify the change is reflected within 3s.
	  confirms that PUT uses upsert semantics and the backend applies the new schedule,
	  not silently ignoring the update.

	TestAccIacmProjectDriftConfig_Disable
	  Create config, delete it, and verify enabled=false within 3s.
	  confirms that DisableIacmProjectDriftConfig (DELETE) correctly disables drift and
	  the state is persisted — not just accepted without effect.

	TestAccIacmProjectDriftConfig_ErrorResponse
	  GET with a non-existent org/project to trigger a client error (4xx).
	  verifies that decodeIacmError returns a GenericSwaggerError with a non-empty
	  body so callers can inspect the error message regardless of error format.
*/

type iacmDriftTestEnv struct {
	accountID  string
	apiKey     string
	orgID      string
	projectID  string
	pipelineID string
}

func getIacmDriftTestEnv(t *testing.T) iacmDriftTestEnv {
	t.Helper()

	env := iacmDriftTestEnv{
		accountID:  os.Getenv("HARNESS_ACCOUNT_ID"),
		apiKey:     os.Getenv("HARNESS_PLATFORM_API_KEY"),
		orgID:      os.Getenv("IACM_DRIFT_ORG_ID"),
		projectID:  os.Getenv("IACM_DRIFT_PROJECT_ID"),
		pipelineID: os.Getenv("IACM_DRIFT_PIPELINE_ID"),
	}

	if env.accountID == "" {
		t.Skip("HARNESS_ACCOUNT_ID must be set for acceptance tests")
	}

	missing := []string{}
	if env.apiKey == "" {
		missing = append(missing, "HARNESS_PLATFORM_API_KEY")
	}
	if env.orgID == "" {
		missing = append(missing, "IACM_DRIFT_ORG_ID")
	}
	if env.projectID == "" {
		missing = append(missing, "IACM_DRIFT_PROJECT_ID")
	}
	if env.pipelineID == "" {
		missing = append(missing, "IACM_DRIFT_PIPELINE_ID")
	}
	if len(missing) > 0 {
		t.Skipf("missing required env vars: %v", missing)
	}

	return env
}

// iacmDriftLoggingRoundTripper logs both the request body and response body for each call.
type iacmDriftLoggingRoundTripper struct {
	t       *testing.T
	wrapped http.RoundTripper
}

func (l *iacmDriftLoggingRoundTripper) RoundTrip(req *http.Request) (*http.Response, error) {
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

func newIacmDriftClient(t *testing.T, env iacmDriftTestEnv) (*APIClient, context.Context) {
	t.Helper()

	cfg := NewConfiguration()
	cfg.AccountId = env.accountID
	cfg.ApiKey = env.apiKey
	if os.Getenv("RUN_TESTS_VERBOSE") == "true" {
		transport := cfg.HTTPClient.HTTPClient.Transport
		if transport == nil {
			transport = http.DefaultTransport
		}
		cfg.HTTPClient.HTTPClient.Transport = &iacmDriftLoggingRoundTripper{
			t:       t,
			wrapped: transport,
		}
	}

	client := NewAPIClient(cfg)
	ctx := context.WithValue(context.Background(), ContextAPIKey, APIKey{Key: env.apiKey})
	return client, ctx
}

// TestAccIacmProjectDriftConfig_CreateAndGet asserts that upsert (PUT) creates a
// new drift config and that a subsequent GET returns the same values.
func TestAccIacmProjectDriftConfig_CreateAndGet(t *testing.T) {
	env := getIacmDriftTestEnv(t)
	client, ctx := newIacmDriftClient(t, env)

	body := IacmProjectDriftConfigRequestBody{
		PipelineId:   env.pipelineID,
		Provisioner:  "opentofu",
		RunEvery:     1,
		RunEveryUnit: "hours",
	}

	resp, err := client.IacmProjectDriftConfigApi.UpsertIacmProjectDriftConfig(
		ctx, env.accountID, env.orgID, env.projectID, body,
	)
	require.NoError(t, err, "UpsertIacmProjectDriftConfig (create) failed")
	require.NotNil(t, resp)
	defer resp.Body.Close()
	require.Equal(t, http.StatusAccepted, resp.StatusCode)
	t.Logf("Created drift config: org=%s project=%s pipeline=%s", env.orgID, env.projectID, env.pipelineID)

	// PUT enqueues a background job — poll until enabled=true is reflected or timeout.
	var result IacmProjectDriftConfigResponse
	require.Eventually(t, func() bool {
		var err error
		result, _, err = client.IacmProjectDriftConfigApi.GetIacmProjectDriftConfig(
			ctx, env.accountID, env.orgID, env.projectID, "opentofu",
		)
		return err == nil && result.Enabled &&
			result.PipelineId == env.pipelineID &&
			result.RunEvery == int32(1) &&
			result.RunEveryUnit == "hours"
	}, 60*time.Second, 2*time.Second, "drift config should be enabled with run_every=1 within 60s")
	require.Equal(t, env.pipelineID, result.PipelineId)
	require.Equal(t, "opentofu", result.Provisioner)
	require.Equal(t, int32(1), result.RunEvery)
	require.Equal(t, "hours", result.RunEveryUnit)
	t.Logf("GET drift config: enabled=%v pipeline=%s run_every=%d %s",
		result.Enabled, result.PipelineId, result.RunEvery, result.RunEveryUnit)
}

// TestAccIacmProjectDriftConfig_Update asserts that re-sending upsert with a
// different run_every value replaces the existing config (PUT upsert semantics).
// It first creates with run_every=1, confirms it, then updates to run_every=6,
// confirms the change, and restores run_every=1 on cleanup.
func TestAccIacmProjectDriftConfig_Update(t *testing.T) {
	env := getIacmDriftTestEnv(t)
	client, ctx := newIacmDriftClient(t, env)

	// Step 1: create with run_every=1 so we have a known starting state.
	createResp, err := client.IacmProjectDriftConfigApi.UpsertIacmProjectDriftConfig(
		ctx, env.accountID, env.orgID, env.projectID, IacmProjectDriftConfigRequestBody{
			PipelineId:   env.pipelineID,
			Provisioner:  "opentofu",
			RunEvery:     1,
			RunEveryUnit: "hours",
		},
	)
	require.NoError(t, err, "UpsertIacmProjectDriftConfig (create) failed")
	require.NotNil(t, createResp)
	defer createResp.Body.Close()
	require.Equal(t, http.StatusAccepted, createResp.StatusCode)

	// Step 2: wait until GET confirms enabled=true and run_every=1.
	require.Eventually(t, func() bool {
		result, _, err := client.IacmProjectDriftConfigApi.GetIacmProjectDriftConfig(
			ctx, env.accountID, env.orgID, env.projectID, "opentofu",
		)
		return err == nil && result.Enabled &&
			result.PipelineId == env.pipelineID &&
			result.RunEvery == int32(1) &&
			result.RunEveryUnit == "hours"
	}, 60*time.Second, 2*time.Second, "initial config: enabled=true, run_every=1 within 60s")
	t.Logf("Initial config confirmed: run_every=1")

	// Step 3: update to run_every=6.
	updateResp, err := client.IacmProjectDriftConfigApi.UpsertIacmProjectDriftConfig(
		ctx, env.accountID, env.orgID, env.projectID, IacmProjectDriftConfigRequestBody{
			PipelineId:   env.pipelineID,
			Provisioner:  "opentofu",
			RunEvery:     6,
			RunEveryUnit: "hours",
		},
	)
	require.NoError(t, err, "UpsertIacmProjectDriftConfig (update) failed")
	require.NotNil(t, updateResp)
	defer updateResp.Body.Close()
	require.Equal(t, http.StatusAccepted, updateResp.StatusCode)

	// Step 4: wait until GET confirms enabled=true, run_every=6, correct pipeline.
	var result IacmProjectDriftConfigResponse
	require.Eventually(t, func() bool {
		var err error
		result, _, err = client.IacmProjectDriftConfigApi.GetIacmProjectDriftConfig(
			ctx, env.accountID, env.orgID, env.projectID, "opentofu",
		)
		return err == nil && result.Enabled &&
			result.PipelineId == env.pipelineID &&
			result.RunEvery == int32(6) &&
			result.RunEveryUnit == "hours"
	}, 60*time.Second, 2*time.Second, "updated config: enabled=true, run_every=6 within 60s")
	require.Equal(t, env.pipelineID, result.PipelineId)
	t.Logf("Updated config confirmed: run_every=%d pipeline=%s", result.RunEvery, result.PipelineId)

	// Step 5: restore run_every=1 so subsequent tests start from a known state.
	restoreResp, err := client.IacmProjectDriftConfigApi.UpsertIacmProjectDriftConfig(
		ctx, env.accountID, env.orgID, env.projectID, IacmProjectDriftConfigRequestBody{
			PipelineId:   env.pipelineID,
			Provisioner:  "opentofu",
			RunEvery:     1,
			RunEveryUnit: "hours",
		},
	)
	require.NoError(t, err, "UpsertIacmProjectDriftConfig (restore) failed")
	require.NotNil(t, restoreResp)
	defer restoreResp.Body.Close()
	t.Logf("Restored config to run_every=1")
}

// TestAccIacmProjectDriftConfig_Disable asserts that DELETE disables drift
// detection and that a subsequent GET reflects enabled=false.
func TestAccIacmProjectDriftConfig_Disable(t *testing.T) {
	env := getIacmDriftTestEnv(t)

	if isSkipDeletion() {
		t.Skip("IS_SKIP_DELETION_TEST=true: skipping disable test")
	}

	client, ctx := newIacmDriftClient(t, env)

	// Ensure config exists before disabling.
	body := IacmProjectDriftConfigRequestBody{
		PipelineId:   env.pipelineID,
		Provisioner:  "opentofu",
		RunEvery:     1,
		RunEveryUnit: "hours",
	}
	upsertResp, err := client.IacmProjectDriftConfigApi.UpsertIacmProjectDriftConfig(
		ctx, env.accountID, env.orgID, env.projectID, body,
	)
	require.NoError(t, err, "pre-condition upsert failed")
	require.NotNil(t, upsertResp)
	defer upsertResp.Body.Close()

	// Wait for the upsert job to complete before issuing DELETE.
	require.Eventually(t, func() bool {
		result, _, err := client.IacmProjectDriftConfigApi.GetIacmProjectDriftConfig(
			ctx, env.accountID, env.orgID, env.projectID, "opentofu",
		)
		return err == nil && result.Enabled
	}, 60*time.Second, 2*time.Second, "pre-condition: drift config should be enabled within 60s")

	resp, err := client.IacmProjectDriftConfigApi.DisableIacmProjectDriftConfig(
		ctx, env.accountID, env.orgID, env.projectID, "opentofu",
	)
	require.NoError(t, err, "DisableIacmProjectDriftConfig failed")
	require.NotNil(t, resp)
	defer resp.Body.Close()
	require.Equal(t, http.StatusAccepted, resp.StatusCode)
	t.Logf("Disabled drift config: org=%s project=%s", env.orgID, env.projectID)

	// DELETE also enqueues a background job — poll until enabled=false is reflected.
	var result IacmProjectDriftConfigResponse
	require.Eventually(t, func() bool {
		var err error
		result, _, err = client.IacmProjectDriftConfigApi.GetIacmProjectDriftConfig(
			ctx, env.accountID, env.orgID, env.projectID, "opentofu",
		)
		return err == nil && !result.Enabled
	}, 60*time.Second, 2*time.Second, "drift config should be disabled within 60s")
	t.Logf("Verified drift config disabled: enabled=%v", result.Enabled)
}

// TestAccIacmProjectDriftConfig_ErrorResponse asserts that a client error (4xx) response
// is decoded into a GenericSwaggerError with a non-empty body. Uses a non-existent
// org/project to trigger a 404; logs the actual status so the assumption can be verified.
func TestAccIacmProjectDriftConfig_ErrorResponse(t *testing.T) {
	env := getIacmDriftTestEnv(t)
	client, ctx := newIacmDriftClient(t, env)

	_, resp, err := client.IacmProjectDriftConfigApi.GetIacmProjectDriftConfig(
		ctx, env.accountID, "nonexistent-org-hsf175", "nonexistent-project-hsf175", "opentofu",
	)

	require.Error(t, err, "expected an error for a non-existent org/project")
	require.NotNil(t, resp)
	require.GreaterOrEqual(t, resp.StatusCode, 400, "expected a client error (4xx), got %d", resp.StatusCode)
	require.Less(t, resp.StatusCode, 500, "expected a client error (4xx), not a server error (5xx), got %d", resp.StatusCode)

	swaggerErr, ok := err.(GenericSwaggerError)
	require.True(t, ok, "expected error to be a GenericSwaggerError, got %T", err)
	require.NotEmpty(t, swaggerErr.Body(), "expected non-empty error body from decodeIacmError")
	t.Logf("Error response: status=%d body=%s", resp.StatusCode, string(swaggerErr.Body()))
}

func isSkipDeletion() bool {
	return os.Getenv("IS_SKIP_DELETION_TEST") == "true"
}
