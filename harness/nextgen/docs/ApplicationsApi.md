# {{classname}}

All URIs are relative to *https://app.harness.io/gateway*

Method | HTTP request | Description
------------- | ------------- | -------------
[**AgentApplicationServiceCreate**](ApplicationsApi.md#AgentApplicationServiceCreate) | **Post** /api/v1/agents/{agentIdentifier}/applications | Create an application
[**AgentApplicationServiceDelete**](ApplicationsApi.md#AgentApplicationServiceDelete) | **Delete** /api/v1/agents/{agentIdentifier}/applications/{request.name} | Delete an application
[**AgentApplicationServiceDeleteResource**](ApplicationsApi.md#AgentApplicationServiceDeleteResource) | **Delete** /api/v1/agents/{agentIdentifier}/applications/{request.name}/resource | Delete resource
[**AgentApplicationServiceGet**](ApplicationsApi.md#AgentApplicationServiceGet) | **Get** /api/v1/agents/{agentIdentifier}/applications/{query.name} | Get application
[**AgentApplicationServiceGetApplicationSyncWindows**](ApplicationsApi.md#AgentApplicationServiceGetApplicationSyncWindows) | **Get** /api/v1/agents/{agentIdentifier}/applications/{query.name}/syncwindows | Get sync windows of the application
[**AgentApplicationServiceGetManifests**](ApplicationsApi.md#AgentApplicationServiceGetManifests) | **Get** /api/v1/agents/{agentIdentifier}/applications/{query.name}/manifests | Get application manifest
[**AgentApplicationServiceGetResource**](ApplicationsApi.md#AgentApplicationServiceGetResource) | **Get** /api/v1/agents/{agentIdentifier}/applications/{request.name}/resource | Get an application resource
[**AgentApplicationServiceList**](ApplicationsApi.md#AgentApplicationServiceList) | **Get** /api/v1/agents/{agentIdentifier}/applications | List applications for a specific agent
[**AgentApplicationServiceListResourceActions**](ApplicationsApi.md#AgentApplicationServiceListResourceActions) | **Get** /api/v1/agents/{agentIdentifier}/applications/{request.name}/resource/actions | List resource actions
[**AgentApplicationServiceListResourceEvents**](ApplicationsApi.md#AgentApplicationServiceListResourceEvents) | **Get** /api/v1/agents/{agentIdentifier}/applications/{query.name}/events | List resource events
[**AgentApplicationServiceManagedResources**](ApplicationsApi.md#AgentApplicationServiceManagedResources) | **Get** /api/v1/agents/{agentIdentifier}/applications/{query.applicationName}/managed-resources | Get ManagedResources
[**AgentApplicationServiceParent**](ApplicationsApi.md#AgentApplicationServiceParent) | **Get** /api/v1/agents/{agentIdentifier}/applications/parent/{query.name} | Get parent application
[**AgentApplicationServicePatch**](ApplicationsApi.md#AgentApplicationServicePatch) | **Patch** /api/v1/agents/{agentIdentifier}/applications/{request.name} | Patch an application
[**AgentApplicationServicePatchResource**](ApplicationsApi.md#AgentApplicationServicePatchResource) | **Post** /api/v1/agents/{agentIdentifier}/applications/{request.name}/resource | Patch an application resource
[**AgentApplicationServicePodLogs**](ApplicationsApi.md#AgentApplicationServicePodLogs) | **Get** /api/v1/agents/{agentIdentifier}/applications/{query.name}/pods/{query.podName}/logs | PodLogs returns stream of log entries for the specified pod(s).
[**AgentApplicationServicePodLogs2**](ApplicationsApi.md#AgentApplicationServicePodLogs2) | **Get** /api/v1/agents/{agentIdentifier}/applications/{query.name}/logs | PodLogs returns stream of log entries for the specified pod(s).
[**AgentApplicationServiceResourceTree**](ApplicationsApi.md#AgentApplicationServiceResourceTree) | **Get** /api/v1/agents/{agentIdentifier}/applications/{query.applicationName}/resource-tree | Get resource tree
[**AgentApplicationServiceRevisionMetadata**](ApplicationsApi.md#AgentApplicationServiceRevisionMetadata) | **Get** /api/v1/agents/{agentIdentifier}/applications/{query.name}/revisions/{query.revision}/metadata | Get revision metadata
[**AgentApplicationServiceRollback**](ApplicationsApi.md#AgentApplicationServiceRollback) | **Post** /api/v1/agents/{agentIdentifier}/applications/{request.name}/rollback | Rollback syncs an application to its target state
[**AgentApplicationServiceRunResourceAction**](ApplicationsApi.md#AgentApplicationServiceRunResourceAction) | **Post** /api/v1/agents/{agentIdentifier}/applications/{request.name}/resource/actions | Run resource action
[**AgentApplicationServiceSync**](ApplicationsApi.md#AgentApplicationServiceSync) | **Post** /api/v1/agents/{agentIdentifier}/applications/{request.name}/sync | Sync an application
[**AgentApplicationServiceTerminateOperation**](ApplicationsApi.md#AgentApplicationServiceTerminateOperation) | **Delete** /api/v1/agents/{agentIdentifier}/applications/{request.name}/operation | Terminate operation
[**AgentApplicationServiceUpdate**](ApplicationsApi.md#AgentApplicationServiceUpdate) | **Put** /api/v1/agents/{agentIdentifier}/applications/{request.application.metadata.name} | Update an application
[**AgentApplicationServiceUpdateSpec**](ApplicationsApi.md#AgentApplicationServiceUpdateSpec) | **Put** /api/v1/agents/{agentIdentifier}/applications/{request.name}/spec | Update application spec
[**AgentApplicationServiceWatch**](ApplicationsApi.md#AgentApplicationServiceWatch) | **Get** /api/v1/agents/{agentIdentifier}/stream/applications | Watch returns stream of application change events
[**AgentApplicationServiceWatchResourceTree**](ApplicationsApi.md#AgentApplicationServiceWatchResourceTree) | **Get** /api/v1/agents/{agentIdentifier}/stream/applications/{query.applicationName}/resource-tree | WatchResourceTree returns stream of application resource tree
[**ApplicationServiceExists**](ApplicationsApi.md#ApplicationServiceExists) | **Get** /api/v1/applications/{name}/exists | Check if an application exists
[**ApplicationServiceIsMultiSourceEnabled**](ApplicationsApi.md#ApplicationServiceIsMultiSourceEnabled) | **Get** /api/v1/applications/multisource | Check if multisource feature is enabled
[**ApplicationServiceListAppSets**](ApplicationsApi.md#ApplicationServiceListAppSets) | **Get** /api/v1/applications/sets | List ApplicationSets
[**ApplicationServiceListAppSync**](ApplicationsApi.md#ApplicationServiceListAppSync) | **Post** /api/v1/applications/sync | List application sync status
[**ApplicationServiceListAttributes**](ApplicationsApi.md#ApplicationServiceListAttributes) | **Get** /api/v1/applications/attributes | List attributes for applications
[**ApplicationServiceListNs**](ApplicationsApi.md#ApplicationServiceListNs) | **Post** /api/v1/applications/ns | List namespaces for a specific query

# **AgentApplicationServiceCreate**
> Servicev1Application AgentApplicationServiceCreate(ctx, body, agentIdentifier, accountIdentifier, optional)
Create creates an application

Creates application in project.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**ApplicationsApplicationCreateRequest**](ApplicationsApplicationCreateRequest.md)|  | 
  **agentIdentifier** | **string**| Agent identifier for entity. | 
  **accountIdentifier** | **string**| Account Identifier for the Entity. | 
 **optional** | ***ApplicationsApiAgentApplicationServiceCreateOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a ApplicationsApiAgentApplicationServiceCreateOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **orgIdentifier** | **optional.**| Organization Identifier for the Entity. | 
 **projectIdentifier** | **optional.**| Project Identifier for the Entity. | 
 **clusterIdentifier** | **optional.**|  | 
 **repoIdentifier** | **optional.**|  | 
 **skipRepoValidation** | **optional.**|  | 
 **repoIdentifiers** | [**optional.Interface of []string**](string.md)|  | 

### Return type

[**Servicev1Application**](servicev1Application.md)

### Authorization

[ApiKey](../README.md#ApiKey)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **AgentApplicationServiceDelete**
> ApplicationsApplicationResponse AgentApplicationServiceDelete(ctx, agentIdentifier, requestName, accountIdentifier, orgIdentifier, projectIdentifier, optional)
Delete deletes an application

Delete deletes an application.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **agentIdentifier** | **string**| Agent identifier for entity. | 
  **requestName** | **string**|  | 
  **accountIdentifier** | **string**| Account Identifier for the Entity. | 
  **orgIdentifier** | **string**| Organization Identifier for the Entity. | 
  **projectIdentifier** | **string**| Project Identifier for the Entity. | 
 **optional** | ***ApplicationsApiAgentApplicationServiceDeleteOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a ApplicationsApiAgentApplicationServiceDeleteOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------





 **requestCascade** | **optional.Bool**|  | 
 **requestPropagationPolicy** | **optional.String**|  | 
 **requestAppNamespace** | **optional.String**|  | 
 **requestProject** | **optional.String**|  | 
 **optionsRemoveExistingFinalizers** | **optional.Bool**|  | 

### Return type

[**ApplicationsApplicationResponse**](applicationsApplicationResponse.md)

### Authorization

[ApiKey](../README.md#ApiKey)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **AgentApplicationServiceDeleteResource**
> ApplicationsApplicationResponse AgentApplicationServiceDeleteResource(ctx, agentIdentifier, requestName, accountIdentifier, orgIdentifier, projectIdentifier, optional)
DeleteResource deletes a single application resource

DeleteResource deletes a single application resource.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **agentIdentifier** | **string**| Agent identifier for entity. | 
  **requestName** | **string**|  | 
  **accountIdentifier** | **string**| Account Identifier for the Entity. | 
  **orgIdentifier** | **string**| Organization Identifier for the Entity. | 
  **projectIdentifier** | **string**| Project Identifier for the Entity. | 
 **optional** | ***ApplicationsApiAgentApplicationServiceDeleteResourceOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a ApplicationsApiAgentApplicationServiceDeleteResourceOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------





 **requestNamespace** | **optional.String**|  | 
 **requestResourceName** | **optional.String**|  | 
 **requestVersion** | **optional.String**|  | 
 **requestGroup** | **optional.String**|  | 
 **requestKind** | **optional.String**|  | 
 **requestForce** | **optional.Bool**|  | 
 **requestOrphan** | **optional.Bool**|  | 
 **requestAppNamespace** | **optional.String**|  | 
 **requestProject** | **optional.String**|  | 

### Return type

[**ApplicationsApplicationResponse**](applicationsApplicationResponse.md)

### Authorization

[ApiKey](../README.md#ApiKey)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **AgentApplicationServiceGet**
> Servicev1Application AgentApplicationServiceGet(ctx, agentIdentifier, queryName, accountIdentifier, orgIdentifier, projectIdentifier, optional)
Get returns an application by name

 Get returns an application by name

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **agentIdentifier** | **string**| Agent identifier for entity. | 
  **queryName** | **string**| the application&#x27;s name | 
  **accountIdentifier** | **string**| Account Identifier for the Entity. | 
  **orgIdentifier** | **string**| Organization Identifier for the Entity. | 
  **projectIdentifier** | **string**| Project Identifier for the Entity. | 
 **optional** | ***ApplicationsApiAgentApplicationServiceGetOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a ApplicationsApiAgentApplicationServiceGetOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------





 **queryRefresh** | **optional.String**| forces application reconciliation if set to true. | 
 **queryProject** | [**optional.Interface of []string**](string.md)| the project names to restrict returned list applications. | 
 **queryResourceVersion** | **optional.String**| when specified with a watch call, shows changes that occur after that particular version of a resource. | 
 **querySelector** | **optional.String**| the selector to to restrict returned list to applications only with matched labels. | 
 **queryRepo** | **optional.String**| the repoURL to restrict returned list applications. | 
 **queryAppNamespace** | **optional.String**|  | 

### Return type

[**Servicev1Application**](servicev1Application.md)

### Authorization

[ApiKey](../README.md#ApiKey)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **AgentApplicationServiceGetApplicationSyncWindows**
> ApplicationsApplicationSyncWindowsResponse AgentApplicationServiceGetApplicationSyncWindows(ctx, agentIdentifier, queryName, accountIdentifier, orgIdentifier, projectIdentifier)
Get returns sync windows of the application

GetApplicationSyncWindows returns sync windows of the application.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **agentIdentifier** | **string**| Agent identifier for entity. | 
  **queryName** | **string**|  | 
  **accountIdentifier** | **string**| Account Identifier for the Entity. | 
  **orgIdentifier** | **string**| Organization Identifier for the Entity. | 
  **projectIdentifier** | **string**| Project Identifier for the Entity. | 

### Return type

[**ApplicationsApplicationSyncWindowsResponse**](applicationsApplicationSyncWindowsResponse.md)

### Authorization

[ApiKey](../README.md#ApiKey)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **AgentApplicationServiceGetManifests**
> RepositoriesManifestResponse AgentApplicationServiceGetManifests(ctx, agentIdentifier, queryName, accountIdentifier, orgIdentifier, projectIdentifier, optional)
GetManifests returns application manifests

GetManifests returns application manifests.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **agentIdentifier** | **string**| Agent identifier for entity. | 
  **queryName** | **string**|  | 
  **accountIdentifier** | **string**| Account Identifier for the Entity. | 
  **orgIdentifier** | **string**| Organization Identifier for the Entity. | 
  **projectIdentifier** | **string**| Project Identifier for the Entity. | 
 **optional** | ***ApplicationsApiAgentApplicationServiceGetManifestsOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a ApplicationsApiAgentApplicationServiceGetManifestsOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------





 **queryRevision** | **optional.String**|  | 
 **queryAppNamespace** | **optional.String**|  | 
 **queryProject** | **optional.String**|  | 
 **querySourcePositions** | [**optional.Interface of []string**](string.md)|  | 
 **queryRevisions** | [**optional.Interface of []string**](string.md)|  | 

### Return type

[**RepositoriesManifestResponse**](repositoriesManifestResponse.md)

### Authorization

[ApiKey](../README.md#ApiKey)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **AgentApplicationServiceGetResource**
> ApplicationsApplicationResourceResponse AgentApplicationServiceGetResource(ctx, agentIdentifier, requestName, accountIdentifier, orgIdentifier, projectIdentifier, optional)
GetResource returns single application resource

GetResource returns single application resource.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **agentIdentifier** | **string**| Agent identifier for entity. | 
  **requestName** | **string**|  | 
  **accountIdentifier** | **string**| Account Identifier for the Entity. | 
  **orgIdentifier** | **string**| Organization Identifier for the Entity. | 
  **projectIdentifier** | **string**| Project Identifier for the Entity. | 
 **optional** | ***ApplicationsApiAgentApplicationServiceGetResourceOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a ApplicationsApiAgentApplicationServiceGetResourceOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------





 **requestNamespace** | **optional.String**|  | 
 **requestResourceName** | **optional.String**|  | 
 **requestVersion** | **optional.String**|  | 
 **requestGroup** | **optional.String**|  | 
 **requestKind** | **optional.String**|  | 
 **requestAppNamespace** | **optional.String**|  | 
 **requestProject** | **optional.String**|  | 

### Return type

[**ApplicationsApplicationResourceResponse**](applicationsApplicationResourceResponse.md)

### Authorization

[ApiKey](../README.md#ApiKey)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **AgentApplicationServiceList**
> ApplicationsApplicationList AgentApplicationServiceList(ctx, agentIdentifier, accountIdentifier, orgIdentifier, projectIdentifier, optional)
List returns list of applications for a specific agent

List returns list of applications for a specific agent.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **agentIdentifier** | **string**| Agent identifier for entity. | 
  **accountIdentifier** | **string**| Account Identifier for the Entity. | 
  **orgIdentifier** | **string**| Organization Identifier for the Entity. | 
  **projectIdentifier** | **string**| Project Identifier for the Entity. | 
 **optional** | ***ApplicationsApiAgentApplicationServiceListOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a ApplicationsApiAgentApplicationServiceListOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------




 **queryName** | **optional.String**| the application&#x27;s name. | 
 **queryRefresh** | **optional.String**| forces application reconciliation if set to true. | 
 **queryProject** | [**optional.Interface of []string**](string.md)| the project names to restrict returned list applications. | 
 **queryResourceVersion** | **optional.String**| when specified with a watch call, shows changes that occur after that particular version of a resource. | 
 **querySelector** | **optional.String**| the selector to to restrict returned list to applications only with matched labels. | 
 **queryRepo** | **optional.String**| the repoURL to restrict returned list applications. | 
 **queryAppNamespace** | **optional.String**|  | 

### Return type

[**ApplicationsApplicationList**](applicationsApplicationList.md)

### Authorization

[ApiKey](../README.md#ApiKey)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **AgentApplicationServiceListResourceActions**
> ApplicationsResourceActionsListResponse AgentApplicationServiceListResourceActions(ctx, agentIdentifier, requestName, accountIdentifier, orgIdentifier, projectIdentifier, optional)
ListResourceActions returns list of resource actions

ListResourceActions returns list of resource actions.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **agentIdentifier** | **string**| Agent identifier for entity. | 
  **requestName** | **string**|  | 
  **accountIdentifier** | **string**| Account Identifier for the Entity. | 
  **orgIdentifier** | **string**| Organization Identifier for the Entity. | 
  **projectIdentifier** | **string**| Project Identifier for the Entity. | 
 **optional** | ***ApplicationsApiAgentApplicationServiceListResourceActionsOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a ApplicationsApiAgentApplicationServiceListResourceActionsOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------





 **requestNamespace** | **optional.String**|  | 
 **requestResourceName** | **optional.String**|  | 
 **requestVersion** | **optional.String**|  | 
 **requestGroup** | **optional.String**|  | 
 **requestKind** | **optional.String**|  | 
 **requestAppNamespace** | **optional.String**|  | 
 **requestProject** | **optional.String**|  | 

### Return type

[**ApplicationsResourceActionsListResponse**](applicationsResourceActionsListResponse.md)

### Authorization

[ApiKey](../README.md#ApiKey)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **AgentApplicationServiceListResourceEvents**
> ApplicationsEventList AgentApplicationServiceListResourceEvents(ctx, agentIdentifier, queryName, accountIdentifier, orgIdentifier, projectIdentifier, optional)
ListResourceEvents returns a list of event resources

ListResourceEvents returns list of event resources.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **agentIdentifier** | **string**| Agent identifier for entity. | 
  **queryName** | **string**|  | 
  **accountIdentifier** | **string**| Account Identifier for the Entity. | 
  **orgIdentifier** | **string**| Organization Identifier for the Entity. | 
  **projectIdentifier** | **string**| Project Identifier for the Entity. | 
 **optional** | ***ApplicationsApiAgentApplicationServiceListResourceEventsOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a ApplicationsApiAgentApplicationServiceListResourceEventsOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------





 **queryResourceNamespace** | **optional.String**|  | 
 **queryResourceName** | **optional.String**|  | 
 **queryResourceUID** | **optional.String**|  | 
 **queryAppNamespace** | **optional.String**|  | 
 **queryProject** | **optional.String**|  | 

### Return type

[**ApplicationsEventList**](applicationsEventList.md)

### Authorization

[ApiKey](../README.md#ApiKey)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **AgentApplicationServiceManagedResources**
> ApplicationsManagedResourcesResponse AgentApplicationServiceManagedResources(ctx, agentIdentifier, queryApplicationName, accountIdentifier, orgIdentifier, projectIdentifier, optional)
ManagedResources returns list of managed resources

ManagedResources returns list of managed resources.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **agentIdentifier** | **string**| Agent identifier for entity. | 
  **queryApplicationName** | **string**|  | 
  **accountIdentifier** | **string**| Account Identifier for the Entity. | 
  **orgIdentifier** | **string**| Organization Identifier for the Entity. | 
  **projectIdentifier** | **string**| Project Identifier for the Entity. | 
 **optional** | ***ApplicationsApiAgentApplicationServiceManagedResourcesOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a ApplicationsApiAgentApplicationServiceManagedResourcesOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------





 **queryNamespace** | **optional.String**|  | 
 **queryName** | **optional.String**|  | 
 **queryVersion** | **optional.String**|  | 
 **queryGroup** | **optional.String**|  | 
 **queryKind** | **optional.String**|  | 
 **queryAppNamespace** | **optional.String**|  | 
 **queryProject** | **optional.String**|  | 

### Return type

[**ApplicationsManagedResourcesResponse**](applicationsManagedResourcesResponse.md)

### Authorization

[ApiKey](../README.md#ApiKey)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **AgentApplicationServiceParent**
> Servicev1Application AgentApplicationServiceParent(ctx, agentIdentifier, queryName, optional)
Get returns an application by name

 Parent returns a parent application ( app of apps or  appset) by child name

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **agentIdentifier** | **string**| Agent identifier for entity. | 
  **queryName** | **string**| the application&#x27;s name | 
 **optional** | ***ApplicationsApiAgentApplicationServiceParentOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a ApplicationsApiAgentApplicationServiceParentOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **accountIdentifier** | **optional.String**| Account Identifier for the Entity. | 
 **orgIdentifier** | **optional.String**| Organization Identifier for the Entity. | 
 **projectIdentifier** | **optional.String**| Project Identifier for the Entity. | 
 **queryRefresh** | **optional.String**| forces application reconciliation if set to true. | 
 **queryProject** | [**optional.Interface of []string**](string.md)| the project names to restrict returned list applications. | 
 **queryResourceVersion** | **optional.String**| when specified with a watch call, shows changes that occur after that particular version of a resource. | 
 **querySelector** | **optional.String**| the selector to to restrict returned list to applications only with matched labels. | 
 **queryRepo** | **optional.String**| the repoURL to restrict returned list applications. | 
 **queryAppNamespace** | **optional.String**|  | 

### Return type

[**Servicev1Application**](servicev1Application.md)

### Authorization

[ApiKey](../README.md#ApiKey)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **AgentApplicationServicePatch**
> Servicev1Application AgentApplicationServicePatch(ctx, body, agentIdentifier, requestName)
Patch patch an application

Patch applys a patches to an application.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**Servicev1ApplicationPatchRequest**](Servicev1ApplicationPatchRequest.md)|  | 
  **agentIdentifier** | **string**| Agent identifier for entity. | 
  **requestName** | **string**|  | 

### Return type

[**Servicev1Application**](servicev1Application.md)

### Authorization

[ApiKey](../README.md#ApiKey)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **AgentApplicationServicePatchResource**
> ApplicationsApplicationResourceResponse AgentApplicationServicePatchResource(ctx, body, agentIdentifier, requestName, accountIdentifier, orgIdentifier, projectIdentifier)
PatchResource patch single application resource

PatchResource patch single application resource.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**ApplicationsApplicationResourcePatchRequest**](ApplicationsApplicationResourcePatchRequest.md)|  | 
  **agentIdentifier** | **string**| Agent identifier for entity. | 
  **requestName** | **string**|  | 
  **accountIdentifier** | **string**| Account Identifier for the Entity. | 
  **orgIdentifier** | **string**| Organization Identifier for the Entity. | 
  **projectIdentifier** | **string**| Project Identifier for the Entity. | 

### Return type

[**ApplicationsApplicationResourceResponse**](applicationsApplicationResourceResponse.md)

### Authorization

[ApiKey](../README.md#ApiKey)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **AgentApplicationServicePodLogs**
> StreamResultOfApplicationsLogEntry AgentApplicationServicePodLogs(ctx, agentIdentifier, queryName, queryPodName, accountIdentifier, orgIdentifier, projectIdentifier, optional)
PodLogs returns stream of log entries for the specified pod(s).

PodLogs returns stream of log entries for the specified pod(s).

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **agentIdentifier** | **string**| Agent identifier for entity. | 
  **queryName** | **string**|  | 
  **queryPodName** | **string**|  | 
  **accountIdentifier** | **string**| Account Identifier for the Entity. | 
  **orgIdentifier** | **string**| Organization Identifier for the Entity. | 
  **projectIdentifier** | **string**| Project Identifier for the Entity. | 
 **optional** | ***ApplicationsApiAgentApplicationServicePodLogsOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a ApplicationsApiAgentApplicationServicePodLogsOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------






 **queryNamespace** | **optional.String**|  | 
 **queryContainer** | **optional.String**|  | 
 **querySinceSeconds** | **optional.String**|  | 
 **querySinceTimeSeconds** | **optional.String**| Represents seconds of UTC time since Unix epoch 1970-01-01T00:00:00Z. Must be from 0001-01-01T00:00:00Z to 9999-12-31T23:59:59Z inclusive. | 
 **querySinceTimeNanos** | **optional.Int32**| Non-negative fractions of a second at nanosecond resolution. Negative second values with fractions must still have non-negative nanos values that count forward in time. Must be from 0 to 999,999,999 inclusive. This field may be limited in precision depending on context. | 
 **queryTailLines** | **optional.String**|  | 
 **queryFollow** | **optional.Bool**|  | 
 **queryUntilTime** | **optional.String**|  | 
 **queryFilter** | **optional.String**|  | 
 **queryKind** | **optional.String**|  | 
 **queryGroup** | **optional.String**|  | 
 **queryResourceName** | **optional.String**|  | 
 **queryPrevious** | **optional.Bool**|  | 
 **queryAppNamespace** | **optional.String**|  | 
 **queryProject** | **optional.String**|  | 

### Return type

[**StreamResultOfApplicationsLogEntry**](Stream result of applicationsLogEntry.md)

### Authorization

[ApiKey](../README.md#ApiKey)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **AgentApplicationServicePodLogs2**
> StreamResultOfApplicationsLogEntry AgentApplicationServicePodLogs2(ctx, agentIdentifier, queryName, accountIdentifier, orgIdentifier, projectIdentifier, optional)
PodLogs returns stream of log entries for the specified pod(s).

PodLogs returns stream of log entries for the specified pod(s).

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **agentIdentifier** | **string**| Agent identifier for entity. | 
  **queryName** | **string**|  | 
  **accountIdentifier** | **string**| Account Identifier for the Entity. | 
  **orgIdentifier** | **string**| Organization Identifier for the Entity. | 
  **projectIdentifier** | **string**| Project Identifier for the Entity. | 
 **optional** | ***ApplicationsApiAgentApplicationServicePodLogs2Opts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a ApplicationsApiAgentApplicationServicePodLogs2Opts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------





 **queryNamespace** | **optional.String**|  | 
 **queryPodName** | **optional.String**|  | 
 **queryContainer** | **optional.String**|  | 
 **querySinceSeconds** | **optional.String**|  | 
 **querySinceTimeSeconds** | **optional.String**| Represents seconds of UTC time since Unix epoch 1970-01-01T00:00:00Z. Must be from 0001-01-01T00:00:00Z to 9999-12-31T23:59:59Z inclusive. | 
 **querySinceTimeNanos** | **optional.Int32**| Non-negative fractions of a second at nanosecond resolution. Negative second values with fractions must still have non-negative nanos values that count forward in time. Must be from 0 to 999,999,999 inclusive. This field may be limited in precision depending on context. | 
 **queryTailLines** | **optional.String**|  | 
 **queryFollow** | **optional.Bool**|  | 
 **queryUntilTime** | **optional.String**|  | 
 **queryFilter** | **optional.String**|  | 
 **queryKind** | **optional.String**|  | 
 **queryGroup** | **optional.String**|  | 
 **queryResourceName** | **optional.String**|  | 
 **queryPrevious** | **optional.Bool**|  | 
 **queryAppNamespace** | **optional.String**|  | 
 **queryProject** | **optional.String**|  | 

### Return type

[**StreamResultOfApplicationsLogEntry**](Stream result of applicationsLogEntry.md)

### Authorization

[ApiKey](../README.md#ApiKey)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **AgentApplicationServiceResourceTree**
> ApplicationsApplicationTree AgentApplicationServiceResourceTree(ctx, agentIdentifier, queryName, accountIdentifier, orgIdentifier, projectIdentifier, optional)
ResourceTree returns resource tree

ResourceTree returns resource tree.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **agentIdentifier** | **string**| Agent identifier for entity. | 
  **queryName** | **string**|  | 
  **accountIdentifier** | **string**| Account Identifier for the Entity. | 
  **orgIdentifier** | **string**| Organization Identifier for the Entity. | 
  **projectIdentifier** | **string**| Project Identifier for the Entity. | 
 **optional** | ***ApplicationsApiAgentApplicationServiceResourceTreeOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a ApplicationsApiAgentApplicationServiceResourceTreeOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------





 **queryApplicationName** | **optional.String**|  | 
 **queryNamespace** | **optional.String**|  | 
 **queryName** | **optional.String**|  | 
 **queryVersion** | **optional.String**|  | 
 **queryGroup** | **optional.String**|  | 
 **queryKind** | **optional.String**|  | 
 **queryAppNamespace** | **optional.String**|  | 
 **queryProject** | **optional.String**|  | 

### Return type

[**ApplicationsApplicationTree**](applicationsApplicationTree.md)

### Authorization

[ApiKey](../README.md#ApiKey)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **AgentApplicationServiceRevisionMetadata**
> RepositoriesRevisionMetadata AgentApplicationServiceRevisionMetadata(ctx, agentIdentifier, queryName, queryRevision, accountIdentifier, orgIdentifier, projectIdentifier)
Get the meta-data (author, date, tags, message) for a specific revision of the application

RevisionMetadata returns metadata for a specific revision of the application.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **agentIdentifier** | **string**| Agent identifier for entity. | 
  **queryName** | **string**| the application&#x27;s name | 
  **queryRevision** | **string**| the revision of the app | 
  **accountIdentifier** | **string**| Account Identifier for the Entity. | 
  **orgIdentifier** | **string**| Organization Identifier for the Entity. | 
  **projectIdentifier** | **string**| Project Identifier for the Entity. | 

### Return type

[**RepositoriesRevisionMetadata**](repositoriesRevisionMetadata.md)

### Authorization

[ApiKey](../README.md#ApiKey)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **AgentApplicationServiceRollback**
> Servicev1Application AgentApplicationServiceRollback(ctx, body, agentIdentifier, requestName, accountIdentifier, orgIdentifier, projectIdentifier)
Rollback syncs an application to its target state Harness Event type (rollback)

Rollback syncs an application to its target state Harness Event type (rollback).

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**ApplicationsApplicationRollbackRequest**](ApplicationsApplicationRollbackRequest.md)|  | 
  **agentIdentifier** | **string**| Agent identifier for entity. | 
  **requestName** | **string**|  | 
  **accountIdentifier** | **string**| Account Identifier for the Entity. | 
  **orgIdentifier** | **string**| Organization Identifier for the Entity. | 
  **projectIdentifier** | **string**| Project Identifier for the Entity. | 

### Return type

[**Servicev1Application**](servicev1Application.md)

### Authorization

[ApiKey](../README.md#ApiKey)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **AgentApplicationServiceRunResourceAction**
> ApplicationsApplicationResponse AgentApplicationServiceRunResourceAction(ctx, body, agentIdentifier, requestName, accountIdentifier, orgIdentifier, projectIdentifier)
RunResourceAction run resource action

RunResourceAction run resource action.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**ApplicationsResourceActionRunRequest**](ApplicationsResourceActionRunRequest.md)|  | 
  **agentIdentifier** | **string**| Agent identifier for entity. | 
  **requestName** | **string**|  | 
  **accountIdentifier** | **string**| Account Identifier for the Entity. | 
  **orgIdentifier** | **string**| Organization Identifier for the Entity. | 
  **projectIdentifier** | **string**| Project Identifier for the Entity. | 

### Return type

[**ApplicationsApplicationResponse**](applicationsApplicationResponse.md)

### Authorization

[ApiKey](../README.md#ApiKey)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **AgentApplicationServiceSync**
> Servicev1Application AgentApplicationServiceSync(ctx, body, agentIdentifier, requestName, accountIdentifier, orgIdentifier, projectIdentifier)
Sync syncs an application to its target state Harness Event type (deploy)

Delete deletes an application.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**ApplicationsApplicationSyncRequest**](ApplicationsApplicationSyncRequest.md)|  | 
  **agentIdentifier** | **string**| Agent identifier for entity. | 
  **requestName** | **string**|  | 
  **accountIdentifier** | **string**| Account Identifier for the Entity. | 
  **orgIdentifier** | **string**| Organization Identifier for the Entity. | 
  **projectIdentifier** | **string**| Project Identifier for the Entity. | 

### Return type

[**Servicev1Application**](servicev1Application.md)

### Authorization

[ApiKey](../README.md#ApiKey)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **AgentApplicationServiceTerminateOperation**
> ApplicationsOperationTerminateResponse AgentApplicationServiceTerminateOperation(ctx, agentIdentifier, requestName, accountIdentifier, orgIdentifier, projectIdentifier)
TerminateOperation terminates the currently running operation

TerminateOperation terminates the currently running operation.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **agentIdentifier** | **string**| Agent identifier for entity. | 
  **requestName** | **string**|  | 
  **accountIdentifier** | **string**| Account Identifier for the Entity. | 
  **orgIdentifier** | **string**| Organization Identifier for the Entity. | 
  **projectIdentifier** | **string**| Project Identifier for the Entity. | 

### Return type

[**ApplicationsOperationTerminateResponse**](applicationsOperationTerminateResponse.md)

### Authorization

[ApiKey](../README.md#ApiKey)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **AgentApplicationServiceUpdate**
> Servicev1Application AgentApplicationServiceUpdate(ctx, body, agentIdentifier, requestApplicationMetadataName, accountIdentifier, orgIdentifier, projectIdentifier, optional)
Update updates an application

Update updates an application.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**ApplicationsApplicationUpdateRequest**](ApplicationsApplicationUpdateRequest.md)|  | 
  **agentIdentifier** | **string**| Agent identifier for entity. | 
  **requestApplicationMetadataName** | **string**| Name must be unique within a namespace. Is required when creating resources, although some resources may allow a client to request the generation of an appropriate name automatically. Name is primarily intended for creation idempotence and configuration definition. Cannot be updated. More info: http://kubernetes.io/docs/user-guide/identifiers#names +optional | 
  **accountIdentifier** | **string**| Account Identifier for the Entity. | 
  **orgIdentifier** | **string**| Organization Identifier for the Entity. | 
  **projectIdentifier** | **string**| Project Identifier for the Entity. | 
 **optional** | ***ApplicationsApiAgentApplicationServiceUpdateOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a ApplicationsApiAgentApplicationServiceUpdateOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------






 **clusterIdentifier** | **optional.**|  | 
 **repoIdentifier** | **optional.**|  | 
 **skipRepoValidation** | **optional.**|  | 
 **repoIdentifiers** | [**optional.Interface of []string**](string.md)|  | 

### Return type

[**Servicev1Application**](servicev1Application.md)

### Authorization

[ApiKey](../README.md#ApiKey)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **AgentApplicationServiceUpdateSpec**
> ApplicationsApplicationSpec AgentApplicationServiceUpdateSpec(ctx, body, agentIdentifier, requestName, accountIdentifier, orgIdentifier, projectIdentifier)
UpdateSpec updates an application spec

UpdateSpec updates an application spec.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**ApplicationsApplicationUpdateSpecRequest**](ApplicationsApplicationUpdateSpecRequest.md)|  | 
  **agentIdentifier** | **string**| Agent identifier for entity. | 
  **requestName** | **string**|  | 
  **accountIdentifier** | **string**| Account Identifier for the Entity. | 
  **orgIdentifier** | **string**| Organization Identifier for the Entity. | 
  **projectIdentifier** | **string**| Project Identifier for the Entity. | 

### Return type

[**ApplicationsApplicationSpec**](applicationsApplicationSpec.md)

### Authorization

[ApiKey](../README.md#ApiKey)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **AgentApplicationServiceWatch**
> StreamResultOfApplicationsApplicationWatchEvent AgentApplicationServiceWatch(ctx, agentIdentifier, accountIdentifier, orgIdentifier, projectIdentifier, optional)
Watch returns stream of application change events

Watch returns stream of application change events.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **agentIdentifier** | **string**| Agent identifier for entity. | 
  **accountIdentifier** | **string**| Account Identifier for the Entity. | 
  **orgIdentifier** | **string**| Organization Identifier for the Entity. | 
  **projectIdentifier** | **string**| Project Identifier for the Entity. | 
 **optional** | ***ApplicationsApiAgentApplicationServiceWatchOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a ApplicationsApiAgentApplicationServiceWatchOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------




 **queryName** | **optional.String**| the application&#x27;s name. | 
 **queryRefresh** | **optional.String**| forces application reconciliation if set to true. | 
 **queryProject** | [**optional.Interface of []string**](string.md)| the project names to restrict returned list applications. | 
 **queryResourceVersion** | **optional.String**| when specified with a watch call, shows changes that occur after that particular version of a resource. | 
 **querySelector** | **optional.String**| the selector to to restrict returned list to applications only with matched labels. | 
 **queryRepo** | **optional.String**| the repoURL to restrict returned list applications. | 
 **queryAppNamespace** | **optional.String**|  | 

### Return type

[**StreamResultOfApplicationsApplicationWatchEvent**](Stream result of applicationsApplicationWatchEvent.md)

### Authorization

[ApiKey](../README.md#ApiKey)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **AgentApplicationServiceWatchResourceTree**
> StreamResultOfApplicationsApplicationTree AgentApplicationServiceWatchResourceTree(ctx, agentIdentifier, queryApplicationName, accountIdentifier, orgIdentifier, projectIdentifier, optional)
WatchResourceTree returns stream of application resource tree

WatchResourceTree returns stream of application resource tree.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **agentIdentifier** | **string**| Agent identifier for entity. | 
  **queryApplicationName** | **string**|  | 
  **accountIdentifier** | **string**| Account Identifier for the Entity. | 
  **orgIdentifier** | **string**| Organization Identifier for the Entity. | 
  **projectIdentifier** | **string**| Project Identifier for the Entity. | 
 **optional** | ***ApplicationsApiAgentApplicationServiceWatchResourceTreeOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a ApplicationsApiAgentApplicationServiceWatchResourceTreeOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------





 **queryNamespace** | **optional.String**|  | 
 **queryName** | **optional.String**|  | 
 **queryVersion** | **optional.String**|  | 
 **queryGroup** | **optional.String**|  | 
 **queryKind** | **optional.String**|  | 
 **queryAppNamespace** | **optional.String**|  | 
 **queryProject** | **optional.String**|  | 

### Return type

[**StreamResultOfApplicationsApplicationTree**](Stream result of applicationsApplicationTree.md)

### Authorization

[ApiKey](../README.md#ApiKey)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ApplicationServiceExists**
> bool ApplicationServiceExists(ctx, name, accountIdentifier, orgIdentifier, projectIdentifier, optional)
Checks whether an app with the given name exists

Checks whether an app with the given name exists

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **name** | **string**|  | 
  **accountIdentifier** | **string**| Account Identifier for the Entity. | 
  **orgIdentifier** | **string**| Organization Identifier for the Entity. | 
  **projectIdentifier** | **string**| Project Identifier for the Entity. | 
 **optional** | ***ApplicationsApiApplicationServiceExistsOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a ApplicationsApiApplicationServiceExistsOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------




 **agentIdentifier** | **optional.String**| Agent identifier for entity. | 

### Return type

**bool**

### Authorization

[ApiKey](../README.md#ApiKey)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ApplicationServiceListAppSync**
> V1ApplicationSyncStatuslist ApplicationServiceListAppSync(ctx, body)
List returns list of application sync status

List returns list of application sync status

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**V1ApplicationSyncStatusQuery**](V1ApplicationSyncStatusQuery.md)|  | 

### Return type

[**V1ApplicationSyncStatuslist**](v1ApplicationSyncStatuslist.md)

### Authorization

[ApiKey](../README.md#ApiKey)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ApplicationServiceListAttributes**
> V1ApplicationAttributesList ApplicationServiceListAttributes(ctx, optional)


Returns list of attributes for applications

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***ApplicationsApiApplicationServiceListAttributesOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a ApplicationsApiApplicationServiceListAttributesOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **accountIdentifier** | **optional.String**| Account Identifier for the Entity. | 
 **orgIdentifier** | **optional.String**| Organization Identifier for the Entity. | 
 **projectIdentifier** | **optional.String**| Project Identifier for the Entity. | 
 **applicationIdentifiers** | [**optional.Interface of []string**](string.md)| List of application identifiers | 

### Return type

[**V1ApplicationAttributesList**](v1ApplicationAttributesList.md)

### Authorization

[ApiKey](../README.md#ApiKey)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ApplicationServiceListNs**
> Gitopsservicev1NamespaceList ApplicationServiceListNs(ctx, body)
ListNs returns list of namespaces for a query.

List returns list of namespaces for a specific agent.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**Servicev1ApplicationQuery**](Servicev1ApplicationQuery.md)|  | 

### Return type

[**Gitopsservicev1NamespaceList**](gitopsservicev1NamespaceList.md)

### Authorization

[ApiKey](../README.md#ApiKey)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

