# BaseResources

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Cpu** | [**Cpu**](Cpu.md) |  | 
**Memory** | [**Memory**](Memory.md) |  | 
**PublicNetworkSpeed** | [**NetworkSpeed**](NetworkSpeed.md) |  | 

## Methods

### NewBaseResources

`func NewBaseResources(cpu Cpu, memory Memory, publicNetworkSpeed NetworkSpeed, ) *BaseResources`

NewBaseResources instantiates a new BaseResources object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewBaseResourcesWithDefaults

`func NewBaseResourcesWithDefaults() *BaseResources`

NewBaseResourcesWithDefaults instantiates a new BaseResources object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCpu

`func (o *BaseResources) GetCpu() Cpu`

GetCpu returns the Cpu field if non-nil, zero value otherwise.

### GetCpuOk

`func (o *BaseResources) GetCpuOk() (*Cpu, bool)`

GetCpuOk returns a tuple with the Cpu field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCpu

`func (o *BaseResources) SetCpu(v Cpu)`

SetCpu sets Cpu field to given value.


### GetMemory

`func (o *BaseResources) GetMemory() Memory`

GetMemory returns the Memory field if non-nil, zero value otherwise.

### GetMemoryOk

`func (o *BaseResources) GetMemoryOk() (*Memory, bool)`

GetMemoryOk returns a tuple with the Memory field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMemory

`func (o *BaseResources) SetMemory(v Memory)`

SetMemory sets Memory field to given value.


### GetPublicNetworkSpeed

`func (o *BaseResources) GetPublicNetworkSpeed() NetworkSpeed`

GetPublicNetworkSpeed returns the PublicNetworkSpeed field if non-nil, zero value otherwise.

### GetPublicNetworkSpeedOk

`func (o *BaseResources) GetPublicNetworkSpeedOk() (*NetworkSpeed, bool)`

GetPublicNetworkSpeedOk returns a tuple with the PublicNetworkSpeed field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPublicNetworkSpeed

`func (o *BaseResources) SetPublicNetworkSpeed(v NetworkSpeed)`

SetPublicNetworkSpeed sets PublicNetworkSpeed field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


