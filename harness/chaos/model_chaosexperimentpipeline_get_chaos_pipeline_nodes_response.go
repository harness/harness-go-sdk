/*
 * Chaos Manager API
 *
 * No description provided (generated by Swagger Codegen https://github.com/swagger-api/swagger-codegen)
 *
 * API version: 1.0
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */
package chaos

type ChaosexperimentpipelineGetChaosPipelineNodesResponse struct {
	ChaosExecutionNode *ChaosExecutionNodesChaosExecutionNode `json:"chaosExecutionNode,omitempty"`
	RefEntities *ChaosexperimentpipelineReferenceEntities `json:"refEntities,omitempty"`
}
