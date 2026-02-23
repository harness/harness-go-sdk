# {{classname}}

All URIs are relative to *https://app.harness.io/gateway*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateAlert**](AutoStoppingAlertsApi.md#CreateAlert) | **Post** /accounts/{account_id}/alerts | Create Alert
[**DeleteAlert**](AutoStoppingAlertsApi.md#DeleteAlert) | **Delete** /accounts/{account_id}/alerts/{alert_id} | Delete Alert
[**GetAlert**](AutoStoppingAlertsApi.md#GetAlert) | **Get** /accounts/{account_id}/alerts/{alert_id} | Get Alert
[**ListAlerts**](AutoStoppingAlertsApi.md#ListAlerts) | **Post** /accounts/{account_id}/alerts/list | List Alerts
[**UpdateAlert**](AutoStoppingAlertsApi.md#UpdateAlert) | **Put** /accounts/{account_id}/alerts | Update Alert

# **CreateAlert**
> AlertCreateResponse CreateAlert(ctx, body, accountId, accountIdentifier)
Create Alert

Creates a new alert configuration for AutoStopping events. Alerts can notify users via email or Slack when specific events occur, such as warmup failures, cooldown failures, or rule lifecycle changes.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**AlertRequest**](AlertRequest.md)| Alert configuration to create | 
  **accountId** | **string**| Unique identifier of the Harness account | 
  **accountIdentifier** | **string**| Account Identifier for authentication and authorization | 

### Return type

[**AlertCreateResponse**](AlertCreateResponse.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **DeleteAlert**
> AlertMessageResponse DeleteAlert(ctx, accountId, alertId, accountIdentifier)
Delete Alert

Deletes an alert by its ID. This is a soft delete operation.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **accountId** | **string**| Unique identifier of the Harness account | 
  **alertId** | **string**| Unique identifier of the alert to delete | 
  **accountIdentifier** | **string**| Account Identifier for authentication and authorization | 

### Return type

[**AlertMessageResponse**](AlertMessageResponse.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetAlert**
> AlertResponse GetAlert(ctx, accountId, alertId, accountIdentifier)
Get Alert

Retrieves the details of a specific alert by its ID.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **accountId** | **string**| Unique identifier of the Harness account | 
  **alertId** | **string**| Unique identifier of the alert to retrieve | 
  **accountIdentifier** | **string**| Account Identifier for authentication and authorization | 

### Return type

[**AlertResponse**](AlertResponse.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ListAlerts**
> AlertListResponse ListAlerts(ctx, body, accountId, accountIdentifier)
List Alerts

Retrieves a paginated list of alerts matching the specified search criteria. Supports filtering by entity type, entity IDs, enabled status, events, and name search.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**AlertSearchRequest**](AlertSearchRequest.md)| Search criteria for filtering alerts | 
  **accountId** | **string**| Unique identifier of the Harness account | 
  **accountIdentifier** | **string**| Account Identifier for authentication and authorization | 

### Return type

[**AlertListResponse**](AlertListResponse.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **UpdateAlert**
> AlertMessageResponse UpdateAlert(ctx, body, accountId, accountIdentifier)
Update Alert

Updates an existing alert configuration. The alert ID must be provided in the request body.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**AlertUpdateRequest**](AlertUpdateRequest.md)| Alert configuration to update. The alert ID is required. | 
  **accountId** | **string**| Unique identifier of the Harness account | 
  **accountIdentifier** | **string**| Account Identifier for authentication and authorization | 

### Return type

[**AlertMessageResponse**](AlertMessageResponse.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

