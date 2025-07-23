# WebhookRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Description** | Pointer to **string** |  | [optional] 
**Enabled** | **bool** |  | 
**ExtraHeaders** | Pointer to [**[]ExtraHeader**](ExtraHeader.md) |  | [optional] 
**Identifier** | **string** |  | 
**Insecure** | **bool** |  | 
**Name** | **string** |  | 
**SecretIdentifier** | Pointer to **string** |  | [optional] 
**SecretSpaceId** | Pointer to **int64** |  | [optional] 
**SecretSpacePath** | Pointer to **string** |  | [optional] 
**Triggers** | Pointer to [**[]Trigger**](Trigger.md) |  | [optional] 
**Url** | **string** |  | 

## Methods

### NewWebhookRequest

`func NewWebhookRequest(enabled bool, identifier string, insecure bool, name string, url string, ) *WebhookRequest`

NewWebhookRequest instantiates a new WebhookRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewWebhookRequestWithDefaults

`func NewWebhookRequestWithDefaults() *WebhookRequest`

NewWebhookRequestWithDefaults instantiates a new WebhookRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDescription

`func (o *WebhookRequest) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *WebhookRequest) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *WebhookRequest) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *WebhookRequest) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetEnabled

`func (o *WebhookRequest) GetEnabled() bool`

GetEnabled returns the Enabled field if non-nil, zero value otherwise.

### GetEnabledOk

`func (o *WebhookRequest) GetEnabledOk() (*bool, bool)`

GetEnabledOk returns a tuple with the Enabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnabled

`func (o *WebhookRequest) SetEnabled(v bool)`

SetEnabled sets Enabled field to given value.


### GetExtraHeaders

`func (o *WebhookRequest) GetExtraHeaders() []ExtraHeader`

GetExtraHeaders returns the ExtraHeaders field if non-nil, zero value otherwise.

### GetExtraHeadersOk

`func (o *WebhookRequest) GetExtraHeadersOk() (*[]ExtraHeader, bool)`

GetExtraHeadersOk returns a tuple with the ExtraHeaders field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExtraHeaders

`func (o *WebhookRequest) SetExtraHeaders(v []ExtraHeader)`

SetExtraHeaders sets ExtraHeaders field to given value.

### HasExtraHeaders

`func (o *WebhookRequest) HasExtraHeaders() bool`

HasExtraHeaders returns a boolean if a field has been set.

### GetIdentifier

`func (o *WebhookRequest) GetIdentifier() string`

GetIdentifier returns the Identifier field if non-nil, zero value otherwise.

### GetIdentifierOk

`func (o *WebhookRequest) GetIdentifierOk() (*string, bool)`

GetIdentifierOk returns a tuple with the Identifier field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIdentifier

`func (o *WebhookRequest) SetIdentifier(v string)`

SetIdentifier sets Identifier field to given value.


### GetInsecure

`func (o *WebhookRequest) GetInsecure() bool`

GetInsecure returns the Insecure field if non-nil, zero value otherwise.

### GetInsecureOk

`func (o *WebhookRequest) GetInsecureOk() (*bool, bool)`

GetInsecureOk returns a tuple with the Insecure field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInsecure

`func (o *WebhookRequest) SetInsecure(v bool)`

SetInsecure sets Insecure field to given value.


### GetName

`func (o *WebhookRequest) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *WebhookRequest) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *WebhookRequest) SetName(v string)`

SetName sets Name field to given value.


### GetSecretIdentifier

`func (o *WebhookRequest) GetSecretIdentifier() string`

GetSecretIdentifier returns the SecretIdentifier field if non-nil, zero value otherwise.

### GetSecretIdentifierOk

`func (o *WebhookRequest) GetSecretIdentifierOk() (*string, bool)`

GetSecretIdentifierOk returns a tuple with the SecretIdentifier field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSecretIdentifier

`func (o *WebhookRequest) SetSecretIdentifier(v string)`

SetSecretIdentifier sets SecretIdentifier field to given value.

### HasSecretIdentifier

`func (o *WebhookRequest) HasSecretIdentifier() bool`

HasSecretIdentifier returns a boolean if a field has been set.

### GetSecretSpaceId

`func (o *WebhookRequest) GetSecretSpaceId() int64`

GetSecretSpaceId returns the SecretSpaceId field if non-nil, zero value otherwise.

### GetSecretSpaceIdOk

`func (o *WebhookRequest) GetSecretSpaceIdOk() (*int64, bool)`

GetSecretSpaceIdOk returns a tuple with the SecretSpaceId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSecretSpaceId

`func (o *WebhookRequest) SetSecretSpaceId(v int64)`

SetSecretSpaceId sets SecretSpaceId field to given value.

### HasSecretSpaceId

`func (o *WebhookRequest) HasSecretSpaceId() bool`

HasSecretSpaceId returns a boolean if a field has been set.

### GetSecretSpacePath

`func (o *WebhookRequest) GetSecretSpacePath() string`

GetSecretSpacePath returns the SecretSpacePath field if non-nil, zero value otherwise.

### GetSecretSpacePathOk

`func (o *WebhookRequest) GetSecretSpacePathOk() (*string, bool)`

GetSecretSpacePathOk returns a tuple with the SecretSpacePath field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSecretSpacePath

`func (o *WebhookRequest) SetSecretSpacePath(v string)`

SetSecretSpacePath sets SecretSpacePath field to given value.

### HasSecretSpacePath

`func (o *WebhookRequest) HasSecretSpacePath() bool`

HasSecretSpacePath returns a boolean if a field has been set.

### GetTriggers

`func (o *WebhookRequest) GetTriggers() []Trigger`

GetTriggers returns the Triggers field if non-nil, zero value otherwise.

### GetTriggersOk

`func (o *WebhookRequest) GetTriggersOk() (*[]Trigger, bool)`

GetTriggersOk returns a tuple with the Triggers field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTriggers

`func (o *WebhookRequest) SetTriggers(v []Trigger)`

SetTriggers sets Triggers field to given value.

### HasTriggers

`func (o *WebhookRequest) HasTriggers() bool`

HasTriggers returns a boolean if a field has been set.

### GetUrl

`func (o *WebhookRequest) GetUrl() string`

GetUrl returns the Url field if non-nil, zero value otherwise.

### GetUrlOk

`func (o *WebhookRequest) GetUrlOk() (*string, bool)`

GetUrlOk returns a tuple with the Url field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUrl

`func (o *WebhookRequest) SetUrl(v string)`

SetUrl sets Url field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


