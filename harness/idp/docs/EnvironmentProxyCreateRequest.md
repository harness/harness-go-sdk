# EnvironmentProxyCreateRequest

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**EnvironmentName** | **string** |  | [default to null]
**EnvironmentIdentifier** | **string** |  | [default to null]
**Owner** | **string** |  | [default to null]
**EnvironmentBlueprintIdentifier** | **string** |  | [default to null]
**EnvironmentBlueprintVersion** | **string** | Version of the environment blueprint (e.g., v1.0.0) | [default to null]
**BasedOnIdentifier** | **string** | Reference to the environment this is based on in format environment:account.orgIdentifier.projectIdentifier/identifier | [optional] [default to null]
**Overrides** | **string** |  | [default to null]
**TargetState** | **string** | Target state for the environment | [optional] [default to null]
**Inputs** | **string** | Inputs for the environment as YAML/JSON string (e.g., ttl, resources) | [optional] [default to null]
**Type_** | **string** | Type of the environment (e.g., ephemeral, persistent) | [optional] [default to null]
**Tags** | **[]string** |  | [optional] [default to null]
**Description** | **string** |  | [optional] [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)

