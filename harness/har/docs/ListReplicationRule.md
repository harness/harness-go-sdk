# ListReplicationRule

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ItemCount** | **int64** | The total number of items | 
**PageCount** | **int64** | The total number of pages | 
**PageIndex** | **int64** | The current page | 
**PageSize** | **int32** | The number of items per page | 
**Rules** | [**[]ReplicationRule**](ReplicationRule.md) | A list of Replication Rules | 

## Methods

### NewListReplicationRule

`func NewListReplicationRule(itemCount int64, pageCount int64, pageIndex int64, pageSize int32, rules []ReplicationRule, ) *ListReplicationRule`

NewListReplicationRule instantiates a new ListReplicationRule object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewListReplicationRuleWithDefaults

`func NewListReplicationRuleWithDefaults() *ListReplicationRule`

NewListReplicationRuleWithDefaults instantiates a new ListReplicationRule object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetItemCount

`func (o *ListReplicationRule) GetItemCount() int64`

GetItemCount returns the ItemCount field if non-nil, zero value otherwise.

### GetItemCountOk

`func (o *ListReplicationRule) GetItemCountOk() (*int64, bool)`

GetItemCountOk returns a tuple with the ItemCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetItemCount

`func (o *ListReplicationRule) SetItemCount(v int64)`

SetItemCount sets ItemCount field to given value.


### GetPageCount

`func (o *ListReplicationRule) GetPageCount() int64`

GetPageCount returns the PageCount field if non-nil, zero value otherwise.

### GetPageCountOk

`func (o *ListReplicationRule) GetPageCountOk() (*int64, bool)`

GetPageCountOk returns a tuple with the PageCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPageCount

`func (o *ListReplicationRule) SetPageCount(v int64)`

SetPageCount sets PageCount field to given value.


### GetPageIndex

`func (o *ListReplicationRule) GetPageIndex() int64`

GetPageIndex returns the PageIndex field if non-nil, zero value otherwise.

### GetPageIndexOk

`func (o *ListReplicationRule) GetPageIndexOk() (*int64, bool)`

GetPageIndexOk returns a tuple with the PageIndex field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPageIndex

`func (o *ListReplicationRule) SetPageIndex(v int64)`

SetPageIndex sets PageIndex field to given value.


### GetPageSize

`func (o *ListReplicationRule) GetPageSize() int32`

GetPageSize returns the PageSize field if non-nil, zero value otherwise.

### GetPageSizeOk

`func (o *ListReplicationRule) GetPageSizeOk() (*int32, bool)`

GetPageSizeOk returns a tuple with the PageSize field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPageSize

`func (o *ListReplicationRule) SetPageSize(v int32)`

SetPageSize sets PageSize field to given value.


### GetRules

`func (o *ListReplicationRule) GetRules() []ReplicationRule`

GetRules returns the Rules field if non-nil, zero value otherwise.

### GetRulesOk

`func (o *ListReplicationRule) GetRulesOk() (*[]ReplicationRule, bool)`

GetRulesOk returns a tuple with the Rules field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRules

`func (o *ListReplicationRule) SetRules(v []ReplicationRule)`

SetRules sets Rules field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


