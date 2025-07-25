/*
 * Harness Artifact Registry API
 *
 * No description provided (generated by Swagger Codegen https://github.com/swagger-api/swagger-codegen)
 *
 * API version: 1.0.0
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */
package har
// Trigger : refers to trigger
type Trigger string

// List of Trigger
const (
	CREATION_Trigger Trigger = "ARTIFACT_CREATION"
	DELETION_Trigger Trigger = "ARTIFACT_DELETION"
)
