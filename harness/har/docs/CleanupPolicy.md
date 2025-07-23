# CleanupPolicy

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ExpireDays** | Pointer to **int32** |  | [optional] 
**Name** | Pointer to **string** |  | [optional] 
**PackagePrefix** | Pointer to **[]string** |  | [optional] 
**VersionPrefix** | Pointer to **[]string** |  | [optional] 

## Methods

### NewCleanupPolicy

`func NewCleanupPolicy() *CleanupPolicy`

NewCleanupPolicy instantiates a new CleanupPolicy object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCleanupPolicyWithDefaults

`func NewCleanupPolicyWithDefaults() *CleanupPolicy`

NewCleanupPolicyWithDefaults instantiates a new CleanupPolicy object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetExpireDays

`func (o *CleanupPolicy) GetExpireDays() int32`

GetExpireDays returns the ExpireDays field if non-nil, zero value otherwise.

### GetExpireDaysOk

`func (o *CleanupPolicy) GetExpireDaysOk() (*int32, bool)`

GetExpireDaysOk returns a tuple with the ExpireDays field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExpireDays

`func (o *CleanupPolicy) SetExpireDays(v int32)`

SetExpireDays sets ExpireDays field to given value.

### HasExpireDays

`func (o *CleanupPolicy) HasExpireDays() bool`

HasExpireDays returns a boolean if a field has been set.

### GetName

`func (o *CleanupPolicy) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *CleanupPolicy) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *CleanupPolicy) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *CleanupPolicy) HasName() bool`

HasName returns a boolean if a field has been set.

### GetPackagePrefix

`func (o *CleanupPolicy) GetPackagePrefix() []string`

GetPackagePrefix returns the PackagePrefix field if non-nil, zero value otherwise.

### GetPackagePrefixOk

`func (o *CleanupPolicy) GetPackagePrefixOk() (*[]string, bool)`

GetPackagePrefixOk returns a tuple with the PackagePrefix field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPackagePrefix

`func (o *CleanupPolicy) SetPackagePrefix(v []string)`

SetPackagePrefix sets PackagePrefix field to given value.

### HasPackagePrefix

`func (o *CleanupPolicy) HasPackagePrefix() bool`

HasPackagePrefix returns a boolean if a field has been set.

### GetVersionPrefix

`func (o *CleanupPolicy) GetVersionPrefix() []string`

GetVersionPrefix returns the VersionPrefix field if non-nil, zero value otherwise.

### GetVersionPrefixOk

`func (o *CleanupPolicy) GetVersionPrefixOk() (*[]string, bool)`

GetVersionPrefixOk returns a tuple with the VersionPrefix field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVersionPrefix

`func (o *CleanupPolicy) SetVersionPrefix(v []string)`

SetVersionPrefix sets VersionPrefix field to given value.

### HasVersionPrefix

`func (o *CleanupPolicy) HasVersionPrefix() bool`

HasVersionPrefix returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


