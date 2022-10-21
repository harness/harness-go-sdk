/*
 * Harness NextGen Software Delivery Platform API Reference
 *
 * This is the Open Api Spec 3 for the NextGen Manager. This is under active development. Beware of the breaking change with respect to the generated code stub
 *
 * API version: 3.0
 * Contact: contact@harness.io
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */
package nextgen

type V1AgentMetadata struct {
	Namespace                string                      `json:"namespace,omitempty"`
	HighAvailability         bool                        `json:"highAvailability,omitempty"`
	DeployedApplicationCount int32                       `json:"deployedApplicationCount,omitempty"`
	ExistingInstallation     bool                        `json:"existingInstallation,omitempty"`
	MappedProjects           *Servicev1AppProjectMapping `json:"mappedProjects,omitempty"`
}