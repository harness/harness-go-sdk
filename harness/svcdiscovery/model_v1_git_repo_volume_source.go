/*
 * Service Discovery API
 *
 * No description provided (generated by Swagger Codegen https://github.com/swagger-api/swagger-codegen)
 *
 * API version: 1.0
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */
package svcdiscovery

type V1GitRepoVolumeSource struct {
	// directory is the target directory name. Must not contain or start with '..'.  If '.' is supplied, the volume directory will be the git repository.  Otherwise, if specified, the volume will contain the git repository in the subdirectory with the given name. +optional
	Directory string `json:"directory,omitempty"`
	// repository is the URL
	Repository string `json:"repository,omitempty"`
	// revision is the commit hash for the specified revision. +optional
	Revision string `json:"revision,omitempty"`
}
