# \WebhooksAPI

All URIs are relative to */gateway/har/api/v1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateWebhook**](WebhooksAPI.md#CreateWebhook) | **Post** /registry/{registry_ref}/+/webhooks | CreateWebhook
[**DeleteWebhook**](WebhooksAPI.md#DeleteWebhook) | **Delete** /registry/{registry_ref}/+/webhooks/{webhook_identifier} | DeleteWebhook
[**GetWebhook**](WebhooksAPI.md#GetWebhook) | **Get** /registry/{registry_ref}/+/webhooks/{webhook_identifier} | GetWebhook
[**GetWebhookExecution**](WebhooksAPI.md#GetWebhookExecution) | **Get** /registry/{registry_ref}/+/webhooks/{webhook_identifier}/executions/{webhook_execution_id} | GetWebhookExecution
[**ListWebhookExecutions**](WebhooksAPI.md#ListWebhookExecutions) | **Get** /registry/{registry_ref}/+/webhooks/{webhook_identifier}/executions | ListWebhookExecutions
[**ListWebhooks**](WebhooksAPI.md#ListWebhooks) | **Get** /registry/{registry_ref}/+/webhooks | ListWebhooks
[**ReTriggerWebhookExecution**](WebhooksAPI.md#ReTriggerWebhookExecution) | **Get** /registry/{registry_ref}/+/webhooks/{webhook_identifier}/executions/{webhook_execution_id}/retrigger | ReTriggerWebhookExecution
[**UpdateWebhook**](WebhooksAPI.md#UpdateWebhook) | **Put** /registry/{registry_ref}/+/webhooks/{webhook_identifier} | UpdateWebhook



## CreateWebhook

> InlineResponse2011 CreateWebhook(ctx, registryRef).WebhookRequest(webhookRequest).Execute()

CreateWebhook



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
	webhookRequest := *openapiclient.NewWebhookRequest(false, "Identifier_example", false, "Name_example", "Url_example") // WebhookRequest | request for create and update webhook (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.WebhooksAPI.CreateWebhook(context.Background(), registryRef).WebhookRequest(webhookRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `WebhooksAPI.CreateWebhook``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CreateWebhook`: InlineResponse2011
	fmt.Fprintf(os.Stdout, "Response from `WebhooksAPI.CreateWebhook`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**registryRef** | **string** | Unique registry path. | 

### Other Parameters

Other parameters are passed through a pointer to a apiCreateWebhookRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **webhookRequest** | [**WebhookRequest**](WebhookRequest.md) | request for create and update webhook | 

### Return type

[**InlineResponse2011**](InlineResponse2011.md)

### Authorization

[XApiKeyAuth](../README.md#XApiKeyAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteWebhook

> InlineResponse200 DeleteWebhook(ctx, registryRef, webhookIdentifier).Execute()

DeleteWebhook



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
	webhookIdentifier := "webhookIdentifier_example" // string | Unique webhook identifier.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.WebhooksAPI.DeleteWebhook(context.Background(), registryRef, webhookIdentifier).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `WebhooksAPI.DeleteWebhook``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `DeleteWebhook`: InlineResponse200
	fmt.Fprintf(os.Stdout, "Response from `WebhooksAPI.DeleteWebhook`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**registryRef** | **string** | Unique registry path. | 
**webhookIdentifier** | **string** | Unique webhook identifier. | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteWebhookRequest struct via the builder pattern


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


## GetWebhook

> InlineResponse2011 GetWebhook(ctx, registryRef, webhookIdentifier).Execute()

GetWebhook



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
	webhookIdentifier := "webhookIdentifier_example" // string | Unique webhook identifier.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.WebhooksAPI.GetWebhook(context.Background(), registryRef, webhookIdentifier).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `WebhooksAPI.GetWebhook``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetWebhook`: InlineResponse2011
	fmt.Fprintf(os.Stdout, "Response from `WebhooksAPI.GetWebhook`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**registryRef** | **string** | Unique registry path. | 
**webhookIdentifier** | **string** | Unique webhook identifier. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetWebhookRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**InlineResponse2011**](InlineResponse2011.md)

### Authorization

[XApiKeyAuth](../README.md#XApiKeyAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetWebhookExecution

> InlineResponse20021 GetWebhookExecution(ctx, registryRef, webhookIdentifier, webhookExecutionId).Execute()

GetWebhookExecution



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
	webhookIdentifier := "webhookIdentifier_example" // string | Unique webhook identifier.
	webhookExecutionId := "webhookExecutionId_example" // string | Unique webhook execution identifier.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.WebhooksAPI.GetWebhookExecution(context.Background(), registryRef, webhookIdentifier, webhookExecutionId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `WebhooksAPI.GetWebhookExecution``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetWebhookExecution`: InlineResponse20021
	fmt.Fprintf(os.Stdout, "Response from `WebhooksAPI.GetWebhookExecution`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**registryRef** | **string** | Unique registry path. | 
**webhookIdentifier** | **string** | Unique webhook identifier. | 
**webhookExecutionId** | **string** | Unique webhook execution identifier. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetWebhookExecutionRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------




### Return type

[**InlineResponse20021**](InlineResponse20021.md)

### Authorization

[XApiKeyAuth](../README.md#XApiKeyAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListWebhookExecutions

> InlineResponse20020 ListWebhookExecutions(ctx, registryRef, webhookIdentifier).Page(page).Size(size).Execute()

ListWebhookExecutions



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
	webhookIdentifier := "webhookIdentifier_example" // string | Unique webhook identifier.
	page := int64(789) // int64 | Current page number (optional) (default to 1)
	size := int64(789) // int64 | Number of items per page (optional) (default to 20)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.WebhooksAPI.ListWebhookExecutions(context.Background(), registryRef, webhookIdentifier).Page(page).Size(size).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `WebhooksAPI.ListWebhookExecutions``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ListWebhookExecutions`: InlineResponse20020
	fmt.Fprintf(os.Stdout, "Response from `WebhooksAPI.ListWebhookExecutions`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**registryRef** | **string** | Unique registry path. | 
**webhookIdentifier** | **string** | Unique webhook identifier. | 

### Other Parameters

Other parameters are passed through a pointer to a apiListWebhookExecutionsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **page** | **int64** | Current page number | [default to 1]
 **size** | **int64** | Number of items per page | [default to 20]

### Return type

[**InlineResponse20020**](InlineResponse20020.md)

### Authorization

[XApiKeyAuth](../README.md#XApiKeyAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListWebhooks

> InlineResponse20019 ListWebhooks(ctx, registryRef).Page(page).Size(size).SortOrder(sortOrder).SortField(sortField).SearchTerm(searchTerm).Execute()

ListWebhooks



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
	sortOrder := "sortOrder_example" // string | sortOrder (optional)
	sortField := "sortField_example" // string | sortField (optional)
	searchTerm := "searchTerm_example" // string | search Term. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.WebhooksAPI.ListWebhooks(context.Background(), registryRef).Page(page).Size(size).SortOrder(sortOrder).SortField(sortField).SearchTerm(searchTerm).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `WebhooksAPI.ListWebhooks``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ListWebhooks`: InlineResponse20019
	fmt.Fprintf(os.Stdout, "Response from `WebhooksAPI.ListWebhooks`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**registryRef** | **string** | Unique registry path. | 

### Other Parameters

Other parameters are passed through a pointer to a apiListWebhooksRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **page** | **int64** | Current page number | [default to 1]
 **size** | **int64** | Number of items per page | [default to 20]
 **sortOrder** | **string** | sortOrder | 
 **sortField** | **string** | sortField | 
 **searchTerm** | **string** | search Term. | 

### Return type

[**InlineResponse20019**](InlineResponse20019.md)

### Authorization

[XApiKeyAuth](../README.md#XApiKeyAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ReTriggerWebhookExecution

> InlineResponse20021 ReTriggerWebhookExecution(ctx, registryRef, webhookIdentifier, webhookExecutionId).Execute()

ReTriggerWebhookExecution



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
	webhookIdentifier := "webhookIdentifier_example" // string | Unique webhook identifier.
	webhookExecutionId := "webhookExecutionId_example" // string | Unique webhook execution identifier.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.WebhooksAPI.ReTriggerWebhookExecution(context.Background(), registryRef, webhookIdentifier, webhookExecutionId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `WebhooksAPI.ReTriggerWebhookExecution``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ReTriggerWebhookExecution`: InlineResponse20021
	fmt.Fprintf(os.Stdout, "Response from `WebhooksAPI.ReTriggerWebhookExecution`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**registryRef** | **string** | Unique registry path. | 
**webhookIdentifier** | **string** | Unique webhook identifier. | 
**webhookExecutionId** | **string** | Unique webhook execution identifier. | 

### Other Parameters

Other parameters are passed through a pointer to a apiReTriggerWebhookExecutionRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------




### Return type

[**InlineResponse20021**](InlineResponse20021.md)

### Authorization

[XApiKeyAuth](../README.md#XApiKeyAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateWebhook

> InlineResponse2011 UpdateWebhook(ctx, registryRef, webhookIdentifier).WebhookRequest(webhookRequest).Execute()

UpdateWebhook



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
	webhookIdentifier := "webhookIdentifier_example" // string | Unique webhook identifier.
	webhookRequest := *openapiclient.NewWebhookRequest(false, "Identifier_example", false, "Name_example", "Url_example") // WebhookRequest | request for create and update webhook (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.WebhooksAPI.UpdateWebhook(context.Background(), registryRef, webhookIdentifier).WebhookRequest(webhookRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `WebhooksAPI.UpdateWebhook``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `UpdateWebhook`: InlineResponse2011
	fmt.Fprintf(os.Stdout, "Response from `WebhooksAPI.UpdateWebhook`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**registryRef** | **string** | Unique registry path. | 
**webhookIdentifier** | **string** | Unique webhook identifier. | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateWebhookRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **webhookRequest** | [**WebhookRequest**](WebhookRequest.md) | request for create and update webhook | 

### Return type

[**InlineResponse2011**](InlineResponse2011.md)

### Authorization

[XApiKeyAuth](../README.md#XApiKeyAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

