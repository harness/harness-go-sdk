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

// Returns secret reference access key and secret key of AWS Secret Manager.
type AwsSmCredentialSpecOidcConfig struct {
	// List of Delegate Selectors that belong to the same Delegate and are used to connect to the Secret Manager.
	DelegateSelectors []string `json:"delegateSelectors"`
	IamRoleArn        string   `json:"iamRoleArn,omitempty"`
}
