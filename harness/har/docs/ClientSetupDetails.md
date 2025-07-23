# ClientSetupDetails

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**MainHeader** | **string** |  | 
**SecHeader** | **string** |  | 
**Sections** | [**[]ClientSetupSection**](ClientSetupSection.md) |  | 

## Methods

### NewClientSetupDetails

`func NewClientSetupDetails(mainHeader string, secHeader string, sections []ClientSetupSection, ) *ClientSetupDetails`

NewClientSetupDetails instantiates a new ClientSetupDetails object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewClientSetupDetailsWithDefaults

`func NewClientSetupDetailsWithDefaults() *ClientSetupDetails`

NewClientSetupDetailsWithDefaults instantiates a new ClientSetupDetails object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetMainHeader

`func (o *ClientSetupDetails) GetMainHeader() string`

GetMainHeader returns the MainHeader field if non-nil, zero value otherwise.

### GetMainHeaderOk

`func (o *ClientSetupDetails) GetMainHeaderOk() (*string, bool)`

GetMainHeaderOk returns a tuple with the MainHeader field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMainHeader

`func (o *ClientSetupDetails) SetMainHeader(v string)`

SetMainHeader sets MainHeader field to given value.


### GetSecHeader

`func (o *ClientSetupDetails) GetSecHeader() string`

GetSecHeader returns the SecHeader field if non-nil, zero value otherwise.

### GetSecHeaderOk

`func (o *ClientSetupDetails) GetSecHeaderOk() (*string, bool)`

GetSecHeaderOk returns a tuple with the SecHeader field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSecHeader

`func (o *ClientSetupDetails) SetSecHeader(v string)`

SetSecHeader sets SecHeader field to given value.


### GetSections

`func (o *ClientSetupDetails) GetSections() []ClientSetupSection`

GetSections returns the Sections field if non-nil, zero value otherwise.

### GetSectionsOk

`func (o *ClientSetupDetails) GetSectionsOk() (*[]ClientSetupSection, bool)`

GetSectionsOk returns a tuple with the Sections field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSections

`func (o *ClientSetupDetails) SetSections(v []ClientSetupSection)`

SetSections sets Sections field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


