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

// Properties of the Files Filter defined in Harness
type FilesFilterPropertiesDto struct {
	// This specifies the file usage
	FileUsage    string                  `json:"fileUsage,omitempty"`
	CreatedBy    *EmbeddedUserDetailsDto `json:"createdBy,omitempty"`
	ReferencedBy *ReferencedByDto        `json:"referencedBy,omitempty"`
	// Filter tags as a key-value pair.
	Tags map[string]string `json:"tags,omitempty"`
	// This specifies the corresponding Entity of the filter.
	FilterType string `json:"filterType"`
}
