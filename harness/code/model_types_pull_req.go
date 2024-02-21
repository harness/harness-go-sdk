/*
 * API Specification
 *
 * No description provided (generated by Swagger Codegen https://github.com/swagger-api/swagger-codegen)
 *
 * API version: 0.0.0
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */
package code

type TypesPullReq struct {
	Author           *TypesPrincipalInfo `json:"author,omitempty"`
	Created          int32               `json:"created,omitempty"`
	Description      string              `json:"description,omitempty"`
	Edited           int32               `json:"edited,omitempty"`
	IsDraft          bool                `json:"is_draft,omitempty"`
	MergeBaseSha     string              `json:"merge_base_sha,omitempty"`
	MergeCheckStatus string              `json:"merge_check_status,omitempty"`
	MergeConflicts   []string            `json:"merge_conflicts,omitempty"`
	MergeMethod      *EnumMergeMethod    `json:"merge_method,omitempty"`
	MergeSha         string              `json:"merge_sha,omitempty"`
	MergeTargetSha   string              `json:"merge_target_sha,omitempty"`
	Merged           int32               `json:"merged,omitempty"`
	Merger           *TypesPrincipalInfo `json:"merger,omitempty"`
	Number           int32               `json:"number,omitempty"`
	SourceBranch     string              `json:"source_branch,omitempty"`
	SourceRepoId     int32               `json:"source_repo_id,omitempty"`
	SourceSha        string              `json:"source_sha,omitempty"`
	State            *EnumPullReqState   `json:"state,omitempty"`
	Stats            *TypesPullReqStats  `json:"stats,omitempty"`
	TargetBranch     string              `json:"target_branch,omitempty"`
	TargetRepoId     int32               `json:"target_repo_id,omitempty"`
	Title            string              `json:"title,omitempty"`
}