/*
 * API Specification
 *
 * No description provided (generated by Swagger Codegen https://github.com/swagger-api/swagger-codegen)
 *
 * API version: 0.0.0
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */
package code

type OpenapiCommitFilesRequest struct {
	Actions     []RepoCommitFileAction `json:"actions,omitempty"`
	Branch      string                 `json:"branch,omitempty"`
	BypassRules bool                   `json:"bypass_rules,omitempty"`
	DryRunRules bool                   `json:"dry_run_rules,omitempty"`
	Message     string                 `json:"message,omitempty"`
	NewBranch   string                 `json:"new_branch,omitempty"`
	Title       string                 `json:"title,omitempty"`
}
