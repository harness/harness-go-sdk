# UpstreamConfigAuth

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**SecretIdentifier** | Pointer to **string** |  | [optional] 
**SecretSpaceId** | Pointer to **int32** |  | [optional] 
**SecretSpacePath** | Pointer to **string** |  | [optional] 
**UserName** | **string** |  | 
**AccessKey** | Pointer to **string** |  | [optional] 
**AccessKeySecretIdentifier** | Pointer to **string** |  | [optional] 
**AccessKeySecretSpaceId** | Pointer to **int64** |  | [optional] 
**AccessKeySecretSpacePath** | Pointer to **string** |  | [optional] 
**SecretKeyIdentifier** | **string** |  | 
**SecretKeySpaceId** | Pointer to **int64** |  | [optional] 
**SecretKeySpacePath** | Pointer to **string** |  | [optional] 

## Methods

### NewUpstreamConfigAuth

`func NewUpstreamConfigAuth(userName string, secretKeyIdentifier string, ) *UpstreamConfigAuth`

NewUpstreamConfigAuth instantiates a new UpstreamConfigAuth object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUpstreamConfigAuthWithDefaults

`func NewUpstreamConfigAuthWithDefaults() *UpstreamConfigAuth`

NewUpstreamConfigAuthWithDefaults instantiates a new UpstreamConfigAuth object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSecretIdentifier

`func (o *UpstreamConfigAuth) GetSecretIdentifier() string`

GetSecretIdentifier returns the SecretIdentifier field if non-nil, zero value otherwise.

### GetSecretIdentifierOk

`func (o *UpstreamConfigAuth) GetSecretIdentifierOk() (*string, bool)`

GetSecretIdentifierOk returns a tuple with the SecretIdentifier field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSecretIdentifier

`func (o *UpstreamConfigAuth) SetSecretIdentifier(v string)`

SetSecretIdentifier sets SecretIdentifier field to given value.

### HasSecretIdentifier

`func (o *UpstreamConfigAuth) HasSecretIdentifier() bool`

HasSecretIdentifier returns a boolean if a field has been set.

### GetSecretSpaceId

`func (o *UpstreamConfigAuth) GetSecretSpaceId() int32`

GetSecretSpaceId returns the SecretSpaceId field if non-nil, zero value otherwise.

### GetSecretSpaceIdOk

`func (o *UpstreamConfigAuth) GetSecretSpaceIdOk() (*int32, bool)`

GetSecretSpaceIdOk returns a tuple with the SecretSpaceId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSecretSpaceId

`func (o *UpstreamConfigAuth) SetSecretSpaceId(v int32)`

SetSecretSpaceId sets SecretSpaceId field to given value.

### HasSecretSpaceId

`func (o *UpstreamConfigAuth) HasSecretSpaceId() bool`

HasSecretSpaceId returns a boolean if a field has been set.

### GetSecretSpacePath

`func (o *UpstreamConfigAuth) GetSecretSpacePath() string`

GetSecretSpacePath returns the SecretSpacePath field if non-nil, zero value otherwise.

### GetSecretSpacePathOk

`func (o *UpstreamConfigAuth) GetSecretSpacePathOk() (*string, bool)`

GetSecretSpacePathOk returns a tuple with the SecretSpacePath field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSecretSpacePath

`func (o *UpstreamConfigAuth) SetSecretSpacePath(v string)`

SetSecretSpacePath sets SecretSpacePath field to given value.

### HasSecretSpacePath

`func (o *UpstreamConfigAuth) HasSecretSpacePath() bool`

HasSecretSpacePath returns a boolean if a field has been set.

### GetUserName

`func (o *UpstreamConfigAuth) GetUserName() string`

GetUserName returns the UserName field if non-nil, zero value otherwise.

### GetUserNameOk

`func (o *UpstreamConfigAuth) GetUserNameOk() (*string, bool)`

GetUserNameOk returns a tuple with the UserName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserName

`func (o *UpstreamConfigAuth) SetUserName(v string)`

SetUserName sets UserName field to given value.


### GetAccessKey

`func (o *UpstreamConfigAuth) GetAccessKey() string`

GetAccessKey returns the AccessKey field if non-nil, zero value otherwise.

### GetAccessKeyOk

`func (o *UpstreamConfigAuth) GetAccessKeyOk() (*string, bool)`

GetAccessKeyOk returns a tuple with the AccessKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccessKey

`func (o *UpstreamConfigAuth) SetAccessKey(v string)`

SetAccessKey sets AccessKey field to given value.

### HasAccessKey

`func (o *UpstreamConfigAuth) HasAccessKey() bool`

HasAccessKey returns a boolean if a field has been set.

### GetAccessKeySecretIdentifier

`func (o *UpstreamConfigAuth) GetAccessKeySecretIdentifier() string`

GetAccessKeySecretIdentifier returns the AccessKeySecretIdentifier field if non-nil, zero value otherwise.

### GetAccessKeySecretIdentifierOk

`func (o *UpstreamConfigAuth) GetAccessKeySecretIdentifierOk() (*string, bool)`

GetAccessKeySecretIdentifierOk returns a tuple with the AccessKeySecretIdentifier field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccessKeySecretIdentifier

`func (o *UpstreamConfigAuth) SetAccessKeySecretIdentifier(v string)`

SetAccessKeySecretIdentifier sets AccessKeySecretIdentifier field to given value.

### HasAccessKeySecretIdentifier

`func (o *UpstreamConfigAuth) HasAccessKeySecretIdentifier() bool`

HasAccessKeySecretIdentifier returns a boolean if a field has been set.

### GetAccessKeySecretSpaceId

`func (o *UpstreamConfigAuth) GetAccessKeySecretSpaceId() int64`

GetAccessKeySecretSpaceId returns the AccessKeySecretSpaceId field if non-nil, zero value otherwise.

### GetAccessKeySecretSpaceIdOk

`func (o *UpstreamConfigAuth) GetAccessKeySecretSpaceIdOk() (*int64, bool)`

GetAccessKeySecretSpaceIdOk returns a tuple with the AccessKeySecretSpaceId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccessKeySecretSpaceId

`func (o *UpstreamConfigAuth) SetAccessKeySecretSpaceId(v int64)`

SetAccessKeySecretSpaceId sets AccessKeySecretSpaceId field to given value.

### HasAccessKeySecretSpaceId

`func (o *UpstreamConfigAuth) HasAccessKeySecretSpaceId() bool`

HasAccessKeySecretSpaceId returns a boolean if a field has been set.

### GetAccessKeySecretSpacePath

`func (o *UpstreamConfigAuth) GetAccessKeySecretSpacePath() string`

GetAccessKeySecretSpacePath returns the AccessKeySecretSpacePath field if non-nil, zero value otherwise.

### GetAccessKeySecretSpacePathOk

`func (o *UpstreamConfigAuth) GetAccessKeySecretSpacePathOk() (*string, bool)`

GetAccessKeySecretSpacePathOk returns a tuple with the AccessKeySecretSpacePath field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccessKeySecretSpacePath

`func (o *UpstreamConfigAuth) SetAccessKeySecretSpacePath(v string)`

SetAccessKeySecretSpacePath sets AccessKeySecretSpacePath field to given value.

### HasAccessKeySecretSpacePath

`func (o *UpstreamConfigAuth) HasAccessKeySecretSpacePath() bool`

HasAccessKeySecretSpacePath returns a boolean if a field has been set.

### GetSecretKeyIdentifier

`func (o *UpstreamConfigAuth) GetSecretKeyIdentifier() string`

GetSecretKeyIdentifier returns the SecretKeyIdentifier field if non-nil, zero value otherwise.

### GetSecretKeyIdentifierOk

`func (o *UpstreamConfigAuth) GetSecretKeyIdentifierOk() (*string, bool)`

GetSecretKeyIdentifierOk returns a tuple with the SecretKeyIdentifier field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSecretKeyIdentifier

`func (o *UpstreamConfigAuth) SetSecretKeyIdentifier(v string)`

SetSecretKeyIdentifier sets SecretKeyIdentifier field to given value.


### GetSecretKeySpaceId

`func (o *UpstreamConfigAuth) GetSecretKeySpaceId() int64`

GetSecretKeySpaceId returns the SecretKeySpaceId field if non-nil, zero value otherwise.

### GetSecretKeySpaceIdOk

`func (o *UpstreamConfigAuth) GetSecretKeySpaceIdOk() (*int64, bool)`

GetSecretKeySpaceIdOk returns a tuple with the SecretKeySpaceId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSecretKeySpaceId

`func (o *UpstreamConfigAuth) SetSecretKeySpaceId(v int64)`

SetSecretKeySpaceId sets SecretKeySpaceId field to given value.

### HasSecretKeySpaceId

`func (o *UpstreamConfigAuth) HasSecretKeySpaceId() bool`

HasSecretKeySpaceId returns a boolean if a field has been set.

### GetSecretKeySpacePath

`func (o *UpstreamConfigAuth) GetSecretKeySpacePath() string`

GetSecretKeySpacePath returns the SecretKeySpacePath field if non-nil, zero value otherwise.

### GetSecretKeySpacePathOk

`func (o *UpstreamConfigAuth) GetSecretKeySpacePathOk() (*string, bool)`

GetSecretKeySpacePathOk returns a tuple with the SecretKeySpacePath field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSecretKeySpacePath

`func (o *UpstreamConfigAuth) SetSecretKeySpacePath(v string)`

SetSecretKeySpacePath sets SecretKeySpacePath field to given value.

### HasSecretKeySpacePath

`func (o *UpstreamConfigAuth) HasSecretKeySpacePath() bool`

HasSecretKeySpacePath returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


