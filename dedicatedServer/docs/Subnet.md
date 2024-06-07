# Subnet

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Quantity** | Pointer to **int32** |  | [optional] 
**SubnetSize** | Pointer to **string** |  | [optional] 
**NetworkType** | Pointer to [**NetworkType**](NetworkType.md) |  | [optional] 

## Methods

### NewSubnet

`func NewSubnet() *Subnet`

NewSubnet instantiates a new Subnet object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSubnetWithDefaults

`func NewSubnetWithDefaults() *Subnet`

NewSubnetWithDefaults instantiates a new Subnet object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetQuantity

`func (o *Subnet) GetQuantity() int32`

GetQuantity returns the Quantity field if non-nil, zero value otherwise.

### GetQuantityOk

`func (o *Subnet) GetQuantityOk() (*int32, bool)`

GetQuantityOk returns a tuple with the Quantity field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQuantity

`func (o *Subnet) SetQuantity(v int32)`

SetQuantity sets Quantity field to given value.

### HasQuantity

`func (o *Subnet) HasQuantity() bool`

HasQuantity returns a boolean if a field has been set.

### GetSubnetSize

`func (o *Subnet) GetSubnetSize() string`

GetSubnetSize returns the SubnetSize field if non-nil, zero value otherwise.

### GetSubnetSizeOk

`func (o *Subnet) GetSubnetSizeOk() (*string, bool)`

GetSubnetSizeOk returns a tuple with the SubnetSize field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubnetSize

`func (o *Subnet) SetSubnetSize(v string)`

SetSubnetSize sets SubnetSize field to given value.

### HasSubnetSize

`func (o *Subnet) HasSubnetSize() bool`

HasSubnetSize returns a boolean if a field has been set.

### GetNetworkType

`func (o *Subnet) GetNetworkType() NetworkType`

GetNetworkType returns the NetworkType field if non-nil, zero value otherwise.

### GetNetworkTypeOk

`func (o *Subnet) GetNetworkTypeOk() (*NetworkType, bool)`

GetNetworkTypeOk returns a tuple with the NetworkType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNetworkType

`func (o *Subnet) SetNetworkType(v NetworkType)`

SetNetworkType sets NetworkType field to given value.

### HasNetworkType

`func (o *Subnet) HasNetworkType() bool`

HasNetworkType returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


