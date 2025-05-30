/*
 * Harness NextGen Software Delivery Platform API Reference
 *
 * This is the Open Api Spec 3 for the NextGen Manager. This is under active development. Beware of the breaking change with respect to the generated code stub
 *
 * API version: 3.0
 * Contact: contact@harness.io
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */
package nextgen

// PullRequestGenerator defines connection info specific to GitHub.
type ApplicationsPullRequestGeneratorGithub struct {
	// GitHub org or user to scan. Required.
	Owner string `json:"owner,omitempty"`
	// GitHub repo name to scan. Required.
	Repo string `json:"repo,omitempty"`
	// The GitHub API URL to talk to. If blank, use https://api.github.com/.
	Api string `json:"api,omitempty"`
	TokenRef *ApplicationsSecretRef `json:"tokenRef,omitempty"`
	// AppSecretName is a reference to a GitHub App repo-creds secret with permission to access pull requests.
	AppSecretName string `json:"appSecretName,omitempty"`
	Labels []string `json:"labels,omitempty"`
}
