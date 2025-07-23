# ClientSetupStepConfig

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Steps** | Pointer to [**[]ClientSetupStep**](ClientSetupStep.md) |  | [optional] 

## Methods

### NewClientSetupStepConfig

`func NewClientSetupStepConfig() *ClientSetupStepConfig`

NewClientSetupStepConfig instantiates a new ClientSetupStepConfig object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewClientSetupStepConfigWithDefaults

`func NewClientSetupStepConfigWithDefaults() *ClientSetupStepConfig`

NewClientSetupStepConfigWithDefaults instantiates a new ClientSetupStepConfig object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSteps

`func (o *ClientSetupStepConfig) GetSteps() []ClientSetupStep`

GetSteps returns the Steps field if non-nil, zero value otherwise.

### GetStepsOk

`func (o *ClientSetupStepConfig) GetStepsOk() (*[]ClientSetupStep, bool)`

GetStepsOk returns a tuple with the Steps field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSteps

`func (o *ClientSetupStepConfig) SetSteps(v []ClientSetupStep)`

SetSteps sets Steps field to given value.

### HasSteps

`func (o *ClientSetupStepConfig) HasSteps() bool`

HasSteps returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


