# AnthropicModelConnector

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Authentication** | [***AnthropicModelAuthentication**](AnthropicModelAuthentication.md) | Authentication configuration for the Anthropic Model Connector. | [default to null]
**Url** | **string** | Anthropic API URL. Only relevant for Token auth. | [optional] [default to null]
**DelegateSelectors** | **[]string** | Delegate selectors for the connector. | [optional] [default to null]
**ExecuteOnDelegate** | **bool** | Whether to execute on delegate. | [optional] [default to null]
**IgnoreTestConnection** | **bool** | Whether to skip test connection on creation. | [optional] [default to null]
**Model** | **string** | Claude model identifier (e.g. claude-opus-4-6). | [optional] [default to null]
**HarnessManaged** | **bool** | Whether the connector is managed by Harness. | [optional] [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)
