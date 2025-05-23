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

// SCMProviderGeneratorAWSCodeCommit defines connection info specific to AWS CodeCommit.
type ApplicationsScmProviderGeneratorAwsCodeCommit struct {
	TagFilters []ApplicationsTagFilter `json:"tagFilters,omitempty"`
	// Role provides the AWS IAM role to assume, for cross-account repo discovery if not provided, AppSet controller will use its pod/node identity to discover.
	Role string `json:"role,omitempty"`
	// Region provides the AWS region to discover repos. if not provided, AppSet controller will infer the current region from environment.
	Region string `json:"region,omitempty"`
	// Scan all branches instead of just the default branch.
	AllBranches bool `json:"allBranches,omitempty"`
}
