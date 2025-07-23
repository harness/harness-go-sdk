# ClientSetupSection

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Header** | Pointer to **string** |  | [optional] 
**SecHeader** | Pointer to **string** |  | [optional] 
**Type** | [**SectionType**](SectionType.md) |  | 
**Steps** | Pointer to [**[]ClientSetupStep**](ClientSetupStep.md) |  | [optional] 
**Tabs** | Pointer to [**[]TabSetupStep**](TabSetupStep.md) |  | [optional] 

## Methods

### NewClientSetupSection

`func NewClientSetupSection(type_ SectionType, ) *ClientSetupSection`

NewClientSetupSection instantiates a new ClientSetupSection object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewClientSetupSectionWithDefaults

`func NewClientSetupSectionWithDefaults() *ClientSetupSection`

NewClientSetupSectionWithDefaults instantiates a new ClientSetupSection object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetHeader

`func (o *ClientSetupSection) GetHeader() string`

GetHeader returns the Header field if non-nil, zero value otherwise.

### GetHeaderOk

`func (o *ClientSetupSection) GetHeaderOk() (*string, bool)`

GetHeaderOk returns a tuple with the Header field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHeader

`func (o *ClientSetupSection) SetHeader(v string)`

SetHeader sets Header field to given value.

### HasHeader

`func (o *ClientSetupSection) HasHeader() bool`

HasHeader returns a boolean if a field has been set.

### GetSecHeader

`func (o *ClientSetupSection) GetSecHeader() string`

GetSecHeader returns the SecHeader field if non-nil, zero value otherwise.

### GetSecHeaderOk

`func (o *ClientSetupSection) GetSecHeaderOk() (*string, bool)`

GetSecHeaderOk returns a tuple with the SecHeader field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSecHeader

`func (o *ClientSetupSection) SetSecHeader(v string)`

SetSecHeader sets SecHeader field to given value.

### HasSecHeader

`func (o *ClientSetupSection) HasSecHeader() bool`

HasSecHeader returns a boolean if a field has been set.

### GetType

`func (o *ClientSetupSection) GetType() SectionType`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *ClientSetupSection) GetTypeOk() (*SectionType, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *ClientSetupSection) SetType(v SectionType)`

SetType sets Type field to given value.


### GetSteps

`func (o *ClientSetupSection) GetSteps() []ClientSetupStep`

GetSteps returns the Steps field if non-nil, zero value otherwise.

### GetStepsOk

`func (o *ClientSetupSection) GetStepsOk() (*[]ClientSetupStep, bool)`

GetStepsOk returns a tuple with the Steps field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSteps

`func (o *ClientSetupSection) SetSteps(v []ClientSetupStep)`

SetSteps sets Steps field to given value.

### HasSteps

`func (o *ClientSetupSection) HasSteps() bool`

HasSteps returns a boolean if a field has been set.

### GetTabs

`func (o *ClientSetupSection) GetTabs() []TabSetupStep`

GetTabs returns the Tabs field if non-nil, zero value otherwise.

### GetTabsOk

`func (o *ClientSetupSection) GetTabsOk() (*[]TabSetupStep, bool)`

GetTabsOk returns a tuple with the Tabs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTabs

`func (o *ClientSetupSection) SetTabs(v []TabSetupStep)`

SetTabs sets Tabs field to given value.

### HasTabs

`func (o *ClientSetupSection) HasTabs() bool`

HasTabs returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


