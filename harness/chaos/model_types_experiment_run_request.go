/*
 * Chaos Manager API
 *
 * No description provided (generated by Swagger Codegen https://github.com/swagger-api/swagger-codegen)
 *
 * API version: 1.0
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */
package chaos

type TypesExperimentRunRequest struct {
	InputsetIdentity string `json:"inputsetIdentity,omitempty"`
	RuntimeInputs *TemplateChaosExperimentInputsetSpec `json:"runtimeInputs,omitempty"`
}
