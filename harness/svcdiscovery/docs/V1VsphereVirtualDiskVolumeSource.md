# V1VsphereVirtualDiskVolumeSource

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**FsType** | **string** | fsType is filesystem type to mount. Must be a filesystem type supported by the host operating system. Ex. \&quot;ext4\&quot;, \&quot;xfs\&quot;, \&quot;ntfs\&quot;. Implicitly inferred to be \&quot;ext4\&quot; if unspecified. +optional | [optional] [default to null]
**StoragePolicyID** | **string** | storagePolicyID is the storage Policy Based Management (SPBM) profile ID associated with the StoragePolicyName. +optional | [optional] [default to null]
**StoragePolicyName** | **string** | storagePolicyName is the storage Policy Based Management (SPBM) profile name. +optional | [optional] [default to null]
**VolumePath** | **string** | volumePath is the path that identifies vSphere volume vmdk | [optional] [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)

