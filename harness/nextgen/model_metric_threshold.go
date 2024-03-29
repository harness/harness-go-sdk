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

import "encoding/json"

type MetricThreshold struct {
	GroupName        string                   `json:"groupName,omitempty"`
	MetricName       string                   `json:"metricName,omitempty"`
	MetricIdentifier string                   `json:"identifier,omitempty"`
	MetricType       string                   `json:"metricType,omitempty"`
	Criteria         *MetricThresholdCriteria `json:"criteria,omitempty"`

	Type_           MetricThresholdType        `json:"type,omitempty"`
	FailImmediately *FailMetricThresholdSpec   `json:"-"`
	IgnoreThreshold *IgnoreMetricThresholdSpec `json:"-"`

	Spec json.RawMessage `json:"spec"`
}
