# nextgen{{classname}}

All URIs are relative to *https://app.harness.io*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CloneDashboard**](DashboardsApi.md#CloneDashboard) | **Post** /dashboard/clone | Clones a Dashboard
[**GetDashboardElements**](DashboardsApi.md#GetDashboardElements) | **Get** /dashboard/dashboards/{dashboard_id}/elements | Gets all the elements within a Dashboard
[**GetDashboardFilters**](DashboardsApi.md#GetDashboardFilters) | **Get** /dashboard/dashboards/{dashboard_id}/filters | Gets all the filters within a Dashboard
[**UpdateDashboardFilter**](DashboardsApi.md#UpdateDashboardFilter) | **Patch** /dashboard/dashboards/{dashboard_id}/filters/{filter_id} | Updates a specified Dashboard Filter

# **CloneDashboard**
> ClonedDashboardResponse CloneDashboard(ctx, body, optional)


Clone a dashboard.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**CloneDashboardRequestBody**](CloneDashboardRequestBody.md)| Clone a Dashboard | 
 **optional** | ***DashboardsApiCloneDashboardOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a DashboardsApiCloneDashboardOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **accountId** | **optional.**|  | 

### Return type

[**ClonedDashboardResponse**](ClonedDashboardResponse.md)

### Authorization

[x-api-key](../README.md#x-api-key)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetDashboardElements**
> GetDashboardElementsResponse GetDashboardElements(ctx, dashboardId, optional)


Get all elements within a dashboard by ID.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **dashboardId** | **string**|  | 
 **optional** | ***DashboardsApiGetDashboardElementsOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a DashboardsApiGetDashboardElementsOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **accountId** | **optional.String**|  | 

### Return type

[**GetDashboardElementsResponse**](GetDashboardElementsResponse.md)

### Authorization

[x-api-key](../README.md#x-api-key)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetDashboardFilters**
> GetDashboardFiltersResponse GetDashboardFilters(ctx, dashboardId, optional)


Get all filters within a dashboard by ID.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **dashboardId** | **string**|  | 
 **optional** | ***DashboardsApiGetDashboardFiltersOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a DashboardsApiGetDashboardFiltersOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **accountId** | **optional.String**|  | 

### Return type

[**GetDashboardFiltersResponse**](GetDashboardFiltersResponse.md)

### Authorization

[x-api-key](../README.md#x-api-key)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **UpdateDashboardFilter**
> GetDashboardFilterResponse UpdateDashboardFilter(ctx, body, dashboardId, filterId)


Update a specific filter within a dashboard by ID.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**UpdateDashboardFilterRequest**](UpdateDashboardFilterRequest.md)|  | 
  **dashboardId** | **string**|  | 
  **filterId** | **string**|  | 

### Return type

[**GetDashboardFilterResponse**](GetDashboardFilterResponse.md)

### Authorization

[x-api-key](../README.md#x-api-key)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

