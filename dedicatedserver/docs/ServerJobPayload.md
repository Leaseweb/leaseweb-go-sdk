# ServerJobPayload

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

### NewServerJobPayload

`func NewServerJobPayload(jobType string, ) *ServerJobPayload`

NewServerJobPayload instantiates a new ServerJobPayload object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewServerJobPayloadWithDefaults

`func NewServerJobPayloadWithDefaults() *ServerJobPayload`

NewServerJobPayloadWithDefaults instantiates a new ServerJobPayload object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetFileserverBaseUrl

`func (o *ServerJobPayload) GetFileserverBaseUrl() string`

GetFileserverBaseUrl returns the FileserverBaseUrl field if non-nil, zero value otherwise.

### GetFileserverBaseUrlOk

`func (o *ServerJobPayload) GetFileserverBaseUrlOk() (*string, bool)`

GetFileserverBaseUrlOk returns a tuple with the FileserverBaseUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFileserverBaseUrl

`func (o *ServerJobPayload) SetFileserverBaseUrl(v string)`

SetFileserverBaseUrl sets FileserverBaseUrl field to given value.

### HasFileserverBaseUrl

`func (o *ServerJobPayload) HasFileserverBaseUrl() bool`

HasFileserverBaseUrl returns a boolean if a field has been set.

### GetPop

`func (o *ServerJobPayload) GetPop() string`

GetPop returns the Pop field if non-nil, zero value otherwise.

### GetPopOk

`func (o *ServerJobPayload) GetPopOk() (*string, bool)`

GetPopOk returns a tuple with the Pop field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPop

`func (o *ServerJobPayload) SetPop(v string)`

SetPop sets Pop field to given value.

### HasPop

`func (o *ServerJobPayload) HasPop() bool`

HasPop returns a boolean if a field has been set.

### GetPowerCycle

`func (o *ServerJobPayload) GetPowerCycle() bool`

GetPowerCycle returns the PowerCycle field if non-nil, zero value otherwise.

### GetPowerCycleOk

`func (o *ServerJobPayload) GetPowerCycleOk() (*bool, bool)`

GetPowerCycleOk returns a tuple with the PowerCycle field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPowerCycle

`func (o *ServerJobPayload) SetPowerCycle(v bool)`

SetPowerCycle sets PowerCycle field to given value.

### HasPowerCycle

`func (o *ServerJobPayload) HasPowerCycle() bool`

HasPowerCycle returns a boolean if a field has been set.

### GetIsUnassignedServer

`func (o *ServerJobPayload) GetIsUnassignedServer() bool`

GetIsUnassignedServer returns the IsUnassignedServer field if non-nil, zero value otherwise.

### GetIsUnassignedServerOk

`func (o *ServerJobPayload) GetIsUnassignedServerOk() (*bool, bool)`

GetIsUnassignedServerOk returns a tuple with the IsUnassignedServer field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsUnassignedServer

`func (o *ServerJobPayload) SetIsUnassignedServer(v bool)`

SetIsUnassignedServer sets IsUnassignedServer field to given value.

### HasIsUnassignedServer

`func (o *ServerJobPayload) HasIsUnassignedServer() bool`

HasIsUnassignedServer returns a boolean if a field has been set.

### GetServerId

`func (o *ServerJobPayload) GetServerId() string`

GetServerId returns the ServerId field if non-nil, zero value otherwise.

### GetServerIdOk

`func (o *ServerJobPayload) GetServerIdOk() (*string, bool)`

GetServerIdOk returns a tuple with the ServerId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServerId

`func (o *ServerJobPayload) SetServerId(v string)`

SetServerId sets ServerId field to given value.

### HasServerId

`func (o *ServerJobPayload) HasServerId() bool`

HasServerId returns a boolean if a field has been set.

### GetJobType

`func (o *ServerJobPayload) GetJobType() string`

GetJobType returns the JobType field if non-nil, zero value otherwise.

### GetJobTypeOk

`func (o *ServerJobPayload) GetJobTypeOk() (*string, bool)`

GetJobTypeOk returns a tuple with the JobType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetJobType

`func (o *ServerJobPayload) SetJobType(v string)`

SetJobType sets JobType field to given value.


### GetConfigurable

`func (o *ServerJobPayload) GetConfigurable() bool`

GetConfigurable returns the Configurable field if non-nil, zero value otherwise.

### GetConfigurableOk

`func (o *ServerJobPayload) GetConfigurableOk() (*bool, bool)`

GetConfigurableOk returns a tuple with the Configurable field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConfigurable

`func (o *ServerJobPayload) SetConfigurable(v bool)`

SetConfigurable sets Configurable field to given value.

### HasConfigurable

`func (o *ServerJobPayload) HasConfigurable() bool`

HasConfigurable returns a boolean if a field has been set.

### GetDevice

`func (o *ServerJobPayload) GetDevice() string`

GetDevice returns the Device field if non-nil, zero value otherwise.

### GetDeviceOk

`func (o *ServerJobPayload) GetDeviceOk() (*string, bool)`

GetDeviceOk returns a tuple with the Device field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDevice

`func (o *ServerJobPayload) SetDevice(v string)`

SetDevice sets Device field to given value.

### HasDevice

`func (o *ServerJobPayload) HasDevice() bool`

HasDevice returns a boolean if a field has been set.

### GetNumberOfDisks

`func (o *ServerJobPayload) GetNumberOfDisks() int32`

GetNumberOfDisks returns the NumberOfDisks field if non-nil, zero value otherwise.

### GetNumberOfDisksOk

`func (o *ServerJobPayload) GetNumberOfDisksOk() (*int32, bool)`

GetNumberOfDisksOk returns a tuple with the NumberOfDisks field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNumberOfDisks

`func (o *ServerJobPayload) SetNumberOfDisks(v int32)`

SetNumberOfDisks sets NumberOfDisks field to given value.

### HasNumberOfDisks

`func (o *ServerJobPayload) HasNumberOfDisks() bool`

HasNumberOfDisks returns a boolean if a field has been set.

### SetNumberOfDisksNil

`func (o *ServerJobPayload) SetNumberOfDisksNil(b bool)`

 SetNumberOfDisksNil sets the value for NumberOfDisks to be an explicit nil

### UnsetNumberOfDisks
`func (o *ServerJobPayload) UnsetNumberOfDisks()`

UnsetNumberOfDisks ensures that no value is present for NumberOfDisks, not even an explicit nil
### GetOperatingSystemId

`func (o *ServerJobPayload) GetOperatingSystemId() string`

GetOperatingSystemId returns the OperatingSystemId field if non-nil, zero value otherwise.

### GetOperatingSystemIdOk

`func (o *ServerJobPayload) GetOperatingSystemIdOk() (*string, bool)`

GetOperatingSystemIdOk returns a tuple with the OperatingSystemId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOperatingSystemId

`func (o *ServerJobPayload) SetOperatingSystemId(v string)`

SetOperatingSystemId sets OperatingSystemId field to given value.

### HasOperatingSystemId

`func (o *ServerJobPayload) HasOperatingSystemId() bool`

HasOperatingSystemId returns a boolean if a field has been set.

### GetOs

`func (o *ServerJobPayload) GetOs() Os`

GetOs returns the Os field if non-nil, zero value otherwise.

### GetOsOk

`func (o *ServerJobPayload) GetOsOk() (*Os, bool)`

GetOsOk returns a tuple with the Os field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOs

`func (o *ServerJobPayload) SetOs(v Os)`

SetOs sets Os field to given value.

### HasOs

`func (o *ServerJobPayload) HasOs() bool`

HasOs returns a boolean if a field has been set.

### GetPartitions

`func (o *ServerJobPayload) GetPartitions() []Partition`

GetPartitions returns the Partitions field if non-nil, zero value otherwise.

### GetPartitionsOk

`func (o *ServerJobPayload) GetPartitionsOk() (*[]Partition, bool)`

GetPartitionsOk returns a tuple with the Partitions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPartitions

`func (o *ServerJobPayload) SetPartitions(v []Partition)`

SetPartitions sets Partitions field to given value.

### HasPartitions

`func (o *ServerJobPayload) HasPartitions() bool`

HasPartitions returns a boolean if a field has been set.

### GetRaidLevel

`func (o *ServerJobPayload) GetRaidLevel() int32`

GetRaidLevel returns the RaidLevel field if non-nil, zero value otherwise.

### GetRaidLevelOk

`func (o *ServerJobPayload) GetRaidLevelOk() (*int32, bool)`

GetRaidLevelOk returns a tuple with the RaidLevel field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRaidLevel

`func (o *ServerJobPayload) SetRaidLevel(v int32)`

SetRaidLevel sets RaidLevel field to given value.

### HasRaidLevel

`func (o *ServerJobPayload) HasRaidLevel() bool`

HasRaidLevel returns a boolean if a field has been set.

### SetRaidLevelNil

`func (o *ServerJobPayload) SetRaidLevelNil(b bool)`

 SetRaidLevelNil sets the value for RaidLevel to be an explicit nil

### UnsetRaidLevel
`func (o *ServerJobPayload) UnsetRaidLevel()`

UnsetRaidLevel ensures that no value is present for RaidLevel, not even an explicit nil
### GetTimezone

`func (o *ServerJobPayload) GetTimezone() string`

GetTimezone returns the Timezone field if non-nil, zero value otherwise.

### GetTimezoneOk

`func (o *ServerJobPayload) GetTimezoneOk() (*string, bool)`

GetTimezoneOk returns a tuple with the Timezone field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimezone

`func (o *ServerJobPayload) SetTimezone(v string)`

SetTimezone sets Timezone field to given value.

### HasTimezone

`func (o *ServerJobPayload) HasTimezone() bool`

HasTimezone returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


