/*
 * Infrastructure as Code Management
 *
 * Services for Harness IaCM Module.
 *
 * API version: 0.1.0
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */
package nextgen

type IacmDeleteDefaultPipelineRequestBody struct {
	// The default pipeline is associated with this operation
	Operation string `json:"operation"`
	// The default pipeline is associated with this provisioner
	Provisioner string `json:"provisioner"`
	// The default pipeline is associated with this workspace if specified
	Workspace string `json:"workspace,omitempty"`
}
