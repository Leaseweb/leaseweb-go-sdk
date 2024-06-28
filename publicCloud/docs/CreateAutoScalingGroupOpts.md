# CreateAutoScalingGroupOpts

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**DesiredAmount** | Pointer to **int32** | Required for \&quot;MANUAL\&quot; and \&quot;SCHEDULED\&quot; auto scaling group. Number of instances to be launched | [optional] 
**MinimumAmount** | Pointer to **int32** | Required for \&quot;CPU_BASED\&quot;. The minimum number of instances that should be running | [optional] 
**MaximumAmount** | Pointer to **int32** | Required for \&quot;CPU_BASED\&quot; auto scaling group. The maximum number of instances that can be running | [optional] 
**CpuThreshold** | Pointer to **int32** | Required for \&quot;CPU_BASED\&quot; auto scaling group. The target average CPU utilization for scaling | [optional] 
**WarmupTime** | Pointer to **int32** | Required for \&quot;CPU_BASED\&quot; auto scaling group. Warm-up time in seconds for new instances | [optional] 
**CooldownTime** | Pointer to **int32** | Required for \&quot;CPU_BASED\&quot; auto scaling group. Cool-down time in seconds for new instances | [optional] 
**InstanceId** | **string** | The instance on which instances will be based on. This instance needs to be either Running or Stopped | 
**Reference** | **string** | The identifying name set to the auto scaling group | 
**Type** | **string** |  | 
**StartsAt** | Pointer to **time.Time** | Required for \&quot;SCHEDULED\&quot; auto scaling group. Date and time (UTC) that the instances need to be launched | [optional] 
**EndsAt** | Pointer to **time.Time** | Required for \&quot;SCHEDULED\&quot; auto scaling group. Date and time (UTC) that the instances need to be terminated | [optional] 

## Methods

### NewCreateAutoScalingGroupOpts

`func NewCreateAutoScalingGroupOpts(instanceId string, reference string, type_ string, ) *CreateAutoScalingGroupOpts`

NewCreateAutoScalingGroupOpts instantiates a new CreateAutoScalingGroupOpts object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCreateAutoScalingGroupOptsWithDefaults

`func NewCreateAutoScalingGroupOptsWithDefaults() *CreateAutoScalingGroupOpts`

NewCreateAutoScalingGroupOptsWithDefaults instantiates a new CreateAutoScalingGroupOpts object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDesiredAmount

`func (o *CreateAutoScalingGroupOpts) GetDesiredAmount() int32`

GetDesiredAmount returns the DesiredAmount field if non-nil, zero value otherwise.

### GetDesiredAmountOk

`func (o *CreateAutoScalingGroupOpts) GetDesiredAmountOk() (*int32, bool)`

GetDesiredAmountOk returns a tuple with the DesiredAmount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDesiredAmount

`func (o *CreateAutoScalingGroupOpts) SetDesiredAmount(v int32)`

SetDesiredAmount sets DesiredAmount field to given value.

### HasDesiredAmount

`func (o *CreateAutoScalingGroupOpts) HasDesiredAmount() bool`

HasDesiredAmount returns a boolean if a field has been set.

### GetMinimumAmount

`func (o *CreateAutoScalingGroupOpts) GetMinimumAmount() int32`

GetMinimumAmount returns the MinimumAmount field if non-nil, zero value otherwise.

### GetMinimumAmountOk

`func (o *CreateAutoScalingGroupOpts) GetMinimumAmountOk() (*int32, bool)`

GetMinimumAmountOk returns a tuple with the MinimumAmount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMinimumAmount

`func (o *CreateAutoScalingGroupOpts) SetMinimumAmount(v int32)`

SetMinimumAmount sets MinimumAmount field to given value.

### HasMinimumAmount

`func (o *CreateAutoScalingGroupOpts) HasMinimumAmount() bool`

HasMinimumAmount returns a boolean if a field has been set.

### GetMaximumAmount

`func (o *CreateAutoScalingGroupOpts) GetMaximumAmount() int32`

GetMaximumAmount returns the MaximumAmount field if non-nil, zero value otherwise.

### GetMaximumAmountOk

`func (o *CreateAutoScalingGroupOpts) GetMaximumAmountOk() (*int32, bool)`

GetMaximumAmountOk returns a tuple with the MaximumAmount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaximumAmount

`func (o *CreateAutoScalingGroupOpts) SetMaximumAmount(v int32)`

SetMaximumAmount sets MaximumAmount field to given value.

### HasMaximumAmount

`func (o *CreateAutoScalingGroupOpts) HasMaximumAmount() bool`

HasMaximumAmount returns a boolean if a field has been set.

### GetCpuThreshold

`func (o *CreateAutoScalingGroupOpts) GetCpuThreshold() int32`

GetCpuThreshold returns the CpuThreshold field if non-nil, zero value otherwise.

### GetCpuThresholdOk

`func (o *CreateAutoScalingGroupOpts) GetCpuThresholdOk() (*int32, bool)`

GetCpuThresholdOk returns a tuple with the CpuThreshold field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCpuThreshold

`func (o *CreateAutoScalingGroupOpts) SetCpuThreshold(v int32)`

SetCpuThreshold sets CpuThreshold field to given value.

### HasCpuThreshold

`func (o *CreateAutoScalingGroupOpts) HasCpuThreshold() bool`

HasCpuThreshold returns a boolean if a field has been set.

### GetWarmupTime

`func (o *CreateAutoScalingGroupOpts) GetWarmupTime() int32`

GetWarmupTime returns the WarmupTime field if non-nil, zero value otherwise.

### GetWarmupTimeOk

`func (o *CreateAutoScalingGroupOpts) GetWarmupTimeOk() (*int32, bool)`

GetWarmupTimeOk returns a tuple with the WarmupTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWarmupTime

`func (o *CreateAutoScalingGroupOpts) SetWarmupTime(v int32)`

SetWarmupTime sets WarmupTime field to given value.

### HasWarmupTime

`func (o *CreateAutoScalingGroupOpts) HasWarmupTime() bool`

HasWarmupTime returns a boolean if a field has been set.

### GetCooldownTime

`func (o *CreateAutoScalingGroupOpts) GetCooldownTime() int32`

GetCooldownTime returns the CooldownTime field if non-nil, zero value otherwise.

### GetCooldownTimeOk

`func (o *CreateAutoScalingGroupOpts) GetCooldownTimeOk() (*int32, bool)`

GetCooldownTimeOk returns a tuple with the CooldownTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCooldownTime

`func (o *CreateAutoScalingGroupOpts) SetCooldownTime(v int32)`

SetCooldownTime sets CooldownTime field to given value.

### HasCooldownTime

`func (o *CreateAutoScalingGroupOpts) HasCooldownTime() bool`

HasCooldownTime returns a boolean if a field has been set.

### GetInstanceId

`func (o *CreateAutoScalingGroupOpts) GetInstanceId() string`

GetInstanceId returns the InstanceId field if non-nil, zero value otherwise.

### GetInstanceIdOk

`func (o *CreateAutoScalingGroupOpts) GetInstanceIdOk() (*string, bool)`

GetInstanceIdOk returns a tuple with the InstanceId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInstanceId

`func (o *CreateAutoScalingGroupOpts) SetInstanceId(v string)`

SetInstanceId sets InstanceId field to given value.


### GetReference

`func (o *CreateAutoScalingGroupOpts) GetReference() string`

GetReference returns the Reference field if non-nil, zero value otherwise.

### GetReferenceOk

`func (o *CreateAutoScalingGroupOpts) GetReferenceOk() (*string, bool)`

GetReferenceOk returns a tuple with the Reference field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReference

`func (o *CreateAutoScalingGroupOpts) SetReference(v string)`

SetReference sets Reference field to given value.


### GetType

`func (o *CreateAutoScalingGroupOpts) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *CreateAutoScalingGroupOpts) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *CreateAutoScalingGroupOpts) SetType(v string)`

SetType sets Type field to given value.


### GetStartsAt

`func (o *CreateAutoScalingGroupOpts) GetStartsAt() time.Time`

GetStartsAt returns the StartsAt field if non-nil, zero value otherwise.

### GetStartsAtOk

`func (o *CreateAutoScalingGroupOpts) GetStartsAtOk() (*time.Time, bool)`

GetStartsAtOk returns a tuple with the StartsAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStartsAt

`func (o *CreateAutoScalingGroupOpts) SetStartsAt(v time.Time)`

SetStartsAt sets StartsAt field to given value.

### HasStartsAt

`func (o *CreateAutoScalingGroupOpts) HasStartsAt() bool`

HasStartsAt returns a boolean if a field has been set.

### GetEndsAt

`func (o *CreateAutoScalingGroupOpts) GetEndsAt() time.Time`

GetEndsAt returns the EndsAt field if non-nil, zero value otherwise.

### GetEndsAtOk

`func (o *CreateAutoScalingGroupOpts) GetEndsAtOk() (*time.Time, bool)`

GetEndsAtOk returns a tuple with the EndsAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEndsAt

`func (o *CreateAutoScalingGroupOpts) SetEndsAt(v time.Time)`

SetEndsAt sets EndsAt field to given value.

### HasEndsAt

`func (o *CreateAutoScalingGroupOpts) HasEndsAt() bool`

HasEndsAt returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


