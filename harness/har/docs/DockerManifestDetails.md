# DockerManifestDetails

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CreatedAt** | Pointer to **string** |  | [optional] 
**Digest** | **string** |  | 
**DownloadsCount** | Pointer to **int64** |  | [optional] 
**OsArch** | **string** |  | 
**Size** | Pointer to **string** |  | [optional] 
**StoDetails** | Pointer to [**STODetails**](STODetails.md) |  | [optional] 
**StoExecutionId** | Pointer to **string** |  | [optional] 
**StoPipelineId** | Pointer to **string** |  | [optional] 

## Methods

### NewDockerManifestDetails

`func NewDockerManifestDetails(digest string, osArch string, ) *DockerManifestDetails`

NewDockerManifestDetails instantiates a new DockerManifestDetails object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDockerManifestDetailsWithDefaults

`func NewDockerManifestDetailsWithDefaults() *DockerManifestDetails`

NewDockerManifestDetailsWithDefaults instantiates a new DockerManifestDetails object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCreatedAt

`func (o *DockerManifestDetails) GetCreatedAt() string`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *DockerManifestDetails) GetCreatedAtOk() (*string, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *DockerManifestDetails) SetCreatedAt(v string)`

SetCreatedAt sets CreatedAt field to given value.

### HasCreatedAt

`func (o *DockerManifestDetails) HasCreatedAt() bool`

HasCreatedAt returns a boolean if a field has been set.

### GetDigest

`func (o *DockerManifestDetails) GetDigest() string`

GetDigest returns the Digest field if non-nil, zero value otherwise.

### GetDigestOk

`func (o *DockerManifestDetails) GetDigestOk() (*string, bool)`

GetDigestOk returns a tuple with the Digest field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDigest

`func (o *DockerManifestDetails) SetDigest(v string)`

SetDigest sets Digest field to given value.


### GetDownloadsCount

`func (o *DockerManifestDetails) GetDownloadsCount() int64`

GetDownloadsCount returns the DownloadsCount field if non-nil, zero value otherwise.

### GetDownloadsCountOk

`func (o *DockerManifestDetails) GetDownloadsCountOk() (*int64, bool)`

GetDownloadsCountOk returns a tuple with the DownloadsCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDownloadsCount

`func (o *DockerManifestDetails) SetDownloadsCount(v int64)`

SetDownloadsCount sets DownloadsCount field to given value.

### HasDownloadsCount

`func (o *DockerManifestDetails) HasDownloadsCount() bool`

HasDownloadsCount returns a boolean if a field has been set.

### GetOsArch

`func (o *DockerManifestDetails) GetOsArch() string`

GetOsArch returns the OsArch field if non-nil, zero value otherwise.

### GetOsArchOk

`func (o *DockerManifestDetails) GetOsArchOk() (*string, bool)`

GetOsArchOk returns a tuple with the OsArch field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOsArch

`func (o *DockerManifestDetails) SetOsArch(v string)`

SetOsArch sets OsArch field to given value.


### GetSize

`func (o *DockerManifestDetails) GetSize() string`

GetSize returns the Size field if non-nil, zero value otherwise.

### GetSizeOk

`func (o *DockerManifestDetails) GetSizeOk() (*string, bool)`

GetSizeOk returns a tuple with the Size field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSize

`func (o *DockerManifestDetails) SetSize(v string)`

SetSize sets Size field to given value.

### HasSize

`func (o *DockerManifestDetails) HasSize() bool`

HasSize returns a boolean if a field has been set.

### GetStoDetails

`func (o *DockerManifestDetails) GetStoDetails() STODetails`

GetStoDetails returns the StoDetails field if non-nil, zero value otherwise.

### GetStoDetailsOk

`func (o *DockerManifestDetails) GetStoDetailsOk() (*STODetails, bool)`

GetStoDetailsOk returns a tuple with the StoDetails field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStoDetails

`func (o *DockerManifestDetails) SetStoDetails(v STODetails)`

SetStoDetails sets StoDetails field to given value.

### HasStoDetails

`func (o *DockerManifestDetails) HasStoDetails() bool`

HasStoDetails returns a boolean if a field has been set.

### GetStoExecutionId

`func (o *DockerManifestDetails) GetStoExecutionId() string`

GetStoExecutionId returns the StoExecutionId field if non-nil, zero value otherwise.

### GetStoExecutionIdOk

`func (o *DockerManifestDetails) GetStoExecutionIdOk() (*string, bool)`

GetStoExecutionIdOk returns a tuple with the StoExecutionId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStoExecutionId

`func (o *DockerManifestDetails) SetStoExecutionId(v string)`

SetStoExecutionId sets StoExecutionId field to given value.

### HasStoExecutionId

`func (o *DockerManifestDetails) HasStoExecutionId() bool`

HasStoExecutionId returns a boolean if a field has been set.

### GetStoPipelineId

`func (o *DockerManifestDetails) GetStoPipelineId() string`

GetStoPipelineId returns the StoPipelineId field if non-nil, zero value otherwise.

### GetStoPipelineIdOk

`func (o *DockerManifestDetails) GetStoPipelineIdOk() (*string, bool)`

GetStoPipelineIdOk returns a tuple with the StoPipelineId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStoPipelineId

`func (o *DockerManifestDetails) SetStoPipelineId(v string)`

SetStoPipelineId sets StoPipelineId field to given value.

### HasStoPipelineId

`func (o *DockerManifestDetails) HasStoPipelineId() bool`

HasStoPipelineId returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


