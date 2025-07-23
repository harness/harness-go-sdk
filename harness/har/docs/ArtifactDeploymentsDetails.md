# ArtifactDeploymentsDetails

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Deployments** | [**ArtifactDeploymentsList**](ArtifactDeploymentsList.md) |  | 
**DeploymentsStats** | Pointer to [**DeploymentStats**](DeploymentStats.md) |  | [optional] 

## Methods

### NewArtifactDeploymentsDetails

`func NewArtifactDeploymentsDetails(deployments ArtifactDeploymentsList, ) *ArtifactDeploymentsDetails`

NewArtifactDeploymentsDetails instantiates a new ArtifactDeploymentsDetails object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewArtifactDeploymentsDetailsWithDefaults

`func NewArtifactDeploymentsDetailsWithDefaults() *ArtifactDeploymentsDetails`

NewArtifactDeploymentsDetailsWithDefaults instantiates a new ArtifactDeploymentsDetails object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDeployments

`func (o *ArtifactDeploymentsDetails) GetDeployments() ArtifactDeploymentsList`

GetDeployments returns the Deployments field if non-nil, zero value otherwise.

### GetDeploymentsOk

`func (o *ArtifactDeploymentsDetails) GetDeploymentsOk() (*ArtifactDeploymentsList, bool)`

GetDeploymentsOk returns a tuple with the Deployments field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeployments

`func (o *ArtifactDeploymentsDetails) SetDeployments(v ArtifactDeploymentsList)`

SetDeployments sets Deployments field to given value.


### GetDeploymentsStats

`func (o *ArtifactDeploymentsDetails) GetDeploymentsStats() DeploymentStats`

GetDeploymentsStats returns the DeploymentsStats field if non-nil, zero value otherwise.

### GetDeploymentsStatsOk

`func (o *ArtifactDeploymentsDetails) GetDeploymentsStatsOk() (*DeploymentStats, bool)`

GetDeploymentsStatsOk returns a tuple with the DeploymentsStats field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeploymentsStats

`func (o *ArtifactDeploymentsDetails) SetDeploymentsStats(v DeploymentStats)`

SetDeploymentsStats sets DeploymentsStats field to given value.

### HasDeploymentsStats

`func (o *ArtifactDeploymentsDetails) HasDeploymentsStats() bool`

HasDeploymentsStats returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


