# Subnet

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** | Subnet identifier consisting of network IP and prefix length separated by underscore (e.g. 192.0.2.0_24) | 
**NetworkIp** | **string** | Network IP of the subnet | 
**PrefixLength** | **int32** | Address prefix length | [default to 0]
**Gateway** | **string** | The gateway IP to be used in network settings | 

## Methods

### NewSubnet

`func NewSubnet(id string, networkIp string, prefixLength int32, gateway string, ) *Subnet`

NewSubnet instantiates a new Subnet object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSubnetWithDefaults

`func NewSubnetWithDefaults() *Subnet`

NewSubnetWithDefaults instantiates a new Subnet object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *Subnet) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *Subnet) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *Subnet) SetId(v string)`

SetId sets Id field to given value.


### GetNetworkIp

`func (o *Subnet) GetNetworkIp() string`

GetNetworkIp returns the NetworkIp field if non-nil, zero value otherwise.

### GetNetworkIpOk

`func (o *Subnet) GetNetworkIpOk() (*string, bool)`

GetNetworkIpOk returns a tuple with the NetworkIp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNetworkIp

`func (o *Subnet) SetNetworkIp(v string)`

SetNetworkIp sets NetworkIp field to given value.


### GetPrefixLength

`func (o *Subnet) GetPrefixLength() int32`

GetPrefixLength returns the PrefixLength field if non-nil, zero value otherwise.

### GetPrefixLengthOk

`func (o *Subnet) GetPrefixLengthOk() (*int32, bool)`

GetPrefixLengthOk returns a tuple with the PrefixLength field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrefixLength

`func (o *Subnet) SetPrefixLength(v int32)`

SetPrefixLength sets PrefixLength field to given value.


### GetGateway

`func (o *Subnet) GetGateway() string`

GetGateway returns the Gateway field if non-nil, zero value otherwise.

### GetGatewayOk

`func (o *Subnet) GetGatewayOk() (*string, bool)`

GetGatewayOk returns a tuple with the Gateway field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGateway

`func (o *Subnet) SetGateway(v string)`

SetGateway sets Gateway field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


