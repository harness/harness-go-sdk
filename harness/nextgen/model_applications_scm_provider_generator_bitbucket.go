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

// SCMProviderGeneratorBitbucket defines connection info specific to Bitbucket Cloud (API version 2).
type ApplicationsScmProviderGeneratorBitbucket struct {
	// Bitbucket workspace to scan. Required.
	Owner string `json:"owner,omitempty"`
	User string `json:"user,omitempty"`
	AppPasswordRef *ApplicationsSecretRef `json:"appPasswordRef,omitempty"`
	// Scan all branches instead of just the main branch.
	AllBranches bool `json:"allBranches,omitempty"`
}
