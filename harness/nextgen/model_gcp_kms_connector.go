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

// This contains GCP KMS SecretManager configuration.
type GcpKmsConnector struct {
	// ID of the project on GCP.
	ProjectId string `json:"projectId"`
	// Region for GCP KMS
	Region string `json:"region"`
	// Name of the Key Ring where Google Cloud Symmetric Key is created.
	KeyRing string `json:"keyRing"`
	// Name of the Google Cloud Symmetric Key.
	KeyName     string `json:"keyName"`
	Credentials string `json:"credentials"`
	// List of Delegate Selectors that belong to the same Delegate and are used to connect to the Secret Manager.
	DelegateSelectors []string `json:"delegateSelectors,omitempty"`
	Default_          bool     `json:"default,omitempty"`
}
