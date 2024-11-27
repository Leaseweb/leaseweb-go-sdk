# Payload

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**FileserverBaseUrl** | Pointer to **string** |  | [optional] 
**JobType** | **string** |  | 
**Pop** | Pointer to **string** |  | [optional] 
**PowerCycle** | Pointer to **bool** |  | [optional] 
**IsUnassignedServer** | Pointer to **bool** |  | [optional] 
**ServerId** | Pointer to **string** | Id of the server | [optional] 

## Methods

### NewPayload

`func NewPayload(jobType string, ) *Payload`

NewPayload instantiates a new Payload object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPayloadWithDefaults

`func NewPayloadWithDefaults() *Payload`

NewPayloadWithDefaults instantiates a new Payload object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetFileserverBaseUrl

`func (o *Payload) GetFileserverBaseUrl() string`

GetFileserverBaseUrl returns the FileserverBaseUrl field if non-nil, zero value otherwise.

### GetFileserverBaseUrlOk

`func (o *Payload) GetFileserverBaseUrlOk() (*string, bool)`

GetFileserverBaseUrlOk returns a tuple with the FileserverBaseUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFileserverBaseUrl

`func (o *Payload) SetFileserverBaseUrl(v string)`

SetFileserverBaseUrl sets FileserverBaseUrl field to given value.

### HasFileserverBaseUrl

`func (o *Payload) HasFileserverBaseUrl() bool`

HasFileserverBaseUrl returns a boolean if a field has been set.

### GetJobType

`func (o *Payload) GetJobType() string`

GetJobType returns the JobType field if non-nil, zero value otherwise.

### GetJobTypeOk

`func (o *Payload) GetJobTypeOk() (*string, bool)`

GetJobTypeOk returns a tuple with the JobType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetJobType

`func (o *Payload) SetJobType(v string)`

SetJobType sets JobType field to given value.


### GetPop

`func (o *Payload) GetPop() string`

GetPop returns the Pop field if non-nil, zero value otherwise.

### GetPopOk

`func (o *Payload) GetPopOk() (*string, bool)`

GetPopOk returns a tuple with the Pop field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPop

`func (o *Payload) SetPop(v string)`

SetPop sets Pop field to given value.

### HasPop

`func (o *Payload) HasPop() bool`

HasPop returns a boolean if a field has been set.

### GetPowerCycle

`func (o *Payload) GetPowerCycle() bool`

GetPowerCycle returns the PowerCycle field if non-nil, zero value otherwise.

### GetPowerCycleOk

`func (o *Payload) GetPowerCycleOk() (*bool, bool)`

GetPowerCycleOk returns a tuple with the PowerCycle field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPowerCycle

`func (o *Payload) SetPowerCycle(v bool)`

SetPowerCycle sets PowerCycle field to given value.

### HasPowerCycle

`func (o *Payload) HasPowerCycle() bool`

HasPowerCycle returns a boolean if a field has been set.

### GetIsUnassignedServer

`func (o *Payload) GetIsUnassignedServer() bool`

GetIsUnassignedServer returns the IsUnassignedServer field if non-nil, zero value otherwise.

### GetIsUnassignedServerOk

`func (o *Payload) GetIsUnassignedServerOk() (*bool, bool)`

GetIsUnassignedServerOk returns a tuple with the IsUnassignedServer field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsUnassignedServer

`func (o *Payload) SetIsUnassignedServer(v bool)`

SetIsUnassignedServer sets IsUnassignedServer field to given value.

### HasIsUnassignedServer

`func (o *Payload) HasIsUnassignedServer() bool`

HasIsUnassignedServer returns a boolean if a field has been set.

### GetServerId

`func (o *Payload) GetServerId() string`

GetServerId returns the ServerId field if non-nil, zero value otherwise.

### GetServerIdOk

`func (o *Payload) GetServerIdOk() (*string, bool)`

GetServerIdOk returns a tuple with the ServerId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServerId

`func (o *Payload) SetServerId(v string)`

SetServerId sets ServerId field to given value.

### HasServerId

`func (o *Payload) HasServerId() bool`

HasServerId returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


