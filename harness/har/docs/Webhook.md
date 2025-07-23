# Webhook

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CreatedAt** | Pointer to **string** |  | [optional] 
**CreatedBy** | Pointer to **int64** |  | [optional] 
**Description** | Pointer to **string** |  | [optional] 
**Enabled** | **bool** |  | 
**ExtraHeaders** | Pointer to [**[]ExtraHeader**](ExtraHeader.md) |  | [optional] 
**Identifier** | **string** |  | 
**Insecure** | **bool** |  | 
**Internal** | Pointer to **bool** |  | [optional] 
**LatestExecutionResult** | Pointer to [**WebhookExecResult**](WebhookExecResult.md) |  | [optional] 
**ModifiedAt** | Pointer to **string** |  | [optional] 
**Name** | **string** |  | 
**SecretIdentifier** | Pointer to **string** |  | [optional] 
**SecretSpaceId** | Pointer to **int64** |  | [optional] 
**SecretSpacePath** | Pointer to **string** |  | [optional] 
**Triggers** | Pointer to [**[]Trigger**](Trigger.md) |  | [optional] 
**Url** | **string** |  | 
**Version** | Pointer to **int64** |  | [optional] 

## Methods

### NewWebhook

`func NewWebhook(enabled bool, identifier string, insecure bool, name string, url string, ) *Webhook`

NewWebhook instantiates a new Webhook object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewWebhookWithDefaults

`func NewWebhookWithDefaults() *Webhook`

NewWebhookWithDefaults instantiates a new Webhook object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCreatedAt

`func (o *Webhook) GetCreatedAt() string`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *Webhook) GetCreatedAtOk() (*string, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *Webhook) SetCreatedAt(v string)`

SetCreatedAt sets CreatedAt field to given value.

### HasCreatedAt

`func (o *Webhook) HasCreatedAt() bool`

HasCreatedAt returns a boolean if a field has been set.

### GetCreatedBy

`func (o *Webhook) GetCreatedBy() int64`

GetCreatedBy returns the CreatedBy field if non-nil, zero value otherwise.

### GetCreatedByOk

`func (o *Webhook) GetCreatedByOk() (*int64, bool)`

GetCreatedByOk returns a tuple with the CreatedBy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedBy

`func (o *Webhook) SetCreatedBy(v int64)`

SetCreatedBy sets CreatedBy field to given value.

### HasCreatedBy

`func (o *Webhook) HasCreatedBy() bool`

HasCreatedBy returns a boolean if a field has been set.

### GetDescription

`func (o *Webhook) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *Webhook) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *Webhook) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *Webhook) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetEnabled

`func (o *Webhook) GetEnabled() bool`

GetEnabled returns the Enabled field if non-nil, zero value otherwise.

### GetEnabledOk

`func (o *Webhook) GetEnabledOk() (*bool, bool)`

GetEnabledOk returns a tuple with the Enabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnabled

`func (o *Webhook) SetEnabled(v bool)`

SetEnabled sets Enabled field to given value.


### GetExtraHeaders

`func (o *Webhook) GetExtraHeaders() []ExtraHeader`

GetExtraHeaders returns the ExtraHeaders field if non-nil, zero value otherwise.

### GetExtraHeadersOk

`func (o *Webhook) GetExtraHeadersOk() (*[]ExtraHeader, bool)`

GetExtraHeadersOk returns a tuple with the ExtraHeaders field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExtraHeaders

`func (o *Webhook) SetExtraHeaders(v []ExtraHeader)`

SetExtraHeaders sets ExtraHeaders field to given value.

### HasExtraHeaders

`func (o *Webhook) HasExtraHeaders() bool`

HasExtraHeaders returns a boolean if a field has been set.

### GetIdentifier

`func (o *Webhook) GetIdentifier() string`

GetIdentifier returns the Identifier field if non-nil, zero value otherwise.

### GetIdentifierOk

`func (o *Webhook) GetIdentifierOk() (*string, bool)`

GetIdentifierOk returns a tuple with the Identifier field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIdentifier

`func (o *Webhook) SetIdentifier(v string)`

SetIdentifier sets Identifier field to given value.


### GetInsecure

`func (o *Webhook) GetInsecure() bool`

GetInsecure returns the Insecure field if non-nil, zero value otherwise.

### GetInsecureOk

`func (o *Webhook) GetInsecureOk() (*bool, bool)`

GetInsecureOk returns a tuple with the Insecure field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInsecure

`func (o *Webhook) SetInsecure(v bool)`

SetInsecure sets Insecure field to given value.


### GetInternal

`func (o *Webhook) GetInternal() bool`

GetInternal returns the Internal field if non-nil, zero value otherwise.

### GetInternalOk

`func (o *Webhook) GetInternalOk() (*bool, bool)`

GetInternalOk returns a tuple with the Internal field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInternal

`func (o *Webhook) SetInternal(v bool)`

SetInternal sets Internal field to given value.

### HasInternal

`func (o *Webhook) HasInternal() bool`

HasInternal returns a boolean if a field has been set.

### GetLatestExecutionResult

`func (o *Webhook) GetLatestExecutionResult() WebhookExecResult`

GetLatestExecutionResult returns the LatestExecutionResult field if non-nil, zero value otherwise.

### GetLatestExecutionResultOk

`func (o *Webhook) GetLatestExecutionResultOk() (*WebhookExecResult, bool)`

GetLatestExecutionResultOk returns a tuple with the LatestExecutionResult field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLatestExecutionResult

`func (o *Webhook) SetLatestExecutionResult(v WebhookExecResult)`

SetLatestExecutionResult sets LatestExecutionResult field to given value.

### HasLatestExecutionResult

`func (o *Webhook) HasLatestExecutionResult() bool`

HasLatestExecutionResult returns a boolean if a field has been set.

### GetModifiedAt

`func (o *Webhook) GetModifiedAt() string`

GetModifiedAt returns the ModifiedAt field if non-nil, zero value otherwise.

### GetModifiedAtOk

`func (o *Webhook) GetModifiedAtOk() (*string, bool)`

GetModifiedAtOk returns a tuple with the ModifiedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetModifiedAt

`func (o *Webhook) SetModifiedAt(v string)`

SetModifiedAt sets ModifiedAt field to given value.

### HasModifiedAt

`func (o *Webhook) HasModifiedAt() bool`

HasModifiedAt returns a boolean if a field has been set.

### GetName

`func (o *Webhook) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *Webhook) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *Webhook) SetName(v string)`

SetName sets Name field to given value.


### GetSecretIdentifier

`func (o *Webhook) GetSecretIdentifier() string`

GetSecretIdentifier returns the SecretIdentifier field if non-nil, zero value otherwise.

### GetSecretIdentifierOk

`func (o *Webhook) GetSecretIdentifierOk() (*string, bool)`

GetSecretIdentifierOk returns a tuple with the SecretIdentifier field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSecretIdentifier

`func (o *Webhook) SetSecretIdentifier(v string)`

SetSecretIdentifier sets SecretIdentifier field to given value.

### HasSecretIdentifier

`func (o *Webhook) HasSecretIdentifier() bool`

HasSecretIdentifier returns a boolean if a field has been set.

### GetSecretSpaceId

`func (o *Webhook) GetSecretSpaceId() int64`

GetSecretSpaceId returns the SecretSpaceId field if non-nil, zero value otherwise.

### GetSecretSpaceIdOk

`func (o *Webhook) GetSecretSpaceIdOk() (*int64, bool)`

GetSecretSpaceIdOk returns a tuple with the SecretSpaceId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSecretSpaceId

`func (o *Webhook) SetSecretSpaceId(v int64)`

SetSecretSpaceId sets SecretSpaceId field to given value.

### HasSecretSpaceId

`func (o *Webhook) HasSecretSpaceId() bool`

HasSecretSpaceId returns a boolean if a field has been set.

### GetSecretSpacePath

`func (o *Webhook) GetSecretSpacePath() string`

GetSecretSpacePath returns the SecretSpacePath field if non-nil, zero value otherwise.

### GetSecretSpacePathOk

`func (o *Webhook) GetSecretSpacePathOk() (*string, bool)`

GetSecretSpacePathOk returns a tuple with the SecretSpacePath field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSecretSpacePath

`func (o *Webhook) SetSecretSpacePath(v string)`

SetSecretSpacePath sets SecretSpacePath field to given value.

### HasSecretSpacePath

`func (o *Webhook) HasSecretSpacePath() bool`

HasSecretSpacePath returns a boolean if a field has been set.

### GetTriggers

`func (o *Webhook) GetTriggers() []Trigger`

GetTriggers returns the Triggers field if non-nil, zero value otherwise.

### GetTriggersOk

`func (o *Webhook) GetTriggersOk() (*[]Trigger, bool)`

GetTriggersOk returns a tuple with the Triggers field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTriggers

`func (o *Webhook) SetTriggers(v []Trigger)`

SetTriggers sets Triggers field to given value.

### HasTriggers

`func (o *Webhook) HasTriggers() bool`

HasTriggers returns a boolean if a field has been set.

### GetUrl

`func (o *Webhook) GetUrl() string`

GetUrl returns the Url field if non-nil, zero value otherwise.

### GetUrlOk

`func (o *Webhook) GetUrlOk() (*string, bool)`

GetUrlOk returns a tuple with the Url field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUrl

`func (o *Webhook) SetUrl(v string)`

SetUrl sets Url field to given value.


### GetVersion

`func (o *Webhook) GetVersion() int64`

GetVersion returns the Version field if non-nil, zero value otherwise.

### GetVersionOk

`func (o *Webhook) GetVersionOk() (*int64, bool)`

GetVersionOk returns a tuple with the Version field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVersion

`func (o *Webhook) SetVersion(v int64)`

SetVersion sets Version field to given value.

### HasVersion

`func (o *Webhook) HasVersion() bool`

HasVersion returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


