# STODetails

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Critical** | Pointer to **int32** |  | [optional] 
**ExecutionId** | Pointer to **string** |  | [optional] 
**High** | Pointer to **int32** |  | [optional] 
**Ignored** | Pointer to **int32** |  | [optional] 
**Info** | Pointer to **int32** |  | [optional] 
**LastScanned** | Pointer to **string** |  | [optional] 
**Low** | Pointer to **int32** |  | [optional] 
**Medium** | Pointer to **int32** |  | [optional] 
**PipelineId** | Pointer to **string** |  | [optional] 
**Total** | Pointer to **int32** |  | [optional] 

## Methods

### NewSTODetails

`func NewSTODetails() *STODetails`

NewSTODetails instantiates a new STODetails object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSTODetailsWithDefaults

`func NewSTODetailsWithDefaults() *STODetails`

NewSTODetailsWithDefaults instantiates a new STODetails object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCritical

`func (o *STODetails) GetCritical() int32`

GetCritical returns the Critical field if non-nil, zero value otherwise.

### GetCriticalOk

`func (o *STODetails) GetCriticalOk() (*int32, bool)`

GetCriticalOk returns a tuple with the Critical field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCritical

`func (o *STODetails) SetCritical(v int32)`

SetCritical sets Critical field to given value.

### HasCritical

`func (o *STODetails) HasCritical() bool`

HasCritical returns a boolean if a field has been set.

### GetExecutionId

`func (o *STODetails) GetExecutionId() string`

GetExecutionId returns the ExecutionId field if non-nil, zero value otherwise.

### GetExecutionIdOk

`func (o *STODetails) GetExecutionIdOk() (*string, bool)`

GetExecutionIdOk returns a tuple with the ExecutionId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExecutionId

`func (o *STODetails) SetExecutionId(v string)`

SetExecutionId sets ExecutionId field to given value.

### HasExecutionId

`func (o *STODetails) HasExecutionId() bool`

HasExecutionId returns a boolean if a field has been set.

### GetHigh

`func (o *STODetails) GetHigh() int32`

GetHigh returns the High field if non-nil, zero value otherwise.

### GetHighOk

`func (o *STODetails) GetHighOk() (*int32, bool)`

GetHighOk returns a tuple with the High field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHigh

`func (o *STODetails) SetHigh(v int32)`

SetHigh sets High field to given value.

### HasHigh

`func (o *STODetails) HasHigh() bool`

HasHigh returns a boolean if a field has been set.

### GetIgnored

`func (o *STODetails) GetIgnored() int32`

GetIgnored returns the Ignored field if non-nil, zero value otherwise.

### GetIgnoredOk

`func (o *STODetails) GetIgnoredOk() (*int32, bool)`

GetIgnoredOk returns a tuple with the Ignored field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIgnored

`func (o *STODetails) SetIgnored(v int32)`

SetIgnored sets Ignored field to given value.

### HasIgnored

`func (o *STODetails) HasIgnored() bool`

HasIgnored returns a boolean if a field has been set.

### GetInfo

`func (o *STODetails) GetInfo() int32`

GetInfo returns the Info field if non-nil, zero value otherwise.

### GetInfoOk

`func (o *STODetails) GetInfoOk() (*int32, bool)`

GetInfoOk returns a tuple with the Info field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInfo

`func (o *STODetails) SetInfo(v int32)`

SetInfo sets Info field to given value.

### HasInfo

`func (o *STODetails) HasInfo() bool`

HasInfo returns a boolean if a field has been set.

### GetLastScanned

`func (o *STODetails) GetLastScanned() string`

GetLastScanned returns the LastScanned field if non-nil, zero value otherwise.

### GetLastScannedOk

`func (o *STODetails) GetLastScannedOk() (*string, bool)`

GetLastScannedOk returns a tuple with the LastScanned field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastScanned

`func (o *STODetails) SetLastScanned(v string)`

SetLastScanned sets LastScanned field to given value.

### HasLastScanned

`func (o *STODetails) HasLastScanned() bool`

HasLastScanned returns a boolean if a field has been set.

### GetLow

`func (o *STODetails) GetLow() int32`

GetLow returns the Low field if non-nil, zero value otherwise.

### GetLowOk

`func (o *STODetails) GetLowOk() (*int32, bool)`

GetLowOk returns a tuple with the Low field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLow

`func (o *STODetails) SetLow(v int32)`

SetLow sets Low field to given value.

### HasLow

`func (o *STODetails) HasLow() bool`

HasLow returns a boolean if a field has been set.

### GetMedium

`func (o *STODetails) GetMedium() int32`

GetMedium returns the Medium field if non-nil, zero value otherwise.

### GetMediumOk

`func (o *STODetails) GetMediumOk() (*int32, bool)`

GetMediumOk returns a tuple with the Medium field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMedium

`func (o *STODetails) SetMedium(v int32)`

SetMedium sets Medium field to given value.

### HasMedium

`func (o *STODetails) HasMedium() bool`

HasMedium returns a boolean if a field has been set.

### GetPipelineId

`func (o *STODetails) GetPipelineId() string`

GetPipelineId returns the PipelineId field if non-nil, zero value otherwise.

### GetPipelineIdOk

`func (o *STODetails) GetPipelineIdOk() (*string, bool)`

GetPipelineIdOk returns a tuple with the PipelineId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPipelineId

`func (o *STODetails) SetPipelineId(v string)`

SetPipelineId sets PipelineId field to given value.

### HasPipelineId

`func (o *STODetails) HasPipelineId() bool`

HasPipelineId returns a boolean if a field has been set.

### GetTotal

`func (o *STODetails) GetTotal() int32`

GetTotal returns the Total field if non-nil, zero value otherwise.

### GetTotalOk

`func (o *STODetails) GetTotalOk() (*int32, bool)`

GetTotalOk returns a tuple with the Total field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotal

`func (o *STODetails) SetTotal(v int32)`

SetTotal sets Total field to given value.

### HasTotal

`func (o *STODetails) HasTotal() bool`

HasTotal returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


