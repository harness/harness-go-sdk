/*
 * Chaos Manager API
 *
 * No description provided (generated by Swagger Codegen https://github.com/swagger-api/swagger-codegen)
 *
 * API version: 1.0
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */
package chaos

type TargetserviceDiscoveredServiceSpec struct {
	Id string `json:"id,omitempty"`
	Kind string `json:"kind,omitempty"`
	Label map[string]string `json:"label,omitempty"`
	Name string `json:"name,omitempty"`
	Namespace string `json:"namespace,omitempty"`
	Uuid string `json:"uuid,omitempty"`
}
