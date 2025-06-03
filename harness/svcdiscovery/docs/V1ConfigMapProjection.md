# V1ConfigMapProjection

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Items** | [**[]V1KeyToPath**](v1.KeyToPath.md) | items if unspecified, each key-value pair in the Data field of the referenced ConfigMap will be projected into the volume as a file whose name is the key and content is the value. If specified, the listed keys will be projected into the specified paths, and unlisted keys will not be present. If a key is specified which is not present in the ConfigMap, the volume setup will error unless it is marked optional. Paths must be relative and may not contain the &#x27;..&#x27; path or start with &#x27;..&#x27;. +optional | [optional] [default to null]
**Name** | **string** | Name of the referent. More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/names/#names TODO: Add other useful fields. apiVersion, kind, uid? +optional | [optional] [default to null]
**Optional** | **bool** | optional specify whether the ConfigMap or its keys must be defined +optional | [optional] [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)

