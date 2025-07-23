# ListRegistryArtifact

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Artifacts** | [**[]RegistryArtifactMetadata**](RegistryArtifactMetadata.md) | A list of Artifact | 
**ItemCount** | Pointer to **int64** | The total number of items | [optional] 
**PageCount** | Pointer to **int64** | The total number of pages | [optional] 
**PageIndex** | Pointer to **int64** | The current page | [optional] 
**PageSize** | Pointer to **int32** | The number of items per page | [optional] 

## Methods

### NewListRegistryArtifact

`func NewListRegistryArtifact(artifacts []RegistryArtifactMetadata, ) *ListRegistryArtifact`

NewListRegistryArtifact instantiates a new ListRegistryArtifact object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewListRegistryArtifactWithDefaults

`func NewListRegistryArtifactWithDefaults() *ListRegistryArtifact`

NewListRegistryArtifactWithDefaults instantiates a new ListRegistryArtifact object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetArtifacts

`func (o *ListRegistryArtifact) GetArtifacts() []RegistryArtifactMetadata`

GetArtifacts returns the Artifacts field if non-nil, zero value otherwise.

### GetArtifactsOk

`func (o *ListRegistryArtifact) GetArtifactsOk() (*[]RegistryArtifactMetadata, bool)`

GetArtifactsOk returns a tuple with the Artifacts field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetArtifacts

`func (o *ListRegistryArtifact) SetArtifacts(v []RegistryArtifactMetadata)`

SetArtifacts sets Artifacts field to given value.


### GetItemCount

`func (o *ListRegistryArtifact) GetItemCount() int64`

GetItemCount returns the ItemCount field if non-nil, zero value otherwise.

### GetItemCountOk

`func (o *ListRegistryArtifact) GetItemCountOk() (*int64, bool)`

GetItemCountOk returns a tuple with the ItemCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetItemCount

`func (o *ListRegistryArtifact) SetItemCount(v int64)`

SetItemCount sets ItemCount field to given value.

### HasItemCount

`func (o *ListRegistryArtifact) HasItemCount() bool`

HasItemCount returns a boolean if a field has been set.

### GetPageCount

`func (o *ListRegistryArtifact) GetPageCount() int64`

GetPageCount returns the PageCount field if non-nil, zero value otherwise.

### GetPageCountOk

`func (o *ListRegistryArtifact) GetPageCountOk() (*int64, bool)`

GetPageCountOk returns a tuple with the PageCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPageCount

`func (o *ListRegistryArtifact) SetPageCount(v int64)`

SetPageCount sets PageCount field to given value.

### HasPageCount

`func (o *ListRegistryArtifact) HasPageCount() bool`

HasPageCount returns a boolean if a field has been set.

### GetPageIndex

`func (o *ListRegistryArtifact) GetPageIndex() int64`

GetPageIndex returns the PageIndex field if non-nil, zero value otherwise.

### GetPageIndexOk

`func (o *ListRegistryArtifact) GetPageIndexOk() (*int64, bool)`

GetPageIndexOk returns a tuple with the PageIndex field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPageIndex

`func (o *ListRegistryArtifact) SetPageIndex(v int64)`

SetPageIndex sets PageIndex field to given value.

### HasPageIndex

`func (o *ListRegistryArtifact) HasPageIndex() bool`

HasPageIndex returns a boolean if a field has been set.

### GetPageSize

`func (o *ListRegistryArtifact) GetPageSize() int32`

GetPageSize returns the PageSize field if non-nil, zero value otherwise.

### GetPageSizeOk

`func (o *ListRegistryArtifact) GetPageSizeOk() (*int32, bool)`

GetPageSizeOk returns a tuple with the PageSize field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPageSize

`func (o *ListRegistryArtifact) SetPageSize(v int32)`

SetPageSize sets PageSize field to given value.

### HasPageSize

`func (o *ListRegistryArtifact) HasPageSize() bool`

HasPageSize returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


