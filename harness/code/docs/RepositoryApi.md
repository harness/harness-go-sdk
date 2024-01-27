# code{{classname}}

All URIs are relative to */api/v1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CalculateCommitDivergence**](RepositoryApi.md#CalculateCommitDivergence) | **Post** /repos/{repo_ref}/commits/calculate-divergence | 
[**CodeOwnersValidate**](RepositoryApi.md#CodeOwnersValidate) | **Get** /repos/{repo_ref}/codeowners/validate | 
[**CommitFiles**](RepositoryApi.md#CommitFiles) | **Post** /repos/{repo_ref}/commits | 
[**CreateBranch**](RepositoryApi.md#CreateBranch) | **Post** /repos/{repo_ref}/branches | 
[**CreateRepository**](RepositoryApi.md#CreateRepository) | **Post** /repos | 
[**CreateTag**](RepositoryApi.md#CreateTag) | **Post** /repos/{repo_ref}/tags | 
[**DeleteBranch**](RepositoryApi.md#DeleteBranch) | **Delete** /repos/{repo_ref}/branches/{branch_name} | 
[**DeleteRepository**](RepositoryApi.md#DeleteRepository) | **Delete** /repos/{repo_ref} | 
[**DeleteTag**](RepositoryApi.md#DeleteTag) | **Delete** /repos/{repo_ref}/tags/{tag_name} | 
[**DiffStats**](RepositoryApi.md#DiffStats) | **Get** /repos/{repo_ref}/diff-stats/{range} | 
[**FindRepository**](RepositoryApi.md#FindRepository) | **Get** /repos/{repo_ref} | 
[**GetBlame**](RepositoryApi.md#GetBlame) | **Get** /repos/{repo_ref}/blame/{path} | 
[**GetBranch**](RepositoryApi.md#GetBranch) | **Get** /repos/{repo_ref}/branches/{branch_name} | 
[**GetCommit**](RepositoryApi.md#GetCommit) | **Get** /repos/{repo_ref}/commits/{commit_sha} | 
[**GetCommitDiff**](RepositoryApi.md#GetCommitDiff) | **Get** /repos/{repo_ref}/commits/{commit_sha}/diff | 
[**GetContent**](RepositoryApi.md#GetContent) | **Get** /repos/{repo_ref}/content/{path} | 
[**GetRaw**](RepositoryApi.md#GetRaw) | **Get** /repos/{repo_ref}/raw/{path} | 
[**ImportRepository**](RepositoryApi.md#ImportRepository) | **Post** /repos/import | 
[**ListBranches**](RepositoryApi.md#ListBranches) | **Get** /repos/{repo_ref}/branches | 
[**ListCommits**](RepositoryApi.md#ListCommits) | **Get** /repos/{repo_ref}/commits | 
[**ListRepositoryServiceAccounts**](RepositoryApi.md#ListRepositoryServiceAccounts) | **Get** /repos/{repo_ref}/service-accounts | 
[**ListTags**](RepositoryApi.md#ListTags) | **Get** /repos/{repo_ref}/tags | 
[**MergeCheck**](RepositoryApi.md#MergeCheck) | **Post** /repos/{repo_ref}/merge-check/{range} | 
[**MoveRepository**](RepositoryApi.md#MoveRepository) | **Post** /repos/{repo_ref}/move | 
[**PathDetails**](RepositoryApi.md#PathDetails) | **Post** /repos/{repo_ref}/path-details | 
[**RawDiff**](RepositoryApi.md#RawDiff) | **Get** /repos/{repo_ref}/diff/{range} | 
[**RuleAdd**](RepositoryApi.md#RuleAdd) | **Post** /repos/{repo_ref}/rules | 
[**RuleDelete**](RepositoryApi.md#RuleDelete) | **Delete** /repos/{repo_ref}/rules/{rule_uid} | 
[**RuleGet**](RepositoryApi.md#RuleGet) | **Get** /repos/{repo_ref}/rules/{rule_uid} | 
[**RuleList**](RepositoryApi.md#RuleList) | **Get** /repos/{repo_ref}/rules | 
[**RuleUpdate**](RepositoryApi.md#RuleUpdate) | **Patch** /repos/{repo_ref}/rules/{rule_uid} | 
[**UpdateRepository**](RepositoryApi.md#UpdateRepository) | **Patch** /repos/{repo_ref} | 

# **CalculateCommitDivergence**
> []RepoCommitDivergence CalculateCommitDivergence(ctx, repoRef, optional)


### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **repoRef** | **string**|  | 
 **optional** | ***RepositoryApiCalculateCommitDivergenceOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a RepositoryApiCalculateCommitDivergenceOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **body** | [**optional.Interface of OpenapiCalculateCommitDivergenceRequest**](OpenapiCalculateCommitDivergenceRequest.md)|  | 

### Return type

[**[]RepoCommitDivergence**](RepoCommitDivergence.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **CodeOwnersValidate**
> CodeOwnersValidate(ctx, repoRef, optional)


### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **repoRef** | **string**|  | 
 **optional** | ***RepositoryApiCodeOwnersValidateOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a RepositoryApiCodeOwnersValidateOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **gitRef** | **optional.String**| The git reference (branch / tag / commitID) that will be used to retrieve the data. If no value is provided the default branch of the repository is used. | [default to {Repository Default Branch}]

### Return type

 (empty response body)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **CommitFiles**
> TypesCommitFilesResponse CommitFiles(ctx, repoRef, optional)


### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **repoRef** | **string**|  | 
 **optional** | ***RepositoryApiCommitFilesOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a RepositoryApiCommitFilesOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **body** | [**optional.Interface of OpenapiCommitFilesRequest**](OpenapiCommitFilesRequest.md)|  | 

### Return type

[**TypesCommitFilesResponse**](TypesCommitFilesResponse.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **CreateBranch**
> RepoBranch CreateBranch(ctx, repoRef, optional)


### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **repoRef** | **string**|  | 
 **optional** | ***RepositoryApiCreateBranchOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a RepositoryApiCreateBranchOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **body** | [**optional.Interface of OpenapiCreateBranchRequest**](OpenapiCreateBranchRequest.md)|  | 

### Return type

[**RepoBranch**](RepoBranch.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **CreateRepository**
> TypesRepository CreateRepository(ctx, optional)


### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***RepositoryApiCreateRepositoryOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a RepositoryApiCreateRepositoryOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **body** | [**optional.Interface of OpenapiCreateRepositoryRequest**](OpenapiCreateRepositoryRequest.md)|  | 
 **spacePath** | **optional.**| path of parent space (Not needed in standalone). | [default to false]

### Return type

[**TypesRepository**](TypesRepository.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **CreateTag**
> RepoCommitTag CreateTag(ctx, repoRef, optional)


### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **repoRef** | **string**|  | 
 **optional** | ***RepositoryApiCreateTagOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a RepositoryApiCreateTagOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **body** | [**optional.Interface of OpenapiCreateTagRequest**](OpenapiCreateTagRequest.md)|  | 

### Return type

[**RepoCommitTag**](RepoCommitTag.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **DeleteBranch**
> DeleteBranch(ctx, repoRef, branchName, optional)


### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **repoRef** | **string**|  | 
  **branchName** | **string**|  | 
 **optional** | ***RepositoryApiDeleteBranchOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a RepositoryApiDeleteBranchOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **bypassRules** | **optional.Bool**| Bypass rule violations if possible. | [default to false]

### Return type

 (empty response body)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **DeleteRepository**
> DeleteRepository(ctx, repoRef)


### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **repoRef** | **string**|  | 

### Return type

 (empty response body)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **DeleteTag**
> DeleteTag(ctx, repoRef, tagName, optional)


### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **repoRef** | **string**|  | 
  **tagName** | **string**|  | 
 **optional** | ***RepositoryApiDeleteTagOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a RepositoryApiDeleteTagOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **bypassRules** | **optional.Bool**| Bypass rule violations if possible. | [default to false]

### Return type

 (empty response body)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **DiffStats**
> TypesDiffStats DiffStats(ctx, repoRef, range_)


### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **repoRef** | **string**|  | 
  **range_** | **string**|  | 

### Return type

[**TypesDiffStats**](TypesDiffStats.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **FindRepository**
> TypesRepository FindRepository(ctx, repoRef)


### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **repoRef** | **string**|  | 

### Return type

[**TypesRepository**](TypesRepository.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetBlame**
> []GitBlamePart GetBlame(ctx, repoRef, path, optional)


### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **repoRef** | **string**|  | 
  **path** | **string**|  | 
 **optional** | ***RepositoryApiGetBlameOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a RepositoryApiGetBlameOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **gitRef** | **optional.String**| The git reference (branch / tag / commitID) that will be used to retrieve the data. If no value is provided the default branch of the repository is used. | [default to {Repository Default Branch}]
 **lineFrom** | **optional.Int32**| Line number from which the file data is considered | [default to 0]
 **lineTo** | **optional.Int32**| Line number to which the file data is considered | [default to 0]

### Return type

[**[]GitBlamePart**](GitBlamePart.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetBranch**
> RepoBranch GetBranch(ctx, repoRef, branchName)


### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **repoRef** | **string**|  | 
  **branchName** | **string**|  | 

### Return type

[**RepoBranch**](RepoBranch.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetCommit**
> TypesCommit GetCommit(ctx, repoRef, commitSha)


### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **repoRef** | **string**|  | 
  **commitSha** | **string**|  | 

### Return type

[**TypesCommit**](TypesCommit.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetCommitDiff**
> string GetCommitDiff(ctx, repoRef, commitSha)


### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **repoRef** | **string**|  | 
  **commitSha** | **string**|  | 

### Return type

**string**

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: text/plain, application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetContent**
> OpenapiGetContentOutput GetContent(ctx, repoRef, path, optional)


### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **repoRef** | **string**|  | 
  **path** | **string**|  | 
 **optional** | ***RepositoryApiGetContentOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a RepositoryApiGetContentOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **gitRef** | **optional.String**| The git reference (branch / tag / commitID) that will be used to retrieve the data. If no value is provided the default branch of the repository is used. | [default to {Repository Default Branch}]
 **includeCommit** | **optional.Bool**| Indicates whether optional commit information should be included in the response. | [default to false]

### Return type

[**OpenapiGetContentOutput**](OpenapiGetContentOutput.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetRaw**
> GetRaw(ctx, repoRef, path, optional)


### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **repoRef** | **string**|  | 
  **path** | **string**|  | 
 **optional** | ***RepositoryApiGetRawOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a RepositoryApiGetRawOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **gitRef** | **optional.String**| The git reference (branch / tag / commitID) that will be used to retrieve the data. If no value is provided the default branch of the repository is used. | [default to {Repository Default Branch}]

### Return type

 (empty response body)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ImportRepository**
> TypesRepository ImportRepository(ctx, optional)


### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***RepositoryApiImportRepositoryOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a RepositoryApiImportRepositoryOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **body** | [**optional.Interface of ReposImportBody**](ReposImportBody.md)|  | 
 **spacePath** | **optional.**| path of parent space (Not needed in standalone). | [default to false]

### Return type

[**TypesRepository**](TypesRepository.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ListBranches**
> []RepoBranch ListBranches(ctx, repoRef, optional)


### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **repoRef** | **string**|  | 
 **optional** | ***RepositoryApiListBranchesOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a RepositoryApiListBranchesOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **includeCommit** | **optional.Bool**| Indicates whether optional commit information should be included in the response. | [default to false]
 **query** | **optional.String**| The substring by which the branches are filtered. | 
 **order** | **optional.String**| The order of the output. | [default to asc]
 **sort** | **optional.String**| The data by which the branches are sorted. | [default to name]
 **page** | **optional.Int32**| The page to return. | [default to 1]
 **limit** | **optional.Int32**| The maximum number of results to return. | [default to 30]

### Return type

[**[]RepoBranch**](RepoBranch.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ListCommits**
> []TypesListCommitResponse ListCommits(ctx, repoRef, optional)


### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **repoRef** | **string**|  | 
 **optional** | ***RepositoryApiListCommitsOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a RepositoryApiListCommitsOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **gitRef** | **optional.String**| The git reference (branch / tag / commitID) that will be used to retrieve the data. If no value is provided the default branch of the repository is used. | [default to {Repository Default Branch}]
 **after** | **optional.String**| The result should only contain commits that occurred after the provided reference. | 
 **path** | **optional.String**| Path for which commit information should be retrieved | 
 **since** | **optional.Int32**| Epoch since when commit information should be retrieved. | 
 **until** | **optional.Int32**| Epoch until when commit information should be retrieved. | 
 **committer** | **optional.String**| Committer pattern for which commit information should be retrieved. | 
 **page** | **optional.Int32**| The page to return. | [default to 1]
 **limit** | **optional.Int32**| The maximum number of results to return. | [default to 30]

### Return type

[**[]TypesListCommitResponse**](TypesListCommitResponse.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ListRepositoryServiceAccounts**
> []TypesServiceAccount ListRepositoryServiceAccounts(ctx, repoRef)


### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **repoRef** | **string**|  | 

### Return type

[**[]TypesServiceAccount**](TypesServiceAccount.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ListTags**
> []RepoCommitTag ListTags(ctx, repoRef, optional)


### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **repoRef** | **string**|  | 
 **optional** | ***RepositoryApiListTagsOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a RepositoryApiListTagsOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **includeCommit** | **optional.Bool**| Indicates whether optional commit information should be included in the response. | [default to false]
 **query** | **optional.String**| The substring by which the tags are filtered. | 
 **order** | **optional.String**| The order of the output. | [default to asc]
 **sort** | **optional.String**| The data by which the tags are sorted. | [default to name]
 **page** | **optional.Int32**| The page to return. | [default to 1]
 **limit** | **optional.Int32**| The maximum number of results to return. | [default to 30]

### Return type

[**[]RepoCommitTag**](RepoCommitTag.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **MergeCheck**
> RepoMergeCheck MergeCheck(ctx, repoRef, range_)


### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **repoRef** | **string**|  | 
  **range_** | **string**|  | 

### Return type

[**RepoMergeCheck**](RepoMergeCheck.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **MoveRepository**
> TypesRepository MoveRepository(ctx, repoRef, optional)


### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **repoRef** | **string**|  | 
 **optional** | ***RepositoryApiMoveRepositoryOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a RepositoryApiMoveRepositoryOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **body** | [**optional.Interface of OpenapiMoveRepoRequest**](OpenapiMoveRepoRequest.md)|  | 

### Return type

[**TypesRepository**](TypesRepository.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **PathDetails**
> RepoPathsDetailsOutput PathDetails(ctx, repoRef, optional)


### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **repoRef** | **string**|  | 
 **optional** | ***RepositoryApiPathDetailsOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a RepositoryApiPathDetailsOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **body** | [**optional.Interface of OpenapiPathsDetailsRequest**](OpenapiPathsDetailsRequest.md)|  | 
 **gitRef** | **optional.**| The git reference (branch / tag / commitID) that will be used to retrieve the data. If no value is provided the default branch of the repository is used. | [default to {Repository Default Branch}]

### Return type

[**RepoPathsDetailsOutput**](RepoPathsDetailsOutput.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **RawDiff**
> []GitFileDiff RawDiff(ctx, repoRef, range_)


### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **repoRef** | **string**|  | 
  **range_** | **string**|  | 

### Return type

[**[]GitFileDiff**](GitFileDiff.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json, text/plain

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **RuleAdd**
> OpenapiRule RuleAdd(ctx, repoRef, optional)


### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **repoRef** | **string**|  | 
 **optional** | ***RepositoryApiRuleAddOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a RepositoryApiRuleAddOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **body** | [**optional.Interface of RepoRefRulesBody**](RepoRefRulesBody.md)|  | 

### Return type

[**OpenapiRule**](OpenapiRule.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **RuleDelete**
> RuleDelete(ctx, repoRef, ruleUid)


### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **repoRef** | **string**|  | 
  **ruleUid** | **string**|  | 

### Return type

 (empty response body)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **RuleGet**
> []OpenapiRule RuleGet(ctx, repoRef, ruleUid)


### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **repoRef** | **string**|  | 
  **ruleUid** | **string**|  | 

### Return type

[**[]OpenapiRule**](OpenapiRule.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **RuleList**
> []OpenapiRule RuleList(ctx, repoRef, optional)


### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **repoRef** | **string**|  | 
 **optional** | ***RepositoryApiRuleListOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a RepositoryApiRuleListOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **query** | **optional.String**| The substring by which the repository protection rules are filtered. | 
 **order** | **optional.String**| The order of the output. | [default to asc]
 **sort** | **optional.String**| The field by which the protection rules are sorted. | [default to created_at]
 **page** | **optional.Int32**| The page to return. | [default to 1]
 **limit** | **optional.Int32**| The maximum number of results to return. | [default to 30]

### Return type

[**[]OpenapiRule**](OpenapiRule.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **RuleUpdate**
> OpenapiRule RuleUpdate(ctx, repoRef, ruleUid, optional)


### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **repoRef** | **string**|  | 
  **ruleUid** | **string**|  | 
 **optional** | ***RepositoryApiRuleUpdateOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a RepositoryApiRuleUpdateOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **body** | [**optional.Interface of RulesRuleUidBody**](RulesRuleUidBody.md)|  | 

### Return type

[**OpenapiRule**](OpenapiRule.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **UpdateRepository**
> TypesRepository UpdateRepository(ctx, repoRef, optional)


### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **repoRef** | **string**|  | 
 **optional** | ***RepositoryApiUpdateRepositoryOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a RepositoryApiUpdateRepositoryOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **body** | [**optional.Interface of OpenapiUpdateRepoRequest**](OpenapiUpdateRepoRequest.md)|  | 

### Return type

[**TypesRepository**](TypesRepository.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

