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

// This contains the Vault Connector configuration.
type VaultConnector struct {
	ConnectorType string `json:"connectorType"`
	// This is the authentication token for Vault.
	AuthToken string `json:"authToken,omitempty"`
	// This is the location of the Vault directory where Secret will be stored.
	BasePath string `json:"basePath,omitempty"`
	// URL of the HashiCorp Vault.
	VaultUrl   string `json:"vaultUrl"`
	IsReadOnly bool   `json:"isReadOnly,omitempty"`
	// This is the time interval for token renewal.
	RenewalIntervalMinutes int64 `json:"renewalIntervalMinutes"`
	// Manually entered Secret Engine.
	SecretEngineManuallyConfigured bool `json:"secretEngineManuallyConfigured,omitempty"`
	// Name of the Secret Engine.
	SecretEngineName string `json:"secretEngineName,omitempty"`
	// ID of App Role.
	AppRoleId string `json:"appRoleId,omitempty"`
	// ID of the Secret.
	SecretId  string `json:"secretId,omitempty"`
	IsDefault bool   `json:"isDefault,omitempty"`
	// Version of Secret Engine.
	SecretEngineVersion int32 `json:"secretEngineVersion,omitempty"`
	// List of Delegate Selectors that belong to the same Delegate and are used to connect to the Secret Manager.
	DelegateSelectors []string `json:"delegateSelectors,omitempty"`
	// This is the Vault namespace where Secret will be created.
	Namespace string `json:"namespace,omitempty"`
	// This is the location at which auth token is to be read from.
	SinkPath string `json:"sinkPath,omitempty"`
	// Boolean value to indicate if Vault Agent is used for authentication.
	UseVaultAgent bool `json:"useVaultAgent,omitempty"`
	// Boolean value to indicate if Aws Iam is used for authentication.
	UseAwsIam bool `json:"useAwsIam,omitempty"`
	// This is the Aws region where aws iam auth will happen.
	AwsRegion string `json:"awsRegion,omitempty"`
	// This is the Vault role defined to bind to aws iam account/role being accessed.
	VaultAwsIamRole string `json:"vaultAwsIamRole,omitempty"`
	// This is the Aws Iam Header Server ID that has been configured for this Aws Iam instance.
	XvaultAwsIamServerId string `json:"xvaultAwsIamServerId,omitempty"`
	// Boolean value to indicate if K8s Auth is used for authentication.
	UseK8sAuth bool `json:"useK8sAuth,omitempty"`
	// This is the role where K8s auth will happen.
	VaultK8sAuthRole string `json:"vaultK8sAuthRole,omitempty"`
	// This is the SA token path where the token is mounted in the K8s Pod.
	ServiceAccountTokenPath string `json:"serviceAccountTokenPath,omitempty"`
	// This is the path where kubernetes auth is enabled in Vault.
	K8sAuthEndpoint string `json:"k8sAuthEndpoint,omitempty"`
	// Boolean value to indicate if appRole token renewal is enabled or not.
	RenewAppRoleToken bool   `json:"renewAppRoleToken,omitempty"`
	AccessType        string `json:"accessType,omitempty"`
	Default_          bool   `json:"default,omitempty"`
	ReadOnly          bool   `json:"readOnly,omitempty"`
	// Boolean value to indicate if JWT Auth is used for authentication.
	UseJwtAuth bool `json:"useJwtAuth,omitempty"`
	// This is the role name which is created to perform JWT auth method.
	JwtAuthRole string `json:"jwtAuthRole,omitempty"`
	// This specifies mount path where JWT auth method is enabled.
	JwtAuthPath       string `json:"jwtAuthPath,omitempty"`
	ExecuteOnDelegate bool   `json:"executeOnDelegate"`
}
