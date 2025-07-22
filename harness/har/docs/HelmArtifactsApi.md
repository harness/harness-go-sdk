# {{classname}}

All URIs are relative to *https://app.harness.io/api/v1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetHelmArtifactDetails**](HelmArtifactsApi.md#GetHelmArtifactDetails) | **Get** /registry/{registry_ref}/artifact/{artifact}/version/{version}/helm/details | Describe Helm Artifact Detail
[**GetHelmArtifactManifest**](HelmArtifactsApi.md#GetHelmArtifactManifest) | **Get** /registry/{registry_ref}/artifact/{artifact}/version/{version}/helm/manifest | Describe Helm Artifact Manifest

# **GetHelmArtifactDetails**
> InlineResponse20012 GetHelmArtifactDetails(ctx, registryRef, artifact, version)
Describe Helm Artifact Detail

Get Helm Artifact Details

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **registryRef** | **string**| Unique registry path. | 
  **artifact** | **string**| Name of artifact. | 
  **version** | **string**| Name of Artifact Version. | 

### Return type

[**InlineResponse20012**](inline_response_200_12.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetHelmArtifactManifest**
> InlineResponse20013 GetHelmArtifactManifest(ctx, registryRef, artifact, version)
Describe Helm Artifact Manifest

Get Helm Artifact Manifest

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **registryRef** | **string**| Unique registry path. | 
  **artifact** | **string**| Name of artifact. | 
  **version** | **string**| Name of Artifact Version. | 

### Return type

[**InlineResponse20013**](inline_response_200_13.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

