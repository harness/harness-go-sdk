/*
 * CD NextGen API Reference
 *
 * This is the Open Api Spec 3 for the NextGen Manager. This is under active development. Beware of the breaking change with respect to the generated code stub  # Authentication  <!-- ReDoc-Inject: <security-definitions> -->
 *
 * API version: 3.0
 * Contact: contact@harness.io
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */
package nextgen

// This has details of encrypted text secret.
type SecretTextSpec struct {
	ErrorMessageForInvalidYaml string `json:"errorMessageForInvalidYaml,omitempty"`
	Type_                      string `json:"type"`
	// Identifier of the Secret Manager used to manage the secret.
	SecretManagerIdentifier string `json:"secretManagerIdentifier"`
	// This has details to specify if the secret value is inline or referenced.
	ValueType string `json:"valueType"`
	// Value of the Secret
	Value string `json:"value,omitempty"`
}