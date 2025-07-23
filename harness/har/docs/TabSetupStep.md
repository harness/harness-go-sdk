# TabSetupStep

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Header** | Pointer to **string** |  | [optional] 
**Sections** | Pointer to [**[]ClientSetupSection**](ClientSetupSection.md) |  | [optional] 

## Methods

### NewTabSetupStep

`func NewTabSetupStep() *TabSetupStep`

NewTabSetupStep instantiates a new TabSetupStep object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTabSetupStepWithDefaults

`func NewTabSetupStepWithDefaults() *TabSetupStep`

NewTabSetupStepWithDefaults instantiates a new TabSetupStep object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetHeader

`func (o *TabSetupStep) GetHeader() string`

GetHeader returns the Header field if non-nil, zero value otherwise.

### GetHeaderOk

`func (o *TabSetupStep) GetHeaderOk() (*string, bool)`

GetHeaderOk returns a tuple with the Header field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHeader

`func (o *TabSetupStep) SetHeader(v string)`

SetHeader sets Header field to given value.

### HasHeader

`func (o *TabSetupStep) HasHeader() bool`

HasHeader returns a boolean if a field has been set.

### GetSections

`func (o *TabSetupStep) GetSections() []ClientSetupSection`

GetSections returns the Sections field if non-nil, zero value otherwise.

### GetSectionsOk

`func (o *TabSetupStep) GetSectionsOk() (*[]ClientSetupSection, bool)`

GetSectionsOk returns a tuple with the Sections field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSections

`func (o *TabSetupStep) SetSections(v []ClientSetupSection)`

SetSections sets Sections field to given value.

### HasSections

`func (o *TabSetupStep) HasSections() bool`

HasSections returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


