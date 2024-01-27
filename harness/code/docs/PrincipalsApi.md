# code{{classname}}

All URIs are relative to */api/v1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**ListPrincipals**](PrincipalsApi.md#ListPrincipals) | **Get** /principals | 

# **ListPrincipals**
> []TypesPrincipalInfo ListPrincipals(ctx, optional)


### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***PrincipalsApiListPrincipalsOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a PrincipalsApiListPrincipalsOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **query** | **optional.String**| The substring by which the principals are filtered. | 
 **page** | **optional.Int32**| The page to return. | [default to 1]
 **limit** | **optional.Int32**| The maximum number of results to return. | [default to 30]
 **type_** | [**optional.Interface of []string**](string.md)| The types of principals to include. | 

### Return type

[**[]TypesPrincipalInfo**](TypesPrincipalInfo.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

