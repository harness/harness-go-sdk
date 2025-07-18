/*
 * Chaos Manager API
 *
 * No description provided (generated by Swagger Codegen https://github.com/swagger-api/swagger-codegen)
 *
 * API version: 1.0
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */
package chaos

type ModelWorkflowType string

// List of model.WorkflowType
const (
	ALL_ModelWorkflowType ModelWorkflowType = "All"
	GAMEDAY_WORKFLOW_ModelWorkflowType ModelWorkflowType = "GamedayWorkflow"
	WORKFLOW_ModelWorkflowType ModelWorkflowType = "Workflow"
	CRON_WORKFLOW_ModelWorkflowType ModelWorkflowType = "CronWorkflow"
	CHAOS_ENGINE_ModelWorkflowType ModelWorkflowType = "ChaosEngine"
	CHAOS_SCHEDULE_ModelWorkflowType ModelWorkflowType = "ChaosSchedule"
)
