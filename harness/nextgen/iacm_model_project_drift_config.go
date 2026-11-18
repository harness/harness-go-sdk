/*
 * Infrastructure as Code Management
 *
 * Services for Harness IaCM Module.
 */
package nextgen

const iacmProjectDriftConfigPath = "/iacm/api/orgs/{org}/projects/{project}/settings/drift-config"

type IacmProjectDriftConfigRequestBody struct {
	// Pipeline to run for drift detection
	PipelineId string `json:"pipeline_id"`
	// Provisioner type to enable drift for
	Provisioner string `json:"provisioner"`
	// Run drift detection every N units (see RunEveryUnit)
	RunEvery int32 `json:"run_every"`
	// Unit for RunEvery interval
	RunEveryUnit string `json:"run_every_unit"`
}

type IacmProjectDriftConfigResponse struct {
	// Whether project-level drift automation is currently enabled for this provisioner
	Enabled bool `json:"enabled"`
	// Pipeline configured to run for drift detection
	PipelineId string `json:"pipeline_id,omitempty"`
	// Provisioner this configuration applies to
	Provisioner string `json:"provisioner"`
	// Repeat interval; set when drift is enabled
	RunEvery int32 `json:"run_every,omitempty"`
	// Unit for RunEvery interval
	RunEveryUnit string `json:"run_every_unit,omitempty"`
}
