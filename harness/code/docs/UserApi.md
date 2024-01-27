# code{{classname}}

All URIs are relative to */api/v1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateToken**](UserApi.md#CreateToken) | **Post** /user/token | 
[**GetUser**](UserApi.md#GetUser) | **Get** /user | 
[**MembershipSpaces**](UserApi.md#MembershipSpaces) | **Get** /user/memberships | 
[**UpdateUser**](UserApi.md#UpdateUser) | **Patch** /user | 

# **CreateToken**
> TypesTokenResponse CreateToken(ctx, optional)


### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***UserApiCreateTokenOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a UserApiCreateTokenOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **body** | [**optional.Interface of OpenapiCreateTokenRequest**](OpenapiCreateTokenRequest.md)|  | 

### Return type

[**TypesTokenResponse**](TypesTokenResponse.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetUser**
> TypesUser GetUser(ctx, )


### Required Parameters
This endpoint does not need any parameter.

### Return type

[**TypesUser**](TypesUser.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **MembershipSpaces**
> []TypesMembershipSpace MembershipSpaces(ctx, optional)


### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***UserApiMembershipSpacesOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a UserApiMembershipSpacesOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **query** | **optional.String**| The substring by which the spaces the users is a member of are filtered. | 
 **order** | **optional.String**| The order of the output. | [default to asc]
 **sort** | **optional.String**| The field by which the spaces the user is a member of are sorted. | [default to uid]
 **page** | **optional.Int32**| The page to return. | [default to 1]
 **limit** | **optional.Int32**| The maximum number of results to return. | [default to 30]

### Return type

[**[]TypesMembershipSpace**](TypesMembershipSpace.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **UpdateUser**
> TypesUser UpdateUser(ctx, optional)


### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***UserApiUpdateUserOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a UserApiUpdateUserOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **body** | [**optional.Interface of UserUpdateInput**](UserUpdateInput.md)|  | 

### Return type

[**TypesUser**](TypesUser.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

