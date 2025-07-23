# RegistryMetadata

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ArtifactsCount** | Pointer to **int64** |  | [optional] 
**Description** | Pointer to **string** |  | [optional] 
**DownloadsCount** | Pointer to **int64** |  | [optional] 
**Identifier** | **string** |  | 
**Labels** | Pointer to **[]string** |  | [optional] 
**LastModified** | Pointer to **string** |  | [optional] 
**PackageType** | [**PackageType**](PackageType.md) |  | 
**Path** | Pointer to **string** |  | [optional] 
**RegistrySize** | Pointer to **string** |  | [optional] 
**Type** | [**RegistryType**](RegistryType.md) |  | 
**Url** | **string** |  | 

## Methods

### NewRegistryMetadata

`func NewRegistryMetadata(identifier string, packageType PackageType, type_ RegistryType, url string, ) *RegistryMetadata`

NewRegistryMetadata instantiates a new RegistryMetadata object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRegistryMetadataWithDefaults

`func NewRegistryMetadataWithDefaults() *RegistryMetadata`

NewRegistryMetadataWithDefaults instantiates a new RegistryMetadata object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetArtifactsCount

`func (o *RegistryMetadata) GetArtifactsCount() int64`

GetArtifactsCount returns the ArtifactsCount field if non-nil, zero value otherwise.

### GetArtifactsCountOk

`func (o *RegistryMetadata) GetArtifactsCountOk() (*int64, bool)`

GetArtifactsCountOk returns a tuple with the ArtifactsCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetArtifactsCount

`func (o *RegistryMetadata) SetArtifactsCount(v int64)`

SetArtifactsCount sets ArtifactsCount field to given value.

### HasArtifactsCount

`func (o *RegistryMetadata) HasArtifactsCount() bool`

HasArtifactsCount returns a boolean if a field has been set.

### GetDescription

`func (o *RegistryMetadata) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *RegistryMetadata) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *RegistryMetadata) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *RegistryMetadata) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetDownloadsCount

`func (o *RegistryMetadata) GetDownloadsCount() int64`

GetDownloadsCount returns the DownloadsCount field if non-nil, zero value otherwise.

### GetDownloadsCountOk

`func (o *RegistryMetadata) GetDownloadsCountOk() (*int64, bool)`

GetDownloadsCountOk returns a tuple with the DownloadsCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDownloadsCount

`func (o *RegistryMetadata) SetDownloadsCount(v int64)`

SetDownloadsCount sets DownloadsCount field to given value.

### HasDownloadsCount

`func (o *RegistryMetadata) HasDownloadsCount() bool`

HasDownloadsCount returns a boolean if a field has been set.

### GetIdentifier

`func (o *RegistryMetadata) GetIdentifier() string`

GetIdentifier returns the Identifier field if non-nil, zero value otherwise.

### GetIdentifierOk

`func (o *RegistryMetadata) GetIdentifierOk() (*string, bool)`

GetIdentifierOk returns a tuple with the Identifier field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIdentifier

`func (o *RegistryMetadata) SetIdentifier(v string)`

SetIdentifier sets Identifier field to given value.


### GetLabels

`func (o *RegistryMetadata) GetLabels() []string`

GetLabels returns the Labels field if non-nil, zero value otherwise.

### GetLabelsOk

`func (o *RegistryMetadata) GetLabelsOk() (*[]string, bool)`

GetLabelsOk returns a tuple with the Labels field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLabels

`func (o *RegistryMetadata) SetLabels(v []string)`

SetLabels sets Labels field to given value.

### HasLabels

`func (o *RegistryMetadata) HasLabels() bool`

HasLabels returns a boolean if a field has been set.

### GetLastModified

`func (o *RegistryMetadata) GetLastModified() string`

GetLastModified returns the LastModified field if non-nil, zero value otherwise.

### GetLastModifiedOk

`func (o *RegistryMetadata) GetLastModifiedOk() (*string, bool)`

GetLastModifiedOk returns a tuple with the LastModified field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastModified

`func (o *RegistryMetadata) SetLastModified(v string)`

SetLastModified sets LastModified field to given value.

### HasLastModified

`func (o *RegistryMetadata) HasLastModified() bool`

HasLastModified returns a boolean if a field has been set.

### GetPackageType

`func (o *RegistryMetadata) GetPackageType() PackageType`

GetPackageType returns the PackageType field if non-nil, zero value otherwise.

### GetPackageTypeOk

`func (o *RegistryMetadata) GetPackageTypeOk() (*PackageType, bool)`

GetPackageTypeOk returns a tuple with the PackageType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPackageType

`func (o *RegistryMetadata) SetPackageType(v PackageType)`

SetPackageType sets PackageType field to given value.


### GetPath

`func (o *RegistryMetadata) GetPath() string`

GetPath returns the Path field if non-nil, zero value otherwise.

### GetPathOk

`func (o *RegistryMetadata) GetPathOk() (*string, bool)`

GetPathOk returns a tuple with the Path field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPath

`func (o *RegistryMetadata) SetPath(v string)`

SetPath sets Path field to given value.

### HasPath

`func (o *RegistryMetadata) HasPath() bool`

HasPath returns a boolean if a field has been set.

### GetRegistrySize

`func (o *RegistryMetadata) GetRegistrySize() string`

GetRegistrySize returns the RegistrySize field if non-nil, zero value otherwise.

### GetRegistrySizeOk

`func (o *RegistryMetadata) GetRegistrySizeOk() (*string, bool)`

GetRegistrySizeOk returns a tuple with the RegistrySize field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRegistrySize

`func (o *RegistryMetadata) SetRegistrySize(v string)`

SetRegistrySize sets RegistrySize field to given value.

### HasRegistrySize

`func (o *RegistryMetadata) HasRegistrySize() bool`

HasRegistrySize returns a boolean if a field has been set.

### GetType

`func (o *RegistryMetadata) GetType() RegistryType`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *RegistryMetadata) GetTypeOk() (*RegistryType, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *RegistryMetadata) SetType(v RegistryType)`

SetType sets Type field to given value.


### GetUrl

`func (o *RegistryMetadata) GetUrl() string`

GetUrl returns the Url field if non-nil, zero value otherwise.

### GetUrlOk

`func (o *RegistryMetadata) GetUrlOk() (*string, bool)`

GetUrlOk returns a tuple with the Url field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUrl

`func (o *RegistryMetadata) SetUrl(v string)`

SetUrl sets Url field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


