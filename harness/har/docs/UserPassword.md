# UserPassword

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**SecretIdentifier** | Pointer to **string** |  | [optional] 
**SecretSpaceId** | Pointer to **int32** |  | [optional] 
**SecretSpacePath** | Pointer to **string** |  | [optional] 
**UserName** | **string** |  | 

## Methods

### NewUserPassword

`func NewUserPassword(userName string, ) *UserPassword`

NewUserPassword instantiates a new UserPassword object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUserPasswordWithDefaults

`func NewUserPasswordWithDefaults() *UserPassword`

NewUserPasswordWithDefaults instantiates a new UserPassword object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSecretIdentifier

`func (o *UserPassword) GetSecretIdentifier() string`

GetSecretIdentifier returns the SecretIdentifier field if non-nil, zero value otherwise.

### GetSecretIdentifierOk

`func (o *UserPassword) GetSecretIdentifierOk() (*string, bool)`

GetSecretIdentifierOk returns a tuple with the SecretIdentifier field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSecretIdentifier

`func (o *UserPassword) SetSecretIdentifier(v string)`

SetSecretIdentifier sets SecretIdentifier field to given value.

### HasSecretIdentifier

`func (o *UserPassword) HasSecretIdentifier() bool`

HasSecretIdentifier returns a boolean if a field has been set.

### GetSecretSpaceId

`func (o *UserPassword) GetSecretSpaceId() int32`

GetSecretSpaceId returns the SecretSpaceId field if non-nil, zero value otherwise.

### GetSecretSpaceIdOk

`func (o *UserPassword) GetSecretSpaceIdOk() (*int32, bool)`

GetSecretSpaceIdOk returns a tuple with the SecretSpaceId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSecretSpaceId

`func (o *UserPassword) SetSecretSpaceId(v int32)`

SetSecretSpaceId sets SecretSpaceId field to given value.

### HasSecretSpaceId

`func (o *UserPassword) HasSecretSpaceId() bool`

HasSecretSpaceId returns a boolean if a field has been set.

### GetSecretSpacePath

`func (o *UserPassword) GetSecretSpacePath() string`

GetSecretSpacePath returns the SecretSpacePath field if non-nil, zero value otherwise.

### GetSecretSpacePathOk

`func (o *UserPassword) GetSecretSpacePathOk() (*string, bool)`

GetSecretSpacePathOk returns a tuple with the SecretSpacePath field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSecretSpacePath

`func (o *UserPassword) SetSecretSpacePath(v string)`

SetSecretSpacePath sets SecretSpacePath field to given value.

### HasSecretSpacePath

`func (o *UserPassword) HasSecretSpacePath() bool`

HasSecretSpacePath returns a boolean if a field has been set.

### GetUserName

`func (o *UserPassword) GetUserName() string`

GetUserName returns the UserName field if non-nil, zero value otherwise.

### GetUserNameOk

`func (o *UserPassword) GetUserNameOk() (*string, bool)`

GetUserNameOk returns a tuple with the UserName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserName

`func (o *UserPassword) SetUserName(v string)`

SetUserName sets UserName field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


