/*
 * Chaos Manager API
 *
 * No description provided (generated by Swagger Codegen https://github.com/swagger-api/swagger-codegen)
 *
 * API version: 1.0
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */
package chaos

type PipelinesChaosPipelineInput struct {
	Description string `json:"description,omitempty"`
	ExperimentSpec []PipelinesExperimentSpec `json:"experimentSpec,omitempty"`
	Identifier string `json:"identifier,omitempty"`
	Name string `json:"name,omitempty"`
	Tags map[string]string `json:"tags,omitempty"`
}
