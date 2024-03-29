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

// This is the Dynatrace Metric Health Source spec entity defined in Harness
type DynatraceHealthSource struct {
	ConnectorRef      string                      `json:"connectorRef"`
	MetricPacks       []TimeSeriesMetricPackDto   `json:"metricPacks,omitempty"`
	Feature           string                      `json:"feature"`
	ServiceId         string                      `json:"serviceId"`
	ServiceName       string                      `json:"serviceName,omitempty"`
	ServiceMethodIds  []string                    `json:"serviceMethodIds,omitempty"`
	MetricDefinitions []DynatraceMetricDefinition `json:"metricDefinitions,omitempty"`
}
