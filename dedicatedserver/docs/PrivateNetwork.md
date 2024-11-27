# PrivateNetwork

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **string** | Private network id | [optional] 
**LinkSpeed** | Pointer to [**LinkSpeed**](LinkSpeed.md) |  | [optional] 
**Status** | Pointer to **string** | Configuration status | [optional] 
**Subnet** | Pointer to **string** |  | [optional] 
**VlanId** | Pointer to **string** | VLAN id | [optional] 

## Methods

### NewPrivateNetwork

`func NewPrivateNetwork() *PrivateNetwork`

NewPrivateNetwork instantiates a new PrivateNetwork object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPrivateNetworkWithDefaults

`func NewPrivateNetworkWithDefaults() *PrivateNetwork`

NewPrivateNetworkWithDefaults instantiates a new PrivateNetwork object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *PrivateNetwork) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *PrivateNetwork) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *PrivateNetwork) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *PrivateNetwork) HasId() bool`

HasId returns a boolean if a field has been set.

### GetLinkSpeed

`func (o *PrivateNetwork) GetLinkSpeed() LinkSpeed`

GetLinkSpeed returns the LinkSpeed field if non-nil, zero value otherwise.

### GetLinkSpeedOk

`func (o *PrivateNetwork) GetLinkSpeedOk() (*LinkSpeed, bool)`

GetLinkSpeedOk returns a tuple with the LinkSpeed field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLinkSpeed

`func (o *PrivateNetwork) SetLinkSpeed(v LinkSpeed)`

SetLinkSpeed sets LinkSpeed field to given value.

### HasLinkSpeed

`func (o *PrivateNetwork) HasLinkSpeed() bool`

HasLinkSpeed returns a boolean if a field has been set.

### GetStatus

`func (o *PrivateNetwork) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *PrivateNetwork) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *PrivateNetwork) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *PrivateNetwork) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetSubnet

`func (o *PrivateNetwork) GetSubnet() string`

GetSubnet returns the Subnet field if non-nil, zero value otherwise.

### GetSubnetOk

`func (o *PrivateNetwork) GetSubnetOk() (*string, bool)`

GetSubnetOk returns a tuple with the Subnet field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubnet

`func (o *PrivateNetwork) SetSubnet(v string)`

SetSubnet sets Subnet field to given value.

### HasSubnet

`func (o *PrivateNetwork) HasSubnet() bool`

HasSubnet returns a boolean if a field has been set.

### GetVlanId

`func (o *PrivateNetwork) GetVlanId() string`

GetVlanId returns the VlanId field if non-nil, zero value otherwise.

### GetVlanIdOk

`func (o *PrivateNetwork) GetVlanIdOk() (*string, bool)`

GetVlanIdOk returns a tuple with the VlanId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVlanId

`func (o *PrivateNetwork) SetVlanId(v string)`

SetVlanId sets VlanId field to given value.

### HasVlanId

`func (o *PrivateNetwork) HasVlanId() bool`

HasVlanId returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


