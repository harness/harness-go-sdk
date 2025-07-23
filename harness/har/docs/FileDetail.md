# FileDetail

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Checksums** | **[]string** |  | 
**CreatedAt** | **string** |  | 
**DownloadCommand** | **string** |  | 
**Name** | **string** |  | 
**Size** | **string** |  | 

## Methods

### NewFileDetail

`func NewFileDetail(checksums []string, createdAt string, downloadCommand string, name string, size string, ) *FileDetail`

NewFileDetail instantiates a new FileDetail object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewFileDetailWithDefaults

`func NewFileDetailWithDefaults() *FileDetail`

NewFileDetailWithDefaults instantiates a new FileDetail object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetChecksums

`func (o *FileDetail) GetChecksums() []string`

GetChecksums returns the Checksums field if non-nil, zero value otherwise.

### GetChecksumsOk

`func (o *FileDetail) GetChecksumsOk() (*[]string, bool)`

GetChecksumsOk returns a tuple with the Checksums field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetChecksums

`func (o *FileDetail) SetChecksums(v []string)`

SetChecksums sets Checksums field to given value.


### GetCreatedAt

`func (o *FileDetail) GetCreatedAt() string`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *FileDetail) GetCreatedAtOk() (*string, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *FileDetail) SetCreatedAt(v string)`

SetCreatedAt sets CreatedAt field to given value.


### GetDownloadCommand

`func (o *FileDetail) GetDownloadCommand() string`

GetDownloadCommand returns the DownloadCommand field if non-nil, zero value otherwise.

### GetDownloadCommandOk

`func (o *FileDetail) GetDownloadCommandOk() (*string, bool)`

GetDownloadCommandOk returns a tuple with the DownloadCommand field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDownloadCommand

`func (o *FileDetail) SetDownloadCommand(v string)`

SetDownloadCommand sets DownloadCommand field to given value.


### GetName

`func (o *FileDetail) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *FileDetail) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *FileDetail) SetName(v string)`

SetName sets Name field to given value.


### GetSize

`func (o *FileDetail) GetSize() string`

GetSize returns the Size field if non-nil, zero value otherwise.

### GetSizeOk

`func (o *FileDetail) GetSizeOk() (*string, bool)`

GetSizeOk returns a tuple with the Size field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSize

`func (o *FileDetail) SetSize(v string)`

SetSize sets Size field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


