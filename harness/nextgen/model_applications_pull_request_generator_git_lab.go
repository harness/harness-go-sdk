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

// PullRequestGeneratorGitLab defines connection info specific to GitLab.
type ApplicationsPullRequestGeneratorGitLab struct {
	// GitLab project to scan. Required.
	Project string `json:"project,omitempty"`
	// The GitLab API URL to talk to. If blank, uses https://gitlab.com/.
	Api string `json:"api,omitempty"`
	TokenRef *ApplicationsSecretRef `json:"tokenRef,omitempty"`
	Labels []string `json:"labels,omitempty"`
	PullRequestState string `json:"pullRequestState,omitempty"`
	Insecure bool `json:"insecure,omitempty"`
	CaRef *ApplicationsConfigMapKeyRef `json:"caRef,omitempty"`
}
