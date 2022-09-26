# {{classname}}

All URIs are relative to */*

Method | HTTP request | Description
------------- | ------------- | -------------
[**AgentClusterServiceCreate**](AgentClusterServiceApi.md#AgentClusterServiceCreate) | **Post** /api/v1/agents/{agentIdentifier}/clusters | Create creates a cluster
[**AgentClusterServiceDelete**](AgentClusterServiceApi.md#AgentClusterServiceDelete) | **Delete** /api/v1/agents/{agentIdentifier}/clusters/{identifier} | Delete deletes a cluster
[**AgentClusterServiceGet**](AgentClusterServiceApi.md#AgentClusterServiceGet) | **Get** /api/v1/agents/{agentIdentifier}/clusters/{identifier} | Get returns a cluster by identifier
[**AgentClusterServiceList**](AgentClusterServiceApi.md#AgentClusterServiceList) | **Get** /api/v1/agents/{agentIdentifier}/clusters | List returns list of clusters
[**AgentClusterServiceUpdate**](AgentClusterServiceApi.md#AgentClusterServiceUpdate) | **Put** /api/v1/agents/{agentIdentifier}/clusters/{identifier} | Update updates a cluster

# **AgentClusterServiceCreate**
> Servicev1Cluster AgentClusterServiceCreate(ctx, body, agentIdentifier, optional)
Create creates a cluster

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**ClusterClusterCreateRequest**](ClusterClusterCreateRequest.md)|  | 
  **agentIdentifier** | **string**|  | 
 **optional** | ***AgentClusterServiceApiAgentClusterServiceCreateOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a AgentClusterServiceApiAgentClusterServiceCreateOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **accountIdentifier** | **optional.**|  | 
 **orgIdentifier** | **optional.**|  | 
 **projectIdentifier** | **optional.**|  | 
 **identifier** | **optional.**|  | 

### Return type

[**Servicev1Cluster**](servicev1Cluster.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **AgentClusterServiceDelete**
> ClusterClusterResponse AgentClusterServiceDelete(ctx, agentIdentifier, identifier, optional)
Delete deletes a cluster

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **agentIdentifier** | **string**|  | 
  **identifier** | **string**|  | 
 **optional** | ***AgentClusterServiceApiAgentClusterServiceDeleteOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a AgentClusterServiceApiAgentClusterServiceDeleteOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **accountIdentifier** | **optional.String**|  | 
 **orgIdentifier** | **optional.String**|  | 
 **projectIdentifier** | **optional.String**|  | 
 **queryServer** | **optional.String**|  | 
 **queryName** | **optional.String**|  | 

### Return type

[**ClusterClusterResponse**](clusterClusterResponse.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **AgentClusterServiceGet**
> Servicev1Cluster AgentClusterServiceGet(ctx, agentIdentifier, identifier, optional)
Get returns a cluster by identifier

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **agentIdentifier** | **string**|  | 
  **identifier** | **string**|  | 
 **optional** | ***AgentClusterServiceApiAgentClusterServiceGetOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a AgentClusterServiceApiAgentClusterServiceGetOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **accountIdentifier** | **optional.String**|  | 
 **orgIdentifier** | **optional.String**|  | 
 **projectIdentifier** | **optional.String**|  | 
 **queryServer** | **optional.String**|  | 
 **queryName** | **optional.String**|  | 

### Return type

[**Servicev1Cluster**](servicev1Cluster.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **AgentClusterServiceList**
> V1alpha1ClusterList AgentClusterServiceList(ctx, agentIdentifier, optional)
List returns list of clusters

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **agentIdentifier** | **string**|  | 
 **optional** | ***AgentClusterServiceApiAgentClusterServiceListOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a AgentClusterServiceApiAgentClusterServiceListOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **accountIdentifier** | **optional.String**|  | 
 **orgIdentifier** | **optional.String**|  | 
 **projectIdentifier** | **optional.String**|  | 
 **identifier** | **optional.String**|  | 
 **queryServer** | **optional.String**|  | 
 **queryName** | **optional.String**|  | 

### Return type

[**V1alpha1ClusterList**](v1alpha1ClusterList.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **AgentClusterServiceUpdate**
> Servicev1Cluster AgentClusterServiceUpdate(ctx, body, agentIdentifier, identifier, optional)
Update updates a cluster

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**ClusterClusterUpdateRequest**](ClusterClusterUpdateRequest.md)|  | 
  **agentIdentifier** | **string**|  | 
  **identifier** | **string**|  | 
 **optional** | ***AgentClusterServiceApiAgentClusterServiceUpdateOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a AgentClusterServiceApiAgentClusterServiceUpdateOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **accountIdentifier** | **optional.**|  | 
 **orgIdentifier** | **optional.**|  | 
 **projectIdentifier** | **optional.**|  | 

### Return type

[**Servicev1Cluster**](servicev1Cluster.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

