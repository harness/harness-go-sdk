/*
 * Chaos Manager API
 *
 * No description provided (generated by Swagger Codegen https://github.com/swagger-api/swagger-codegen)
 *
 * API version: 1.0
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */
package chaos

type InfraV2Env struct {
	Key string `json:"key,omitempty"`
	Name string `json:"name,omitempty"`
	Value string `json:"value,omitempty"`
	ValueFrom *InfraV2EnvValueFrom `json:"valueFrom,omitempty"`
}
