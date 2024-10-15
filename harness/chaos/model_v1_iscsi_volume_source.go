/*
 * No description provided (generated by Swagger Codegen https://github.com/swagger-api/swagger-codegen)
 *
 * API version: 1.0.0
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */
package swagger

type V1IscsiVolumeSource struct {
	// whether support iSCSI Discovery CHAP authentication +optional
	ChapAuthDiscovery bool `json:"chapAuthDiscovery,omitempty"`
	// whether support iSCSI Session CHAP authentication +optional
	ChapAuthSession bool `json:"chapAuthSession,omitempty"`
	// Filesystem type of the volume that you want to mount. Tip: Ensure that the filesystem type is supported by the host operating system. Examples: \"ext4\", \"xfs\", \"ntfs\". Implicitly inferred to be \"ext4\" if unspecified. More info: https://kubernetes.io/docs/concepts/storage/volumes#iscsi TODO: how do we prevent errors in the filesystem from compromising the machine +optional
	FsType string `json:"fsType,omitempty"`
	// Custom iSCSI Initiator Name. If initiatorName is specified with iscsiInterface simultaneously, new iSCSI interface <target portal>:<volume name> will be created for the connection. +optional
	InitiatorName string `json:"initiatorName,omitempty"`
	// Target iSCSI Qualified Name.
	Iqn string `json:"iqn,omitempty"`
	// iSCSI Interface Name that uses an iSCSI transport. Defaults to 'default' (tcp). +optional
	IscsiInterface string `json:"iscsiInterface,omitempty"`
	// iSCSI Target Lun number.
	Lun int32 `json:"lun,omitempty"`
	// iSCSI Target Portal List. The portal is either an IP or ip_addr:port if the port is other than default (typically TCP ports 860 and 3260). +optional
	Portals []string `json:"portals,omitempty"`
	// ReadOnly here will force the ReadOnly setting in VolumeMounts. Defaults to false. +optional
	ReadOnly bool `json:"readOnly,omitempty"`
	// CHAP Secret for iSCSI target and initiator authentication +optional
	SecretRef *AllOfv1IscsiVolumeSourceSecretRef `json:"secretRef,omitempty"`
	// iSCSI Target Portal. The Portal is either an IP or ip_addr:port if the port is other than default (typically TCP ports 860 and 3260).
	TargetPortal string `json:"targetPortal,omitempty"`
}