# code{{classname}}

All URIs are relative to */api/v1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CodeownersPullReq**](PullreqApi.md#CodeownersPullReq) | **Get** /repos/{repo_ref}/pullreq/{pullreq_number}/codeowners | 
[**CommentCreatePullReq**](PullreqApi.md#CommentCreatePullReq) | **Post** /repos/{repo_ref}/pullreq/{pullreq_number}/comments | 
[**CommentDeletePullReq**](PullreqApi.md#CommentDeletePullReq) | **Delete** /repos/{repo_ref}/pullreq/{pullreq_number}/comments/{pullreq_comment_id} | 
[**CommentStatusPullReq**](PullreqApi.md#CommentStatusPullReq) | **Put** /repos/{repo_ref}/pullreq/{pullreq_number}/comments/{pullreq_comment_id}/status | 
[**CommentUpdatePullReq**](PullreqApi.md#CommentUpdatePullReq) | **Patch** /repos/{repo_ref}/pullreq/{pullreq_number}/comments/{pullreq_comment_id} | 
[**CreatePullReq**](PullreqApi.md#CreatePullReq) | **Post** /repos/{repo_ref}/pullreq | 
[**FileViewAddPullReq**](PullreqApi.md#FileViewAddPullReq) | **Put** /repos/{repo_ref}/pullreq/{pullreq_number}/file-views | 
[**FileViewDeletePullReq**](PullreqApi.md#FileViewDeletePullReq) | **Delete** /repos/{repo_ref}/pullreq/{pullreq_number}/file-views/{file_path} | 
[**FileViewListPullReq**](PullreqApi.md#FileViewListPullReq) | **Get** /repos/{repo_ref}/pullreq/{pullreq_number}/file-views | 
[**GetPullReq**](PullreqApi.md#GetPullReq) | **Get** /repos/{repo_ref}/pullreq/{pullreq_number} | 
[**ListPullReq**](PullreqApi.md#ListPullReq) | **Get** /repos/{repo_ref}/pullreq | 
[**ListPullReqActivities**](PullreqApi.md#ListPullReqActivities) | **Get** /repos/{repo_ref}/pullreq/{pullreq_number}/activities | 
[**ListPullReqCommits**](PullreqApi.md#ListPullReqCommits) | **Get** /repos/{repo_ref}/pullreq/{pullreq_number}/commits | 
[**MergePullReqOp**](PullreqApi.md#MergePullReqOp) | **Post** /repos/{repo_ref}/pullreq/{pullreq_number}/merge | 
[**PullReqMetaData**](PullreqApi.md#PullReqMetaData) | **Get** /repos/{repo_ref}/pullreq/{pullreq_number}/metadata | 
[**ReviewSubmitPullReq**](PullreqApi.md#ReviewSubmitPullReq) | **Post** /repos/{repo_ref}/pullreq/{pullreq_number}/reviews | 
[**ReviewerAddPullReq**](PullreqApi.md#ReviewerAddPullReq) | **Put** /repos/{repo_ref}/pullreq/{pullreq_number}/reviewers | 
[**ReviewerDeletePullReq**](PullreqApi.md#ReviewerDeletePullReq) | **Delete** /repos/{repo_ref}/pullreq/{pullreq_number}/reviewers/{pullreq_reviewer_id} | 
[**ReviewerListPullReq**](PullreqApi.md#ReviewerListPullReq) | **Get** /repos/{repo_ref}/pullreq/{pullreq_number}/reviewers | 
[**StatePullReq**](PullreqApi.md#StatePullReq) | **Post** /repos/{repo_ref}/pullreq/{pullreq_number}/state | 
[**UpdatePullReq**](PullreqApi.md#UpdatePullReq) | **Patch** /repos/{repo_ref}/pullreq/{pullreq_number} | 

# **CodeownersPullReq**
> TypesCodeOwnerEvaluation CodeownersPullReq(ctx, repoRef, pullreqNumber)


### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **repoRef** | **string**|  | 
  **pullreqNumber** | **int32**|  | 

### Return type

[**TypesCodeOwnerEvaluation**](TypesCodeOwnerEvaluation.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **CommentCreatePullReq**
> TypesPullReqActivity CommentCreatePullReq(ctx, repoRef, pullreqNumber, optional)


### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **repoRef** | **string**|  | 
  **pullreqNumber** | **int32**|  | 
 **optional** | ***PullreqApiCommentCreatePullReqOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a PullreqApiCommentCreatePullReqOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **body** | [**optional.Interface of OpenapiCommentCreatePullReqRequest**](OpenapiCommentCreatePullReqRequest.md)|  | 

### Return type

[**TypesPullReqActivity**](TypesPullReqActivity.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **CommentDeletePullReq**
> CommentDeletePullReq(ctx, repoRef, pullreqNumber, pullreqCommentId)


### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **repoRef** | **string**|  | 
  **pullreqNumber** | **int32**|  | 
  **pullreqCommentId** | **int32**|  | 

### Return type

 (empty response body)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **CommentStatusPullReq**
> TypesPullReqActivity CommentStatusPullReq(ctx, repoRef, pullreqNumber, pullreqCommentId, optional)


### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **repoRef** | **string**|  | 
  **pullreqNumber** | **int32**|  | 
  **pullreqCommentId** | **int32**|  | 
 **optional** | ***PullreqApiCommentStatusPullReqOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a PullreqApiCommentStatusPullReqOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **body** | [**optional.Interface of OpenapiCommentStatusPullReqRequest**](OpenapiCommentStatusPullReqRequest.md)|  | 

### Return type

[**TypesPullReqActivity**](TypesPullReqActivity.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **CommentUpdatePullReq**
> TypesPullReqActivity CommentUpdatePullReq(ctx, repoRef, pullreqNumber, pullreqCommentId, optional)


### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **repoRef** | **string**|  | 
  **pullreqNumber** | **int32**|  | 
  **pullreqCommentId** | **int32**|  | 
 **optional** | ***PullreqApiCommentUpdatePullReqOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a PullreqApiCommentUpdatePullReqOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **body** | [**optional.Interface of OpenapiCommentUpdatePullReqRequest**](OpenapiCommentUpdatePullReqRequest.md)|  | 

### Return type

[**TypesPullReqActivity**](TypesPullReqActivity.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **CreatePullReq**
> TypesPullReq CreatePullReq(ctx, repoRef, optional)


### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **repoRef** | **string**|  | 
 **optional** | ***PullreqApiCreatePullReqOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a PullreqApiCreatePullReqOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **body** | [**optional.Interface of OpenapiCreatePullReqRequest**](OpenapiCreatePullReqRequest.md)|  | 

### Return type

[**TypesPullReq**](TypesPullReq.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **FileViewAddPullReq**
> TypesPullReqFileView FileViewAddPullReq(ctx, repoRef, pullreqNumber, optional)


### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **repoRef** | **string**|  | 
  **pullreqNumber** | **int32**|  | 
 **optional** | ***PullreqApiFileViewAddPullReqOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a PullreqApiFileViewAddPullReqOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **body** | [**optional.Interface of OpenapiFileViewAddPullReqRequest**](OpenapiFileViewAddPullReqRequest.md)|  | 

### Return type

[**TypesPullReqFileView**](TypesPullReqFileView.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **FileViewDeletePullReq**
> FileViewDeletePullReq(ctx, repoRef, pullreqNumber, filePath)


### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **repoRef** | **string**|  | 
  **pullreqNumber** | **int32**|  | 
  **filePath** | **string**|  | 

### Return type

 (empty response body)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **FileViewListPullReq**
> []TypesPullReqFileView FileViewListPullReq(ctx, repoRef, pullreqNumber)


### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **repoRef** | **string**|  | 
  **pullreqNumber** | **int32**|  | 

### Return type

[**[]TypesPullReqFileView**](TypesPullReqFileView.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetPullReq**
> TypesPullReq GetPullReq(ctx, repoRef, pullreqNumber)


### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **repoRef** | **string**|  | 
  **pullreqNumber** | **int32**|  | 

### Return type

[**TypesPullReq**](TypesPullReq.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ListPullReq**
> []TypesPullReq ListPullReq(ctx, repoRef, optional)


### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **repoRef** | **string**|  | 
 **optional** | ***PullreqApiListPullReqOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a PullreqApiListPullReqOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **state** | [**optional.Interface of []string**](string.md)| The state of the pull requests to include in the result. | 
 **sourceRepoRef** | **optional.String**| Source repository ref of the pull requests. | 
 **sourceBranch** | **optional.String**| Source branch of the pull requests. | 
 **targetBranch** | **optional.String**| Target branch of the pull requests. | 
 **query** | **optional.String**| The substring by which the pull requests are filtered. | 
 **createdBy** | **optional.Int32**| The principal ID who created pull requests. | 
 **order** | **optional.String**| The order of the output. | [default to asc]
 **sort** | **optional.String**| The data by which the pull requests are sorted. | [default to number]
 **page** | **optional.Int32**| The page to return. | [default to 1]
 **limit** | **optional.Int32**| The maximum number of results to return. | [default to 30]

### Return type

[**[]TypesPullReq**](TypesPullReq.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ListPullReqActivities**
> []TypesPullReqActivity ListPullReqActivities(ctx, repoRef, pullreqNumber, optional)


### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **repoRef** | **string**|  | 
  **pullreqNumber** | **int32**|  | 
 **optional** | ***PullreqApiListPullReqActivitiesOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a PullreqApiListPullReqActivitiesOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **kind** | [**optional.Interface of []string**](string.md)| The kind of the pull request activity to include in the result. | 
 **type_** | [**optional.Interface of []string**](string.md)| The type of the pull request activity to include in the result. | 
 **after** | **optional.Int32**| The result should contain only entries created at and after this timestamp (unix millis). | 
 **before** | **optional.Int32**| The result should contain only entries created before this timestamp (unix millis). | 
 **limit** | **optional.Int32**| The maximum number of results to return. | [default to 30]

### Return type

[**[]TypesPullReqActivity**](TypesPullReqActivity.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ListPullReqCommits**
> []TypesCommit ListPullReqCommits(ctx, repoRef, pullreqNumber, optional)


### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **repoRef** | **string**|  | 
  **pullreqNumber** | **int32**|  | 
 **optional** | ***PullreqApiListPullReqCommitsOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a PullreqApiListPullReqCommitsOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **page** | **optional.Int32**| The page to return. | [default to 1]
 **limit** | **optional.Int32**| The maximum number of results to return. | [default to 30]

### Return type

[**[]TypesCommit**](TypesCommit.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **MergePullReqOp**
> TypesMergeResponse MergePullReqOp(ctx, repoRef, pullreqNumber, optional)


### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **repoRef** | **string**|  | 
  **pullreqNumber** | **int32**|  | 
 **optional** | ***PullreqApiMergePullReqOpOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a PullreqApiMergePullReqOpOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **body** | [**optional.Interface of OpenapiMergePullReq**](OpenapiMergePullReq.md)|  | 

### Return type

[**TypesMergeResponse**](TypesMergeResponse.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **PullReqMetaData**
> TypesPullReqStats PullReqMetaData(ctx, repoRef, pullreqNumber)


### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **repoRef** | **string**|  | 
  **pullreqNumber** | **int32**|  | 

### Return type

[**TypesPullReqStats**](TypesPullReqStats.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ReviewSubmitPullReq**
> ReviewSubmitPullReq(ctx, repoRef, pullreqNumber, optional)


### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **repoRef** | **string**|  | 
  **pullreqNumber** | **int32**|  | 
 **optional** | ***PullreqApiReviewSubmitPullReqOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a PullreqApiReviewSubmitPullReqOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **body** | [**optional.Interface of OpenapiReviewSubmitPullReqRequest**](OpenapiReviewSubmitPullReqRequest.md)|  | 

### Return type

 (empty response body)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ReviewerAddPullReq**
> TypesPullReqReviewer ReviewerAddPullReq(ctx, repoRef, pullreqNumber, optional)


### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **repoRef** | **string**|  | 
  **pullreqNumber** | **int32**|  | 
 **optional** | ***PullreqApiReviewerAddPullReqOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a PullreqApiReviewerAddPullReqOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **body** | [**optional.Interface of OpenapiReviewerAddPullReqRequest**](OpenapiReviewerAddPullReqRequest.md)|  | 

### Return type

[**TypesPullReqReviewer**](TypesPullReqReviewer.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ReviewerDeletePullReq**
> ReviewerDeletePullReq(ctx, repoRef, pullreqNumber, pullreqReviewerId)


### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **repoRef** | **string**|  | 
  **pullreqNumber** | **int32**|  | 
  **pullreqReviewerId** | **int32**|  | 

### Return type

 (empty response body)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ReviewerListPullReq**
> []TypesPullReqReviewer ReviewerListPullReq(ctx, repoRef, pullreqNumber)


### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **repoRef** | **string**|  | 
  **pullreqNumber** | **int32**|  | 

### Return type

[**[]TypesPullReqReviewer**](TypesPullReqReviewer.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **StatePullReq**
> TypesPullReq StatePullReq(ctx, repoRef, pullreqNumber, optional)


### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **repoRef** | **string**|  | 
  **pullreqNumber** | **int32**|  | 
 **optional** | ***PullreqApiStatePullReqOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a PullreqApiStatePullReqOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **body** | [**optional.Interface of OpenapiStatePullReqRequest**](OpenapiStatePullReqRequest.md)|  | 

### Return type

[**TypesPullReq**](TypesPullReq.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **UpdatePullReq**
> TypesPullReq UpdatePullReq(ctx, repoRef, pullreqNumber, optional)


### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **repoRef** | **string**|  | 
  **pullreqNumber** | **int32**|  | 
 **optional** | ***PullreqApiUpdatePullReqOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a PullreqApiUpdatePullReqOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **body** | [**optional.Interface of OpenapiUpdatePullReqRequest**](OpenapiUpdatePullReqRequest.md)|  | 

### Return type

[**TypesPullReq**](TypesPullReq.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

