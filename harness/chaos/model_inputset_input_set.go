/*
 * Chaos Manager API
 *
 * No description provided (generated by Swagger Codegen https://github.com/swagger-api/swagger-codegen)
 *
 * API version: 1.0
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */
package chaos

type InputsetInputSet struct {
	AccountID string `json:"accountID"`
	// creation timestamp of the input set
	CreatedAt int32 `json:"createdAt,omitempty"`
	// user ID of the user who created the input set
	CreatedBy string `json:"createdBy,omitempty"`
	// Description of the input set
	Description string `json:"description,omitempty"`
	// Foreign key to link with experiment
	ExperimentID string `json:"experimentID,omitempty"`
	// Mongo ID (primary key)
	Id string `json:"id,omitempty"`
	// Human readable ID
	Identity string `json:"identity,omitempty"`
	// TODO: this is not needed, and on delete, input set should be deleted from the DB, makes no sense for storing for audit purpose
	IsRemoved bool `json:"isRemoved,omitempty"`
	// Name of the input set
	Name string `json:"name,omitempty"`
	OrgID string `json:"orgID,omitempty"`
	ProjectID string `json:"projectID,omitempty"`
	// Type of input set Type string `bson:\"type\"` Foreign key to link with probes TODO: not sure if required ProbeID string `bson:\"probe_id\"` For fault level variables, key = step
	Spec string `json:"spec,omitempty"`
	// updation timestamp of the input set
	UpdatedAt int32 `json:"updatedAt,omitempty"`
	// user ID of the user who updated the input set
	UpdatedBy string `json:"updatedBy,omitempty"`
	// Version
	Version string `json:"version,omitempty"`
}
