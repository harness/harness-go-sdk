/*
 * Chaos Manager API
 *
 * No description provided (generated by Swagger Codegen https://github.com/swagger-api/swagger-codegen)
 *
 * API version: 1.0
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */
package chaos

type ExecutionChaosData struct {
	ActionData *ExecutionActionData `json:"actionData,omitempty"`
	FaultData *ExecutionFaultData `json:"faultData,omitempty"`
	ProbeData *ExecutionProbeData `json:"probeData,omitempty"`
}
