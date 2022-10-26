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

// ResourceIgnoreDifferences contains resource filter and list of json paths which should be ignored during comparison with live state.
type ApplicationsResourceIgnoreDifferences struct {
	Group                 string   `json:"group,omitempty"`
	Kind                  string   `json:"kind,omitempty"`
	Name                  string   `json:"name,omitempty"`
	Namespace             string   `json:"namespace,omitempty"`
	JsonPointers          []string `json:"jsonPointers,omitempty"`
	JqPathExpressions     []string `json:"jqPathExpressions,omitempty"`
	ManagedFieldsManagers []string `json:"managedFieldsManagers,omitempty"`
}