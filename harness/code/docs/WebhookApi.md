# code{{classname}}

All URIs are relative to */api/v1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateWebhook**](WebhookApi.md#CreateWebhook) | **Post** /repos/{repo_ref}/webhooks | 
[**DeleteWebhook**](WebhookApi.md#DeleteWebhook) | **Delete** /repos/{repo_ref}/webhooks/{webhook_id} | 
[**GetWebhook**](WebhookApi.md#GetWebhook) | **Get** /repos/{repo_ref}/webhooks/{webhook_id} | 
[**GetWebhookExecution**](WebhookApi.md#GetWebhookExecution) | **Get** /repos/{repo_ref}/webhooks/{webhook_id}/executions/{webhook_execution_id} | 
[**ListWebhookExecutions**](WebhookApi.md#ListWebhookExecutions) | **Get** /repos/{repo_ref}/webhooks/{webhook_id}/executions | 
[**ListWebhooks**](WebhookApi.md#ListWebhooks) | **Get** /repos/{repo_ref}/webhooks | 
[**UpdateWebhook**](WebhookApi.md#UpdateWebhook) | **Patch** /repos/{repo_ref}/webhooks/{webhook_id} | 

# **CreateWebhook**
> OpenapiWebhookType CreateWebhook(ctx, repoRef, optional)


### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **repoRef** | **string**|  | 
 **optional** | ***WebhookApiCreateWebhookOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a WebhookApiCreateWebhookOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **body** | [**optional.Interface of OpenapiCreateWebhookRequest**](OpenapiCreateWebhookRequest.md)|  | 

### Return type

[**OpenapiWebhookType**](OpenapiWebhookType.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **DeleteWebhook**
> DeleteWebhook(ctx, repoRef, webhookId)


### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **repoRef** | **string**|  | 
  **webhookId** | **int32**|  | 

### Return type

 (empty response body)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetWebhook**
> OpenapiWebhookType GetWebhook(ctx, repoRef, webhookId)


### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **repoRef** | **string**|  | 
  **webhookId** | **int32**|  | 

### Return type

[**OpenapiWebhookType**](OpenapiWebhookType.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetWebhookExecution**
> TypesWebhookExecution GetWebhookExecution(ctx, repoRef, webhookId, webhookExecutionId, optional)


### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **repoRef** | **string**|  | 
  **webhookId** | **int32**|  | 
  **webhookExecutionId** | **int32**|  | 
 **optional** | ***WebhookApiGetWebhookExecutionOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a WebhookApiGetWebhookExecutionOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **page** | **optional.Int32**| The page to return. | [default to 1]
 **limit** | **optional.Int32**| The maximum number of results to return. | [default to 30]

### Return type

[**TypesWebhookExecution**](TypesWebhookExecution.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ListWebhookExecutions**
> []TypesWebhookExecution ListWebhookExecutions(ctx, repoRef, webhookId, optional)


### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **repoRef** | **string**|  | 
  **webhookId** | **int32**|  | 
 **optional** | ***WebhookApiListWebhookExecutionsOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a WebhookApiListWebhookExecutionsOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **page** | **optional.Int32**| The page to return. | [default to 1]
 **limit** | **optional.Int32**| The maximum number of results to return. | [default to 30]

### Return type

[**[]TypesWebhookExecution**](TypesWebhookExecution.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ListWebhooks**
> []OpenapiWebhookType ListWebhooks(ctx, repoRef, optional)


### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **repoRef** | **string**|  | 
 **optional** | ***WebhookApiListWebhooksOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a WebhookApiListWebhooksOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **query** | **optional.String**| The substring which is used to filter the spaces by their path name. | 
 **sort** | **optional.String**| The data by which the webhooks are sorted. | [default to id]
 **order** | **optional.String**| The order of the output. | [default to asc]
 **page** | **optional.Int32**| The page to return. | [default to 1]
 **limit** | **optional.Int32**| The maximum number of results to return. | [default to 30]

### Return type

[**[]OpenapiWebhookType**](OpenapiWebhookType.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **UpdateWebhook**
> OpenapiWebhookType UpdateWebhook(ctx, repoRef, webhookId, optional)


### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **repoRef** | **string**|  | 
  **webhookId** | **int32**|  | 
 **optional** | ***WebhookApiUpdateWebhookOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a WebhookApiUpdateWebhookOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **body** | [**optional.Interface of OpenapiUpdateWebhookRequest**](OpenapiUpdateWebhookRequest.md)|  | 

### Return type

[**OpenapiWebhookType**](OpenapiWebhookType.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

