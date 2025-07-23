# ClientSetupStep

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Commands** | Pointer to [**[]ClientSetupStepCommand**](ClientSetupStepCommand.md) |  | [optional] 
**Header** | Pointer to **string** |  | [optional] 
**Type** | Pointer to [**ClientSetupStepType**](ClientSetupStepType.md) |  | [optional] 

## Methods

### NewClientSetupStep

`func NewClientSetupStep() *ClientSetupStep`

NewClientSetupStep instantiates a new ClientSetupStep object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewClientSetupStepWithDefaults

`func NewClientSetupStepWithDefaults() *ClientSetupStep`

NewClientSetupStepWithDefaults instantiates a new ClientSetupStep object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCommands

`func (o *ClientSetupStep) GetCommands() []ClientSetupStepCommand`

GetCommands returns the Commands field if non-nil, zero value otherwise.

### GetCommandsOk

`func (o *ClientSetupStep) GetCommandsOk() (*[]ClientSetupStepCommand, bool)`

GetCommandsOk returns a tuple with the Commands field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCommands

`func (o *ClientSetupStep) SetCommands(v []ClientSetupStepCommand)`

SetCommands sets Commands field to given value.

### HasCommands

`func (o *ClientSetupStep) HasCommands() bool`

HasCommands returns a boolean if a field has been set.

### GetHeader

`func (o *ClientSetupStep) GetHeader() string`

GetHeader returns the Header field if non-nil, zero value otherwise.

### GetHeaderOk

`func (o *ClientSetupStep) GetHeaderOk() (*string, bool)`

GetHeaderOk returns a tuple with the Header field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHeader

`func (o *ClientSetupStep) SetHeader(v string)`

SetHeader sets Header field to given value.

### HasHeader

`func (o *ClientSetupStep) HasHeader() bool`

HasHeader returns a boolean if a field has been set.

### GetType

`func (o *ClientSetupStep) GetType() ClientSetupStepType`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *ClientSetupStep) GetTypeOk() (*ClientSetupStepType, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *ClientSetupStep) SetType(v ClientSetupStepType)`

SetType sets Type field to given value.

### HasType

`func (o *ClientSetupStep) HasType() bool`

HasType returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


