/*
 * Chaos Manager API
 *
 * No description provided (generated by Swagger Codegen https://github.com/swagger-api/swagger-codegen)
 *
 * API version: 1.0
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */
package chaos

type ChaosExperimentRunConditionDetails struct {
	ConditionID string `json:"conditionID,omitempty"`
	ConditionName string `json:"conditionName,omitempty"`
	Message string `json:"message,omitempty"`
	Phase string `json:"phase,omitempty"`
}
