# Resources

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Cpu** | [**Cpu**](Cpu.md) |  | 
**Memory** | [**Memory**](Memory.md) |  | 
**PublicNetworkSpeed** | [**NetworkSpeed**](NetworkSpeed.md) |  | 
**PrivateNetworkSpeed** | [**NetworkSpeed**](NetworkSpeed.md) |  | 

## Methods

### NewResources

`func NewResources(cpu Cpu, memory Memory, publicNetworkSpeed NetworkSpeed, privateNetworkSpeed NetworkSpeed, ) *Resources`

NewResources instantiates a new Resources object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewResourcesWithDefaults

`func NewResourcesWithDefaults() *Resources`

NewResourcesWithDefaults instantiates a new Resources object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCpu

`func (o *Resources) GetCpu() Cpu`

GetCpu returns the Cpu field if non-nil, zero value otherwise.

### GetCpuOk

`func (o *Resources) GetCpuOk() (*Cpu, bool)`

GetCpuOk returns a tuple with the Cpu field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCpu

`func (o *Resources) SetCpu(v Cpu)`

SetCpu sets Cpu field to given value.


### GetMemory

`func (o *Resources) GetMemory() Memory`

GetMemory returns the Memory field if non-nil, zero value otherwise.

### GetMemoryOk

`func (o *Resources) GetMemoryOk() (*Memory, bool)`

GetMemoryOk returns a tuple with the Memory field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMemory

`func (o *Resources) SetMemory(v Memory)`

SetMemory sets Memory field to given value.


### GetPublicNetworkSpeed

`func (o *Resources) GetPublicNetworkSpeed() NetworkSpeed`

GetPublicNetworkSpeed returns the PublicNetworkSpeed field if non-nil, zero value otherwise.

### GetPublicNetworkSpeedOk

`func (o *Resources) GetPublicNetworkSpeedOk() (*NetworkSpeed, bool)`

GetPublicNetworkSpeedOk returns a tuple with the PublicNetworkSpeed field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPublicNetworkSpeed

`func (o *Resources) SetPublicNetworkSpeed(v NetworkSpeed)`

SetPublicNetworkSpeed sets PublicNetworkSpeed field to given value.


### GetPrivateNetworkSpeed

`func (o *Resources) GetPrivateNetworkSpeed() NetworkSpeed`

GetPrivateNetworkSpeed returns the PrivateNetworkSpeed field if non-nil, zero value otherwise.

### GetPrivateNetworkSpeedOk

`func (o *Resources) GetPrivateNetworkSpeedOk() (*NetworkSpeed, bool)`

GetPrivateNetworkSpeedOk returns a tuple with the PrivateNetworkSpeed field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrivateNetworkSpeed

`func (o *Resources) SetPrivateNetworkSpeed(v NetworkSpeed)`

SetPrivateNetworkSpeed sets PrivateNetworkSpeed field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


