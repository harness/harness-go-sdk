/*
 * Harness Artifact Registry API
 *
 * No description provided (generated by Swagger Codegen https://github.com/swagger-api/swagger-codegen)
 *
 * API version: 1.0.0
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */
package har

// Harness Artifact Registry
type Registry struct {
	AllowedPattern []string        `json:"allowedPattern,omitempty"`
	BlockedPattern []string        `json:"blockedPattern,omitempty"`
	CleanupPolicy  []CleanupPolicy `json:"cleanupPolicy,omitempty"`
	Config         *RegistryConfig `json:"config,omitempty"`
	CreatedAt      string          `json:"createdAt,omitempty"`
	Description    string          `json:"description,omitempty"`
	Identifier     string          `json:"identifier"`
	Labels         []string        `json:"labels,omitempty"`
	ModifiedAt     string          `json:"modifiedAt,omitempty"`
	PackageType    *PackageType    `json:"packageType"`
	Url            string          `json:"url"`
}