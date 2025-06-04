# V1QuobyteVolumeSource

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Group** | **string** | group to map volume access to Default is no group +optional | [optional] [default to null]
**ReadOnly** | **bool** | readOnly here will force the Quobyte volume to be mounted with read-only permissions. Defaults to false. +optional | [optional] [default to null]
**Registry** | **string** | registry represents a single or multiple Quobyte Registry services specified as a string as host:port pair (multiple entries are separated with commas) which acts as the central registry for volumes | [optional] [default to null]
**Tenant** | **string** | tenant owning the given Quobyte volume in the Backend Used with dynamically provisioned Quobyte volumes, value is set by the plugin +optional | [optional] [default to null]
**User** | **string** | user to map volume access to Defaults to serivceaccount user +optional | [optional] [default to null]
**Volume** | **string** | volume is a string that references an already created Quobyte volume by name. | [optional] [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)

