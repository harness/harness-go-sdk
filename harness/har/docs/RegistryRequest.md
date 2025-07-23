# RegistryRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AllowedPattern** | Pointer to **[]string** |  | [optional] 
**BlockedPattern** | Pointer to **[]string** |  | [optional] 
**CleanupPolicy** | Pointer to [**[]CleanupPolicy**](CleanupPolicy.md) |  | [optional] 
**Config** | Pointer to [**RegistryConfig**](RegistryConfig.md) |  | [optional] 
**Description** | Pointer to **string** |  | [optional] 
**Identifier** | **string** |  | 
**Labels** | Pointer to **[]string** |  | [optional] 
**PackageType** | [**PackageType**](PackageType.md) |  | 
**ParentRef** | Pointer to **string** |  | [optional] 
**Scanners** | Pointer to [**[]Scanner**](Scanner.md) |  | [optional] 

## Methods

### NewRegistryRequest

`func NewRegistryRequest(identifier string, packageType PackageType, ) *RegistryRequest`

NewRegistryRequest instantiates a new RegistryRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRegistryRequestWithDefaults

`func NewRegistryRequestWithDefaults() *RegistryRequest`

NewRegistryRequestWithDefaults instantiates a new RegistryRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAllowedPattern

`func (o *RegistryRequest) GetAllowedPattern() []string`

GetAllowedPattern returns the AllowedPattern field if non-nil, zero value otherwise.

### GetAllowedPatternOk

`func (o *RegistryRequest) GetAllowedPatternOk() (*[]string, bool)`

GetAllowedPatternOk returns a tuple with the AllowedPattern field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAllowedPattern

`func (o *RegistryRequest) SetAllowedPattern(v []string)`

SetAllowedPattern sets AllowedPattern field to given value.

### HasAllowedPattern

`func (o *RegistryRequest) HasAllowedPattern() bool`

HasAllowedPattern returns a boolean if a field has been set.

### GetBlockedPattern

`func (o *RegistryRequest) GetBlockedPattern() []string`

GetBlockedPattern returns the BlockedPattern field if non-nil, zero value otherwise.

### GetBlockedPatternOk

`func (o *RegistryRequest) GetBlockedPatternOk() (*[]string, bool)`

GetBlockedPatternOk returns a tuple with the BlockedPattern field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBlockedPattern

`func (o *RegistryRequest) SetBlockedPattern(v []string)`

SetBlockedPattern sets BlockedPattern field to given value.

### HasBlockedPattern

`func (o *RegistryRequest) HasBlockedPattern() bool`

HasBlockedPattern returns a boolean if a field has been set.

### GetCleanupPolicy

`func (o *RegistryRequest) GetCleanupPolicy() []CleanupPolicy`

GetCleanupPolicy returns the CleanupPolicy field if non-nil, zero value otherwise.

### GetCleanupPolicyOk

`func (o *RegistryRequest) GetCleanupPolicyOk() (*[]CleanupPolicy, bool)`

GetCleanupPolicyOk returns a tuple with the CleanupPolicy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCleanupPolicy

`func (o *RegistryRequest) SetCleanupPolicy(v []CleanupPolicy)`

SetCleanupPolicy sets CleanupPolicy field to given value.

### HasCleanupPolicy

`func (o *RegistryRequest) HasCleanupPolicy() bool`

HasCleanupPolicy returns a boolean if a field has been set.

### GetConfig

`func (o *RegistryRequest) GetConfig() RegistryConfig`

GetConfig returns the Config field if non-nil, zero value otherwise.

### GetConfigOk

`func (o *RegistryRequest) GetConfigOk() (*RegistryConfig, bool)`

GetConfigOk returns a tuple with the Config field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConfig

`func (o *RegistryRequest) SetConfig(v RegistryConfig)`

SetConfig sets Config field to given value.

### HasConfig

`func (o *RegistryRequest) HasConfig() bool`

HasConfig returns a boolean if a field has been set.

### GetDescription

`func (o *RegistryRequest) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *RegistryRequest) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *RegistryRequest) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *RegistryRequest) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetIdentifier

`func (o *RegistryRequest) GetIdentifier() string`

GetIdentifier returns the Identifier field if non-nil, zero value otherwise.

### GetIdentifierOk

`func (o *RegistryRequest) GetIdentifierOk() (*string, bool)`

GetIdentifierOk returns a tuple with the Identifier field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIdentifier

`func (o *RegistryRequest) SetIdentifier(v string)`

SetIdentifier sets Identifier field to given value.


### GetLabels

`func (o *RegistryRequest) GetLabels() []string`

GetLabels returns the Labels field if non-nil, zero value otherwise.

### GetLabelsOk

`func (o *RegistryRequest) GetLabelsOk() (*[]string, bool)`

GetLabelsOk returns a tuple with the Labels field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLabels

`func (o *RegistryRequest) SetLabels(v []string)`

SetLabels sets Labels field to given value.

### HasLabels

`func (o *RegistryRequest) HasLabels() bool`

HasLabels returns a boolean if a field has been set.

### GetPackageType

`func (o *RegistryRequest) GetPackageType() PackageType`

GetPackageType returns the PackageType field if non-nil, zero value otherwise.

### GetPackageTypeOk

`func (o *RegistryRequest) GetPackageTypeOk() (*PackageType, bool)`

GetPackageTypeOk returns a tuple with the PackageType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPackageType

`func (o *RegistryRequest) SetPackageType(v PackageType)`

SetPackageType sets PackageType field to given value.


### GetParentRef

`func (o *RegistryRequest) GetParentRef() string`

GetParentRef returns the ParentRef field if non-nil, zero value otherwise.

### GetParentRefOk

`func (o *RegistryRequest) GetParentRefOk() (*string, bool)`

GetParentRefOk returns a tuple with the ParentRef field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetParentRef

`func (o *RegistryRequest) SetParentRef(v string)`

SetParentRef sets ParentRef field to given value.

### HasParentRef

`func (o *RegistryRequest) HasParentRef() bool`

HasParentRef returns a boolean if a field has been set.

### GetScanners

`func (o *RegistryRequest) GetScanners() []Scanner`

GetScanners returns the Scanners field if non-nil, zero value otherwise.

### GetScannersOk

`func (o *RegistryRequest) GetScannersOk() (*[]Scanner, bool)`

GetScannersOk returns a tuple with the Scanners field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScanners

`func (o *RegistryRequest) SetScanners(v []Scanner)`

SetScanners sets Scanners field to given value.

### HasScanners

`func (o *RegistryRequest) HasScanners() bool`

HasScanners returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


