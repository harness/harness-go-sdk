/*
 * Chaos Manager API
 *
 * No description provided (generated by Swagger Codegen https://github.com/swagger-api/swagger-codegen)
 *
 * API version: 1.0
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */
package chaos

type ChaosExecutionNodesChaosData struct {
	ActionData *ChaosExecutionNodesActionData `json:"actionData,omitempty"`
	FaultData *ChaosExecutionNodesFaultData `json:"faultData,omitempty"`
	ProbeData *ChaosExecutionNodesProbeData `json:"probeData,omitempty"`
}
