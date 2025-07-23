# DockerArtifactDetail

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CreatedAt** | Pointer to **string** |  | [optional] 
**DownloadsCount** | Pointer to **int64** |  | [optional] 
**ImageName** | **string** |  | 
**ModifiedAt** | Pointer to **string** |  | [optional] 
**PackageType** | [**PackageType**](PackageType.md) |  | 
**PullCommand** | Pointer to **string** |  | [optional] 
**RegistryPath** | **string** |  | 
**Size** | Pointer to **string** |  | [optional] 
**Url** | **string** |  | 
**Version** | **string** |  | 

## Methods

### NewDockerArtifactDetail

`func NewDockerArtifactDetail(imageName string, packageType PackageType, registryPath string, url string, version string, ) *DockerArtifactDetail`

NewDockerArtifactDetail instantiates a new DockerArtifactDetail object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDockerArtifactDetailWithDefaults

`func NewDockerArtifactDetailWithDefaults() *DockerArtifactDetail`

NewDockerArtifactDetailWithDefaults instantiates a new DockerArtifactDetail object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCreatedAt

`func (o *DockerArtifactDetail) GetCreatedAt() string`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *DockerArtifactDetail) GetCreatedAtOk() (*string, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *DockerArtifactDetail) SetCreatedAt(v string)`

SetCreatedAt sets CreatedAt field to given value.

### HasCreatedAt

`func (o *DockerArtifactDetail) HasCreatedAt() bool`

HasCreatedAt returns a boolean if a field has been set.

### GetDownloadsCount

`func (o *DockerArtifactDetail) GetDownloadsCount() int64`

GetDownloadsCount returns the DownloadsCount field if non-nil, zero value otherwise.

### GetDownloadsCountOk

`func (o *DockerArtifactDetail) GetDownloadsCountOk() (*int64, bool)`

GetDownloadsCountOk returns a tuple with the DownloadsCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDownloadsCount

`func (o *DockerArtifactDetail) SetDownloadsCount(v int64)`

SetDownloadsCount sets DownloadsCount field to given value.

### HasDownloadsCount

`func (o *DockerArtifactDetail) HasDownloadsCount() bool`

HasDownloadsCount returns a boolean if a field has been set.

### GetImageName

`func (o *DockerArtifactDetail) GetImageName() string`

GetImageName returns the ImageName field if non-nil, zero value otherwise.

### GetImageNameOk

`func (o *DockerArtifactDetail) GetImageNameOk() (*string, bool)`

GetImageNameOk returns a tuple with the ImageName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetImageName

`func (o *DockerArtifactDetail) SetImageName(v string)`

SetImageName sets ImageName field to given value.


### GetModifiedAt

`func (o *DockerArtifactDetail) GetModifiedAt() string`

GetModifiedAt returns the ModifiedAt field if non-nil, zero value otherwise.

### GetModifiedAtOk

`func (o *DockerArtifactDetail) GetModifiedAtOk() (*string, bool)`

GetModifiedAtOk returns a tuple with the ModifiedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetModifiedAt

`func (o *DockerArtifactDetail) SetModifiedAt(v string)`

SetModifiedAt sets ModifiedAt field to given value.

### HasModifiedAt

`func (o *DockerArtifactDetail) HasModifiedAt() bool`

HasModifiedAt returns a boolean if a field has been set.

### GetPackageType

`func (o *DockerArtifactDetail) GetPackageType() PackageType`

GetPackageType returns the PackageType field if non-nil, zero value otherwise.

### GetPackageTypeOk

`func (o *DockerArtifactDetail) GetPackageTypeOk() (*PackageType, bool)`

GetPackageTypeOk returns a tuple with the PackageType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPackageType

`func (o *DockerArtifactDetail) SetPackageType(v PackageType)`

SetPackageType sets PackageType field to given value.


### GetPullCommand

`func (o *DockerArtifactDetail) GetPullCommand() string`

GetPullCommand returns the PullCommand field if non-nil, zero value otherwise.

### GetPullCommandOk

`func (o *DockerArtifactDetail) GetPullCommandOk() (*string, bool)`

GetPullCommandOk returns a tuple with the PullCommand field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPullCommand

`func (o *DockerArtifactDetail) SetPullCommand(v string)`

SetPullCommand sets PullCommand field to given value.

### HasPullCommand

`func (o *DockerArtifactDetail) HasPullCommand() bool`

HasPullCommand returns a boolean if a field has been set.

### GetRegistryPath

`func (o *DockerArtifactDetail) GetRegistryPath() string`

GetRegistryPath returns the RegistryPath field if non-nil, zero value otherwise.

### GetRegistryPathOk

`func (o *DockerArtifactDetail) GetRegistryPathOk() (*string, bool)`

GetRegistryPathOk returns a tuple with the RegistryPath field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRegistryPath

`func (o *DockerArtifactDetail) SetRegistryPath(v string)`

SetRegistryPath sets RegistryPath field to given value.


### GetSize

`func (o *DockerArtifactDetail) GetSize() string`

GetSize returns the Size field if non-nil, zero value otherwise.

### GetSizeOk

`func (o *DockerArtifactDetail) GetSizeOk() (*string, bool)`

GetSizeOk returns a tuple with the Size field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSize

`func (o *DockerArtifactDetail) SetSize(v string)`

SetSize sets Size field to given value.

### HasSize

`func (o *DockerArtifactDetail) HasSize() bool`

HasSize returns a boolean if a field has been set.

### GetUrl

`func (o *DockerArtifactDetail) GetUrl() string`

GetUrl returns the Url field if non-nil, zero value otherwise.

### GetUrlOk

`func (o *DockerArtifactDetail) GetUrlOk() (*string, bool)`

GetUrlOk returns a tuple with the Url field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUrl

`func (o *DockerArtifactDetail) SetUrl(v string)`

SetUrl sets Url field to given value.


### GetVersion

`func (o *DockerArtifactDetail) GetVersion() string`

GetVersion returns the Version field if non-nil, zero value otherwise.

### GetVersionOk

`func (o *DockerArtifactDetail) GetVersionOk() (*string, bool)`

GetVersionOk returns a tuple with the Version field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVersion

`func (o *DockerArtifactDetail) SetVersion(v string)`

SetVersion sets Version field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


