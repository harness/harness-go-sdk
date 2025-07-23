# ReplicationRule

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AllowedPatterns** | **[]string** |  | 
**BlockedPatterns** | **[]string** |  | 
**CreatedAt** | **string** |  | 
**Destination** | [**ReplicationRegistry**](ReplicationRegistry.md) |  | 
**DestinationType** | **string** |  | 
**Identifier** | **string** |  | 
**ModifiedAt** | **string** |  | 
**ParentRef** | **string** |  | 
**Source** | [**ReplicationRegistry**](ReplicationRegistry.md) |  | 
**SourceType** | **string** |  | 

## Methods

### NewReplicationRule

`func NewReplicationRule(allowedPatterns []string, blockedPatterns []string, createdAt string, destination ReplicationRegistry, destinationType string, identifier string, modifiedAt string, parentRef string, source ReplicationRegistry, sourceType string, ) *ReplicationRule`

NewReplicationRule instantiates a new ReplicationRule object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewReplicationRuleWithDefaults

`func NewReplicationRuleWithDefaults() *ReplicationRule`

NewReplicationRuleWithDefaults instantiates a new ReplicationRule object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAllowedPatterns

`func (o *ReplicationRule) GetAllowedPatterns() []string`

GetAllowedPatterns returns the AllowedPatterns field if non-nil, zero value otherwise.

### GetAllowedPatternsOk

`func (o *ReplicationRule) GetAllowedPatternsOk() (*[]string, bool)`

GetAllowedPatternsOk returns a tuple with the AllowedPatterns field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAllowedPatterns

`func (o *ReplicationRule) SetAllowedPatterns(v []string)`

SetAllowedPatterns sets AllowedPatterns field to given value.


### GetBlockedPatterns

`func (o *ReplicationRule) GetBlockedPatterns() []string`

GetBlockedPatterns returns the BlockedPatterns field if non-nil, zero value otherwise.

### GetBlockedPatternsOk

`func (o *ReplicationRule) GetBlockedPatternsOk() (*[]string, bool)`

GetBlockedPatternsOk returns a tuple with the BlockedPatterns field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBlockedPatterns

`func (o *ReplicationRule) SetBlockedPatterns(v []string)`

SetBlockedPatterns sets BlockedPatterns field to given value.


### GetCreatedAt

`func (o *ReplicationRule) GetCreatedAt() string`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *ReplicationRule) GetCreatedAtOk() (*string, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *ReplicationRule) SetCreatedAt(v string)`

SetCreatedAt sets CreatedAt field to given value.


### GetDestination

`func (o *ReplicationRule) GetDestination() ReplicationRegistry`

GetDestination returns the Destination field if non-nil, zero value otherwise.

### GetDestinationOk

`func (o *ReplicationRule) GetDestinationOk() (*ReplicationRegistry, bool)`

GetDestinationOk returns a tuple with the Destination field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDestination

`func (o *ReplicationRule) SetDestination(v ReplicationRegistry)`

SetDestination sets Destination field to given value.


### GetDestinationType

`func (o *ReplicationRule) GetDestinationType() string`

GetDestinationType returns the DestinationType field if non-nil, zero value otherwise.

### GetDestinationTypeOk

`func (o *ReplicationRule) GetDestinationTypeOk() (*string, bool)`

GetDestinationTypeOk returns a tuple with the DestinationType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDestinationType

`func (o *ReplicationRule) SetDestinationType(v string)`

SetDestinationType sets DestinationType field to given value.


### GetIdentifier

`func (o *ReplicationRule) GetIdentifier() string`

GetIdentifier returns the Identifier field if non-nil, zero value otherwise.

### GetIdentifierOk

`func (o *ReplicationRule) GetIdentifierOk() (*string, bool)`

GetIdentifierOk returns a tuple with the Identifier field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIdentifier

`func (o *ReplicationRule) SetIdentifier(v string)`

SetIdentifier sets Identifier field to given value.


### GetModifiedAt

`func (o *ReplicationRule) GetModifiedAt() string`

GetModifiedAt returns the ModifiedAt field if non-nil, zero value otherwise.

### GetModifiedAtOk

`func (o *ReplicationRule) GetModifiedAtOk() (*string, bool)`

GetModifiedAtOk returns a tuple with the ModifiedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetModifiedAt

`func (o *ReplicationRule) SetModifiedAt(v string)`

SetModifiedAt sets ModifiedAt field to given value.


### GetParentRef

`func (o *ReplicationRule) GetParentRef() string`

GetParentRef returns the ParentRef field if non-nil, zero value otherwise.

### GetParentRefOk

`func (o *ReplicationRule) GetParentRefOk() (*string, bool)`

GetParentRefOk returns a tuple with the ParentRef field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetParentRef

`func (o *ReplicationRule) SetParentRef(v string)`

SetParentRef sets ParentRef field to given value.


### GetSource

`func (o *ReplicationRule) GetSource() ReplicationRegistry`

GetSource returns the Source field if non-nil, zero value otherwise.

### GetSourceOk

`func (o *ReplicationRule) GetSourceOk() (*ReplicationRegistry, bool)`

GetSourceOk returns a tuple with the Source field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSource

`func (o *ReplicationRule) SetSource(v ReplicationRegistry)`

SetSource sets Source field to given value.


### GetSourceType

`func (o *ReplicationRule) GetSourceType() string`

GetSourceType returns the SourceType field if non-nil, zero value otherwise.

### GetSourceTypeOk

`func (o *ReplicationRule) GetSourceTypeOk() (*string, bool)`

GetSourceTypeOk returns a tuple with the SourceType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSourceType

`func (o *ReplicationRule) SetSourceType(v string)`

SetSourceType sets SourceType field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


