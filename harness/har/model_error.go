/*
 * Harness Artifact Registry API
 *
 * No description provided (generated by Swagger Codegen https://github.com/swagger-api/swagger-codegen)
 *
 * API version: 1.0.0
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */
package har

type ModelError struct {
	// The http error code
	Code string `json:"code"`
	// Additional details about the error
	Details *interface{} `json:"details,omitempty"`
	// The reason the request failed
	Message string `json:"message"`
}
