/*
 * Service Discovery API
 *
 * No description provided (generated by Swagger Codegen https://github.com/swagger-api/swagger-codegen)
 *
 * API version: 1.0
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */
package svcdiscovery

type ApiCreateNetworkMapRequest struct {
	Connections []DatabaseConnection `json:"connections,omitempty"`
	Description string `json:"description,omitempty"`
	Identity string `json:"identity"`
	Name string `json:"name"`
	Resources []DatabaseNetworkMapEntity `json:"resources"`
	Rules *DatabaseNetworkMapRules `json:"rules,omitempty"`
	Tags []string `json:"tags,omitempty"`
}
