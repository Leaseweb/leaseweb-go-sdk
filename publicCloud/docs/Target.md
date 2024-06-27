# Target

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **string** | The instance unique identifier | [optional] 
**Reference** | Pointer to **string** | The identifying name set to the instance | [optional] 
**OperatingSystem** | Pointer to [**OperatingSystem**](OperatingSystem.md) |  | [optional] 
**State** | Pointer to [**State**](State.md) |  | [optional] 
**HealthCheckStatus** | Pointer to [**HealthCheckStatus**](HealthCheckStatus.md) |  | [optional] 
**Ips** | Pointer to [**[]Ip**](Ip.md) |  | [optional] 

## Methods

### NewTarget

`func NewTarget() *Target`

NewTarget instantiates a new Target object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTargetWithDefaults

`func NewTargetWithDefaults() *Target`

NewTargetWithDefaults instantiates a new Target object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *Target) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *Target) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *Target) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *Target) HasId() bool`

HasId returns a boolean if a field has been set.

### GetReference

`func (o *Target) GetReference() string`

GetReference returns the Reference field if non-nil, zero value otherwise.

### GetReferenceOk

`func (o *Target) GetReferenceOk() (*string, bool)`

GetReferenceOk returns a tuple with the Reference field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReference

`func (o *Target) SetReference(v string)`

SetReference sets Reference field to given value.

### HasReference

`func (o *Target) HasReference() bool`

HasReference returns a boolean if a field has been set.

### GetOperatingSystem

`func (o *Target) GetOperatingSystem() OperatingSystem`

GetOperatingSystem returns the OperatingSystem field if non-nil, zero value otherwise.

### GetOperatingSystemOk

`func (o *Target) GetOperatingSystemOk() (*OperatingSystem, bool)`

GetOperatingSystemOk returns a tuple with the OperatingSystem field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOperatingSystem

`func (o *Target) SetOperatingSystem(v OperatingSystem)`

SetOperatingSystem sets OperatingSystem field to given value.

### HasOperatingSystem

`func (o *Target) HasOperatingSystem() bool`

HasOperatingSystem returns a boolean if a field has been set.

### GetState

`func (o *Target) GetState() State`

GetState returns the State field if non-nil, zero value otherwise.

### GetStateOk

`func (o *Target) GetStateOk() (*State, bool)`

GetStateOk returns a tuple with the State field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetState

`func (o *Target) SetState(v State)`

SetState sets State field to given value.

### HasState

`func (o *Target) HasState() bool`

HasState returns a boolean if a field has been set.

### GetHealthCheckStatus

`func (o *Target) GetHealthCheckStatus() HealthCheckStatus`

GetHealthCheckStatus returns the HealthCheckStatus field if non-nil, zero value otherwise.

### GetHealthCheckStatusOk

`func (o *Target) GetHealthCheckStatusOk() (*HealthCheckStatus, bool)`

GetHealthCheckStatusOk returns a tuple with the HealthCheckStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHealthCheckStatus

`func (o *Target) SetHealthCheckStatus(v HealthCheckStatus)`

SetHealthCheckStatus sets HealthCheckStatus field to given value.

### HasHealthCheckStatus

`func (o *Target) HasHealthCheckStatus() bool`

HasHealthCheckStatus returns a boolean if a field has been set.

### GetIps

`func (o *Target) GetIps() []Ip`

GetIps returns the Ips field if non-nil, zero value otherwise.

### GetIpsOk

`func (o *Target) GetIpsOk() (*[]Ip, bool)`

GetIpsOk returns a tuple with the Ips field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIps

`func (o *Target) SetIps(v []Ip)`

SetIps sets Ips field to given value.

### HasIps

`func (o *Target) HasIps() bool`

HasIps returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


