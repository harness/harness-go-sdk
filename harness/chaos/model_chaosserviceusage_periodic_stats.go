/*
 * Chaos Manager API
 *
 * No description provided (generated by Swagger Codegen https://github.com/swagger-api/swagger-codegen)
 *
 * API version: 1.0
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */
package chaos

type ChaosserviceusagePeriodicStats struct {
	ServiceStatsByType *ChaosserviceusageServiceStatsByType `json:"serviceStatsByType,omitempty"`
	Timestamp int32 `json:"timestamp,omitempty"`
	TotalConsumption float64 `json:"totalConsumption,omitempty"`
}
