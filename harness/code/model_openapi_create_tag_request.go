/*
 * API Specification
 *
 * No description provided (generated by Swagger Codegen https://github.com/swagger-api/swagger-codegen)
 *
 * API version: 0.0.0
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */
package code

type OpenapiCreateTagRequest struct {
	BypassRules bool   `json:"bypass_rules,omitempty"`
	Message     string `json:"message,omitempty"`
	Name        string `json:"name,omitempty"`
	Target      string `json:"target,omitempty"`
}
