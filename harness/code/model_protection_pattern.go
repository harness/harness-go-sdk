/*
 * API Specification
 *
 * No description provided (generated by Swagger Codegen https://github.com/swagger-api/swagger-codegen)
 *
 * API version: 0.0.0
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */
package code

type ProtectionPattern struct {
	Default_ bool     `json:"default,omitempty"`
	Exclude  []string `json:"exclude,omitempty"`
	Include  []string `json:"include,omitempty"`
}
