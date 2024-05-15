# NetworkSettings

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Autonegotiation** | Pointer to **string** |  | [optional] 
**Broadcast** | Pointer to **string** |  | [optional] 
**Driver** | Pointer to **string** |  | [optional] 
**Driverversion** | Pointer to **string** |  | [optional] 
**Duplex** | Pointer to **string** |  | [optional] 
**Firmware** | Pointer to **string** |  | [optional] 
**Ip** | Pointer to **string** |  | [optional] 
**Latency** | Pointer to **string** |  | [optional] 
**Link** | Pointer to **string** |  | [optional] 
**Multicast** | Pointer to **string** |  | [optional] 
**Port** | Pointer to **string** |  | [optional] 
**Speed** | Pointer to **string** |  | [optional] 

## Methods

### NewNetworkSettings

`func NewNetworkSettings() *NetworkSettings`

NewNetworkSettings instantiates a new NetworkSettings object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewNetworkSettingsWithDefaults

`func NewNetworkSettingsWithDefaults() *NetworkSettings`

NewNetworkSettingsWithDefaults instantiates a new NetworkSettings object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAutonegotiation

`func (o *NetworkSettings) GetAutonegotiation() string`

GetAutonegotiation returns the Autonegotiation field if non-nil, zero value otherwise.

### GetAutonegotiationOk

`func (o *NetworkSettings) GetAutonegotiationOk() (*string, bool)`

GetAutonegotiationOk returns a tuple with the Autonegotiation field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAutonegotiation

`func (o *NetworkSettings) SetAutonegotiation(v string)`

SetAutonegotiation sets Autonegotiation field to given value.

### HasAutonegotiation

`func (o *NetworkSettings) HasAutonegotiation() bool`

HasAutonegotiation returns a boolean if a field has been set.

### GetBroadcast

`func (o *NetworkSettings) GetBroadcast() string`

GetBroadcast returns the Broadcast field if non-nil, zero value otherwise.

### GetBroadcastOk

`func (o *NetworkSettings) GetBroadcastOk() (*string, bool)`

GetBroadcastOk returns a tuple with the Broadcast field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBroadcast

`func (o *NetworkSettings) SetBroadcast(v string)`

SetBroadcast sets Broadcast field to given value.

### HasBroadcast

`func (o *NetworkSettings) HasBroadcast() bool`

HasBroadcast returns a boolean if a field has been set.

### GetDriver

`func (o *NetworkSettings) GetDriver() string`

GetDriver returns the Driver field if non-nil, zero value otherwise.

### GetDriverOk

`func (o *NetworkSettings) GetDriverOk() (*string, bool)`

GetDriverOk returns a tuple with the Driver field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDriver

`func (o *NetworkSettings) SetDriver(v string)`

SetDriver sets Driver field to given value.

### HasDriver

`func (o *NetworkSettings) HasDriver() bool`

HasDriver returns a boolean if a field has been set.

### GetDriverversion

`func (o *NetworkSettings) GetDriverversion() string`

GetDriverversion returns the Driverversion field if non-nil, zero value otherwise.

### GetDriverversionOk

`func (o *NetworkSettings) GetDriverversionOk() (*string, bool)`

GetDriverversionOk returns a tuple with the Driverversion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDriverversion

`func (o *NetworkSettings) SetDriverversion(v string)`

SetDriverversion sets Driverversion field to given value.

### HasDriverversion

`func (o *NetworkSettings) HasDriverversion() bool`

HasDriverversion returns a boolean if a field has been set.

### GetDuplex

`func (o *NetworkSettings) GetDuplex() string`

GetDuplex returns the Duplex field if non-nil, zero value otherwise.

### GetDuplexOk

`func (o *NetworkSettings) GetDuplexOk() (*string, bool)`

GetDuplexOk returns a tuple with the Duplex field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDuplex

`func (o *NetworkSettings) SetDuplex(v string)`

SetDuplex sets Duplex field to given value.

### HasDuplex

`func (o *NetworkSettings) HasDuplex() bool`

HasDuplex returns a boolean if a field has been set.

### GetFirmware

`func (o *NetworkSettings) GetFirmware() string`

GetFirmware returns the Firmware field if non-nil, zero value otherwise.

### GetFirmwareOk

`func (o *NetworkSettings) GetFirmwareOk() (*string, bool)`

GetFirmwareOk returns a tuple with the Firmware field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFirmware

`func (o *NetworkSettings) SetFirmware(v string)`

SetFirmware sets Firmware field to given value.

### HasFirmware

`func (o *NetworkSettings) HasFirmware() bool`

HasFirmware returns a boolean if a field has been set.

### GetIp

`func (o *NetworkSettings) GetIp() string`

GetIp returns the Ip field if non-nil, zero value otherwise.

### GetIpOk

`func (o *NetworkSettings) GetIpOk() (*string, bool)`

GetIpOk returns a tuple with the Ip field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIp

`func (o *NetworkSettings) SetIp(v string)`

SetIp sets Ip field to given value.

### HasIp

`func (o *NetworkSettings) HasIp() bool`

HasIp returns a boolean if a field has been set.

### GetLatency

`func (o *NetworkSettings) GetLatency() string`

GetLatency returns the Latency field if non-nil, zero value otherwise.

### GetLatencyOk

`func (o *NetworkSettings) GetLatencyOk() (*string, bool)`

GetLatencyOk returns a tuple with the Latency field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLatency

`func (o *NetworkSettings) SetLatency(v string)`

SetLatency sets Latency field to given value.

### HasLatency

`func (o *NetworkSettings) HasLatency() bool`

HasLatency returns a boolean if a field has been set.

### GetLink

`func (o *NetworkSettings) GetLink() string`

GetLink returns the Link field if non-nil, zero value otherwise.

### GetLinkOk

`func (o *NetworkSettings) GetLinkOk() (*string, bool)`

GetLinkOk returns a tuple with the Link field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLink

`func (o *NetworkSettings) SetLink(v string)`

SetLink sets Link field to given value.

### HasLink

`func (o *NetworkSettings) HasLink() bool`

HasLink returns a boolean if a field has been set.

### GetMulticast

`func (o *NetworkSettings) GetMulticast() string`

GetMulticast returns the Multicast field if non-nil, zero value otherwise.

### GetMulticastOk

`func (o *NetworkSettings) GetMulticastOk() (*string, bool)`

GetMulticastOk returns a tuple with the Multicast field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMulticast

`func (o *NetworkSettings) SetMulticast(v string)`

SetMulticast sets Multicast field to given value.

### HasMulticast

`func (o *NetworkSettings) HasMulticast() bool`

HasMulticast returns a boolean if a field has been set.

### GetPort

`func (o *NetworkSettings) GetPort() string`

GetPort returns the Port field if non-nil, zero value otherwise.

### GetPortOk

`func (o *NetworkSettings) GetPortOk() (*string, bool)`

GetPortOk returns a tuple with the Port field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPort

`func (o *NetworkSettings) SetPort(v string)`

SetPort sets Port field to given value.

### HasPort

`func (o *NetworkSettings) HasPort() bool`

HasPort returns a boolean if a field has been set.

### GetSpeed

`func (o *NetworkSettings) GetSpeed() string`

GetSpeed returns the Speed field if non-nil, zero value otherwise.

### GetSpeedOk

`func (o *NetworkSettings) GetSpeedOk() (*string, bool)`

GetSpeedOk returns a tuple with the Speed field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSpeed

`func (o *NetworkSettings) SetSpeed(v string)`

SetSpeed sets Speed field to given value.

### HasSpeed

`func (o *NetworkSettings) HasSpeed() bool`

HasSpeed returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


