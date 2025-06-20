/*
 * Service Discovery API
 *
 * No description provided (generated by Swagger Codegen https://github.com/swagger-api/swagger-codegen)
 *
 * API version: 1.0
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */
package svcdiscovery

type ApiListPodResponse struct {
	CorrelationID string `json:"correlationID,omitempty"`
	Items []DatabasePodCollection `json:"items,omitempty"`
	Page *ApiPagination `json:"page,omitempty"`
}
