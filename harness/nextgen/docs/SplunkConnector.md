# SplunkConnector

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**SplunkUrl** | **string** | URL of the Splunk server | [default to null]
**AccountId** | **string** | Splunk account identifier | [default to null]
**DelegateSelectors** | **[]string** | Tags to filter delegates for connection | [optional] [default to null]
**AuthType** | **string** | Authentication type. Valid values: "UsernamePassword", "Bearer Token(HTTP Header)", "HEC Token", "Anonymous" | [optional] [default to null]
**Username** | **string** | Username for authentication (required for UsernamePassword auth) | [optional] [default to null]
**PasswordRef** | **string** | Reference to secret containing password (required for UsernamePassword auth) | [optional] [default to null]
**TokenRef** | **string** | Reference to secret containing token (required for BearerToken and HECToken auth) | [optional] [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)

