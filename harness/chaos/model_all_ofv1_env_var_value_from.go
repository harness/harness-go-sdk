/*
 * Chaos Manager API
 *
 * No description provided (generated by Swagger Codegen https://github.com/swagger-api/swagger-codegen)
 *
 * API version: 1.0
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */
package chaos

// Source for the environment variable's value. Cannot be used if value is not empty. +optional
type AllOfv1EnvVarValueFrom struct {
	// Selects a key of a ConfigMap. +optional
	ConfigMapKeyRef *interface{} `json:"configMapKeyRef,omitempty"`
	// Selects a field of the pod: supports metadata.name, metadata.namespace, `metadata.labels['<KEY>']`, `metadata.annotations['<KEY>']`, spec.nodeName, spec.serviceAccountName, status.hostIP, status.podIP, status.podIPs. +optional
	FieldRef *interface{} `json:"fieldRef,omitempty"`
	// Selects a resource of the container: only resources limits and requests (limits.cpu, limits.memory, limits.ephemeral-storage, requests.cpu, requests.memory and requests.ephemeral-storage) are currently supported. +optional
	ResourceFieldRef *interface{} `json:"resourceFieldRef,omitempty"`
	// Selects a key of a secret in the pod's namespace +optional
	SecretKeyRef *interface{} `json:"secretKeyRef,omitempty"`
}
