# SslCertificate

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**PrivateKey** | **string** | Client Private Key. Required only if protocol is HTTPS | 
**Certificate** | **string** | Client Certificate. Required only if protocol is HTTPS | 
**Chain** | Pointer to **string** | CA certificate. Not required, but can be added if protocol is HTTPS | [optional] 

## Methods

### NewSslCertificate

`func NewSslCertificate(privateKey string, certificate string, ) *SslCertificate`

NewSslCertificate instantiates a new SslCertificate object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSslCertificateWithDefaults

`func NewSslCertificateWithDefaults() *SslCertificate`

NewSslCertificateWithDefaults instantiates a new SslCertificate object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetPrivateKey

`func (o *SslCertificate) GetPrivateKey() string`

GetPrivateKey returns the PrivateKey field if non-nil, zero value otherwise.

### GetPrivateKeyOk

`func (o *SslCertificate) GetPrivateKeyOk() (*string, bool)`

GetPrivateKeyOk returns a tuple with the PrivateKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrivateKey

`func (o *SslCertificate) SetPrivateKey(v string)`

SetPrivateKey sets PrivateKey field to given value.


### GetCertificate

`func (o *SslCertificate) GetCertificate() string`

GetCertificate returns the Certificate field if non-nil, zero value otherwise.

### GetCertificateOk

`func (o *SslCertificate) GetCertificateOk() (*string, bool)`

GetCertificateOk returns a tuple with the Certificate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCertificate

`func (o *SslCertificate) SetCertificate(v string)`

SetCertificate sets Certificate field to given value.


### GetChain

`func (o *SslCertificate) GetChain() string`

GetChain returns the Chain field if non-nil, zero value otherwise.

### GetChainOk

`func (o *SslCertificate) GetChainOk() (*string, bool)`

GetChainOk returns a tuple with the Chain field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetChain

`func (o *SslCertificate) SetChain(v string)`

SetChain sets Chain field to given value.

### HasChain

`func (o *SslCertificate) HasChain() bool`

HasChain returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


