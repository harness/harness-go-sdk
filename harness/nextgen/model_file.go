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

// This is details of the file entity defined in Harness.
type File struct {
	// Account Identifier for the Entity.
	AccountIdentifier string `json:"accountIdentifier,omitempty"`
	// Organization Identifier for the Entity.
	OrgIdentifier string `json:"orgIdentifier,omitempty"`
	// Project Identifier for the Entity.
	ProjectIdentifier string `json:"projectIdentifier,omitempty"`
	// Identifier of the File
	Identifier string `json:"identifier,omitempty"`
	// Name of the File
	Name string `json:"name,omitempty"`
	// This specifies the file usage
	FileUsage string `json:"fileUsage,omitempty"`
	// This specifies the type of the File
	Type_ string `json:"type"`
	// This specifies parent identifier
	ParentIdentifier string `json:"parentIdentifier,omitempty"`
	// Description of the File
	Description string `json:"description,omitempty"`
	// Tags
	Tags []NgTag `json:"tags,omitempty"`
	// Mime type of the File
	MimeType string `json:"mimeType,omitempty"`
}
