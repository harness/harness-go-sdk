/*
 * Harness NextGen Software Delivery Platform API Reference
 *
 * This is the Open Api Spec 3 for the DB Service
 *
 * API version: 1.0
 * Contact: contact@harness.io
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */
package dbops

// The status of changeset deployment
type DeployedStateOutput struct {
	Status      *CommandExecutionStatus `json:"status"`
	ChangeSetId string                  `json:"changeSetId"`
	Author      string                  `json:"author"`
	FileName    string                  `json:"fileName"`
	DeployedAt  int64                   `json:"deployedAt,omitempty"`
	Tag         string                  `json:"tag,omitempty"`
	Command     *Command                `json:"command"`
	Metadata    *ExecutionMetadata      `json:"metadata"`
	// if changeset run as part of current step execution
	DeployedInCurrentExecution bool `json:"deployedInCurrentExecution"`
}