# SBOMDetails

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AllowListViolations** | Pointer to **int32** |  | [optional] 
**ArtifactId** | Pointer to **string** |  | [optional] 
**ArtifactSourceId** | Pointer to **string** |  | [optional] 
**AvgScore** | Pointer to **string** |  | [optional] 
**ComponentsCount** | Pointer to **int32** |  | [optional] 
**DenyListViolations** | Pointer to **int32** |  | [optional] 
**MaxScore** | Pointer to **string** |  | [optional] 
**OrchestrationId** | Pointer to **string** |  | [optional] 
**OrgId** | Pointer to **string** |  | [optional] 
**ProjectId** | Pointer to **string** |  | [optional] 

## Methods

### NewSBOMDetails

`func NewSBOMDetails() *SBOMDetails`

NewSBOMDetails instantiates a new SBOMDetails object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSBOMDetailsWithDefaults

`func NewSBOMDetailsWithDefaults() *SBOMDetails`

NewSBOMDetailsWithDefaults instantiates a new SBOMDetails object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAllowListViolations

`func (o *SBOMDetails) GetAllowListViolations() int32`

GetAllowListViolations returns the AllowListViolations field if non-nil, zero value otherwise.

### GetAllowListViolationsOk

`func (o *SBOMDetails) GetAllowListViolationsOk() (*int32, bool)`

GetAllowListViolationsOk returns a tuple with the AllowListViolations field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAllowListViolations

`func (o *SBOMDetails) SetAllowListViolations(v int32)`

SetAllowListViolations sets AllowListViolations field to given value.

### HasAllowListViolations

`func (o *SBOMDetails) HasAllowListViolations() bool`

HasAllowListViolations returns a boolean if a field has been set.

### GetArtifactId

`func (o *SBOMDetails) GetArtifactId() string`

GetArtifactId returns the ArtifactId field if non-nil, zero value otherwise.

### GetArtifactIdOk

`func (o *SBOMDetails) GetArtifactIdOk() (*string, bool)`

GetArtifactIdOk returns a tuple with the ArtifactId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetArtifactId

`func (o *SBOMDetails) SetArtifactId(v string)`

SetArtifactId sets ArtifactId field to given value.

### HasArtifactId

`func (o *SBOMDetails) HasArtifactId() bool`

HasArtifactId returns a boolean if a field has been set.

### GetArtifactSourceId

`func (o *SBOMDetails) GetArtifactSourceId() string`

GetArtifactSourceId returns the ArtifactSourceId field if non-nil, zero value otherwise.

### GetArtifactSourceIdOk

`func (o *SBOMDetails) GetArtifactSourceIdOk() (*string, bool)`

GetArtifactSourceIdOk returns a tuple with the ArtifactSourceId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetArtifactSourceId

`func (o *SBOMDetails) SetArtifactSourceId(v string)`

SetArtifactSourceId sets ArtifactSourceId field to given value.

### HasArtifactSourceId

`func (o *SBOMDetails) HasArtifactSourceId() bool`

HasArtifactSourceId returns a boolean if a field has been set.

### GetAvgScore

`func (o *SBOMDetails) GetAvgScore() string`

GetAvgScore returns the AvgScore field if non-nil, zero value otherwise.

### GetAvgScoreOk

`func (o *SBOMDetails) GetAvgScoreOk() (*string, bool)`

GetAvgScoreOk returns a tuple with the AvgScore field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAvgScore

`func (o *SBOMDetails) SetAvgScore(v string)`

SetAvgScore sets AvgScore field to given value.

### HasAvgScore

`func (o *SBOMDetails) HasAvgScore() bool`

HasAvgScore returns a boolean if a field has been set.

### GetComponentsCount

`func (o *SBOMDetails) GetComponentsCount() int32`

GetComponentsCount returns the ComponentsCount field if non-nil, zero value otherwise.

### GetComponentsCountOk

`func (o *SBOMDetails) GetComponentsCountOk() (*int32, bool)`

GetComponentsCountOk returns a tuple with the ComponentsCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetComponentsCount

`func (o *SBOMDetails) SetComponentsCount(v int32)`

SetComponentsCount sets ComponentsCount field to given value.

### HasComponentsCount

`func (o *SBOMDetails) HasComponentsCount() bool`

HasComponentsCount returns a boolean if a field has been set.

### GetDenyListViolations

`func (o *SBOMDetails) GetDenyListViolations() int32`

GetDenyListViolations returns the DenyListViolations field if non-nil, zero value otherwise.

### GetDenyListViolationsOk

`func (o *SBOMDetails) GetDenyListViolationsOk() (*int32, bool)`

GetDenyListViolationsOk returns a tuple with the DenyListViolations field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDenyListViolations

`func (o *SBOMDetails) SetDenyListViolations(v int32)`

SetDenyListViolations sets DenyListViolations field to given value.

### HasDenyListViolations

`func (o *SBOMDetails) HasDenyListViolations() bool`

HasDenyListViolations returns a boolean if a field has been set.

### GetMaxScore

`func (o *SBOMDetails) GetMaxScore() string`

GetMaxScore returns the MaxScore field if non-nil, zero value otherwise.

### GetMaxScoreOk

`func (o *SBOMDetails) GetMaxScoreOk() (*string, bool)`

GetMaxScoreOk returns a tuple with the MaxScore field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxScore

`func (o *SBOMDetails) SetMaxScore(v string)`

SetMaxScore sets MaxScore field to given value.

### HasMaxScore

`func (o *SBOMDetails) HasMaxScore() bool`

HasMaxScore returns a boolean if a field has been set.

### GetOrchestrationId

`func (o *SBOMDetails) GetOrchestrationId() string`

GetOrchestrationId returns the OrchestrationId field if non-nil, zero value otherwise.

### GetOrchestrationIdOk

`func (o *SBOMDetails) GetOrchestrationIdOk() (*string, bool)`

GetOrchestrationIdOk returns a tuple with the OrchestrationId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrchestrationId

`func (o *SBOMDetails) SetOrchestrationId(v string)`

SetOrchestrationId sets OrchestrationId field to given value.

### HasOrchestrationId

`func (o *SBOMDetails) HasOrchestrationId() bool`

HasOrchestrationId returns a boolean if a field has been set.

### GetOrgId

`func (o *SBOMDetails) GetOrgId() string`

GetOrgId returns the OrgId field if non-nil, zero value otherwise.

### GetOrgIdOk

`func (o *SBOMDetails) GetOrgIdOk() (*string, bool)`

GetOrgIdOk returns a tuple with the OrgId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrgId

`func (o *SBOMDetails) SetOrgId(v string)`

SetOrgId sets OrgId field to given value.

### HasOrgId

`func (o *SBOMDetails) HasOrgId() bool`

HasOrgId returns a boolean if a field has been set.

### GetProjectId

`func (o *SBOMDetails) GetProjectId() string`

GetProjectId returns the ProjectId field if non-nil, zero value otherwise.

### GetProjectIdOk

`func (o *SBOMDetails) GetProjectIdOk() (*string, bool)`

GetProjectIdOk returns a tuple with the ProjectId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProjectId

`func (o *SBOMDetails) SetProjectId(v string)`

SetProjectId sets ProjectId field to given value.

### HasProjectId

`func (o *SBOMDetails) HasProjectId() bool`

HasProjectId returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


