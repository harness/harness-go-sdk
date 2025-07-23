# ArtifactVersionMetadata

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**DeploymentMetadata** | Pointer to [**DeploymentMetadata**](DeploymentMetadata.md) |  | [optional] 
**DigestCount** | Pointer to **int64** |  | [optional] 
**DownloadsCount** | Pointer to **int64** |  | [optional] 
**FileCount** | Pointer to **int64** |  | [optional] 
**LastModified** | Pointer to **string** |  | [optional] 
**Name** | **string** |  | 
**PackageType** | Pointer to [**PackageType**](PackageType.md) |  | [optional] 
**PullCommand** | Pointer to **string** |  | [optional] 
**RegistryIdentifier** | **string** |  | 
**RegistryPath** | **string** |  | 
**Size** | Pointer to **string** |  | [optional] 

## Methods

### NewArtifactVersionMetadata

`func NewArtifactVersionMetadata(name string, registryIdentifier string, registryPath string, ) *ArtifactVersionMetadata`

NewArtifactVersionMetadata instantiates a new ArtifactVersionMetadata object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewArtifactVersionMetadataWithDefaults

`func NewArtifactVersionMetadataWithDefaults() *ArtifactVersionMetadata`

NewArtifactVersionMetadataWithDefaults instantiates a new ArtifactVersionMetadata object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDeploymentMetadata

`func (o *ArtifactVersionMetadata) GetDeploymentMetadata() DeploymentMetadata`

GetDeploymentMetadata returns the DeploymentMetadata field if non-nil, zero value otherwise.

### GetDeploymentMetadataOk

`func (o *ArtifactVersionMetadata) GetDeploymentMetadataOk() (*DeploymentMetadata, bool)`

GetDeploymentMetadataOk returns a tuple with the DeploymentMetadata field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeploymentMetadata

`func (o *ArtifactVersionMetadata) SetDeploymentMetadata(v DeploymentMetadata)`

SetDeploymentMetadata sets DeploymentMetadata field to given value.

### HasDeploymentMetadata

`func (o *ArtifactVersionMetadata) HasDeploymentMetadata() bool`

HasDeploymentMetadata returns a boolean if a field has been set.

### GetDigestCount

`func (o *ArtifactVersionMetadata) GetDigestCount() int64`

GetDigestCount returns the DigestCount field if non-nil, zero value otherwise.

### GetDigestCountOk

`func (o *ArtifactVersionMetadata) GetDigestCountOk() (*int64, bool)`

GetDigestCountOk returns a tuple with the DigestCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDigestCount

`func (o *ArtifactVersionMetadata) SetDigestCount(v int64)`

SetDigestCount sets DigestCount field to given value.

### HasDigestCount

`func (o *ArtifactVersionMetadata) HasDigestCount() bool`

HasDigestCount returns a boolean if a field has been set.

### GetDownloadsCount

`func (o *ArtifactVersionMetadata) GetDownloadsCount() int64`

GetDownloadsCount returns the DownloadsCount field if non-nil, zero value otherwise.

### GetDownloadsCountOk

`func (o *ArtifactVersionMetadata) GetDownloadsCountOk() (*int64, bool)`

GetDownloadsCountOk returns a tuple with the DownloadsCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDownloadsCount

`func (o *ArtifactVersionMetadata) SetDownloadsCount(v int64)`

SetDownloadsCount sets DownloadsCount field to given value.

### HasDownloadsCount

`func (o *ArtifactVersionMetadata) HasDownloadsCount() bool`

HasDownloadsCount returns a boolean if a field has been set.

### GetFileCount

`func (o *ArtifactVersionMetadata) GetFileCount() int64`

GetFileCount returns the FileCount field if non-nil, zero value otherwise.

### GetFileCountOk

`func (o *ArtifactVersionMetadata) GetFileCountOk() (*int64, bool)`

GetFileCountOk returns a tuple with the FileCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFileCount

`func (o *ArtifactVersionMetadata) SetFileCount(v int64)`

SetFileCount sets FileCount field to given value.

### HasFileCount

`func (o *ArtifactVersionMetadata) HasFileCount() bool`

HasFileCount returns a boolean if a field has been set.

### GetLastModified

`func (o *ArtifactVersionMetadata) GetLastModified() string`

GetLastModified returns the LastModified field if non-nil, zero value otherwise.

### GetLastModifiedOk

`func (o *ArtifactVersionMetadata) GetLastModifiedOk() (*string, bool)`

GetLastModifiedOk returns a tuple with the LastModified field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastModified

`func (o *ArtifactVersionMetadata) SetLastModified(v string)`

SetLastModified sets LastModified field to given value.

### HasLastModified

`func (o *ArtifactVersionMetadata) HasLastModified() bool`

HasLastModified returns a boolean if a field has been set.

### GetName

`func (o *ArtifactVersionMetadata) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *ArtifactVersionMetadata) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *ArtifactVersionMetadata) SetName(v string)`

SetName sets Name field to given value.


### GetPackageType

`func (o *ArtifactVersionMetadata) GetPackageType() PackageType`

GetPackageType returns the PackageType field if non-nil, zero value otherwise.

### GetPackageTypeOk

`func (o *ArtifactVersionMetadata) GetPackageTypeOk() (*PackageType, bool)`

GetPackageTypeOk returns a tuple with the PackageType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPackageType

`func (o *ArtifactVersionMetadata) SetPackageType(v PackageType)`

SetPackageType sets PackageType field to given value.

### HasPackageType

`func (o *ArtifactVersionMetadata) HasPackageType() bool`

HasPackageType returns a boolean if a field has been set.

### GetPullCommand

`func (o *ArtifactVersionMetadata) GetPullCommand() string`

GetPullCommand returns the PullCommand field if non-nil, zero value otherwise.

### GetPullCommandOk

`func (o *ArtifactVersionMetadata) GetPullCommandOk() (*string, bool)`

GetPullCommandOk returns a tuple with the PullCommand field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPullCommand

`func (o *ArtifactVersionMetadata) SetPullCommand(v string)`

SetPullCommand sets PullCommand field to given value.

### HasPullCommand

`func (o *ArtifactVersionMetadata) HasPullCommand() bool`

HasPullCommand returns a boolean if a field has been set.

### GetRegistryIdentifier

`func (o *ArtifactVersionMetadata) GetRegistryIdentifier() string`

GetRegistryIdentifier returns the RegistryIdentifier field if non-nil, zero value otherwise.

### GetRegistryIdentifierOk

`func (o *ArtifactVersionMetadata) GetRegistryIdentifierOk() (*string, bool)`

GetRegistryIdentifierOk returns a tuple with the RegistryIdentifier field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRegistryIdentifier

`func (o *ArtifactVersionMetadata) SetRegistryIdentifier(v string)`

SetRegistryIdentifier sets RegistryIdentifier field to given value.


### GetRegistryPath

`func (o *ArtifactVersionMetadata) GetRegistryPath() string`

GetRegistryPath returns the RegistryPath field if non-nil, zero value otherwise.

### GetRegistryPathOk

`func (o *ArtifactVersionMetadata) GetRegistryPathOk() (*string, bool)`

GetRegistryPathOk returns a tuple with the RegistryPath field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRegistryPath

`func (o *ArtifactVersionMetadata) SetRegistryPath(v string)`

SetRegistryPath sets RegistryPath field to given value.


### GetSize

`func (o *ArtifactVersionMetadata) GetSize() string`

GetSize returns the Size field if non-nil, zero value otherwise.

### GetSizeOk

`func (o *ArtifactVersionMetadata) GetSizeOk() (*string, bool)`

GetSizeOk returns a tuple with the Size field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSize

`func (o *ArtifactVersionMetadata) SetSize(v string)`

SetSize sets Size field to given value.

### HasSize

`func (o *ArtifactVersionMetadata) HasSize() bool`

HasSize returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


