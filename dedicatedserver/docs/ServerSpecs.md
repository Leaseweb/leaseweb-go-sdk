# ServerSpecs

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Chassis** | Pointer to **string** | The chassis description of the server | [optional] 
**Cpu** | Pointer to [**Cpu**](Cpu.md) |  | [optional] 
**HardwareRaidCapable** | Pointer to **bool** | Hardware RAID capability of the server | [optional] 
**Hdd** | Pointer to [**[]Hdd**](Hdd.md) | List of hard disk drives of the server | [optional] 
**PciCards** | Pointer to [**[]PciCard**](PciCard.md) | List of PCI cards of the server | [optional] 
**Ram** | Pointer to [**Ram**](Ram.md) |  | [optional] 

## Methods

### NewServerSpecs

`func NewServerSpecs() *ServerSpecs`

NewServerSpecs instantiates a new ServerSpecs object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewServerSpecsWithDefaults

`func NewServerSpecsWithDefaults() *ServerSpecs`

NewServerSpecsWithDefaults instantiates a new ServerSpecs object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetChassis

`func (o *ServerSpecs) GetChassis() string`

GetChassis returns the Chassis field if non-nil, zero value otherwise.

### GetChassisOk

`func (o *ServerSpecs) GetChassisOk() (*string, bool)`

GetChassisOk returns a tuple with the Chassis field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetChassis

`func (o *ServerSpecs) SetChassis(v string)`

SetChassis sets Chassis field to given value.

### HasChassis

`func (o *ServerSpecs) HasChassis() bool`

HasChassis returns a boolean if a field has been set.

### GetCpu

`func (o *ServerSpecs) GetCpu() Cpu`

GetCpu returns the Cpu field if non-nil, zero value otherwise.

### GetCpuOk

`func (o *ServerSpecs) GetCpuOk() (*Cpu, bool)`

GetCpuOk returns a tuple with the Cpu field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCpu

`func (o *ServerSpecs) SetCpu(v Cpu)`

SetCpu sets Cpu field to given value.

### HasCpu

`func (o *ServerSpecs) HasCpu() bool`

HasCpu returns a boolean if a field has been set.

### GetHardwareRaidCapable

`func (o *ServerSpecs) GetHardwareRaidCapable() bool`

GetHardwareRaidCapable returns the HardwareRaidCapable field if non-nil, zero value otherwise.

### GetHardwareRaidCapableOk

`func (o *ServerSpecs) GetHardwareRaidCapableOk() (*bool, bool)`

GetHardwareRaidCapableOk returns a tuple with the HardwareRaidCapable field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHardwareRaidCapable

`func (o *ServerSpecs) SetHardwareRaidCapable(v bool)`

SetHardwareRaidCapable sets HardwareRaidCapable field to given value.

### HasHardwareRaidCapable

`func (o *ServerSpecs) HasHardwareRaidCapable() bool`

HasHardwareRaidCapable returns a boolean if a field has been set.

### GetHdd

`func (o *ServerSpecs) GetHdd() []Hdd`

GetHdd returns the Hdd field if non-nil, zero value otherwise.

### GetHddOk

`func (o *ServerSpecs) GetHddOk() (*[]Hdd, bool)`

GetHddOk returns a tuple with the Hdd field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHdd

`func (o *ServerSpecs) SetHdd(v []Hdd)`

SetHdd sets Hdd field to given value.

### HasHdd

`func (o *ServerSpecs) HasHdd() bool`

HasHdd returns a boolean if a field has been set.

### GetPciCards

`func (o *ServerSpecs) GetPciCards() []PciCard`

GetPciCards returns the PciCards field if non-nil, zero value otherwise.

### GetPciCardsOk

`func (o *ServerSpecs) GetPciCardsOk() (*[]PciCard, bool)`

GetPciCardsOk returns a tuple with the PciCards field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPciCards

`func (o *ServerSpecs) SetPciCards(v []PciCard)`

SetPciCards sets PciCards field to given value.

### HasPciCards

`func (o *ServerSpecs) HasPciCards() bool`

HasPciCards returns a boolean if a field has been set.

### GetRam

`func (o *ServerSpecs) GetRam() Ram`

GetRam returns the Ram field if non-nil, zero value otherwise.

### GetRamOk

`func (o *ServerSpecs) GetRamOk() (*Ram, bool)`

GetRamOk returns a tuple with the Ram field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRam

`func (o *ServerSpecs) SetRam(v Ram)`

SetRam sets Ram field to given value.

### HasRam

`func (o *ServerSpecs) HasRam() bool`

HasRam returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


