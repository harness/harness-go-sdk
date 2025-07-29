# nextgen{{classname}}

All URIs are relative to */*

Method | HTTP request | Description
------------- | ------------- | -------------
[**AgentClusterServiceCreate**](ClustersApi.md#AgentClusterServiceCreate) | **Post** /api/v1/agents/{agentIdentifier}/clusters | Create creates a cluster
[**AgentClusterServiceCreateHosted**](ClustersApi.md#AgentClusterServiceCreateHosted) | **Post** /api/v1/agents/{agentIdentifier}/hosted/cluster | CreateHosted creates a harness hosted cluster
[**AgentClusterServiceDelete**](ClustersApi.md#AgentClusterServiceDelete) | **Delete** /api/v1/agents/{agentIdentifier}/clusters/{identifier} | Delete deletes a cluster
[**AgentClusterServiceGet**](ClustersApi.md#AgentClusterServiceGet) | **Get** /api/v1/agents/{agentIdentifier}/clusters/{identifier} | Get returns a cluster by identifier
[**AgentClusterServiceGetByName**](ClustersApi.md#AgentClusterServiceGetByName) | **Get** /api/v1/agents/{agentIdentifier}/cluster_byname | Get a cluster managed by an agent by name
[**AgentClusterServiceGetByUrl**](ClustersApi.md#AgentClusterServiceGetByUrl) | **Get** /api/v1/agents/{agentIdentifier}/cluster_byurl | Get a cluster managed by an agent by URL
[**AgentClusterServiceList**](ClustersApi.md#AgentClusterServiceList) | **Get** /api/v1/agents/{agentIdentifier}/clusters | List returns list of clusters
[**AgentClusterServiceUpdate**](ClustersApi.md#AgentClusterServiceUpdate) | **Put** /api/v1/agents/{agentIdentifier}/clusters/{identifier} | Update updates a cluster
[**AgentGPGKeyServiceList**](ClustersApi.md#AgentGPGKeyServiceList) | **Get** /api/v1/agents/{agentIdentifier}/gpgkeys | List all available repository certificates
[**ClusterServiceExists**](ClustersApi.md#ClusterServiceExists) | **Get** /api/v1/clusters/exists | Check if a cluster exists
[**ClusterServiceListClusters**](ClustersApi.md#ClusterServiceListClusters) | **Post** /api/v1/clusters | List clusters

# **AgentClusterServiceCreate**
> Servicev1Cluster AgentClusterServiceCreate(ctx, body, agentIdentifier, optional)
Create creates a cluster

Create cluster.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**ClustersClusterCreateRequest**](ClustersClusterCreateRequest.md)|  | 
  **agentIdentifier** | **string**| Agent identifier for entity. | 
 **optional** | ***ClustersApiAgentClusterServiceCreateOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a ClustersApiAgentClusterServiceCreateOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **accountIdentifier** | **optional.**| Account Identifier for the Entity. | 
 **orgIdentifier** | **optional.**| Organization Identifier for the Entity. | 
 **projectIdentifier** | **optional.**| Project Identifier for the Entity. | 
 **identifier** | **optional.**|  | 

### Return type

[**Servicev1Cluster**](servicev1Cluster.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **AgentClusterServiceCreateHosted**
> Servicev1Cluster AgentClusterServiceCreateHosted(ctx, agentIdentifier)
CreateHosted creates a harness hosted cluster

Creates Harness hosted cluster.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **agentIdentifier** | **string**| Agent identifier for entity. | 

### Return type

[**Servicev1Cluster**](servicev1Cluster.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **AgentClusterServiceDelete**
> ClustersClusterResponse AgentClusterServiceDelete(ctx, agentIdentifier, identifier, optional)
Delete deletes a cluster

Delete cluster.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **agentIdentifier** | **string**| Agent identifier for entity. | 
  **identifier** | **string**|  | 
 **optional** | ***ClustersApiAgentClusterServiceDeleteOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a ClustersApiAgentClusterServiceDeleteOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **accountIdentifier** | **optional.String**| Account Identifier for the Entity. | 
 **orgIdentifier** | **optional.String**| Organization Identifier for the Entity. | 
 **projectIdentifier** | **optional.String**| Project Identifier for the Entity. | 
 **queryServer** | **optional.String**|  | 
 **queryName** | **optional.String**|  | 
 **queryIdType** | **optional.String**| type is the type of the specified cluster identifier ( \&quot;server\&quot; - default, \&quot;name\&quot; ). | 
 **queryIdValue** | **optional.String**| value holds the cluster server URL or cluster name. | 
 **queryProject** | **optional.String**|  | 
 **forceDelete** | **optional.Bool**|  | 

### Return type

[**ClustersClusterResponse**](clustersClusterResponse.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **AgentClusterServiceGet**
> Servicev1Cluster AgentClusterServiceGet(ctx, agentIdentifier, identifier, optional)
Get returns a cluster by identifier

Get cluster.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **agentIdentifier** | **string**| Agent identifier for entity. | 
  **identifier** | **string**|  | 
 **optional** | ***ClustersApiAgentClusterServiceGetOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a ClustersApiAgentClusterServiceGetOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **accountIdentifier** | **optional.String**| Account Identifier for the Entity. | 
 **orgIdentifier** | **optional.String**| Organization Identifier for the Entity. | 
 **projectIdentifier** | **optional.String**| Project Identifier for the Entity. | 
 **queryServer** | **optional.String**|  | 
 **queryName** | **optional.String**|  | 
 **queryIdType** | **optional.String**| type is the type of the specified cluster identifier ( \&quot;server\&quot; - default, \&quot;name\&quot; ). | 
 **queryIdValue** | **optional.String**| value holds the cluster server URL or cluster name. | 
 **queryProject** | **optional.String**|  | 

### Return type

[**Servicev1Cluster**](servicev1Cluster.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **AgentClusterServiceGetByName**
> Servicev1Cluster AgentClusterServiceGetByName(ctx, agentIdentifier, optional)
Get a cluster managed by an agent by name

Get a cluster managed by agent using name

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **agentIdentifier** | **string**| Agent identifier for entity. | 
 **optional** | ***ClustersApiAgentClusterServiceGetByNameOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a ClustersApiAgentClusterServiceGetByNameOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **accountIdentifier** | **optional.String**| Account Identifier for the Entity. | 
 **orgIdentifier** | **optional.String**| Organization Identifier for the Entity. | 
 **projectIdentifier** | **optional.String**| Project Identifier for the Entity. | 
 **name** | **optional.String**|  | 
 **fetchScopePrefixedIdentifier** | **optional.Bool**| Indicates whether the identifier of cluster fetched must contain harness scope prefix (account./org.) | 

### Return type

[**Servicev1Cluster**](servicev1Cluster.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **AgentClusterServiceGetByUrl**
> Servicev1Cluster AgentClusterServiceGetByUrl(ctx, agentIdentifier, optional)
Get a cluster managed by an agent by URL

Get a cluster managed by agent using URL

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **agentIdentifier** | **string**| Agent identifier for entity. | 
 **optional** | ***ClustersApiAgentClusterServiceGetByUrlOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a ClustersApiAgentClusterServiceGetByUrlOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **accountIdentifier** | **optional.String**| Account Identifier for the Entity. | 
 **orgIdentifier** | **optional.String**| Organization Identifier for the Entity. | 
 **projectIdentifier** | **optional.String**| Project Identifier for the Entity. | 
 **server** | **optional.String**|  | 
 **fetchScopePrefixedIdentifier** | **optional.Bool**| Indicates whether the identifier of cluster fetched must contain harness scope prefix (account./org.) | 

### Return type

[**Servicev1Cluster**](servicev1Cluster.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **AgentClusterServiceList**
> ClustersClusterList AgentClusterServiceList(ctx, agentIdentifier, optional)
List returns list of clusters

List clusters.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **agentIdentifier** | **string**| Agent identifier for entity. | 
 **optional** | ***ClustersApiAgentClusterServiceListOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a ClustersApiAgentClusterServiceListOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **accountIdentifier** | **optional.String**| Account Identifier for the Entity. | 
 **orgIdentifier** | **optional.String**| Organization Identifier for the Entity. | 
 **projectIdentifier** | **optional.String**| Project Identifier for the Entity. | 
 **identifier** | **optional.String**|  | 
 **queryServer** | **optional.String**|  | 
 **queryName** | **optional.String**|  | 
 **queryIdType** | **optional.String**| type is the type of the specified cluster identifier ( \&quot;server\&quot; - default, \&quot;name\&quot; ). | 
 **queryIdValue** | **optional.String**| value holds the cluster server URL or cluster name. | 
 **queryProject** | **optional.String**|  | 

### Return type

[**ClustersClusterList**](clustersClusterList.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **AgentClusterServiceUpdate**
> Servicev1Cluster AgentClusterServiceUpdate(ctx, body, agentIdentifier, identifier, optional)
Update updates a cluster

Update cluster.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**ClustersClusterUpdateRequest**](ClustersClusterUpdateRequest.md)|  | 
  **agentIdentifier** | **string**| Agent identifier for entity. | 
  **identifier** | **string**|  | 
 **optional** | ***ClustersApiAgentClusterServiceUpdateOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a ClustersApiAgentClusterServiceUpdateOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **accountIdentifier** | **optional.**| Account Identifier for the Entity. | 
 **orgIdentifier** | **optional.**| Organization Identifier for the Entity. | 
 **projectIdentifier** | **optional.**| Project Identifier for the Entity. | 
 **forceUpdate** | **optional.**|  | 

### Return type

[**Servicev1Cluster**](servicev1Cluster.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **AgentGPGKeyServiceList**
> GpgkeysGnuPgPublicKeyList AgentGPGKeyServiceList(ctx, agentIdentifier, optional)
List all available repository certificates

List all available repository certificates.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **agentIdentifier** | **string**| Agent identifier for entity. | 
 **optional** | ***ClustersApiAgentGPGKeyServiceListOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a ClustersApiAgentGPGKeyServiceListOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **accountIdentifier** | **optional.String**| Account Identifier for the Entity. | 
 **orgIdentifier** | **optional.String**| Organization Identifier for the Entity. | 
 **projectIdentifier** | **optional.String**| Project Identifier for the Entity. | 
 **queryKeyID** | **optional.String**| The GPG key ID to query for. | 

### Return type

[**GpgkeysGnuPgPublicKeyList**](gpgkeysGnuPGPublicKeyList.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ClusterServiceExists**
> bool ClusterServiceExists(ctx, optional)
Check if a cluster exists

Exists checks whether a cluster with the given identifier exists.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***ClustersApiClusterServiceExistsOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a ClustersApiClusterServiceExistsOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **accountIdentifier** | **optional.String**| Account Identifier for the Entity. | 
 **orgIdentifier** | **optional.String**| Organization Identifier for the Entity. | 
 **projectIdentifier** | **optional.String**| Project Identifier for the Entity. | 
 **agentIdentifier** | **optional.String**| Agent identifier for entity. | 
 **server** | **optional.String**|  | 

### Return type

**bool**

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ClusterServiceListClusters**
> V1Clusterlist ClusterServiceListClusters(ctx, body)
List clusters

ListClusters returns a list of clusters.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**Servicev1ClusterQuery**](Servicev1ClusterQuery.md)|  | 

### Return type

[**V1Clusterlist**](v1Clusterlist.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

