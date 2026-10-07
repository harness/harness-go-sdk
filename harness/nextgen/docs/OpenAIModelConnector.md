# OpenAIModelConnector

## Properties

| Name | Type | Description |
|------|------|-------------|
| **Authentication** | [**OpenAIModelAuthentication**](OpenAIModelAuthentication.md) | Authentication configuration |
| **Url** | **string** | OpenAI API endpoint URL |
| **DelegateSelectors** | **[]string** | Delegate selectors for connection |
| **ExecuteOnDelegate** | **bool** | Whether to execute on delegate |
| **IgnoreTestConnection** | **bool** | Skip test connection on create/update |
| **Model** | **string** | OpenAI model identifier (e.g. gpt-4o) |
| **HarnessManaged** | **bool** | Whether the connector is Harness-managed |
