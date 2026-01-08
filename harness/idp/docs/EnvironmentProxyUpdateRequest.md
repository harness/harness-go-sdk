# EnvironmentProxyUpdateRequest

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**EnvironmentBlueprintVersion** | **string** | Version of the environment blueprint (e.g., v1.0.0) | [optional] [default to null]
**BasedOnIdentifier** | **string** | Reference to the environment this is based on in format environment:account.orgIdentifier.projectIdentifier/identifier | [optional] [default to null]
**Overrides** | **string** |  | [optional] [default to null]
**TargetState** | **string** | Target state for the environment | [optional] [default to null]
**Inputs** | **string** | Inputs for the environment as YAML/JSON string (e.g., ttl, resources) | [optional] [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)

