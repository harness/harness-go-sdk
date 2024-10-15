# V1SecretVolumeSource

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**DefaultMode** | **int32** | Optional: mode bits used to set permissions on created files by default. Must be an octal value between 0000 and 0777 or a decimal value between 0 and 511. YAML accepts both octal and decimal values, JSON requires decimal values for mode bits. Defaults to 0644. Directories within the path are not affected by this setting. This might be in conflict with other options that affect the file mode, like fsGroup, and the result can be other mode bits set. +optional | [optional] [default to null]
**Items** | [**[]V1KeyToPath**](v1.KeyToPath.md) | If unspecified, each key-value pair in the Data field of the referenced Secret will be projected into the volume as a file whose name is the key and content is the value. If specified, the listed keys will be projected into the specified paths, and unlisted keys will not be present. If a key is specified which is not present in the Secret, the volume setup will error unless it is marked optional. Paths must be relative and may not contain the &#x27;..&#x27; path or start with &#x27;..&#x27;. +optional | [optional] [default to null]
**Optional** | **bool** | Specify whether the Secret or its keys must be defined +optional | [optional] [default to null]
**SecretName** | **string** | Name of the secret in the pod&#x27;s namespace to use. More info: https://kubernetes.io/docs/concepts/storage/volumes#secret +optional | [optional] [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)
