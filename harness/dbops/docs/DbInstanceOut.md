# DbInstanceOut

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Identifier** | **string** | identifier of the database instance | [default to null]
**Name** | **string** | name of the database instance | [default to null]
**Created** | **int64** | epoch seconds when the database instance was created | [default to null]
**Updated** | **int64** | epoch seconds when the database instance was last updated | [optional] [default to null]
**Tags** | **map[string]string** | tags attached to the database instance | [optional] [default to null]
**Branch** | **string** | branch where the instance is stored | [optional] [default to null]
**Connector** | **string** | DB Connector | [default to null]
**Context** | **string** | Liquibase context | [optional] [default to null]
**LastAppliedTag** | **string** | Most recent tag applied to the database | [optional] [default to null]
**ToOnboard** | **bool** |  | [optional] [default to null]
**LastDeployedChangeSetTag** | **string** | Tag on last deployed changeSet | [default to null]
**SubstituteProperties** | **map[string]string** | Placeholder replacement in migration scripts. | [optional] [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)

