/*
	go test \
		-v \
		-run TestCiExecutionConfig \
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

type ciTestContext struct {
	t         *testing.T
	client    *APIClient
	ctx       context.Context
	infraType string
}

func TestCiExecutionConfig(t *testing.T) {
	client, ctx := getClientWithContext()

	tc := &ciTestContext{
		t:         t,
		client:    client,
		ctx:       ctx,
		infraType: "K8",
	}

	testUpdateImageType := "liteEngineTag"
	testUpdateImageTag := "harness/ci-lite-engine:updatedTmpNow"

	// reset the test field first to ensure clean state from any previous failed run
	testCiResetConfig(tc, testUpdateImageType)
	testCiGetDefaultConfig(tc)
	testCiGetCustomerConfigOverridesOnly(tc)

	testCiUpdateConfig(tc, testUpdateImageType, testUpdateImageTag)
	testCiGetCustomerConfigOverridesOnlyAfterUpdate(tc, testUpdateImageType, testUpdateImageTag)
	testCiGetCustomerConfigAllAfterUpdate(tc, testUpdateImageType, testUpdateImageTag)

	testCiResetConfig(tc, testUpdateImageType)
	testCiGetCustomerConfigOverridesOnlyAfterReset(tc, testUpdateImageType)
}

// pass: HTTP 200, status SUCCESS, data contains at least one image tag
// fail: any HTTP error, or data is nil (no images returned)
func testCiGetDefaultConfig(tc *ciTestContext) {
	tc.t.Log("=== GetDefaultConfig ===")
	resp, err := tc.client.CiExecutionConfigApi.GetDefaultConfig(tc.ctx, tc.infraType)
	printCiResult(tc.t, resp, err)
	require.NotEmpty(tc.t, resp.Data, "expected default config data, got empty map")
	require.NotEmpty(tc.t, resp.Data["liteEngineTag"], "expected liteEngineTag to be set in default config")
}

// pass: HTTP 200, status SUCCESS, data nil or liteEngineTag empty (no overrides set yet)
// fail: HTTP error, or liteEngineTag already has an override (dirty state from a previous run)
func testCiGetCustomerConfigOverridesOnly(tc *ciTestContext) {
	tc.t.Log("=== GetCustomerConfig before update (overridesOnly=true, expect no overrides) ===")
	resp, err := tc.client.CiExecutionConfigApi.GetCustomerConfig(tc.ctx, tc.infraType, true)
	printCiResult(tc.t, resp, err)
	require.True(tc.t, resp.Data == nil || resp.Data["liteEngineTag"] == "",
		"expected no liteEngineTag override before update, got %q", resp.Data["liteEngineTag"])
}

// pass: HTTP 200, status SUCCESS — the API accepted the override
// fail: any HTTP error (e.g. 400 bad field name, 401 auth, 403 forbidden)
func testCiUpdateConfig(tc *ciTestContext, imageField, imageTag string) {
	tc.t.Logf("=== UpdateConfig (set %s = %s) ===", imageField, imageTag)
	body := []ExecutionConfigUpdate{
		{Field: imageField, Value: imageTag},
	}
	resp, err := tc.client.CiExecutionConfigApi.UpdateConfig(tc.ctx, tc.infraType, body)
	printCiResult(tc.t, resp, err)
}

// pass: HTTP 200, data contains imageField set to imageTag (override is visible)
// fail: HTTP error, data nil, or imageField value does not match imageTag
func testCiGetCustomerConfigOverridesOnlyAfterUpdate(tc *ciTestContext, imageField, imageTag string) {
	tc.t.Log("=== GetCustomerConfig after update (overridesOnly=true, expect override present) ===")
	resp, err := tc.client.CiExecutionConfigApi.GetCustomerConfig(tc.ctx, tc.infraType, true)
	printCiResult(tc.t, resp, err)
	checkCiField(tc.t, resp, imageField, imageTag)
}

// pass: HTTP 200, full config returned with imageField set to imageTag
// fail: HTTP error, data nil, or imageField value does not match imageTag
func testCiGetCustomerConfigAllAfterUpdate(tc *ciTestContext, imageField, imageTag string) {
	tc.t.Log("=== GetCustomerConfig after update (overridesOnly=false, expect override present) ===")
	resp, err := tc.client.CiExecutionConfigApi.GetCustomerConfig(tc.ctx, tc.infraType, false)
	printCiResult(tc.t, resp, err)
	checkCiField(tc.t, resp, imageField, imageTag)
}

// pass: HTTP 200, status SUCCESS — the API accepted the reset
// fail: any HTTP error
func testCiResetConfig(tc *ciTestContext, imageField string) {
	tc.t.Logf("=== ResetConfig (reset %s) ===", imageField)
	body := []ExecutionConfigUpdate{
		{Field: imageField},
	}
	resp, err := tc.client.CiExecutionConfigApi.ResetConfig(tc.ctx, tc.infraType, body)
	printCiResult(tc.t, resp, err)
}

// pass: HTTP 200, data nil or imageField empty — override was removed by reset
// fail: HTTP error, or imageField still holds the previously set override value
func testCiGetCustomerConfigOverridesOnlyAfterReset(tc *ciTestContext, imageField string) {
	tc.t.Log("=== GetCustomerConfig after reset (overridesOnly=true, expect no overrides) ===")
	resp, err := tc.client.CiExecutionConfigApi.GetCustomerConfig(tc.ctx, tc.infraType, true)
	printCiResult(tc.t, resp, err)
	require.True(tc.t, resp.Data == nil || resp.Data[imageField] == "",
		"%s still has override value %q after reset", imageField, resp.Data[imageField])
	tc.t.Logf("PASS: %s override is gone after reset", imageField)
}

func checkCiField(t *testing.T, resp ExecutionConfigResponse, imageField, imageTag string) {
	require.NotEmpty(t, resp.Data, "%s — response data is empty", imageField)
	require.Equal(t, imageTag, resp.Data[imageField],
		"%s mismatch: got %q, expected %q", imageField, resp.Data[imageField], imageTag)
	t.Logf("PASS: %s = %q", imageField, imageTag)
}

func printCiResult(t *testing.T, resp ExecutionConfigResponse, err error) {
	require.NoError(t, err)
	t.Logf("Status: %s", resp.Status)
	if len(resp.Data) > 0 {
		b, _ := json.MarshalIndent(resp.Data, "", "  ")
		t.Log(string(b))
	} else {
		t.Log("Data: (empty — all defaults in use)")
	}
}
