# {{classname}}

All URIs are relative to *https://app.harness.io/gateway*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetSettingValue**](SettingApi.md#GetSettingValue) | **Get** /ng/api/settings/{identifier} | Get a setting value by identifier
[**GetSettingsList**](SettingApi.md#GetSettingsList) | **Get** /ng/api/settings | Get list of settings under the specified category
[**UpdateSettingValue**](SettingApi.md#UpdateSettingValue) | **Put** /ng/api/settings | Update settings

# **GetSettingValue**
> ResponseDtoSettingValueResponseDto GetSettingValue(ctx, identifier, accountIdentifier, optional)
Get a setting value by identifier

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **identifier** | **string**| This is the Identifier of the Entity | 
  **accountIdentifier** | **string**| Account Identifier for the Entity. | 
 **optional** | ***SettingApiGetSettingValueOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a SettingApiGetSettingValueOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **orgIdentifier** | **optional.String**| Organization Identifier for the Entity. | 
 **projectIdentifier** | **optional.String**| Project Identifier for the Entity. | 

### Return type

[**ResponseDtoSettingValueResponseDto**](ResponseDTOSettingValueResponseDTO.md)

### Authorization

[ApiKey](../README.md#ApiKey)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json, application/yaml

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetSettingsList**
> ResponseDtoListSettingResponseDto GetSettingsList(ctx, accountIdentifier, category, optional)
Get list of settings under the specified category

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **accountIdentifier** | **string**| Account Identifier for the Entity. | 
  **category** | **string**| Category of the Setting. | 
 **optional** | ***SettingApiGetSettingsListOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a SettingApiGetSettingsListOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **orgIdentifier** | **optional.String**| Organization Identifier for the Entity. | 
 **projectIdentifier** | **optional.String**| Project Identifier for the Entity. | 
 **group** | **optional.String**| Group Id of the setting | 
 **includeParentScopes** | **optional.Bool**| Flag to include the settings which only exist at the parent scopes | 

### Return type

[**ResponseDtoListSettingResponseDto**](ResponseDTOListSettingResponseDTO.md)

### Authorization

[ApiKey](../README.md#ApiKey)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json, application/yaml

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **UpdateSettingValue**
> ResponseDtoListSettingUpdateResponseDto UpdateSettingValue(ctx, body, accountIdentifier, optional)
Update settings

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**[]SettingRequestDto**](SettingRequestDTO.md)| List of update requests for settings | 
  **accountIdentifier** | **string**| Account Identifier for the Entity. | 
 **optional** | ***SettingApiUpdateSettingValueOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a SettingApiUpdateSettingValueOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **orgIdentifier** | **optional.**| Organization Identifier for the Entity. | 
 **projectIdentifier** | **optional.**| Project Identifier for the Entity. | 

### Return type

[**ResponseDtoListSettingUpdateResponseDto**](ResponseDTOListSettingUpdateResponseDTO.md)

### Authorization

[ApiKey](../README.md#ApiKey)

### HTTP request headers

 - **Content-Type**: application/json, application/yaml
 - **Accept**: application/json, application/yaml

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

