package api

import (
	"fmt"
	"os"
	"testing"

	"github.com/harness-io/harness-go-sdk/harness/api/graphql"
	"github.com/harness-io/harness-go-sdk/harness/utils"
	"github.com/stretchr/testify/require"
)

func TestCreatePhysicalDataCenterCloudProvider(t *testing.T) {
	c := getClient()

	expectedName := fmt.Sprintf("%s_%s", t.Name(), utils.RandStringBytes(4))

	input := &graphql.PhysicalDataCenterCloudProvider{}
	input.Name = expectedName

	cp, err := c.CloudProviders().CreatePhysicalDataCenterCloudProvider(input)
	require.NoError(t, err)
	require.NotNil(t, cp)
	require.Equal(t, expectedName, cp.Name)

	err = c.CloudProviders().DeleteCloudProvider(cp.Id)
	require.NoError(t, err)
}

func TestCreateKubernetesCloudProvider(t *testing.T) {
	c := getClient()
	expectedName := fmt.Sprintf("%s_%s", t.Name(), utils.RandStringBytes(4))

	input := &graphql.KubernetesCloudProvider{}
	input.Name = expectedName
	input.ClusterDetailsType = graphql.ClusterDetailsTypes.InheritClusterDetails
	input.InheritClusterDetails = &graphql.InheritClusterDetails{
		DelegateSelectors: []string{"Primary"},
	}
	input.SkipValidation = true

	cp, err := c.CloudProviders().CreateKubernetesCloudProvider(input)
	require.NoError(t, err)
	require.NotNil(t, cp)
	require.Equal(t, expectedName, cp.Name)

	err = c.CloudProviders().DeleteCloudProvider(cp.Id)
	require.NoError(t, err)

}

func TestCreateAwsCloudProvider(t *testing.T) {
	c := getClient()
	expectedName := fmt.Sprintf("%s_%s", t.Name(), utils.RandStringBytes(4))

	input := &graphql.AwsCloudProvider{}
	input.Name = expectedName
	input.CredentialsType = graphql.AwsCredentialsTypes.Ec2Iam
	input.Ec2IamCredentials = &graphql.Ec2IamCredentials{
		DelegateSelector: "test-account",
	}

	cp, err := c.CloudProviders().CreateAwsCloudProvider(input)
	require.NoError(t, err)
	require.NotNil(t, cp)
	require.Equal(t, expectedName, cp.Name)

	err = c.CloudProviders().DeleteCloudProvider(cp.Id)
	require.NoError(t, err)
}

func TestCreateSpotInstCloudProvider(t *testing.T) {
	c := getClient()
	expectedName := fmt.Sprintf("%s_%s", t.Name(), utils.RandStringBytes(4))

	secret, err := createEncryptedTextSecret(expectedName, os.Getenv("HARNESS_TEST_SPOT_TOKEN"))
	require.NoError(t, err)
	require.NotNil(t, secret)

	input := &graphql.SpotInstCloudProvider{}
	input.Name = expectedName
	input.AccountId = os.Getenv("HARNESS_TEST_SPOT_ACCT_ID")
	input.TokenSecretId = secret.Id

	cp, err := c.CloudProviders().CreateSpotInstCloudProvider(input)
	require.NoError(t, err)
	require.NotNil(t, cp)
	require.Equal(t, expectedName, cp.Name)

	err = c.CloudProviders().DeleteCloudProvider(cp.Id)
	require.NoError(t, err)

	err = c.Secrets().DeleteSecret(secret.Id, secret.SecretType)
	require.NoError(t, err)
}

func TestPcfCenterCloudProvider(t *testing.T) {
	c := getClient()

	expectedName := fmt.Sprintf("%s_%s", t.Name(), utils.RandStringBytes(4))

	input := &graphql.PcfCloudProvider{}
	input.Name = expectedName
	input.EndpointUrl = "https://example.com"
	input.PasswordSecretId = "abc123"
	input.SkipValidation = true
	input.UserName = "foo123"

	cp, err := c.CloudProviders().CreatePcfCloudProvider(input)
	require.NoError(t, err)
	require.NotNil(t, cp)
	require.Equal(t, expectedName, cp.Name)

	err = c.CloudProviders().DeleteCloudProvider(cp.Id)
	require.NoError(t, err)
}

func TestGcpCenterCloudProvider(t *testing.T) {
	c := getClient()

	expectedName := fmt.Sprintf("%s_%s", t.Name(), utils.RandStringBytes(4))

	input := &graphql.GcpCloudProvider{}
	input.Name = expectedName
	input.DelegateSelectors = []string{"Primary"}
	input.UseDelegate = true
	input.SkipValidation = true

	cp, err := c.CloudProviders().CreateGcpCloudProvider(input)
	require.NoError(t, err)
	require.NotNil(t, cp)
	require.Equal(t, expectedName, cp.Name)

	err = c.CloudProviders().DeleteCloudProvider(cp.Id)
	require.NoError(t, err)
}

func TestAzureCenterCloudProvider(t *testing.T) {
	c := getClient()

	expectedName := fmt.Sprintf("%s_%s", t.Name(), utils.RandStringBytes(4))

	secret, err := createEncryptedTextSecret(expectedName, os.Getenv("HARNESS_TEST_AZURE_CLIENT_SECRET"))
	require.NoError(t, err)
	require.NotNil(t, secret)

	input := &graphql.AzureCloudProvider{}
	input.Name = expectedName
	input.ClientId = os.Getenv("HARNESS_TEST_AZURE_CLIENT_ID")
	input.KeySecretId = secret.Id
	input.TenantId = os.Getenv("HARNESS_TEST_AZURE_TENANT_ID")

	cp, err := c.CloudProviders().CreateAzureCloudProvider(input)
	require.NoError(t, err)
	require.NotNil(t, cp)
	require.Equal(t, expectedName, cp.Name)

	err = c.CloudProviders().DeleteCloudProvider(cp.Id)
	require.NoError(t, err)

	err = c.Secrets().DeleteSecret(secret.Id, secret.SecretType)
	require.NoError(t, err)
}
