/*
 * Harness Artifact Registry API
 *
 * No description provided (generated by Swagger Codegen https://github.com/swagger-api/swagger-codegen)
 *
 * API version: 1.0.0
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */
package har

type DeploymentMetadata struct {
	NonProdEnvCount int32 `json:"nonProdEnvCount"`
	ProdEnvCount    int32 `json:"prodEnvCount"`
}