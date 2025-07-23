# ArtifactDeploymentsDetail

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Count** | Pointer to **int32** |  | [optional] 
**EnvIdentifier** | Pointer to **string** |  | [optional] 
**EnvName** | Pointer to **string** |  | [optional] 
**EnvType** | Pointer to [**EnvironmentType**](EnvironmentType.md) |  | [optional] 
**InfraIdentifier** | Pointer to **string** |  | [optional] 
**InfraName** | Pointer to **string** |  | [optional] 
**LastDeployedAt** | Pointer to **string** |  | [optional] 
**LastDeployedById** | Pointer to **string** |  | [optional] 
**LastDeployedByName** | Pointer to **string** |  | [optional] 
**LastPipelineExecutionId** | Pointer to **string** |  | [optional] 
**LastPipelineExecutionName** | Pointer to **string** |  | [optional] 
**OrgIdentifier** | Pointer to **string** |  | [optional] 
**PipelineId** | Pointer to **string** |  | [optional] 
**ProjectIdentifier** | Pointer to **string** |  | [optional] 
**ServiceIdentifier** | Pointer to **string** |  | [optional] 
**ServiceName** | Pointer to **string** |  | [optional] 

## Methods

### NewArtifactDeploymentsDetail

`func NewArtifactDeploymentsDetail() *ArtifactDeploymentsDetail`

NewArtifactDeploymentsDetail instantiates a new ArtifactDeploymentsDetail object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewArtifactDeploymentsDetailWithDefaults

`func NewArtifactDeploymentsDetailWithDefaults() *ArtifactDeploymentsDetail`

NewArtifactDeploymentsDetailWithDefaults instantiates a new ArtifactDeploymentsDetail object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCount

`func (o *ArtifactDeploymentsDetail) GetCount() int32`

GetCount returns the Count field if non-nil, zero value otherwise.

### GetCountOk

`func (o *ArtifactDeploymentsDetail) GetCountOk() (*int32, bool)`

GetCountOk returns a tuple with the Count field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCount

`func (o *ArtifactDeploymentsDetail) SetCount(v int32)`

SetCount sets Count field to given value.

### HasCount

`func (o *ArtifactDeploymentsDetail) HasCount() bool`

HasCount returns a boolean if a field has been set.

### GetEnvIdentifier

`func (o *ArtifactDeploymentsDetail) GetEnvIdentifier() string`

GetEnvIdentifier returns the EnvIdentifier field if non-nil, zero value otherwise.

### GetEnvIdentifierOk

`func (o *ArtifactDeploymentsDetail) GetEnvIdentifierOk() (*string, bool)`

GetEnvIdentifierOk returns a tuple with the EnvIdentifier field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnvIdentifier

`func (o *ArtifactDeploymentsDetail) SetEnvIdentifier(v string)`

SetEnvIdentifier sets EnvIdentifier field to given value.

### HasEnvIdentifier

`func (o *ArtifactDeploymentsDetail) HasEnvIdentifier() bool`

HasEnvIdentifier returns a boolean if a field has been set.

### GetEnvName

`func (o *ArtifactDeploymentsDetail) GetEnvName() string`

GetEnvName returns the EnvName field if non-nil, zero value otherwise.

### GetEnvNameOk

`func (o *ArtifactDeploymentsDetail) GetEnvNameOk() (*string, bool)`

GetEnvNameOk returns a tuple with the EnvName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnvName

`func (o *ArtifactDeploymentsDetail) SetEnvName(v string)`

SetEnvName sets EnvName field to given value.

### HasEnvName

`func (o *ArtifactDeploymentsDetail) HasEnvName() bool`

HasEnvName returns a boolean if a field has been set.

### GetEnvType

`func (o *ArtifactDeploymentsDetail) GetEnvType() EnvironmentType`

GetEnvType returns the EnvType field if non-nil, zero value otherwise.

### GetEnvTypeOk

`func (o *ArtifactDeploymentsDetail) GetEnvTypeOk() (*EnvironmentType, bool)`

GetEnvTypeOk returns a tuple with the EnvType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnvType

`func (o *ArtifactDeploymentsDetail) SetEnvType(v EnvironmentType)`

SetEnvType sets EnvType field to given value.

### HasEnvType

`func (o *ArtifactDeploymentsDetail) HasEnvType() bool`

HasEnvType returns a boolean if a field has been set.

### GetInfraIdentifier

`func (o *ArtifactDeploymentsDetail) GetInfraIdentifier() string`

GetInfraIdentifier returns the InfraIdentifier field if non-nil, zero value otherwise.

### GetInfraIdentifierOk

`func (o *ArtifactDeploymentsDetail) GetInfraIdentifierOk() (*string, bool)`

GetInfraIdentifierOk returns a tuple with the InfraIdentifier field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInfraIdentifier

`func (o *ArtifactDeploymentsDetail) SetInfraIdentifier(v string)`

SetInfraIdentifier sets InfraIdentifier field to given value.

### HasInfraIdentifier

`func (o *ArtifactDeploymentsDetail) HasInfraIdentifier() bool`

HasInfraIdentifier returns a boolean if a field has been set.

### GetInfraName

`func (o *ArtifactDeploymentsDetail) GetInfraName() string`

GetInfraName returns the InfraName field if non-nil, zero value otherwise.

### GetInfraNameOk

`func (o *ArtifactDeploymentsDetail) GetInfraNameOk() (*string, bool)`

GetInfraNameOk returns a tuple with the InfraName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInfraName

`func (o *ArtifactDeploymentsDetail) SetInfraName(v string)`

SetInfraName sets InfraName field to given value.

### HasInfraName

`func (o *ArtifactDeploymentsDetail) HasInfraName() bool`

HasInfraName returns a boolean if a field has been set.

### GetLastDeployedAt

`func (o *ArtifactDeploymentsDetail) GetLastDeployedAt() string`

GetLastDeployedAt returns the LastDeployedAt field if non-nil, zero value otherwise.

### GetLastDeployedAtOk

`func (o *ArtifactDeploymentsDetail) GetLastDeployedAtOk() (*string, bool)`

GetLastDeployedAtOk returns a tuple with the LastDeployedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastDeployedAt

`func (o *ArtifactDeploymentsDetail) SetLastDeployedAt(v string)`

SetLastDeployedAt sets LastDeployedAt field to given value.

### HasLastDeployedAt

`func (o *ArtifactDeploymentsDetail) HasLastDeployedAt() bool`

HasLastDeployedAt returns a boolean if a field has been set.

### GetLastDeployedById

`func (o *ArtifactDeploymentsDetail) GetLastDeployedById() string`

GetLastDeployedById returns the LastDeployedById field if non-nil, zero value otherwise.

### GetLastDeployedByIdOk

`func (o *ArtifactDeploymentsDetail) GetLastDeployedByIdOk() (*string, bool)`

GetLastDeployedByIdOk returns a tuple with the LastDeployedById field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastDeployedById

`func (o *ArtifactDeploymentsDetail) SetLastDeployedById(v string)`

SetLastDeployedById sets LastDeployedById field to given value.

### HasLastDeployedById

`func (o *ArtifactDeploymentsDetail) HasLastDeployedById() bool`

HasLastDeployedById returns a boolean if a field has been set.

### GetLastDeployedByName

`func (o *ArtifactDeploymentsDetail) GetLastDeployedByName() string`

GetLastDeployedByName returns the LastDeployedByName field if non-nil, zero value otherwise.

### GetLastDeployedByNameOk

`func (o *ArtifactDeploymentsDetail) GetLastDeployedByNameOk() (*string, bool)`

GetLastDeployedByNameOk returns a tuple with the LastDeployedByName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastDeployedByName

`func (o *ArtifactDeploymentsDetail) SetLastDeployedByName(v string)`

SetLastDeployedByName sets LastDeployedByName field to given value.

### HasLastDeployedByName

`func (o *ArtifactDeploymentsDetail) HasLastDeployedByName() bool`

HasLastDeployedByName returns a boolean if a field has been set.

### GetLastPipelineExecutionId

`func (o *ArtifactDeploymentsDetail) GetLastPipelineExecutionId() string`

GetLastPipelineExecutionId returns the LastPipelineExecutionId field if non-nil, zero value otherwise.

### GetLastPipelineExecutionIdOk

`func (o *ArtifactDeploymentsDetail) GetLastPipelineExecutionIdOk() (*string, bool)`

GetLastPipelineExecutionIdOk returns a tuple with the LastPipelineExecutionId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastPipelineExecutionId

`func (o *ArtifactDeploymentsDetail) SetLastPipelineExecutionId(v string)`

SetLastPipelineExecutionId sets LastPipelineExecutionId field to given value.

### HasLastPipelineExecutionId

`func (o *ArtifactDeploymentsDetail) HasLastPipelineExecutionId() bool`

HasLastPipelineExecutionId returns a boolean if a field has been set.

### GetLastPipelineExecutionName

`func (o *ArtifactDeploymentsDetail) GetLastPipelineExecutionName() string`

GetLastPipelineExecutionName returns the LastPipelineExecutionName field if non-nil, zero value otherwise.

### GetLastPipelineExecutionNameOk

`func (o *ArtifactDeploymentsDetail) GetLastPipelineExecutionNameOk() (*string, bool)`

GetLastPipelineExecutionNameOk returns a tuple with the LastPipelineExecutionName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastPipelineExecutionName

`func (o *ArtifactDeploymentsDetail) SetLastPipelineExecutionName(v string)`

SetLastPipelineExecutionName sets LastPipelineExecutionName field to given value.

### HasLastPipelineExecutionName

`func (o *ArtifactDeploymentsDetail) HasLastPipelineExecutionName() bool`

HasLastPipelineExecutionName returns a boolean if a field has been set.

### GetOrgIdentifier

`func (o *ArtifactDeploymentsDetail) GetOrgIdentifier() string`

GetOrgIdentifier returns the OrgIdentifier field if non-nil, zero value otherwise.

### GetOrgIdentifierOk

`func (o *ArtifactDeploymentsDetail) GetOrgIdentifierOk() (*string, bool)`

GetOrgIdentifierOk returns a tuple with the OrgIdentifier field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrgIdentifier

`func (o *ArtifactDeploymentsDetail) SetOrgIdentifier(v string)`

SetOrgIdentifier sets OrgIdentifier field to given value.

### HasOrgIdentifier

`func (o *ArtifactDeploymentsDetail) HasOrgIdentifier() bool`

HasOrgIdentifier returns a boolean if a field has been set.

### GetPipelineId

`func (o *ArtifactDeploymentsDetail) GetPipelineId() string`

GetPipelineId returns the PipelineId field if non-nil, zero value otherwise.

### GetPipelineIdOk

`func (o *ArtifactDeploymentsDetail) GetPipelineIdOk() (*string, bool)`

GetPipelineIdOk returns a tuple with the PipelineId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPipelineId

`func (o *ArtifactDeploymentsDetail) SetPipelineId(v string)`

SetPipelineId sets PipelineId field to given value.

### HasPipelineId

`func (o *ArtifactDeploymentsDetail) HasPipelineId() bool`

HasPipelineId returns a boolean if a field has been set.

### GetProjectIdentifier

`func (o *ArtifactDeploymentsDetail) GetProjectIdentifier() string`

GetProjectIdentifier returns the ProjectIdentifier field if non-nil, zero value otherwise.

### GetProjectIdentifierOk

`func (o *ArtifactDeploymentsDetail) GetProjectIdentifierOk() (*string, bool)`

GetProjectIdentifierOk returns a tuple with the ProjectIdentifier field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProjectIdentifier

`func (o *ArtifactDeploymentsDetail) SetProjectIdentifier(v string)`

SetProjectIdentifier sets ProjectIdentifier field to given value.

### HasProjectIdentifier

`func (o *ArtifactDeploymentsDetail) HasProjectIdentifier() bool`

HasProjectIdentifier returns a boolean if a field has been set.

### GetServiceIdentifier

`func (o *ArtifactDeploymentsDetail) GetServiceIdentifier() string`

GetServiceIdentifier returns the ServiceIdentifier field if non-nil, zero value otherwise.

### GetServiceIdentifierOk

`func (o *ArtifactDeploymentsDetail) GetServiceIdentifierOk() (*string, bool)`

GetServiceIdentifierOk returns a tuple with the ServiceIdentifier field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServiceIdentifier

`func (o *ArtifactDeploymentsDetail) SetServiceIdentifier(v string)`

SetServiceIdentifier sets ServiceIdentifier field to given value.

### HasServiceIdentifier

`func (o *ArtifactDeploymentsDetail) HasServiceIdentifier() bool`

HasServiceIdentifier returns a boolean if a field has been set.

### GetServiceName

`func (o *ArtifactDeploymentsDetail) GetServiceName() string`

GetServiceName returns the ServiceName field if non-nil, zero value otherwise.

### GetServiceNameOk

`func (o *ArtifactDeploymentsDetail) GetServiceNameOk() (*string, bool)`

GetServiceNameOk returns a tuple with the ServiceName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServiceName

`func (o *ArtifactDeploymentsDetail) SetServiceName(v string)`

SetServiceName sets ServiceName field to given value.

### HasServiceName

`func (o *ArtifactDeploymentsDetail) HasServiceName() bool`

HasServiceName returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


