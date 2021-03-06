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

// This contains Physical Data Center connector details
type PhysicalDataCenterConnectorDto struct {
	Hosts             []HostDto `json:"hosts,omitempty"`
	SshKeyRef         string    `json:"sshKeyRef"`
	DelegateSelectors []string  `json:"delegateSelectors,omitempty"`
}
