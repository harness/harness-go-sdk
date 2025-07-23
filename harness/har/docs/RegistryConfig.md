# RegistryConfig

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Type** | [**RegistryType**](RegistryType.md) |  | 
**UpstreamProxies** | Pointer to **[]string** |  | [optional] 
**Auth** | Pointer to [**NullableUpstreamConfigAuth**](UpstreamConfigAuth.md) |  | [optional] 
**AuthType** | [**AuthType**](AuthType.md) |  | 
**Source** | Pointer to **string** |  | [optional] 
**Url** | Pointer to **string** |  | [optional] 

## Methods

### NewRegistryConfig

`func NewRegistryConfig(type_ RegistryType, authType AuthType, ) *RegistryConfig`

NewRegistryConfig instantiates a new RegistryConfig object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRegistryConfigWithDefaults

`func NewRegistryConfigWithDefaults() *RegistryConfig`

NewRegistryConfigWithDefaults instantiates a new RegistryConfig object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetType

`func (o *RegistryConfig) GetType() RegistryType`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *RegistryConfig) GetTypeOk() (*RegistryType, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *RegistryConfig) SetType(v RegistryType)`

SetType sets Type field to given value.


### GetUpstreamProxies

`func (o *RegistryConfig) GetUpstreamProxies() []string`

GetUpstreamProxies returns the UpstreamProxies field if non-nil, zero value otherwise.

### GetUpstreamProxiesOk

`func (o *RegistryConfig) GetUpstreamProxiesOk() (*[]string, bool)`

GetUpstreamProxiesOk returns a tuple with the UpstreamProxies field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpstreamProxies

`func (o *RegistryConfig) SetUpstreamProxies(v []string)`

SetUpstreamProxies sets UpstreamProxies field to given value.

### HasUpstreamProxies

`func (o *RegistryConfig) HasUpstreamProxies() bool`

HasUpstreamProxies returns a boolean if a field has been set.

### GetAuth

`func (o *RegistryConfig) GetAuth() UpstreamConfigAuth`

GetAuth returns the Auth field if non-nil, zero value otherwise.

### GetAuthOk

`func (o *RegistryConfig) GetAuthOk() (*UpstreamConfigAuth, bool)`

GetAuthOk returns a tuple with the Auth field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAuth

`func (o *RegistryConfig) SetAuth(v UpstreamConfigAuth)`

SetAuth sets Auth field to given value.

### HasAuth

`func (o *RegistryConfig) HasAuth() bool`

HasAuth returns a boolean if a field has been set.

### SetAuthNil

`func (o *RegistryConfig) SetAuthNil(b bool)`

 SetAuthNil sets the value for Auth to be an explicit nil

### UnsetAuth
`func (o *RegistryConfig) UnsetAuth()`

UnsetAuth ensures that no value is present for Auth, not even an explicit nil
### GetAuthType

`func (o *RegistryConfig) GetAuthType() AuthType`

GetAuthType returns the AuthType field if non-nil, zero value otherwise.

### GetAuthTypeOk

`func (o *RegistryConfig) GetAuthTypeOk() (*AuthType, bool)`

GetAuthTypeOk returns a tuple with the AuthType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAuthType

`func (o *RegistryConfig) SetAuthType(v AuthType)`

SetAuthType sets AuthType field to given value.


### GetSource

`func (o *RegistryConfig) GetSource() string`

GetSource returns the Source field if non-nil, zero value otherwise.

### GetSourceOk

`func (o *RegistryConfig) GetSourceOk() (*string, bool)`

GetSourceOk returns a tuple with the Source field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSource

`func (o *RegistryConfig) SetSource(v string)`

SetSource sets Source field to given value.

### HasSource

`func (o *RegistryConfig) HasSource() bool`

HasSource returns a boolean if a field has been set.

### GetUrl

`func (o *RegistryConfig) GetUrl() string`

GetUrl returns the Url field if non-nil, zero value otherwise.

### GetUrlOk

`func (o *RegistryConfig) GetUrlOk() (*string, bool)`

GetUrlOk returns a tuple with the Url field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUrl

`func (o *RegistryConfig) SetUrl(v string)`

SetUrl sets Url field to given value.

### HasUrl

`func (o *RegistryConfig) HasUrl() bool`

HasUrl returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


