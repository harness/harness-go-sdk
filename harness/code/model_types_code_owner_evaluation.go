/*
 * API Specification
 *
 * No description provided (generated by Swagger Codegen https://github.com/swagger-api/swagger-codegen)
 *
 * API version: 0.0.0
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */
package code

type TypesCodeOwnerEvaluation struct {
	EvaluationEntries []TypesCodeOwnerEvaluationEntry `json:"evaluation_entries,omitempty"`
	FileSha           string                          `json:"file_sha,omitempty"`
}
