/*
 * Chaos Manager API
 *
 * No description provided (generated by Swagger Codegen https://github.com/swagger-api/swagger-codegen)
 *
 * API version: 1.0
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */
package chaos

type RecommendationFailure struct {
	ContainerName string `json:"containerName,omitempty"`
	EventIdentifier string `json:"eventIdentifier,omitempty"`
	KubernetesDoc string `json:"kubernetesDoc,omitempty"`
	Text string `json:"text,omitempty"`
}
