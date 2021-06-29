package api

import (
	"fmt"
	"os"
	"testing"

	"github.com/harness-io/harness-go-sdk/harness/api/graphql"
	"github.com/harness-io/harness-go-sdk/harness/utils"
	"github.com/stretchr/testify/require"
)

func TestGetSpotInstCloudProviderByName(t *testing.T) {
	c := getClient()
	expectedName := fmt.Sprintf("%s_%s", t.Name(), utils.RandStringBytes(4))

	cp, err := createSpotInstCloudProvider(expectedName)
	require.NoError(t, err)

	foundCP, err := c.CloudProviders().GetSpotInstCloudProviderByName(expectedName)
	require.NoError(t, err)
	require.NotNil(t, foundCP)
	require.Equal(t, expectedName, foundCP.Name)

	err = c.CloudProviders().DeleteCloudProvider(cp.Id)
	require.NoError(t, err)
}

func TestCreateSpotInstCloudProvider(t *testing.T) {
	c := getClient()
	expectedName := fmt.Sprintf("%s_%s", t.Name(), utils.RandStringBytes(4))

	cp, err := createSpotInstCloudProvider(expectedName)
	require.NoError(t, err)
	require.NotNil(t, cp)
	require.Equal(t, expectedName, cp.Name)

	err = c.CloudProviders().DeleteCloudProvider(cp.Id)
	require.NoError(t, err)

	err = c.Secrets().DeleteSecret(expectedName, graphql.SecretTypes.EncryptedText)
	require.NoError(t, err)
}

func createSpotInstCloudProvider(name string) (*graphql.SpotInstCloudProvider, error) {
	c := getClient()

	secret, err := createEncryptedTextSecret(name, os.Getenv("HARNESS_TEST_SPOT_TOKEN"))
	if err != nil {
		return nil, err
	}

	input := &graphql.SpotInstCloudProvider{}
	input.Name = name
	input.AccountId = os.Getenv("HARNESS_TEST_SPOT_ACCT_ID")
	input.TokenSecretId = secret.Id

	cp, err := c.CloudProviders().CreateSpotInstCloudProvider(input)
	if err != nil {
		return nil, err
	}

	return cp, nil
}
