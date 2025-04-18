/*
 * Chaos Manager API
 *
 * No description provided (generated by Swagger Codegen https://github.com/swagger-api/swagger-codegen)
 *
 * API version: 1.0
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */
package chaos

type V1VsphereVirtualDiskVolumeSource struct {
	// Filesystem type to mount. Must be a filesystem type supported by the host operating system. Ex. \"ext4\", \"xfs\", \"ntfs\". Implicitly inferred to be \"ext4\" if unspecified. +optional
	FsType string `json:"fsType,omitempty"`
	// Storage Policy Based Management (SPBM) profile ID associated with the StoragePolicyName. +optional
	StoragePolicyID string `json:"storagePolicyID,omitempty"`
	// Storage Policy Based Management (SPBM) profile name. +optional
	StoragePolicyName string `json:"storagePolicyName,omitempty"`
	// Path that identifies vSphere volume vmdk
	VolumePath string `json:"volumePath,omitempty"`
}
