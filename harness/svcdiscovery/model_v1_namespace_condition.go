/*
 * Service Discovery API
 *
 * No description provided (generated by Swagger Codegen https://github.com/swagger-api/swagger-codegen)
 *
 * API version: 1.0
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */
package svcdiscovery

type V1NamespaceCondition struct {
	// +optional
	LastTransitionTime string `json:"lastTransitionTime,omitempty"`
	// +optional
	Message string `json:"message,omitempty"`
	// +optional
	Reason string `json:"reason,omitempty"`
	// Status of the condition, one of True, False, Unknown.
	Status string `json:"status,omitempty"`
	// Type of namespace controller condition.
	Type_ string `json:"type,omitempty"`
}
