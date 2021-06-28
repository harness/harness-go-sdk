package api

import (
	"fmt"

	"github.com/harness-io/harness-go-sdk/harness/api/graphql"
)

type CloudProviderClient struct {
	APIClient *Client
}

func (c *Client) CloudProviders() *CloudProviderClient {
	return &CloudProviderClient{
		APIClient: c,
	}
}

func (c *CloudProviderClient) CreatePhysicalDataCenterCloudProvider(provider *graphql.PhysicalDataCenterCloudProvider) (*graphql.PhysicalDataCenterCloudProvider, error) {

	input := &graphql.CreateCloudProviderInput{
		CloudProviderType:               graphql.CloudProviderTypes.PhysicalDataCenter,
		PhysicalDataCenterCloudProvider: provider,
	}
	query := &GraphQLQuery{
		Query: fmt.Sprintf(`mutation($provider: CreateCloudProviderInput!) {
			createCloudProvider(input: $provider) {
				cloudProvider {
					%[1]s
				}
			}
		}`, getPhyisicalDataCenterCloudProviderFields()),
		Variables: map[string]interface{}{
			"provider": &input,
		},
	}

	res := &struct {
		CreateCloudProvider struct {
			CloudProvider graphql.PhysicalDataCenterCloudProvider
		}
	}{}

	err := c.APIClient.ExecuteGraphQLQuery(query, &res)

	if err != nil {
		return nil, err
	}

	return &res.CreateCloudProvider.CloudProvider, nil
}

func (c *CloudProviderClient) CreateKubernetesCloudProvider(provider *graphql.KubernetesCloudProvider) (*graphql.KubernetesCloudProvider, error) {

	input := &graphql.CreateCloudProviderInput{
		CloudProviderType: graphql.CloudProviderTypes.KubernetesCluster,
		K8sCloudProvider:  provider,
	}
	query := &GraphQLQuery{
		Query: fmt.Sprintf(`mutation($provider: CreateCloudProviderInput!) {
			createCloudProvider(input: $provider) {
				cloudProvider {
					%[1]s
				}
			}
		}`, getKubernetesCloudProviderFields()),
		Variables: map[string]interface{}{
			"provider": &input,
		},
	}

	res := &struct {
		CreateCloudProvider struct {
			CloudProvider graphql.KubernetesCloudProvider
		}
	}{}

	err := c.APIClient.ExecuteGraphQLQuery(query, &res)

	if err != nil {
		return nil, err
	}

	return &res.CreateCloudProvider.CloudProvider, nil
}

func (c *CloudProviderClient) CreateAwsCloudProvider(provider *graphql.AwsCloudProvider) (*graphql.AwsCloudProvider, error) {

	input := &graphql.CreateCloudProviderInput{
		CloudProviderType: graphql.CloudProviderTypes.Aws,
		AwsCloudProvider:  provider,
	}
	query := &GraphQLQuery{
		Query: fmt.Sprintf(`mutation($provider: CreateCloudProviderInput!) {
			createCloudProvider(input: $provider) {
				cloudProvider {
						%[1]s
				}
			}
		}`, getAwsCloudProviderFields()),
		Variables: map[string]interface{}{
			"provider": &input,
		},
	}

	res := &struct {
		CreateCloudProvider struct {
			CloudProvider graphql.AwsCloudProvider
		}
	}{}

	err := c.APIClient.ExecuteGraphQLQuery(query, &res)

	if err != nil {
		return nil, err
	}

	return &res.CreateCloudProvider.CloudProvider, nil
}

func (c *CloudProviderClient) DeleteCloudProvider(id string) error {

	query := &GraphQLQuery{
		Query: `mutation($input: DeleteCloudProviderInput!) {
			deleteCloudProvider(input: $input) {
				clientMutationId
			}
		}`,
		Variables: map[string]interface{}{
			"input": struct {
				CloudProviderId string `json:"cloudProviderId"`
			}{
				CloudProviderId: id,
			},
		},
	}

	res := &struct {
		DeleteCloudProvider struct {
			ClientMutationId string
		}
	}{}

	err := c.APIClient.ExecuteGraphQLQuery(query, &res)

	if err != nil {
		return err
	}

	return nil
}

func (c *CloudProviderClient) CreateSpotInstCloudProvider(provider *graphql.SpotInstCloudProvider) (*graphql.SpotInstCloudProvider, error) {

	input := &graphql.CreateCloudProviderInput{
		CloudProviderType:     graphql.CloudProviderTypes.SpotInst,
		SpotInstCloudProvider: provider,
	}
	query := &GraphQLQuery{
		Query: fmt.Sprintf(`mutation($provider: CreateCloudProviderInput!) {
			createCloudProvider(input: $provider) {
				cloudProvider {
					%[1]s
				}
			}
		}`, getSpotInstCloudProviderFields()),
		Variables: map[string]interface{}{
			"provider": &input,
		},
	}

	res := &struct {
		CreateCloudProvider struct {
			CloudProvider graphql.SpotInstCloudProvider
		}
	}{}

	err := c.APIClient.ExecuteGraphQLQuery(query, &res)

	if err != nil {
		return nil, err
	}

	return &res.CreateCloudProvider.CloudProvider, nil
}

func (c *CloudProviderClient) CreatePcfCloudProvider(provider *graphql.PcfCloudProvider) (*graphql.PcfCloudProvider, error) {

	input := &graphql.CreateCloudProviderInput{
		CloudProviderType: graphql.CloudProviderTypes.Pcf,
		PcfCloudProvider:  provider,
	}
	query := &GraphQLQuery{
		Query: fmt.Sprintf(`mutation($provider: CreateCloudProviderInput!) {
			createCloudProvider(input: $provider) {
				cloudProvider {
					%[1]s
				}
			}
		}`, getPcfCloudProviderFields()),
		Variables: map[string]interface{}{
			"provider": &input,
		},
	}

	res := &struct {
		CreateCloudProvider struct {
			CloudProvider graphql.PcfCloudProvider
		}
	}{}

	err := c.APIClient.ExecuteGraphQLQuery(query, &res)

	if err != nil {
		return nil, err
	}

	return &res.CreateCloudProvider.CloudProvider, nil
}

func (c *CloudProviderClient) CreateGcpCloudProvider(provider *graphql.GcpCloudProvider) (*graphql.GcpCloudProvider, error) {

	input := &graphql.CreateCloudProviderInput{
		CloudProviderType: graphql.CloudProviderTypes.Gcp,
		GCPCloudProvider:  provider,
	}
	query := &GraphQLQuery{
		Query: fmt.Sprintf(`mutation($provider: CreateCloudProviderInput!) {
			createCloudProvider(input: $provider) {
				cloudProvider {
					%[1]s
				}
			}
		}`, getGcpCloudProviderFields()),
		Variables: map[string]interface{}{
			"provider": &input,
		},
	}

	res := &struct {
		CreateCloudProvider struct {
			CloudProvider graphql.GcpCloudProvider
		}
	}{}

	err := c.APIClient.ExecuteGraphQLQuery(query, &res)

	if err != nil {
		return nil, err
	}

	return &res.CreateCloudProvider.CloudProvider, nil
}

func (c *CloudProviderClient) CreateAzureCloudProvider(provider *graphql.AzureCloudProvider) (*graphql.AzureCloudProvider, error) {

	input := &graphql.CreateCloudProviderInput{
		CloudProviderType:  graphql.CloudProviderTypes.Azure,
		AzureCloudProvider: provider,
	}
	query := &GraphQLQuery{
		Query: fmt.Sprintf(`mutation($provider: CreateCloudProviderInput!) {
			createCloudProvider(input: $provider) {
				cloudProvider {
					%[1]s
				}
			}
		}`, getAzureCloudProviderFields()),
		Variables: map[string]interface{}{
			"provider": &input,
		},
	}

	res := &struct {
		CreateCloudProvider struct {
			CloudProvider graphql.AzureCloudProvider
		}
	}{}

	err := c.APIClient.ExecuteGraphQLQuery(query, &res)

	if err != nil {
		return nil, err
	}

	return &res.CreateCloudProvider.CloudProvider, nil
}

const (
	commonCloudProviderFields = `
		id
		name
		description
		createdAt
		createdBy {
			id
			name
		}
		type
		isContinuousEfficiencyEnabled
	`

	ceHealthStatusFields = `
		ceHealthStatus {
			clusterHealthStatusList {
				clusterId
				clusterName
				errors
				lastEventTimestamp
				messages
			}
			isCEConnector
			isHealthy
			messages
		}
	`
)

func getPhyisicalDataCenterCloudProviderFields() string {
	return fmt.Sprintf(`
	... on PhysicalDataCenterCloudProvider {
		%[1]s
	}
	`, commonCloudProviderFields)
}

func getKubernetesCloudProviderFields() string {
	return fmt.Sprintf(`
		%[1]s
		... on KubernetesCloudProvider {
			%[2]s
			skipK8sEventCollection
		}
`, commonCloudProviderFields, ceHealthStatusFields)
}

func getAwsCloudProviderFields() string {
	return fmt.Sprintf(`
		%[1]s
		... on AwsCloudProvider {
			%[2]s
		}
	`, commonCloudProviderFields, ceHealthStatusFields)
}

func getSpotInstCloudProviderFields() string {
	return fmt.Sprintf(`
	... on SpotInstCloudProvider {
		%[1]s
	}
	`, commonCloudProviderFields)
}

func getPcfCloudProviderFields() string {
	return fmt.Sprintf(`
	... on PcfCloudProvider {
		%[1]s
	}
	`, commonCloudProviderFields)
}

func getGcpCloudProviderFields() string {
	return fmt.Sprintf(`
		%[1]s
		... on GcpCloudProvider {
			delegateSelectors
		}
	`, commonCloudProviderFields)
}

func getAzureCloudProviderFields() string {
	return fmt.Sprintf(`
		%[1]s
	`, commonCloudProviderFields)
}
