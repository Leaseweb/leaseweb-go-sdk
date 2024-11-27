# PrivateNetwork

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**PrivateNetworkId** | **string** |  | 
**Status** | **string** |  | 
**Subnet** | **string** |  | 

## Methods

### NewPrivateNetwork

`func NewPrivateNetwork(privateNetworkId string, status string, subnet string, ) *PrivateNetwork`

NewPrivateNetwork instantiates a new PrivateNetwork object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPrivateNetworkWithDefaults

`func NewPrivateNetworkWithDefaults() *PrivateNetwork`

NewPrivateNetworkWithDefaults instantiates a new PrivateNetwork object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetPrivateNetworkId

`func (o *PrivateNetwork) GetPrivateNetworkId() string`

GetPrivateNetworkId returns the PrivateNetworkId field if non-nil, zero value otherwise.

### GetPrivateNetworkIdOk

`func (o *PrivateNetwork) GetPrivateNetworkIdOk() (*string, bool)`

GetPrivateNetworkIdOk returns a tuple with the PrivateNetworkId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrivateNetworkId

`func (o *PrivateNetwork) SetPrivateNetworkId(v string)`

SetPrivateNetworkId sets PrivateNetworkId field to given value.


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



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


