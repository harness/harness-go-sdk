/*
	go test \
		-v \
		-run TestIdpExecutionConfig \
		./harness/idp
*/

package idp

import (
	"context"
	"encoding/json"
	"os"
	"testing"

	"github.com/stretchr/testify/require"
)

func getIdpClientWithContext(t *testing.T) (*APIClient, context.Context) {
	t.Helper()
	accountID := os.Getenv("HARNESS_ACCOUNT_ID")
	apiKey := os.Getenv("HARNESS_PLATFORM_API_KEY")
	if accountID == "" || apiKey == "" {
		t.Skip("HARNESS_ACCOUNT_ID and HARNESS_PLATFORM_API_KEY must be set")
	}
	cfg := NewConfiguration()
	client := NewAPIClient(cfg)
	ctx := context.WithValue(context.Background(), ContextAPIKey, APIKey{Key: client.ApiKey})
	return client, ctx
}

type idpExecTestContext struct {
	t         *testing.T
	client    *APIClient
	ctx       context.Context
	infraType string
}

func TestIdpExecutionConfig(t *testing.T) {
	client, ctx := getIdpClientWithContext(t)

	tc := &idpExecTestContext{
		t:         t,
		client:    client,
		ctx:       ctx,
		infraType: "K8",
	}

	testUpdateImageType := "registerCatalog"
	testUpdateImageTag := "harness/registercatalog:updatedTmpNow"

	// reset the test field first to ensure clean state from any previous failed run
	testIdpResetConfig(tc, testUpdateImageType)
	// fetch the harness default images for the infra type
	testIdpGetDefaultConfig(tc)
	// verify no customer overrides exist before making any changes
	testIdpGetCustomerConfigOverridesOnly(tc)
	// override registerCatalog with a custom image tag
	testIdpUpdateConfig(tc, testUpdateImageType, testUpdateImageTag)
	// confirm the override appears when fetching overrides only
	testIdpGetCustomerConfigOverridesOnlyAfterUpdate(tc, testUpdateImageType, testUpdateImageTag)
	// confirm the override appears in the full config
	testIdpGetCustomerConfigAllAfterUpdate(tc, testUpdateImageType, testUpdateImageTag)
	// reset registerCatalog back to the harness default
	testIdpResetConfig(tc, testUpdateImageType)
	// confirm the override is gone after reset
	testIdpGetCustomerConfigOverridesOnlyAfterReset(tc, testUpdateImageType)
}

// pass: HTTP 200, status SUCCESS, data contains at least one image tag
// fail: any HTTP error, or data is empty (no images returned)
func testIdpGetDefaultConfig(tc *idpExecTestContext) {
	tc.t.Log("=== GetDefaultConfig ===")
	resp, err := tc.client.ExecutionConfigApi.GetDefaultConfig(tc.ctx, tc.infraType)
	printIdpExecResult(tc.t, resp, err)
	require.NotEmpty(tc.t, resp.Data, "expected default config data, got empty map")
	require.NotEmpty(tc.t, resp.Data["registerCatalog"], "expected registerCatalog to be set in default config")
}

// pass: HTTP 200, status SUCCESS, data nil or registerCatalog empty (no overrides set yet)
// fail: HTTP error, or registerCatalog already has an override (dirty state from a previous run)
func testIdpGetCustomerConfigOverridesOnly(tc *idpExecTestContext) {
	tc.t.Log("=== GetCustomerConfig before update (overridesOnly=true, expect no overrides) ===")
	resp, err := tc.client.ExecutionConfigApi.GetCustomerConfig(tc.ctx, tc.infraType, true)
	printIdpExecResult(tc.t, resp, err)
	require.True(tc.t, resp.Data == nil || resp.Data["registerCatalog"] == "",
		"expected no registerCatalog override before update, got %q", resp.Data["registerCatalog"])
}

// pass: HTTP 200, status SUCCESS — the API accepted the override
// fail: any HTTP error (e.g. 400 bad field name, 401 auth, 403 forbidden)
func testIdpUpdateConfig(tc *idpExecTestContext, imageField, imageTag string) {
	tc.t.Logf("=== UpdateConfig (set %s = %s) ===", imageField, imageTag)
	body := []ExecutionConfigUpdate{
		{Field: imageField, Value: imageTag},
	}
	resp, err := tc.client.ExecutionConfigApi.UpdateConfig(tc.ctx, tc.infraType, body)
	printIdpExecResult(tc.t, resp, err)
}

// pass: HTTP 200, data contains imageField set to imageTag (override is visible)
// fail: HTTP error, data empty, or imageField value does not match imageTag
func testIdpGetCustomerConfigOverridesOnlyAfterUpdate(tc *idpExecTestContext, imageField, imageTag string) {
	tc.t.Log("=== GetCustomerConfig after update (overridesOnly=true, expect override present) ===")
	resp, err := tc.client.ExecutionConfigApi.GetCustomerConfig(tc.ctx, tc.infraType, true)
	printIdpExecResult(tc.t, resp, err)
	checkIdpExecField(tc.t, resp, imageField, imageTag)
}

// pass: HTTP 200, full config returned with imageField set to imageTag
// fail: HTTP error, data empty, or imageField value does not match imageTag
func testIdpGetCustomerConfigAllAfterUpdate(tc *idpExecTestContext, imageField, imageTag string) {
	tc.t.Log("=== GetCustomerConfig after update (overridesOnly=false, expect override present) ===")
	resp, err := tc.client.ExecutionConfigApi.GetCustomerConfig(tc.ctx, tc.infraType, false)
	printIdpExecResult(tc.t, resp, err)
	checkIdpExecField(tc.t, resp, imageField, imageTag)
}

// pass: HTTP 200, status SUCCESS — the API accepted the reset
// fail: any HTTP error
func testIdpResetConfig(tc *idpExecTestContext, imageField string) {
	tc.t.Logf("=== ResetConfig (reset %s) ===", imageField)
	body := []ExecutionConfigUpdate{
		{Field: imageField},
	}
	resp, err := tc.client.ExecutionConfigApi.ResetConfig(tc.ctx, tc.infraType, body)
	printIdpExecResult(tc.t, resp, err)
}

// pass: HTTP 200, data nil or imageField empty — override was removed by reset
// fail: HTTP error, or imageField still holds the previously set override value
func testIdpGetCustomerConfigOverridesOnlyAfterReset(tc *idpExecTestContext, imageField string) {
	tc.t.Log("=== GetCustomerConfig after reset (overridesOnly=true, expect no overrides) ===")
	resp, err := tc.client.ExecutionConfigApi.GetCustomerConfig(tc.ctx, tc.infraType, true)
	printIdpExecResult(tc.t, resp, err)
	require.True(tc.t, resp.Data == nil || resp.Data[imageField] == "",
		"%s still has override value %q after reset", imageField, resp.Data[imageField])
	tc.t.Logf("PASS: %s override is gone after reset", imageField)
}

func checkIdpExecField(t *testing.T, resp ExecutionConfigResponse, imageField, imageTag string) {
	require.NotEmpty(t, resp.Data, "%s — response data is empty", imageField)
	require.Equal(t, imageTag, resp.Data[imageField],
		"%s mismatch: got %q, expected %q", imageField, resp.Data[imageField], imageTag)
	t.Logf("PASS: %s = %q", imageField, imageTag)
}

func printIdpExecResult(t *testing.T, resp ExecutionConfigResponse, err error) {
	require.NoError(t, err)
	t.Logf("Status: %s", resp.Status)
	if len(resp.Data) > 0 {
		b, _ := json.MarshalIndent(resp.Data, "", "  ")
		t.Log(string(b))
	} else {
		t.Log("Data: (empty — all defaults in use)")
	}
}
