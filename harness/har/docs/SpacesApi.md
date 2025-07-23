# \SpacesAPI

All URIs are relative to */gateway/har/api/v1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetAllHarnessArtifacts**](SpacesAPI.md#GetAllHarnessArtifacts) | **Get** /spaces/{space_ref}/artifacts | List Harness Artifacts
[**GetAllRegistries**](SpacesAPI.md#GetAllRegistries) | **Get** /spaces/{space_ref}/registries | List registries
[**GetArtifactStatsForSpace**](SpacesAPI.md#GetArtifactStatsForSpace) | **Get** /spaces/{space_ref}/artifact/stats | Get artifact stats
[**GetStorageDetails**](SpacesAPI.md#GetStorageDetails) | **Get** /spaces/{space_ref}/details | Get storage details for given space



## GetAllHarnessArtifacts

> InlineResponse20025 GetAllHarnessArtifacts(ctx, spaceRef).RegIdentifier(regIdentifier).Page(page).Size(size).SortOrder(sortOrder).SortField(sortField).SearchTerm(searchTerm).LatestVersion(latestVersion).DeployedArtifact(deployedArtifact).PackageType(packageType).Execute()

List Harness Artifacts



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
	spaceRef := "spaceRef_example" // string | Unique space path.
	regIdentifier := []string{"Inner_example"} // []string | Registry Identifier (optional)
	page := int64(789) // int64 | Current page number (optional) (default to 1)
	size := int64(789) // int64 | Number of items per page (optional) (default to 20)
	sortOrder := "sortOrder_example" // string | sortOrder (optional)
	sortField := "sortField_example" // string | sortField (optional)
	searchTerm := "searchTerm_example" // string | search Term. (optional)
	latestVersion := true // bool | Latest Version Filter. (optional)
	deployedArtifact := true // bool | Deployed Artifact Filter. (optional)
	packageType := []string{"Inner_example"} // []string | Registry Package Type (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.SpacesAPI.GetAllHarnessArtifacts(context.Background(), spaceRef).RegIdentifier(regIdentifier).Page(page).Size(size).SortOrder(sortOrder).SortField(sortField).SearchTerm(searchTerm).LatestVersion(latestVersion).DeployedArtifact(deployedArtifact).PackageType(packageType).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SpacesAPI.GetAllHarnessArtifacts``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetAllHarnessArtifacts`: InlineResponse20025
	fmt.Fprintf(os.Stdout, "Response from `SpacesAPI.GetAllHarnessArtifacts`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**spaceRef** | **string** | Unique space path. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetAllHarnessArtifactsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **regIdentifier** | **[]string** | Registry Identifier | 
 **page** | **int64** | Current page number | [default to 1]
 **size** | **int64** | Number of items per page | [default to 20]
 **sortOrder** | **string** | sortOrder | 
 **sortField** | **string** | sortField | 
 **searchTerm** | **string** | search Term. | 
 **latestVersion** | **bool** | Latest Version Filter. | 
 **deployedArtifact** | **bool** | Deployed Artifact Filter. | 
 **packageType** | **[]string** | Registry Package Type | 

### Return type

[**InlineResponse20025**](InlineResponse20025.md)

### Authorization

[XApiKeyAuth](../README.md#XApiKeyAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetAllRegistries

> InlineResponse20027 GetAllRegistries(ctx, spaceRef).PackageType(packageType).Type_(type_).Page(page).Size(size).SortOrder(sortOrder).SortField(sortField).SearchTerm(searchTerm).Recursive(recursive).Execute()

List registries



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
	spaceRef := "spaceRef_example" // string | Unique space path.
	packageType := []string{"Inner_example"} // []string | Registry Package Type (optional)
	type_ := "type__example" // string | Registry Type (optional)
	page := int64(789) // int64 | Current page number (optional) (default to 1)
	size := int64(789) // int64 | Number of items per page (optional) (default to 20)
	sortOrder := "sortOrder_example" // string | sortOrder (optional)
	sortField := "sortField_example" // string | sortField (optional)
	searchTerm := "searchTerm_example" // string | search Term. (optional)
	recursive := true // bool | Whether to list registries recursively. (optional) (default to false)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.SpacesAPI.GetAllRegistries(context.Background(), spaceRef).PackageType(packageType).Type_(type_).Page(page).Size(size).SortOrder(sortOrder).SortField(sortField).SearchTerm(searchTerm).Recursive(recursive).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SpacesAPI.GetAllRegistries``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetAllRegistries`: InlineResponse20027
	fmt.Fprintf(os.Stdout, "Response from `SpacesAPI.GetAllRegistries`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**spaceRef** | **string** | Unique space path. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetAllRegistriesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **packageType** | **[]string** | Registry Package Type | 
 **type_** | **string** | Registry Type | 
 **page** | **int64** | Current page number | [default to 1]
 **size** | **int64** | Number of items per page | [default to 20]
 **sortOrder** | **string** | sortOrder | 
 **sortField** | **string** | sortField | 
 **searchTerm** | **string** | search Term. | 
 **recursive** | **bool** | Whether to list registries recursively. | [default to false]

### Return type

[**InlineResponse20027**](InlineResponse20027.md)

### Authorization

[XApiKeyAuth](../README.md#XApiKeyAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetArtifactStatsForSpace

> InlineResponse2002 GetArtifactStatsForSpace(ctx, spaceRef).From(from).To(to).Execute()

Get artifact stats



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
	spaceRef := "spaceRef_example" // string | Unique space path.
	from := "from_example" // string | Date. Format - MM/DD/YYYY (optional)
	to := "to_example" // string | Date. Format - MM/DD/YYYY (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.SpacesAPI.GetArtifactStatsForSpace(context.Background(), spaceRef).From(from).To(to).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SpacesAPI.GetArtifactStatsForSpace``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetArtifactStatsForSpace`: InlineResponse2002
	fmt.Fprintf(os.Stdout, "Response from `SpacesAPI.GetArtifactStatsForSpace`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**spaceRef** | **string** | Unique space path. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetArtifactStatsForSpaceRequest struct via the builder pattern


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


## GetStorageDetails

> InlineResponse20026 GetStorageDetails(ctx, spaceRef).Execute()

Get storage details for given space



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
	spaceRef := "spaceRef_example" // string | Unique space path.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.SpacesAPI.GetStorageDetails(context.Background(), spaceRef).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SpacesAPI.GetStorageDetails``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetStorageDetails`: InlineResponse20026
	fmt.Fprintf(os.Stdout, "Response from `SpacesAPI.GetStorageDetails`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**spaceRef** | **string** | Unique space path. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetStorageDetailsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**InlineResponse20026**](InlineResponse20026.md)

### Authorization

[XApiKeyAuth](../README.md#XApiKeyAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

