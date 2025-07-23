# DockerLayerEntry

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Command** | **string** |  | 
**Size** | Pointer to **string** |  | [optional] 

## Methods

### NewDockerLayerEntry

`func NewDockerLayerEntry(command string, ) *DockerLayerEntry`

NewDockerLayerEntry instantiates a new DockerLayerEntry object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDockerLayerEntryWithDefaults

`func NewDockerLayerEntryWithDefaults() *DockerLayerEntry`

NewDockerLayerEntryWithDefaults instantiates a new DockerLayerEntry object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCommand

`func (o *DockerLayerEntry) GetCommand() string`

GetCommand returns the Command field if non-nil, zero value otherwise.

### GetCommandOk

`func (o *DockerLayerEntry) GetCommandOk() (*string, bool)`

GetCommandOk returns a tuple with the Command field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCommand

`func (o *DockerLayerEntry) SetCommand(v string)`

SetCommand sets Command field to given value.


### GetSize

`func (o *DockerLayerEntry) GetSize() string`

GetSize returns the Size field if non-nil, zero value otherwise.

### GetSizeOk

`func (o *DockerLayerEntry) GetSizeOk() (*string, bool)`

GetSizeOk returns a tuple with the Size field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSize

`func (o *DockerLayerEntry) SetSize(v string)`

SetSize sets Size field to given value.

### HasSize

`func (o *DockerLayerEntry) HasSize() bool`

HasSize returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


