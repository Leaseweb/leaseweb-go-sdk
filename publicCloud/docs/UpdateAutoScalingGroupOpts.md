# UpdateAutoScalingGroupOpts

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**DesiredAmount** | Pointer to **int32** | When \&quot;SCHEDULED\&quot;, the number of instances that need to be running at the specified date and time.    When \&quot;MANUAL\&quot;, the number of instances that to be launched immediately. | [optional] 
**MinimumAmount** | Pointer to **int32** | The minimum number of instances that should be running | [optional] 
**MaximumAmount** | Pointer to **int32** | Only for \&quot;CPU_BASED\&quot; auto scaling group. The maximum number of instances that can be running | [optional] 
**CpuThreshold** | Pointer to **int32** | Only for \&quot;CPU_BASED\&quot; auto scaling group. The target average CPU utilization for scaling | [optional] 
**WarmupTime** | Pointer to **int32** | Only for \&quot;CPU_BASED\&quot; auto scaling group. Warm-up time in seconds for new instances | [optional] 
**CooldownTime** | Pointer to **int32** | Only for \&quot;CPU_BASED\&quot; auto scaling group. Cool-down time in seconds for new instances | [optional] 
**Reference** | Pointer to **string** | The identifying name set to the auto scaling group | [optional] 
**StartsAt** | Pointer to **time.Time** | Only for \&quot;SCHEDULED\&quot; auto scaling group. Date and time (UTC) that the instances need to be launched. Must be changed along with \&quot;endsAt\&quot; | [optional] 
**EndsAt** | Pointer to **time.Time** | Only for \&quot;SCHEDULED\&quot; auto scaling group. Date and time (UTC) that the instances need to be terminated. Must be changed along with \&quot;startsAt\&quot; | [optional] 

## Methods

### NewUpdateAutoScalingGroupOpts

`func NewUpdateAutoScalingGroupOpts() *UpdateAutoScalingGroupOpts`

NewUpdateAutoScalingGroupOpts instantiates a new UpdateAutoScalingGroupOpts object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUpdateAutoScalingGroupOptsWithDefaults

`func NewUpdateAutoScalingGroupOptsWithDefaults() *UpdateAutoScalingGroupOpts`

NewUpdateAutoScalingGroupOptsWithDefaults instantiates a new UpdateAutoScalingGroupOpts object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDesiredAmount

`func (o *UpdateAutoScalingGroupOpts) GetDesiredAmount() int32`

GetDesiredAmount returns the DesiredAmount field if non-nil, zero value otherwise.

### GetDesiredAmountOk

`func (o *UpdateAutoScalingGroupOpts) GetDesiredAmountOk() (*int32, bool)`

GetDesiredAmountOk returns a tuple with the DesiredAmount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDesiredAmount

`func (o *UpdateAutoScalingGroupOpts) SetDesiredAmount(v int32)`

SetDesiredAmount sets DesiredAmount field to given value.

### HasDesiredAmount

`func (o *UpdateAutoScalingGroupOpts) HasDesiredAmount() bool`

HasDesiredAmount returns a boolean if a field has been set.

### GetMinimumAmount

`func (o *UpdateAutoScalingGroupOpts) GetMinimumAmount() int32`

GetMinimumAmount returns the MinimumAmount field if non-nil, zero value otherwise.

### GetMinimumAmountOk

`func (o *UpdateAutoScalingGroupOpts) GetMinimumAmountOk() (*int32, bool)`

GetMinimumAmountOk returns a tuple with the MinimumAmount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMinimumAmount

`func (o *UpdateAutoScalingGroupOpts) SetMinimumAmount(v int32)`

SetMinimumAmount sets MinimumAmount field to given value.

### HasMinimumAmount

`func (o *UpdateAutoScalingGroupOpts) HasMinimumAmount() bool`

HasMinimumAmount returns a boolean if a field has been set.

### GetMaximumAmount

`func (o *UpdateAutoScalingGroupOpts) GetMaximumAmount() int32`

GetMaximumAmount returns the MaximumAmount field if non-nil, zero value otherwise.

### GetMaximumAmountOk

`func (o *UpdateAutoScalingGroupOpts) GetMaximumAmountOk() (*int32, bool)`

GetMaximumAmountOk returns a tuple with the MaximumAmount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaximumAmount

`func (o *UpdateAutoScalingGroupOpts) SetMaximumAmount(v int32)`

SetMaximumAmount sets MaximumAmount field to given value.

### HasMaximumAmount

`func (o *UpdateAutoScalingGroupOpts) HasMaximumAmount() bool`

HasMaximumAmount returns a boolean if a field has been set.

### GetCpuThreshold

`func (o *UpdateAutoScalingGroupOpts) GetCpuThreshold() int32`

GetCpuThreshold returns the CpuThreshold field if non-nil, zero value otherwise.

### GetCpuThresholdOk

`func (o *UpdateAutoScalingGroupOpts) GetCpuThresholdOk() (*int32, bool)`

GetCpuThresholdOk returns a tuple with the CpuThreshold field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCpuThreshold

`func (o *UpdateAutoScalingGroupOpts) SetCpuThreshold(v int32)`

SetCpuThreshold sets CpuThreshold field to given value.

### HasCpuThreshold

`func (o *UpdateAutoScalingGroupOpts) HasCpuThreshold() bool`

HasCpuThreshold returns a boolean if a field has been set.

### GetWarmupTime

`func (o *UpdateAutoScalingGroupOpts) GetWarmupTime() int32`

GetWarmupTime returns the WarmupTime field if non-nil, zero value otherwise.

### GetWarmupTimeOk

`func (o *UpdateAutoScalingGroupOpts) GetWarmupTimeOk() (*int32, bool)`

GetWarmupTimeOk returns a tuple with the WarmupTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWarmupTime

`func (o *UpdateAutoScalingGroupOpts) SetWarmupTime(v int32)`

SetWarmupTime sets WarmupTime field to given value.

### HasWarmupTime

`func (o *UpdateAutoScalingGroupOpts) HasWarmupTime() bool`

HasWarmupTime returns a boolean if a field has been set.

### GetCooldownTime

`func (o *UpdateAutoScalingGroupOpts) GetCooldownTime() int32`

GetCooldownTime returns the CooldownTime field if non-nil, zero value otherwise.

### GetCooldownTimeOk

`func (o *UpdateAutoScalingGroupOpts) GetCooldownTimeOk() (*int32, bool)`

GetCooldownTimeOk returns a tuple with the CooldownTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCooldownTime

`func (o *UpdateAutoScalingGroupOpts) SetCooldownTime(v int32)`

SetCooldownTime sets CooldownTime field to given value.

### HasCooldownTime

`func (o *UpdateAutoScalingGroupOpts) HasCooldownTime() bool`

HasCooldownTime returns a boolean if a field has been set.

### GetReference

`func (o *UpdateAutoScalingGroupOpts) GetReference() string`

GetReference returns the Reference field if non-nil, zero value otherwise.

### GetReferenceOk

`func (o *UpdateAutoScalingGroupOpts) GetReferenceOk() (*string, bool)`

GetReferenceOk returns a tuple with the Reference field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReference

`func (o *UpdateAutoScalingGroupOpts) SetReference(v string)`

SetReference sets Reference field to given value.

### HasReference

`func (o *UpdateAutoScalingGroupOpts) HasReference() bool`

HasReference returns a boolean if a field has been set.

### GetStartsAt

`func (o *UpdateAutoScalingGroupOpts) GetStartsAt() time.Time`

GetStartsAt returns the StartsAt field if non-nil, zero value otherwise.

### GetStartsAtOk

`func (o *UpdateAutoScalingGroupOpts) GetStartsAtOk() (*time.Time, bool)`

GetStartsAtOk returns a tuple with the StartsAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStartsAt

`func (o *UpdateAutoScalingGroupOpts) SetStartsAt(v time.Time)`

SetStartsAt sets StartsAt field to given value.

### HasStartsAt

`func (o *UpdateAutoScalingGroupOpts) HasStartsAt() bool`

HasStartsAt returns a boolean if a field has been set.

### GetEndsAt

`func (o *UpdateAutoScalingGroupOpts) GetEndsAt() time.Time`

GetEndsAt returns the EndsAt field if non-nil, zero value otherwise.

### GetEndsAtOk

`func (o *UpdateAutoScalingGroupOpts) GetEndsAtOk() (*time.Time, bool)`

GetEndsAtOk returns a tuple with the EndsAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEndsAt

`func (o *UpdateAutoScalingGroupOpts) SetEndsAt(v time.Time)`

SetEndsAt sets EndsAt field to given value.

### HasEndsAt

`func (o *UpdateAutoScalingGroupOpts) HasEndsAt() bool`

HasEndsAt returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


