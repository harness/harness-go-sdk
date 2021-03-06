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

// This is the view of the Pipeline Summary for Pipeline entity defined in Harness.
type PmsPipelineSummaryResponse struct {
	Name                  string                            `json:"name,omitempty"`
	Identifier            string                            `json:"identifier,omitempty"`
	Description           string                            `json:"description,omitempty"`
	Tags                  map[string]string                 `json:"tags,omitempty"`
	Version               int64                             `json:"version,omitempty"`
	NumOfStages           int32                             `json:"numOfStages,omitempty"`
	CreatedAt             int64                             `json:"createdAt,omitempty"`
	LastUpdatedAt         int64                             `json:"lastUpdatedAt,omitempty"`
	Modules               []string                          `json:"modules,omitempty"`
	ExecutionSummaryInfo  *ExecutionSummaryInfo             `json:"executionSummaryInfo,omitempty"`
	Filters               map[string]map[string]interface{} `json:"filters,omitempty"`
	StageNames            []string                          `json:"stageNames,omitempty"`
	GitDetails            *PipelineEntityGitDetails         `json:"gitDetails,omitempty"`
	EntityValidityDetails *PipelineEntityGitDetails         `json:"entityValidityDetails,omitempty"`
}
