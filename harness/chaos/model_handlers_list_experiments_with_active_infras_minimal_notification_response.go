/*
 * Chaos Manager API
 *
 * No description provided (generated by Swagger Codegen https://github.com/swagger-api/swagger-codegen)
 *
 * API version: 1.0
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */
package chaos

type HandlersListExperimentsWithActiveInfrasMinimalNotificationResponse struct {
	// Details related to the workflows
	Experiments []HandlersMinimalNotificationWorkflow `json:"experiments,omitempty"`
	// Total number of workflows
	TotalNoOfExperiments int32 `json:"totalNoOfExperiments,omitempty"`
}
