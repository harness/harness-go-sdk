# ListWebhooksExecutions

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Executions** | [**[]WebhookExecution**](WebhookExecution.md) | A list of Registries webhooks executions | 
**ItemCount** | Pointer to **int64** | The total number of items | [optional] 
**PageCount** | Pointer to **int64** | The total number of pages | [optional] 
**PageIndex** | Pointer to **int64** | The current page | [optional] 
**PageSize** | Pointer to **int32** | The number of items per page | [optional] 

## Methods

### NewListWebhooksExecutions

`func NewListWebhooksExecutions(executions []WebhookExecution, ) *ListWebhooksExecutions`

NewListWebhooksExecutions instantiates a new ListWebhooksExecutions object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewListWebhooksExecutionsWithDefaults

`func NewListWebhooksExecutionsWithDefaults() *ListWebhooksExecutions`

NewListWebhooksExecutionsWithDefaults instantiates a new ListWebhooksExecutions object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetExecutions

`func (o *ListWebhooksExecutions) GetExecutions() []WebhookExecution`

GetExecutions returns the Executions field if non-nil, zero value otherwise.

### GetExecutionsOk

`func (o *ListWebhooksExecutions) GetExecutionsOk() (*[]WebhookExecution, bool)`

GetExecutionsOk returns a tuple with the Executions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExecutions

`func (o *ListWebhooksExecutions) SetExecutions(v []WebhookExecution)`

SetExecutions sets Executions field to given value.


### GetItemCount

`func (o *ListWebhooksExecutions) GetItemCount() int64`

GetItemCount returns the ItemCount field if non-nil, zero value otherwise.

### GetItemCountOk

`func (o *ListWebhooksExecutions) GetItemCountOk() (*int64, bool)`

GetItemCountOk returns a tuple with the ItemCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetItemCount

`func (o *ListWebhooksExecutions) SetItemCount(v int64)`

SetItemCount sets ItemCount field to given value.

### HasItemCount

`func (o *ListWebhooksExecutions) HasItemCount() bool`

HasItemCount returns a boolean if a field has been set.

### GetPageCount

`func (o *ListWebhooksExecutions) GetPageCount() int64`

GetPageCount returns the PageCount field if non-nil, zero value otherwise.

### GetPageCountOk

`func (o *ListWebhooksExecutions) GetPageCountOk() (*int64, bool)`

GetPageCountOk returns a tuple with the PageCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPageCount

`func (o *ListWebhooksExecutions) SetPageCount(v int64)`

SetPageCount sets PageCount field to given value.

### HasPageCount

`func (o *ListWebhooksExecutions) HasPageCount() bool`

HasPageCount returns a boolean if a field has been set.

### GetPageIndex

`func (o *ListWebhooksExecutions) GetPageIndex() int64`

GetPageIndex returns the PageIndex field if non-nil, zero value otherwise.

### GetPageIndexOk

`func (o *ListWebhooksExecutions) GetPageIndexOk() (*int64, bool)`

GetPageIndexOk returns a tuple with the PageIndex field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPageIndex

`func (o *ListWebhooksExecutions) SetPageIndex(v int64)`

SetPageIndex sets PageIndex field to given value.

### HasPageIndex

`func (o *ListWebhooksExecutions) HasPageIndex() bool`

HasPageIndex returns a boolean if a field has been set.

### GetPageSize

`func (o *ListWebhooksExecutions) GetPageSize() int32`

GetPageSize returns the PageSize field if non-nil, zero value otherwise.

### GetPageSizeOk

`func (o *ListWebhooksExecutions) GetPageSizeOk() (*int32, bool)`

GetPageSizeOk returns a tuple with the PageSize field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPageSize

`func (o *ListWebhooksExecutions) SetPageSize(v int32)`

SetPageSize sets PageSize field to given value.

### HasPageSize

`func (o *ListWebhooksExecutions) HasPageSize() bool`

HasPageSize returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


