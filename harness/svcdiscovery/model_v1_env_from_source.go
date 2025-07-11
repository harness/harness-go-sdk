/*
 * Service Discovery API
 *
 * No description provided (generated by Swagger Codegen https://github.com/swagger-api/swagger-codegen)
 *
 * API version: 1.0
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */
package svcdiscovery

type V1EnvFromSource struct {
	ConfigMapRef *V1ConfigMapEnvSource `json:"configMapRef,omitempty"`
	// An optional identifier to prepend to each key in the ConfigMap. Must be a C_IDENTIFIER. +optional
	Prefix string `json:"prefix,omitempty"`
	SecretRef *V1SecretEnvSource `json:"secretRef,omitempty"`
}
