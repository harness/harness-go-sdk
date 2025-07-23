# ListMigrationImage

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Images** | [**[]MigrationImage**](MigrationImage.md) | A list of Artifact versions | 
**ItemCount** | **int64** | The total number of items | 
**PageCount** | **int64** | The total number of pages | 
**PageIndex** | **int64** | The current page | 
**PageSize** | **int32** | The number of items per page | 

## Methods

### NewListMigrationImage

`func NewListMigrationImage(images []MigrationImage, itemCount int64, pageCount int64, pageIndex int64, pageSize int32, ) *ListMigrationImage`

NewListMigrationImage instantiates a new ListMigrationImage object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewListMigrationImageWithDefaults

`func NewListMigrationImageWithDefaults() *ListMigrationImage`

NewListMigrationImageWithDefaults instantiates a new ListMigrationImage object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetImages

`func (o *ListMigrationImage) GetImages() []MigrationImage`

GetImages returns the Images field if non-nil, zero value otherwise.

### GetImagesOk

`func (o *ListMigrationImage) GetImagesOk() (*[]MigrationImage, bool)`

GetImagesOk returns a tuple with the Images field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetImages

`func (o *ListMigrationImage) SetImages(v []MigrationImage)`

SetImages sets Images field to given value.


### GetItemCount

`func (o *ListMigrationImage) GetItemCount() int64`

GetItemCount returns the ItemCount field if non-nil, zero value otherwise.

### GetItemCountOk

`func (o *ListMigrationImage) GetItemCountOk() (*int64, bool)`

GetItemCountOk returns a tuple with the ItemCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetItemCount

`func (o *ListMigrationImage) SetItemCount(v int64)`

SetItemCount sets ItemCount field to given value.


### GetPageCount

`func (o *ListMigrationImage) GetPageCount() int64`

GetPageCount returns the PageCount field if non-nil, zero value otherwise.

### GetPageCountOk

`func (o *ListMigrationImage) GetPageCountOk() (*int64, bool)`

GetPageCountOk returns a tuple with the PageCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPageCount

`func (o *ListMigrationImage) SetPageCount(v int64)`

SetPageCount sets PageCount field to given value.


### GetPageIndex

`func (o *ListMigrationImage) GetPageIndex() int64`

GetPageIndex returns the PageIndex field if non-nil, zero value otherwise.

### GetPageIndexOk

`func (o *ListMigrationImage) GetPageIndexOk() (*int64, bool)`

GetPageIndexOk returns a tuple with the PageIndex field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPageIndex

`func (o *ListMigrationImage) SetPageIndex(v int64)`

SetPageIndex sets PageIndex field to given value.


### GetPageSize

`func (o *ListMigrationImage) GetPageSize() int32`

GetPageSize returns the PageSize field if non-nil, zero value otherwise.

### GetPageSizeOk

`func (o *ListMigrationImage) GetPageSizeOk() (*int32, bool)`

GetPageSizeOk returns a tuple with the PageSize field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPageSize

`func (o *ListMigrationImage) SetPageSize(v int32)`

SetPageSize sets PageSize field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


