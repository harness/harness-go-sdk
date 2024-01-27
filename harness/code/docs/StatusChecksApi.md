# code{{classname}}

All URIs are relative to */api/v1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**ListStatusCheckRecent**](StatusChecksApi.md#ListStatusCheckRecent) | **Get** /repos/{repo_ref}/checks/recent | 
[**ListStatusCheckResults**](StatusChecksApi.md#ListStatusCheckResults) | **Get** /repos/{repo_ref}/checks/commits/{commit_sha} | 
[**ReportStatusCheckResults**](StatusChecksApi.md#ReportStatusCheckResults) | **Put** /repos/{repo_ref}/checks/commits/{commit_sha} | 

# **ListStatusCheckRecent**
> []string ListStatusCheckRecent(ctx, repoRef, optional)


### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **repoRef** | **string**|  | 
 **optional** | ***StatusChecksApiListStatusCheckRecentOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a StatusChecksApiListStatusCheckRecentOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **query** | **optional.String**| The substring which is used to filter the status checks by their UID. | 
 **since** | **optional.Int32**| The timestamp (in Unix time millis) since the status checks have been run. | 

### Return type

**[]string**

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ListStatusCheckResults**
> []TypesCheck ListStatusCheckResults(ctx, repoRef, commitSha, optional)


### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **repoRef** | **string**|  | 
  **commitSha** | **string**|  | 
 **optional** | ***StatusChecksApiListStatusCheckResultsOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a StatusChecksApiListStatusCheckResultsOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **page** | **optional.Int32**| The page to return. | [default to 1]
 **limit** | **optional.Int32**| The maximum number of results to return. | [default to 30]
 **query** | **optional.String**| The substring which is used to filter the status checks by their UID. | 

### Return type

[**[]TypesCheck**](TypesCheck.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ReportStatusCheckResults**
> TypesCheck ReportStatusCheckResults(ctx, repoRef, commitSha, optional)


### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **repoRef** | **string**|  | 
  **commitSha** | **string**|  | 
 **optional** | ***StatusChecksApiReportStatusCheckResultsOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a StatusChecksApiReportStatusCheckResultsOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **body** | [**optional.Interface of CommitsCommitShaBody**](CommitsCommitShaBody.md)|  | 

### Return type

[**TypesCheck**](TypesCheck.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

