/*
 * Harness Artifact Registry API
 *
 * No description provided (generated by Swagger Codegen https://github.com/swagger-api/swagger-codegen)
 *
 * API version: 1.0.0
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */
package har

// Configuration for Harness Artifact UpstreamProxies
type UpstreamConfig struct {
	Auth     *OneOfUpstreamConfigAuth `json:"auth,omitempty"`
	AuthType *AuthType                `json:"authType"`
	Source   string                   `json:"source,omitempty"`
	Url      string                   `json:"url"`
}
