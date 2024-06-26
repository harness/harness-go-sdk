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

type QueryParamsDto struct {
	ServiceInstanceField        string `json:"serviceInstanceField,omitempty"`
	Index                       string `json:"index,omitempty"`
	TimeStampIdentifier         string `json:"timeStampIdentifier,omitempty"`
	TimeStampFormat             string `json:"timeStampFormat,omitempty"`
	MessageIdentifier           string `json:"messageIdentifier,omitempty"`
	HealthSourceMetricName      string `json:"healthSourceMetricName,omitempty"`
	HealthSourceMetricNamespace string `json:"healthSourceMetricNamespace,omitempty"`
	AggregationType             string `json:"aggregationType,omitempty"`
}
