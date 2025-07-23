# \ArtifactsAPI

All URIs are relative to */gateway/har/api/v1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**DeleteArtifact**](ArtifactsAPI.md#DeleteArtifact) | **Delete** /registry/{registry_ref}/+/artifact/{artifact}/+ | Delete Artifact
[**DeleteArtifactVersion**](ArtifactsAPI.md#DeleteArtifactVersion) | **Delete** /registry/{registry_ref}/+/artifact/{artifact}/+/version/{version} | Delete an Artifact Version
[**GetAllArtifactVersions**](ArtifactsAPI.md#GetAllArtifactVersions) | **Get** /registry/{registry_ref}/+/artifact/{artifact}/+/versions | List Artifact Versions
[**GetArtifactDeployments**](ArtifactsAPI.md#GetArtifactDeployments) | **Get** /registry/{registry_ref}/+/artifact/{artifact}/+/version/{version}/deploymentdetails | Describe Artifact Deployments
[**GetArtifactDetails**](ArtifactsAPI.md#GetArtifactDetails) | **Get** /registry/{registry_ref}/+/artifact/{artifact}/+/version/{version}/details | Describe Artifact Details
[**GetArtifactFile**](ArtifactsAPI.md#GetArtifactFile) | **Get** /registry/{registry_ref}/+/artifact/{artifact}/+/version/{version}/file/{file_name} | Get Artifact file
[**GetArtifactFiles**](ArtifactsAPI.md#GetArtifactFiles) | **Get** /registry/{registry_ref}/+/artifact/{artifact}/+/version/{version}/files | Describe Artifact files
[**GetArtifactStats**](ArtifactsAPI.md#GetArtifactStats) | **Get** /registry/{registry_ref}/+/artifact/{artifact}/+/stats | Get Artifact Stats
[**GetArtifactStatsForRegistry**](ArtifactsAPI.md#GetArtifactStatsForRegistry) | **Get** /registry/{registry_ref}/+/artifact/stats | Get Artifact Stats
[**GetArtifactSummary**](ArtifactsAPI.md#GetArtifactSummary) | **Get** /registry/{registry_ref}/+/artifact/{artifact}/+/summary | Get Artifact Summary
[**GetArtifactVersionSummary**](ArtifactsAPI.md#GetArtifactVersionSummary) | **Get** /registry/{registry_ref}/+/artifact/{artifact}/+/version/{version}/summary | Get Artifact Version Summary
[**ListArtifactLabels**](ArtifactsAPI.md#ListArtifactLabels) | **Get** /registry/{registry_ref}/+/artifact/labels | List Artifact Labels
[**RedirectHarnessArtifact**](ArtifactsAPI.md#RedirectHarnessArtifact) | **Get** /registry/{registry_identifier}/artifact/{artifact}/+/redirect | Redirect to Harness Artifact Page
[**UpdateArtifactLabels**](ArtifactsAPI.md#UpdateArtifactLabels) | **Put** /registry/{registry_ref}/+/artifact/{artifact}/+/labels | Update Artifact Labels



## DeleteArtifact

> InlineResponse200 DeleteArtifact(ctx, registryRef, artifact).Execute()

Delete Artifact



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

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ArtifactsAPI.DeleteArtifact(context.Background(), registryRef, artifact).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ArtifactsAPI.DeleteArtifact``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `DeleteArtifact`: InlineResponse200
	fmt.Fprintf(os.Stdout, "Response from `ArtifactsAPI.DeleteArtifact`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**registryRef** | **string** | Unique registry path. | 
**artifact** | **string** | Name of artifact. | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteArtifactRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**InlineResponse200**](InlineResponse200.md)

### Authorization

[XApiKeyAuth](../README.md#XApiKeyAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteArtifactVersion

> InlineResponse200 DeleteArtifactVersion(ctx, registryRef, artifact, version).Execute()

Delete an Artifact Version



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
	resp, r, err := apiClient.ArtifactsAPI.DeleteArtifactVersion(context.Background(), registryRef, artifact, version).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ArtifactsAPI.DeleteArtifactVersion``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `DeleteArtifactVersion`: InlineResponse200
	fmt.Fprintf(os.Stdout, "Response from `ArtifactsAPI.DeleteArtifactVersion`: %v\n", resp)
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

Other parameters are passed through a pointer to a apiDeleteArtifactVersionRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------




### Return type

[**InlineResponse200**](InlineResponse200.md)

### Authorization

[XApiKeyAuth](../README.md#XApiKeyAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetAllArtifactVersions

> InlineResponse20015 GetAllArtifactVersions(ctx, registryRef, artifact).Page(page).Size(size).SortOrder(sortOrder).SortField(sortField).SearchTerm(searchTerm).Execute()

List Artifact Versions



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
	page := int64(789) // int64 | Current page number (optional) (default to 1)
	size := int64(789) // int64 | Number of items per page (optional) (default to 20)
	sortOrder := "sortOrder_example" // string | sortOrder (optional)
	sortField := "sortField_example" // string | sortField (optional)
	searchTerm := "searchTerm_example" // string | search Term. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ArtifactsAPI.GetAllArtifactVersions(context.Background(), registryRef, artifact).Page(page).Size(size).SortOrder(sortOrder).SortField(sortField).SearchTerm(searchTerm).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ArtifactsAPI.GetAllArtifactVersions``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetAllArtifactVersions`: InlineResponse20015
	fmt.Fprintf(os.Stdout, "Response from `ArtifactsAPI.GetAllArtifactVersions`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**registryRef** | **string** | Unique registry path. | 
**artifact** | **string** | Name of artifact. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetAllArtifactVersionsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **page** | **int64** | Current page number | [default to 1]
 **size** | **int64** | Number of items per page | [default to 20]
 **sortOrder** | **string** | sortOrder | 
 **sortField** | **string** | sortField | 
 **searchTerm** | **string** | search Term. | 

### Return type

[**InlineResponse20015**](InlineResponse20015.md)

### Authorization

[XApiKeyAuth](../README.md#XApiKeyAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetArtifactDeployments

> InlineResponse2003 GetArtifactDeployments(ctx, registryRef, artifact, version).EnvType(envType).Page(page).Size(size).SortOrder(sortOrder).SortField(sortField).SearchTerm(searchTerm).Execute()

Describe Artifact Deployments



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
	envType := "envType_example" // string | env type (optional)
	page := int64(789) // int64 | Current page number (optional) (default to 1)
	size := int64(789) // int64 | Number of items per page (optional) (default to 20)
	sortOrder := "sortOrder_example" // string | sortOrder (optional)
	sortField := "sortField_example" // string | sortField (optional)
	searchTerm := "searchTerm_example" // string | search Term. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ArtifactsAPI.GetArtifactDeployments(context.Background(), registryRef, artifact, version).EnvType(envType).Page(page).Size(size).SortOrder(sortOrder).SortField(sortField).SearchTerm(searchTerm).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ArtifactsAPI.GetArtifactDeployments``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetArtifactDeployments`: InlineResponse2003
	fmt.Fprintf(os.Stdout, "Response from `ArtifactsAPI.GetArtifactDeployments`: %v\n", resp)
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

Other parameters are passed through a pointer to a apiGetArtifactDeploymentsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **envType** | **string** | env type | 
 **page** | **int64** | Current page number | [default to 1]
 **size** | **int64** | Number of items per page | [default to 20]
 **sortOrder** | **string** | sortOrder | 
 **sortField** | **string** | sortField | 
 **searchTerm** | **string** | search Term. | 

### Return type

[**InlineResponse2003**](InlineResponse2003.md)

### Authorization

[XApiKeyAuth](../README.md#XApiKeyAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetArtifactDetails

> InlineResponse2004 GetArtifactDetails(ctx, registryRef, artifact, version).ChildVersion(childVersion).Execute()

Describe Artifact Details



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
	childVersion := "childVersion_example" // string | Child version incase of Docker artifacts. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ArtifactsAPI.GetArtifactDetails(context.Background(), registryRef, artifact, version).ChildVersion(childVersion).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ArtifactsAPI.GetArtifactDetails``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetArtifactDetails`: InlineResponse2004
	fmt.Fprintf(os.Stdout, "Response from `ArtifactsAPI.GetArtifactDetails`: %v\n", resp)
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

Other parameters are passed through a pointer to a apiGetArtifactDetailsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **childVersion** | **string** | Child version incase of Docker artifacts. | 

### Return type

[**InlineResponse2004**](InlineResponse2004.md)

### Authorization

[XApiKeyAuth](../README.md#XApiKeyAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetArtifactFile

> InlineResponse20010 GetArtifactFile(ctx, registryRef, artifact, version, fileName).Execute()

Get Artifact file



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
	fileName := "fileName_example" // string | Name of Artifact File.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ArtifactsAPI.GetArtifactFile(context.Background(), registryRef, artifact, version, fileName).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ArtifactsAPI.GetArtifactFile``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetArtifactFile`: InlineResponse20010
	fmt.Fprintf(os.Stdout, "Response from `ArtifactsAPI.GetArtifactFile`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**registryRef** | **string** | Unique registry path. | 
**artifact** | **string** | Name of artifact. | 
**version** | **string** | Name of Artifact Version. | 
**fileName** | **string** | Name of Artifact File. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetArtifactFileRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------





### Return type

[**InlineResponse20010**](InlineResponse20010.md)

### Authorization

[XApiKeyAuth](../README.md#XApiKeyAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetArtifactFiles

> InlineResponse20011 GetArtifactFiles(ctx, registryRef, artifact, version).Page(page).Size(size).SortOrder(sortOrder).SortField(sortField).SearchTerm(searchTerm).Execute()

Describe Artifact files



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
	page := int64(789) // int64 | Current page number (optional) (default to 1)
	size := int64(789) // int64 | Number of items per page (optional) (default to 20)
	sortOrder := "sortOrder_example" // string | sortOrder (optional)
	sortField := "sortField_example" // string | sortField (optional)
	searchTerm := "searchTerm_example" // string | search Term. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ArtifactsAPI.GetArtifactFiles(context.Background(), registryRef, artifact, version).Page(page).Size(size).SortOrder(sortOrder).SortField(sortField).SearchTerm(searchTerm).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ArtifactsAPI.GetArtifactFiles``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetArtifactFiles`: InlineResponse20011
	fmt.Fprintf(os.Stdout, "Response from `ArtifactsAPI.GetArtifactFiles`: %v\n", resp)
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

Other parameters are passed through a pointer to a apiGetArtifactFilesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **page** | **int64** | Current page number | [default to 1]
 **size** | **int64** | Number of items per page | [default to 20]
 **sortOrder** | **string** | sortOrder | 
 **sortField** | **string** | sortField | 
 **searchTerm** | **string** | search Term. | 

### Return type

[**InlineResponse20011**](InlineResponse20011.md)

### Authorization

[XApiKeyAuth](../README.md#XApiKeyAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetArtifactStats

> InlineResponse2002 GetArtifactStats(ctx, registryRef, artifact).From(from).To(to).Execute()

Get Artifact Stats



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
	from := "from_example" // string | Date. Format - MM/DD/YYYY (optional)
	to := "to_example" // string | Date. Format - MM/DD/YYYY (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ArtifactsAPI.GetArtifactStats(context.Background(), registryRef, artifact).From(from).To(to).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ArtifactsAPI.GetArtifactStats``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetArtifactStats`: InlineResponse2002
	fmt.Fprintf(os.Stdout, "Response from `ArtifactsAPI.GetArtifactStats`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**registryRef** | **string** | Unique registry path. | 
**artifact** | **string** | Name of artifact. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetArtifactStatsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **from** | **string** | Date. Format - MM/DD/YYYY | 
 **to** | **string** | Date. Format - MM/DD/YYYY | 

### Return type

[**InlineResponse2002**](InlineResponse2002.md)

### Authorization

[XApiKeyAuth](../README.md#XApiKeyAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetArtifactStatsForRegistry

> InlineResponse2002 GetArtifactStatsForRegistry(ctx, registryRef).From(from).To(to).Execute()

Get Artifact Stats



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
	from := "from_example" // string | Date. Format - MM/DD/YYYY (optional)
	to := "to_example" // string | Date. Format - MM/DD/YYYY (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ArtifactsAPI.GetArtifactStatsForRegistry(context.Background(), registryRef).From(from).To(to).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ArtifactsAPI.GetArtifactStatsForRegistry``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetArtifactStatsForRegistry`: InlineResponse2002
	fmt.Fprintf(os.Stdout, "Response from `ArtifactsAPI.GetArtifactStatsForRegistry`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**registryRef** | **string** | Unique registry path. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetArtifactStatsForRegistryRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **from** | **string** | Date. Format - MM/DD/YYYY | 
 **to** | **string** | Date. Format - MM/DD/YYYY | 

### Return type

[**InlineResponse2002**](InlineResponse2002.md)

### Authorization

[XApiKeyAuth](../README.md#XApiKeyAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetArtifactSummary

> InlineResponse2001 GetArtifactSummary(ctx, registryRef, artifact).Execute()

Get Artifact Summary



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

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ArtifactsAPI.GetArtifactSummary(context.Background(), registryRef, artifact).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ArtifactsAPI.GetArtifactSummary``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetArtifactSummary`: InlineResponse2001
	fmt.Fprintf(os.Stdout, "Response from `ArtifactsAPI.GetArtifactSummary`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**registryRef** | **string** | Unique registry path. | 
**artifact** | **string** | Name of artifact. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetArtifactSummaryRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**InlineResponse2001**](InlineResponse2001.md)

### Authorization

[XApiKeyAuth](../README.md#XApiKeyAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetArtifactVersionSummary

> InlineResponse20014 GetArtifactVersionSummary(ctx, registryRef, artifact, version).Digest(digest).Execute()

Get Artifact Version Summary



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
	digest := "digest_example" // string | Digest. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ArtifactsAPI.GetArtifactVersionSummary(context.Background(), registryRef, artifact, version).Digest(digest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ArtifactsAPI.GetArtifactVersionSummary``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetArtifactVersionSummary`: InlineResponse20014
	fmt.Fprintf(os.Stdout, "Response from `ArtifactsAPI.GetArtifactVersionSummary`: %v\n", resp)
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

Other parameters are passed through a pointer to a apiGetArtifactVersionSummaryRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **digest** | **string** | Digest. | 

### Return type

[**InlineResponse20014**](InlineResponse20014.md)

### Authorization

[XApiKeyAuth](../README.md#XApiKeyAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListArtifactLabels

> InlineResponse20016 ListArtifactLabels(ctx, registryRef).Page(page).Size(size).SearchTerm(searchTerm).Execute()

List Artifact Labels



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
	page := int64(789) // int64 | Current page number (optional) (default to 1)
	size := int64(789) // int64 | Number of items per page (optional) (default to 20)
	searchTerm := "searchTerm_example" // string | search Term. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ArtifactsAPI.ListArtifactLabels(context.Background(), registryRef).Page(page).Size(size).SearchTerm(searchTerm).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ArtifactsAPI.ListArtifactLabels``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ListArtifactLabels`: InlineResponse20016
	fmt.Fprintf(os.Stdout, "Response from `ArtifactsAPI.ListArtifactLabels`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**registryRef** | **string** | Unique registry path. | 

### Other Parameters

Other parameters are passed through a pointer to a apiListArtifactLabelsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **page** | **int64** | Current page number | [default to 1]
 **size** | **int64** | Number of items per page | [default to 20]
 **searchTerm** | **string** | search Term. | 

### Return type

[**InlineResponse20016**](InlineResponse20016.md)

### Authorization

[XApiKeyAuth](../README.md#XApiKeyAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## RedirectHarnessArtifact

> RedirectHarnessArtifact(ctx, registryIdentifier, artifact).AccountIdentifier(accountIdentifier).Version(version).Execute()

Redirect to Harness Artifact Page



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
	registryIdentifier := "registryIdentifier_example" // string | Unique registry Identifier in a account.
	artifact := "artifact_example" // string | Name of artifact.
	accountIdentifier := "accountIdentifier_example" // string | Account Identifier (optional)
	version := "version_example" // string | Version (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.ArtifactsAPI.RedirectHarnessArtifact(context.Background(), registryIdentifier, artifact).AccountIdentifier(accountIdentifier).Version(version).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ArtifactsAPI.RedirectHarnessArtifact``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**registryIdentifier** | **string** | Unique registry Identifier in a account. | 
**artifact** | **string** | Name of artifact. | 

### Other Parameters

Other parameters are passed through a pointer to a apiRedirectHarnessArtifactRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **accountIdentifier** | **string** | Account Identifier | 
 **version** | **string** | Version | 

### Return type

 (empty response body)

### Authorization

[XApiKeyAuth](../README.md#XApiKeyAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateArtifactLabels

> InlineResponse2001 UpdateArtifactLabels(ctx, registryRef, artifact).ArtifactLabelRequest(artifactLabelRequest).Execute()

Update Artifact Labels



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
	artifactLabelRequest := *openapiclient.NewArtifactLabelRequest([]string{"Labels_example"}) // ArtifactLabelRequest | request to update artifact labels (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ArtifactsAPI.UpdateArtifactLabels(context.Background(), registryRef, artifact).ArtifactLabelRequest(artifactLabelRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ArtifactsAPI.UpdateArtifactLabels``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `UpdateArtifactLabels`: InlineResponse2001
	fmt.Fprintf(os.Stdout, "Response from `ArtifactsAPI.UpdateArtifactLabels`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**registryRef** | **string** | Unique registry path. | 
**artifact** | **string** | Name of artifact. | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateArtifactLabelsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **artifactLabelRequest** | [**ArtifactLabelRequest**](ArtifactLabelRequest.md) | request to update artifact labels | 

### Return type

[**InlineResponse2001**](InlineResponse2001.md)

### Authorization

[XApiKeyAuth](../README.md#XApiKeyAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

