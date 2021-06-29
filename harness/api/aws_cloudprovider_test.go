package api

import (
	"fmt"
	"os"
	"testing"

	"github.com/harness-io/harness-go-sdk/harness/api/graphql"
	"github.com/harness-io/harness-go-sdk/harness/utils"
	"github.com/stretchr/testify/require"
)

func TestCreateAwsCloudProvider(t *testing.T) {
	c := getClient()
	expectedName := fmt.Sprintf("%s_%s", t.Name(), utils.RandStringBytes(4))

	cp, err := createAwsCloudProvider(expectedName)
	require.NoError(t, err)
	require.NotNil(t, cp)
	require.Equal(t, expectedName, cp.Name)

	err = c.CloudProviders().DeleteCloudProvider(cp.Id)
	require.NoError(t, err)
}

func TestGetAwsCloudProviderByName(t *testing.T) {
	c := getClient()
	expectedName := fmt.Sprintf("%s_%s", t.Name(), utils.RandStringBytes(4))

	cp, err := createAwsCloudProvider(expectedName)
	require.NoError(t, err)

	foundCP, err := c.CloudProviders().GetAwsCloudProviderByName(expectedName)
	require.NoError(t, err)
	require.NotNil(t, foundCP)
	require.Equal(t, expectedName, foundCP.Name)

	err = c.CloudProviders().DeleteCloudProvider(cp.Id)
	require.NoError(t, err)
}

func createAwsCloudProvider(name string) (*graphql.AwsCloudProvider, error) {

	c := getClient()
	expectedName := name

	secret, err := createEncryptedTextSecret(expectedName, os.Getenv("HARNESS_TEST_AWS_SECRET_ACCESS_KEY"))
	if err != nil {
		return nil, err
	}

	input := &graphql.AwsCloudProvider{}
	input.Name = expectedName
	input.CredentialsType = graphql.AwsCredentialsTypes.Manual
	input.ManualCredentials = &graphql.AwsManualCredentials{
		AccessKey:         os.Getenv("HARNESS_TEST_AWS_ACCESS_KEY_ID"),
		SecretKeySecretId: secret.Id,
	}

	return c.CloudProviders().CreateAwsCloudProvider(input)
}
