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

// ECRAuthorizationTokenSpec represents externalSecret for ECR External Secret Operator generator.
type RepositoriesEcrAuthorizationTokenGenerator struct {
	Region string `json:"region,omitempty"`
	SecretRef *RepositoriesAwsSecretRef `json:"secretRef,omitempty"`
	JwtAuth *RepositoriesServiceAccountSelector `json:"jwtAuth,omitempty"`
	Role string `json:"role,omitempty"`
}
