# DockerLayersSummary

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Digest** | **string** |  | 
**Layers** | Pointer to [**[]DockerLayerEntry**](DockerLayerEntry.md) |  | [optional] 
**OsArch** | Pointer to **string** |  | [optional] 

## Methods

### NewDockerLayersSummary

`func NewDockerLayersSummary(digest string, ) *DockerLayersSummary`

NewDockerLayersSummary instantiates a new DockerLayersSummary object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDockerLayersSummaryWithDefaults

`func NewDockerLayersSummaryWithDefaults() *DockerLayersSummary`

NewDockerLayersSummaryWithDefaults instantiates a new DockerLayersSummary object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDigest

`func (o *DockerLayersSummary) GetDigest() string`

GetDigest returns the Digest field if non-nil, zero value otherwise.

### GetDigestOk

`func (o *DockerLayersSummary) GetDigestOk() (*string, bool)`

GetDigestOk returns a tuple with the Digest field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDigest

`func (o *DockerLayersSummary) SetDigest(v string)`

SetDigest sets Digest field to given value.


### GetLayers

`func (o *DockerLayersSummary) GetLayers() []DockerLayerEntry`

GetLayers returns the Layers field if non-nil, zero value otherwise.

### GetLayersOk

`func (o *DockerLayersSummary) GetLayersOk() (*[]DockerLayerEntry, bool)`

GetLayersOk returns a tuple with the Layers field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLayers

`func (o *DockerLayersSummary) SetLayers(v []DockerLayerEntry)`

SetLayers sets Layers field to given value.

### HasLayers

`func (o *DockerLayersSummary) HasLayers() bool`

HasLayers returns a boolean if a field has been set.

### GetOsArch

`func (o *DockerLayersSummary) GetOsArch() string`

GetOsArch returns the OsArch field if non-nil, zero value otherwise.

### GetOsArchOk

`func (o *DockerLayersSummary) GetOsArchOk() (*string, bool)`

GetOsArchOk returns a tuple with the OsArch field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOsArch

`func (o *DockerLayersSummary) SetOsArch(v string)`

SetOsArch sets OsArch field to given value.

### HasOsArch

`func (o *DockerLayersSummary) HasOsArch() bool`

HasOsArch returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


