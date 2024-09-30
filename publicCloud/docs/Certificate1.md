# Certificate1

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**PrivateKey** | Pointer to **string** | Client Private Key. Required only if protocol is HTTPS | [optional] 
**Certificate** | Pointer to **string** | Client Certificate. Required only if protocol is HTTPS | [optional] 
**Chain** | Pointer to **NullableString** | CA certificate. Not required, but can be added if protocol is HTTPS | [optional] 

## Methods

### NewCertificate1

`func NewCertificate1() *Certificate1`

NewCertificate1 instantiates a new Certificate1 object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCertificate1WithDefaults

`func NewCertificate1WithDefaults() *Certificate1`

NewCertificate1WithDefaults instantiates a new Certificate1 object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetPrivateKey

`func (o *Certificate1) GetPrivateKey() string`

GetPrivateKey returns the PrivateKey field if non-nil, zero value otherwise.

### GetPrivateKeyOk

`func (o *Certificate1) GetPrivateKeyOk() (*string, bool)`

GetPrivateKeyOk returns a tuple with the PrivateKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrivateKey

`func (o *Certificate1) SetPrivateKey(v string)`

SetPrivateKey sets PrivateKey field to given value.

### HasPrivateKey

`func (o *Certificate1) HasPrivateKey() bool`

HasPrivateKey returns a boolean if a field has been set.

### GetCertificate

`func (o *Certificate1) GetCertificate() string`

GetCertificate returns the Certificate field if non-nil, zero value otherwise.

### GetCertificateOk

`func (o *Certificate1) GetCertificateOk() (*string, bool)`

GetCertificateOk returns a tuple with the Certificate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCertificate

`func (o *Certificate1) SetCertificate(v string)`

SetCertificate sets Certificate field to given value.

### HasCertificate

`func (o *Certificate1) HasCertificate() bool`

HasCertificate returns a boolean if a field has been set.

### GetChain

`func (o *Certificate1) GetChain() string`

GetChain returns the Chain field if non-nil, zero value otherwise.

### GetChainOk

`func (o *Certificate1) GetChainOk() (*string, bool)`

GetChainOk returns a tuple with the Chain field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetChain

`func (o *Certificate1) SetChain(v string)`

SetChain sets Chain field to given value.

### HasChain

`func (o *Certificate1) HasChain() bool`

HasChain returns a boolean if a field has been set.

### SetChainNil

`func (o *Certificate1) SetChainNil(b bool)`

 SetChainNil sets the value for Chain to be an explicit nil

### UnsetChain
`func (o *Certificate1) UnsetChain()`

UnsetChain ensures that no value is present for Chain, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


