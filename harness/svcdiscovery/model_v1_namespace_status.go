/*
 * Service Discovery API
 *
 * No description provided (generated by Swagger Codegen https://github.com/swagger-api/swagger-codegen)
 *
 * API version: 1.0
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */
package svcdiscovery

type V1NamespaceStatus struct {
	// Represents the latest available observations of a namespace's current state. +optional +patchMergeKey=type +patchStrategy=merge
	Conditions []V1NamespaceCondition `json:"conditions,omitempty"`
	// Phase is the current lifecycle phase of the namespace. More info: https://kubernetes.io/docs/tasks/administer-cluster/namespaces/ +optional
	Phase string `json:"phase,omitempty"`
}
