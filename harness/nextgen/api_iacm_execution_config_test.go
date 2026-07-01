/*
	go test \
		-v \
		-run TestIacmExecutionConfig \
		./harness/nextgen \
		-account $HARNESS_ACCOUNT_ID \
		-api-key $HARNESS_PLATFORM_API_KEY
*/

package nextgen

import (
	"context"
	"encoding/json"
	"testing"

	"github.com/stretchr/testify/require"
)

type iacmTestContext struct {
	t         *testing.T
	client    *APIClient
	ctx       context.Context
	infraType string
}

func TestIacmExecutionConfig(t *testing.T) {
	client, ctx := getClientWithContext()

	tc := &iacmTestContext{
		t:         t,
		client:    client,
		ctx:       ctx,
		infraType: "k8",
	}

	testUpdateImageType := "liteEngineTag"
	testUpdateImageTag := "harness/ci-lite-engine:updatedTmpNow"

	// reset the test field first to ensure clean state from any previous failed run
	testIacmResetConfig(tc, testUpdateImageType)
	// fetch the harness default images for the infra type
	testIacmGetDefaultConfig(tc)
	// verify no customer overrides exist before making any changes
	testIacmGetCustomerConfigOverridesOnly(tc)
	// override liteEngineTag with a custom image tag
	testIacmUpdateConfig(tc, testUpdateImageType, testUpdateImageTag)
	// confirm the override appears when fetching overrides only
	testIacmGetCustomerConfigOverridesOnlyAfterUpdate(tc, testUpdateImageType, testUpdateImageTag)
	// confirm the override appears in the full config
	testIacmGetCustomerConfigAllAfterUpdate(tc, testUpdateImageType, testUpdateImageTag)
	// reset liteEngineTag back to the harness default
	testIacmResetConfig(tc, testUpdateImageType)
	// confirm the override is gone after reset
	testIacmGetCustomerConfigOverridesOnlyAfterReset(tc, testUpdateImageType)
}

// pass: HTTP 200, status SUCCESS, data contains at least one image tag
// fail: any HTTP error, or data is nil (no images returned)
func testIacmGetDefaultConfig(tc *iacmTestContext) {
	tc.t.Log("=== GetDefaultConfig ===")
	resp, err := tc.client.IacmExecutionConfigApi.GetDefaultConfig(tc.ctx, tc.infraType)
	printIacmResult(tc.t, resp, err)
	require.NotEmpty(tc.t, resp.Data, "expected default config data, got empty map")
	require.NotEmpty(tc.t, resp.Data["liteEngineTag"], "expected liteEngineTag to be set in default config")
}

// pass: HTTP 200, status SUCCESS, data nil or liteEngineTag empty (no overrides set yet)
// fail: HTTP error, or liteEngineTag already has an override (dirty state from a previous run)
func testIacmGetCustomerConfigOverridesOnly(tc *iacmTestContext) {
	tc.t.Log("=== GetCustomerConfig before update (overridesOnly=true, expect no overrides) ===")
	resp, err := tc.client.IacmExecutionConfigApi.GetCustomerConfig(tc.ctx, tc.infraType, true)
	printIacmResult(tc.t, resp, err)
	require.True(tc.t, resp.Data == nil || resp.Data["liteEngineTag"] == "",
		"expected no liteEngineTag override before update, got %q", resp.Data["liteEngineTag"])
}

// pass: HTTP 200, status SUCCESS — the API accepted the override
// fail: any HTTP error (e.g. 400 bad field name, 401 auth, 403 forbidden)
func testIacmUpdateConfig(tc *iacmTestContext, imageField, imageTag string) {
	tc.t.Logf("=== UpdateConfig (set %s = %s) ===", imageField, imageTag)
	body := []ExecutionConfigUpdate{
		{Field: imageField, Value: imageTag},
	}
	resp, err := tc.client.IacmExecutionConfigApi.UpdateConfig(tc.ctx, tc.infraType, body)
	printIacmResult(tc.t, resp, err)
}

// pass: HTTP 200, data contains imageField set to imageTag (override is visible)
// fail: HTTP error, data nil, or imageField value does not match imageTag
func testIacmGetCustomerConfigOverridesOnlyAfterUpdate(tc *iacmTestContext, imageField, imageTag string) {
	tc.t.Log("=== GetCustomerConfig after update (overridesOnly=true, expect override present) ===")
	resp, err := tc.client.IacmExecutionConfigApi.GetCustomerConfig(tc.ctx, tc.infraType, true)
	printIacmResult(tc.t, resp, err)
	checkIacmField(tc.t, resp, imageField, imageTag)
}

// pass: HTTP 200, full config returned with imageField set to imageTag
// fail: HTTP error, data nil, or imageField value does not match imageTag
func testIacmGetCustomerConfigAllAfterUpdate(tc *iacmTestContext, imageField, imageTag string) {
	tc.t.Log("=== GetCustomerConfig after update (overridesOnly=false, expect override present) ===")
	resp, err := tc.client.IacmExecutionConfigApi.GetCustomerConfig(tc.ctx, tc.infraType, false)
	printIacmResult(tc.t, resp, err)
	checkIacmField(tc.t, resp, imageField, imageTag)
}

// pass: HTTP 200, status SUCCESS — the API accepted the reset
// fail: any HTTP error
func testIacmResetConfig(tc *iacmTestContext, imageField string) {
	tc.t.Logf("=== ResetConfig (reset %s) ===", imageField)
	body := []ExecutionConfigUpdate{
		{Field: imageField},
	}
	resp, err := tc.client.IacmExecutionConfigApi.ResetConfig(tc.ctx, tc.infraType, body)
	printIacmResult(tc.t, resp, err)
}

// pass: HTTP 200, data nil or imageField empty — override was removed by reset
// fail: HTTP error, or imageField still holds the previously set override value
func testIacmGetCustomerConfigOverridesOnlyAfterReset(tc *iacmTestContext, imageField string) {
	tc.t.Log("=== GetCustomerConfig after reset (overridesOnly=true, expect no overrides) ===")
	resp, err := tc.client.IacmExecutionConfigApi.GetCustomerConfig(tc.ctx, tc.infraType, true)
	printIacmResult(tc.t, resp, err)
	require.True(tc.t, resp.Data == nil || resp.Data[imageField] == "",
		"%s still has override value %q after reset", imageField, resp.Data[imageField])
	tc.t.Logf("PASS: %s override is gone after reset", imageField)
}

func checkIacmField(t *testing.T, resp ExecutionConfigResponse, imageField, imageTag string) {
	require.NotEmpty(t, resp.Data, "%s — response data is empty", imageField)
	require.Equal(t, imageTag, resp.Data[imageField],
		"%s mismatch: got %q, expected %q", imageField, resp.Data[imageField], imageTag)
	t.Logf("PASS: %s = %q", imageField, imageTag)
}

func printIacmResult(t *testing.T, resp ExecutionConfigResponse, err error) {
	require.NoError(t, err)
	t.Logf("Status: %s", resp.Status)
	if len(resp.Data) > 0 {
		b, _ := json.MarshalIndent(resp.Data, "", "  ")
		t.Log(string(b))
	} else {
		t.Log("Data: (empty — all defaults in use)")
	}
}
