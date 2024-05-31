# StoreCredentialOpts

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Type** | [**CredentialType**](CredentialType.md) |  | 
**Username** | **string** | Can contain only alphanumeric values and characters &#x60;@&#x60;, &#x60;.&#x60;, &#x60;-&#x60; and &#x60;_&#x60; | 
**Password** | **string** | The password you&#39;d like to store | 

## Methods

### NewStoreCredentialOpts

`func NewStoreCredentialOpts(type_ CredentialType, username string, password string, ) *StoreCredentialOpts`

NewStoreCredentialOpts instantiates a new StoreCredentialOpts object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewStoreCredentialOptsWithDefaults

`func NewStoreCredentialOptsWithDefaults() *StoreCredentialOpts`

NewStoreCredentialOptsWithDefaults instantiates a new StoreCredentialOpts object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetType

`func (o *StoreCredentialOpts) GetType() CredentialType`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *StoreCredentialOpts) GetTypeOk() (*CredentialType, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *StoreCredentialOpts) SetType(v CredentialType)`

SetType sets Type field to given value.


### GetUsername

`func (o *StoreCredentialOpts) GetUsername() string`

GetUsername returns the Username field if non-nil, zero value otherwise.

### GetUsernameOk

`func (o *StoreCredentialOpts) GetUsernameOk() (*string, bool)`

GetUsernameOk returns a tuple with the Username field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUsername

`func (o *StoreCredentialOpts) SetUsername(v string)`

SetUsername sets Username field to given value.


### GetPassword

`func (o *StoreCredentialOpts) GetPassword() string`

GetPassword returns the Password field if non-nil, zero value otherwise.

### GetPasswordOk

`func (o *StoreCredentialOpts) GetPasswordOk() (*string, bool)`

GetPasswordOk returns a tuple with the Password field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPassword

`func (o *StoreCredentialOpts) SetPassword(v string)`

SetPassword sets Password field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


