# IpDetails

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Ip** | **string** | Ip Address | 
**PrefixLength** | **string** | The number of leading bits in the IP address | 
**Version** | [**IpVersion**](IpVersion.md) |  | 
**NullRouted** | **bool** | Whether or not the IP has been nulled | 
**MainIp** | **bool** |  | 
**NetworkType** | [**NetworkType**](NetworkType.md) |  | 
**ReverseLookup** | **NullableString** |  | 
**Ddos** | [**NullableDdos**](Ddos.md) |  | 

## Methods

### NewIpDetails

`func NewIpDetails(ip string, prefixLength string, version IpVersion, nullRouted bool, mainIp bool, networkType NetworkType, reverseLookup NullableString, ddos NullableDdos, ) *IpDetails`

NewIpDetails instantiates a new IpDetails object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewIpDetailsWithDefaults

`func NewIpDetailsWithDefaults() *IpDetails`

NewIpDetailsWithDefaults instantiates a new IpDetails object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetIp

`func (o *IpDetails) GetIp() string`

GetIp returns the Ip field if non-nil, zero value otherwise.

### GetIpOk

`func (o *IpDetails) GetIpOk() (*string, bool)`

GetIpOk returns a tuple with the Ip field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIp

`func (o *IpDetails) SetIp(v string)`

SetIp sets Ip field to given value.


### GetPrefixLength

`func (o *IpDetails) GetPrefixLength() string`

GetPrefixLength returns the PrefixLength field if non-nil, zero value otherwise.

### GetPrefixLengthOk

`func (o *IpDetails) GetPrefixLengthOk() (*string, bool)`

GetPrefixLengthOk returns a tuple with the PrefixLength field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrefixLength

`func (o *IpDetails) SetPrefixLength(v string)`

SetPrefixLength sets PrefixLength field to given value.


### GetVersion

`func (o *IpDetails) GetVersion() IpVersion`

GetVersion returns the Version field if non-nil, zero value otherwise.

### GetVersionOk

`func (o *IpDetails) GetVersionOk() (*IpVersion, bool)`

GetVersionOk returns a tuple with the Version field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVersion

`func (o *IpDetails) SetVersion(v IpVersion)`

SetVersion sets Version field to given value.


### GetNullRouted

`func (o *IpDetails) GetNullRouted() bool`

GetNullRouted returns the NullRouted field if non-nil, zero value otherwise.

### GetNullRoutedOk

`func (o *IpDetails) GetNullRoutedOk() (*bool, bool)`

GetNullRoutedOk returns a tuple with the NullRouted field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNullRouted

`func (o *IpDetails) SetNullRouted(v bool)`

SetNullRouted sets NullRouted field to given value.


### GetMainIp

`func (o *IpDetails) GetMainIp() bool`

GetMainIp returns the MainIp field if non-nil, zero value otherwise.

### GetMainIpOk

`func (o *IpDetails) GetMainIpOk() (*bool, bool)`

GetMainIpOk returns a tuple with the MainIp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMainIp

`func (o *IpDetails) SetMainIp(v bool)`

SetMainIp sets MainIp field to given value.


### GetNetworkType

`func (o *IpDetails) GetNetworkType() NetworkType`

GetNetworkType returns the NetworkType field if non-nil, zero value otherwise.

### GetNetworkTypeOk

`func (o *IpDetails) GetNetworkTypeOk() (*NetworkType, bool)`

GetNetworkTypeOk returns a tuple with the NetworkType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNetworkType

`func (o *IpDetails) SetNetworkType(v NetworkType)`

SetNetworkType sets NetworkType field to given value.


### GetReverseLookup

`func (o *IpDetails) GetReverseLookup() string`

GetReverseLookup returns the ReverseLookup field if non-nil, zero value otherwise.

### GetReverseLookupOk

`func (o *IpDetails) GetReverseLookupOk() (*string, bool)`

GetReverseLookupOk returns a tuple with the ReverseLookup field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReverseLookup

`func (o *IpDetails) SetReverseLookup(v string)`

SetReverseLookup sets ReverseLookup field to given value.


### SetReverseLookupNil

`func (o *IpDetails) SetReverseLookupNil(b bool)`

 SetReverseLookupNil sets the value for ReverseLookup to be an explicit nil

### UnsetReverseLookup
`func (o *IpDetails) UnsetReverseLookup()`

UnsetReverseLookup ensures that no value is present for ReverseLookup, not even an explicit nil
### GetDdos

`func (o *IpDetails) GetDdos() Ddos`

GetDdos returns the Ddos field if non-nil, zero value otherwise.

### GetDdosOk

`func (o *IpDetails) GetDdosOk() (*Ddos, bool)`

GetDdosOk returns a tuple with the Ddos field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDdos

`func (o *IpDetails) SetDdos(v Ddos)`

SetDdos sets Ddos field to given value.


### SetDdosNil

`func (o *IpDetails) SetDdosNil(b bool)`

 SetDdosNil sets the value for Ddos to be an explicit nil

### UnsetDdos
`func (o *IpDetails) UnsetDdos()`

UnsetDdos ensures that no value is present for Ddos, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


