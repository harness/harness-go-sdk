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

type MetricResponseMapping struct {
	MetricValueJsonPath string `json:"metricValueJsonPath,omitempty"`
	TimestampJsonPath string `json:"timestampJsonPath,omitempty"`
	ServiceInstanceJsonPath string `json:"serviceInstanceJsonPath,omitempty"`
	ServiceInstanceListJsonPath string `json:"serviceInstanceListJsonPath,omitempty"`
	RelativeMetricListJsonPath string `json:"relativeMetricListJsonPath,omitempty"`
	RelativeTimestampJsonPath string `json:"relativeTimestampJsonPath,omitempty"`
	RelativeMetricValueJsonPath string `json:"relativeMetricValueJsonPath,omitempty"`
	RelativeServiceInstanceValueJsonPath string `json:"relativeServiceInstanceValueJsonPath,omitempty"`
	TimestampFormat string `json:"timestampFormat,omitempty"`
}
