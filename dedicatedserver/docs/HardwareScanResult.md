# HardwareScanResult

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Chassis** | Pointer to [**HardwareScanChassis**](HardwareScanChassis.md) |  | [optional] 
**Cpu** | Pointer to [**[]HardwareScanCpu**](HardwareScanCpu.md) |  | [optional] 
**Disks** | Pointer to [**[]HardwareScanDisk**](HardwareScanDisk.md) |  | [optional] 
**Ipmi** | Pointer to [**HardwareScanResultIpmi**](HardwareScanResultIpmi.md) |  | [optional] 
**Memory** | Pointer to [**[]HardwareScanMemory**](HardwareScanMemory.md) |  | [optional] 
**Network** | Pointer to [**[]HardwareScanResultNetwork**](HardwareScanResultNetwork.md) |  | [optional] 

## Methods

### NewHardwareScanResult

`func NewHardwareScanResult() *HardwareScanResult`

NewHardwareScanResult instantiates a new HardwareScanResult object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewHardwareScanResultWithDefaults

`func NewHardwareScanResultWithDefaults() *HardwareScanResult`

NewHardwareScanResultWithDefaults instantiates a new HardwareScanResult object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetChassis

`func (o *HardwareScanResult) GetChassis() HardwareScanChassis`

GetChassis returns the Chassis field if non-nil, zero value otherwise.

### GetChassisOk

`func (o *HardwareScanResult) GetChassisOk() (*HardwareScanChassis, bool)`

GetChassisOk returns a tuple with the Chassis field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetChassis

`func (o *HardwareScanResult) SetChassis(v HardwareScanChassis)`

SetChassis sets Chassis field to given value.

### HasChassis

`func (o *HardwareScanResult) HasChassis() bool`

HasChassis returns a boolean if a field has been set.

### GetCpu

`func (o *HardwareScanResult) GetCpu() []HardwareScanCpu`

GetCpu returns the Cpu field if non-nil, zero value otherwise.

### GetCpuOk

`func (o *HardwareScanResult) GetCpuOk() (*[]HardwareScanCpu, bool)`

GetCpuOk returns a tuple with the Cpu field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCpu

`func (o *HardwareScanResult) SetCpu(v []HardwareScanCpu)`

SetCpu sets Cpu field to given value.

### HasCpu

`func (o *HardwareScanResult) HasCpu() bool`

HasCpu returns a boolean if a field has been set.

### GetDisks

`func (o *HardwareScanResult) GetDisks() []HardwareScanDisk`

GetDisks returns the Disks field if non-nil, zero value otherwise.

### GetDisksOk

`func (o *HardwareScanResult) GetDisksOk() (*[]HardwareScanDisk, bool)`

GetDisksOk returns a tuple with the Disks field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisks

`func (o *HardwareScanResult) SetDisks(v []HardwareScanDisk)`

SetDisks sets Disks field to given value.

### HasDisks

`func (o *HardwareScanResult) HasDisks() bool`

HasDisks returns a boolean if a field has been set.

### GetIpmi

`func (o *HardwareScanResult) GetIpmi() HardwareScanResultIpmi`

GetIpmi returns the Ipmi field if non-nil, zero value otherwise.

### GetIpmiOk

`func (o *HardwareScanResult) GetIpmiOk() (*HardwareScanResultIpmi, bool)`

GetIpmiOk returns a tuple with the Ipmi field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIpmi

`func (o *HardwareScanResult) SetIpmi(v HardwareScanResultIpmi)`

SetIpmi sets Ipmi field to given value.

### HasIpmi

`func (o *HardwareScanResult) HasIpmi() bool`

HasIpmi returns a boolean if a field has been set.

### GetMemory

`func (o *HardwareScanResult) GetMemory() []HardwareScanMemory`

GetMemory returns the Memory field if non-nil, zero value otherwise.

### GetMemoryOk

`func (o *HardwareScanResult) GetMemoryOk() (*[]HardwareScanMemory, bool)`

GetMemoryOk returns a tuple with the Memory field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMemory

`func (o *HardwareScanResult) SetMemory(v []HardwareScanMemory)`

SetMemory sets Memory field to given value.

### HasMemory

`func (o *HardwareScanResult) HasMemory() bool`

HasMemory returns a boolean if a field has been set.

### GetNetwork

`func (o *HardwareScanResult) GetNetwork() []HardwareScanResultNetwork`

GetNetwork returns the Network field if non-nil, zero value otherwise.

### GetNetworkOk

`func (o *HardwareScanResult) GetNetworkOk() (*[]HardwareScanResultNetwork, bool)`

GetNetworkOk returns a tuple with the Network field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNetwork

`func (o *HardwareScanResult) SetNetwork(v []HardwareScanResultNetwork)`

SetNetwork sets Network field to given value.

### HasNetwork

`func (o *HardwareScanResult) HasNetwork() bool`

HasNetwork returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


