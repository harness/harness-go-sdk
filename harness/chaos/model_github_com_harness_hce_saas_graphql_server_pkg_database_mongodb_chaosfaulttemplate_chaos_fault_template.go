/*
 * Chaos Manager API
 *
 * No description provided (generated by Swagger Codegen https://github.com/swagger-api/swagger-codegen)
 *
 * API version: 1.0
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */
package chaos

type GithubComHarnessHceSaasGraphqlServerPkgDatabaseMongodbChaosfaulttemplateChaosFaultTemplate struct {
	AccountID string `json:"accountID"`
	// creation timestamp of the revision
	CreatedAt int32 `json:"createdAt,omitempty"`
	// user ID of the user who created the revision
	CreatedBy string `json:"createdBy,omitempty"`
	// Hub reference to sync the changes whenever there are changes in the faults
	HubRef string `json:"hubRef,omitempty"`
	// Mongo ID (primary key)
	Id string `json:"id,omitempty"`
	// Unique identifier (human-readable) immutable Initially it will be same as name
	Identity string `json:"identity,omitempty"`
	// isDefault indicates if it is the default version for predefined faults, latest should be set as default
	IsDefault bool `json:"isDefault,omitempty"`
	// isRemoved indicates if the document is deleted
	IsRemoved bool `json:"isRemoved,omitempty"`
	// Fault name to sync the changes from the hub HubRef + Name should be unique
	Name string `json:"name,omitempty"`
	OrgID string `json:"orgID,omitempty"`
	ProjectID string `json:"projectID,omitempty"`
	// Revision is the version of fault template, it increments every time a new version of fault is published
	Revision int32 `json:"revision,omitempty"`
	// template of the fault
	Template string `json:"template,omitempty"`
	// Variables for template
	Variables []TemplateVariable `json:"variables,omitempty"`
}
