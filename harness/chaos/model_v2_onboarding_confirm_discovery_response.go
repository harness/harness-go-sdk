/*
 * Chaos Manager API
 *
 * No description provided (generated by Swagger Codegen https://github.com/swagger-api/swagger-codegen)
 *
 * API version: 1.0
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */
package chaos

type V2OnboardingConfirmDiscoveryResponse struct {
	AgentCreated bool `json:"agentCreated,omitempty"`
	AgentLinked bool `json:"agentLinked,omitempty"`
	DiscoveryCompleted bool `json:"discoveryCompleted,omitempty"`
}
