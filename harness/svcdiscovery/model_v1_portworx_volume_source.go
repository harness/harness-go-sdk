/*
 * Service Discovery API
 *
 * No description provided (generated by Swagger Codegen https://github.com/swagger-api/swagger-codegen)
 *
 * API version: 1.0
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */
package svcdiscovery

type V1PortworxVolumeSource struct {
	// fSType represents the filesystem type to mount Must be a filesystem type supported by the host operating system. Ex. \"ext4\", \"xfs\". Implicitly inferred to be \"ext4\" if unspecified.
	FsType string `json:"fsType,omitempty"`
	// readOnly defaults to false (read/write). ReadOnly here will force the ReadOnly setting in VolumeMounts. +optional
	ReadOnly bool `json:"readOnly,omitempty"`
	// volumeID uniquely identifies a Portworx volume
	VolumeID string `json:"volumeID,omitempty"`
}
