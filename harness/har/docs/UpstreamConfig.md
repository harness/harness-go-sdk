# UpstreamConfig

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Auth** | [***OneOfUpstreamConfigAuth**](OneOfUpstreamConfigAuth.md) |  | [optional] [default to null]
**AuthType** | [***AuthType**](AuthType.md) |  | [default to null]
**Source** | **string** |  | [optional] [default to null]
**Url** | **string** |  | [optional] [default to null]
**MetadataCacheTTL** | **int64** | Time-to-live in seconds for cached FOUND metadata entries. Honoured only for UPSTREAM registries with package types that support resource cache (currently Maven). Must be between 0 and 604800 (7 days). | [optional] [default to null]
**NegativeCacheTTL** | **int64** | Time-to-live in seconds for cached NOT_FOUND entries. Honoured only for UPSTREAM registries with package types that support resource cache (currently Maven). Must be between 0 and 604800 (7 days). | [optional] [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)

