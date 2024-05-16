# InstallOperatingSystemPayload

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**FileserverBaseUrl** | Pointer to **string** |  | [optional] 
**Pop** | Pointer to **string** |  | [optional] 
**PowerCycle** | Pointer to **bool** |  | [optional] 
**IsUnassignedServer** | Pointer to **bool** |  | [optional] 
**ServerId** | Pointer to **string** | Id of the server | [optional] 
**JobType** | **string** |  | 
**Configurable** | Pointer to **bool** |  | [optional] 
**Device** | Pointer to **string** |  | [optional] 
**NumberOfDisks** | Pointer to **NullableInt32** |  | [optional] 
**OperatingSystemId** | Pointer to **string** |  | [optional] 
**Os** | Pointer to [**Os**](Os.md) |  | [optional] 
**Partitions** | Pointer to [**[]Partition**](Partition.md) |  | [optional] 
**RaidLevel** | Pointer to **NullableInt32** |  | [optional] 
**Timezone** | Pointer to **string** | Timezone represented as Geographical_Area/City | [optional] 

## Methods

### NewInstallOperatingSystemPayload

`func NewInstallOperatingSystemPayload(jobType string, ) *InstallOperatingSystemPayload`

NewInstallOperatingSystemPayload instantiates a new InstallOperatingSystemPayload object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewInstallOperatingSystemPayloadWithDefaults

`func NewInstallOperatingSystemPayloadWithDefaults() *InstallOperatingSystemPayload`

NewInstallOperatingSystemPayloadWithDefaults instantiates a new InstallOperatingSystemPayload object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetFileserverBaseUrl

`func (o *InstallOperatingSystemPayload) GetFileserverBaseUrl() string`

GetFileserverBaseUrl returns the FileserverBaseUrl field if non-nil, zero value otherwise.

### GetFileserverBaseUrlOk

`func (o *InstallOperatingSystemPayload) GetFileserverBaseUrlOk() (*string, bool)`

GetFileserverBaseUrlOk returns a tuple with the FileserverBaseUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFileserverBaseUrl

`func (o *InstallOperatingSystemPayload) SetFileserverBaseUrl(v string)`

SetFileserverBaseUrl sets FileserverBaseUrl field to given value.

### HasFileserverBaseUrl

`func (o *InstallOperatingSystemPayload) HasFileserverBaseUrl() bool`

HasFileserverBaseUrl returns a boolean if a field has been set.

### GetPop

`func (o *InstallOperatingSystemPayload) GetPop() string`

GetPop returns the Pop field if non-nil, zero value otherwise.

### GetPopOk

`func (o *InstallOperatingSystemPayload) GetPopOk() (*string, bool)`

GetPopOk returns a tuple with the Pop field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPop

`func (o *InstallOperatingSystemPayload) SetPop(v string)`

SetPop sets Pop field to given value.

### HasPop

`func (o *InstallOperatingSystemPayload) HasPop() bool`

HasPop returns a boolean if a field has been set.

### GetPowerCycle

`func (o *InstallOperatingSystemPayload) GetPowerCycle() bool`

GetPowerCycle returns the PowerCycle field if non-nil, zero value otherwise.

### GetPowerCycleOk

`func (o *InstallOperatingSystemPayload) GetPowerCycleOk() (*bool, bool)`

GetPowerCycleOk returns a tuple with the PowerCycle field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPowerCycle

`func (o *InstallOperatingSystemPayload) SetPowerCycle(v bool)`

SetPowerCycle sets PowerCycle field to given value.

### HasPowerCycle

`func (o *InstallOperatingSystemPayload) HasPowerCycle() bool`

HasPowerCycle returns a boolean if a field has been set.

### GetIsUnassignedServer

`func (o *InstallOperatingSystemPayload) GetIsUnassignedServer() bool`

GetIsUnassignedServer returns the IsUnassignedServer field if non-nil, zero value otherwise.

### GetIsUnassignedServerOk

`func (o *InstallOperatingSystemPayload) GetIsUnassignedServerOk() (*bool, bool)`

GetIsUnassignedServerOk returns a tuple with the IsUnassignedServer field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsUnassignedServer

`func (o *InstallOperatingSystemPayload) SetIsUnassignedServer(v bool)`

SetIsUnassignedServer sets IsUnassignedServer field to given value.

### HasIsUnassignedServer

`func (o *InstallOperatingSystemPayload) HasIsUnassignedServer() bool`

HasIsUnassignedServer returns a boolean if a field has been set.

### GetServerId

`func (o *InstallOperatingSystemPayload) GetServerId() string`

GetServerId returns the ServerId field if non-nil, zero value otherwise.

### GetServerIdOk

`func (o *InstallOperatingSystemPayload) GetServerIdOk() (*string, bool)`

GetServerIdOk returns a tuple with the ServerId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServerId

`func (o *InstallOperatingSystemPayload) SetServerId(v string)`

SetServerId sets ServerId field to given value.

### HasServerId

`func (o *InstallOperatingSystemPayload) HasServerId() bool`

HasServerId returns a boolean if a field has been set.

### GetJobType

`func (o *InstallOperatingSystemPayload) GetJobType() string`

GetJobType returns the JobType field if non-nil, zero value otherwise.

### GetJobTypeOk

`func (o *InstallOperatingSystemPayload) GetJobTypeOk() (*string, bool)`

GetJobTypeOk returns a tuple with the JobType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetJobType

`func (o *InstallOperatingSystemPayload) SetJobType(v string)`

SetJobType sets JobType field to given value.


### GetConfigurable

`func (o *InstallOperatingSystemPayload) GetConfigurable() bool`

GetConfigurable returns the Configurable field if non-nil, zero value otherwise.

### GetConfigurableOk

`func (o *InstallOperatingSystemPayload) GetConfigurableOk() (*bool, bool)`

GetConfigurableOk returns a tuple with the Configurable field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConfigurable

`func (o *InstallOperatingSystemPayload) SetConfigurable(v bool)`

SetConfigurable sets Configurable field to given value.

### HasConfigurable

`func (o *InstallOperatingSystemPayload) HasConfigurable() bool`

HasConfigurable returns a boolean if a field has been set.

### GetDevice

`func (o *InstallOperatingSystemPayload) GetDevice() string`

GetDevice returns the Device field if non-nil, zero value otherwise.

### GetDeviceOk

`func (o *InstallOperatingSystemPayload) GetDeviceOk() (*string, bool)`

GetDeviceOk returns a tuple with the Device field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDevice

`func (o *InstallOperatingSystemPayload) SetDevice(v string)`

SetDevice sets Device field to given value.

### HasDevice

`func (o *InstallOperatingSystemPayload) HasDevice() bool`

HasDevice returns a boolean if a field has been set.

### GetNumberOfDisks

`func (o *InstallOperatingSystemPayload) GetNumberOfDisks() int32`

GetNumberOfDisks returns the NumberOfDisks field if non-nil, zero value otherwise.

### GetNumberOfDisksOk

`func (o *InstallOperatingSystemPayload) GetNumberOfDisksOk() (*int32, bool)`

GetNumberOfDisksOk returns a tuple with the NumberOfDisks field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNumberOfDisks

`func (o *InstallOperatingSystemPayload) SetNumberOfDisks(v int32)`

SetNumberOfDisks sets NumberOfDisks field to given value.

### HasNumberOfDisks

`func (o *InstallOperatingSystemPayload) HasNumberOfDisks() bool`

HasNumberOfDisks returns a boolean if a field has been set.

### SetNumberOfDisksNil

`func (o *InstallOperatingSystemPayload) SetNumberOfDisksNil(b bool)`

 SetNumberOfDisksNil sets the value for NumberOfDisks to be an explicit nil

### UnsetNumberOfDisks
`func (o *InstallOperatingSystemPayload) UnsetNumberOfDisks()`

UnsetNumberOfDisks ensures that no value is present for NumberOfDisks, not even an explicit nil
### GetOperatingSystemId

`func (o *InstallOperatingSystemPayload) GetOperatingSystemId() string`

GetOperatingSystemId returns the OperatingSystemId field if non-nil, zero value otherwise.

### GetOperatingSystemIdOk

`func (o *InstallOperatingSystemPayload) GetOperatingSystemIdOk() (*string, bool)`

GetOperatingSystemIdOk returns a tuple with the OperatingSystemId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOperatingSystemId

`func (o *InstallOperatingSystemPayload) SetOperatingSystemId(v string)`

SetOperatingSystemId sets OperatingSystemId field to given value.

### HasOperatingSystemId

`func (o *InstallOperatingSystemPayload) HasOperatingSystemId() bool`

HasOperatingSystemId returns a boolean if a field has been set.

### GetOs

`func (o *InstallOperatingSystemPayload) GetOs() Os`

GetOs returns the Os field if non-nil, zero value otherwise.

### GetOsOk

`func (o *InstallOperatingSystemPayload) GetOsOk() (*Os, bool)`

GetOsOk returns a tuple with the Os field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOs

`func (o *InstallOperatingSystemPayload) SetOs(v Os)`

SetOs sets Os field to given value.

### HasOs

`func (o *InstallOperatingSystemPayload) HasOs() bool`

HasOs returns a boolean if a field has been set.

### GetPartitions

`func (o *InstallOperatingSystemPayload) GetPartitions() []Partition`

GetPartitions returns the Partitions field if non-nil, zero value otherwise.

### GetPartitionsOk

`func (o *InstallOperatingSystemPayload) GetPartitionsOk() (*[]Partition, bool)`

GetPartitionsOk returns a tuple with the Partitions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPartitions

`func (o *InstallOperatingSystemPayload) SetPartitions(v []Partition)`

SetPartitions sets Partitions field to given value.

### HasPartitions

`func (o *InstallOperatingSystemPayload) HasPartitions() bool`

HasPartitions returns a boolean if a field has been set.

### GetRaidLevel

`func (o *InstallOperatingSystemPayload) GetRaidLevel() int32`

GetRaidLevel returns the RaidLevel field if non-nil, zero value otherwise.

### GetRaidLevelOk

`func (o *InstallOperatingSystemPayload) GetRaidLevelOk() (*int32, bool)`

GetRaidLevelOk returns a tuple with the RaidLevel field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRaidLevel

`func (o *InstallOperatingSystemPayload) SetRaidLevel(v int32)`

SetRaidLevel sets RaidLevel field to given value.

### HasRaidLevel

`func (o *InstallOperatingSystemPayload) HasRaidLevel() bool`

HasRaidLevel returns a boolean if a field has been set.

### SetRaidLevelNil

`func (o *InstallOperatingSystemPayload) SetRaidLevelNil(b bool)`

 SetRaidLevelNil sets the value for RaidLevel to be an explicit nil

### UnsetRaidLevel
`func (o *InstallOperatingSystemPayload) UnsetRaidLevel()`

UnsetRaidLevel ensures that no value is present for RaidLevel, not even an explicit nil
### GetTimezone

`func (o *InstallOperatingSystemPayload) GetTimezone() string`

GetTimezone returns the Timezone field if non-nil, zero value otherwise.

### GetTimezoneOk

`func (o *InstallOperatingSystemPayload) GetTimezoneOk() (*string, bool)`

GetTimezoneOk returns a tuple with the Timezone field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimezone

`func (o *InstallOperatingSystemPayload) SetTimezone(v string)`

SetTimezone sets Timezone field to given value.

### HasTimezone

`func (o *InstallOperatingSystemPayload) HasTimezone() bool`

HasTimezone returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


