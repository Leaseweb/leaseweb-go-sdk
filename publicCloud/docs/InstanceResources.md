# InstanceResources

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Cpu** | [**Cpu**](Cpu.md) |  | 
**Memory** | [**Memory**](Memory.md) |  | 
**PublicNetworkSpeed** | [**NetworkSpeed**](NetworkSpeed.md) |  | 
**PrivateNetworkSpeed** | [**NetworkSpeed**](NetworkSpeed.md) |  | 

## Methods

### NewInstanceResources

`func NewInstanceResources(cpu Cpu, memory Memory, publicNetworkSpeed NetworkSpeed, privateNetworkSpeed NetworkSpeed, ) *InstanceResources`

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


### GetPublicNetworkSpeed

`func (o *InstanceResources) GetPublicNetworkSpeed() NetworkSpeed`

GetPublicNetworkSpeed returns the PublicNetworkSpeed field if non-nil, zero value otherwise.

### GetPublicNetworkSpeedOk

`func (o *InstanceResources) GetPublicNetworkSpeedOk() (*NetworkSpeed, bool)`

GetPublicNetworkSpeedOk returns a tuple with the PublicNetworkSpeed field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPublicNetworkSpeed

`func (o *InstanceResources) SetPublicNetworkSpeed(v NetworkSpeed)`

SetPublicNetworkSpeed sets PublicNetworkSpeed field to given value.


### GetPrivateNetworkSpeed

`func (o *InstanceResources) GetPrivateNetworkSpeed() NetworkSpeed`

GetPrivateNetworkSpeed returns the PrivateNetworkSpeed field if non-nil, zero value otherwise.

### GetPrivateNetworkSpeedOk

`func (o *InstanceResources) GetPrivateNetworkSpeedOk() (*NetworkSpeed, bool)`

GetPrivateNetworkSpeedOk returns a tuple with the PrivateNetworkSpeed field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrivateNetworkSpeed

`func (o *InstanceResources) SetPrivateNetworkSpeed(v NetworkSpeed)`

SetPrivateNetworkSpeed sets PrivateNetworkSpeed field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


