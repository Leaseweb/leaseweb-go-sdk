# CreateNetworkEquipmentCredentialOpts

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Password** | **string** | The password for the credentials | 
**Type** | [**CredentialType**](CredentialType.md) |  | 
**Username** | **string** | The username for the credentials | 

## Methods

### NewCreateNetworkEquipmentCredentialOpts

`func NewCreateNetworkEquipmentCredentialOpts(password string, type_ CredentialType, username string, ) *CreateNetworkEquipmentCredentialOpts`

NewCreateNetworkEquipmentCredentialOpts instantiates a new CreateNetworkEquipmentCredentialOpts object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCreateNetworkEquipmentCredentialOptsWithDefaults

`func NewCreateNetworkEquipmentCredentialOptsWithDefaults() *CreateNetworkEquipmentCredentialOpts`

NewCreateNetworkEquipmentCredentialOptsWithDefaults instantiates a new CreateNetworkEquipmentCredentialOpts object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetPassword

`func (o *CreateNetworkEquipmentCredentialOpts) GetPassword() string`

GetPassword returns the Password field if non-nil, zero value otherwise.

### GetPasswordOk

`func (o *CreateNetworkEquipmentCredentialOpts) GetPasswordOk() (*string, bool)`

GetPasswordOk returns a tuple with the Password field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPassword

`func (o *CreateNetworkEquipmentCredentialOpts) SetPassword(v string)`

SetPassword sets Password field to given value.


### GetType

`func (o *CreateNetworkEquipmentCredentialOpts) GetType() CredentialType`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *CreateNetworkEquipmentCredentialOpts) GetTypeOk() (*CredentialType, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *CreateNetworkEquipmentCredentialOpts) SetType(v CredentialType)`

SetType sets Type field to given value.


### GetUsername

`func (o *CreateNetworkEquipmentCredentialOpts) GetUsername() string`

GetUsername returns the Username field if non-nil, zero value otherwise.

### GetUsernameOk

`func (o *CreateNetworkEquipmentCredentialOpts) GetUsernameOk() (*string, bool)`

GetUsernameOk returns a tuple with the Username field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUsername

`func (o *CreateNetworkEquipmentCredentialOpts) SetUsername(v string)`

SetUsername sets Username field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


