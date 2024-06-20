# LoadBalancerAutoScalingGroup

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CustomerId** | Pointer to **string** | The customer ID who owns the auto scaling group | [optional] 
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

## Methods

### NewLoadBalancerAutoScalingGroup

`func NewLoadBalancerAutoScalingGroup() *LoadBalancerAutoScalingGroup`

NewLoadBalancerAutoScalingGroup instantiates a new LoadBalancerAutoScalingGroup object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewLoadBalancerAutoScalingGroupWithDefaults

`func NewLoadBalancerAutoScalingGroupWithDefaults() *LoadBalancerAutoScalingGroup`

NewLoadBalancerAutoScalingGroupWithDefaults instantiates a new LoadBalancerAutoScalingGroup object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCustomerId

`func (o *LoadBalancerAutoScalingGroup) GetCustomerId() string`

GetCustomerId returns the CustomerId field if non-nil, zero value otherwise.

### GetCustomerIdOk

`func (o *LoadBalancerAutoScalingGroup) GetCustomerIdOk() (*string, bool)`

GetCustomerIdOk returns a tuple with the CustomerId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomerId

`func (o *LoadBalancerAutoScalingGroup) SetCustomerId(v string)`

SetCustomerId sets CustomerId field to given value.

### HasCustomerId

`func (o *LoadBalancerAutoScalingGroup) HasCustomerId() bool`

HasCustomerId returns a boolean if a field has been set.

### GetId

`func (o *LoadBalancerAutoScalingGroup) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *LoadBalancerAutoScalingGroup) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *LoadBalancerAutoScalingGroup) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *LoadBalancerAutoScalingGroup) HasId() bool`

HasId returns a boolean if a field has been set.

### GetType

`func (o *LoadBalancerAutoScalingGroup) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *LoadBalancerAutoScalingGroup) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *LoadBalancerAutoScalingGroup) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *LoadBalancerAutoScalingGroup) HasType() bool`

HasType returns a boolean if a field has been set.

### GetState

`func (o *LoadBalancerAutoScalingGroup) GetState() string`

GetState returns the State field if non-nil, zero value otherwise.

### GetStateOk

`func (o *LoadBalancerAutoScalingGroup) GetStateOk() (*string, bool)`

GetStateOk returns a tuple with the State field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetState

`func (o *LoadBalancerAutoScalingGroup) SetState(v string)`

SetState sets State field to given value.

### HasState

`func (o *LoadBalancerAutoScalingGroup) HasState() bool`

HasState returns a boolean if a field has been set.

### GetDesiredAmount

`func (o *LoadBalancerAutoScalingGroup) GetDesiredAmount() int32`

GetDesiredAmount returns the DesiredAmount field if non-nil, zero value otherwise.

### GetDesiredAmountOk

`func (o *LoadBalancerAutoScalingGroup) GetDesiredAmountOk() (*int32, bool)`

GetDesiredAmountOk returns a tuple with the DesiredAmount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDesiredAmount

`func (o *LoadBalancerAutoScalingGroup) SetDesiredAmount(v int32)`

SetDesiredAmount sets DesiredAmount field to given value.

### HasDesiredAmount

`func (o *LoadBalancerAutoScalingGroup) HasDesiredAmount() bool`

HasDesiredAmount returns a boolean if a field has been set.

### SetDesiredAmountNil

`func (o *LoadBalancerAutoScalingGroup) SetDesiredAmountNil(b bool)`

 SetDesiredAmountNil sets the value for DesiredAmount to be an explicit nil

### UnsetDesiredAmount
`func (o *LoadBalancerAutoScalingGroup) UnsetDesiredAmount()`

UnsetDesiredAmount ensures that no value is present for DesiredAmount, not even an explicit nil
### GetMinimumAmount

`func (o *LoadBalancerAutoScalingGroup) GetMinimumAmount() int32`

GetMinimumAmount returns the MinimumAmount field if non-nil, zero value otherwise.

### GetMinimumAmountOk

`func (o *LoadBalancerAutoScalingGroup) GetMinimumAmountOk() (*int32, bool)`

GetMinimumAmountOk returns a tuple with the MinimumAmount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMinimumAmount

`func (o *LoadBalancerAutoScalingGroup) SetMinimumAmount(v int32)`

SetMinimumAmount sets MinimumAmount field to given value.

### HasMinimumAmount

`func (o *LoadBalancerAutoScalingGroup) HasMinimumAmount() bool`

HasMinimumAmount returns a boolean if a field has been set.

### SetMinimumAmountNil

`func (o *LoadBalancerAutoScalingGroup) SetMinimumAmountNil(b bool)`

 SetMinimumAmountNil sets the value for MinimumAmount to be an explicit nil

### UnsetMinimumAmount
`func (o *LoadBalancerAutoScalingGroup) UnsetMinimumAmount()`

UnsetMinimumAmount ensures that no value is present for MinimumAmount, not even an explicit nil
### GetMaximumAmount

`func (o *LoadBalancerAutoScalingGroup) GetMaximumAmount() int32`

GetMaximumAmount returns the MaximumAmount field if non-nil, zero value otherwise.

### GetMaximumAmountOk

`func (o *LoadBalancerAutoScalingGroup) GetMaximumAmountOk() (*int32, bool)`

GetMaximumAmountOk returns a tuple with the MaximumAmount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaximumAmount

`func (o *LoadBalancerAutoScalingGroup) SetMaximumAmount(v int32)`

SetMaximumAmount sets MaximumAmount field to given value.

### HasMaximumAmount

`func (o *LoadBalancerAutoScalingGroup) HasMaximumAmount() bool`

HasMaximumAmount returns a boolean if a field has been set.

### SetMaximumAmountNil

`func (o *LoadBalancerAutoScalingGroup) SetMaximumAmountNil(b bool)`

 SetMaximumAmountNil sets the value for MaximumAmount to be an explicit nil

### UnsetMaximumAmount
`func (o *LoadBalancerAutoScalingGroup) UnsetMaximumAmount()`

UnsetMaximumAmount ensures that no value is present for MaximumAmount, not even an explicit nil
### GetCpuThreshold

`func (o *LoadBalancerAutoScalingGroup) GetCpuThreshold() int32`

GetCpuThreshold returns the CpuThreshold field if non-nil, zero value otherwise.

### GetCpuThresholdOk

`func (o *LoadBalancerAutoScalingGroup) GetCpuThresholdOk() (*int32, bool)`

GetCpuThresholdOk returns a tuple with the CpuThreshold field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCpuThreshold

`func (o *LoadBalancerAutoScalingGroup) SetCpuThreshold(v int32)`

SetCpuThreshold sets CpuThreshold field to given value.

### HasCpuThreshold

`func (o *LoadBalancerAutoScalingGroup) HasCpuThreshold() bool`

HasCpuThreshold returns a boolean if a field has been set.

### SetCpuThresholdNil

`func (o *LoadBalancerAutoScalingGroup) SetCpuThresholdNil(b bool)`

 SetCpuThresholdNil sets the value for CpuThreshold to be an explicit nil

### UnsetCpuThreshold
`func (o *LoadBalancerAutoScalingGroup) UnsetCpuThreshold()`

UnsetCpuThreshold ensures that no value is present for CpuThreshold, not even an explicit nil
### GetWarmupTime

`func (o *LoadBalancerAutoScalingGroup) GetWarmupTime() int32`

GetWarmupTime returns the WarmupTime field if non-nil, zero value otherwise.

### GetWarmupTimeOk

`func (o *LoadBalancerAutoScalingGroup) GetWarmupTimeOk() (*int32, bool)`

GetWarmupTimeOk returns a tuple with the WarmupTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWarmupTime

`func (o *LoadBalancerAutoScalingGroup) SetWarmupTime(v int32)`

SetWarmupTime sets WarmupTime field to given value.

### HasWarmupTime

`func (o *LoadBalancerAutoScalingGroup) HasWarmupTime() bool`

HasWarmupTime returns a boolean if a field has been set.

### SetWarmupTimeNil

`func (o *LoadBalancerAutoScalingGroup) SetWarmupTimeNil(b bool)`

 SetWarmupTimeNil sets the value for WarmupTime to be an explicit nil

### UnsetWarmupTime
`func (o *LoadBalancerAutoScalingGroup) UnsetWarmupTime()`

UnsetWarmupTime ensures that no value is present for WarmupTime, not even an explicit nil
### GetRegion

`func (o *LoadBalancerAutoScalingGroup) GetRegion() string`

GetRegion returns the Region field if non-nil, zero value otherwise.

### GetRegionOk

`func (o *LoadBalancerAutoScalingGroup) GetRegionOk() (*string, bool)`

GetRegionOk returns a tuple with the Region field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRegion

`func (o *LoadBalancerAutoScalingGroup) SetRegion(v string)`

SetRegion sets Region field to given value.

### HasRegion

`func (o *LoadBalancerAutoScalingGroup) HasRegion() bool`

HasRegion returns a boolean if a field has been set.

### GetReference

`func (o *LoadBalancerAutoScalingGroup) GetReference() string`

GetReference returns the Reference field if non-nil, zero value otherwise.

### GetReferenceOk

`func (o *LoadBalancerAutoScalingGroup) GetReferenceOk() (*string, bool)`

GetReferenceOk returns a tuple with the Reference field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReference

`func (o *LoadBalancerAutoScalingGroup) SetReference(v string)`

SetReference sets Reference field to given value.

### HasReference

`func (o *LoadBalancerAutoScalingGroup) HasReference() bool`

HasReference returns a boolean if a field has been set.

### GetCreatedAt

`func (o *LoadBalancerAutoScalingGroup) GetCreatedAt() time.Time`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *LoadBalancerAutoScalingGroup) GetCreatedAtOk() (*time.Time, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *LoadBalancerAutoScalingGroup) SetCreatedAt(v time.Time)`

SetCreatedAt sets CreatedAt field to given value.

### HasCreatedAt

`func (o *LoadBalancerAutoScalingGroup) HasCreatedAt() bool`

HasCreatedAt returns a boolean if a field has been set.

### GetUpdatedAt

`func (o *LoadBalancerAutoScalingGroup) GetUpdatedAt() time.Time`

GetUpdatedAt returns the UpdatedAt field if non-nil, zero value otherwise.

### GetUpdatedAtOk

`func (o *LoadBalancerAutoScalingGroup) GetUpdatedAtOk() (*time.Time, bool)`

GetUpdatedAtOk returns a tuple with the UpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedAt

`func (o *LoadBalancerAutoScalingGroup) SetUpdatedAt(v time.Time)`

SetUpdatedAt sets UpdatedAt field to given value.

### HasUpdatedAt

`func (o *LoadBalancerAutoScalingGroup) HasUpdatedAt() bool`

HasUpdatedAt returns a boolean if a field has been set.

### GetStartsAt

`func (o *LoadBalancerAutoScalingGroup) GetStartsAt() time.Time`

GetStartsAt returns the StartsAt field if non-nil, zero value otherwise.

### GetStartsAtOk

`func (o *LoadBalancerAutoScalingGroup) GetStartsAtOk() (*time.Time, bool)`

GetStartsAtOk returns a tuple with the StartsAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStartsAt

`func (o *LoadBalancerAutoScalingGroup) SetStartsAt(v time.Time)`

SetStartsAt sets StartsAt field to given value.

### HasStartsAt

`func (o *LoadBalancerAutoScalingGroup) HasStartsAt() bool`

HasStartsAt returns a boolean if a field has been set.

### SetStartsAtNil

`func (o *LoadBalancerAutoScalingGroup) SetStartsAtNil(b bool)`

 SetStartsAtNil sets the value for StartsAt to be an explicit nil

### UnsetStartsAt
`func (o *LoadBalancerAutoScalingGroup) UnsetStartsAt()`

UnsetStartsAt ensures that no value is present for StartsAt, not even an explicit nil
### GetEndsAt

`func (o *LoadBalancerAutoScalingGroup) GetEndsAt() time.Time`

GetEndsAt returns the EndsAt field if non-nil, zero value otherwise.

### GetEndsAtOk

`func (o *LoadBalancerAutoScalingGroup) GetEndsAtOk() (*time.Time, bool)`

GetEndsAtOk returns a tuple with the EndsAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEndsAt

`func (o *LoadBalancerAutoScalingGroup) SetEndsAt(v time.Time)`

SetEndsAt sets EndsAt field to given value.

### HasEndsAt

`func (o *LoadBalancerAutoScalingGroup) HasEndsAt() bool`

HasEndsAt returns a boolean if a field has been set.

### SetEndsAtNil

`func (o *LoadBalancerAutoScalingGroup) SetEndsAtNil(b bool)`

 SetEndsAtNil sets the value for EndsAt to be an explicit nil

### UnsetEndsAt
`func (o *LoadBalancerAutoScalingGroup) UnsetEndsAt()`

UnsetEndsAt ensures that no value is present for EndsAt, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


