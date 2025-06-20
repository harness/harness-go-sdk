/*
 * Chaos Manager API
 *
 * No description provided (generated by Swagger Codegen https://github.com/swagger-api/swagger-codegen)
 *
 * API version: 1.0
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */
package chaos

type ExecutionChaosStepType string

// List of execution.ChaosStepType
const (
	FAULT_ExecutionChaosStepType ExecutionChaosStepType = "FAULT"
	PROBE_ExecutionChaosStepType ExecutionChaosStepType = "PROBE"
	ACTION_ExecutionChaosStepType ExecutionChaosStepType = "ACTION"
	EXPERIMENT_ExecutionChaosStepType ExecutionChaosStepType = "EXPERIMENT"
)
