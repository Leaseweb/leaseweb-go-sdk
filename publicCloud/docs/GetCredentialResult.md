# GetCredentialResult

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Type** | Pointer to [**CredentialType**](CredentialType.md) |  | [optional] 
**Username** | Pointer to **string** |  | [optional] 
**Password** | Pointer to **string** |  | [optional] 

## Methods

### NewGetCredentialResult

`func NewGetCredentialResult() *GetCredentialResult`

NewGetCredentialResult instantiates a new GetCredentialResult object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetCredentialResultWithDefaults

`func NewGetCredentialResultWithDefaults() *GetCredentialResult`

NewGetCredentialResultWithDefaults instantiates a new GetCredentialResult object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetType

`func (o *GetCredentialResult) GetType() CredentialType`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *GetCredentialResult) GetTypeOk() (*CredentialType, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *GetCredentialResult) SetType(v CredentialType)`

SetType sets Type field to given value.

### HasType

`func (o *GetCredentialResult) HasType() bool`

HasType returns a boolean if a field has been set.

### GetUsername

`func (o *GetCredentialResult) GetUsername() string`

GetUsername returns the Username field if non-nil, zero value otherwise.

### GetUsernameOk

`func (o *GetCredentialResult) GetUsernameOk() (*string, bool)`

GetUsernameOk returns a tuple with the Username field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUsername

`func (o *GetCredentialResult) SetUsername(v string)`

SetUsername sets Username field to given value.

### HasUsername

`func (o *GetCredentialResult) HasUsername() bool`

HasUsername returns a boolean if a field has been set.

### GetPassword

`func (o *GetCredentialResult) GetPassword() string`

GetPassword returns the Password field if non-nil, zero value otherwise.

### GetPasswordOk

`func (o *GetCredentialResult) GetPasswordOk() (*string, bool)`

GetPasswordOk returns a tuple with the Password field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPassword

`func (o *GetCredentialResult) SetPassword(v string)`

SetPassword sets Password field to given value.

### HasPassword

`func (o *GetCredentialResult) HasPassword() bool`

HasPassword returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


