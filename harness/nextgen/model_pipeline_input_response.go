/*
 * Harness NextGen Software Delivery Platform API Reference
 *
 * This is the Open Api Spec 3 for the NextGen Manager. This is under active development. Beware of the breaking change with respect to the generated code stub  # Authentication  <!-- ReDoc-Inject: <security-definitions> -->
 *
 * API version: 3.0
 * Contact: contact@harness.io
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */
package nextgen

type PipelineInputResponse struct {
	ErrorInfo *PreFlightEntityErrorInfo `json:"errorInfo,omitempty"`
	Success   bool                      `json:"success,omitempty"`
	Fqn       string                    `json:"fqn,omitempty"`
	StageName string                    `json:"stageName,omitempty"`
	StepName  string                    `json:"stepName,omitempty"`
}
