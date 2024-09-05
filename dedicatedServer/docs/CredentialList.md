# CredentialList

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Metadata** | Pointer to [**Metadata**](Metadata.md) |  | [optional] 
**Credentials** | Pointer to [**[]CredentialWithoutPassword**](CredentialWithoutPassword.md) | An array of credentials | [optional] 

## Methods

### NewCredentialList

`func NewCredentialList() *CredentialList`

NewCredentialList instantiates a new CredentialList object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCredentialListWithDefaults

`func NewCredentialListWithDefaults() *CredentialList`

NewCredentialListWithDefaults instantiates a new CredentialList object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetMetadata

`func (o *CredentialList) GetMetadata() Metadata`

GetMetadata returns the Metadata field if non-nil, zero value otherwise.

### GetMetadataOk

`func (o *CredentialList) GetMetadataOk() (*Metadata, bool)`

GetMetadataOk returns a tuple with the Metadata field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetadata

`func (o *CredentialList) SetMetadata(v Metadata)`

SetMetadata sets Metadata field to given value.

### HasMetadata

`func (o *CredentialList) HasMetadata() bool`

HasMetadata returns a boolean if a field has been set.

### GetCredentials

`func (o *CredentialList) GetCredentials() []CredentialWithoutPassword`

GetCredentials returns the Credentials field if non-nil, zero value otherwise.

### GetCredentialsOk

`func (o *CredentialList) GetCredentialsOk() (*[]CredentialWithoutPassword, bool)`

GetCredentialsOk returns a tuple with the Credentials field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCredentials

`func (o *CredentialList) SetCredentials(v []CredentialWithoutPassword)`

SetCredentials sets Credentials field to given value.

### HasCredentials

`func (o *CredentialList) HasCredentials() bool`

HasCredentials returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


