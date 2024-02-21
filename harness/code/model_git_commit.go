/*
 * API Specification
 *
 * No description provided (generated by Swagger Codegen https://github.com/swagger-api/swagger-codegen)
 *
 * API version: 0.0.0
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */
package code

type GitCommit struct {
	Author    *GitSignature `json:"author,omitempty"`
	Committer *GitSignature `json:"committer,omitempty"`
	Message   string        `json:"message,omitempty"`
	Sha       string        `json:"sha,omitempty"`
	Title     string        `json:"title,omitempty"`
}