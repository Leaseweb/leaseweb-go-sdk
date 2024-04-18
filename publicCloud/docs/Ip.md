# Ip

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Ip** | Pointer to **string** | Ip Address | [optional] 
**PrefixLength** | Pointer to **string** | The number of leading bits in the IP address | [optional] 
**Version** | Pointer to **int32** | Ip version | [optional] 
**NullRouted** | Pointer to **bool** | Whether or not the IP has been nulled | [optional] 
**MainIp** | Pointer to **bool** |  | [optional] 
**ReverseLookup** | Pointer to **NullableString** |  | [optional] 
**Ddos** | Pointer to [**Ddos**](Ddos.md) |  | [optional] 

## Methods

### NewIp

`func NewIp() *Ip`

NewIp instantiates a new Ip object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewIpWithDefaults

`func NewIpWithDefaults() *Ip`

NewIpWithDefaults instantiates a new Ip object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetIp

`func (o *Ip) GetIp() string`

GetIp returns the Ip field if non-nil, zero value otherwise.

### GetIpOk

`func (o *Ip) GetIpOk() (*string, bool)`

GetIpOk returns a tuple with the Ip field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIp

`func (o *Ip) SetIp(v string)`

SetIp sets Ip field to given value.

### HasIp

`func (o *Ip) HasIp() bool`

HasIp returns a boolean if a field has been set.

### GetPrefixLength

`func (o *Ip) GetPrefixLength() string`

GetPrefixLength returns the PrefixLength field if non-nil, zero value otherwise.

### GetPrefixLengthOk

`func (o *Ip) GetPrefixLengthOk() (*string, bool)`

GetPrefixLengthOk returns a tuple with the PrefixLength field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrefixLength

`func (o *Ip) SetPrefixLength(v string)`

SetPrefixLength sets PrefixLength field to given value.

### HasPrefixLength

`func (o *Ip) HasPrefixLength() bool`

HasPrefixLength returns a boolean if a field has been set.

### GetVersion

`func (o *Ip) GetVersion() int32`

GetVersion returns the Version field if non-nil, zero value otherwise.

### GetVersionOk

`func (o *Ip) GetVersionOk() (*int32, bool)`

GetVersionOk returns a tuple with the Version field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVersion

`func (o *Ip) SetVersion(v int32)`

SetVersion sets Version field to given value.

### HasVersion

`func (o *Ip) HasVersion() bool`

HasVersion returns a boolean if a field has been set.

### GetNullRouted

`func (o *Ip) GetNullRouted() bool`

GetNullRouted returns the NullRouted field if non-nil, zero value otherwise.

### GetNullRoutedOk

`func (o *Ip) GetNullRoutedOk() (*bool, bool)`

GetNullRoutedOk returns a tuple with the NullRouted field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNullRouted

`func (o *Ip) SetNullRouted(v bool)`

SetNullRouted sets NullRouted field to given value.

### HasNullRouted

`func (o *Ip) HasNullRouted() bool`

HasNullRouted returns a boolean if a field has been set.

### GetMainIp

`func (o *Ip) GetMainIp() bool`

GetMainIp returns the MainIp field if non-nil, zero value otherwise.

### GetMainIpOk

`func (o *Ip) GetMainIpOk() (*bool, bool)`

GetMainIpOk returns a tuple with the MainIp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMainIp

`func (o *Ip) SetMainIp(v bool)`

SetMainIp sets MainIp field to given value.

### HasMainIp

`func (o *Ip) HasMainIp() bool`

HasMainIp returns a boolean if a field has been set.

### GetReverseLookup

`func (o *Ip) GetReverseLookup() string`

GetReverseLookup returns the ReverseLookup field if non-nil, zero value otherwise.

### GetReverseLookupOk

`func (o *Ip) GetReverseLookupOk() (*string, bool)`

GetReverseLookupOk returns a tuple with the ReverseLookup field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReverseLookup

`func (o *Ip) SetReverseLookup(v string)`

SetReverseLookup sets ReverseLookup field to given value.

### HasReverseLookup

`func (o *Ip) HasReverseLookup() bool`

HasReverseLookup returns a boolean if a field has been set.

### SetReverseLookupNil

`func (o *Ip) SetReverseLookupNil(b bool)`

 SetReverseLookupNil sets the value for ReverseLookup to be an explicit nil

### UnsetReverseLookup
`func (o *Ip) UnsetReverseLookup()`

UnsetReverseLookup ensures that no value is present for ReverseLookup, not even an explicit nil
### GetDdos

`func (o *Ip) GetDdos() Ddos`

GetDdos returns the Ddos field if non-nil, zero value otherwise.

### GetDdosOk

`func (o *Ip) GetDdosOk() (*Ddos, bool)`

GetDdosOk returns a tuple with the Ddos field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDdos

`func (o *Ip) SetDdos(v Ddos)`

SetDdos sets Ddos field to given value.

### HasDdos

`func (o *Ip) HasDdos() bool`

HasDdos returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


