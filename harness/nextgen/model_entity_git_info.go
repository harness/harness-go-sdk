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

// EntityGitInfo contains git info of the entity
type EntityGitInfo struct {
	ObjectId  string `json:"objectId,omitempty"`  // Object Id of the Entity
	Branch    string `json:"branch,omitempty"`    // Branch Name
	FilePath  string `json:"filePath,omitempty"`  // File Path of the Entity
	RepoName  string `json:"repoName,omitempty"`  // Name of the repo
	CommitId  string `json:"commitId,omitempty"`  // Commit Id of the Entity
	FileUrl   string `json:"fileUrl,omitempty"`   // File Url of the entity
	RepoUrl   string `json:"repoUrl,omitempty"`   // Repo url of the entity
}
