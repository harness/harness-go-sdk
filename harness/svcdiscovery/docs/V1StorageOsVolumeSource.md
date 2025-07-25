# V1StorageOsVolumeSource

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**FsType** | **string** | fsType is the filesystem type to mount. Must be a filesystem type supported by the host operating system. Ex. \&quot;ext4\&quot;, \&quot;xfs\&quot;, \&quot;ntfs\&quot;. Implicitly inferred to be \&quot;ext4\&quot; if unspecified. +optional | [optional] [default to null]
**ReadOnly** | **bool** | readOnly defaults to false (read/write). ReadOnly here will force the ReadOnly setting in VolumeMounts. +optional | [optional] [default to null]
**SecretRef** | [***V1LocalObjectReference**](v1.LocalObjectReference.md) |  | [optional] [default to null]
**VolumeName** | **string** | volumeName is the human-readable name of the StorageOS volume.  Volume names are only unique within a namespace. | [optional] [default to null]
**VolumeNamespace** | **string** | volumeNamespace specifies the scope of the volume within StorageOS.  If no namespace is specified then the Pod&#x27;s namespace will be used.  This allows the Kubernetes name scoping to be mirrored within StorageOS for tighter integration. Set VolumeName to any name to override the default behaviour. Set to \&quot;default\&quot; if you are not using namespaces within StorageOS. Namespaces that do not pre-exist within StorageOS will be created. +optional | [optional] [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)

