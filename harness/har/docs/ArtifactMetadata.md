# ArtifactMetadata

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**DeploymentMetadata** | Pointer to [**DeploymentMetadata**](DeploymentMetadata.md) |  | [optional] 
**DownloadsCount** | Pointer to **int64** |  | [optional] 
**Labels** | Pointer to **[]string** |  | [optional] 
**LastModified** | Pointer to **string** |  | [optional] 
**LatestVersion** | **string** |  | 
**Name** | **string** |  | 
**PackageType** | Pointer to [**PackageType**](PackageType.md) |  | [optional] 
**PullCommand** | Pointer to **string** |  | [optional] 
**RegistryIdentifier** | **string** |  | 
**RegistryPath** | **string** |  | 
**Version** | **string** |  | 

## Methods

### NewArtifactMetadata

`func NewArtifactMetadata(latestVersion string, name string, registryIdentifier string, registryPath string, version string, ) *ArtifactMetadata`

NewArtifactMetadata instantiates a new ArtifactMetadata object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewArtifactMetadataWithDefaults

`func NewArtifactMetadataWithDefaults() *ArtifactMetadata`

NewArtifactMetadataWithDefaults instantiates a new ArtifactMetadata object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDeploymentMetadata

`func (o *ArtifactMetadata) GetDeploymentMetadata() DeploymentMetadata`

GetDeploymentMetadata returns the DeploymentMetadata field if non-nil, zero value otherwise.

### GetDeploymentMetadataOk

`func (o *ArtifactMetadata) GetDeploymentMetadataOk() (*DeploymentMetadata, bool)`

GetDeploymentMetadataOk returns a tuple with the DeploymentMetadata field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeploymentMetadata

`func (o *ArtifactMetadata) SetDeploymentMetadata(v DeploymentMetadata)`

SetDeploymentMetadata sets DeploymentMetadata field to given value.

### HasDeploymentMetadata

`func (o *ArtifactMetadata) HasDeploymentMetadata() bool`

HasDeploymentMetadata returns a boolean if a field has been set.

### GetDownloadsCount

`func (o *ArtifactMetadata) GetDownloadsCount() int64`

GetDownloadsCount returns the DownloadsCount field if non-nil, zero value otherwise.

### GetDownloadsCountOk

`func (o *ArtifactMetadata) GetDownloadsCountOk() (*int64, bool)`

GetDownloadsCountOk returns a tuple with the DownloadsCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDownloadsCount

`func (o *ArtifactMetadata) SetDownloadsCount(v int64)`

SetDownloadsCount sets DownloadsCount field to given value.

### HasDownloadsCount

`func (o *ArtifactMetadata) HasDownloadsCount() bool`

HasDownloadsCount returns a boolean if a field has been set.

### GetLabels

`func (o *ArtifactMetadata) GetLabels() []string`

GetLabels returns the Labels field if non-nil, zero value otherwise.

### GetLabelsOk

`func (o *ArtifactMetadata) GetLabelsOk() (*[]string, bool)`

GetLabelsOk returns a tuple with the Labels field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLabels

`func (o *ArtifactMetadata) SetLabels(v []string)`

SetLabels sets Labels field to given value.

### HasLabels

`func (o *ArtifactMetadata) HasLabels() bool`

HasLabels returns a boolean if a field has been set.

### GetLastModified

`func (o *ArtifactMetadata) GetLastModified() string`

GetLastModified returns the LastModified field if non-nil, zero value otherwise.

### GetLastModifiedOk

`func (o *ArtifactMetadata) GetLastModifiedOk() (*string, bool)`

GetLastModifiedOk returns a tuple with the LastModified field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastModified

`func (o *ArtifactMetadata) SetLastModified(v string)`

SetLastModified sets LastModified field to given value.

### HasLastModified

`func (o *ArtifactMetadata) HasLastModified() bool`

HasLastModified returns a boolean if a field has been set.

### GetLatestVersion

`func (o *ArtifactMetadata) GetLatestVersion() string`

GetLatestVersion returns the LatestVersion field if non-nil, zero value otherwise.

### GetLatestVersionOk

`func (o *ArtifactMetadata) GetLatestVersionOk() (*string, bool)`

GetLatestVersionOk returns a tuple with the LatestVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLatestVersion

`func (o *ArtifactMetadata) SetLatestVersion(v string)`

SetLatestVersion sets LatestVersion field to given value.


### GetName

`func (o *ArtifactMetadata) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *ArtifactMetadata) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *ArtifactMetadata) SetName(v string)`

SetName sets Name field to given value.


### GetPackageType

`func (o *ArtifactMetadata) GetPackageType() PackageType`

GetPackageType returns the PackageType field if non-nil, zero value otherwise.

### GetPackageTypeOk

`func (o *ArtifactMetadata) GetPackageTypeOk() (*PackageType, bool)`

GetPackageTypeOk returns a tuple with the PackageType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPackageType

`func (o *ArtifactMetadata) SetPackageType(v PackageType)`

SetPackageType sets PackageType field to given value.

### HasPackageType

`func (o *ArtifactMetadata) HasPackageType() bool`

HasPackageType returns a boolean if a field has been set.

### GetPullCommand

`func (o *ArtifactMetadata) GetPullCommand() string`

GetPullCommand returns the PullCommand field if non-nil, zero value otherwise.

### GetPullCommandOk

`func (o *ArtifactMetadata) GetPullCommandOk() (*string, bool)`

GetPullCommandOk returns a tuple with the PullCommand field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPullCommand

`func (o *ArtifactMetadata) SetPullCommand(v string)`

SetPullCommand sets PullCommand field to given value.

### HasPullCommand

`func (o *ArtifactMetadata) HasPullCommand() bool`

HasPullCommand returns a boolean if a field has been set.

### GetRegistryIdentifier

`func (o *ArtifactMetadata) GetRegistryIdentifier() string`

GetRegistryIdentifier returns the RegistryIdentifier field if non-nil, zero value otherwise.

### GetRegistryIdentifierOk

`func (o *ArtifactMetadata) GetRegistryIdentifierOk() (*string, bool)`

GetRegistryIdentifierOk returns a tuple with the RegistryIdentifier field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRegistryIdentifier

`func (o *ArtifactMetadata) SetRegistryIdentifier(v string)`

SetRegistryIdentifier sets RegistryIdentifier field to given value.


### GetRegistryPath

`func (o *ArtifactMetadata) GetRegistryPath() string`

GetRegistryPath returns the RegistryPath field if non-nil, zero value otherwise.

### GetRegistryPathOk

`func (o *ArtifactMetadata) GetRegistryPathOk() (*string, bool)`

GetRegistryPathOk returns a tuple with the RegistryPath field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRegistryPath

`func (o *ArtifactMetadata) SetRegistryPath(v string)`

SetRegistryPath sets RegistryPath field to given value.


### GetVersion

`func (o *ArtifactMetadata) GetVersion() string`

GetVersion returns the Version field if non-nil, zero value otherwise.

### GetVersionOk

`func (o *ArtifactMetadata) GetVersionOk() (*string, bool)`

GetVersionOk returns a tuple with the Version field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVersion

`func (o *ArtifactMetadata) SetVersion(v string)`

SetVersion sets Version field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


