/*
 * Harness NextGen Software Delivery Platform API Reference
 *
 * This is the Open Api Spec 3 for the NextGen Manager. This is under active development. Beware of the breaking change with respect to the generated code stub
 *
 * API version: 3.0
 * Contact: contact@harness.io
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */
package nextgen

// ApplicationsApplicationSetResourceIgnoreDifferences ApplicationSetResourceIgnoreDifferences configures how the ApplicationSet controller will ignore differences in live applications when applying changes from generated applications.
type ApplicationsApplicationSetResourceIgnoreDifferences struct {
	// Name is the name of the application to ignore differences for. If not specified, the rule applies to all applications.
	Name string `json:"name,omitempty"`
	// JSONPointers is a list of JSON pointers to fields to ignore differences for.
	JsonPointers []string `json:"jsonPointers,omitempty"`
	// JQPathExpressions is a list of JQ path expressions to fields to ignore differences for.
	JqPathExpressions []string `json:"jqPathExpressions,omitempty"`
}
