# DockerManifests

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ImageName** | **string** |  | 
**Manifests** | Pointer to [**[]DockerManifestDetails**](DockerManifestDetails.md) |  | [optional] 
**Version** | **string** |  | 

## Methods

### NewDockerManifests

`func NewDockerManifests(imageName string, version string, ) *DockerManifests`

NewDockerManifests instantiates a new DockerManifests object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDockerManifestsWithDefaults

`func NewDockerManifestsWithDefaults() *DockerManifests`

NewDockerManifestsWithDefaults instantiates a new DockerManifests object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetImageName

`func (o *DockerManifests) GetImageName() string`

GetImageName returns the ImageName field if non-nil, zero value otherwise.

### GetImageNameOk

`func (o *DockerManifests) GetImageNameOk() (*string, bool)`

GetImageNameOk returns a tuple with the ImageName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetImageName

`func (o *DockerManifests) SetImageName(v string)`

SetImageName sets ImageName field to given value.


### GetManifests

`func (o *DockerManifests) GetManifests() []DockerManifestDetails`

GetManifests returns the Manifests field if non-nil, zero value otherwise.

### GetManifestsOk

`func (o *DockerManifests) GetManifestsOk() (*[]DockerManifestDetails, bool)`

GetManifestsOk returns a tuple with the Manifests field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetManifests

`func (o *DockerManifests) SetManifests(v []DockerManifestDetails)`

SetManifests sets Manifests field to given value.

### HasManifests

`func (o *DockerManifests) HasManifests() bool`

HasManifests returns a boolean if a field has been set.

### GetVersion

`func (o *DockerManifests) GetVersion() string`

GetVersion returns the Version field if non-nil, zero value otherwise.

### GetVersionOk

`func (o *DockerManifests) GetVersionOk() (*string, bool)`

GetVersionOk returns a tuple with the Version field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVersion

`func (o *DockerManifests) SetVersion(v string)`

SetVersion sets Version field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


