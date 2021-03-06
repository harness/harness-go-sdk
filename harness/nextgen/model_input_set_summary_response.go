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

// This is the view of the Input Set Summary.
type InputSetSummaryResponse struct {
	// Input Set Identifier
	Identifier string `json:"identifier,omitempty"`
	// Input Set Name
	Name string `json:"name,omitempty"`
	// Pipeline Identifier for the entity.
	PipelineIdentifier string `json:"pipelineIdentifier,omitempty"`
	// Input Set description
	Description string `json:"description,omitempty"`
	// Type of Input Set. The default value is ALL.
	InputSetType string `json:"inputSetType,omitempty"`
	// Input Set tags
	Tags       map[string]string         `json:"tags,omitempty"`
	Version    int64                     `json:"version,omitempty"`
	GitDetails *PipelineEntityGitDetails `json:"gitDetails,omitempty"`
	// Time at which the entity was created
	CreatedAt int64 `json:"createdAt,omitempty"`
	// Time at which the entity was last updated
	LastUpdatedAt int64 `json:"lastUpdatedAt,omitempty"`
	// This field is true if a Pipeline update has made this Input Set invalid, and cannot be used for Pipeline Execution
	IsOutdated           bool                  `json:"isOutdated,omitempty"`
	InputSetErrorDetails *InputSetErrorWrapper `json:"inputSetErrorDetails,omitempty"`
	// This contains the invalid references in the Overlay Input Set, along with a message why they are invalid
	OverlaySetErrorDetails map[string]string         `json:"overlaySetErrorDetails,omitempty"`
	EntityValidityDetails  *PipelineEntityGitDetails `json:"entityValidityDetails,omitempty"`
	// Modules in which the Pipeline belongs
	Modules []string `json:"modules,omitempty"`
}
