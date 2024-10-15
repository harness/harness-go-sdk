/*
 * No description provided (generated by Swagger Codegen https://github.com/swagger-api/swagger-codegen)
 *
 * API version: 1.0.0
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */
package swagger

// Selects a resource of the container: only resources limits and requests (limits.cpu, limits.memory, limits.ephemeral-storage, requests.cpu, requests.memory and requests.ephemeral-storage) are currently supported. +optional
type AllOfv1EnvVarSourceResourceFieldRef struct {
	// Container name: required for volumes, optional for env vars +optional
	ContainerName string `json:"containerName,omitempty"`
	// Specifies the output format of the exposed resources, defaults to \"1\" +optional
	Divisor *interface{} `json:"divisor,omitempty"`
	// Required: resource to select
	Resource string `json:"resource,omitempty"`
}