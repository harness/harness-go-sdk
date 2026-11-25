/*
 * Infrastructure as Code Management
 *
 * Services for Harness IaCM Module.
 */
package nextgen

const iacmWorkspaceDriftConfigPath = "/iacm/api/orgs/{org}/projects/{project}/workspaces/{workspace}/drift-config"

type IacmWorkspaceDriftConfigRequestBody struct {
	// Pipeline to run for drift detection
	PipelineId string `json:"pipeline_id"`
	// Provisioner type to enable drift for
	Provisioner string `json:"provisioner"`
	// Run drift detection every N units (see RunEveryUnit)
	RunEvery int32 `json:"run_every"`
	// Unit for RunEvery interval
	RunEveryUnit string `json:"run_every_unit"`
}

type IacmWorkspaceDriftConfigResponse struct {
	// Whether workspace-level drift automation is currently enabled
	Enabled bool `json:"enabled"`
	// Pipeline configured to run for drift detection
	PipelineId string `json:"pipeline_id,omitempty"`
	// Provisioner this configuration applies to
	Provisioner string `json:"provisioner,omitempty"`
	// Repeat interval; set when drift is enabled
	RunEvery int32 `json:"run_every,omitempty"`
	// Unit for RunEvery interval
	RunEveryUnit string `json:"run_every_unit,omitempty"`
	// Generated trigger identifier
	TriggerId string `json:"trigger_id,omitempty"`
}
