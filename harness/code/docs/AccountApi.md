# code{{classname}}

All URIs are relative to */api/v1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**OnLogin**](AccountApi.md#OnLogin) | **Post** /login | 
[**OnRegister**](AccountApi.md#OnRegister) | **Post** /register | 
[**OpLogout**](AccountApi.md#OpLogout) | **Post** /logout | 

# **OnLogin**
> TypesTokenResponse OnLogin(ctx, optional)


### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***AccountApiOnLoginOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a AccountApiOnLoginOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **body** | [**optional.Interface of OpenapiLoginRequest**](OpenapiLoginRequest.md)|  | 
 **includeCookie** | **optional.**| If set to true the token is also returned as a cookie. | [default to false]

### Return type

[**TypesTokenResponse**](TypesTokenResponse.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **OnRegister**
> TypesTokenResponse OnRegister(ctx, optional)


### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***AccountApiOnRegisterOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a AccountApiOnRegisterOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **body** | [**optional.Interface of OpenapiRegisterRequest**](OpenapiRegisterRequest.md)|  | 
 **includeCookie** | **optional.**| If set to true the token is also returned as a cookie. | [default to false]

### Return type

[**TypesTokenResponse**](TypesTokenResponse.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **OpLogout**
> OpLogout(ctx, )


### Required Parameters
This endpoint does not need any parameter.

### Return type

 (empty response body)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

