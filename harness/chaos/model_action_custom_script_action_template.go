/*
 * Chaos Manager API
 *
 * No description provided (generated by Swagger Codegen https://github.com/swagger-api/swagger-codegen)
 *
 * API version: 1.0
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */
package chaos

type ActionCustomScriptActionTemplate struct {
	Args []string `json:"args,omitempty"`
	Command string `json:"command,omitempty"`
	Env *interface{} `json:"env,omitempty"`
}
