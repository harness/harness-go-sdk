# V1StorageOsVolumeSource

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**FsType** | **string** | Filesystem type to mount. Must be a filesystem type supported by the host operating system. Ex. \&quot;ext4\&quot;, \&quot;xfs\&quot;, \&quot;ntfs\&quot;. Implicitly inferred to be \&quot;ext4\&quot; if unspecified. +optional | [optional] [default to null]
**ReadOnly** | **bool** | Defaults to false (read/write). ReadOnly here will force the ReadOnly setting in VolumeMounts. +optional | [optional] [default to null]
**SecretRef** | [***AllOfv1StorageOsVolumeSourceSecretRef**](AllOfv1StorageOsVolumeSourceSecretRef.md) | SecretRef specifies the secret to use for obtaining the StorageOS API credentials.  If not specified, default values will be attempted. +optional | [optional] [default to null]
**VolumeName** | **string** | VolumeName is the human-readable name of the StorageOS volume.  Volume names are only unique within a namespace. | [optional] [default to null]
**VolumeNamespace** | **string** | VolumeNamespace specifies the scope of the volume within StorageOS.  If no namespace is specified then the Pod&#x27;s namespace will be used.  This allows the Kubernetes name scoping to be mirrored within StorageOS for tighter integration. Set VolumeName to any name to override the default behaviour. Set to \&quot;default\&quot; if you are not using namespaces within StorageOS. Namespaces that do not pre-exist within StorageOS will be created. +optional | [optional] [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)

