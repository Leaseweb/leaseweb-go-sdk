# StoreCredentialResult

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Type** | Pointer to [**CredentialType**](CredentialType.md) |  | [optional] 
**Username** | Pointer to **string** | The provided username | [optional] 
**Password** | Pointer to **string** | The provided password | [optional] 

## Methods

### NewStoreCredentialResult

`func NewStoreCredentialResult() *StoreCredentialResult`

NewStoreCredentialResult instantiates a new StoreCredentialResult object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewStoreCredentialResultWithDefaults

`func NewStoreCredentialResultWithDefaults() *StoreCredentialResult`

NewStoreCredentialResultWithDefaults instantiates a new StoreCredentialResult object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetType

`func (o *StoreCredentialResult) GetType() CredentialType`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *StoreCredentialResult) GetTypeOk() (*CredentialType, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *StoreCredentialResult) SetType(v CredentialType)`

SetType sets Type field to given value.

### HasType

`func (o *StoreCredentialResult) HasType() bool`

HasType returns a boolean if a field has been set.

### GetUsername

`func (o *StoreCredentialResult) GetUsername() string`

GetUsername returns the Username field if non-nil, zero value otherwise.

### GetUsernameOk

`func (o *StoreCredentialResult) GetUsernameOk() (*string, bool)`

GetUsernameOk returns a tuple with the Username field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUsername

`func (o *StoreCredentialResult) SetUsername(v string)`

SetUsername sets Username field to given value.

### HasUsername

`func (o *StoreCredentialResult) HasUsername() bool`

HasUsername returns a boolean if a field has been set.

### GetPassword

`func (o *StoreCredentialResult) GetPassword() string`

GetPassword returns the Password field if non-nil, zero value otherwise.

### GetPasswordOk

`func (o *StoreCredentialResult) GetPasswordOk() (*string, bool)`

GetPasswordOk returns a tuple with the Password field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPassword

`func (o *StoreCredentialResult) SetPassword(v string)`

SetPassword sets Password field to given value.

### HasPassword

`func (o *StoreCredentialResult) HasPassword() bool`

HasPassword returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


