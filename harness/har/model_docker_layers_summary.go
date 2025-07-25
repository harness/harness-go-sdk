/*
 * Harness Artifact Registry API
 *
 * No description provided (generated by Swagger Codegen https://github.com/swagger-api/swagger-codegen)
 *
 * API version: 1.0.0
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */
package har

// Harness Layers Summary
type DockerLayersSummary struct {
	Digest string `json:"digest"`
	Layers []DockerLayerEntry `json:"layers,omitempty"`
	OsArch string `json:"osArch,omitempty"`
}
