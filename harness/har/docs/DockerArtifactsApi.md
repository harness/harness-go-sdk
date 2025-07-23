# \DockerArtifactsAPI

All URIs are relative to */gateway/har/api/v1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetDockerArtifactDetails**](DockerArtifactsAPI.md#GetDockerArtifactDetails) | **Get** /registry/{registry_ref}/+/artifact/{artifact}/+/version/{version}/docker/details | Describe Docker Artifact Detail
[**GetDockerArtifactIntegrationDetails**](DockerArtifactsAPI.md#GetDockerArtifactIntegrationDetails) | **Get** /registry/{registry_ref}/+/artifact/{artifact}/+/version/{version}/docker/integrationdetails | Describe Docker Artifact Integration Detail
[**GetDockerArtifactLayers**](DockerArtifactsAPI.md#GetDockerArtifactLayers) | **Get** /registry/{registry_ref}/+/artifact/{artifact}/+/version/{version}/docker/layers | Describe Docker Artifact Layers
[**GetDockerArtifactManifest**](DockerArtifactsAPI.md#GetDockerArtifactManifest) | **Get** /registry/{registry_ref}/+/artifact/{artifact}/+/version/{version}/docker/manifest | Describe Docker Artifact Manifest
[**GetDockerArtifactManifests**](DockerArtifactsAPI.md#GetDockerArtifactManifests) | **Get** /registry/{registry_ref}/+/artifact/{artifact}/+/version/{version}/docker/manifests | Describe Docker Artifact Manifests



## GetDockerArtifactDetails

> InlineResponse2005 GetDockerArtifactDetails(ctx, registryRef, artifact, version).Digest(digest).Execute()

Describe Docker Artifact Detail



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
	digest := "digest_example" // string | Digest.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DockerArtifactsAPI.GetDockerArtifactDetails(context.Background(), registryRef, artifact, version).Digest(digest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DockerArtifactsAPI.GetDockerArtifactDetails``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetDockerArtifactDetails`: InlineResponse2005
	fmt.Fprintf(os.Stdout, "Response from `DockerArtifactsAPI.GetDockerArtifactDetails`: %v\n", resp)
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

Other parameters are passed through a pointer to a apiGetDockerArtifactDetailsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **digest** | **string** | Digest. | 

### Return type

[**InlineResponse2005**](InlineResponse2005.md)

### Authorization

[XApiKeyAuth](../README.md#XApiKeyAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetDockerArtifactIntegrationDetails

> InlineResponse2006 GetDockerArtifactIntegrationDetails(ctx, registryRef, artifact, version).Digest(digest).Execute()

Describe Docker Artifact Integration Detail



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
	digest := "digest_example" // string | Digest.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DockerArtifactsAPI.GetDockerArtifactIntegrationDetails(context.Background(), registryRef, artifact, version).Digest(digest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DockerArtifactsAPI.GetDockerArtifactIntegrationDetails``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetDockerArtifactIntegrationDetails`: InlineResponse2006
	fmt.Fprintf(os.Stdout, "Response from `DockerArtifactsAPI.GetDockerArtifactIntegrationDetails`: %v\n", resp)
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

Other parameters are passed through a pointer to a apiGetDockerArtifactIntegrationDetailsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **digest** | **string** | Digest. | 

### Return type

[**InlineResponse2006**](InlineResponse2006.md)

### Authorization

[XApiKeyAuth](../README.md#XApiKeyAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetDockerArtifactLayers

> InlineResponse2007 GetDockerArtifactLayers(ctx, registryRef, artifact, version).Digest(digest).Execute()

Describe Docker Artifact Layers



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
	digest := "digest_example" // string | Digest.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DockerArtifactsAPI.GetDockerArtifactLayers(context.Background(), registryRef, artifact, version).Digest(digest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DockerArtifactsAPI.GetDockerArtifactLayers``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetDockerArtifactLayers`: InlineResponse2007
	fmt.Fprintf(os.Stdout, "Response from `DockerArtifactsAPI.GetDockerArtifactLayers`: %v\n", resp)
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

Other parameters are passed through a pointer to a apiGetDockerArtifactLayersRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **digest** | **string** | Digest. | 

### Return type

[**InlineResponse2007**](InlineResponse2007.md)

### Authorization

[XApiKeyAuth](../README.md#XApiKeyAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetDockerArtifactManifest

> InlineResponse2008 GetDockerArtifactManifest(ctx, registryRef, artifact, version).Digest(digest).Execute()

Describe Docker Artifact Manifest



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
	digest := "digest_example" // string | Digest.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DockerArtifactsAPI.GetDockerArtifactManifest(context.Background(), registryRef, artifact, version).Digest(digest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DockerArtifactsAPI.GetDockerArtifactManifest``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetDockerArtifactManifest`: InlineResponse2008
	fmt.Fprintf(os.Stdout, "Response from `DockerArtifactsAPI.GetDockerArtifactManifest`: %v\n", resp)
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

Other parameters are passed through a pointer to a apiGetDockerArtifactManifestRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **digest** | **string** | Digest. | 

### Return type

[**InlineResponse2008**](InlineResponse2008.md)

### Authorization

[XApiKeyAuth](../README.md#XApiKeyAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetDockerArtifactManifests

> InlineResponse2009 GetDockerArtifactManifests(ctx, registryRef, artifact, version).Execute()

Describe Docker Artifact Manifests



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
	resp, r, err := apiClient.DockerArtifactsAPI.GetDockerArtifactManifests(context.Background(), registryRef, artifact, version).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DockerArtifactsAPI.GetDockerArtifactManifests``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetDockerArtifactManifests`: InlineResponse2009
	fmt.Fprintf(os.Stdout, "Response from `DockerArtifactsAPI.GetDockerArtifactManifests`: %v\n", resp)
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

Other parameters are passed through a pointer to a apiGetDockerArtifactManifestsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------




### Return type

[**InlineResponse2009**](InlineResponse2009.md)

### Authorization

[XApiKeyAuth](../README.md#XApiKeyAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

