/*
 * Service Discovery API
 *
 * No description provided (generated by Swagger Codegen https://github.com/swagger-api/swagger-codegen)
 *
 * API version: 1.0
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */
package svcdiscovery

type V1AzureFileVolumeSource struct {
	// readOnly defaults to false (read/write). ReadOnly here will force the ReadOnly setting in VolumeMounts. +optional
	ReadOnly bool `json:"readOnly,omitempty"`
	// secretName is the  name of secret that contains Azure Storage Account Name and Key
	SecretName string `json:"secretName,omitempty"`
	// shareName is the azure share Name
	ShareName string `json:"shareName,omitempty"`
}
