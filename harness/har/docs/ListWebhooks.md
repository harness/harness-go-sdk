# ListWebhooks

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ItemCount** | Pointer to **int64** | The total number of items | [optional] 
**PageCount** | Pointer to **int64** | The total number of pages | [optional] 
**PageIndex** | Pointer to **int64** | The current page | [optional] 
**PageSize** | Pointer to **int32** | The number of items per page | [optional] 
**Webhooks** | [**[]Webhook**](Webhook.md) | A list of Registries webhooks | 

## Methods

### NewListWebhooks

`func NewListWebhooks(webhooks []Webhook, ) *ListWebhooks`

NewListWebhooks instantiates a new ListWebhooks object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewListWebhooksWithDefaults

`func NewListWebhooksWithDefaults() *ListWebhooks`

NewListWebhooksWithDefaults instantiates a new ListWebhooks object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetItemCount

`func (o *ListWebhooks) GetItemCount() int64`

GetItemCount returns the ItemCount field if non-nil, zero value otherwise.

### GetItemCountOk

`func (o *ListWebhooks) GetItemCountOk() (*int64, bool)`

GetItemCountOk returns a tuple with the ItemCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetItemCount

`func (o *ListWebhooks) SetItemCount(v int64)`

SetItemCount sets ItemCount field to given value.

### HasItemCount

`func (o *ListWebhooks) HasItemCount() bool`

HasItemCount returns a boolean if a field has been set.

### GetPageCount

`func (o *ListWebhooks) GetPageCount() int64`

GetPageCount returns the PageCount field if non-nil, zero value otherwise.

### GetPageCountOk

`func (o *ListWebhooks) GetPageCountOk() (*int64, bool)`

GetPageCountOk returns a tuple with the PageCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPageCount

`func (o *ListWebhooks) SetPageCount(v int64)`

SetPageCount sets PageCount field to given value.

### HasPageCount

`func (o *ListWebhooks) HasPageCount() bool`

HasPageCount returns a boolean if a field has been set.

### GetPageIndex

`func (o *ListWebhooks) GetPageIndex() int64`

GetPageIndex returns the PageIndex field if non-nil, zero value otherwise.

### GetPageIndexOk

`func (o *ListWebhooks) GetPageIndexOk() (*int64, bool)`

GetPageIndexOk returns a tuple with the PageIndex field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPageIndex

`func (o *ListWebhooks) SetPageIndex(v int64)`

SetPageIndex sets PageIndex field to given value.

### HasPageIndex

`func (o *ListWebhooks) HasPageIndex() bool`

HasPageIndex returns a boolean if a field has been set.

### GetPageSize

`func (o *ListWebhooks) GetPageSize() int32`

GetPageSize returns the PageSize field if non-nil, zero value otherwise.

### GetPageSizeOk

`func (o *ListWebhooks) GetPageSizeOk() (*int32, bool)`

GetPageSizeOk returns a tuple with the PageSize field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPageSize

`func (o *ListWebhooks) SetPageSize(v int32)`

SetPageSize sets PageSize field to given value.

### HasPageSize

`func (o *ListWebhooks) HasPageSize() bool`

HasPageSize returns a boolean if a field has been set.

### GetWebhooks

`func (o *ListWebhooks) GetWebhooks() []Webhook`

GetWebhooks returns the Webhooks field if non-nil, zero value otherwise.

### GetWebhooksOk

`func (o *ListWebhooks) GetWebhooksOk() (*[]Webhook, bool)`

GetWebhooksOk returns a tuple with the Webhooks field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWebhooks

`func (o *ListWebhooks) SetWebhooks(v []Webhook)`

SetWebhooks sets Webhooks field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


