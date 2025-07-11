/*
 * Chaos Manager API
 *
 * No description provided (generated by Swagger Codegen https://github.com/swagger-api/swagger-codegen)
 *
 * API version: 1.0
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */
package chaos

type V1NewrelicMetrics struct {
	// Comparator check for the correctness of the probe output
	Comparator *AllOfv1NewrelicMetricsComparator `json:"comparator,omitempty"`
	// Mode on which metrcis will get evaluated(min/max/mean)
	MetricsEvaluationMode string `json:"metricsEvaluationMode,omitempty"`
	// NRQL query string
	Query string `json:"query,omitempty"`
	// NRQL Metrics that will get evaluated
	QueryMetrics string `json:"queryMetrics,omitempty"`
}
