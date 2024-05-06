# Task

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Description** | Pointer to **string** | Description of the task | [optional] 
**ErrorMessage** | Pointer to **NullableString** | An optional error message | [optional] 
**Flow** | Pointer to **string** | The flow this task is part of | [optional] 
**OnError** | Pointer to **string** | The behaviour if this task fails | [optional] 
**Status** | Pointer to **string** | The status of the task | [optional] 
**StatusTimestamps** | Pointer to [**map[string]time.Time**](time.Time.md) | Timestamp for each state change | [optional] 
**Uuid** | Pointer to **string** | Unique ID for this task | [optional] 

## Methods

### NewTask

`func NewTask() *Task`

NewTask instantiates a new Task object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTaskWithDefaults

`func NewTaskWithDefaults() *Task`

NewTaskWithDefaults instantiates a new Task object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDescription

`func (o *Task) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *Task) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *Task) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *Task) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetErrorMessage

`func (o *Task) GetErrorMessage() string`

GetErrorMessage returns the ErrorMessage field if non-nil, zero value otherwise.

### GetErrorMessageOk

`func (o *Task) GetErrorMessageOk() (*string, bool)`

GetErrorMessageOk returns a tuple with the ErrorMessage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetErrorMessage

`func (o *Task) SetErrorMessage(v string)`

SetErrorMessage sets ErrorMessage field to given value.

### HasErrorMessage

`func (o *Task) HasErrorMessage() bool`

HasErrorMessage returns a boolean if a field has been set.

### SetErrorMessageNil

`func (o *Task) SetErrorMessageNil(b bool)`

 SetErrorMessageNil sets the value for ErrorMessage to be an explicit nil

### UnsetErrorMessage
`func (o *Task) UnsetErrorMessage()`

UnsetErrorMessage ensures that no value is present for ErrorMessage, not even an explicit nil
### GetFlow

`func (o *Task) GetFlow() string`

GetFlow returns the Flow field if non-nil, zero value otherwise.

### GetFlowOk

`func (o *Task) GetFlowOk() (*string, bool)`

GetFlowOk returns a tuple with the Flow field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFlow

`func (o *Task) SetFlow(v string)`

SetFlow sets Flow field to given value.

### HasFlow

`func (o *Task) HasFlow() bool`

HasFlow returns a boolean if a field has been set.

### GetOnError

`func (o *Task) GetOnError() string`

GetOnError returns the OnError field if non-nil, zero value otherwise.

### GetOnErrorOk

`func (o *Task) GetOnErrorOk() (*string, bool)`

GetOnErrorOk returns a tuple with the OnError field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOnError

`func (o *Task) SetOnError(v string)`

SetOnError sets OnError field to given value.

### HasOnError

`func (o *Task) HasOnError() bool`

HasOnError returns a boolean if a field has been set.

### GetStatus

`func (o *Task) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *Task) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *Task) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *Task) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetStatusTimestamps

`func (o *Task) GetStatusTimestamps() map[string]time.Time`

GetStatusTimestamps returns the StatusTimestamps field if non-nil, zero value otherwise.

### GetStatusTimestampsOk

`func (o *Task) GetStatusTimestampsOk() (*map[string]time.Time, bool)`

GetStatusTimestampsOk returns a tuple with the StatusTimestamps field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatusTimestamps

`func (o *Task) SetStatusTimestamps(v map[string]time.Time)`

SetStatusTimestamps sets StatusTimestamps field to given value.

### HasStatusTimestamps

`func (o *Task) HasStatusTimestamps() bool`

HasStatusTimestamps returns a boolean if a field has been set.

### GetUuid

`func (o *Task) GetUuid() string`

GetUuid returns the Uuid field if non-nil, zero value otherwise.

### GetUuidOk

`func (o *Task) GetUuidOk() (*string, bool)`

GetUuidOk returns a tuple with the Uuid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUuid

`func (o *Task) SetUuid(v string)`

SetUuid sets Uuid field to given value.

### HasUuid

`func (o *Task) HasUuid() bool`

HasUuid returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


