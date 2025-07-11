/*
 * Chaos Manager API
 *
 * No description provided (generated by Swagger Codegen https://github.com/swagger-api/swagger-codegen)
 *
 * API version: 1.0
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */
package chaos

type ProbeDatadogProbeTemplate struct {
	DatadogCredentialsSecretName string `json:"datadogCredentialsSecretName,omitempty"`
	DatadogSite string `json:"datadogSite,omitempty"`
	Metrics *ProbeDatadogMetricsTemplate `json:"metrics,omitempty"`
	SyntheticsTest *ProbeSyntheticsTestTemplate `json:"syntheticsTest,omitempty"`
}
