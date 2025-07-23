# RegistryArtifactMetadata

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**DownloadsCount** | Pointer to **int64** |  | [optional] 
**Labels** | Pointer to **[]string** |  | [optional] 
**LastModified** | Pointer to **string** |  | [optional] 
**LatestVersion** | **string** |  | 
**Name** | **string** |  | 
**PackageType** | Pointer to [**PackageType**](PackageType.md) |  | [optional] 
**RegistryIdentifier** | **string** |  | 
**RegistryPath** | **string** |  | 

## Methods

### NewRegistryArtifactMetadata

`func NewRegistryArtifactMetadata(latestVersion string, name string, registryIdentifier string, registryPath string, ) *RegistryArtifactMetadata`

NewRegistryArtifactMetadata instantiates a new RegistryArtifactMetadata object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRegistryArtifactMetadataWithDefaults

`func NewRegistryArtifactMetadataWithDefaults() *RegistryArtifactMetadata`

NewRegistryArtifactMetadataWithDefaults instantiates a new RegistryArtifactMetadata object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDownloadsCount

`func (o *RegistryArtifactMetadata) GetDownloadsCount() int64`

GetDownloadsCount returns the DownloadsCount field if non-nil, zero value otherwise.

### GetDownloadsCountOk

`func (o *RegistryArtifactMetadata) GetDownloadsCountOk() (*int64, bool)`

GetDownloadsCountOk returns a tuple with the DownloadsCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDownloadsCount

`func (o *RegistryArtifactMetadata) SetDownloadsCount(v int64)`

SetDownloadsCount sets DownloadsCount field to given value.

### HasDownloadsCount

`func (o *RegistryArtifactMetadata) HasDownloadsCount() bool`

HasDownloadsCount returns a boolean if a field has been set.

### GetLabels

`func (o *RegistryArtifactMetadata) GetLabels() []string`

GetLabels returns the Labels field if non-nil, zero value otherwise.

### GetLabelsOk

`func (o *RegistryArtifactMetadata) GetLabelsOk() (*[]string, bool)`

GetLabelsOk returns a tuple with the Labels field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLabels

`func (o *RegistryArtifactMetadata) SetLabels(v []string)`

SetLabels sets Labels field to given value.

### HasLabels

`func (o *RegistryArtifactMetadata) HasLabels() bool`

HasLabels returns a boolean if a field has been set.

### GetLastModified

`func (o *RegistryArtifactMetadata) GetLastModified() string`

GetLastModified returns the LastModified field if non-nil, zero value otherwise.

### GetLastModifiedOk

`func (o *RegistryArtifactMetadata) GetLastModifiedOk() (*string, bool)`

GetLastModifiedOk returns a tuple with the LastModified field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastModified

`func (o *RegistryArtifactMetadata) SetLastModified(v string)`

SetLastModified sets LastModified field to given value.

### HasLastModified

`func (o *RegistryArtifactMetadata) HasLastModified() bool`

HasLastModified returns a boolean if a field has been set.

### GetLatestVersion

`func (o *RegistryArtifactMetadata) GetLatestVersion() string`

GetLatestVersion returns the LatestVersion field if non-nil, zero value otherwise.

### GetLatestVersionOk

`func (o *RegistryArtifactMetadata) GetLatestVersionOk() (*string, bool)`

GetLatestVersionOk returns a tuple with the LatestVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLatestVersion

`func (o *RegistryArtifactMetadata) SetLatestVersion(v string)`

SetLatestVersion sets LatestVersion field to given value.


### GetName

`func (o *RegistryArtifactMetadata) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *RegistryArtifactMetadata) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *RegistryArtifactMetadata) SetName(v string)`

SetName sets Name field to given value.


### GetPackageType

`func (o *RegistryArtifactMetadata) GetPackageType() PackageType`

GetPackageType returns the PackageType field if non-nil, zero value otherwise.

### GetPackageTypeOk

`func (o *RegistryArtifactMetadata) GetPackageTypeOk() (*PackageType, bool)`

GetPackageTypeOk returns a tuple with the PackageType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPackageType

`func (o *RegistryArtifactMetadata) SetPackageType(v PackageType)`

SetPackageType sets PackageType field to given value.

### HasPackageType

`func (o *RegistryArtifactMetadata) HasPackageType() bool`

HasPackageType returns a boolean if a field has been set.

### GetRegistryIdentifier

`func (o *RegistryArtifactMetadata) GetRegistryIdentifier() string`

GetRegistryIdentifier returns the RegistryIdentifier field if non-nil, zero value otherwise.

### GetRegistryIdentifierOk

`func (o *RegistryArtifactMetadata) GetRegistryIdentifierOk() (*string, bool)`

GetRegistryIdentifierOk returns a tuple with the RegistryIdentifier field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRegistryIdentifier

`func (o *RegistryArtifactMetadata) SetRegistryIdentifier(v string)`

SetRegistryIdentifier sets RegistryIdentifier field to given value.


### GetRegistryPath

`func (o *RegistryArtifactMetadata) GetRegistryPath() string`

GetRegistryPath returns the RegistryPath field if non-nil, zero value otherwise.

### GetRegistryPathOk

`func (o *RegistryArtifactMetadata) GetRegistryPathOk() (*string, bool)`

GetRegistryPathOk returns a tuple with the RegistryPath field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRegistryPath

`func (o *RegistryArtifactMetadata) SetRegistryPath(v string)`

SetRegistryPath sets RegistryPath field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


