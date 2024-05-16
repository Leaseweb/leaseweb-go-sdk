# NetworkInterface

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Mac** | Pointer to **NullableString** | Represents a MAC Address in the standard colon delimited format. Eg. &#x60;01:23:45:67:89:0A&#x60; | [optional] 
**Ip** | Pointer to **NullableString** |  | [optional] 
**NullRouted** | Pointer to **NullableBool** |  | [optional] 
**Gateway** | Pointer to **NullableString** |  | [optional] 
**Ports** | Pointer to [**[]Port**](Port.md) |  | [optional] 
**LocationId** | Pointer to **NullableString** |  | [optional] 

## Methods

### NewNetworkInterface

`func NewNetworkInterface() *NetworkInterface`

NewNetworkInterface instantiates a new NetworkInterface object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewNetworkInterfaceWithDefaults

`func NewNetworkInterfaceWithDefaults() *NetworkInterface`

NewNetworkInterfaceWithDefaults instantiates a new NetworkInterface object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetMac

`func (o *NetworkInterface) GetMac() string`

GetMac returns the Mac field if non-nil, zero value otherwise.

### GetMacOk

`func (o *NetworkInterface) GetMacOk() (*string, bool)`

GetMacOk returns a tuple with the Mac field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMac

`func (o *NetworkInterface) SetMac(v string)`

SetMac sets Mac field to given value.

### HasMac

`func (o *NetworkInterface) HasMac() bool`

HasMac returns a boolean if a field has been set.

### SetMacNil

`func (o *NetworkInterface) SetMacNil(b bool)`

 SetMacNil sets the value for Mac to be an explicit nil

### UnsetMac
`func (o *NetworkInterface) UnsetMac()`

UnsetMac ensures that no value is present for Mac, not even an explicit nil
### GetIp

`func (o *NetworkInterface) GetIp() string`

GetIp returns the Ip field if non-nil, zero value otherwise.

### GetIpOk

`func (o *NetworkInterface) GetIpOk() (*string, bool)`

GetIpOk returns a tuple with the Ip field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIp

`func (o *NetworkInterface) SetIp(v string)`

SetIp sets Ip field to given value.

### HasIp

`func (o *NetworkInterface) HasIp() bool`

HasIp returns a boolean if a field has been set.

### SetIpNil

`func (o *NetworkInterface) SetIpNil(b bool)`

 SetIpNil sets the value for Ip to be an explicit nil

### UnsetIp
`func (o *NetworkInterface) UnsetIp()`

UnsetIp ensures that no value is present for Ip, not even an explicit nil
### GetNullRouted

`func (o *NetworkInterface) GetNullRouted() bool`

GetNullRouted returns the NullRouted field if non-nil, zero value otherwise.

### GetNullRoutedOk

`func (o *NetworkInterface) GetNullRoutedOk() (*bool, bool)`

GetNullRoutedOk returns a tuple with the NullRouted field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNullRouted

`func (o *NetworkInterface) SetNullRouted(v bool)`

SetNullRouted sets NullRouted field to given value.

### HasNullRouted

`func (o *NetworkInterface) HasNullRouted() bool`

HasNullRouted returns a boolean if a field has been set.

### SetNullRoutedNil

`func (o *NetworkInterface) SetNullRoutedNil(b bool)`

 SetNullRoutedNil sets the value for NullRouted to be an explicit nil

### UnsetNullRouted
`func (o *NetworkInterface) UnsetNullRouted()`

UnsetNullRouted ensures that no value is present for NullRouted, not even an explicit nil
### GetGateway

`func (o *NetworkInterface) GetGateway() string`

GetGateway returns the Gateway field if non-nil, zero value otherwise.

### GetGatewayOk

`func (o *NetworkInterface) GetGatewayOk() (*string, bool)`

GetGatewayOk returns a tuple with the Gateway field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGateway

`func (o *NetworkInterface) SetGateway(v string)`

SetGateway sets Gateway field to given value.

### HasGateway

`func (o *NetworkInterface) HasGateway() bool`

HasGateway returns a boolean if a field has been set.

### SetGatewayNil

`func (o *NetworkInterface) SetGatewayNil(b bool)`

 SetGatewayNil sets the value for Gateway to be an explicit nil

### UnsetGateway
`func (o *NetworkInterface) UnsetGateway()`

UnsetGateway ensures that no value is present for Gateway, not even an explicit nil
### GetPorts

`func (o *NetworkInterface) GetPorts() []Port`

GetPorts returns the Ports field if non-nil, zero value otherwise.

### GetPortsOk

`func (o *NetworkInterface) GetPortsOk() (*[]Port, bool)`

GetPortsOk returns a tuple with the Ports field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPorts

`func (o *NetworkInterface) SetPorts(v []Port)`

SetPorts sets Ports field to given value.

### HasPorts

`func (o *NetworkInterface) HasPorts() bool`

HasPorts returns a boolean if a field has been set.

### GetLocationId

`func (o *NetworkInterface) GetLocationId() string`

GetLocationId returns the LocationId field if non-nil, zero value otherwise.

### GetLocationIdOk

`func (o *NetworkInterface) GetLocationIdOk() (*string, bool)`

GetLocationIdOk returns a tuple with the LocationId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLocationId

`func (o *NetworkInterface) SetLocationId(v string)`

SetLocationId sets LocationId field to given value.

### HasLocationId

`func (o *NetworkInterface) HasLocationId() bool`

HasLocationId returns a boolean if a field has been set.

### SetLocationIdNil

`func (o *NetworkInterface) SetLocationIdNil(b bool)`

 SetLocationIdNil sets the value for LocationId to be an explicit nil

### UnsetLocationId
`func (o *NetworkInterface) UnsetLocationId()`

UnsetLocationId ensures that no value is present for LocationId, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


