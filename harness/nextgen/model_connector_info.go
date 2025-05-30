/*
 * Harness NextGen Software Delivery Platform API Reference
 *
 * This is the Open Api Spec 3 for the NextGen Manager. This is under active development. Beware of the breaking change with respect to the generated code stub  # Authentication  <!-- ReDoc-Inject: <security-definitions> -->
 *
 * API version: 3.0
 * Contact: contact@harness.io
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */
package nextgen

// This has the Connector details defined in Harness

import (
	"encoding/json"
)

type ConnectorInfo struct {
	// Name of the Connector.
	Name string `json:"name"`
	// Identifier of the Connector.
	Identifier string `json:"identifier"`
	// Description of the entity
	Description string `json:"description,omitempty"`
	// Organization Identifier for the Entity.
	OrgIdentifier string `json:"orgIdentifier,omitempty"`
	// Project Identifier for the Entity.
	ProjectIdentifier string `json:"projectIdentifier,omitempty"`
	// Tags
	Tags map[string]string `json:"tags,omitempty"`
	// Type of the Connector.
	Type_               ConnectorType                   `json:"type"`
	AppDynamics         *AppDynamicsConnectorDto        `json:"-"`
	Artifactory         *ArtifactoryConnector           `json:"-"`
	Aws                 *AwsConnector                   `json:"-"`
	AwsCC               *CeAwsConnector                 `json:"-"`
	AwsKms              *AwsKmsConnector                `json:"-"`
	AwsSecretManager    *AwsSecretManager               `json:"-"`
	AzureCloudCost      *CeAzureConnector               `json:"-"`
	Azure               *AzureConnector                 `json:"-"`
	AzureArtifacts      *AzureArtifactsConnector        `json:"-"`
	AzureKeyVault       *AzureKeyVaultConnector         `json:"-"`
	BitBucket           *BitbucketConnector             `json:"-"`
	CustomSecretManager *CustomSecretManager            `json:"_"`
	Datadog             *DatadogConnectorDto            `json:"-"`
	DockerRegistry      *DockerConnector                `json:"-"`
	JDBC                *JdbcConnector                  `json:"-"`
	Dynatrace           *DynatraceConnectorDto          `json:"-"`
	Gcp                 *GcpConnector                   `json:"-"`
	GcpCloudCost        *GcpCloudCostConnectorDto       `json:"-"`
	GcpKms              *GcpKmsConnector                `json:"-"`
	Git                 *GitConfig                      `json:"-"`
	Github              *GithubConnector                `json:"-"`
	Gitlab              *GitlabConnector                `json:"-"`
	HttpHelm            *HttpHelmConnector              `json:"-"`
	OciHelm             *OciHelmConnector               `json:"-"`
	Jira                *JiraConnector                  `json:"-"`
	Jenkins             *JenkinsConnector               `json:"-"`
	K8sCluster          *KubernetesClusterConfig        `json:"-"`
	K8sClusterCloudCost *CeKubernetesClusterConfigDto   `json:"-"`
	NewRelic            *NewRelicConnectorDto           `json:"-"`
	Nexus               *NexusConnector                 `json:"-"`
	PagerDuty           *PagerDutyConnectorDto          `json:"-"`
	Prometheus          *PrometheusConnectorDto         `json:"-"`
	Splunk              *SplunkConnector                `json:"-"`
	SumoLogic           *SumoLogicConnectorDto          `json:"-"`
	Spot                *SpotConnector                  `json:"-"`
	Tas                 *TasConnector                   `json:"-"`
	TerraformCloud      *TerraformCloudConnector        `json:"-"`
	Vault               *VaultConnector                 `json:"-"`
	GcpSecretManager    *GcpSecretManager               `json:"-"`
	ServiceNow          *ServiceNowConnector            `json:"-"`
	Spec                json.RawMessage                 `json:"spec"`
	ElasticSearch       *ElkConnectorDto                `json:"-"`
	Rancher             *RancherConnector               `json:"-"`
	CustomHealth        *CustomHealthConnectorDto       `json:"-"`
	Pdc                 *PhysicalDataCenterConnectorDto `json:"-"`
	AzureRepo           *AzureRepoConfig                `json:"-"`
}
