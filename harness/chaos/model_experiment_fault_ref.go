/*
 * Chaos Manager API
 *
 * No description provided (generated by Swagger Codegen https://github.com/swagger-api/swagger-codegen)
 *
 * API version: 1.0
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */
package chaos

type ExperimentFaultRef struct {
	// HubRef of the fault reference
	HubRef string `json:"hubRef,omitempty"`
	// Identity of the fault reference
	Identity string `json:"identity,omitempty"`
	// InfraID contains the infrastructure id
	InfraId string `json:"infraId,omitempty"`
	// Name of the fault reference
	Name string `json:"name,omitempty"`
	// Revision of the fault reference
	Revision int32 `json:"revision,omitempty"`
	// Variables to store the variables
	Values []TemplateVariableMinimum `json:"values,omitempty"`
}
