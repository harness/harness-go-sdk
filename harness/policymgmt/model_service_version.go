/*
 * Governance Policy Management API
 *
 * Read and manage OPA Governance policies, policy sets and evaluations
 *
 * API version: 1.2.0
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */
package policymgmt

type ServiceVersion struct {
	// Build identifier
	Commit string `json:"commit"`
	// Version number
	Version string `json:"version"`
}
