# ServerJob

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**BareMetalId** | Pointer to **string** | Id of the server (deprecated, use serverId instead) | [optional] 
**CreatedAt** | Pointer to **time.Time** | Creation timestamp | [optional] 
**Flow** | Pointer to **string** | Job flow | [optional] 
**IsRunning** | Pointer to **bool** | Describes whether the job is running | [optional] 
**Node** | Pointer to **string** | Node ID for this server | [optional] 
**Payload** | Pointer to [**ServerJobPayload**](ServerJobPayload.md) |  | [optional] 
**Progress** | Pointer to [**Progress**](Progress.md) |  | [optional] 
**ServerId** | Pointer to **string** | Id of the server | [optional] 
**Status** | Pointer to **string** | Status of the job | [optional] 
**Tasks** | Pointer to [**[]Task**](Task.md) | List of tasks in the job | [optional] 
**Type** | Pointer to [**JobType**](JobType.md) |  | [optional] 
**UpdatedAt** | Pointer to **time.Time** | Update timestamp | [optional] 
**Uuid** | Pointer to **string** | Unique identifier of the job | [optional] 

## Methods

### NewServerJob

`func NewServerJob() *ServerJob`

NewServerJob instantiates a new ServerJob object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewServerJobWithDefaults

`func NewServerJobWithDefaults() *ServerJob`

NewServerJobWithDefaults instantiates a new ServerJob object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetBareMetalId

`func (o *ServerJob) GetBareMetalId() string`

GetBareMetalId returns the BareMetalId field if non-nil, zero value otherwise.

### GetBareMetalIdOk

`func (o *ServerJob) GetBareMetalIdOk() (*string, bool)`

GetBareMetalIdOk returns a tuple with the BareMetalId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBareMetalId

`func (o *ServerJob) SetBareMetalId(v string)`

SetBareMetalId sets BareMetalId field to given value.

### HasBareMetalId

`func (o *ServerJob) HasBareMetalId() bool`

HasBareMetalId returns a boolean if a field has been set.

### GetCreatedAt

`func (o *ServerJob) GetCreatedAt() time.Time`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *ServerJob) GetCreatedAtOk() (*time.Time, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *ServerJob) SetCreatedAt(v time.Time)`

SetCreatedAt sets CreatedAt field to given value.

### HasCreatedAt

`func (o *ServerJob) HasCreatedAt() bool`

HasCreatedAt returns a boolean if a field has been set.

### GetFlow

`func (o *ServerJob) GetFlow() string`

GetFlow returns the Flow field if non-nil, zero value otherwise.

### GetFlowOk

`func (o *ServerJob) GetFlowOk() (*string, bool)`

GetFlowOk returns a tuple with the Flow field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFlow

`func (o *ServerJob) SetFlow(v string)`

SetFlow sets Flow field to given value.

### HasFlow

`func (o *ServerJob) HasFlow() bool`

HasFlow returns a boolean if a field has been set.

### GetIsRunning

`func (o *ServerJob) GetIsRunning() bool`

GetIsRunning returns the IsRunning field if non-nil, zero value otherwise.

### GetIsRunningOk

`func (o *ServerJob) GetIsRunningOk() (*bool, bool)`

GetIsRunningOk returns a tuple with the IsRunning field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsRunning

`func (o *ServerJob) SetIsRunning(v bool)`

SetIsRunning sets IsRunning field to given value.

### HasIsRunning

`func (o *ServerJob) HasIsRunning() bool`

HasIsRunning returns a boolean if a field has been set.

### GetNode

`func (o *ServerJob) GetNode() string`

GetNode returns the Node field if non-nil, zero value otherwise.

### GetNodeOk

`func (o *ServerJob) GetNodeOk() (*string, bool)`

GetNodeOk returns a tuple with the Node field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNode

`func (o *ServerJob) SetNode(v string)`

SetNode sets Node field to given value.

### HasNode

`func (o *ServerJob) HasNode() bool`

HasNode returns a boolean if a field has been set.

### GetPayload

`func (o *ServerJob) GetPayload() ServerJobPayload`

GetPayload returns the Payload field if non-nil, zero value otherwise.

### GetPayloadOk

`func (o *ServerJob) GetPayloadOk() (*ServerJobPayload, bool)`

GetPayloadOk returns a tuple with the Payload field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPayload

`func (o *ServerJob) SetPayload(v ServerJobPayload)`

SetPayload sets Payload field to given value.

### HasPayload

`func (o *ServerJob) HasPayload() bool`

HasPayload returns a boolean if a field has been set.

### GetProgress

`func (o *ServerJob) GetProgress() Progress`

GetProgress returns the Progress field if non-nil, zero value otherwise.

### GetProgressOk

`func (o *ServerJob) GetProgressOk() (*Progress, bool)`

GetProgressOk returns a tuple with the Progress field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProgress

`func (o *ServerJob) SetProgress(v Progress)`

SetProgress sets Progress field to given value.

### HasProgress

`func (o *ServerJob) HasProgress() bool`

HasProgress returns a boolean if a field has been set.

### GetServerId

`func (o *ServerJob) GetServerId() string`

GetServerId returns the ServerId field if non-nil, zero value otherwise.

### GetServerIdOk

`func (o *ServerJob) GetServerIdOk() (*string, bool)`

GetServerIdOk returns a tuple with the ServerId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServerId

`func (o *ServerJob) SetServerId(v string)`

SetServerId sets ServerId field to given value.

### HasServerId

`func (o *ServerJob) HasServerId() bool`

HasServerId returns a boolean if a field has been set.

### GetStatus

`func (o *ServerJob) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *ServerJob) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *ServerJob) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *ServerJob) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetTasks

`func (o *ServerJob) GetTasks() []Task`

GetTasks returns the Tasks field if non-nil, zero value otherwise.

### GetTasksOk

`func (o *ServerJob) GetTasksOk() (*[]Task, bool)`

GetTasksOk returns a tuple with the Tasks field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTasks

`func (o *ServerJob) SetTasks(v []Task)`

SetTasks sets Tasks field to given value.

### HasTasks

`func (o *ServerJob) HasTasks() bool`

HasTasks returns a boolean if a field has been set.

### GetType

`func (o *ServerJob) GetType() JobType`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *ServerJob) GetTypeOk() (*JobType, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *ServerJob) SetType(v JobType)`

SetType sets Type field to given value.

### HasType

`func (o *ServerJob) HasType() bool`

HasType returns a boolean if a field has been set.

### GetUpdatedAt

`func (o *ServerJob) GetUpdatedAt() time.Time`

GetUpdatedAt returns the UpdatedAt field if non-nil, zero value otherwise.

### GetUpdatedAtOk

`func (o *ServerJob) GetUpdatedAtOk() (*time.Time, bool)`

GetUpdatedAtOk returns a tuple with the UpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedAt

`func (o *ServerJob) SetUpdatedAt(v time.Time)`

SetUpdatedAt sets UpdatedAt field to given value.

### HasUpdatedAt

`func (o *ServerJob) HasUpdatedAt() bool`

HasUpdatedAt returns a boolean if a field has been set.

### GetUuid

`func (o *ServerJob) GetUuid() string`

GetUuid returns the Uuid field if non-nil, zero value otherwise.

### GetUuidOk

`func (o *ServerJob) GetUuidOk() (*string, bool)`

GetUuidOk returns a tuple with the Uuid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUuid

`func (o *ServerJob) SetUuid(v string)`

SetUuid sets Uuid field to given value.

### HasUuid

`func (o *ServerJob) HasUuid() bool`

HasUuid returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


