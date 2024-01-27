# code{{classname}}

All URIs are relative to */api/v1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateConnector**](ConnectorApi.md#CreateConnector) | **Post** /connectors | 
[**DeleteConnector**](ConnectorApi.md#DeleteConnector) | **Delete** /connectors/{connector_ref} | 
[**FindConnector**](ConnectorApi.md#FindConnector) | **Get** /connectors/{connector_ref} | 
[**UpdateConnector**](ConnectorApi.md#UpdateConnector) | **Patch** /connectors/{connector_ref} | 

# **CreateConnector**
> TypesConnector CreateConnector(ctx, optional)


### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***ConnectorApiCreateConnectorOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a ConnectorApiCreateConnectorOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **body** | [**optional.Interface of OpenapiCreateConnectorRequest**](OpenapiCreateConnectorRequest.md)|  | 

### Return type

[**TypesConnector**](TypesConnector.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **DeleteConnector**
> DeleteConnector(ctx, connectorRef)


### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **connectorRef** | **string**|  | 

### Return type

 (empty response body)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **FindConnector**
> TypesConnector FindConnector(ctx, connectorRef)


### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **connectorRef** | **string**|  | 

### Return type

[**TypesConnector**](TypesConnector.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **UpdateConnector**
> TypesConnector UpdateConnector(ctx, connectorRef, optional)


### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **connectorRef** | **string**|  | 
 **optional** | ***ConnectorApiUpdateConnectorOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a ConnectorApiUpdateConnectorOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **body** | [**optional.Interface of OpenapiUpdateConnectorRequest**](OpenapiUpdateConnectorRequest.md)|  | 

### Return type

[**TypesConnector**](TypesConnector.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

