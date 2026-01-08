# {{classname}}

All URIs are relative to */gateway/v1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**InfrastructureGet**](InfrastructureApi.md#InfrastructureGet) | **Get** /api/infrastructures/{id} | Get Infrastructure

# **InfrastructureGet**
> InfrastructureOutput InfrastructureGet(ctx, id, harnessAccount, optional)
Get Infrastructure

Endpoint for retrieving a specific infrastructure.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **id** | **string**| ID of the infrastructure | 
  **harnessAccount** | **string**| Account is the internal customer account ID. | 
 **optional** | ***InfrastructureApiInfrastructureGetOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a InfrastructureApiInfrastructureGetOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **orgIdentifier** | **optional.String**| Org is the organization identifier. | 
 **projectIdentifier** | **optional.String**| Project is the project identifier. | 

### Return type

[**InfrastructureOutput**](InfrastructureOutput.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

