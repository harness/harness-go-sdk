# OpenAIModelAuthentication

## Properties

| Name | Type | Description |
|------|------|-------------|
| **Type_** | **OpenAIModelAuthType** | Authentication type: `Token` or `Vertex` |
| **Token** | [**OpenAIModelTokenSpec**](#openaimodeltokenspec) | Token-based authentication (when Type_ = Token) |
| **Vertex** | [**OpenAIModelVertexSpec**](#openaimodelvertexspec) | GCP Vertex AI authentication (when Type_ = Vertex) |

## OpenAIModelTokenSpec

| Name | Type | Description |
|------|------|-------------|
| **TokenRef** | **string** | Reference to a Harness secret containing the OpenAI API token |

## OpenAIModelVertexSpec

| Name | Type | Description |
|------|------|-------------|
| **ServiceAccountKeyRef** | **string** | Reference to a Harness secret containing the GCP service account key |
| **ProjectId** | **string** | GCP project ID |
| **Region** | **string** | GCP region for the Vertex AI endpoint |
