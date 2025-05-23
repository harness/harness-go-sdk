/*
 * Chaos Manager API
 *
 * No description provided (generated by Swagger Codegen https://github.com/swagger-api/swagger-codegen)
 *
 * API version: 1.0
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */
package chaos

type ImageRegistryImageRegistryV2 struct {
	CreatedAt string `json:"createdAt,omitempty"`
	CreatedBy *ImageRegistryUserDetails `json:"createdBy,omitempty"`
	CustomImages *ImageRegistryCustomImagesRequest `json:"customImages,omitempty"`
	Identifier *GithubComHarnessHceSaasHceSdkTypesApiK8sifsImageRegistryScopedIdentifiers `json:"identifier,omitempty"`
	InfraID string `json:"infraID,omitempty"`
	IsDefault bool `json:"isDefault,omitempty"`
	IsOverrideAllowed bool `json:"isOverrideAllowed,omitempty"`
	IsPrivate bool `json:"isPrivate,omitempty"`
	RegistryAccount string `json:"registryAccount,omitempty"`
	RegistryServer string `json:"registryServer,omitempty"`
	SecretName string `json:"secretName,omitempty"`
	UpdatedAt string `json:"updatedAt,omitempty"`
	UpdatedBy *ImageRegistryUserDetails `json:"updatedBy,omitempty"`
	UseCustomImages bool `json:"useCustomImages,omitempty"`
}
