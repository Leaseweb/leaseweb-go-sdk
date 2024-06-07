# AutoScalingGroup

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **string** | The Auto Scaling Group unique identifier | [optional] 
**Type** | Pointer to **string** | Auto Scaling Group type | [optional] 
**State** | Pointer to **string** | The Auto Scaling Group&#39;s current state. | [optional] 
**DesiredAmount** | Pointer to **NullableInt32** | Number of instances that should be running | [optional] 
**MinimumAmount** | Pointer to **NullableInt32** | The minimum number of instances that should be running | [optional] 
**MaximumAmount** | Pointer to **NullableInt32** | Only for \&quot;CPU_BASED\&quot; auto scaling group. The maximum number of instances that can be running | [optional] 
**CpuThreshold** | Pointer to **NullableInt32** | Only for \&quot;CPU_BASED\&quot; auto scaling group. The target average CPU utilization for scaling | [optional] 
**WarmupTime** | Pointer to **NullableInt32** | Only for \&quot;CPU_BASED\&quot; auto scaling group. Warm-up time in seconds for new instances | [optional] 
**Region** | Pointer to **string** | The region in which the Auto Scaling Group was launched | [optional] 
**Reference** | Pointer to **string** | The identifying name set to the auto scaling group | [optional] 
**CreatedAt** | Pointer to **time.Time** | Date and time when the Auto Scaling Group was created | [optional] 
**UpdatedAt** | Pointer to **time.Time** | Date and time when the Auto Scaling Group was last updated | [optional] 
**StartsAt** | Pointer to **NullableTime** | Only for \&quot;SCHEDULED\&quot; auto scaling group. Date and time (UTC) that the instances need to be launched | [optional] 
**EndsAt** | Pointer to **NullableTime** | Only for \&quot;SCHEDULED\&quot; auto scaling group. Date and time (UTC) that the instances need to be terminated | [optional] 
**LoadBalancer** | Pointer to [**NullableLoadBalancer1**](LoadBalancer1.md) |  | [optional] 

## Methods

### NewAutoScalingGroup

`func NewAutoScalingGroup() *AutoScalingGroup`

NewAutoScalingGroup instantiates a new AutoScalingGroup object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAutoScalingGroupWithDefaults

`func NewAutoScalingGroupWithDefaults() *AutoScalingGroup`

NewAutoScalingGroupWithDefaults instantiates a new AutoScalingGroup object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *AutoScalingGroup) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *AutoScalingGroup) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *AutoScalingGroup) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *AutoScalingGroup) HasId() bool`

HasId returns a boolean if a field has been set.

### GetType

`func (o *AutoScalingGroup) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *AutoScalingGroup) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *AutoScalingGroup) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *AutoScalingGroup) HasType() bool`

HasType returns a boolean if a field has been set.

### GetState

`func (o *AutoScalingGroup) GetState() string`

GetState returns the State field if non-nil, zero value otherwise.

### GetStateOk

`func (o *AutoScalingGroup) GetStateOk() (*string, bool)`

GetStateOk returns a tuple with the State field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetState

`func (o *AutoScalingGroup) SetState(v string)`

SetState sets State field to given value.

### HasState

`func (o *AutoScalingGroup) HasState() bool`

HasState returns a boolean if a field has been set.

### GetDesiredAmount

`func (o *AutoScalingGroup) GetDesiredAmount() int32`

GetDesiredAmount returns the DesiredAmount field if non-nil, zero value otherwise.

### GetDesiredAmountOk

`func (o *AutoScalingGroup) GetDesiredAmountOk() (*int32, bool)`

GetDesiredAmountOk returns a tuple with the DesiredAmount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDesiredAmount

`func (o *AutoScalingGroup) SetDesiredAmount(v int32)`

SetDesiredAmount sets DesiredAmount field to given value.

### HasDesiredAmount

`func (o *AutoScalingGroup) HasDesiredAmount() bool`

HasDesiredAmount returns a boolean if a field has been set.

### SetDesiredAmountNil

`func (o *AutoScalingGroup) SetDesiredAmountNil(b bool)`

 SetDesiredAmountNil sets the value for DesiredAmount to be an explicit nil

### UnsetDesiredAmount
`func (o *AutoScalingGroup) UnsetDesiredAmount()`

UnsetDesiredAmount ensures that no value is present for DesiredAmount, not even an explicit nil
### GetMinimumAmount

`func (o *AutoScalingGroup) GetMinimumAmount() int32`

GetMinimumAmount returns the MinimumAmount field if non-nil, zero value otherwise.

### GetMinimumAmountOk

`func (o *AutoScalingGroup) GetMinimumAmountOk() (*int32, bool)`

GetMinimumAmountOk returns a tuple with the MinimumAmount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMinimumAmount

`func (o *AutoScalingGroup) SetMinimumAmount(v int32)`

SetMinimumAmount sets MinimumAmount field to given value.

### HasMinimumAmount

`func (o *AutoScalingGroup) HasMinimumAmount() bool`

HasMinimumAmount returns a boolean if a field has been set.

### SetMinimumAmountNil

`func (o *AutoScalingGroup) SetMinimumAmountNil(b bool)`

 SetMinimumAmountNil sets the value for MinimumAmount to be an explicit nil

### UnsetMinimumAmount
`func (o *AutoScalingGroup) UnsetMinimumAmount()`

UnsetMinimumAmount ensures that no value is present for MinimumAmount, not even an explicit nil
### GetMaximumAmount

`func (o *AutoScalingGroup) GetMaximumAmount() int32`

GetMaximumAmount returns the MaximumAmount field if non-nil, zero value otherwise.

### GetMaximumAmountOk

`func (o *AutoScalingGroup) GetMaximumAmountOk() (*int32, bool)`

GetMaximumAmountOk returns a tuple with the MaximumAmount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaximumAmount

`func (o *AutoScalingGroup) SetMaximumAmount(v int32)`

SetMaximumAmount sets MaximumAmount field to given value.

### HasMaximumAmount

`func (o *AutoScalingGroup) HasMaximumAmount() bool`

HasMaximumAmount returns a boolean if a field has been set.

### SetMaximumAmountNil

`func (o *AutoScalingGroup) SetMaximumAmountNil(b bool)`

 SetMaximumAmountNil sets the value for MaximumAmount to be an explicit nil

### UnsetMaximumAmount
`func (o *AutoScalingGroup) UnsetMaximumAmount()`

UnsetMaximumAmount ensures that no value is present for MaximumAmount, not even an explicit nil
### GetCpuThreshold

`func (o *AutoScalingGroup) GetCpuThreshold() int32`

GetCpuThreshold returns the CpuThreshold field if non-nil, zero value otherwise.

### GetCpuThresholdOk

`func (o *AutoScalingGroup) GetCpuThresholdOk() (*int32, bool)`

GetCpuThresholdOk returns a tuple with the CpuThreshold field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCpuThreshold

`func (o *AutoScalingGroup) SetCpuThreshold(v int32)`

SetCpuThreshold sets CpuThreshold field to given value.

### HasCpuThreshold

`func (o *AutoScalingGroup) HasCpuThreshold() bool`

HasCpuThreshold returns a boolean if a field has been set.

### SetCpuThresholdNil

`func (o *AutoScalingGroup) SetCpuThresholdNil(b bool)`

 SetCpuThresholdNil sets the value for CpuThreshold to be an explicit nil

### UnsetCpuThreshold
`func (o *AutoScalingGroup) UnsetCpuThreshold()`

UnsetCpuThreshold ensures that no value is present for CpuThreshold, not even an explicit nil
### GetWarmupTime

`func (o *AutoScalingGroup) GetWarmupTime() int32`

GetWarmupTime returns the WarmupTime field if non-nil, zero value otherwise.

### GetWarmupTimeOk

`func (o *AutoScalingGroup) GetWarmupTimeOk() (*int32, bool)`

GetWarmupTimeOk returns a tuple with the WarmupTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWarmupTime

`func (o *AutoScalingGroup) SetWarmupTime(v int32)`

SetWarmupTime sets WarmupTime field to given value.

### HasWarmupTime

`func (o *AutoScalingGroup) HasWarmupTime() bool`

HasWarmupTime returns a boolean if a field has been set.

### SetWarmupTimeNil

`func (o *AutoScalingGroup) SetWarmupTimeNil(b bool)`

 SetWarmupTimeNil sets the value for WarmupTime to be an explicit nil

### UnsetWarmupTime
`func (o *AutoScalingGroup) UnsetWarmupTime()`

UnsetWarmupTime ensures that no value is present for WarmupTime, not even an explicit nil
### GetRegion

`func (o *AutoScalingGroup) GetRegion() string`

GetRegion returns the Region field if non-nil, zero value otherwise.

### GetRegionOk

`func (o *AutoScalingGroup) GetRegionOk() (*string, bool)`

GetRegionOk returns a tuple with the Region field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRegion

`func (o *AutoScalingGroup) SetRegion(v string)`

SetRegion sets Region field to given value.

### HasRegion

`func (o *AutoScalingGroup) HasRegion() bool`

HasRegion returns a boolean if a field has been set.

### GetReference

`func (o *AutoScalingGroup) GetReference() string`

GetReference returns the Reference field if non-nil, zero value otherwise.

### GetReferenceOk

`func (o *AutoScalingGroup) GetReferenceOk() (*string, bool)`

GetReferenceOk returns a tuple with the Reference field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReference

`func (o *AutoScalingGroup) SetReference(v string)`

SetReference sets Reference field to given value.

### HasReference

`func (o *AutoScalingGroup) HasReference() bool`

HasReference returns a boolean if a field has been set.

### GetCreatedAt

`func (o *AutoScalingGroup) GetCreatedAt() time.Time`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *AutoScalingGroup) GetCreatedAtOk() (*time.Time, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *AutoScalingGroup) SetCreatedAt(v time.Time)`

SetCreatedAt sets CreatedAt field to given value.

### HasCreatedAt

`func (o *AutoScalingGroup) HasCreatedAt() bool`

HasCreatedAt returns a boolean if a field has been set.

### GetUpdatedAt

`func (o *AutoScalingGroup) GetUpdatedAt() time.Time`

GetUpdatedAt returns the UpdatedAt field if non-nil, zero value otherwise.

### GetUpdatedAtOk

`func (o *AutoScalingGroup) GetUpdatedAtOk() (*time.Time, bool)`

GetUpdatedAtOk returns a tuple with the UpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedAt

`func (o *AutoScalingGroup) SetUpdatedAt(v time.Time)`

SetUpdatedAt sets UpdatedAt field to given value.

### HasUpdatedAt

`func (o *AutoScalingGroup) HasUpdatedAt() bool`

HasUpdatedAt returns a boolean if a field has been set.

### GetStartsAt

`func (o *AutoScalingGroup) GetStartsAt() time.Time`

GetStartsAt returns the StartsAt field if non-nil, zero value otherwise.

### GetStartsAtOk

`func (o *AutoScalingGroup) GetStartsAtOk() (*time.Time, bool)`

GetStartsAtOk returns a tuple with the StartsAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStartsAt

`func (o *AutoScalingGroup) SetStartsAt(v time.Time)`

SetStartsAt sets StartsAt field to given value.

### HasStartsAt

`func (o *AutoScalingGroup) HasStartsAt() bool`

HasStartsAt returns a boolean if a field has been set.

### SetStartsAtNil

`func (o *AutoScalingGroup) SetStartsAtNil(b bool)`

 SetStartsAtNil sets the value for StartsAt to be an explicit nil

### UnsetStartsAt
`func (o *AutoScalingGroup) UnsetStartsAt()`

UnsetStartsAt ensures that no value is present for StartsAt, not even an explicit nil
### GetEndsAt

`func (o *AutoScalingGroup) GetEndsAt() time.Time`

GetEndsAt returns the EndsAt field if non-nil, zero value otherwise.

### GetEndsAtOk

`func (o *AutoScalingGroup) GetEndsAtOk() (*time.Time, bool)`

GetEndsAtOk returns a tuple with the EndsAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEndsAt

`func (o *AutoScalingGroup) SetEndsAt(v time.Time)`

SetEndsAt sets EndsAt field to given value.

### HasEndsAt

`func (o *AutoScalingGroup) HasEndsAt() bool`

HasEndsAt returns a boolean if a field has been set.

### SetEndsAtNil

`func (o *AutoScalingGroup) SetEndsAtNil(b bool)`

 SetEndsAtNil sets the value for EndsAt to be an explicit nil

### UnsetEndsAt
`func (o *AutoScalingGroup) UnsetEndsAt()`

UnsetEndsAt ensures that no value is present for EndsAt, not even an explicit nil
### GetLoadBalancer

`func (o *AutoScalingGroup) GetLoadBalancer() LoadBalancer1`

GetLoadBalancer returns the LoadBalancer field if non-nil, zero value otherwise.

### GetLoadBalancerOk

`func (o *AutoScalingGroup) GetLoadBalancerOk() (*LoadBalancer1, bool)`

GetLoadBalancerOk returns a tuple with the LoadBalancer field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLoadBalancer

`func (o *AutoScalingGroup) SetLoadBalancer(v LoadBalancer1)`

SetLoadBalancer sets LoadBalancer field to given value.

### HasLoadBalancer

`func (o *AutoScalingGroup) HasLoadBalancer() bool`

HasLoadBalancer returns a boolean if a field has been set.

### SetLoadBalancerNil

`func (o *AutoScalingGroup) SetLoadBalancerNil(b bool)`

 SetLoadBalancerNil sets the value for LoadBalancer to be an explicit nil

### UnsetLoadBalancer
`func (o *AutoScalingGroup) UnsetLoadBalancer()`

UnsetLoadBalancer ensures that no value is present for LoadBalancer, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


