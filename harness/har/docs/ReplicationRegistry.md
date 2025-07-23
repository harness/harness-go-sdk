# ReplicationRegistry

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**RegistryIdentifier** | **string** |  | 
**Namespace** | **string** |  | 
**PasswordSecretId** | Pointer to **string** |  | [optional] 
**PasswordSecretSpaceId** | Pointer to **string** |  | [optional] 
**Url** | **string** |  | 
**Username** | Pointer to **string** |  | [optional] 

## Methods

### NewReplicationRegistry

`func NewReplicationRegistry(registryIdentifier string, namespace string, url string, ) *ReplicationRegistry`

NewReplicationRegistry instantiates a new ReplicationRegistry object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewReplicationRegistryWithDefaults

`func NewReplicationRegistryWithDefaults() *ReplicationRegistry`

NewReplicationRegistryWithDefaults instantiates a new ReplicationRegistry object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetRegistryIdentifier

`func (o *ReplicationRegistry) GetRegistryIdentifier() string`

GetRegistryIdentifier returns the RegistryIdentifier field if non-nil, zero value otherwise.

### GetRegistryIdentifierOk

`func (o *ReplicationRegistry) GetRegistryIdentifierOk() (*string, bool)`

GetRegistryIdentifierOk returns a tuple with the RegistryIdentifier field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRegistryIdentifier

`func (o *ReplicationRegistry) SetRegistryIdentifier(v string)`

SetRegistryIdentifier sets RegistryIdentifier field to given value.


### GetNamespace

`func (o *ReplicationRegistry) GetNamespace() string`

GetNamespace returns the Namespace field if non-nil, zero value otherwise.

### GetNamespaceOk

`func (o *ReplicationRegistry) GetNamespaceOk() (*string, bool)`

GetNamespaceOk returns a tuple with the Namespace field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNamespace

`func (o *ReplicationRegistry) SetNamespace(v string)`

SetNamespace sets Namespace field to given value.


### GetPasswordSecretId

`func (o *ReplicationRegistry) GetPasswordSecretId() string`

GetPasswordSecretId returns the PasswordSecretId field if non-nil, zero value otherwise.

### GetPasswordSecretIdOk

`func (o *ReplicationRegistry) GetPasswordSecretIdOk() (*string, bool)`

GetPasswordSecretIdOk returns a tuple with the PasswordSecretId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPasswordSecretId

`func (o *ReplicationRegistry) SetPasswordSecretId(v string)`

SetPasswordSecretId sets PasswordSecretId field to given value.

### HasPasswordSecretId

`func (o *ReplicationRegistry) HasPasswordSecretId() bool`

HasPasswordSecretId returns a boolean if a field has been set.

### GetPasswordSecretSpaceId

`func (o *ReplicationRegistry) GetPasswordSecretSpaceId() string`

GetPasswordSecretSpaceId returns the PasswordSecretSpaceId field if non-nil, zero value otherwise.

### GetPasswordSecretSpaceIdOk

`func (o *ReplicationRegistry) GetPasswordSecretSpaceIdOk() (*string, bool)`

GetPasswordSecretSpaceIdOk returns a tuple with the PasswordSecretSpaceId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPasswordSecretSpaceId

`func (o *ReplicationRegistry) SetPasswordSecretSpaceId(v string)`

SetPasswordSecretSpaceId sets PasswordSecretSpaceId field to given value.

### HasPasswordSecretSpaceId

`func (o *ReplicationRegistry) HasPasswordSecretSpaceId() bool`

HasPasswordSecretSpaceId returns a boolean if a field has been set.

### GetUrl

`func (o *ReplicationRegistry) GetUrl() string`

GetUrl returns the Url field if non-nil, zero value otherwise.

### GetUrlOk

`func (o *ReplicationRegistry) GetUrlOk() (*string, bool)`

GetUrlOk returns a tuple with the Url field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUrl

`func (o *ReplicationRegistry) SetUrl(v string)`

SetUrl sets Url field to given value.


### GetUsername

`func (o *ReplicationRegistry) GetUsername() string`

GetUsername returns the Username field if non-nil, zero value otherwise.

### GetUsernameOk

`func (o *ReplicationRegistry) GetUsernameOk() (*string, bool)`

GetUsernameOk returns a tuple with the Username field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUsername

`func (o *ReplicationRegistry) SetUsername(v string)`

SetUsername sets Username field to given value.

### HasUsername

`func (o *ReplicationRegistry) HasUsername() bool`

HasUsername returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


