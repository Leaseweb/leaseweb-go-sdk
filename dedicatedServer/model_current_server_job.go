/*
Leaseweb API for dedicated servers

This documents the rest api dedicatedserver provides.

API version: v2
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package dedicatedServer

import (
	"encoding/json"
	"time"
)

// checks if the CurrentServerJob type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &CurrentServerJob{}

// CurrentServerJob struct for CurrentServerJob
type CurrentServerJob struct {
	// Id of the server (deprecated, use serverId instead)
	// Deprecated
	BareMetalId *string `json:"bareMetalId,omitempty"`
	// Creation timestamp
	CreatedAt *time.Time `json:"createdAt,omitempty"`
	// Job flow
	Flow *string `json:"flow,omitempty"`
	// Describes whether the job is running
	IsRunning *bool `json:"isRunning,omitempty"`
	// Node ID for this server
	Node *string `json:"node,omitempty"`
	Payload *ServerJobPayload `json:"payload,omitempty"`
	Progress *Progress `json:"progress,omitempty"`
	// Id of the server
	ServerId *string `json:"serverId,omitempty"`
	// Status of the job
	Status *string `json:"status,omitempty"`
	// List of tasks in the job
	Tasks []Task `json:"tasks,omitempty"`
	Type *JobType `json:"type,omitempty"`
	// Update timestamp
	UpdatedAt *time.Time `json:"updatedAt,omitempty"`
	// Unique identifier of the job
	Uuid *string `json:"uuid,omitempty"`
	Metadata *JobMetadata `json:"metadata,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _CurrentServerJob CurrentServerJob

// NewCurrentServerJob instantiates a new CurrentServerJob object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCurrentServerJob() *CurrentServerJob {
	this := CurrentServerJob{}
	return &this
}

// NewCurrentServerJobWithDefaults instantiates a new CurrentServerJob object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCurrentServerJobWithDefaults() *CurrentServerJob {
	this := CurrentServerJob{}
	return &this
}

// GetBareMetalId returns the BareMetalId field value if set, zero value otherwise.
// Deprecated
func (o *CurrentServerJob) GetBareMetalId() string {
	if o == nil || IsNil(o.BareMetalId) {
		var ret string
		return ret
	}
	return *o.BareMetalId
}

// GetBareMetalIdOk returns a tuple with the BareMetalId field value if set, nil otherwise
// and a boolean to check if the value has been set.
// Deprecated
func (o *CurrentServerJob) GetBareMetalIdOk() (*string, bool) {
	if o == nil || IsNil(o.BareMetalId) {
		return nil, false
	}
	return o.BareMetalId, true
}

// HasBareMetalId returns a boolean if a field has been set.
func (o *CurrentServerJob) HasBareMetalId() bool {
	if o != nil && !IsNil(o.BareMetalId) {
		return true
	}

	return false
}

// SetBareMetalId gets a reference to the given string and assigns it to the BareMetalId field.
// Deprecated
func (o *CurrentServerJob) SetBareMetalId(v string) {
	o.BareMetalId = &v
}

// GetCreatedAt returns the CreatedAt field value if set, zero value otherwise.
func (o *CurrentServerJob) GetCreatedAt() time.Time {
	if o == nil || IsNil(o.CreatedAt) {
		var ret time.Time
		return ret
	}
	return *o.CreatedAt
}

// GetCreatedAtOk returns a tuple with the CreatedAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CurrentServerJob) GetCreatedAtOk() (*time.Time, bool) {
	if o == nil || IsNil(o.CreatedAt) {
		return nil, false
	}
	return o.CreatedAt, true
}

// HasCreatedAt returns a boolean if a field has been set.
func (o *CurrentServerJob) HasCreatedAt() bool {
	if o != nil && !IsNil(o.CreatedAt) {
		return true
	}

	return false
}

// SetCreatedAt gets a reference to the given time.Time and assigns it to the CreatedAt field.
func (o *CurrentServerJob) SetCreatedAt(v time.Time) {
	o.CreatedAt = &v
}

// GetFlow returns the Flow field value if set, zero value otherwise.
func (o *CurrentServerJob) GetFlow() string {
	if o == nil || IsNil(o.Flow) {
		var ret string
		return ret
	}
	return *o.Flow
}

// GetFlowOk returns a tuple with the Flow field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CurrentServerJob) GetFlowOk() (*string, bool) {
	if o == nil || IsNil(o.Flow) {
		return nil, false
	}
	return o.Flow, true
}

// HasFlow returns a boolean if a field has been set.
func (o *CurrentServerJob) HasFlow() bool {
	if o != nil && !IsNil(o.Flow) {
		return true
	}

	return false
}

// SetFlow gets a reference to the given string and assigns it to the Flow field.
func (o *CurrentServerJob) SetFlow(v string) {
	o.Flow = &v
}

// GetIsRunning returns the IsRunning field value if set, zero value otherwise.
func (o *CurrentServerJob) GetIsRunning() bool {
	if o == nil || IsNil(o.IsRunning) {
		var ret bool
		return ret
	}
	return *o.IsRunning
}

// GetIsRunningOk returns a tuple with the IsRunning field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CurrentServerJob) GetIsRunningOk() (*bool, bool) {
	if o == nil || IsNil(o.IsRunning) {
		return nil, false
	}
	return o.IsRunning, true
}

// HasIsRunning returns a boolean if a field has been set.
func (o *CurrentServerJob) HasIsRunning() bool {
	if o != nil && !IsNil(o.IsRunning) {
		return true
	}

	return false
}

// SetIsRunning gets a reference to the given bool and assigns it to the IsRunning field.
func (o *CurrentServerJob) SetIsRunning(v bool) {
	o.IsRunning = &v
}

// GetNode returns the Node field value if set, zero value otherwise.
func (o *CurrentServerJob) GetNode() string {
	if o == nil || IsNil(o.Node) {
		var ret string
		return ret
	}
	return *o.Node
}

// GetNodeOk returns a tuple with the Node field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CurrentServerJob) GetNodeOk() (*string, bool) {
	if o == nil || IsNil(o.Node) {
		return nil, false
	}
	return o.Node, true
}

// HasNode returns a boolean if a field has been set.
func (o *CurrentServerJob) HasNode() bool {
	if o != nil && !IsNil(o.Node) {
		return true
	}

	return false
}

// SetNode gets a reference to the given string and assigns it to the Node field.
func (o *CurrentServerJob) SetNode(v string) {
	o.Node = &v
}

// GetPayload returns the Payload field value if set, zero value otherwise.
func (o *CurrentServerJob) GetPayload() ServerJobPayload {
	if o == nil || IsNil(o.Payload) {
		var ret ServerJobPayload
		return ret
	}
	return *o.Payload
}

// GetPayloadOk returns a tuple with the Payload field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CurrentServerJob) GetPayloadOk() (*ServerJobPayload, bool) {
	if o == nil || IsNil(o.Payload) {
		return nil, false
	}
	return o.Payload, true
}

// HasPayload returns a boolean if a field has been set.
func (o *CurrentServerJob) HasPayload() bool {
	if o != nil && !IsNil(o.Payload) {
		return true
	}

	return false
}

// SetPayload gets a reference to the given ServerJobPayload and assigns it to the Payload field.
func (o *CurrentServerJob) SetPayload(v ServerJobPayload) {
	o.Payload = &v
}

// GetProgress returns the Progress field value if set, zero value otherwise.
func (o *CurrentServerJob) GetProgress() Progress {
	if o == nil || IsNil(o.Progress) {
		var ret Progress
		return ret
	}
	return *o.Progress
}

// GetProgressOk returns a tuple with the Progress field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CurrentServerJob) GetProgressOk() (*Progress, bool) {
	if o == nil || IsNil(o.Progress) {
		return nil, false
	}
	return o.Progress, true
}

// HasProgress returns a boolean if a field has been set.
func (o *CurrentServerJob) HasProgress() bool {
	if o != nil && !IsNil(o.Progress) {
		return true
	}

	return false
}

// SetProgress gets a reference to the given Progress and assigns it to the Progress field.
func (o *CurrentServerJob) SetProgress(v Progress) {
	o.Progress = &v
}

// GetServerId returns the ServerId field value if set, zero value otherwise.
func (o *CurrentServerJob) GetServerId() string {
	if o == nil || IsNil(o.ServerId) {
		var ret string
		return ret
	}
	return *o.ServerId
}

// GetServerIdOk returns a tuple with the ServerId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CurrentServerJob) GetServerIdOk() (*string, bool) {
	if o == nil || IsNil(o.ServerId) {
		return nil, false
	}
	return o.ServerId, true
}

// HasServerId returns a boolean if a field has been set.
func (o *CurrentServerJob) HasServerId() bool {
	if o != nil && !IsNil(o.ServerId) {
		return true
	}

	return false
}

// SetServerId gets a reference to the given string and assigns it to the ServerId field.
func (o *CurrentServerJob) SetServerId(v string) {
	o.ServerId = &v
}

// GetStatus returns the Status field value if set, zero value otherwise.
func (o *CurrentServerJob) GetStatus() string {
	if o == nil || IsNil(o.Status) {
		var ret string
		return ret
	}
	return *o.Status
}

// GetStatusOk returns a tuple with the Status field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CurrentServerJob) GetStatusOk() (*string, bool) {
	if o == nil || IsNil(o.Status) {
		return nil, false
	}
	return o.Status, true
}

// HasStatus returns a boolean if a field has been set.
func (o *CurrentServerJob) HasStatus() bool {
	if o != nil && !IsNil(o.Status) {
		return true
	}

	return false
}

// SetStatus gets a reference to the given string and assigns it to the Status field.
func (o *CurrentServerJob) SetStatus(v string) {
	o.Status = &v
}

// GetTasks returns the Tasks field value if set, zero value otherwise.
func (o *CurrentServerJob) GetTasks() []Task {
	if o == nil || IsNil(o.Tasks) {
		var ret []Task
		return ret
	}
	return o.Tasks
}

// GetTasksOk returns a tuple with the Tasks field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CurrentServerJob) GetTasksOk() ([]Task, bool) {
	if o == nil || IsNil(o.Tasks) {
		return nil, false
	}
	return o.Tasks, true
}

// HasTasks returns a boolean if a field has been set.
func (o *CurrentServerJob) HasTasks() bool {
	if o != nil && !IsNil(o.Tasks) {
		return true
	}

	return false
}

// SetTasks gets a reference to the given []Task and assigns it to the Tasks field.
func (o *CurrentServerJob) SetTasks(v []Task) {
	o.Tasks = v
}

// GetType returns the Type field value if set, zero value otherwise.
func (o *CurrentServerJob) GetType() JobType {
	if o == nil || IsNil(o.Type) {
		var ret JobType
		return ret
	}
	return *o.Type
}

// GetTypeOk returns a tuple with the Type field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CurrentServerJob) GetTypeOk() (*JobType, bool) {
	if o == nil || IsNil(o.Type) {
		return nil, false
	}
	return o.Type, true
}

// HasType returns a boolean if a field has been set.
func (o *CurrentServerJob) HasType() bool {
	if o != nil && !IsNil(o.Type) {
		return true
	}

	return false
}

// SetType gets a reference to the given JobType and assigns it to the Type field.
func (o *CurrentServerJob) SetType(v JobType) {
	o.Type = &v
}

// GetUpdatedAt returns the UpdatedAt field value if set, zero value otherwise.
func (o *CurrentServerJob) GetUpdatedAt() time.Time {
	if o == nil || IsNil(o.UpdatedAt) {
		var ret time.Time
		return ret
	}
	return *o.UpdatedAt
}

// GetUpdatedAtOk returns a tuple with the UpdatedAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CurrentServerJob) GetUpdatedAtOk() (*time.Time, bool) {
	if o == nil || IsNil(o.UpdatedAt) {
		return nil, false
	}
	return o.UpdatedAt, true
}

// HasUpdatedAt returns a boolean if a field has been set.
func (o *CurrentServerJob) HasUpdatedAt() bool {
	if o != nil && !IsNil(o.UpdatedAt) {
		return true
	}

	return false
}

// SetUpdatedAt gets a reference to the given time.Time and assigns it to the UpdatedAt field.
func (o *CurrentServerJob) SetUpdatedAt(v time.Time) {
	o.UpdatedAt = &v
}

// GetUuid returns the Uuid field value if set, zero value otherwise.
func (o *CurrentServerJob) GetUuid() string {
	if o == nil || IsNil(o.Uuid) {
		var ret string
		return ret
	}
	return *o.Uuid
}

// GetUuidOk returns a tuple with the Uuid field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CurrentServerJob) GetUuidOk() (*string, bool) {
	if o == nil || IsNil(o.Uuid) {
		return nil, false
	}
	return o.Uuid, true
}

// HasUuid returns a boolean if a field has been set.
func (o *CurrentServerJob) HasUuid() bool {
	if o != nil && !IsNil(o.Uuid) {
		return true
	}

	return false
}

// SetUuid gets a reference to the given string and assigns it to the Uuid field.
func (o *CurrentServerJob) SetUuid(v string) {
	o.Uuid = &v
}

// GetMetadata returns the Metadata field value if set, zero value otherwise.
func (o *CurrentServerJob) GetMetadata() JobMetadata {
	if o == nil || IsNil(o.Metadata) {
		var ret JobMetadata
		return ret
	}
	return *o.Metadata
}

// GetMetadataOk returns a tuple with the Metadata field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CurrentServerJob) GetMetadataOk() (*JobMetadata, bool) {
	if o == nil || IsNil(o.Metadata) {
		return nil, false
	}
	return o.Metadata, true
}

// HasMetadata returns a boolean if a field has been set.
func (o *CurrentServerJob) HasMetadata() bool {
	if o != nil && !IsNil(o.Metadata) {
		return true
	}

	return false
}

// SetMetadata gets a reference to the given JobMetadata and assigns it to the Metadata field.
func (o *CurrentServerJob) SetMetadata(v JobMetadata) {
	o.Metadata = &v
}

func (o CurrentServerJob) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o CurrentServerJob) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.BareMetalId) {
		toSerialize["bareMetalId"] = o.BareMetalId
	}
	if !IsNil(o.CreatedAt) {
		toSerialize["createdAt"] = o.CreatedAt
	}
	if !IsNil(o.Flow) {
		toSerialize["flow"] = o.Flow
	}
	if !IsNil(o.IsRunning) {
		toSerialize["isRunning"] = o.IsRunning
	}
	if !IsNil(o.Node) {
		toSerialize["node"] = o.Node
	}
	if !IsNil(o.Payload) {
		toSerialize["payload"] = o.Payload
	}
	if !IsNil(o.Progress) {
		toSerialize["progress"] = o.Progress
	}
	if !IsNil(o.ServerId) {
		toSerialize["serverId"] = o.ServerId
	}
	if !IsNil(o.Status) {
		toSerialize["status"] = o.Status
	}
	if !IsNil(o.Tasks) {
		toSerialize["tasks"] = o.Tasks
	}
	if !IsNil(o.Type) {
		toSerialize["type"] = o.Type
	}
	if !IsNil(o.UpdatedAt) {
		toSerialize["updatedAt"] = o.UpdatedAt
	}
	if !IsNil(o.Uuid) {
		toSerialize["uuid"] = o.Uuid
	}
	if !IsNil(o.Metadata) {
		toSerialize["metadata"] = o.Metadata
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *CurrentServerJob) UnmarshalJSON(data []byte) (err error) {
	varCurrentServerJob := _CurrentServerJob{}

	err = json.Unmarshal(data, &varCurrentServerJob)

	if err != nil {
		return err
	}

	*o = CurrentServerJob(varCurrentServerJob)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "bareMetalId")
		delete(additionalProperties, "createdAt")
		delete(additionalProperties, "flow")
		delete(additionalProperties, "isRunning")
		delete(additionalProperties, "node")
		delete(additionalProperties, "payload")
		delete(additionalProperties, "progress")
		delete(additionalProperties, "serverId")
		delete(additionalProperties, "status")
		delete(additionalProperties, "tasks")
		delete(additionalProperties, "type")
		delete(additionalProperties, "updatedAt")
		delete(additionalProperties, "uuid")
		delete(additionalProperties, "metadata")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableCurrentServerJob struct {
	value *CurrentServerJob
	isSet bool
}

func (v NullableCurrentServerJob) Get() *CurrentServerJob {
	return v.value
}

func (v *NullableCurrentServerJob) Set(val *CurrentServerJob) {
	v.value = val
	v.isSet = true
}

func (v NullableCurrentServerJob) IsSet() bool {
	return v.isSet
}

func (v *NullableCurrentServerJob) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCurrentServerJob(val *CurrentServerJob) *NullableCurrentServerJob {
	return &NullableCurrentServerJob{value: val, isSet: true}
}

func (v NullableCurrentServerJob) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCurrentServerJob) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


