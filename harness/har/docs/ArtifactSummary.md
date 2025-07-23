# ArtifactSummary

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CreatedAt** | Pointer to **string** |  | [optional] 
**DownloadsCount** | Pointer to **int64** |  | [optional] 
**ImageName** | **string** |  | 
**Labels** | Pointer to **[]string** |  | [optional] 
**ModifiedAt** | Pointer to **string** |  | [optional] 
**PackageType** | [**PackageType**](PackageType.md) |  | 

## Methods

### NewArtifactSummary

`func NewArtifactSummary(imageName string, packageType PackageType, ) *ArtifactSummary`

NewArtifactSummary instantiates a new ArtifactSummary object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewArtifactSummaryWithDefaults

`func NewArtifactSummaryWithDefaults() *ArtifactSummary`

NewArtifactSummaryWithDefaults instantiates a new ArtifactSummary object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCreatedAt

`func (o *ArtifactSummary) GetCreatedAt() string`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *ArtifactSummary) GetCreatedAtOk() (*string, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *ArtifactSummary) SetCreatedAt(v string)`

SetCreatedAt sets CreatedAt field to given value.

### HasCreatedAt

`func (o *ArtifactSummary) HasCreatedAt() bool`

HasCreatedAt returns a boolean if a field has been set.

### GetDownloadsCount

`func (o *ArtifactSummary) GetDownloadsCount() int64`

GetDownloadsCount returns the DownloadsCount field if non-nil, zero value otherwise.

### GetDownloadsCountOk

`func (o *ArtifactSummary) GetDownloadsCountOk() (*int64, bool)`

GetDownloadsCountOk returns a tuple with the DownloadsCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDownloadsCount

`func (o *ArtifactSummary) SetDownloadsCount(v int64)`

SetDownloadsCount sets DownloadsCount field to given value.

### HasDownloadsCount

`func (o *ArtifactSummary) HasDownloadsCount() bool`

HasDownloadsCount returns a boolean if a field has been set.

### GetImageName

`func (o *ArtifactSummary) GetImageName() string`

GetImageName returns the ImageName field if non-nil, zero value otherwise.

### GetImageNameOk

`func (o *ArtifactSummary) GetImageNameOk() (*string, bool)`

GetImageNameOk returns a tuple with the ImageName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetImageName

`func (o *ArtifactSummary) SetImageName(v string)`

SetImageName sets ImageName field to given value.


### GetLabels

`func (o *ArtifactSummary) GetLabels() []string`

GetLabels returns the Labels field if non-nil, zero value otherwise.

### GetLabelsOk

`func (o *ArtifactSummary) GetLabelsOk() (*[]string, bool)`

GetLabelsOk returns a tuple with the Labels field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLabels

`func (o *ArtifactSummary) SetLabels(v []string)`

SetLabels sets Labels field to given value.

### HasLabels

`func (o *ArtifactSummary) HasLabels() bool`

HasLabels returns a boolean if a field has been set.

### GetModifiedAt

`func (o *ArtifactSummary) GetModifiedAt() string`

GetModifiedAt returns the ModifiedAt field if non-nil, zero value otherwise.

### GetModifiedAtOk

`func (o *ArtifactSummary) GetModifiedAtOk() (*string, bool)`

GetModifiedAtOk returns a tuple with the ModifiedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetModifiedAt

`func (o *ArtifactSummary) SetModifiedAt(v string)`

SetModifiedAt sets ModifiedAt field to given value.

### HasModifiedAt

`func (o *ArtifactSummary) HasModifiedAt() bool`

HasModifiedAt returns a boolean if a field has been set.

### GetPackageType

`func (o *ArtifactSummary) GetPackageType() PackageType`

GetPackageType returns the PackageType field if non-nil, zero value otherwise.

### GetPackageTypeOk

`func (o *ArtifactSummary) GetPackageTypeOk() (*PackageType, bool)`

GetPackageTypeOk returns a tuple with the PackageType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPackageType

`func (o *ArtifactSummary) SetPackageType(v PackageType)`

SetPackageType sets PackageType field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


