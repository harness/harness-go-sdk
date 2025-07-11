/*
 * Service Discovery API
 *
 * No description provided (generated by Swagger Codegen https://github.com/swagger-api/swagger-codegen)
 *
 * API version: 1.0
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */
package svcdiscovery

type ApiInfrastructure struct {
	AgentName string `json:"agentName,omitempty"`
	ConnectorRef string `json:"connectorRef,omitempty"`
	EnvironmentRef string `json:"environmentRef,omitempty"`
	Identifier string `json:"identifier,omitempty"`
	IsCompatible bool `json:"isCompatible,omitempty"`
	IsUsed bool `json:"isUsed,omitempty"`
	Name string `json:"name,omitempty"`
	Namespace string `json:"namespace,omitempty"`
	OrgIdentifier string `json:"orgIdentifier,omitempty"`
	ProjectIdentifier string `json:"projectIdentifier,omitempty"`
	Type_ string `json:"type,omitempty"`
}
