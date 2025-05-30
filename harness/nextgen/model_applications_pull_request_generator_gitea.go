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

// PullRequestGeneratorGitea defines connection info specific to Gitea.
type ApplicationsPullRequestGeneratorGitea struct {
	// Gitea org or user to scan. Required.
	Owner string `json:"owner,omitempty"`
	// Gitea repo name to scan. Required.
	Repo string `json:"repo,omitempty"`
	Api string `json:"api,omitempty"`
	TokenRef *ApplicationsSecretRef `json:"tokenRef,omitempty"`
	// Allow insecure tls, for self-signed certificates; default: false.
	Insecure bool `json:"insecure,omitempty"`
}
