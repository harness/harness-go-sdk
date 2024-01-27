# code{{classname}}

All URIs are relative to */api/v1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**AdminCreateUser**](AdminApi.md#AdminCreateUser) | **Post** /admin/users | 
[**AdminDeleteUser**](AdminApi.md#AdminDeleteUser) | **Delete** /admin/users/{user_uid} | 
[**AdminGetUser**](AdminApi.md#AdminGetUser) | **Get** /admin/users/{user_uid} | 
[**AdminListUsers**](AdminApi.md#AdminListUsers) | **Get** /admin/users | 
[**AdminUpdateUser**](AdminApi.md#AdminUpdateUser) | **Patch** /admin/users/{user_uid} | 
[**UpdateUserAdmin**](AdminApi.md#UpdateUserAdmin) | **Patch** /admin/users/{user_uid}/admin | 

# **AdminCreateUser**
> TypesUser AdminCreateUser(ctx, optional)


### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***AdminApiAdminCreateUserOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a AdminApiAdminCreateUserOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **body** | [**optional.Interface of OpenapiAdminUsersCreateRequest**](OpenapiAdminUsersCreateRequest.md)|  | 

### Return type

[**TypesUser**](TypesUser.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **AdminDeleteUser**
> AdminDeleteUser(ctx, userUid)


### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **userUid** | **string**|  | 

### Return type

 (empty response body)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **AdminGetUser**
> TypesUser AdminGetUser(ctx, userUid)


### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **userUid** | **string**|  | 

### Return type

[**TypesUser**](TypesUser.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **AdminListUsers**
> []TypesUser AdminListUsers(ctx, optional)


### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***AdminApiAdminListUsersOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a AdminApiAdminListUsersOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **sort** | **optional.String**|  | 
 **order** | **optional.String**|  | 
 **page** | **optional.Int32**|  | [default to 1]
 **limit** | **optional.Int32**|  | [default to 30]

### Return type

[**[]TypesUser**](TypesUser.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **AdminUpdateUser**
> TypesUser AdminUpdateUser(ctx, userUid, optional)


### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **userUid** | **string**|  | 
 **optional** | ***AdminApiAdminUpdateUserOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a AdminApiAdminUpdateUserOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **body** | [**optional.Interface of OpenapiAdminUsersUpdateRequest**](OpenapiAdminUsersUpdateRequest.md)|  | 

### Return type

[**TypesUser**](TypesUser.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **UpdateUserAdmin**
> TypesUser UpdateUserAdmin(ctx, userUid, optional)


### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **userUid** | **string**|  | 
 **optional** | ***AdminApiUpdateUserAdminOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a AdminApiUpdateUserAdminOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **body** | [**optional.Interface of OpenapiUpdateAdminRequest**](OpenapiUpdateAdminRequest.md)|  | 

### Return type

[**TypesUser**](TypesUser.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

