# code{{classname}}

All URIs are relative to */api/v1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**RepoArtifactDownload**](UploadApi.md#RepoArtifactDownload) | **Get** /repos/{repo_ref}/uploads/{file_ref} | 
[**RepoArtifactUpload**](UploadApi.md#RepoArtifactUpload) | **Post** /repos/{repo_ref}/uploads | 

# **RepoArtifactDownload**
> RepoArtifactDownload(ctx, repoRef, fileRef)


### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **repoRef** | **string**|  | 
  **fileRef** | **string**|  | 

### Return type

 (empty response body)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **RepoArtifactUpload**
> UploadResult RepoArtifactUpload(ctx, repoRef)


### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **repoRef** | **string**|  | 

### Return type

[**UploadResult**](UploadResult.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

