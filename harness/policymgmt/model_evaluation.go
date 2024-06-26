/*
 * Governance Policy Management API
 *
 * Read and manage OPA Governance policies, policy sets and evaluations
 *
 * API version: 1.2.0
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */
package policymgmt
import (
	"os"
)

type Evaluation struct {
	// The Harness account in which the evaluation was performed
	AccountId string `json:"account_id"`
	// The action that triggered evaluation
	Action string `json:"action"`
	// The time at which the evaluation was performed in Unix time millseconds
	Created int64 `json:"created"`
	// The detailed results of te evaluation
	Details []EvaluationDetail `json:"details"`
	// An arbtrary user-supplied string that globally identifies the entity under evaluation
	Entity string `json:"entity"`
	// Additional arbtrary user-supplied metadta about the entity under evaluation
	EntityMetadata string `json:"entity_metadata"`
	// The ID of this evaluation
	Id int64 `json:"id"`
	// The input provided at evaluation time
	Input **os.File `json:"input"`
	// The Harness organisation in which the evaluation was performed
	OrgId string `json:"org_id"`
	// The Harness project in which the evaluation was performed
	ProjectId string `json:"project_id"`
	// The overall status of the evaluation indicating whether it passed
	Status string `json:"status"`
	// The types of the entity under evaluation
	Type_ string `json:"type"`
}
