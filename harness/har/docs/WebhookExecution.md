# WebhookExecution

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Created** | Pointer to **int64** |  | [optional] 
**Duration** | Pointer to **int64** |  | [optional] 
**Error** | Pointer to **string** |  | [optional] 
**Id** | Pointer to **int64** |  | [optional] 
**Request** | Pointer to [**WebhookExecRequest**](WebhookExecRequest.md) |  | [optional] 
**Response** | Pointer to [**WebhookExecResponse**](WebhookExecResponse.md) |  | [optional] 
**Result** | Pointer to [**WebhookExecResult**](WebhookExecResult.md) |  | [optional] 
**RetriggerOf** | Pointer to **int64** |  | [optional] 
**Retriggerable** | Pointer to **bool** |  | [optional] 
**TriggerType** | Pointer to [**Trigger**](Trigger.md) |  | [optional] 
**WebhookId** | Pointer to **int64** |  | [optional] 

## Methods

### NewWebhookExecution

`func NewWebhookExecution() *WebhookExecution`

NewWebhookExecution instantiates a new WebhookExecution object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewWebhookExecutionWithDefaults

`func NewWebhookExecutionWithDefaults() *WebhookExecution`

NewWebhookExecutionWithDefaults instantiates a new WebhookExecution object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCreated

`func (o *WebhookExecution) GetCreated() int64`

GetCreated returns the Created field if non-nil, zero value otherwise.

### GetCreatedOk

`func (o *WebhookExecution) GetCreatedOk() (*int64, bool)`

GetCreatedOk returns a tuple with the Created field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreated

`func (o *WebhookExecution) SetCreated(v int64)`

SetCreated sets Created field to given value.

### HasCreated

`func (o *WebhookExecution) HasCreated() bool`

HasCreated returns a boolean if a field has been set.

### GetDuration

`func (o *WebhookExecution) GetDuration() int64`

GetDuration returns the Duration field if non-nil, zero value otherwise.

### GetDurationOk

`func (o *WebhookExecution) GetDurationOk() (*int64, bool)`

GetDurationOk returns a tuple with the Duration field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDuration

`func (o *WebhookExecution) SetDuration(v int64)`

SetDuration sets Duration field to given value.

### HasDuration

`func (o *WebhookExecution) HasDuration() bool`

HasDuration returns a boolean if a field has been set.

### GetError

`func (o *WebhookExecution) GetError() string`

GetError returns the Error field if non-nil, zero value otherwise.

### GetErrorOk

`func (o *WebhookExecution) GetErrorOk() (*string, bool)`

GetErrorOk returns a tuple with the Error field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetError

`func (o *WebhookExecution) SetError(v string)`

SetError sets Error field to given value.

### HasError

`func (o *WebhookExecution) HasError() bool`

HasError returns a boolean if a field has been set.

### GetId

`func (o *WebhookExecution) GetId() int64`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *WebhookExecution) GetIdOk() (*int64, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *WebhookExecution) SetId(v int64)`

SetId sets Id field to given value.

### HasId

`func (o *WebhookExecution) HasId() bool`

HasId returns a boolean if a field has been set.

### GetRequest

`func (o *WebhookExecution) GetRequest() WebhookExecRequest`

GetRequest returns the Request field if non-nil, zero value otherwise.

### GetRequestOk

`func (o *WebhookExecution) GetRequestOk() (*WebhookExecRequest, bool)`

GetRequestOk returns a tuple with the Request field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRequest

`func (o *WebhookExecution) SetRequest(v WebhookExecRequest)`

SetRequest sets Request field to given value.

### HasRequest

`func (o *WebhookExecution) HasRequest() bool`

HasRequest returns a boolean if a field has been set.

### GetResponse

`func (o *WebhookExecution) GetResponse() WebhookExecResponse`

GetResponse returns the Response field if non-nil, zero value otherwise.

### GetResponseOk

`func (o *WebhookExecution) GetResponseOk() (*WebhookExecResponse, bool)`

GetResponseOk returns a tuple with the Response field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResponse

`func (o *WebhookExecution) SetResponse(v WebhookExecResponse)`

SetResponse sets Response field to given value.

### HasResponse

`func (o *WebhookExecution) HasResponse() bool`

HasResponse returns a boolean if a field has been set.

### GetResult

`func (o *WebhookExecution) GetResult() WebhookExecResult`

GetResult returns the Result field if non-nil, zero value otherwise.

### GetResultOk

`func (o *WebhookExecution) GetResultOk() (*WebhookExecResult, bool)`

GetResultOk returns a tuple with the Result field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResult

`func (o *WebhookExecution) SetResult(v WebhookExecResult)`

SetResult sets Result field to given value.

### HasResult

`func (o *WebhookExecution) HasResult() bool`

HasResult returns a boolean if a field has been set.

### GetRetriggerOf

`func (o *WebhookExecution) GetRetriggerOf() int64`

GetRetriggerOf returns the RetriggerOf field if non-nil, zero value otherwise.

### GetRetriggerOfOk

`func (o *WebhookExecution) GetRetriggerOfOk() (*int64, bool)`

GetRetriggerOfOk returns a tuple with the RetriggerOf field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRetriggerOf

`func (o *WebhookExecution) SetRetriggerOf(v int64)`

SetRetriggerOf sets RetriggerOf field to given value.

### HasRetriggerOf

`func (o *WebhookExecution) HasRetriggerOf() bool`

HasRetriggerOf returns a boolean if a field has been set.

### GetRetriggerable

`func (o *WebhookExecution) GetRetriggerable() bool`

GetRetriggerable returns the Retriggerable field if non-nil, zero value otherwise.

### GetRetriggerableOk

`func (o *WebhookExecution) GetRetriggerableOk() (*bool, bool)`

GetRetriggerableOk returns a tuple with the Retriggerable field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRetriggerable

`func (o *WebhookExecution) SetRetriggerable(v bool)`

SetRetriggerable sets Retriggerable field to given value.

### HasRetriggerable

`func (o *WebhookExecution) HasRetriggerable() bool`

HasRetriggerable returns a boolean if a field has been set.

### GetTriggerType

`func (o *WebhookExecution) GetTriggerType() Trigger`

GetTriggerType returns the TriggerType field if non-nil, zero value otherwise.

### GetTriggerTypeOk

`func (o *WebhookExecution) GetTriggerTypeOk() (*Trigger, bool)`

GetTriggerTypeOk returns a tuple with the TriggerType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTriggerType

`func (o *WebhookExecution) SetTriggerType(v Trigger)`

SetTriggerType sets TriggerType field to given value.

### HasTriggerType

`func (o *WebhookExecution) HasTriggerType() bool`

HasTriggerType returns a boolean if a field has been set.

### GetWebhookId

`func (o *WebhookExecution) GetWebhookId() int64`

GetWebhookId returns the WebhookId field if non-nil, zero value otherwise.

### GetWebhookIdOk

`func (o *WebhookExecution) GetWebhookIdOk() (*int64, bool)`

GetWebhookIdOk returns a tuple with the WebhookId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWebhookId

`func (o *WebhookExecution) SetWebhookId(v int64)`

SetWebhookId sets WebhookId field to given value.

### HasWebhookId

`func (o *WebhookExecution) HasWebhookId() bool`

HasWebhookId returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


