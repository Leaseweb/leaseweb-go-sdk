# GetCredentialListByTypeResult

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Credentials** | Pointer to [**[]Credential**](Credential.md) |  | [optional] 
**Metadata** | Pointer to [**Metadata**](Metadata.md) |  | [optional] 

## Methods

### NewGetCredentialListByTypeResult

`func NewGetCredentialListByTypeResult() *GetCredentialListByTypeResult`

NewGetCredentialListByTypeResult instantiates a new GetCredentialListByTypeResult object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetCredentialListByTypeResultWithDefaults

`func NewGetCredentialListByTypeResultWithDefaults() *GetCredentialListByTypeResult`

NewGetCredentialListByTypeResultWithDefaults instantiates a new GetCredentialListByTypeResult object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCredentials

`func (o *GetCredentialListByTypeResult) GetCredentials() []Credential`

GetCredentials returns the Credentials field if non-nil, zero value otherwise.

### GetCredentialsOk

`func (o *GetCredentialListByTypeResult) GetCredentialsOk() (*[]Credential, bool)`

GetCredentialsOk returns a tuple with the Credentials field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCredentials

`func (o *GetCredentialListByTypeResult) SetCredentials(v []Credential)`

SetCredentials sets Credentials field to given value.

### HasCredentials

`func (o *GetCredentialListByTypeResult) HasCredentials() bool`

HasCredentials returns a boolean if a field has been set.

### GetMetadata

`func (o *GetCredentialListByTypeResult) GetMetadata() Metadata`

GetMetadata returns the Metadata field if non-nil, zero value otherwise.

### GetMetadataOk

`func (o *GetCredentialListByTypeResult) GetMetadataOk() (*Metadata, bool)`

GetMetadataOk returns a tuple with the Metadata field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetadata

`func (o *GetCredentialListByTypeResult) SetMetadata(v Metadata)`

SetMetadata sets Metadata field to given value.

### HasMetadata

`func (o *GetCredentialListByTypeResult) HasMetadata() bool`

HasMetadata returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


