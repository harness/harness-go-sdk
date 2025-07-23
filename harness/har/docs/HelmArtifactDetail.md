# HelmArtifactDetail

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Artifact** | Pointer to **string** |  | [optional] 
**CreatedAt** | Pointer to **string** |  | [optional] 
**DownloadsCount** | Pointer to **int64** |  | [optional] 
**ModifiedAt** | Pointer to **string** |  | [optional] 
**PackageType** | [**PackageType**](PackageType.md) |  | 
**PullCommand** | Pointer to **string** |  | [optional] 
**RegistryPath** | **string** |  | 
**Size** | Pointer to **string** |  | [optional] 
**Url** | **string** |  | 
**Version** | **string** |  | 

## Methods

### NewHelmArtifactDetail

`func NewHelmArtifactDetail(packageType PackageType, registryPath string, url string, version string, ) *HelmArtifactDetail`

NewHelmArtifactDetail instantiates a new HelmArtifactDetail object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewHelmArtifactDetailWithDefaults

`func NewHelmArtifactDetailWithDefaults() *HelmArtifactDetail`

NewHelmArtifactDetailWithDefaults instantiates a new HelmArtifactDetail object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetArtifact

`func (o *HelmArtifactDetail) GetArtifact() string`

GetArtifact returns the Artifact field if non-nil, zero value otherwise.

### GetArtifactOk

`func (o *HelmArtifactDetail) GetArtifactOk() (*string, bool)`

GetArtifactOk returns a tuple with the Artifact field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetArtifact

`func (o *HelmArtifactDetail) SetArtifact(v string)`

SetArtifact sets Artifact field to given value.

### HasArtifact

`func (o *HelmArtifactDetail) HasArtifact() bool`

HasArtifact returns a boolean if a field has been set.

### GetCreatedAt

`func (o *HelmArtifactDetail) GetCreatedAt() string`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *HelmArtifactDetail) GetCreatedAtOk() (*string, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *HelmArtifactDetail) SetCreatedAt(v string)`

SetCreatedAt sets CreatedAt field to given value.

### HasCreatedAt

`func (o *HelmArtifactDetail) HasCreatedAt() bool`

HasCreatedAt returns a boolean if a field has been set.

### GetDownloadsCount

`func (o *HelmArtifactDetail) GetDownloadsCount() int64`

GetDownloadsCount returns the DownloadsCount field if non-nil, zero value otherwise.

### GetDownloadsCountOk

`func (o *HelmArtifactDetail) GetDownloadsCountOk() (*int64, bool)`

GetDownloadsCountOk returns a tuple with the DownloadsCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDownloadsCount

`func (o *HelmArtifactDetail) SetDownloadsCount(v int64)`

SetDownloadsCount sets DownloadsCount field to given value.

### HasDownloadsCount

`func (o *HelmArtifactDetail) HasDownloadsCount() bool`

HasDownloadsCount returns a boolean if a field has been set.

### GetModifiedAt

`func (o *HelmArtifactDetail) GetModifiedAt() string`

GetModifiedAt returns the ModifiedAt field if non-nil, zero value otherwise.

### GetModifiedAtOk

`func (o *HelmArtifactDetail) GetModifiedAtOk() (*string, bool)`

GetModifiedAtOk returns a tuple with the ModifiedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetModifiedAt

`func (o *HelmArtifactDetail) SetModifiedAt(v string)`

SetModifiedAt sets ModifiedAt field to given value.

### HasModifiedAt

`func (o *HelmArtifactDetail) HasModifiedAt() bool`

HasModifiedAt returns a boolean if a field has been set.

### GetPackageType

`func (o *HelmArtifactDetail) GetPackageType() PackageType`

GetPackageType returns the PackageType field if non-nil, zero value otherwise.

### GetPackageTypeOk

`func (o *HelmArtifactDetail) GetPackageTypeOk() (*PackageType, bool)`

GetPackageTypeOk returns a tuple with the PackageType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPackageType

`func (o *HelmArtifactDetail) SetPackageType(v PackageType)`

SetPackageType sets PackageType field to given value.


### GetPullCommand

`func (o *HelmArtifactDetail) GetPullCommand() string`

GetPullCommand returns the PullCommand field if non-nil, zero value otherwise.

### GetPullCommandOk

`func (o *HelmArtifactDetail) GetPullCommandOk() (*string, bool)`

GetPullCommandOk returns a tuple with the PullCommand field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPullCommand

`func (o *HelmArtifactDetail) SetPullCommand(v string)`

SetPullCommand sets PullCommand field to given value.

### HasPullCommand

`func (o *HelmArtifactDetail) HasPullCommand() bool`

HasPullCommand returns a boolean if a field has been set.

### GetRegistryPath

`func (o *HelmArtifactDetail) GetRegistryPath() string`

GetRegistryPath returns the RegistryPath field if non-nil, zero value otherwise.

### GetRegistryPathOk

`func (o *HelmArtifactDetail) GetRegistryPathOk() (*string, bool)`

GetRegistryPathOk returns a tuple with the RegistryPath field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRegistryPath

`func (o *HelmArtifactDetail) SetRegistryPath(v string)`

SetRegistryPath sets RegistryPath field to given value.


### GetSize

`func (o *HelmArtifactDetail) GetSize() string`

GetSize returns the Size field if non-nil, zero value otherwise.

### GetSizeOk

`func (o *HelmArtifactDetail) GetSizeOk() (*string, bool)`

GetSizeOk returns a tuple with the Size field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSize

`func (o *HelmArtifactDetail) SetSize(v string)`

SetSize sets Size field to given value.

### HasSize

`func (o *HelmArtifactDetail) HasSize() bool`

HasSize returns a boolean if a field has been set.

### GetUrl

`func (o *HelmArtifactDetail) GetUrl() string`

GetUrl returns the Url field if non-nil, zero value otherwise.

### GetUrlOk

`func (o *HelmArtifactDetail) GetUrlOk() (*string, bool)`

GetUrlOk returns a tuple with the Url field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUrl

`func (o *HelmArtifactDetail) SetUrl(v string)`

SetUrl sets Url field to given value.


### GetVersion

`func (o *HelmArtifactDetail) GetVersion() string`

GetVersion returns the Version field if non-nil, zero value otherwise.

### GetVersionOk

`func (o *HelmArtifactDetail) GetVersionOk() (*string, bool)`

GetVersionOk returns a tuple with the Version field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVersion

`func (o *HelmArtifactDetail) SetVersion(v string)`

SetVersion sets Version field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


