/*
 * Chaos Manager API
 *
 * No description provided (generated by Swagger Codegen https://github.com/swagger-api/swagger-codegen)
 *
 * API version: 1.0
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */
package chaos

type V1PortworxVolumeSource struct {
	// FSType represents the filesystem type to mount Must be a filesystem type supported by the host operating system. Ex. \"ext4\", \"xfs\". Implicitly inferred to be \"ext4\" if unspecified.
	FsType string `json:"fsType,omitempty"`
	// Defaults to false (read/write). ReadOnly here will force the ReadOnly setting in VolumeMounts. +optional
	ReadOnly bool `json:"readOnly,omitempty"`
	// VolumeID uniquely identifies a Portworx volume
	VolumeID string `json:"volumeID,omitempty"`
}
