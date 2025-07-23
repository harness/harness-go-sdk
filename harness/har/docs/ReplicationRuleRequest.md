# ReplicationRuleRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AllowedPatterns** | **[]string** |  | 
**BlockedPatterns** | **[]string** |  | 
**Destination** | [**ReplicationRegistry**](ReplicationRegistry.md) |  | 
**DestinationType** | **string** |  | 
**Source** | [**ReplicationRegistry**](ReplicationRegistry.md) |  | 
**SourceType** | **string** |  | 

## Methods

### NewReplicationRuleRequest

`func NewReplicationRuleRequest(allowedPatterns []string, blockedPatterns []string, destination ReplicationRegistry, destinationType string, source ReplicationRegistry, sourceType string, ) *ReplicationRuleRequest`

NewReplicationRuleRequest instantiates a new ReplicationRuleRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewReplicationRuleRequestWithDefaults

`func NewReplicationRuleRequestWithDefaults() *ReplicationRuleRequest`

NewReplicationRuleRequestWithDefaults instantiates a new ReplicationRuleRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAllowedPatterns

`func (o *ReplicationRuleRequest) GetAllowedPatterns() []string`

GetAllowedPatterns returns the AllowedPatterns field if non-nil, zero value otherwise.

### GetAllowedPatternsOk

`func (o *ReplicationRuleRequest) GetAllowedPatternsOk() (*[]string, bool)`

GetAllowedPatternsOk returns a tuple with the AllowedPatterns field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAllowedPatterns

`func (o *ReplicationRuleRequest) SetAllowedPatterns(v []string)`

SetAllowedPatterns sets AllowedPatterns field to given value.


### GetBlockedPatterns

`func (o *ReplicationRuleRequest) GetBlockedPatterns() []string`

GetBlockedPatterns returns the BlockedPatterns field if non-nil, zero value otherwise.

### GetBlockedPatternsOk

`func (o *ReplicationRuleRequest) GetBlockedPatternsOk() (*[]string, bool)`

GetBlockedPatternsOk returns a tuple with the BlockedPatterns field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBlockedPatterns

`func (o *ReplicationRuleRequest) SetBlockedPatterns(v []string)`

SetBlockedPatterns sets BlockedPatterns field to given value.


### GetDestination

`func (o *ReplicationRuleRequest) GetDestination() ReplicationRegistry`

GetDestination returns the Destination field if non-nil, zero value otherwise.

### GetDestinationOk

`func (o *ReplicationRuleRequest) GetDestinationOk() (*ReplicationRegistry, bool)`

GetDestinationOk returns a tuple with the Destination field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDestination

`func (o *ReplicationRuleRequest) SetDestination(v ReplicationRegistry)`

SetDestination sets Destination field to given value.


### GetDestinationType

`func (o *ReplicationRuleRequest) GetDestinationType() string`

GetDestinationType returns the DestinationType field if non-nil, zero value otherwise.

### GetDestinationTypeOk

`func (o *ReplicationRuleRequest) GetDestinationTypeOk() (*string, bool)`

GetDestinationTypeOk returns a tuple with the DestinationType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDestinationType

`func (o *ReplicationRuleRequest) SetDestinationType(v string)`

SetDestinationType sets DestinationType field to given value.


### GetSource

`func (o *ReplicationRuleRequest) GetSource() ReplicationRegistry`

GetSource returns the Source field if non-nil, zero value otherwise.

### GetSourceOk

`func (o *ReplicationRuleRequest) GetSourceOk() (*ReplicationRegistry, bool)`

GetSourceOk returns a tuple with the Source field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSource

`func (o *ReplicationRuleRequest) SetSource(v ReplicationRegistry)`

SetSource sets Source field to given value.


### GetSourceType

`func (o *ReplicationRuleRequest) GetSourceType() string`

GetSourceType returns the SourceType field if non-nil, zero value otherwise.

### GetSourceTypeOk

`func (o *ReplicationRuleRequest) GetSourceTypeOk() (*string, bool)`

GetSourceTypeOk returns a tuple with the SourceType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSourceType

`func (o *ReplicationRuleRequest) SetSourceType(v string)`

SetSourceType sets SourceType field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


