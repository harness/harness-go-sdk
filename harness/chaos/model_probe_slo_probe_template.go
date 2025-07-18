/*
 * Chaos Manager API
 *
 * No description provided (generated by Swagger Codegen https://github.com/swagger-api/swagger-codegen)
 *
 * API version: 1.0
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */
package chaos

type ProbeSloProbeTemplate struct {
	Comparator *ProbeComparatorTemplate `json:"comparator,omitempty"`
	EvaluationTimeout string `json:"evaluationTimeout,omitempty"`
	EvaluationWindow *ProbeEvaluationWindowTemplate `json:"evaluationWindow,omitempty"`
	InsecureSkipVerify bool `json:"insecureSkipVerify,omitempty"`
	PlatformEndpoint string `json:"platformEndpoint,omitempty"`
	SloIdentifier string `json:"sloIdentifier,omitempty"`
	SloSourceMetadata *ProbeSloSourceMetadataTemplate `json:"sloSourceMetadata,omitempty"`
}
