# AlertSearchRequest

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Page** | **int32** | Page number for pagination (1-indexed) | [optional] [default to 1]
**Limit** | **int32** | Number of records per page | [optional] [default to 10]
**Sort** | [***AlertSearchRequestSort**](AlertSearchRequest_sort.md) |  | [optional] [default to null]
**Filters** | [**[]AlertSearchRequestFilters**](AlertSearchRequest_filters.md) | Advanced filters for querying alerts | [optional] [default to null]
**EntityType** | **string** | Filter by entity type | [optional] [default to null]
**EntityId** | **[]string** | Filter by specific entity IDs (AutoStopping rule IDs) | [optional] [default to null]
**Enabled** | **bool** | Filter by enabled status | [optional] [default to null]
**Events** | **[]string** | Filter by event types | [optional] [default to null]
**Query** | **string** | Search query to filter alerts by name (case-insensitive partial match) | [optional] [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)

