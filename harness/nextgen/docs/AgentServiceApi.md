# {{classname}}

All URIs are relative to */*

Method | HTTP request | Description
------------- | ------------- | -------------
[**AgentServiceCreate**](AgentServiceApi.md#AgentServiceCreate) | **Post** /api/v1/agents | 
[**AgentServiceDelete**](AgentServiceApi.md#AgentServiceDelete) | **Delete** /api/v1/agents/{identifier} | 
[**AgentServiceGet**](AgentServiceApi.md#AgentServiceGet) | **Get** /api/v1/agents/{identifier} | 
[**AgentServiceGetAppProjectMappingList**](AgentServiceApi.md#AgentServiceGetAppProjectMappingList) | **Get** /api/v1/agents/{identifier}/appprojectsmapping | 
[**AgentServiceGetDeployYaml**](AgentServiceApi.md#AgentServiceGetDeployYaml) | **Get** /api/v1/agents/{agentIdentifier}/deploy.yaml | 
[**AgentServiceList**](AgentServiceApi.md#AgentServiceList) | **Get** /api/v1/agents | 
[**AgentServiceRegenerateCredentials**](AgentServiceApi.md#AgentServiceRegenerateCredentials) | **Post** /api/v1/agents/{identifier}/credentials | 
[**AgentServiceUnique**](AgentServiceApi.md#AgentServiceUnique) | **Get** /api/v1/agents/{identifier}/unique | 
[**AgentServiceUpdate**](AgentServiceApi.md#AgentServiceUpdate) | **Put** /api/v1/agents/{agent.identifier} | 

# **AgentServiceCreate**
> V1Agent AgentServiceCreate(ctx, body)


### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**V1Agent**](V1Agent.md)|  | 

### Return type

[**V1Agent**](v1Agent.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **AgentServiceDelete**
> V1Agent AgentServiceDelete(ctx, identifier, optional)


### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **identifier** | **string**|  | 
 **optional** | ***AgentServiceApiAgentServiceDeleteOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a AgentServiceApiAgentServiceDeleteOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **accountIdentifier** | **optional.String**|  | 
 **projectIdentifier** | **optional.String**|  | 
 **orgIdentifier** | **optional.String**|  | 
 **name** | **optional.String**|  | 
 **type_** | **optional.String**|  | [default to AGENT_TYPE_UNSET]
 **tags** | [**optional.Interface of []string**](string.md)|  | 
 **searchTerm** | **optional.String**|  | 
 **pageSize** | **optional.Int32**|  | 
 **pageIndex** | **optional.Int32**|  | 
 **scope** | **optional.String**|  | [default to AGENT_SCOPE_UNSET]

### Return type

[**V1Agent**](v1Agent.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **AgentServiceGet**
> V1Agent AgentServiceGet(ctx, identifier, optional)


### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **identifier** | **string**|  | 
 **optional** | ***AgentServiceApiAgentServiceGetOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a AgentServiceApiAgentServiceGetOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **accountIdentifier** | **optional.String**|  | 
 **projectIdentifier** | **optional.String**|  | 
 **orgIdentifier** | **optional.String**|  | 
 **name** | **optional.String**|  | 
 **type_** | **optional.String**|  | [default to AGENT_TYPE_UNSET]
 **tags** | [**optional.Interface of []string**](string.md)|  | 
 **searchTerm** | **optional.String**|  | 
 **pageSize** | **optional.Int32**|  | 
 **pageIndex** | **optional.Int32**|  | 
 **scope** | **optional.String**|  | [default to AGENT_SCOPE_UNSET]

### Return type

[**V1Agent**](v1Agent.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **AgentServiceGetAppProjectMappingList**
> V1AppProjectMapping AgentServiceGetAppProjectMappingList(ctx, identifier, optional)


### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **identifier** | **string**|  | 
 **optional** | ***AgentServiceApiAgentServiceGetAppProjectMappingListOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a AgentServiceApiAgentServiceGetAppProjectMappingListOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **accountIdentifier** | **optional.String**|  | 
 **projectIdentifier** | **optional.String**|  | 
 **orgIdentifier** | **optional.String**|  | 
 **name** | **optional.String**|  | 
 **type_** | **optional.String**|  | [default to AGENT_TYPE_UNSET]
 **tags** | [**optional.Interface of []string**](string.md)|  | 
 **searchTerm** | **optional.String**|  | 
 **pageSize** | **optional.Int32**|  | 
 **pageIndex** | **optional.Int32**|  | 
 **scope** | **optional.String**|  | [default to AGENT_SCOPE_UNSET]

### Return type

[**V1AppProjectMapping**](v1AppProjectMapping.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **AgentServiceGetDeployYaml**
> string AgentServiceGetDeployYaml(ctx, agentIdentifier, optional)


### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **agentIdentifier** | **string**|  | 
 **optional** | ***AgentServiceApiAgentServiceGetDeployYamlOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a AgentServiceApiAgentServiceGetDeployYamlOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **accountIdentifier** | **optional.String**|  | 
 **orgIdentifier** | **optional.String**|  | 
 **projectIdentifier** | **optional.String**|  | 
 **namespace** | **optional.String**|  | 

### Return type

**string**

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/x-yml

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **AgentServiceList**
> V1AgentList AgentServiceList(ctx, optional)


### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***AgentServiceApiAgentServiceListOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a AgentServiceApiAgentServiceListOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **accountIdentifier** | **optional.String**|  | 
 **projectIdentifier** | **optional.String**|  | 
 **orgIdentifier** | **optional.String**|  | 
 **identifier** | **optional.String**|  | 
 **name** | **optional.String**|  | 
 **type_** | **optional.String**|  | [default to AGENT_TYPE_UNSET]
 **tags** | [**optional.Interface of []string**](string.md)|  | 
 **searchTerm** | **optional.String**|  | 
 **pageSize** | **optional.Int32**|  | 
 **pageIndex** | **optional.Int32**|  | 
 **scope** | **optional.String**|  | [default to AGENT_SCOPE_UNSET]

### Return type

[**V1AgentList**](v1AgentList.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **AgentServiceRegenerateCredentials**
> V1Agent AgentServiceRegenerateCredentials(ctx, identifier)


### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **identifier** | **string**|  | 

### Return type

[**V1Agent**](v1Agent.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **AgentServiceUnique**
> V1UniqueMessage AgentServiceUnique(ctx, identifier, optional)


### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **identifier** | **string**|  | 
 **optional** | ***AgentServiceApiAgentServiceUniqueOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a AgentServiceApiAgentServiceUniqueOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **accountIdentifier** | **optional.String**|  | 
 **projectIdentifier** | **optional.String**|  | 
 **orgIdentifier** | **optional.String**|  | 
 **name** | **optional.String**|  | 
 **type_** | **optional.String**|  | [default to AGENT_TYPE_UNSET]
 **tags** | [**optional.Interface of []string**](string.md)|  | 
 **searchTerm** | **optional.String**|  | 
 **pageSize** | **optional.Int32**|  | 
 **pageIndex** | **optional.Int32**|  | 
 **scope** | **optional.String**|  | [default to AGENT_SCOPE_UNSET]

### Return type

[**V1UniqueMessage**](v1UniqueMessage.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **AgentServiceUpdate**
> V1Agent AgentServiceUpdate(ctx, body, agentIdentifier)


### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**V1Agent**](V1Agent.md)|  | 
  **agentIdentifier** | **string**| The gitops-server generated ID for this gitops-agent | 

### Return type

[**V1Agent**](v1Agent.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

