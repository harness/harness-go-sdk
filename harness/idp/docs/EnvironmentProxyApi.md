# {{classname}}

All URIs are relative to */gateway/v1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateCompileAndExecuteEnvironment**](EnvironmentProxyApi.md#CreateCompileAndExecuteEnvironment) | **Post** /v1/idp-environments | Create Compile And Execute Environment
[**DeleteEnvironment**](EnvironmentProxyApi.md#DeleteEnvironment) | **Delete** /v1/idp-environments/{environment-id} | Delete Environment
[**UpdateCompileAndExecuteEnvironment**](EnvironmentProxyApi.md#UpdateCompileAndExecuteEnvironment) | **Put** /v1/idp-environments/{environment-id} | Update Compile And Execute Environment

# **CreateCompileAndExecuteEnvironment**
> EnvironmentExecuteResponse CreateCompileAndExecuteEnvironment(ctx, body, optional)
Create Compile And Execute Environment

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**EnvironmentProxyCreateRequest**](EnvironmentProxyCreateRequest.md)|  | 
 **optional** | ***EnvironmentProxyApiCreateCompileAndExecuteEnvironmentOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a EnvironmentProxyApiCreateCompileAndExecuteEnvironmentOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **harnessAccount** | **optional.**| Identifier field of the account the resource is scoped to. This is required for Authorization methods other than the x-api-key header. If you are using the x-api-key header, this can be skipped. | 
 **orgIdentifier** | **optional.**| Organization identifier | 
 **projectIdentifier** | **optional.**| Project identifier | 

### Return type

[**EnvironmentExecuteResponse**](EnvironmentExecuteResponse.md)

### Authorization

[x-api-key](../README.md#x-api-key)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json, application/yaml

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **DeleteEnvironment**
> DeleteEnvironment(ctx, environmentId, optional)
Delete Environment

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **environmentId** | **string**|  | 
 **optional** | ***EnvironmentProxyApiDeleteEnvironmentOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a EnvironmentProxyApiDeleteEnvironmentOpts struct
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

# **UpdateCompileAndExecuteEnvironment**
> EnvironmentExecuteResponse UpdateCompileAndExecuteEnvironment(ctx, body, environmentId, optional)
Update Compile And Execute Environment

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**EnvironmentProxyUpdateRequest**](EnvironmentProxyUpdateRequest.md)|  | 
  **environmentId** | **string**|  | 
 **optional** | ***EnvironmentProxyApiUpdateCompileAndExecuteEnvironmentOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a EnvironmentProxyApiUpdateCompileAndExecuteEnvironmentOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **harnessAccount** | **optional.**| Identifier field of the account the resource is scoped to. This is required for Authorization methods other than the x-api-key header. If you are using the x-api-key header, this can be skipped. | 
 **orgIdentifier** | **optional.**| Organization identifier | 
 **projectIdentifier** | **optional.**| Project identifier | 

### Return type

[**EnvironmentExecuteResponse**](EnvironmentExecuteResponse.md)

### Authorization

[x-api-key](../README.md#x-api-key)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json, application/yaml

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

