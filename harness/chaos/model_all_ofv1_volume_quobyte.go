/*
 * No description provided (generated by Swagger Codegen https://github.com/swagger-api/swagger-codegen)
 *
 * API version: 1.0.0
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */
package swagger

// Quobyte represents a Quobyte mount on the host that shares a pod's lifetime +optional
type AllOfv1VolumeQuobyte struct {
	// Group to map volume access to Default is no group +optional
	Group string `json:"group,omitempty"`
	// ReadOnly here will force the Quobyte volume to be mounted with read-only permissions. Defaults to false. +optional
	ReadOnly bool `json:"readOnly,omitempty"`
	// Registry represents a single or multiple Quobyte Registry services specified as a string as host:port pair (multiple entries are separated with commas) which acts as the central registry for volumes
	Registry string `json:"registry,omitempty"`
	// Tenant owning the given Quobyte volume in the Backend Used with dynamically provisioned Quobyte volumes, value is set by the plugin +optional
	Tenant string `json:"tenant,omitempty"`
	// User to map volume access to Defaults to serivceaccount user +optional
	User string `json:"user,omitempty"`
	// Volume is a string that references an already created Quobyte volume by name.
	Volume string `json:"volume,omitempty"`
}