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

type VariableResponseDto struct {
	Variable *VariableDto `json:"variable"`
	// Time when the Variable was first created.
	CreatedAt int64 `json:"createdAt,omitempty"`
	// Time when the Variable was last modified.
	LastModifiedAt int64 `json:"lastModifiedAt,omitempty"`
}
