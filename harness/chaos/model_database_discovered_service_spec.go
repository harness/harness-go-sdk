/*
 * Chaos Manager API
 *
 * No description provided (generated by Swagger Codegen https://github.com/swagger-api/swagger-codegen)
 *
 * API version: 1.0
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */
package chaos

type DatabaseDiscoveredServiceSpec struct {
	Fqdn []string `json:"fqdn,omitempty"`
	HarnessEnvironmentIdentity *DatabaseEnvironmentIdentity `json:"harnessEnvironmentIdentity,omitempty"`
	HarnessServiceIdentity *DatabaseServiceIdentity `json:"harnessServiceIdentity,omitempty"`
	Ip []string `json:"ip,omitempty"`
	Kubernetes *DatabaseDiscoveredServiceKubernetesSpec `json:"kubernetes,omitempty"`
	Port []string `json:"port,omitempty"`
}
