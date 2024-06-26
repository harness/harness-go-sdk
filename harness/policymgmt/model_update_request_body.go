/*
 * Governance Policy Management API
 *
 * Read and manage OPA Governance policies, policy sets and evaluations
 *
 * API version: 1.2.0
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */
package policymgmt

type UpdateRequestBody struct {
	// The name of the policy
	Name string `json:"name,omitempty"`
	// The rego that defines the policy policy
	Rego string `json:"rego,omitempty"`
}
