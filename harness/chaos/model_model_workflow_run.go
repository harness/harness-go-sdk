/*
 * Chaos Manager API
 *
 * No description provided (generated by Swagger Codegen https://github.com/swagger-api/swagger-codegen)
 *
 * API version: 1.0
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */
package chaos

type ModelWorkflowRun struct {
	// Timestamp at which workflow run was created
	CreatedAt string `json:"createdAt,omitempty"`
	// User who has created the experiment run
	CreatedBy *AllOfmodelWorkflowRunCreatedBy `json:"createdBy,omitempty"`
	// Cron syntax of the workflow schedule
	CronSyntax string `json:"cronSyntax,omitempty"`
	// Error Response is the reason why experiment failed to run
	ErrorResponse string `json:"errorResponse,omitempty"`
	// Stores all the workflow run details related to the nodes of DAG graph and chaos results of the experiments
	ExecutionData string `json:"executionData,omitempty"`
	// experimentType is the type of experiment run
	ExperimentType string `json:"experimentType,omitempty"`
	// Number of experiments awaited
	ExperimentsAwaited int32 `json:"experimentsAwaited,omitempty"`
	// Number of experiments failed
	ExperimentsFailed int32 `json:"experimentsFailed,omitempty"`
	// Number of experiments which are not available
	ExperimentsNa int32 `json:"experimentsNa,omitempty"`
	// Number of experiments passed
	ExperimentsPassed int32 `json:"experimentsPassed,omitempty"`
	// Number of experiments stopped
	ExperimentsStopped int32 `json:"experimentsStopped,omitempty"`
	// Harness identifiers
	Identifiers *AllOfmodelWorkflowRunIdentifiers `json:"identifiers,omitempty"`
	// Target infra in which the workflow will run
	Infra *AllOfmodelWorkflowRunInfra `json:"infra,omitempty"`
	// If cron is enabled or disabled
	IsCronEnabled bool `json:"isCronEnabled,omitempty"`
	// Bool value indicating if the workflow run has removed
	IsRemoved bool `json:"isRemoved,omitempty"`
	// Flag to check if single run status is enabled or not
	IsSingleRunCronEnabled bool `json:"isSingleRunCronEnabled,omitempty"`
	// Notify ID of the experiment run
	NotifyID string `json:"notifyID,omitempty"`
	// Phase of the workflow run
	Phase *AllOfmodelWorkflowRunPhase `json:"phase,omitempty"`
	// Probe object containing reference of probeIDs
	Probe []ModelProbeMap `json:"probe,omitempty"`
	// Resiliency score of the workflow
	ResiliencyScore float64 `json:"resiliencyScore,omitempty"`
	// runSequence is the sequence number of experiment run
	RunSequence int32 `json:"runSequence,omitempty"`
	// Security Governance details of the workflow run
	SecurityGovernance *AllOfmodelWorkflowRunSecurityGovernance `json:"securityGovernance,omitempty"`
	// Total number of experiments
	TotalExperiments int32 `json:"totalExperiments,omitempty"`
	// Timestamp at which workflow run was last updated
	UpdatedAt string `json:"updatedAt,omitempty"`
	// User who has updated the workflow
	UpdatedBy *AllOfmodelWorkflowRunUpdatedBy `json:"updatedBy,omitempty"`
	// Array containing weightage and name of each chaos experiment in the workflow
	Weightages []ModelWeightages `json:"weightages,omitempty"`
	// Description of the workflow
	WorkflowDescription string `json:"workflowDescription,omitempty"`
	// ID of the workflow
	WorkflowID string `json:"workflowID,omitempty"`
	// Manifest of the workflow run
	WorkflowManifest string `json:"workflowManifest,omitempty"`
	// Name of the workflow
	WorkflowName string `json:"workflowName,omitempty"`
	// ID of the workflow run which is to be queried
	WorkflowRunID string `json:"workflowRunID,omitempty"`
	// Tag of the workflow
	WorkflowTags []string `json:"workflowTags,omitempty"`
	// Type of the workflow
	WorkflowType string `json:"workflowType,omitempty"`
}
