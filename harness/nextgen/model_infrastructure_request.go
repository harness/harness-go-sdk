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

// This is the InfrastructureRequest entity defined in Harness
type InfrastructureRequest struct {
	// identifier of the infrastructure
	Identifier string `json:"identifier,omitempty"`
	// organisation identifier of the infrastructure
	OrgIdentifier string `json:"orgIdentifier,omitempty"`
	// project identifier of the infrastructure
	ProjectIdentifier string `json:"projectIdentifier,omitempty"`
	// environment reference of the infrastructure
	EnvironmentRef string `json:"environmentRef,omitempty"`
	// name of the infrastructure
	Name string `json:"name,omitempty"`
	// description of the infrastructure
	Description string `json:"description,omitempty"`
	// tags associated with the infrastructure
	Tags map[string]string `json:"tags,omitempty"`
	// type of the infrastructure
	Type_ string `json:"type"`
	// yaml spec of the infrastructure
	Yaml string `json:"yaml,omitempty"`
}