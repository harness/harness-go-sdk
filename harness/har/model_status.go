/*
 * Harness Artifact Registry API
 *
 * No description provided (generated by Swagger Codegen https://github.com/swagger-api/swagger-codegen)
 *
 * API version: 1.0.0
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */
package har

// Status : Indicates if the request was successful or not
type Status string

// List of Status
const (
	SUCCESS_Status Status = "SUCCESS"
	FAILURE_Status Status = "FAILURE"
	ERROR__Status  Status = "ERROR"
)