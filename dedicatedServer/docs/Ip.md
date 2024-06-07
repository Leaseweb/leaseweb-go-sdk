# Ip

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Ddos** | Pointer to [**DDos**](DDos.md) |  | [optional] 
**FloatingIp** | Pointer to **bool** | Indicates if the IP is a Floating IP | [optional] 
**Gateway** | Pointer to **string** | Gateway | [optional] 
**Ip** | Pointer to **string** | IP address in CIDR notation | [optional] 
**MainIp** | Pointer to **bool** | IP address is main | [optional] 
**NetworkType** | Pointer to [**NetworkType**](NetworkType.md) |  | [optional] 
**NullRouted** | Pointer to **bool** | IP address null routed | [optional] 
**ReverseLookup** | Pointer to **string** | The reverse lookup value | [optional] 
**Version** | Pointer to **int32** | IP version | [optional] 

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

### GetDdos

`func (o *Ip) GetDdos() DDos`

GetDdos returns the Ddos field if non-nil, zero value otherwise.

### GetDdosOk

`func (o *Ip) GetDdosOk() (*DDos, bool)`

GetDdosOk returns a tuple with the Ddos field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDdos

`func (o *Ip) SetDdos(v DDos)`

SetDdos sets Ddos field to given value.

### HasDdos

`func (o *Ip) HasDdos() bool`

HasDdos returns a boolean if a field has been set.

### GetFloatingIp

`func (o *Ip) GetFloatingIp() bool`

GetFloatingIp returns the FloatingIp field if non-nil, zero value otherwise.

### GetFloatingIpOk

`func (o *Ip) GetFloatingIpOk() (*bool, bool)`

GetFloatingIpOk returns a tuple with the FloatingIp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFloatingIp

`func (o *Ip) SetFloatingIp(v bool)`

SetFloatingIp sets FloatingIp field to given value.

### HasFloatingIp

`func (o *Ip) HasFloatingIp() bool`

HasFloatingIp returns a boolean if a field has been set.

### GetGateway

`func (o *Ip) GetGateway() string`

GetGateway returns the Gateway field if non-nil, zero value otherwise.

### GetGatewayOk

`func (o *Ip) GetGatewayOk() (*string, bool)`

GetGatewayOk returns a tuple with the Gateway field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGateway

`func (o *Ip) SetGateway(v string)`

SetGateway sets Gateway field to given value.

### HasGateway

`func (o *Ip) HasGateway() bool`

HasGateway returns a boolean if a field has been set.

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

### GetNetworkType

`func (o *Ip) GetNetworkType() NetworkType`

GetNetworkType returns the NetworkType field if non-nil, zero value otherwise.

### GetNetworkTypeOk

`func (o *Ip) GetNetworkTypeOk() (*NetworkType, bool)`

GetNetworkTypeOk returns a tuple with the NetworkType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNetworkType

`func (o *Ip) SetNetworkType(v NetworkType)`

SetNetworkType sets NetworkType field to given value.

### HasNetworkType

`func (o *Ip) HasNetworkType() bool`

HasNetworkType returns a boolean if a field has been set.

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


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


