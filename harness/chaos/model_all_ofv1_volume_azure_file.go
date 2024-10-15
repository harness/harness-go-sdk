/*
 * No description provided (generated by Swagger Codegen https://github.com/swagger-api/swagger-codegen)
 *
 * API version: 1.0.0
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */
package swagger

// AzureFile represents an Azure File Service mount on the host and bind mount to the pod. +optional
type AllOfv1VolumeAzureFile struct {
	// Defaults to false (read/write). ReadOnly here will force the ReadOnly setting in VolumeMounts. +optional
	ReadOnly bool `json:"readOnly,omitempty"`
	// the name of secret that contains Azure Storage Account Name and Key
	SecretName string `json:"secretName,omitempty"`
	// Share Name
	ShareName string `json:"shareName,omitempty"`
}