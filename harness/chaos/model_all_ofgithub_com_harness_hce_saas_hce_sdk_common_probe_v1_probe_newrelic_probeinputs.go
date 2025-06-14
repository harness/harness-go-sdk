/*
 * Chaos Manager API
 *
 * No description provided (generated by Swagger Codegen https://github.com/swagger-api/swagger-codegen)
 *
 * API version: 1.0
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */
package chaos

// inputs needed for the Newrelic probe
type AllOfgithubComHarnessHceSaasHceSdkCommonProbeV1ProbeNewrelicProbeinputs struct {
	// Newrelic accountId
	AccountId string `json:"accountId,omitempty"`
	// APITokenSecretName for authenticating with the Newrelic platform
	ApiTokenSecretName string `json:"apiTokenSecretName,omitempty"`
	// newrelic site URL identifier
	Endpoint string `json:"endpoint,omitempty"`
	// Newrelic metrics parameters
	NewrelicMetrics *interface{} `json:"newrelicMetrics,omitempty"`
}
