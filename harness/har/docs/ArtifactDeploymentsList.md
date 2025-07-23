# ArtifactDeploymentsList

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Deployments** | Pointer to [**[]ArtifactDeploymentsDetail**](ArtifactDeploymentsDetail.md) | A list of Artifact | [optional] 
**ItemCount** | Pointer to **int64** | The total number of items | [optional] 
**PageCount** | Pointer to **int64** | The total number of pages | [optional] 
**PageIndex** | Pointer to **int64** | The current page | [optional] 
**PageSize** | Pointer to **int32** | The number of items per page | [optional] 

## Methods

### NewArtifactDeploymentsList

`func NewArtifactDeploymentsList() *ArtifactDeploymentsList`

NewArtifactDeploymentsList instantiates a new ArtifactDeploymentsList object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewArtifactDeploymentsListWithDefaults

`func NewArtifactDeploymentsListWithDefaults() *ArtifactDeploymentsList`

NewArtifactDeploymentsListWithDefaults instantiates a new ArtifactDeploymentsList object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDeployments

`func (o *ArtifactDeploymentsList) GetDeployments() []ArtifactDeploymentsDetail`

GetDeployments returns the Deployments field if non-nil, zero value otherwise.

### GetDeploymentsOk

`func (o *ArtifactDeploymentsList) GetDeploymentsOk() (*[]ArtifactDeploymentsDetail, bool)`

GetDeploymentsOk returns a tuple with the Deployments field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeployments

`func (o *ArtifactDeploymentsList) SetDeployments(v []ArtifactDeploymentsDetail)`

SetDeployments sets Deployments field to given value.

### HasDeployments

`func (o *ArtifactDeploymentsList) HasDeployments() bool`

HasDeployments returns a boolean if a field has been set.

### GetItemCount

`func (o *ArtifactDeploymentsList) GetItemCount() int64`

GetItemCount returns the ItemCount field if non-nil, zero value otherwise.

### GetItemCountOk

`func (o *ArtifactDeploymentsList) GetItemCountOk() (*int64, bool)`

GetItemCountOk returns a tuple with the ItemCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetItemCount

`func (o *ArtifactDeploymentsList) SetItemCount(v int64)`

SetItemCount sets ItemCount field to given value.

### HasItemCount

`func (o *ArtifactDeploymentsList) HasItemCount() bool`

HasItemCount returns a boolean if a field has been set.

### GetPageCount

`func (o *ArtifactDeploymentsList) GetPageCount() int64`

GetPageCount returns the PageCount field if non-nil, zero value otherwise.

### GetPageCountOk

`func (o *ArtifactDeploymentsList) GetPageCountOk() (*int64, bool)`

GetPageCountOk returns a tuple with the PageCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPageCount

`func (o *ArtifactDeploymentsList) SetPageCount(v int64)`

SetPageCount sets PageCount field to given value.

### HasPageCount

`func (o *ArtifactDeploymentsList) HasPageCount() bool`

HasPageCount returns a boolean if a field has been set.

### GetPageIndex

`func (o *ArtifactDeploymentsList) GetPageIndex() int64`

GetPageIndex returns the PageIndex field if non-nil, zero value otherwise.

### GetPageIndexOk

`func (o *ArtifactDeploymentsList) GetPageIndexOk() (*int64, bool)`

GetPageIndexOk returns a tuple with the PageIndex field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPageIndex

`func (o *ArtifactDeploymentsList) SetPageIndex(v int64)`

SetPageIndex sets PageIndex field to given value.

### HasPageIndex

`func (o *ArtifactDeploymentsList) HasPageIndex() bool`

HasPageIndex returns a boolean if a field has been set.

### GetPageSize

`func (o *ArtifactDeploymentsList) GetPageSize() int32`

GetPageSize returns the PageSize field if non-nil, zero value otherwise.

### GetPageSizeOk

`func (o *ArtifactDeploymentsList) GetPageSizeOk() (*int32, bool)`

GetPageSizeOk returns a tuple with the PageSize field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPageSize

`func (o *ArtifactDeploymentsList) SetPageSize(v int32)`

SetPageSize sets PageSize field to given value.

### HasPageSize

`func (o *ArtifactDeploymentsList) HasPageSize() bool`

HasPageSize returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


