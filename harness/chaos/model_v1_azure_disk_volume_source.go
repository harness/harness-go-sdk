/*
 * Chaos Manager API
 *
 * No description provided (generated by Swagger Codegen https://github.com/swagger-api/swagger-codegen)
 *
 * API version: 1.0
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */
package chaos

type V1AzureDiskVolumeSource struct {
	// Host Caching mode: None, Read Only, Read Write. +optional
	CachingMode *AllOfv1AzureDiskVolumeSourceCachingMode `json:"cachingMode,omitempty"`
	// The Name of the data disk in the blob storage
	DiskName string `json:"diskName,omitempty"`
	// The URI the data disk in the blob storage
	DiskURI string `json:"diskURI,omitempty"`
	// Filesystem type to mount. Must be a filesystem type supported by the host operating system. Ex. \"ext4\", \"xfs\", \"ntfs\". Implicitly inferred to be \"ext4\" if unspecified. +optional
	FsType string `json:"fsType,omitempty"`
	// Expected values Shared: multiple blob disks per storage account  Dedicated: single blob disk per storage account  Managed: azure managed data disk (only in managed availability set). defaults to shared
	Kind *AllOfv1AzureDiskVolumeSourceKind `json:"kind,omitempty"`
	// Defaults to false (read/write). ReadOnly here will force the ReadOnly setting in VolumeMounts. +optional
	ReadOnly bool `json:"readOnly,omitempty"`
}
