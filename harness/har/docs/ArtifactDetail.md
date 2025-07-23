# ArtifactDetail

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CreatedAt** | Pointer to **string** |  | [optional] 
**CreatedBy** | Pointer to **string** |  | [optional] 
**DownloadCount** | Pointer to **int64** |  | [optional] 
**ModifiedAt** | Pointer to **string** |  | [optional] 
**Name** | Pointer to **string** |  | [optional] 
**PackageType** | [**PackageType**](PackageType.md) |  | 
**Size** | Pointer to **string** |  | [optional] 
**Version** | **string** |  | 
**PullCommand** | Pointer to **string** |  | [optional] 
**Description** | Pointer to **string** |  | [optional] 
**ArtifactId** | Pointer to **string** |  | [optional] 
**GroupId** | Pointer to **string** |  | [optional] 
**Metadata** | Pointer to **map[string]interface{}** |  | [optional] 

## Methods

### NewArtifactDetail

`func NewArtifactDetail(packageType PackageType, version string, ) *ArtifactDetail`

NewArtifactDetail instantiates a new ArtifactDetail object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewArtifactDetailWithDefaults

`func NewArtifactDetailWithDefaults() *ArtifactDetail`

NewArtifactDetailWithDefaults instantiates a new ArtifactDetail object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCreatedAt

`func (o *ArtifactDetail) GetCreatedAt() string`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *ArtifactDetail) GetCreatedAtOk() (*string, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *ArtifactDetail) SetCreatedAt(v string)`

SetCreatedAt sets CreatedAt field to given value.

### HasCreatedAt

`func (o *ArtifactDetail) HasCreatedAt() bool`

HasCreatedAt returns a boolean if a field has been set.

### GetCreatedBy

`func (o *ArtifactDetail) GetCreatedBy() string`

GetCreatedBy returns the CreatedBy field if non-nil, zero value otherwise.

### GetCreatedByOk

`func (o *ArtifactDetail) GetCreatedByOk() (*string, bool)`

GetCreatedByOk returns a tuple with the CreatedBy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedBy

`func (o *ArtifactDetail) SetCreatedBy(v string)`

SetCreatedBy sets CreatedBy field to given value.

### HasCreatedBy

`func (o *ArtifactDetail) HasCreatedBy() bool`

HasCreatedBy returns a boolean if a field has been set.

### GetDownloadCount

`func (o *ArtifactDetail) GetDownloadCount() int64`

GetDownloadCount returns the DownloadCount field if non-nil, zero value otherwise.

### GetDownloadCountOk

`func (o *ArtifactDetail) GetDownloadCountOk() (*int64, bool)`

GetDownloadCountOk returns a tuple with the DownloadCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDownloadCount

`func (o *ArtifactDetail) SetDownloadCount(v int64)`

SetDownloadCount sets DownloadCount field to given value.

### HasDownloadCount

`func (o *ArtifactDetail) HasDownloadCount() bool`

HasDownloadCount returns a boolean if a field has been set.

### GetModifiedAt

`func (o *ArtifactDetail) GetModifiedAt() string`

GetModifiedAt returns the ModifiedAt field if non-nil, zero value otherwise.

### GetModifiedAtOk

`func (o *ArtifactDetail) GetModifiedAtOk() (*string, bool)`

GetModifiedAtOk returns a tuple with the ModifiedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetModifiedAt

`func (o *ArtifactDetail) SetModifiedAt(v string)`

SetModifiedAt sets ModifiedAt field to given value.

### HasModifiedAt

`func (o *ArtifactDetail) HasModifiedAt() bool`

HasModifiedAt returns a boolean if a field has been set.

### GetName

`func (o *ArtifactDetail) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *ArtifactDetail) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *ArtifactDetail) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *ArtifactDetail) HasName() bool`

HasName returns a boolean if a field has been set.

### GetPackageType

`func (o *ArtifactDetail) GetPackageType() PackageType`

GetPackageType returns the PackageType field if non-nil, zero value otherwise.

### GetPackageTypeOk

`func (o *ArtifactDetail) GetPackageTypeOk() (*PackageType, bool)`

GetPackageTypeOk returns a tuple with the PackageType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPackageType

`func (o *ArtifactDetail) SetPackageType(v PackageType)`

SetPackageType sets PackageType field to given value.


### GetSize

`func (o *ArtifactDetail) GetSize() string`

GetSize returns the Size field if non-nil, zero value otherwise.

### GetSizeOk

`func (o *ArtifactDetail) GetSizeOk() (*string, bool)`

GetSizeOk returns a tuple with the Size field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSize

`func (o *ArtifactDetail) SetSize(v string)`

SetSize sets Size field to given value.

### HasSize

`func (o *ArtifactDetail) HasSize() bool`

HasSize returns a boolean if a field has been set.

### GetVersion

`func (o *ArtifactDetail) GetVersion() string`

GetVersion returns the Version field if non-nil, zero value otherwise.

### GetVersionOk

`func (o *ArtifactDetail) GetVersionOk() (*string, bool)`

GetVersionOk returns a tuple with the Version field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVersion

`func (o *ArtifactDetail) SetVersion(v string)`

SetVersion sets Version field to given value.


### GetPullCommand

`func (o *ArtifactDetail) GetPullCommand() string`

GetPullCommand returns the PullCommand field if non-nil, zero value otherwise.

### GetPullCommandOk

`func (o *ArtifactDetail) GetPullCommandOk() (*string, bool)`

GetPullCommandOk returns a tuple with the PullCommand field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPullCommand

`func (o *ArtifactDetail) SetPullCommand(v string)`

SetPullCommand sets PullCommand field to given value.

### HasPullCommand

`func (o *ArtifactDetail) HasPullCommand() bool`

HasPullCommand returns a boolean if a field has been set.

### GetDescription

`func (o *ArtifactDetail) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *ArtifactDetail) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *ArtifactDetail) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *ArtifactDetail) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetArtifactId

`func (o *ArtifactDetail) GetArtifactId() string`

GetArtifactId returns the ArtifactId field if non-nil, zero value otherwise.

### GetArtifactIdOk

`func (o *ArtifactDetail) GetArtifactIdOk() (*string, bool)`

GetArtifactIdOk returns a tuple with the ArtifactId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetArtifactId

`func (o *ArtifactDetail) SetArtifactId(v string)`

SetArtifactId sets ArtifactId field to given value.

### HasArtifactId

`func (o *ArtifactDetail) HasArtifactId() bool`

HasArtifactId returns a boolean if a field has been set.

### GetGroupId

`func (o *ArtifactDetail) GetGroupId() string`

GetGroupId returns the GroupId field if non-nil, zero value otherwise.

### GetGroupIdOk

`func (o *ArtifactDetail) GetGroupIdOk() (*string, bool)`

GetGroupIdOk returns a tuple with the GroupId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGroupId

`func (o *ArtifactDetail) SetGroupId(v string)`

SetGroupId sets GroupId field to given value.

### HasGroupId

`func (o *ArtifactDetail) HasGroupId() bool`

HasGroupId returns a boolean if a field has been set.

### GetMetadata

`func (o *ArtifactDetail) GetMetadata() map[string]interface{}`

GetMetadata returns the Metadata field if non-nil, zero value otherwise.

### GetMetadataOk

`func (o *ArtifactDetail) GetMetadataOk() (*map[string]interface{}, bool)`

GetMetadataOk returns a tuple with the Metadata field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetadata

`func (o *ArtifactDetail) SetMetadata(v map[string]interface{})`

SetMetadata sets Metadata field to given value.

### HasMetadata

`func (o *ArtifactDetail) HasMetadata() bool`

HasMetadata returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


