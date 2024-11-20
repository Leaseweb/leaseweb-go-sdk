# UpdateTargetGroupOpts

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | Pointer to **string** | The name of the target group | [optional] 
**Port** | Pointer to **int32** | The port of the target group | [optional] 
**HealthCheck** | Pointer to [**HealthCheckOpts**](HealthCheckOpts.md) |  | [optional] 

## Methods

### NewUpdateTargetGroupOpts

`func NewUpdateTargetGroupOpts() *UpdateTargetGroupOpts`

NewUpdateTargetGroupOpts instantiates a new UpdateTargetGroupOpts object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUpdateTargetGroupOptsWithDefaults

`func NewUpdateTargetGroupOptsWithDefaults() *UpdateTargetGroupOpts`

NewUpdateTargetGroupOptsWithDefaults instantiates a new UpdateTargetGroupOpts object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *UpdateTargetGroupOpts) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *UpdateTargetGroupOpts) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *UpdateTargetGroupOpts) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *UpdateTargetGroupOpts) HasName() bool`

HasName returns a boolean if a field has been set.

### GetPort

`func (o *UpdateTargetGroupOpts) GetPort() int32`

GetPort returns the Port field if non-nil, zero value otherwise.

### GetPortOk

`func (o *UpdateTargetGroupOpts) GetPortOk() (*int32, bool)`

GetPortOk returns a tuple with the Port field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPort

`func (o *UpdateTargetGroupOpts) SetPort(v int32)`

SetPort sets Port field to given value.

### HasPort

`func (o *UpdateTargetGroupOpts) HasPort() bool`

HasPort returns a boolean if a field has been set.

### GetHealthCheck

`func (o *UpdateTargetGroupOpts) GetHealthCheck() HealthCheckOpts`

GetHealthCheck returns the HealthCheck field if non-nil, zero value otherwise.

### GetHealthCheckOk

`func (o *UpdateTargetGroupOpts) GetHealthCheckOk() (*HealthCheckOpts, bool)`

GetHealthCheckOk returns a tuple with the HealthCheck field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHealthCheck

`func (o *UpdateTargetGroupOpts) SetHealthCheck(v HealthCheckOpts)`

SetHealthCheck sets HealthCheck field to given value.

### HasHealthCheck

`func (o *UpdateTargetGroupOpts) HasHealthCheck() bool`

HasHealthCheck returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


