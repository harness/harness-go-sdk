# \HelmArtifactsAPI

All URIs are relative to */gateway/har/api/v1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetHelmArtifactDetails**](HelmArtifactsAPI.md#GetHelmArtifactDetails) | **Get** /registry/{registry_ref}/+/artifact/{artifact}/+/version/{version}/helm/details | Describe Helm Artifact Detail
[**GetHelmArtifactManifest**](HelmArtifactsAPI.md#GetHelmArtifactManifest) | **Get** /registry/{registry_ref}/+/artifact/{artifact}/+/version/{version}/helm/manifest | Describe Helm Artifact Manifest



## GetHelmArtifactDetails

> InlineResponse20012 GetHelmArtifactDetails(ctx, registryRef, artifact, version).Execute()

Describe Helm Artifact Detail



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID"
)

func main() {
	registryRef := "registryRef_example" // string | Unique registry path.
	artifact := "artifact_example" // string | Name of artifact.
	version := "version_example" // string | Name of Artifact Version.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.HelmArtifactsAPI.GetHelmArtifactDetails(context.Background(), registryRef, artifact, version).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `HelmArtifactsAPI.GetHelmArtifactDetails``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetHelmArtifactDetails`: InlineResponse20012
	fmt.Fprintf(os.Stdout, "Response from `HelmArtifactsAPI.GetHelmArtifactDetails`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**registryRef** | **string** | Unique registry path. | 
**artifact** | **string** | Name of artifact. | 
**version** | **string** | Name of Artifact Version. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetHelmArtifactDetailsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------




### Return type

[**InlineResponse20012**](InlineResponse20012.md)

### Authorization

[XApiKeyAuth](../README.md#XApiKeyAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetHelmArtifactManifest

> InlineResponse20013 GetHelmArtifactManifest(ctx, registryRef, artifact, version).Execute()

Describe Helm Artifact Manifest



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID"
)

func main() {
	registryRef := "registryRef_example" // string | Unique registry path.
	artifact := "artifact_example" // string | Name of artifact.
	version := "version_example" // string | Name of Artifact Version.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.HelmArtifactsAPI.GetHelmArtifactManifest(context.Background(), registryRef, artifact, version).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `HelmArtifactsAPI.GetHelmArtifactManifest``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetHelmArtifactManifest`: InlineResponse20013
	fmt.Fprintf(os.Stdout, "Response from `HelmArtifactsAPI.GetHelmArtifactManifest`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**registryRef** | **string** | Unique registry path. | 
**artifact** | **string** | Name of artifact. | 
**version** | **string** | Name of Artifact Version. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetHelmArtifactManifestRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------




### Return type

[**InlineResponse20013**](InlineResponse20013.md)

### Authorization

[XApiKeyAuth](../README.md#XApiKeyAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

