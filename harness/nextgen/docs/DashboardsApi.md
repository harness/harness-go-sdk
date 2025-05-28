# nextgen{{classname}}

All URIs are relative to *https://app.harness.io*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CloneDashboard**](DashboardsApi.md#CloneDashboard) | **Post** /dashboard/clone | 
[**DashboardServiceGetDashboardOverview**](DashboardsApi.md#DashboardServiceGetDashboardOverview) | **Get** /gitops/api/v1/dashboard/overview | Get dashboard overview
[**DashboardServiceRecentlyCreatedCounts**](DashboardsApi.md#DashboardServiceRecentlyCreatedCounts) | **Get** /gitops/api/v1/dashboard/counts | Get recently created app, cluster, repo counts
[**GetDashboardData**](DashboardsApi.md#GetDashboardData) | **Get** /dashboard/api/dashboards/{dashboard_id}/download | Download data within a Dashboard
[**GetDashboardDownloadCsv**](DashboardsApi.md#GetDashboardDownloadCsv) | **Get** /dashboard/download/dashboards/{dashboard_id}/csv | Download a Dashboard CSV
[**GetDashboardDownloadPdf**](DashboardsApi.md#GetDashboardDownloadPdf) | **Get** /dashboard/download/dashboards/{dashboard_id}/pdf | Download a Dashboard PDF
[**GetDashboardElements**](DashboardsApi.md#GetDashboardElements) | **Get** /dashboard/dashboards/{dashboard_id}/elements | 
[**GetDashboardFilters**](DashboardsApi.md#GetDashboardFilters) | **Get** /dashboard/dashboards/{dashboard_id}/filters | 
[**SchedulesRunOnce**](DashboardsApi.md#SchedulesRunOnce) | **Post** /dashboard/schedules/run_once | Runs a schedule delivery once that is then immediately sent via email to recipients
[**UpdateDashboardFilter**](DashboardsApi.md#UpdateDashboardFilter) | **Patch** /dashboard/dashboards/{dashboard_id}/filters/{filter_id} | 

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

# **DashboardServiceGetDashboardOverview**
> V1DashboardOverview DashboardServiceGetDashboardOverview(ctx, optional)
Get dashboard overview

GetDashboardOverview provides an overview of the dashboard including key metrics.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***DashboardsApiDashboardServiceGetDashboardOverviewOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a DashboardsApiDashboardServiceGetDashboardOverviewOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **agentIdentifier** | **optional.String**| Agent identifier for entity. | 
 **projectIdentifier** | **optional.String**| Project Identifier for the Entity. | 
 **orgIdentifier** | **optional.String**| Organization Identifier for the Entity. | 
 **accountIdentifier** | **optional.String**| Account Identifier for the Entity. | 
 **disasterRecoveryIdentifier** | **optional.String**|  | 

### Return type

[**V1DashboardOverview**](v1DashboardOverview.md)

### Authorization

[x-api-key](../README.md#x-api-key)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **DashboardServiceRecentlyCreatedCounts**
> V1RecentlyCreatedOverview DashboardServiceRecentlyCreatedCounts(ctx, optional)
Get recently created app, cluster, repo counts

RecentlyCreatedCounts lists the count of clusters, repositories, and applications created within a specified time period.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***DashboardsApiDashboardServiceRecentlyCreatedCountsOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a DashboardsApiDashboardServiceRecentlyCreatedCountsOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **projectIdentifier** | **optional.String**| Project Identifier for the Entity. | 
 **orgIdentifier** | **optional.String**| Organization Identifier for the Entity. | 
 **accountIdentifier** | **optional.String**| Account Identifier for the Entity. | 
 **endTime** | **optional.Int32**|  | 
 **startTime** | **optional.Int32**|  | 

### Return type

[**V1RecentlyCreatedOverview**](v1RecentlyCreatedOverview.md)

### Authorization

[x-api-key](../README.md#x-api-key)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetDashboardData**
> DashboardDownloadResponse GetDashboardData(ctx, dashboardId, fileType, optional)
Download data within a Dashboard

Download the data of all tiles within a Dashboard.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **dashboardId** | **string**|  | 
  **fileType** | **string**|  | 
 **optional** | ***DashboardsApiGetDashboardDataOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a DashboardsApiGetDashboardDataOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **filters** | **optional.String**|  | 

### Return type

[**DashboardDownloadResponse**](DashboardDownloadResponse.md)

### Authorization

[x-api-key](../README.md#x-api-key)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetDashboardDownloadCsv**
> DownloadCsvDashboardResponse GetDashboardDownloadCsv(ctx, dashboardId, optional)
Download a Dashboard CSV

Download a Dashboard in CSV format.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **dashboardId** | **string**|  | 
 **optional** | ***DashboardsApiGetDashboardDownloadCsvOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a DashboardsApiGetDashboardDownloadCsvOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **accountId** | **optional.String**|  | 
 **filters** | **optional.String**|  | 
 **expandedTables** | **optional.Bool**|  | 

### Return type

[**DownloadCsvDashboardResponse**](DownloadCsvDashboardResponse.md)

### Authorization

[x-api-key](../README.md#x-api-key)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetDashboardDownloadPdf**
> DownloadPdfDashboardResponse GetDashboardDownloadPdf(ctx, dashboardId, optional)
Download a Dashboard PDF

Download a Dashboard in PDF format.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **dashboardId** | **string**|  | 
 **optional** | ***DashboardsApiGetDashboardDownloadPdfOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a DashboardsApiGetDashboardDownloadPdfOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **accountId** | **optional.String**|  | 
 **filters** | **optional.String**|  | 
 **expandedTables** | **optional.Bool**|  | 
 **landscape** | **optional.Bool**|  | 
 **paperSize** | **optional.String**|  | 
 **style** | **optional.String**|  | 

### Return type

[**DownloadPdfDashboardResponse**](DownloadPdfDashboardResponse.md)

### Authorization

[x-api-key](../README.md#x-api-key)

### HTTP request headers

 - **Content-Type**: Not defined
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

# **SchedulesRunOnce**
> ScheduleReportResponse SchedulesRunOnce(ctx, body, optional)
Runs a schedule delivery once that is then immediately sent via email to recipients

Run a schedule delivery once

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**ScheduleReportRequestBody**](ScheduleReportRequestBody.md)| Run a Dashboard Schedule Delivery once | 
 **optional** | ***DashboardsApiSchedulesRunOnceOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a DashboardsApiSchedulesRunOnceOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **accountId** | **optional.**|  | 

### Return type

[**ScheduleReportResponse**](ScheduleReportResponse.md)

### Authorization

[x-api-key](../README.md#x-api-key)

### HTTP request headers

 - **Content-Type**: application/json
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

