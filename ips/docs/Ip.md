# Ip

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Ip** | **string** | IP address | 
**Version** | [**ProtocolVersion**](ProtocolVersion.md) |  | 
**Type** | [**IpType**](IpType.md) |  | 
**PrefixLength** | **int32** | Prefix length of the IP range represented by the record. Note: this is not the same as &#x60;subnet.prefixLength&#x60;. | [default to 0]
**Primary** | **bool** | Boolean indicating if this is the primary IP of the assigned equipment | 
**ReverseLookup** | **NullableString** | Reverse lookup set for the IP. This only applies to IPv4. For IPv6 see [GET /ips/{ip}/reverseLookup](#operation/get/ips/{ip}/reverseLookup). | 
**NullRouted** | **bool** | Boolean to indicate if the IP is null-routed | [default to false]
**NullLevel** | **NullableInt32** | Null route level | 
**UnnullingAllowed** | **bool** | Boolean indicating if the null route can be removed | [default to false]
**EquipmentId** | **string** | ID of the equipment using the IP | 
**AssignedContract** | [**NullableAssignedContract**](AssignedContract.md) |  | 
**Subnet** | [**Subnet**](Subnet.md) |  | [default to {}]

## Methods

### NewIp

`func NewIp(ip string, version ProtocolVersion, type_ IpType, prefixLength int32, primary bool, reverseLookup NullableString, nullRouted bool, nullLevel NullableInt32, unnullingAllowed bool, equipmentId string, assignedContract NullableAssignedContract, subnet Subnet, ) *Ip`

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


### GetVersion

`func (o *Ip) GetVersion() ProtocolVersion`

GetVersion returns the Version field if non-nil, zero value otherwise.

### GetVersionOk

`func (o *Ip) GetVersionOk() (*ProtocolVersion, bool)`

GetVersionOk returns a tuple with the Version field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVersion

`func (o *Ip) SetVersion(v ProtocolVersion)`

SetVersion sets Version field to given value.


### GetType

`func (o *Ip) GetType() IpType`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *Ip) GetTypeOk() (*IpType, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *Ip) SetType(v IpType)`

SetType sets Type field to given value.


### GetPrefixLength

`func (o *Ip) GetPrefixLength() int32`

GetPrefixLength returns the PrefixLength field if non-nil, zero value otherwise.

### GetPrefixLengthOk

`func (o *Ip) GetPrefixLengthOk() (*int32, bool)`

GetPrefixLengthOk returns a tuple with the PrefixLength field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrefixLength

`func (o *Ip) SetPrefixLength(v int32)`

SetPrefixLength sets PrefixLength field to given value.


### GetPrimary

`func (o *Ip) GetPrimary() bool`

GetPrimary returns the Primary field if non-nil, zero value otherwise.

### GetPrimaryOk

`func (o *Ip) GetPrimaryOk() (*bool, bool)`

GetPrimaryOk returns a tuple with the Primary field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrimary

`func (o *Ip) SetPrimary(v bool)`

SetPrimary sets Primary field to given value.


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


### SetReverseLookupNil

`func (o *Ip) SetReverseLookupNil(b bool)`

 SetReverseLookupNil sets the value for ReverseLookup to be an explicit nil

### UnsetReverseLookup
`func (o *Ip) UnsetReverseLookup()`

UnsetReverseLookup ensures that no value is present for ReverseLookup, not even an explicit nil
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


### GetNullLevel

`func (o *Ip) GetNullLevel() int32`

GetNullLevel returns the NullLevel field if non-nil, zero value otherwise.

### GetNullLevelOk

`func (o *Ip) GetNullLevelOk() (*int32, bool)`

GetNullLevelOk returns a tuple with the NullLevel field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNullLevel

`func (o *Ip) SetNullLevel(v int32)`

SetNullLevel sets NullLevel field to given value.


### SetNullLevelNil

`func (o *Ip) SetNullLevelNil(b bool)`

 SetNullLevelNil sets the value for NullLevel to be an explicit nil

### UnsetNullLevel
`func (o *Ip) UnsetNullLevel()`

UnsetNullLevel ensures that no value is present for NullLevel, not even an explicit nil
### GetUnnullingAllowed

`func (o *Ip) GetUnnullingAllowed() bool`

GetUnnullingAllowed returns the UnnullingAllowed field if non-nil, zero value otherwise.

### GetUnnullingAllowedOk

`func (o *Ip) GetUnnullingAllowedOk() (*bool, bool)`

GetUnnullingAllowedOk returns a tuple with the UnnullingAllowed field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUnnullingAllowed

`func (o *Ip) SetUnnullingAllowed(v bool)`

SetUnnullingAllowed sets UnnullingAllowed field to given value.


### GetEquipmentId

`func (o *Ip) GetEquipmentId() string`

GetEquipmentId returns the EquipmentId field if non-nil, zero value otherwise.

### GetEquipmentIdOk

`func (o *Ip) GetEquipmentIdOk() (*string, bool)`

GetEquipmentIdOk returns a tuple with the EquipmentId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEquipmentId

`func (o *Ip) SetEquipmentId(v string)`

SetEquipmentId sets EquipmentId field to given value.


### GetAssignedContract

`func (o *Ip) GetAssignedContract() AssignedContract`

GetAssignedContract returns the AssignedContract field if non-nil, zero value otherwise.

### GetAssignedContractOk

`func (o *Ip) GetAssignedContractOk() (*AssignedContract, bool)`

GetAssignedContractOk returns a tuple with the AssignedContract field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAssignedContract

`func (o *Ip) SetAssignedContract(v AssignedContract)`

SetAssignedContract sets AssignedContract field to given value.


### SetAssignedContractNil

`func (o *Ip) SetAssignedContractNil(b bool)`

 SetAssignedContractNil sets the value for AssignedContract to be an explicit nil

### UnsetAssignedContract
`func (o *Ip) UnsetAssignedContract()`

UnsetAssignedContract ensures that no value is present for AssignedContract, not even an explicit nil
### GetSubnet

`func (o *Ip) GetSubnet() Subnet`

GetSubnet returns the Subnet field if non-nil, zero value otherwise.

### GetSubnetOk

`func (o *Ip) GetSubnetOk() (*Subnet, bool)`

GetSubnetOk returns a tuple with the Subnet field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubnet

`func (o *Ip) SetSubnet(v Subnet)`

SetSubnet sets Subnet field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


