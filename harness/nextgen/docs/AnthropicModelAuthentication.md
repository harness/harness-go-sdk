# AnthropicModelAuthentication

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Type_** | **string** | Authentication type. One of: Token, BedrockApiKey, Vertex, CloudProvider. | [default to null]
**Spec** | **json.RawMessage** | Authentication spec (polymorphic, dispatched by Type_). | [default to null]

## Authentication Types

### Token
Direct Anthropic API key authentication.

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**TokenRef** | **string** | Harness secret reference for the Anthropic API key. | [default to null]

### BedrockApiKey
AWS Bedrock API key authentication.

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ApiKeyRef** | **string** | Harness secret reference for the AWS Bedrock API key. | [default to null]
**Region** | **string** | AWS region (e.g. us-east-1). | [default to null]

### Vertex
Google Cloud Vertex AI authentication.

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ServiceAccountKeyRef** | **string** | Harness secret reference for the GCP service account key. | [default to null]
**ProjectId** | **string** | GCP project ID. | [default to null]
**Region** | **string** | GCP region (e.g. us-central1). | [default to null]

### CloudProvider
References an existing AWS, GCP, or Azure cloud provider connector.

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Type_** | **string** | Cloud provider type. One of: AWS, GCP, Azure. | [default to null]
**ConnectorRef** | **string** | Reference to an existing cloud provider connector. | [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)
