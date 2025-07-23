# \RegistriesAPI

All URIs are relative to */gateway/har/api/v1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateRegistry**](RegistriesAPI.md#CreateRegistry) | **Post** /registry | Create Registry.
[**DeleteRegistry**](RegistriesAPI.md#DeleteRegistry) | **Delete** /registry/{registry_ref}/+ | Delete a Registry
[**GetAllArtifactsByRegistry**](RegistriesAPI.md#GetAllArtifactsByRegistry) | **Get** /registry/{registry_ref}/+/artifacts | List Artifacts for Registry
[**GetClientSetupDetails**](RegistriesAPI.md#GetClientSetupDetails) | **Get** /registry/{registry_ref}/+/client-setup-details | Returns CLI Client Setup Details
[**GetRegistry**](RegistriesAPI.md#GetRegistry) | **Get** /registry/{registry_ref}/+ | Returns Registry Details
[**ModifyRegistry**](RegistriesAPI.md#ModifyRegistry) | **Put** /registry/{registry_ref}/+ | Updates a Registry



## CreateRegistry

> InlineResponse201 CreateRegistry(ctx).SpaceRef(spaceRef).RegistryRequest(registryRequest).Execute()

Create Registry.



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
	spaceRef := "spaceRef_example" // string | Unique space path (optional)
	registryRequest := *openapiclient.NewRegistryRequest("Identifier_example", openapiclient.PackageType("DOCKER")) // RegistryRequest | request for create and update registry (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.RegistriesAPI.CreateRegistry(context.Background()).SpaceRef(spaceRef).RegistryRequest(registryRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `RegistriesAPI.CreateRegistry``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CreateRegistry`: InlineResponse201
	fmt.Fprintf(os.Stdout, "Response from `RegistriesAPI.CreateRegistry`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateRegistryRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **spaceRef** | **string** | Unique space path | 
 **registryRequest** | [**RegistryRequest**](RegistryRequest.md) | request for create and update registry | 

### Return type

[**InlineResponse201**](InlineResponse201.md)

### Authorization

[XApiKeyAuth](../README.md#XApiKeyAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteRegistry

> InlineResponse200 DeleteRegistry(ctx, registryRef).Execute()

Delete a Registry



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

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.RegistriesAPI.DeleteRegistry(context.Background(), registryRef).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `RegistriesAPI.DeleteRegistry``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `DeleteRegistry`: InlineResponse200
	fmt.Fprintf(os.Stdout, "Response from `RegistriesAPI.DeleteRegistry`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**registryRef** | **string** | Unique registry path. | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteRegistryRequest struct via the builder pattern


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


## GetAllArtifactsByRegistry

> InlineResponse20017 GetAllArtifactsByRegistry(ctx, registryRef).Label(label).Page(page).Size(size).SortOrder(sortOrder).SortField(sortField).SearchTerm(searchTerm).Execute()

List Artifacts for Registry



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
	label := []string{"Inner_example"} // []string | Label. (optional)
	page := int64(789) // int64 | Current page number (optional) (default to 1)
	size := int64(789) // int64 | Number of items per page (optional) (default to 20)
	sortOrder := "sortOrder_example" // string | sortOrder (optional)
	sortField := "sortField_example" // string | sortField (optional)
	searchTerm := "searchTerm_example" // string | search Term. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.RegistriesAPI.GetAllArtifactsByRegistry(context.Background(), registryRef).Label(label).Page(page).Size(size).SortOrder(sortOrder).SortField(sortField).SearchTerm(searchTerm).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `RegistriesAPI.GetAllArtifactsByRegistry``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetAllArtifactsByRegistry`: InlineResponse20017
	fmt.Fprintf(os.Stdout, "Response from `RegistriesAPI.GetAllArtifactsByRegistry`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**registryRef** | **string** | Unique registry path. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetAllArtifactsByRegistryRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **label** | **[]string** | Label. | 
 **page** | **int64** | Current page number | [default to 1]
 **size** | **int64** | Number of items per page | [default to 20]
 **sortOrder** | **string** | sortOrder | 
 **sortField** | **string** | sortField | 
 **searchTerm** | **string** | search Term. | 

### Return type

[**InlineResponse20017**](InlineResponse20017.md)

### Authorization

[XApiKeyAuth](../README.md#XApiKeyAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetClientSetupDetails

> InlineResponse20018 GetClientSetupDetails(ctx, registryRef).Artifact(artifact).Version(version).Execute()

Returns CLI Client Setup Details



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
	artifact := "artifact_example" // string | Artifat (optional)
	version := "version_example" // string | Version (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.RegistriesAPI.GetClientSetupDetails(context.Background(), registryRef).Artifact(artifact).Version(version).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `RegistriesAPI.GetClientSetupDetails``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetClientSetupDetails`: InlineResponse20018
	fmt.Fprintf(os.Stdout, "Response from `RegistriesAPI.GetClientSetupDetails`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**registryRef** | **string** | Unique registry path. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetClientSetupDetailsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **artifact** | **string** | Artifat | 
 **version** | **string** | Version | 

### Return type

[**InlineResponse20018**](InlineResponse20018.md)

### Authorization

[XApiKeyAuth](../README.md#XApiKeyAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetRegistry

> InlineResponse201 GetRegistry(ctx, registryRef).Execute()

Returns Registry Details



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

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.RegistriesAPI.GetRegistry(context.Background(), registryRef).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `RegistriesAPI.GetRegistry``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetRegistry`: InlineResponse201
	fmt.Fprintf(os.Stdout, "Response from `RegistriesAPI.GetRegistry`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**registryRef** | **string** | Unique registry path. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetRegistryRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**InlineResponse201**](InlineResponse201.md)

### Authorization

[XApiKeyAuth](../README.md#XApiKeyAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ModifyRegistry

> InlineResponse201 ModifyRegistry(ctx, registryRef).RegistryRequest(registryRequest).Execute()

Updates a Registry



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
	registryRequest := *openapiclient.NewRegistryRequest("Identifier_example", openapiclient.PackageType("DOCKER")) // RegistryRequest | request for create and update registry (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.RegistriesAPI.ModifyRegistry(context.Background(), registryRef).RegistryRequest(registryRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `RegistriesAPI.ModifyRegistry``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ModifyRegistry`: InlineResponse201
	fmt.Fprintf(os.Stdout, "Response from `RegistriesAPI.ModifyRegistry`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**registryRef** | **string** | Unique registry path. | 

### Other Parameters

Other parameters are passed through a pointer to a apiModifyRegistryRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **registryRequest** | [**RegistryRequest**](RegistryRequest.md) | request for create and update registry | 

### Return type

[**InlineResponse201**](InlineResponse201.md)

### Authorization

[XApiKeyAuth](../README.md#XApiKeyAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

