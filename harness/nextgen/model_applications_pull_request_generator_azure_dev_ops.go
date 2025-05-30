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

// PullRequestGeneratorAzureDevOps defines connection info specific to AzureDevOps.
type ApplicationsPullRequestGeneratorAzureDevOps struct {
	// Azure DevOps org to scan. Required.
	Organization string `json:"organization,omitempty"`
	// Azure DevOps project name to scan. Required.
	Project string `json:"project,omitempty"`
	// Azure DevOps repo name to scan. Required.
	Repo string `json:"repo,omitempty"`
	// The Azure DevOps API URL to talk to. If blank, use https://dev.azure.com/.
	Api string `json:"api,omitempty"`
	TokenRef *ApplicationsSecretRef `json:"tokenRef,omitempty"`
	Labels []string `json:"labels,omitempty"`
}
