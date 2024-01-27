# code{{classname}}

All URIs are relative to */api/v1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CancelExecution**](PipelineApi.md#CancelExecution) | **Post** /repos/{repo_ref}/pipelines/{pipeline_uid}/executions/{execution_number}/cancel | 
[**CreateExecution**](PipelineApi.md#CreateExecution) | **Post** /repos/{repo_ref}/pipelines/{pipeline_uid}/executions | 
[**CreatePipeline**](PipelineApi.md#CreatePipeline) | **Post** /repos/{repo_ref}/pipelines | 
[**CreateTrigger**](PipelineApi.md#CreateTrigger) | **Post** /repos/{repo_ref}/pipelines/{pipeline_uid}/triggers | 
[**DeleteExecution**](PipelineApi.md#DeleteExecution) | **Delete** /repos/{repo_ref}/pipelines/{pipeline_uid}/executions/{execution_number} | 
[**DeletePipeline**](PipelineApi.md#DeletePipeline) | **Delete** /repos/{repo_ref}/pipelines/{pipeline_uid} | 
[**DeleteTrigger**](PipelineApi.md#DeleteTrigger) | **Delete** /repos/{repo_ref}/pipelines/{pipeline_uid}/triggers/{trigger_uid} | 
[**FindExecution**](PipelineApi.md#FindExecution) | **Get** /repos/{repo_ref}/pipelines/{pipeline_uid}/executions/{execution_number} | 
[**FindPipeline**](PipelineApi.md#FindPipeline) | **Get** /repos/{repo_ref}/pipelines/{pipeline_uid} | 
[**FindTrigger**](PipelineApi.md#FindTrigger) | **Get** /repos/{repo_ref}/pipelines/{pipeline_uid}/triggers/{trigger_uid} | 
[**ListExecutions**](PipelineApi.md#ListExecutions) | **Get** /repos/{repo_ref}/pipelines/{pipeline_uid}/executions | 
[**ListPipelines**](PipelineApi.md#ListPipelines) | **Get** /repos/{repo_ref}/pipelines | 
[**ListTriggers**](PipelineApi.md#ListTriggers) | **Get** /repos/{repo_ref}/pipelines/{pipeline_uid}/triggers | 
[**UpdatePipeline**](PipelineApi.md#UpdatePipeline) | **Patch** /repos/{repo_ref}/pipelines/{pipeline_uid} | 
[**UpdateTrigger**](PipelineApi.md#UpdateTrigger) | **Patch** /repos/{repo_ref}/pipelines/{pipeline_uid}/triggers/{trigger_uid} | 
[**ViewLogs**](PipelineApi.md#ViewLogs) | **Get** /repos/{repo_ref}/pipelines/{pipeline_uid}/executions/{execution_number}/logs/{stage_number}/{step_number} | 

# **CancelExecution**
> TypesExecution CancelExecution(ctx, repoRef, pipelineUid, executionNumber)


### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **repoRef** | **string**|  | 
  **pipelineUid** | **string**|  | 
  **executionNumber** | **string**|  | 

### Return type

[**TypesExecution**](TypesExecution.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **CreateExecution**
> TypesExecution CreateExecution(ctx, repoRef, pipelineUid, optional)


### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **repoRef** | **string**|  | 
  **pipelineUid** | **string**|  | 
 **optional** | ***PipelineApiCreateExecutionOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a PipelineApiCreateExecutionOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **branch** | **optional.String**| Branch to run the execution for. | 

### Return type

[**TypesExecution**](TypesExecution.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **CreatePipeline**
> TypesPipeline CreatePipeline(ctx, repoRef, optional)


### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **repoRef** | **string**|  | 
 **optional** | ***PipelineApiCreatePipelineOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a PipelineApiCreatePipelineOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **body** | [**optional.Interface of OpenapiCreatePipelineRequest**](OpenapiCreatePipelineRequest.md)|  | 

### Return type

[**TypesPipeline**](TypesPipeline.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **CreateTrigger**
> TypesTrigger CreateTrigger(ctx, repoRef, pipelineUid, optional)


### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **repoRef** | **string**|  | 
  **pipelineUid** | **string**|  | 
 **optional** | ***PipelineApiCreateTriggerOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a PipelineApiCreateTriggerOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **body** | [**optional.Interface of OpenapiCreateTriggerRequest**](OpenapiCreateTriggerRequest.md)|  | 

### Return type

[**TypesTrigger**](TypesTrigger.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **DeleteExecution**
> DeleteExecution(ctx, repoRef, pipelineUid, executionNumber)


### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **repoRef** | **string**|  | 
  **pipelineUid** | **string**|  | 
  **executionNumber** | **string**|  | 

### Return type

 (empty response body)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **DeletePipeline**
> DeletePipeline(ctx, repoRef, pipelineUid)


### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **repoRef** | **string**|  | 
  **pipelineUid** | **string**|  | 

### Return type

 (empty response body)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **DeleteTrigger**
> DeleteTrigger(ctx, repoRef, pipelineUid, triggerUid)


### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **repoRef** | **string**|  | 
  **pipelineUid** | **string**|  | 
  **triggerUid** | **string**|  | 

### Return type

 (empty response body)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **FindExecution**
> TypesExecution FindExecution(ctx, repoRef, pipelineUid, executionNumber)


### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **repoRef** | **string**|  | 
  **pipelineUid** | **string**|  | 
  **executionNumber** | **string**|  | 

### Return type

[**TypesExecution**](TypesExecution.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **FindPipeline**
> TypesPipeline FindPipeline(ctx, repoRef, pipelineUid)


### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **repoRef** | **string**|  | 
  **pipelineUid** | **string**|  | 

### Return type

[**TypesPipeline**](TypesPipeline.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **FindTrigger**
> TypesTrigger FindTrigger(ctx, repoRef, pipelineUid, triggerUid)


### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **repoRef** | **string**|  | 
  **pipelineUid** | **string**|  | 
  **triggerUid** | **string**|  | 

### Return type

[**TypesTrigger**](TypesTrigger.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ListExecutions**
> []TypesExecution ListExecutions(ctx, repoRef, pipelineUid, optional)


### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **repoRef** | **string**|  | 
  **pipelineUid** | **string**|  | 
 **optional** | ***PipelineApiListExecutionsOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a PipelineApiListExecutionsOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **page** | **optional.Int32**| The page to return. | [default to 1]
 **limit** | **optional.Int32**| The maximum number of results to return. | [default to 30]

### Return type

[**[]TypesExecution**](TypesExecution.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ListPipelines**
> []TypesPipeline ListPipelines(ctx, repoRef, optional)


### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **repoRef** | **string**|  | 
 **optional** | ***PipelineApiListPipelinesOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a PipelineApiListPipelinesOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **query** | **optional.String**| The substring which is used to filter the repositories by their path name. | 
 **page** | **optional.Int32**| The page to return. | [default to 1]
 **limit** | **optional.Int32**| The maximum number of results to return. | [default to 30]
 **latest** | **optional.Bool**| Whether to fetch latest build information for each pipeline. | 

### Return type

[**[]TypesPipeline**](TypesPipeline.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ListTriggers**
> []TypesTrigger ListTriggers(ctx, repoRef, pipelineUid, optional)


### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **repoRef** | **string**|  | 
  **pipelineUid** | **string**|  | 
 **optional** | ***PipelineApiListTriggersOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a PipelineApiListTriggersOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **query** | **optional.String**| The substring which is used to filter the repositories by their path name. | 
 **page** | **optional.Int32**| The page to return. | [default to 1]
 **limit** | **optional.Int32**| The maximum number of results to return. | [default to 30]

### Return type

[**[]TypesTrigger**](TypesTrigger.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **UpdatePipeline**
> TypesPipeline UpdatePipeline(ctx, repoRef, pipelineUid, optional)


### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **repoRef** | **string**|  | 
  **pipelineUid** | **string**|  | 
 **optional** | ***PipelineApiUpdatePipelineOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a PipelineApiUpdatePipelineOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **body** | [**optional.Interface of OpenapiUpdatePipelineRequest**](OpenapiUpdatePipelineRequest.md)|  | 

### Return type

[**TypesPipeline**](TypesPipeline.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **UpdateTrigger**
> TypesTrigger UpdateTrigger(ctx, repoRef, pipelineUid, triggerUid, optional)


### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **repoRef** | **string**|  | 
  **pipelineUid** | **string**|  | 
  **triggerUid** | **string**|  | 
 **optional** | ***PipelineApiUpdateTriggerOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a PipelineApiUpdateTriggerOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **body** | [**optional.Interface of OpenapiUpdateTriggerRequest**](OpenapiUpdateTriggerRequest.md)|  | 

### Return type

[**TypesTrigger**](TypesTrigger.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ViewLogs**
> []LivelogLine ViewLogs(ctx, repoRef, pipelineUid, executionNumber, stageNumber, stepNumber)


### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **repoRef** | **string**|  | 
  **pipelineUid** | **string**|  | 
  **executionNumber** | **string**|  | 
  **stageNumber** | **string**|  | 
  **stepNumber** | **string**|  | 

### Return type

[**[]LivelogLine**](LivelogLine.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

