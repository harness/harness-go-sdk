# nextgen{{classname}}

All URIs are relative to *https://app.harness.io*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CloneDashboard**](DashboardsApi.md#CloneDashboard) | **Post** /dashboard/clone | Clones a Dashboard

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
