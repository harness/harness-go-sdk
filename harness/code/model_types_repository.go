/*
 * API Specification
 *
 * No description provided (generated by Swagger Codegen https://github.com/swagger-api/swagger-codegen)
 *
 * API version: 0.0.0
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */
package code

type TypesRepository struct {
	Created        int32  `json:"created,omitempty"`
	CreatedBy      int32  `json:"created_by,omitempty"`
	DefaultBranch  string `json:"default_branch,omitempty"`
	Deleted        int32  `json:"deleted,omitempty"`
	Description    string `json:"description,omitempty"`
	ForkId         int32  `json:"fork_id,omitempty"`
	GitUrl         string `json:"git_url,omitempty"`
	Id             int32  `json:"id,omitempty"`
	Identifier     string `json:"identifier,omitempty"`
	Importing      bool   `json:"importing,omitempty"`
	IsPublic       bool   `json:"is_public,omitempty"`
	NumClosedPulls int32  `json:"num_closed_pulls,omitempty"`
	NumForks       int32  `json:"num_forks,omitempty"`
	NumMergedPulls int32  `json:"num_merged_pulls,omitempty"`
	NumOpenPulls   int32  `json:"num_open_pulls,omitempty"`
	NumPulls       int32  `json:"num_pulls,omitempty"`
	ParentId       int32  `json:"parent_id,omitempty"`
	Path           string `json:"path,omitempty"`
	Size           int32  `json:"size,omitempty"`
	SizeUpdated    int32  `json:"size_updated,omitempty"`
	Updated        int32  `json:"updated,omitempty"`
}
