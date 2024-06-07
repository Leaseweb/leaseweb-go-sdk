# InstanceResources

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Cpu** | Pointer to [**Cpu**](Cpu.md) |  | [optional] 
**Memory** | Pointer to [**Memory**](Memory.md) |  | [optional] 
**PublicNetworkSpeed** | Pointer to [**PublicNetworkSpeed**](PublicNetworkSpeed.md) |  | [optional] 
**PrivateNetworkSpeed** | Pointer to [**PrivateNetworkSpeed**](PrivateNetworkSpeed.md) |  | [optional] 

## Methods

### NewInstanceResources

`func NewInstanceResources() *InstanceResources`

NewInstanceResources instantiates a new InstanceResources object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewInstanceResourcesWithDefaults

`func NewInstanceResourcesWithDefaults() *InstanceResources`

NewInstanceResourcesWithDefaults instantiates a new InstanceResources object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCpu

`func (o *InstanceResources) GetCpu() Cpu`

GetCpu returns the Cpu field if non-nil, zero value otherwise.

### GetCpuOk

`func (o *InstanceResources) GetCpuOk() (*Cpu, bool)`

GetCpuOk returns a tuple with the Cpu field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCpu

`func (o *InstanceResources) SetCpu(v Cpu)`

SetCpu sets Cpu field to given value.

### HasCpu

`func (o *InstanceResources) HasCpu() bool`

HasCpu returns a boolean if a field has been set.

### GetMemory

`func (o *InstanceResources) GetMemory() Memory`

GetMemory returns the Memory field if non-nil, zero value otherwise.

### GetMemoryOk

`func (o *InstanceResources) GetMemoryOk() (*Memory, bool)`

GetMemoryOk returns a tuple with the Memory field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMemory

`func (o *InstanceResources) SetMemory(v Memory)`

SetMemory sets Memory field to given value.

### HasMemory

`func (o *InstanceResources) HasMemory() bool`

HasMemory returns a boolean if a field has been set.

### GetPublicNetworkSpeed

`func (o *InstanceResources) GetPublicNetworkSpeed() PublicNetworkSpeed`

GetPublicNetworkSpeed returns the PublicNetworkSpeed field if non-nil, zero value otherwise.

### GetPublicNetworkSpeedOk

`func (o *InstanceResources) GetPublicNetworkSpeedOk() (*PublicNetworkSpeed, bool)`

GetPublicNetworkSpeedOk returns a tuple with the PublicNetworkSpeed field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPublicNetworkSpeed

`func (o *InstanceResources) SetPublicNetworkSpeed(v PublicNetworkSpeed)`

SetPublicNetworkSpeed sets PublicNetworkSpeed field to given value.

### HasPublicNetworkSpeed

`func (o *InstanceResources) HasPublicNetworkSpeed() bool`

HasPublicNetworkSpeed returns a boolean if a field has been set.

### GetPrivateNetworkSpeed

`func (o *InstanceResources) GetPrivateNetworkSpeed() PrivateNetworkSpeed`

GetPrivateNetworkSpeed returns the PrivateNetworkSpeed field if non-nil, zero value otherwise.

### GetPrivateNetworkSpeedOk

`func (o *InstanceResources) GetPrivateNetworkSpeedOk() (*PrivateNetworkSpeed, bool)`

GetPrivateNetworkSpeedOk returns a tuple with the PrivateNetworkSpeed field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrivateNetworkSpeed

`func (o *InstanceResources) SetPrivateNetworkSpeed(v PrivateNetworkSpeed)`

SetPrivateNetworkSpeed sets PrivateNetworkSpeed field to given value.

### HasPrivateNetworkSpeed

`func (o *InstanceResources) HasPrivateNetworkSpeed() bool`

HasPrivateNetworkSpeed returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


