/*
 * Governance Policy Management API
 *
 * Read and manage OPA Governance policies, policy sets and evaluations
 *
 * API version: 1.2.0
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */
package policymgmt

type ResourceGroupIdentifier struct {
	// resource group account id
	AccountId string `json:"account_id"`
	// resource group identifier
	Identifier string `json:"identifier"`
	// resource group org id
	OrgId string `json:"org_id"`
	// resource group project identifier
	ProjectId string `json:"project_id"`
}