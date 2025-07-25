/*
 * Chaos Manager API
 *
 * No description provided (generated by Swagger Codegen https://github.com/swagger-api/swagger-codegen)
 *
 * API version: 1.0
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */
package chaos

// Comparator check for the correctness of the probe output
type AllOfv1DatadogMetricsComparator struct {
	// Criteria for matching data it supports >=, <=, ==, >, <, != for int and float it supports equal, notEqual, contains for string
	Criteria string `json:"criteria,omitempty"`
	// Type of data it can be int, float, string
	Type_ string `json:"type,omitempty"`
	// Value contains relative value for criteria
	Value string `json:"value,omitempty"`
}
