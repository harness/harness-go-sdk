/*
 * API Specification
 *
 * No description provided (generated by Swagger Codegen https://github.com/swagger-api/swagger-codegen)
 *
 * API version: 0.0.0
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */
package code

type ReposImportBody struct {
	Description  string                  `json:"description,omitempty"`
	Identifier   string                  `json:"identifier,omitempty"`
	ParentRef    string                  `json:"parent_ref,omitempty"`
	Pipelines    *ImporterPipelineOption `json:"pipelines,omitempty"`
	Provider     *ImporterProvider       `json:"provider,omitempty"`
	ProviderRepo string                  `json:"provider_repo,omitempty"`
	Uid          string                  `json:"uid,omitempty"`
}
