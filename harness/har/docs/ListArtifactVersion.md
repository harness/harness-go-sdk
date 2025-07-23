# ListArtifactVersion

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ArtifactVersions** | Pointer to [**[]ArtifactVersionMetadata**](ArtifactVersionMetadata.md) | A list of Artifact versions | [optional] 
**ItemCount** | Pointer to **int64** | The total number of items | [optional] 
**PageCount** | Pointer to **int64** | The total number of pages | [optional] 
**PageIndex** | Pointer to **int64** | The current page | [optional] 
**PageSize** | Pointer to **int32** | The number of items per page | [optional] 

## Methods

### NewListArtifactVersion

`func NewListArtifactVersion() *ListArtifactVersion`

NewListArtifactVersion instantiates a new ListArtifactVersion object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewListArtifactVersionWithDefaults

`func NewListArtifactVersionWithDefaults() *ListArtifactVersion`

NewListArtifactVersionWithDefaults instantiates a new ListArtifactVersion object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetArtifactVersions

`func (o *ListArtifactVersion) GetArtifactVersions() []ArtifactVersionMetadata`

GetArtifactVersions returns the ArtifactVersions field if non-nil, zero value otherwise.

### GetArtifactVersionsOk

`func (o *ListArtifactVersion) GetArtifactVersionsOk() (*[]ArtifactVersionMetadata, bool)`

GetArtifactVersionsOk returns a tuple with the ArtifactVersions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetArtifactVersions

`func (o *ListArtifactVersion) SetArtifactVersions(v []ArtifactVersionMetadata)`

SetArtifactVersions sets ArtifactVersions field to given value.

### HasArtifactVersions

`func (o *ListArtifactVersion) HasArtifactVersions() bool`

HasArtifactVersions returns a boolean if a field has been set.

### GetItemCount

`func (o *ListArtifactVersion) GetItemCount() int64`

GetItemCount returns the ItemCount field if non-nil, zero value otherwise.

### GetItemCountOk

`func (o *ListArtifactVersion) GetItemCountOk() (*int64, bool)`

GetItemCountOk returns a tuple with the ItemCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetItemCount

`func (o *ListArtifactVersion) SetItemCount(v int64)`

SetItemCount sets ItemCount field to given value.

### HasItemCount

`func (o *ListArtifactVersion) HasItemCount() bool`

HasItemCount returns a boolean if a field has been set.

### GetPageCount

`func (o *ListArtifactVersion) GetPageCount() int64`

GetPageCount returns the PageCount field if non-nil, zero value otherwise.

### GetPageCountOk

`func (o *ListArtifactVersion) GetPageCountOk() (*int64, bool)`

GetPageCountOk returns a tuple with the PageCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPageCount

`func (o *ListArtifactVersion) SetPageCount(v int64)`

SetPageCount sets PageCount field to given value.

### HasPageCount

`func (o *ListArtifactVersion) HasPageCount() bool`

HasPageCount returns a boolean if a field has been set.

### GetPageIndex

`func (o *ListArtifactVersion) GetPageIndex() int64`

GetPageIndex returns the PageIndex field if non-nil, zero value otherwise.

### GetPageIndexOk

`func (o *ListArtifactVersion) GetPageIndexOk() (*int64, bool)`

GetPageIndexOk returns a tuple with the PageIndex field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPageIndex

`func (o *ListArtifactVersion) SetPageIndex(v int64)`

SetPageIndex sets PageIndex field to given value.

### HasPageIndex

`func (o *ListArtifactVersion) HasPageIndex() bool`

HasPageIndex returns a boolean if a field has been set.

### GetPageSize

`func (o *ListArtifactVersion) GetPageSize() int32`

GetPageSize returns the PageSize field if non-nil, zero value otherwise.

### GetPageSizeOk

`func (o *ListArtifactVersion) GetPageSizeOk() (*int32, bool)`

GetPageSizeOk returns a tuple with the PageSize field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPageSize

`func (o *ListArtifactVersion) SetPageSize(v int32)`

SetPageSize sets PageSize field to given value.

### HasPageSize

`func (o *ListArtifactVersion) HasPageSize() bool`

HasPageSize returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


