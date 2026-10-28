package helpers

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestEnvVarsAreSet(t *testing.T) {
	if testing.Short() || TestEnvVars.AwsAccessKeyId.Get() == "" {
		t.Skip("skipping env var assertions in unit test mode")
	}
	assert.NotEmpty(t, TestEnvVars.AwsAccessKeyId.Get())
	assert.NotEmpty(t, TestEnvVars.AwsSecretAccessKey.Get())
	assert.NotEmpty(t, TestEnvVars.AzureClientId.Get())
	assert.NotEmpty(t, TestEnvVars.AzureClientSecret.Get())
	assert.NotEmpty(t, TestEnvVars.AzureTenantId.Get())
	assert.NotEmpty(t, TestEnvVars.SpotInstAccountId.Get())
	assert.NotEmpty(t, TestEnvVars.SpotInstToken.Get())
}
