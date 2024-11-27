# Lldp

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Chassis** | Pointer to [**LldpChassis**](LldpChassis.md) |  | [optional] 
**Port** | Pointer to [**Port1**](Port1.md) |  | [optional] 
**Vlan** | Pointer to [**Vlan**](Vlan.md) |  | [optional] 

## Methods

### NewLldp

`func NewLldp() *Lldp`

NewLldp instantiates a new Lldp object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewLldpWithDefaults

`func NewLldpWithDefaults() *Lldp`

NewLldpWithDefaults instantiates a new Lldp object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetChassis

`func (o *Lldp) GetChassis() LldpChassis`

GetChassis returns the Chassis field if non-nil, zero value otherwise.

### GetChassisOk

`func (o *Lldp) GetChassisOk() (*LldpChassis, bool)`

GetChassisOk returns a tuple with the Chassis field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetChassis

`func (o *Lldp) SetChassis(v LldpChassis)`

SetChassis sets Chassis field to given value.

### HasChassis

`func (o *Lldp) HasChassis() bool`

HasChassis returns a boolean if a field has been set.

### GetPort

`func (o *Lldp) GetPort() Port1`

GetPort returns the Port field if non-nil, zero value otherwise.

### GetPortOk

`func (o *Lldp) GetPortOk() (*Port1, bool)`

GetPortOk returns a tuple with the Port field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPort

`func (o *Lldp) SetPort(v Port1)`

SetPort sets Port field to given value.

### HasPort

`func (o *Lldp) HasPort() bool`

HasPort returns a boolean if a field has been set.

### GetVlan

`func (o *Lldp) GetVlan() Vlan`

GetVlan returns the Vlan field if non-nil, zero value otherwise.

### GetVlanOk

`func (o *Lldp) GetVlanOk() (*Vlan, bool)`

GetVlanOk returns a tuple with the Vlan field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVlan

`func (o *Lldp) SetVlan(v Vlan)`

SetVlan sets Vlan field to given value.

### HasVlan

`func (o *Lldp) HasVlan() bool`

HasVlan returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


