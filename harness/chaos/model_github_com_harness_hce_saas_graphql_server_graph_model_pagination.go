/*
 * Chaos Manager API
 *
 * No description provided (generated by Swagger Codegen https://github.com/swagger-api/swagger-codegen)
 *
 * API version: 1.0
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */
package chaos

type GithubComHarnessHceSaasGraphqlServerGraphModelPagination struct {
	// Number of data to be fetched
	Limit int32 `json:"limit,omitempty"`
	// Page number for which data will be fetched
	Page int32 `json:"page,omitempty"`
}
