# InstallationJob

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**BareMetalId** | Pointer to **string** | Id of the server (deprecated, use serverId instead) | [optional] 
**CreatedAt** | Pointer to **time.Time** | Creation timestamp | [optional] 
**Flow** | Pointer to **string** | Job flow | [optional] 
**IsRunning** | Pointer to **bool** | Describes whether the job is running | [optional] 
**Node** | Pointer to **string** | Node ID for this server | [optional] 
**Payload** | Pointer to [**InstallationJobPayload**](InstallationJobPayload.md) |  | [optional] 
**Progress** | Pointer to [**Progress**](Progress.md) |  | [optional] 
**ServerId** | Pointer to **string** | Id of the server | [optional] 
**Status** | Pointer to **string** | Status of the job | [optional] 
**Tasks** | Pointer to [**[]Task**](Task.md) | List of tasks in the job | [optional] 
**Type** | Pointer to [**JobType**](JobType.md) |  | [optional] 
**UpdatedAt** | Pointer to **time.Time** | Update timestamp | [optional] 
**Uuid** | Pointer to **string** | Unique identifier of the job | [optional] 

## Methods

### NewInstallationJob

`func NewInstallationJob() *InstallationJob`

NewInstallationJob instantiates a new InstallationJob object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewInstallationJobWithDefaults

`func NewInstallationJobWithDefaults() *InstallationJob`

NewInstallationJobWithDefaults instantiates a new InstallationJob object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetBareMetalId

`func (o *InstallationJob) GetBareMetalId() string`

GetBareMetalId returns the BareMetalId field if non-nil, zero value otherwise.

### GetBareMetalIdOk

`func (o *InstallationJob) GetBareMetalIdOk() (*string, bool)`

GetBareMetalIdOk returns a tuple with the BareMetalId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBareMetalId

`func (o *InstallationJob) SetBareMetalId(v string)`

SetBareMetalId sets BareMetalId field to given value.

### HasBareMetalId

`func (o *InstallationJob) HasBareMetalId() bool`

HasBareMetalId returns a boolean if a field has been set.

### GetCreatedAt

`func (o *InstallationJob) GetCreatedAt() time.Time`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *InstallationJob) GetCreatedAtOk() (*time.Time, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *InstallationJob) SetCreatedAt(v time.Time)`

SetCreatedAt sets CreatedAt field to given value.

### HasCreatedAt

`func (o *InstallationJob) HasCreatedAt() bool`

HasCreatedAt returns a boolean if a field has been set.

### GetFlow

`func (o *InstallationJob) GetFlow() string`

GetFlow returns the Flow field if non-nil, zero value otherwise.

### GetFlowOk

`func (o *InstallationJob) GetFlowOk() (*string, bool)`

GetFlowOk returns a tuple with the Flow field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFlow

`func (o *InstallationJob) SetFlow(v string)`

SetFlow sets Flow field to given value.

### HasFlow

`func (o *InstallationJob) HasFlow() bool`

HasFlow returns a boolean if a field has been set.

### GetIsRunning

`func (o *InstallationJob) GetIsRunning() bool`

GetIsRunning returns the IsRunning field if non-nil, zero value otherwise.

### GetIsRunningOk

`func (o *InstallationJob) GetIsRunningOk() (*bool, bool)`

GetIsRunningOk returns a tuple with the IsRunning field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsRunning

`func (o *InstallationJob) SetIsRunning(v bool)`

SetIsRunning sets IsRunning field to given value.

### HasIsRunning

`func (o *InstallationJob) HasIsRunning() bool`

HasIsRunning returns a boolean if a field has been set.

### GetNode

`func (o *InstallationJob) GetNode() string`

GetNode returns the Node field if non-nil, zero value otherwise.

### GetNodeOk

`func (o *InstallationJob) GetNodeOk() (*string, bool)`

GetNodeOk returns a tuple with the Node field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNode

`func (o *InstallationJob) SetNode(v string)`

SetNode sets Node field to given value.

### HasNode

`func (o *InstallationJob) HasNode() bool`

HasNode returns a boolean if a field has been set.

### GetPayload

`func (o *InstallationJob) GetPayload() InstallationJobPayload`

GetPayload returns the Payload field if non-nil, zero value otherwise.

### GetPayloadOk

`func (o *InstallationJob) GetPayloadOk() (*InstallationJobPayload, bool)`

GetPayloadOk returns a tuple with the Payload field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPayload

`func (o *InstallationJob) SetPayload(v InstallationJobPayload)`

SetPayload sets Payload field to given value.

### HasPayload

`func (o *InstallationJob) HasPayload() bool`

HasPayload returns a boolean if a field has been set.

### GetProgress

`func (o *InstallationJob) GetProgress() Progress`

GetProgress returns the Progress field if non-nil, zero value otherwise.

### GetProgressOk

`func (o *InstallationJob) GetProgressOk() (*Progress, bool)`

GetProgressOk returns a tuple with the Progress field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProgress

`func (o *InstallationJob) SetProgress(v Progress)`

SetProgress sets Progress field to given value.

### HasProgress

`func (o *InstallationJob) HasProgress() bool`

HasProgress returns a boolean if a field has been set.

### GetServerId

`func (o *InstallationJob) GetServerId() string`

GetServerId returns the ServerId field if non-nil, zero value otherwise.

### GetServerIdOk

`func (o *InstallationJob) GetServerIdOk() (*string, bool)`

GetServerIdOk returns a tuple with the ServerId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServerId

`func (o *InstallationJob) SetServerId(v string)`

SetServerId sets ServerId field to given value.

### HasServerId

`func (o *InstallationJob) HasServerId() bool`

HasServerId returns a boolean if a field has been set.

### GetStatus

`func (o *InstallationJob) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *InstallationJob) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *InstallationJob) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *InstallationJob) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetTasks

`func (o *InstallationJob) GetTasks() []Task`

GetTasks returns the Tasks field if non-nil, zero value otherwise.

### GetTasksOk

`func (o *InstallationJob) GetTasksOk() (*[]Task, bool)`

GetTasksOk returns a tuple with the Tasks field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTasks

`func (o *InstallationJob) SetTasks(v []Task)`

SetTasks sets Tasks field to given value.

### HasTasks

`func (o *InstallationJob) HasTasks() bool`

HasTasks returns a boolean if a field has been set.

### GetType

`func (o *InstallationJob) GetType() JobType`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *InstallationJob) GetTypeOk() (*JobType, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *InstallationJob) SetType(v JobType)`

SetType sets Type field to given value.

### HasType

`func (o *InstallationJob) HasType() bool`

HasType returns a boolean if a field has been set.

### GetUpdatedAt

`func (o *InstallationJob) GetUpdatedAt() time.Time`

GetUpdatedAt returns the UpdatedAt field if non-nil, zero value otherwise.

### GetUpdatedAtOk

`func (o *InstallationJob) GetUpdatedAtOk() (*time.Time, bool)`

GetUpdatedAtOk returns a tuple with the UpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedAt

`func (o *InstallationJob) SetUpdatedAt(v time.Time)`

SetUpdatedAt sets UpdatedAt field to given value.

### HasUpdatedAt

`func (o *InstallationJob) HasUpdatedAt() bool`

HasUpdatedAt returns a boolean if a field has been set.

### GetUuid

`func (o *InstallationJob) GetUuid() string`

GetUuid returns the Uuid field if non-nil, zero value otherwise.

### GetUuidOk

`func (o *InstallationJob) GetUuidOk() (*string, bool)`

GetUuidOk returns a tuple with the Uuid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUuid

`func (o *InstallationJob) SetUuid(v string)`

SetUuid sets Uuid field to given value.

### HasUuid

`func (o *InstallationJob) HasUuid() bool`

HasUuid returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


