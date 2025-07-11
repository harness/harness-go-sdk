/*
 * Chaos Manager API
 *
 * No description provided (generated by Swagger Codegen https://github.com/swagger-api/swagger-codegen)
 *
 * API version: 1.0
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */
package chaos

type ChaoshubFaultCategoriesCount struct {
	AWS int32 `json:"AWS,omitempty"`
	Azure int32 `json:"Azure,omitempty"`
	BYOC int32 `json:"BYOC,omitempty"`
	CloudFoundry int32 `json:"Cloud Foundry,omitempty"`
	GCP int32 `json:"GCP,omitempty"`
	KubeResilience int32 `json:"Kube Resilience,omitempty"`
	Kubernetes int32 `json:"Kubernetes,omitempty"`
	Linux int32 `json:"Linux,omitempty"`
	Load int32 `json:"Load,omitempty"`
	SSH int32 `json:"SSH,omitempty"`
	SpringBoot int32 `json:"Spring Boot,omitempty"`
	VMWare int32 `json:"VMWare,omitempty"`
	Windows int32 `json:"Windows,omitempty"`
}
