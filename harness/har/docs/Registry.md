# Registry

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AllowedPattern** | Pointer to **[]string** |  | [optional] 
**BlockedPattern** | Pointer to **[]string** |  | [optional] 
**CleanupPolicy** | Pointer to [**[]CleanupPolicy**](CleanupPolicy.md) |  | [optional] 
**Config** | Pointer to [**RegistryConfig**](RegistryConfig.md) |  | [optional] 
**CreatedAt** | Pointer to **string** |  | [optional] 
**Description** | Pointer to **string** |  | [optional] 
**Identifier** | **string** |  | 
**Labels** | Pointer to **[]string** |  | [optional] 
**ModifiedAt** | Pointer to **string** |  | [optional] 
**PackageType** | [**PackageType**](PackageType.md) |  | 
**Scanners** | Pointer to [**[]Scanner**](Scanner.md) |  | [optional] 
**Url** | **string** |  | 

## Methods

### NewRegistry

`func NewRegistry(identifier string, packageType PackageType, url string, ) *Registry`

NewRegistry instantiates a new Registry object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRegistryWithDefaults

`func NewRegistryWithDefaults() *Registry`

NewRegistryWithDefaults instantiates a new Registry object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAllowedPattern

`func (o *Registry) GetAllowedPattern() []string`

GetAllowedPattern returns the AllowedPattern field if non-nil, zero value otherwise.

### GetAllowedPatternOk

`func (o *Registry) GetAllowedPatternOk() (*[]string, bool)`

GetAllowedPatternOk returns a tuple with the AllowedPattern field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAllowedPattern

`func (o *Registry) SetAllowedPattern(v []string)`

SetAllowedPattern sets AllowedPattern field to given value.

### HasAllowedPattern

`func (o *Registry) HasAllowedPattern() bool`

HasAllowedPattern returns a boolean if a field has been set.

### GetBlockedPattern

`func (o *Registry) GetBlockedPattern() []string`

GetBlockedPattern returns the BlockedPattern field if non-nil, zero value otherwise.

### GetBlockedPatternOk

`func (o *Registry) GetBlockedPatternOk() (*[]string, bool)`

GetBlockedPatternOk returns a tuple with the BlockedPattern field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBlockedPattern

`func (o *Registry) SetBlockedPattern(v []string)`

SetBlockedPattern sets BlockedPattern field to given value.

### HasBlockedPattern

`func (o *Registry) HasBlockedPattern() bool`

HasBlockedPattern returns a boolean if a field has been set.

### GetCleanupPolicy

`func (o *Registry) GetCleanupPolicy() []CleanupPolicy`

GetCleanupPolicy returns the CleanupPolicy field if non-nil, zero value otherwise.

### GetCleanupPolicyOk

`func (o *Registry) GetCleanupPolicyOk() (*[]CleanupPolicy, bool)`

GetCleanupPolicyOk returns a tuple with the CleanupPolicy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCleanupPolicy

`func (o *Registry) SetCleanupPolicy(v []CleanupPolicy)`

SetCleanupPolicy sets CleanupPolicy field to given value.

### HasCleanupPolicy

`func (o *Registry) HasCleanupPolicy() bool`

HasCleanupPolicy returns a boolean if a field has been set.

### GetConfig

`func (o *Registry) GetConfig() RegistryConfig`

GetConfig returns the Config field if non-nil, zero value otherwise.

### GetConfigOk

`func (o *Registry) GetConfigOk() (*RegistryConfig, bool)`

GetConfigOk returns a tuple with the Config field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConfig

`func (o *Registry) SetConfig(v RegistryConfig)`

SetConfig sets Config field to given value.

### HasConfig

`func (o *Registry) HasConfig() bool`

HasConfig returns a boolean if a field has been set.

### GetCreatedAt

`func (o *Registry) GetCreatedAt() string`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *Registry) GetCreatedAtOk() (*string, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *Registry) SetCreatedAt(v string)`

SetCreatedAt sets CreatedAt field to given value.

### HasCreatedAt

`func (o *Registry) HasCreatedAt() bool`

HasCreatedAt returns a boolean if a field has been set.

### GetDescription

`func (o *Registry) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *Registry) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *Registry) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *Registry) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetIdentifier

`func (o *Registry) GetIdentifier() string`

GetIdentifier returns the Identifier field if non-nil, zero value otherwise.

### GetIdentifierOk

`func (o *Registry) GetIdentifierOk() (*string, bool)`

GetIdentifierOk returns a tuple with the Identifier field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIdentifier

`func (o *Registry) SetIdentifier(v string)`

SetIdentifier sets Identifier field to given value.


### GetLabels

`func (o *Registry) GetLabels() []string`

GetLabels returns the Labels field if non-nil, zero value otherwise.

### GetLabelsOk

`func (o *Registry) GetLabelsOk() (*[]string, bool)`

GetLabelsOk returns a tuple with the Labels field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLabels

`func (o *Registry) SetLabels(v []string)`

SetLabels sets Labels field to given value.

### HasLabels

`func (o *Registry) HasLabels() bool`

HasLabels returns a boolean if a field has been set.

### GetModifiedAt

`func (o *Registry) GetModifiedAt() string`

GetModifiedAt returns the ModifiedAt field if non-nil, zero value otherwise.

### GetModifiedAtOk

`func (o *Registry) GetModifiedAtOk() (*string, bool)`

GetModifiedAtOk returns a tuple with the ModifiedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetModifiedAt

`func (o *Registry) SetModifiedAt(v string)`

SetModifiedAt sets ModifiedAt field to given value.

### HasModifiedAt

`func (o *Registry) HasModifiedAt() bool`

HasModifiedAt returns a boolean if a field has been set.

### GetPackageType

`func (o *Registry) GetPackageType() PackageType`

GetPackageType returns the PackageType field if non-nil, zero value otherwise.

### GetPackageTypeOk

`func (o *Registry) GetPackageTypeOk() (*PackageType, bool)`

GetPackageTypeOk returns a tuple with the PackageType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPackageType

`func (o *Registry) SetPackageType(v PackageType)`

SetPackageType sets PackageType field to given value.


### GetScanners

`func (o *Registry) GetScanners() []Scanner`

GetScanners returns the Scanners field if non-nil, zero value otherwise.

### GetScannersOk

`func (o *Registry) GetScannersOk() (*[]Scanner, bool)`

GetScannersOk returns a tuple with the Scanners field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScanners

`func (o *Registry) SetScanners(v []Scanner)`

SetScanners sets Scanners field to given value.

### HasScanners

`func (o *Registry) HasScanners() bool`

HasScanners returns a boolean if a field has been set.

### GetUrl

`func (o *Registry) GetUrl() string`

GetUrl returns the Url field if non-nil, zero value otherwise.

### GetUrlOk

`func (o *Registry) GetUrlOk() (*string, bool)`

GetUrlOk returns a tuple with the Url field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUrl

`func (o *Registry) SetUrl(v string)`

SetUrl sets Url field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


