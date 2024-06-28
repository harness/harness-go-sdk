/*
 * Infrastructure as Code Management
 *
 * Services for Harness IaCM Module.
 *
 * API version: 0.1.0
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */
package nextgen

// Per-workspace override to the assigned default pipelines.
type IacmDefaultPipelineOverride struct {
	// Default pipeline assigned as the project-level
	ProjectPipeline string `json:"project_pipeline"`
	// Default pipeline assigned as the workspace-level
	WorkspacePipeline string `json:"workspace_pipeline"`
}