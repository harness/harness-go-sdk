/*
 * Chaos Manager API
 *
 * No description provided (generated by Swagger Codegen https://github.com/swagger-api/swagger-codegen)
 *
 * API version: 1.0
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */
package chaos

type TypesListProbeResponse struct {
	Data []TypesGetProbeResponse `json:"data,omitempty"`
	Pagination *GithubComHarnessHceSaasGraphqlServerApiPagination `json:"pagination,omitempty"`
	TotalNoOfProbes int32 `json:"totalNoOfProbes,omitempty"`
}
