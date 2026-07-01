# {{classname}}

All URIs are relative to */gateway/v1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateEntity**](EntitiesApi.md#CreateEntity) | **Post** /v1/entities | Create a Entity
[**CreateEntityVersion**](EntitiesApi.md#CreateEntityVersion) | **Post** /v1/entities/versions | Create an EntityVersion
[**DeleteEntity**](EntitiesApi.md#DeleteEntity) | **Delete** /v1/entities/{scope}/{kind}/{identifier} | Delete a Entity
[**DeleteEntityVersion**](EntitiesApi.md#DeleteEntityVersion) | **Delete** /v1/entities/{scope}/{kind}/{identifier}/versions/{version} | Delete an EntityVersion
[**GetEntities**](EntitiesApi.md#GetEntities) | **Get** /v1/entities | Get Entities
[**GetEntity**](EntitiesApi.md#GetEntity) | **Get** /v1/entities/{scope}/{kind}/{identifier} | Get Entity
[**GetEntityVersion**](EntitiesApi.md#GetEntityVersion) | **Get** /v1/entities/{scope}/{kind}/{identifier}/versions/{version} | Get Entity Version
[**GetEntityVersions**](EntitiesApi.md#GetEntityVersions) | **Get** /v1/entities/{scope}/{kind}/{identifier}/versions | Get Entity Versions
[**ImportEntity**](EntitiesApi.md#ImportEntity) | **Post** /v1/entities/import | Import an Entity
[**MoveEntity**](EntitiesApi.md#MoveEntity) | **Post** /v1/entities/move/{scope}/{kind}/{identifier} | Move an Entity
[**UpdateEntity**](EntitiesApi.md#UpdateEntity) | **Put** /v1/entities/{scope}/{kind}/{identifier} | Update a Entity
[**UpdateEntityVersion**](EntitiesApi.md#UpdateEntityVersion) | **Put** /v1/entities/{scope}/{kind}/{identifier}/versions/{version} | Update an EntityVersion
[**UpdateGitMetadata**](EntitiesApi.md#UpdateGitMetadata) | **Put** /v1/entities/git-metadata/{scope}/{kind}/{identifier} | Update GitMetadata for Remote Entities

# **CreateEntity**
> EntityResponse CreateEntity(ctx, body, optional)
Create a Entity

Create an Entity.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**EntityCreateRequest**](EntityCreateRequest.md)| Entity Create Request Body | 
 **optional** | ***EntitiesApiCreateEntityOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a EntitiesApiCreateEntityOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **orgIdentifier** | **optional.**| Organization identifier | 
 **projectIdentifier** | **optional.**| Project identifier | 
 **harnessAccount** | **optional.**| Identifier field of the account the resource is scoped to. This is required for Authorization methods other than the x-api-key header. If you are using the x-api-key header, this can be skipped. | 
 **convert** | **optional.**| Converts Backstage style YAML to Harness entity YAML | [default to false]
 **dryRun** | **optional.**| Dry run validation of entity creation | [default to false]

### Return type

[**EntityResponse**](EntityResponse.md)

### Authorization

[x-api-key](../README.md#x-api-key)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json, application/yaml

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **CreateEntityVersion**
> EntityVersionResponse CreateEntityVersion(ctx, body, optional)
Create an EntityVersion

Create an EntityVersion.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**EntityVersionCreateRequest**](EntityVersionCreateRequest.md)| Entity Version Create Request Body | 
 **optional** | ***EntitiesApiCreateEntityVersionOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a EntitiesApiCreateEntityVersionOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **orgIdentifier** | **optional.**| Organization identifier | 
 **projectIdentifier** | **optional.**| Project identifier | 
 **harnessAccount** | **optional.**| Identifier field of the account the resource is scoped to. This is required for Authorization methods other than the x-api-key header. If you are using the x-api-key header, this can be skipped. | 

### Return type

[**EntityVersionResponse**](EntityVersionResponse.md)

### Authorization

[x-api-key](../README.md#x-api-key)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json, application/yaml

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **DeleteEntity**
> DeleteEntity(ctx, scope, kind, identifier, optional)
Delete a Entity

Delete an Entity.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **scope** | **string**| Scope | 
  **kind** | **string**| Kind | 
  **identifier** | **string**| Identifier | 
 **optional** | ***EntitiesApiDeleteEntityOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a EntitiesApiDeleteEntityOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **orgIdentifier** | **optional.String**| Organization identifier | 
 **projectIdentifier** | **optional.String**| Project identifier | 
 **harnessAccount** | **optional.String**| Identifier field of the account the resource is scoped to. This is required for Authorization methods other than the x-api-key header. If you are using the x-api-key header, this can be skipped. | 

### Return type

 (empty response body)

### Authorization

[x-api-key](../README.md#x-api-key)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **DeleteEntityVersion**
> DeleteEntityVersion(ctx, scope, kind, identifier, version, optional)
Delete an EntityVersion

Delete an EntityVersion

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **scope** | **string**| Scope | 
  **kind** | **string**| Kind | 
  **identifier** | **string**| Identifier | 
  **version** | **string**| version | 
 **optional** | ***EntitiesApiDeleteEntityVersionOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a EntitiesApiDeleteEntityVersionOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------




 **harnessAccount** | **optional.String**| Identifier field of the account the resource is scoped to. This is required for Authorization methods other than the x-api-key header. If you are using the x-api-key header, this can be skipped. | 
 **orgIdentifier** | **optional.String**| Organization identifier | 
 **projectIdentifier** | **optional.String**| Project identifier | 

### Return type

 (empty response body)

### Authorization

[x-api-key](../README.md#x-api-key)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetEntities**
> []EntityResponse GetEntities(ctx, optional)
Get Entities

Get Entities.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***EntitiesApiGetEntitiesOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a EntitiesApiGetEntitiesOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **harnessAccount** | **optional.String**| Identifier field of the account the resource is scoped to. This is required for Authorization methods other than the x-api-key header. If you are using the x-api-key header, this can be skipped. | 
 **page** | **optional.Int32**| Pagination page number strategy: Specify the page number within the paginated collection related to the number of items in each page  | 
 **limit** | **optional.Int32**| Pagination: Number of items to return | [default to 10]
 **sort** | **optional.String**| Parameter on the basis of which sorting is done. | 
 **searchTerm** | **optional.String**| This would be used to filter resources having attributes matching the search term. | 
 **resolvePlaceholders** | **optional.Bool**| This is used to return entities response with resolved placeholders | [default to false]
 **scopes** | **optional.String**| Filter entities on the scopes | 
 **entityRefs** | **optional.String**| Entity refs on which the Entities are filtered. | 
 **ownedByMe** | **optional.Bool**| Filter entities owned by the user and the groups that the user is part of. | 
 **favorites** | **optional.Bool**| Filter entities that are marked as favorites for the user. | 
 **kind** | **optional.String**| Kinds on which the Entities are filtered. | 
 **type_** | **optional.String**| Types on which the Entities are filtered. | 
 **owner** | **optional.String**| Owners on which the Entities are filtered. | 
 **lifecycle** | **optional.String**| Lifecycles on which the Entities are filtered. | 
 **tags** | **optional.String**| Tags on which the Entities are filtered. | 
 **filter** | **optional.String**| Filter sets that get matched against each entity | 

### Return type

[**[]EntityResponse**](EntityResponse.md)

### Authorization

[x-api-key](../README.md#x-api-key)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json, application/yaml

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetEntity**
> EntityResponse GetEntity(ctx, scope, kind, identifier, optional)
Get Entity

Get Entity.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **scope** | **string**| Scope | 
  **kind** | **string**| Kind | 
  **identifier** | **string**| Identifier | 
 **optional** | ***EntitiesApiGetEntityOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a EntitiesApiGetEntityOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **orgIdentifier** | **optional.String**| Organization identifier | 
 **projectIdentifier** | **optional.String**| Project identifier | 
 **harnessAccount** | **optional.String**| Identifier field of the account the resource is scoped to. This is required for Authorization methods other than the x-api-key header. If you are using the x-api-key header, this can be skipped. | 
 **branchName** | **optional.String**| Name of the branch (for Git Experience). | 
 **connectorRef** | **optional.String**| Identifier of the Harness Connector used for CRUD operations on the Entity (for Git Experience). | 
 **repoName** | **optional.String**| Name of the repository (for Git Experience). | 
 **loadFromCache** | **optional.String**| Flag to enable loading the remote pipeline from git or git cache | [default to false]
 **loadFromFallbackBranch** | **optional.Bool**| Flag to load the pipeline from the created non default branch | [default to false]
 **resolvePlaceholders** | **optional.Bool**| This is used to return entities response with resolved placeholders | [default to false]

### Return type

[**EntityResponse**](EntityResponse.md)

### Authorization

[x-api-key](../README.md#x-api-key)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json, application/yaml

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetEntityVersion**
> EntityVersionResponse GetEntityVersion(ctx, scope, kind, identifier, version, optional)
Get Entity Version

Get Entity Version.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **scope** | **string**| Scope | 
  **kind** | **string**| Kind | 
  **identifier** | **string**| Identifier | 
  **version** | **string**| version | 
 **optional** | ***EntitiesApiGetEntityVersionOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a EntitiesApiGetEntityVersionOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------




 **harnessAccount** | **optional.String**| Identifier field of the account the resource is scoped to. This is required for Authorization methods other than the x-api-key header. If you are using the x-api-key header, this can be skipped. | 
 **orgIdentifier** | **optional.String**| Organization identifier | 
 **projectIdentifier** | **optional.String**| Project identifier | 

### Return type

[**EntityVersionResponse**](EntityVersionResponse.md)

### Authorization

[x-api-key](../README.md#x-api-key)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json, application/yaml

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetEntityVersions**
> []EntityVersionResponse GetEntityVersions(ctx, scope, kind, identifier, optional)
Get Entity Versions

Get Entity Versions.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **scope** | **string**| Scope | 
  **kind** | **string**| Kind | 
  **identifier** | **string**| Identifier | 
 **optional** | ***EntitiesApiGetEntityVersionsOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a EntitiesApiGetEntityVersionsOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **orgIdentifier** | **optional.String**| Organization identifier | 
 **projectIdentifier** | **optional.String**| Project identifier | 
 **harnessAccount** | **optional.String**| Identifier field of the account the resource is scoped to. This is required for Authorization methods other than the x-api-key header. If you are using the x-api-key header, this can be skipped. | 
 **page** | **optional.Int32**| Pagination page number strategy: Specify the page number within the paginated collection related to the number of items in each page  | 
 **limit** | **optional.Int32**| Pagination: Number of items to return | [default to 10]
 **searchTerm** | **optional.String**| This would be used to filter resources having attributes matching the search term. | 
 **deprecated** | **optional.Bool**| Filter versions that are deprecated or not. | 

### Return type

[**[]EntityVersionResponse**](EntityVersionResponse.md)

### Authorization

[x-api-key](../README.md#x-api-key)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json, application/yaml

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ImportEntity**
> EntityResponse ImportEntity(ctx, body, optional)
Import an Entity

Import an Entity.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**GitImportDetails**](GitImportDetails.md)| Git Import Request Body | 
 **optional** | ***EntitiesApiImportEntityOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a EntitiesApiImportEntityOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **orgIdentifier** | **optional.**| Organization identifier | 
 **projectIdentifier** | **optional.**| Project identifier | 
 **harnessAccount** | **optional.**| Identifier field of the account the resource is scoped to. This is required for Authorization methods other than the x-api-key header. If you are using the x-api-key header, this can be skipped. | 

### Return type

[**EntityResponse**](EntityResponse.md)

### Authorization

[x-api-key](../README.md#x-api-key)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json, application/yaml

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **MoveEntity**
> MoveEntity(ctx, body, scope, kind, identifier, optional)
Move an Entity

Move an Entity.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**EntityMoveRequest**](EntityMoveRequest.md)| Entity Move Request Body | 
  **scope** | **string**| Scope | 
  **kind** | **string**| Kind | 
  **identifier** | **string**| Identifier | 
 **optional** | ***EntitiesApiMoveEntityOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a EntitiesApiMoveEntityOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------




 **orgIdentifier** | **optional.**| Organization identifier | 
 **projectIdentifier** | **optional.**| Project identifier | 
 **harnessAccount** | **optional.**| Identifier field of the account the resource is scoped to. This is required for Authorization methods other than the x-api-key header. If you are using the x-api-key header, this can be skipped. | 

### Return type

 (empty response body)

### Authorization

[x-api-key](../README.md#x-api-key)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **UpdateEntity**
> EntityResponse UpdateEntity(ctx, body, scope, kind, identifier, optional)
Update a Entity

Update an Entity.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**EntityUpdateRequest**](EntityUpdateRequest.md)| Entity Update Request Body | 
  **scope** | **string**| Scope | 
  **kind** | **string**| Kind | 
  **identifier** | **string**| Identifier | 
 **optional** | ***EntitiesApiUpdateEntityOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a EntitiesApiUpdateEntityOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------




 **orgIdentifier** | **optional.**| Organization identifier | 
 **projectIdentifier** | **optional.**| Project identifier | 
 **harnessAccount** | **optional.**| Identifier field of the account the resource is scoped to. This is required for Authorization methods other than the x-api-key header. If you are using the x-api-key header, this can be skipped. | 

### Return type

[**EntityResponse**](EntityResponse.md)

### Authorization

[x-api-key](../README.md#x-api-key)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json, application/yaml

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **UpdateEntityVersion**
> EntityVersionResponse UpdateEntityVersion(ctx, body, scope, kind, identifier, version, optional)
Update an EntityVersion

Update the mutable fields of an EntityVersion.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**EntityVersionUpdateRequest**](EntityVersionUpdateRequest.md)| Entity Version Update Request Body | 
  **scope** | **string**| Scope | 
  **kind** | **string**| Kind | 
  **identifier** | **string**| Identifier | 
  **version** | **string**| version | 
 **optional** | ***EntitiesApiUpdateEntityVersionOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a EntitiesApiUpdateEntityVersionOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------





 **harnessAccount** | **optional.**| Identifier field of the account the resource is scoped to. This is required for Authorization methods other than the x-api-key header. If you are using the x-api-key header, this can be skipped. | 
 **orgIdentifier** | **optional.**| Organization identifier | 
 **projectIdentifier** | **optional.**| Project identifier | 

### Return type

[**EntityVersionResponse**](EntityVersionResponse.md)

### Authorization

[x-api-key](../README.md#x-api-key)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json, application/yaml

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **UpdateGitMetadata**
> UpdateGitMetadata(ctx, body, scope, kind, identifier, optional)
Update GitMetadata for Remote Entities

Update GitMetadata for Remote Entities

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**GitMetadataUpdateRequest**](GitMetadataUpdateRequest.md)|  | 
  **scope** | **string**| Scope | 
  **kind** | **string**| Kind | 
  **identifier** | **string**| Identifier | 
 **optional** | ***EntitiesApiUpdateGitMetadataOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a EntitiesApiUpdateGitMetadataOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------




 **orgIdentifier** | **optional.**| Organization identifier | 
 **projectIdentifier** | **optional.**| Project identifier | 
 **harnessAccount** | **optional.**| Identifier field of the account the resource is scoped to. This is required for Authorization methods other than the x-api-key header. If you are using the x-api-key header, this can be skipped. | 

### Return type

 (empty response body)

### Authorization

[x-api-key](../README.md#x-api-key)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

