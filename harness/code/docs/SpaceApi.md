# code{{classname}}

All URIs are relative to */api/v1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateSpace**](SpaceApi.md#CreateSpace) | **Post** /spaces | 
[**DeleteSpace**](SpaceApi.md#DeleteSpace) | **Delete** /spaces/{space_ref} | 
[**ExportProgressSpace**](SpaceApi.md#ExportProgressSpace) | **Get** /spaces/{space_ref}/export-progress | 
[**ExportSpace**](SpaceApi.md#ExportSpace) | **Post** /spaces/{space_ref}/export | 
[**GetSpace**](SpaceApi.md#GetSpace) | **Get** /spaces/{space_ref} | 
[**ImportSpace**](SpaceApi.md#ImportSpace) | **Post** /spaces/import | 
[**ListConnectors**](SpaceApi.md#ListConnectors) | **Get** /spaces/{space_ref}/connectors | 
[**ListRepos**](SpaceApi.md#ListRepos) | **Get** /spaces/{space_ref}/repos | 
[**ListSecrets**](SpaceApi.md#ListSecrets) | **Get** /spaces/{space_ref}/secrets | 
[**ListServiceAccounts**](SpaceApi.md#ListServiceAccounts) | **Get** /spaces/{space_ref}/service-accounts | 
[**ListSpaces**](SpaceApi.md#ListSpaces) | **Get** /spaces/{space_ref}/spaces | 
[**ListTemplates**](SpaceApi.md#ListTemplates) | **Get** /spaces/{space_ref}/templates | 
[**MembershipAdd**](SpaceApi.md#MembershipAdd) | **Post** /spaces/{space_ref}/members | 
[**MembershipDelete**](SpaceApi.md#MembershipDelete) | **Delete** /spaces/{space_ref}/members/{user_uid} | 
[**MembershipList**](SpaceApi.md#MembershipList) | **Get** /spaces/{space_ref}/members | 
[**MembershipUpdate**](SpaceApi.md#MembershipUpdate) | **Patch** /spaces/{space_ref}/members/{user_uid} | 
[**MoveSpace**](SpaceApi.md#MoveSpace) | **Post** /spaces/{space_ref}/move | 
[**UpdateSpace**](SpaceApi.md#UpdateSpace) | **Patch** /spaces/{space_ref} | 

# **CreateSpace**
> TypesSpace CreateSpace(ctx, optional)


### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***SpaceApiCreateSpaceOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a SpaceApiCreateSpaceOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **body** | [**optional.Interface of OpenapiCreateSpaceRequest**](OpenapiCreateSpaceRequest.md)|  | 

### Return type

[**TypesSpace**](TypesSpace.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **DeleteSpace**
> DeleteSpace(ctx, spaceRef)


### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **spaceRef** | **string**|  | 

### Return type

 (empty response body)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ExportProgressSpace**
> SpaceExportProgressOutput ExportProgressSpace(ctx, spaceRef)


### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **spaceRef** | **string**|  | 

### Return type

[**SpaceExportProgressOutput**](SpaceExportProgressOutput.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ExportSpace**
> ExportSpace(ctx, spaceRef, optional)


### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **spaceRef** | **string**|  | 
 **optional** | ***SpaceApiExportSpaceOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a SpaceApiExportSpaceOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **body** | [**optional.Interface of OpenapiExportSpaceRequest**](OpenapiExportSpaceRequest.md)|  | 

### Return type

 (empty response body)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetSpace**
> TypesSpace GetSpace(ctx, spaceRef)


### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **spaceRef** | **string**|  | 

### Return type

[**TypesSpace**](TypesSpace.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ImportSpace**
> TypesSpace ImportSpace(ctx, optional)


### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***SpaceApiImportSpaceOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a SpaceApiImportSpaceOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **body** | [**optional.Interface of SpacesImportBody**](SpacesImportBody.md)|  | 

### Return type

[**TypesSpace**](TypesSpace.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ListConnectors**
> []TypesConnector ListConnectors(ctx, spaceRef, optional)


### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **spaceRef** | **string**|  | 
 **optional** | ***SpaceApiListConnectorsOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a SpaceApiListConnectorsOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **query** | **optional.String**| The substring which is used to filter the repositories by their path name. | 
 **page** | **optional.Int32**| The page to return. | [default to 1]
 **limit** | **optional.Int32**| The maximum number of results to return. | [default to 30]

### Return type

[**[]TypesConnector**](TypesConnector.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ListRepos**
> []TypesRepository ListRepos(ctx, spaceRef, optional)


### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **spaceRef** | **string**|  | 
 **optional** | ***SpaceApiListReposOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a SpaceApiListReposOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **query** | **optional.String**| The substring which is used to filter the repositories by their path name. | 
 **sort** | **optional.String**| The data by which the repositories are sorted. | [default to uid]
 **order** | **optional.String**| The order of the output. | [default to asc]
 **page** | **optional.Int32**| The page to return. | [default to 1]
 **limit** | **optional.Int32**| The maximum number of results to return. | [default to 30]

### Return type

[**[]TypesRepository**](TypesRepository.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ListSecrets**
> []TypesSecret ListSecrets(ctx, spaceRef, optional)


### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **spaceRef** | **string**|  | 
 **optional** | ***SpaceApiListSecretsOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a SpaceApiListSecretsOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **query** | **optional.String**| The substring which is used to filter the repositories by their path name. | 
 **page** | **optional.Int32**| The page to return. | [default to 1]
 **limit** | **optional.Int32**| The maximum number of results to return. | [default to 30]

### Return type

[**[]TypesSecret**](TypesSecret.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ListServiceAccounts**
> []TypesServiceAccount ListServiceAccounts(ctx, spaceRef)


### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **spaceRef** | **string**|  | 

### Return type

[**[]TypesServiceAccount**](TypesServiceAccount.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ListSpaces**
> []TypesSpace ListSpaces(ctx, spaceRef, optional)


### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **spaceRef** | **string**|  | 
 **optional** | ***SpaceApiListSpacesOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a SpaceApiListSpacesOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **query** | **optional.String**| The substring which is used to filter the spaces by their path name. | 
 **sort** | **optional.String**| The data by which the spaces are sorted. | [default to uid]
 **order** | **optional.String**| The order of the output. | [default to asc]
 **page** | **optional.Int32**| The page to return. | [default to 1]
 **limit** | **optional.Int32**| The maximum number of results to return. | [default to 30]

### Return type

[**[]TypesSpace**](TypesSpace.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ListTemplates**
> []TypesTemplate ListTemplates(ctx, spaceRef, optional)


### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **spaceRef** | **string**|  | 
 **optional** | ***SpaceApiListTemplatesOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a SpaceApiListTemplatesOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **query** | **optional.String**| The substring which is used to filter the repositories by their path name. | 
 **page** | **optional.Int32**| The page to return. | [default to 1]
 **limit** | **optional.Int32**| The maximum number of results to return. | [default to 30]

### Return type

[**[]TypesTemplate**](TypesTemplate.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **MembershipAdd**
> TypesMembershipUser MembershipAdd(ctx, spaceRef, optional)


### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **spaceRef** | **string**|  | 
 **optional** | ***SpaceApiMembershipAddOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a SpaceApiMembershipAddOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **body** | [**optional.Interface of SpaceRefMembersBody**](SpaceRefMembersBody.md)|  | 

### Return type

[**TypesMembershipUser**](TypesMembershipUser.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **MembershipDelete**
> MembershipDelete(ctx, spaceRef, userUid)


### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **spaceRef** | **string**|  | 
  **userUid** | **string**|  | 

### Return type

 (empty response body)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **MembershipList**
> []TypesMembershipUser MembershipList(ctx, spaceRef, optional)


### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **spaceRef** | **string**|  | 
 **optional** | ***SpaceApiMembershipListOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a SpaceApiMembershipListOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **query** | **optional.String**| The substring by which the space members are filtered. | 
 **order** | **optional.String**| The order of the output. | [default to asc]
 **sort** | **optional.String**| The field by which the space members are sorted. | [default to name]
 **page** | **optional.Int32**| The page to return. | [default to 1]
 **limit** | **optional.Int32**| The maximum number of results to return. | [default to 30]

### Return type

[**[]TypesMembershipUser**](TypesMembershipUser.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **MembershipUpdate**
> TypesMembershipUser MembershipUpdate(ctx, spaceRef, userUid, optional)


### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **spaceRef** | **string**|  | 
  **userUid** | **string**|  | 
 **optional** | ***SpaceApiMembershipUpdateOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a SpaceApiMembershipUpdateOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **body** | [**optional.Interface of MembersUserUidBody**](MembersUserUidBody.md)|  | 

### Return type

[**TypesMembershipUser**](TypesMembershipUser.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **MoveSpace**
> TypesSpace MoveSpace(ctx, spaceRef, optional)


### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **spaceRef** | **string**|  | 
 **optional** | ***SpaceApiMoveSpaceOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a SpaceApiMoveSpaceOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **body** | [**optional.Interface of OpenapiMoveSpaceRequest**](OpenapiMoveSpaceRequest.md)|  | 

### Return type

[**TypesSpace**](TypesSpace.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **UpdateSpace**
> TypesSpace UpdateSpace(ctx, spaceRef, optional)


### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **spaceRef** | **string**|  | 
 **optional** | ***SpaceApiUpdateSpaceOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a SpaceApiUpdateSpaceOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **body** | [**optional.Interface of OpenapiUpdateSpaceRequest**](OpenapiUpdateSpaceRequest.md)|  | 

### Return type

[**TypesSpace**](TypesSpace.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

