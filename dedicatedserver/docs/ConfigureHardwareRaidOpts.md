# ConfigureHardwareRaidOpts

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CallbackUrl** | Pointer to **string** | Url which will receive callbacks | [optional] 
**Level** | Pointer to **int32** | RAID level to apply, this value is only required if you specify a type HW | [optional] 
**NumberOfDisks** | Pointer to **int32** | The number of disks you want to apply RAID on. If not specified all disks are used. | [optional] 
**PowerCycle** | Pointer to **bool** | If set to &#x60;true&#x60;, server will be power cycled in order to complete the operation | [optional] [default to true]
**Type** | [**RaidType**](RaidType.md) |  | 

## Methods

### NewConfigureHardwareRaidOpts

`func NewConfigureHardwareRaidOpts(type_ RaidType, ) *ConfigureHardwareRaidOpts`

NewConfigureHardwareRaidOpts instantiates a new ConfigureHardwareRaidOpts object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewConfigureHardwareRaidOptsWithDefaults

`func NewConfigureHardwareRaidOptsWithDefaults() *ConfigureHardwareRaidOpts`

NewConfigureHardwareRaidOptsWithDefaults instantiates a new ConfigureHardwareRaidOpts object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCallbackUrl

`func (o *ConfigureHardwareRaidOpts) GetCallbackUrl() string`

GetCallbackUrl returns the CallbackUrl field if non-nil, zero value otherwise.

### GetCallbackUrlOk

`func (o *ConfigureHardwareRaidOpts) GetCallbackUrlOk() (*string, bool)`

GetCallbackUrlOk returns a tuple with the CallbackUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCallbackUrl

`func (o *ConfigureHardwareRaidOpts) SetCallbackUrl(v string)`

SetCallbackUrl sets CallbackUrl field to given value.

### HasCallbackUrl

`func (o *ConfigureHardwareRaidOpts) HasCallbackUrl() bool`

HasCallbackUrl returns a boolean if a field has been set.

### GetLevel

`func (o *ConfigureHardwareRaidOpts) GetLevel() int32`

GetLevel returns the Level field if non-nil, zero value otherwise.

### GetLevelOk

`func (o *ConfigureHardwareRaidOpts) GetLevelOk() (*int32, bool)`

GetLevelOk returns a tuple with the Level field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLevel

`func (o *ConfigureHardwareRaidOpts) SetLevel(v int32)`

SetLevel sets Level field to given value.

### HasLevel

`func (o *ConfigureHardwareRaidOpts) HasLevel() bool`

HasLevel returns a boolean if a field has been set.

### GetNumberOfDisks

`func (o *ConfigureHardwareRaidOpts) GetNumberOfDisks() int32`

GetNumberOfDisks returns the NumberOfDisks field if non-nil, zero value otherwise.

### GetNumberOfDisksOk

`func (o *ConfigureHardwareRaidOpts) GetNumberOfDisksOk() (*int32, bool)`

GetNumberOfDisksOk returns a tuple with the NumberOfDisks field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNumberOfDisks

`func (o *ConfigureHardwareRaidOpts) SetNumberOfDisks(v int32)`

SetNumberOfDisks sets NumberOfDisks field to given value.

### HasNumberOfDisks

`func (o *ConfigureHardwareRaidOpts) HasNumberOfDisks() bool`

HasNumberOfDisks returns a boolean if a field has been set.

### GetPowerCycle

`func (o *ConfigureHardwareRaidOpts) GetPowerCycle() bool`

GetPowerCycle returns the PowerCycle field if non-nil, zero value otherwise.

### GetPowerCycleOk

`func (o *ConfigureHardwareRaidOpts) GetPowerCycleOk() (*bool, bool)`

GetPowerCycleOk returns a tuple with the PowerCycle field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPowerCycle

`func (o *ConfigureHardwareRaidOpts) SetPowerCycle(v bool)`

SetPowerCycle sets PowerCycle field to given value.

### HasPowerCycle

`func (o *ConfigureHardwareRaidOpts) HasPowerCycle() bool`

HasPowerCycle returns a boolean if a field has been set.

### GetType

`func (o *ConfigureHardwareRaidOpts) GetType() RaidType`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *ConfigureHardwareRaidOpts) GetTypeOk() (*RaidType, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *ConfigureHardwareRaidOpts) SetType(v RaidType)`

SetType sets Type field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


