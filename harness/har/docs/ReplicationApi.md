# \ReplicationAPI

All URIs are relative to */gateway/har/api/v1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateReplicationRule**](ReplicationAPI.md#CreateReplicationRule) | **Post** /replication/rules | Create a replication rule
[**DeleteReplicationRule**](ReplicationAPI.md#DeleteReplicationRule) | **Delete** /replication/rules/{id} | Delete a replication rule
[**GetMigrationLogsForImage**](ReplicationAPI.md#GetMigrationLogsForImage) | **Get** /replication/rules/{id}/migration/images/{image_id}/logs | Get migration logs for an image
[**GetReplicationRule**](ReplicationAPI.md#GetReplicationRule) | **Get** /replication/rules/{id} | Get a replication rule
[**ListMigrationImages**](ReplicationAPI.md#ListMigrationImages) | **Get** /replication/rules/{id}/migration/images | List migration images
[**ListReplicationRules**](ReplicationAPI.md#ListReplicationRules) | **Get** /replication/rules | List replication rules
[**StartMigration**](ReplicationAPI.md#StartMigration) | **Post** /replication/rules/{id}/migration/start | Start migration
[**StopMigration**](ReplicationAPI.md#StopMigration) | **Post** /replication/rules/{id}/migration/stop | Stop migration
[**UpdateReplicationRule**](ReplicationAPI.md#UpdateReplicationRule) | **Put** /replication/rules/{id} | Update a replication rule



## CreateReplicationRule

> InlineResponse20023 CreateReplicationRule(ctx).SpaceRef(spaceRef).ReplicationRuleRequest(replicationRuleRequest).Execute()

Create a replication rule



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
	replicationRuleRequest := *openapiclient.NewReplicationRuleRequest([]string{"AllowedPatterns_example"}, []string{"BlockedPatterns_example"}, openapiclient.ReplicationRegistry{JfrogReplicationRegistry: openapiclient.NewJfrogReplicationRegistry("Namespace_example", "Url_example")}, "DestinationType_example", openapiclient.ReplicationRegistry{JfrogReplicationRegistry: openapiclient.NewJfrogReplicationRegistry("Namespace_example", "Url_example")}, "SourceType_example") // ReplicationRuleRequest | request for create and update replication rule (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ReplicationAPI.CreateReplicationRule(context.Background()).SpaceRef(spaceRef).ReplicationRuleRequest(replicationRuleRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ReplicationAPI.CreateReplicationRule``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CreateReplicationRule`: InlineResponse20023
	fmt.Fprintf(os.Stdout, "Response from `ReplicationAPI.CreateReplicationRule`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateReplicationRuleRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **spaceRef** | **string** | Unique space path | 
 **replicationRuleRequest** | [**ReplicationRuleRequest**](ReplicationRuleRequest.md) | request for create and update replication rule | 

### Return type

[**InlineResponse20023**](InlineResponse20023.md)

### Authorization

[XApiKeyAuth](../README.md#XApiKeyAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteReplicationRule

> InlineResponse200 DeleteReplicationRule(ctx, id).Execute()

Delete a replication rule



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
	id := "id_example" // string | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ReplicationAPI.DeleteReplicationRule(context.Background(), id).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ReplicationAPI.DeleteReplicationRule``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `DeleteReplicationRule`: InlineResponse200
	fmt.Fprintf(os.Stdout, "Response from `ReplicationAPI.DeleteReplicationRule`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteReplicationRuleRequest struct via the builder pattern


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


## GetMigrationLogsForImage

> string GetMigrationLogsForImage(ctx, id, imageId).Execute()

Get migration logs for an image

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
	id := "id_example" // string | 
	imageId := "imageId_example" // string | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ReplicationAPI.GetMigrationLogsForImage(context.Background(), id, imageId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ReplicationAPI.GetMigrationLogsForImage``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetMigrationLogsForImage`: string
	fmt.Fprintf(os.Stdout, "Response from `ReplicationAPI.GetMigrationLogsForImage`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** |  | 
**imageId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetMigrationLogsForImageRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

**string**

### Authorization

[XApiKeyAuth](../README.md#XApiKeyAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: text/plain; charset=utf-8, application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetReplicationRule

> InlineResponse20023 GetReplicationRule(ctx, id).Execute()

Get a replication rule



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
	id := "id_example" // string | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ReplicationAPI.GetReplicationRule(context.Background(), id).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ReplicationAPI.GetReplicationRule``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetReplicationRule`: InlineResponse20023
	fmt.Fprintf(os.Stdout, "Response from `ReplicationAPI.GetReplicationRule`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetReplicationRuleRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**InlineResponse20023**](InlineResponse20023.md)

### Authorization

[XApiKeyAuth](../README.md#XApiKeyAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListMigrationImages

> InlineResponse20024 ListMigrationImages(ctx, id).Page(page).Size(size).SortOrder(sortOrder).SortField(sortField).Execute()

List migration images



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
	id := "id_example" // string | 
	page := int64(789) // int64 | Current page number (optional) (default to 1)
	size := int64(789) // int64 | Number of items per page (optional) (default to 20)
	sortOrder := "sortOrder_example" // string | sortOrder (optional)
	sortField := "sortField_example" // string | sortField (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ReplicationAPI.ListMigrationImages(context.Background(), id).Page(page).Size(size).SortOrder(sortOrder).SortField(sortField).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ReplicationAPI.ListMigrationImages``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ListMigrationImages`: InlineResponse20024
	fmt.Fprintf(os.Stdout, "Response from `ReplicationAPI.ListMigrationImages`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiListMigrationImagesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **page** | **int64** | Current page number | [default to 1]
 **size** | **int64** | Number of items per page | [default to 20]
 **sortOrder** | **string** | sortOrder | 
 **sortField** | **string** | sortField | 

### Return type

[**InlineResponse20024**](InlineResponse20024.md)

### Authorization

[XApiKeyAuth](../README.md#XApiKeyAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListReplicationRules

> InlineResponse20022 ListReplicationRules(ctx).SpaceRef(spaceRef).Execute()

List replication rules



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

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ReplicationAPI.ListReplicationRules(context.Background()).SpaceRef(spaceRef).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ReplicationAPI.ListReplicationRules``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ListReplicationRules`: InlineResponse20022
	fmt.Fprintf(os.Stdout, "Response from `ReplicationAPI.ListReplicationRules`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiListReplicationRulesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **spaceRef** | **string** | Unique space path | 

### Return type

[**InlineResponse20022**](InlineResponse20022.md)

### Authorization

[XApiKeyAuth](../README.md#XApiKeyAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## StartMigration

> InlineResponse200 StartMigration(ctx, id).Execute()

Start migration



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
	id := "id_example" // string | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ReplicationAPI.StartMigration(context.Background(), id).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ReplicationAPI.StartMigration``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `StartMigration`: InlineResponse200
	fmt.Fprintf(os.Stdout, "Response from `ReplicationAPI.StartMigration`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiStartMigrationRequest struct via the builder pattern


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


## StopMigration

> InlineResponse200 StopMigration(ctx, id).Execute()

Stop migration



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
	id := "id_example" // string | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ReplicationAPI.StopMigration(context.Background(), id).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ReplicationAPI.StopMigration``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `StopMigration`: InlineResponse200
	fmt.Fprintf(os.Stdout, "Response from `ReplicationAPI.StopMigration`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiStopMigrationRequest struct via the builder pattern


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


## UpdateReplicationRule

> InlineResponse20023 UpdateReplicationRule(ctx, id).ReplicationRuleRequest(replicationRuleRequest).Execute()

Update a replication rule



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
	id := "id_example" // string | 
	replicationRuleRequest := *openapiclient.NewReplicationRuleRequest([]string{"AllowedPatterns_example"}, []string{"BlockedPatterns_example"}, openapiclient.ReplicationRegistry{JfrogReplicationRegistry: openapiclient.NewJfrogReplicationRegistry("Namespace_example", "Url_example")}, "DestinationType_example", openapiclient.ReplicationRegistry{JfrogReplicationRegistry: openapiclient.NewJfrogReplicationRegistry("Namespace_example", "Url_example")}, "SourceType_example") // ReplicationRuleRequest | request for create and update replication rule (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ReplicationAPI.UpdateReplicationRule(context.Background(), id).ReplicationRuleRequest(replicationRuleRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ReplicationAPI.UpdateReplicationRule``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `UpdateReplicationRule`: InlineResponse20023
	fmt.Fprintf(os.Stdout, "Response from `ReplicationAPI.UpdateReplicationRule`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateReplicationRuleRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **replicationRuleRequest** | [**ReplicationRuleRequest**](ReplicationRuleRequest.md) | request for create and update replication rule | 

### Return type

[**InlineResponse20023**](InlineResponse20023.md)

### Authorization

[XApiKeyAuth](../README.md#XApiKeyAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

