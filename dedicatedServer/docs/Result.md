# Result

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Chassis** | Pointer to [**Chassis**](Chassis.md) |  | [optional] 
**Cpu** | Pointer to [**[]Cpu1**](Cpu1.md) |  | [optional] 
**Disks** | Pointer to [**[]Disk**](Disk.md) |  | [optional] 
**Ipmi** | Pointer to [**Ipmi1**](Ipmi1.md) |  | [optional] 
**Memory** | Pointer to [**[]MemoryBank**](MemoryBank.md) |  | [optional] 
**Network** | Pointer to [**[]Network**](Network.md) |  | [optional] 

## Methods

### NewResult

`func NewResult() *Result`

NewResult instantiates a new Result object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewResultWithDefaults

`func NewResultWithDefaults() *Result`

NewResultWithDefaults instantiates a new Result object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetChassis

`func (o *Result) GetChassis() Chassis`

GetChassis returns the Chassis field if non-nil, zero value otherwise.

### GetChassisOk

`func (o *Result) GetChassisOk() (*Chassis, bool)`

GetChassisOk returns a tuple with the Chassis field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetChassis

`func (o *Result) SetChassis(v Chassis)`

SetChassis sets Chassis field to given value.

### HasChassis

`func (o *Result) HasChassis() bool`

HasChassis returns a boolean if a field has been set.

### GetCpu

`func (o *Result) GetCpu() []Cpu1`

GetCpu returns the Cpu field if non-nil, zero value otherwise.

### GetCpuOk

`func (o *Result) GetCpuOk() (*[]Cpu1, bool)`

GetCpuOk returns a tuple with the Cpu field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCpu

`func (o *Result) SetCpu(v []Cpu1)`

SetCpu sets Cpu field to given value.

### HasCpu

`func (o *Result) HasCpu() bool`

HasCpu returns a boolean if a field has been set.

### GetDisks

`func (o *Result) GetDisks() []Disk`

GetDisks returns the Disks field if non-nil, zero value otherwise.

### GetDisksOk

`func (o *Result) GetDisksOk() (*[]Disk, bool)`

GetDisksOk returns a tuple with the Disks field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisks

`func (o *Result) SetDisks(v []Disk)`

SetDisks sets Disks field to given value.

### HasDisks

`func (o *Result) HasDisks() bool`

HasDisks returns a boolean if a field has been set.

### GetIpmi

`func (o *Result) GetIpmi() Ipmi1`

GetIpmi returns the Ipmi field if non-nil, zero value otherwise.

### GetIpmiOk

`func (o *Result) GetIpmiOk() (*Ipmi1, bool)`

GetIpmiOk returns a tuple with the Ipmi field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIpmi

`func (o *Result) SetIpmi(v Ipmi1)`

SetIpmi sets Ipmi field to given value.

### HasIpmi

`func (o *Result) HasIpmi() bool`

HasIpmi returns a boolean if a field has been set.

### GetMemory

`func (o *Result) GetMemory() []MemoryBank`

GetMemory returns the Memory field if non-nil, zero value otherwise.

### GetMemoryOk

`func (o *Result) GetMemoryOk() (*[]MemoryBank, bool)`

GetMemoryOk returns a tuple with the Memory field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMemory

`func (o *Result) SetMemory(v []MemoryBank)`

SetMemory sets Memory field to given value.

### HasMemory

`func (o *Result) HasMemory() bool`

HasMemory returns a boolean if a field has been set.

### GetNetwork

`func (o *Result) GetNetwork() []Network`

GetNetwork returns the Network field if non-nil, zero value otherwise.

### GetNetworkOk

`func (o *Result) GetNetworkOk() (*[]Network, bool)`

GetNetworkOk returns a tuple with the Network field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNetwork

`func (o *Result) SetNetwork(v []Network)`

SetNetwork sets Network field to given value.

### HasNetwork

`func (o *Result) HasNetwork() bool`

HasNetwork returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


