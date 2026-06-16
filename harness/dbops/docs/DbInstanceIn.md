# DbInstanceIn

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Identifier** | **string** | identifier of the database instance | [default to null]
**Name** | **string** | name of the database instance | [optional] [default to null]
**Tags** | **map[string]string** | tags attached to the database instance | [optional] [default to null]
**Branch** | **string** | branch where the instance is stored  | [optional] [default to null]
**CommitSha** | **string** | Commit SHA to pin the changelog to a specific revision. Mutually exclusive with branch and gitTag. Supports Harness expressions (e.g. &lt;+trigger.commitSha&gt;).  | [optional] [default to null]
**GitTag** | **string** | Git tag to pin the changelog to a specific tagged revision. Mutually exclusive with branch and commitSha. Supports Harness expressions (e.g. &lt;+trigger.tag&gt;).  | [optional] [default to null]
**Connector** | **string** | DB Connector | [default to null]
**Context** | **string** | Liquibase context | [optional] [default to null]
**SubstituteProperties** | **map[string]string** | Placeholder replacement in migration scripts. | [optional] [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)

