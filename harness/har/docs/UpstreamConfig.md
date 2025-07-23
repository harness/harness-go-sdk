# UpstreamConfig

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Auth** | Pointer to [**NullableUpstreamConfigAuth**](UpstreamConfigAuth.md) |  | [optional] 
**AuthType** | [**AuthType**](AuthType.md) |  | 
**Source** | Pointer to **string** |  | [optional] 
**Url** | Pointer to **string** |  | [optional] 

## Methods

### NewUpstreamConfig

`func NewUpstreamConfig(authType AuthType, ) *UpstreamConfig`

NewUpstreamConfig instantiates a new UpstreamConfig object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUpstreamConfigWithDefaults

`func NewUpstreamConfigWithDefaults() *UpstreamConfig`

NewUpstreamConfigWithDefaults instantiates a new UpstreamConfig object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAuth

`func (o *UpstreamConfig) GetAuth() UpstreamConfigAuth`

GetAuth returns the Auth field if non-nil, zero value otherwise.

### GetAuthOk

`func (o *UpstreamConfig) GetAuthOk() (*UpstreamConfigAuth, bool)`

GetAuthOk returns a tuple with the Auth field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAuth

`func (o *UpstreamConfig) SetAuth(v UpstreamConfigAuth)`

SetAuth sets Auth field to given value.

### HasAuth

`func (o *UpstreamConfig) HasAuth() bool`

HasAuth returns a boolean if a field has been set.

### SetAuthNil

`func (o *UpstreamConfig) SetAuthNil(b bool)`

 SetAuthNil sets the value for Auth to be an explicit nil

### UnsetAuth
`func (o *UpstreamConfig) UnsetAuth()`

UnsetAuth ensures that no value is present for Auth, not even an explicit nil
### GetAuthType

`func (o *UpstreamConfig) GetAuthType() AuthType`

GetAuthType returns the AuthType field if non-nil, zero value otherwise.

### GetAuthTypeOk

`func (o *UpstreamConfig) GetAuthTypeOk() (*AuthType, bool)`

GetAuthTypeOk returns a tuple with the AuthType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAuthType

`func (o *UpstreamConfig) SetAuthType(v AuthType)`

SetAuthType sets AuthType field to given value.


### GetSource

`func (o *UpstreamConfig) GetSource() string`

GetSource returns the Source field if non-nil, zero value otherwise.

### GetSourceOk

`func (o *UpstreamConfig) GetSourceOk() (*string, bool)`

GetSourceOk returns a tuple with the Source field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSource

`func (o *UpstreamConfig) SetSource(v string)`

SetSource sets Source field to given value.

### HasSource

`func (o *UpstreamConfig) HasSource() bool`

HasSource returns a boolean if a field has been set.

### GetUrl

`func (o *UpstreamConfig) GetUrl() string`

GetUrl returns the Url field if non-nil, zero value otherwise.

### GetUrlOk

`func (o *UpstreamConfig) GetUrlOk() (*string, bool)`

GetUrlOk returns a tuple with the Url field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUrl

`func (o *UpstreamConfig) SetUrl(v string)`

SetUrl sets Url field to given value.

### HasUrl

`func (o *UpstreamConfig) HasUrl() bool`

HasUrl returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


