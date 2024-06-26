/*
 * Governance Policy Management API
 *
 * Read and manage OPA Governance policies, policy sets and evaluations
 *
 * API version: 1.2.0
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */
package policymgmt

type GitErrorResult struct {
	// the explanation of the error
	Explanation string `json:"explanation"`
	// the hint on how to resolve the error
	Hint string `json:"hint"`
	// the message is a human-readable explanation specific to this occurrence of the problem
	Message string `json:"message"`
}
