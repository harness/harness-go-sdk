/*
 * API Specification
 *
 * No description provided (generated by Swagger Codegen https://github.com/swagger-api/swagger-codegen)
 *
 * API version: 0.0.0
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */
package code

type OpenapiCreatePullReqRequest struct {
	Description   string `json:"description,omitempty"`
	IsDraft       bool   `json:"is_draft,omitempty"`
	SourceBranch  string `json:"source_branch,omitempty"`
	SourceRepoRef string `json:"source_repo_ref,omitempty"`
	TargetBranch  string `json:"target_branch,omitempty"`
	Title         string `json:"title,omitempty"`
}
