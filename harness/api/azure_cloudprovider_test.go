package api

import (
	"fmt"
	"os"
	"testing"

	"github.com/harness-io/harness-go-sdk/harness/api/graphql"
	"github.com/harness-io/harness-go-sdk/harness/utils"
	"github.com/stretchr/testify/require"
)

func TestGetAzureCloudProviderByName(t *testing.T) {
	c := getClient()
	expectedName := fmt.Sprintf("%s_%s", t.Name(), utils.RandStringBytes(4))

	cp, err := createAzureCloudProvider(expectedName)
	require.NoError(t, err)

	foundCP, err := c.CloudProviders().GetAzureCloudProviderByName(expectedName)
	require.NoError(t, err)
	require.NotNil(t, foundCP)
	require.Equal(t, expectedName, foundCP.Name)

	err = c.CloudProviders().DeleteCloudProvider(cp.Id)
	require.NoError(t, err)
}

func TestCreateAzureCloudProvider(t *testing.T) {
	c := getClient()

	expectedName := fmt.Sprintf("%s_%s", t.Name(), utils.RandStringBytes(4))

	cp, err := createAzureCloudProvider(expectedName)
	require.NoError(t, err)
	require.NotNil(t, cp)
	require.Equal(t, expectedName, cp.Name)

	err = c.CloudProviders().DeleteCloudProvider(cp.Id)
	require.NoError(t, err)

	secret, err := c.Secrets().GetEncryptedTextByName(expectedName)
	require.NoError(t, err)
	c.Secrets().DeleteSecret(secret.Id, secret.SecretType)
}

func createAzureCloudProvider(name string) (*graphql.AzureCloudProvider, error) {

	c := getClient()
	expectedName := name

	secret, err := createEncryptedTextSecret(expectedName, os.Getenv("HARNESS_TEST_AZURE_CLIENT_SECRET"))
	if err != nil {
		return nil, err
	}

	input := &graphql.AzureCloudProvider{}
	input.Name = expectedName
	input.ClientId = os.Getenv("HARNESS_TEST_AZURE_CLIENT_ID")
	input.KeySecretId = secret.Id
	input.TenantId = os.Getenv("HARNESS_TEST_AZURE_TENANT_ID")

	return c.CloudProviders().CreateAzureCloudProvider(input)
}
