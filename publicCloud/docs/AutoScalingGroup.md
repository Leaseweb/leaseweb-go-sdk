# AutoScalingGroup

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** | The Auto Scaling Group unique identifier | 
**Type** | [**AutoScalingGroupType**](AutoScalingGroupType.md) |  | 
**State** | [**AutoScalingGroupState**](AutoScalingGroupState.md) |  | 
**DesiredAmount** | **NullableInt32** | Number of instances that should be running | 
**Region** | [**RegionName**](RegionName.md) |  | 
**Reference** | **string** | The identifying name set to the auto scaling group | 
**CreatedAt** | **time.Time** | Date and time when the Auto Scaling Group was created | 
**UpdatedAt** | **time.Time** | Date and time when the Auto Scaling Group was last updated | 
**StartsAt** | **NullableTime** | Only for \&quot;SCHEDULED\&quot; auto scaling group. Date and time (UTC) that the instances need to be launched | 
**EndsAt** | **NullableTime** | Only for \&quot;SCHEDULED\&quot; auto scaling group. Date and time (UTC) that the instances need to be terminated | 
**MinimumAmount** | **NullableInt32** | The minimum number of instances that should be running | 
**MaximumAmount** | **NullableInt32** | Only for \&quot;CPU_BASED\&quot; auto scaling group. The maximum number of instances that can be running | 
**CpuThreshold** | **NullableInt32** | Only for \&quot;CPU_BASED\&quot; auto scaling group. The target average CPU utilization for scaling | 
**WarmupTime** | **NullableInt32** | Only for \&quot;CPU_BASED\&quot; auto scaling group. Warm-up time in seconds for new instances | 
**CooldownTime** | **NullableInt32** | Only for \&quot;CPU_BASED\&quot; auto scaling group. Cool-down time in seconds for new instances | 

## Methods

### NewAutoScalingGroup

`func NewAutoScalingGroup(id string, type_ AutoScalingGroupType, state AutoScalingGroupState, desiredAmount NullableInt32, region RegionName, reference string, createdAt time.Time, updatedAt time.Time, startsAt NullableTime, endsAt NullableTime, minimumAmount NullableInt32, maximumAmount NullableInt32, cpuThreshold NullableInt32, warmupTime NullableInt32, cooldownTime NullableInt32, ) *AutoScalingGroup`

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


### GetType

`func (o *AutoScalingGroup) GetType() AutoScalingGroupType`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *AutoScalingGroup) GetTypeOk() (*AutoScalingGroupType, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *AutoScalingGroup) SetType(v AutoScalingGroupType)`

SetType sets Type field to given value.


### GetState

`func (o *AutoScalingGroup) GetState() AutoScalingGroupState`

GetState returns the State field if non-nil, zero value otherwise.

### GetStateOk

`func (o *AutoScalingGroup) GetStateOk() (*AutoScalingGroupState, bool)`

GetStateOk returns a tuple with the State field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetState

`func (o *AutoScalingGroup) SetState(v AutoScalingGroupState)`

SetState sets State field to given value.


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


### SetDesiredAmountNil

`func (o *AutoScalingGroup) SetDesiredAmountNil(b bool)`

 SetDesiredAmountNil sets the value for DesiredAmount to be an explicit nil

### UnsetDesiredAmount
`func (o *AutoScalingGroup) UnsetDesiredAmount()`

UnsetDesiredAmount ensures that no value is present for DesiredAmount, not even an explicit nil
### GetRegion

`func (o *AutoScalingGroup) GetRegion() RegionName`

GetRegion returns the Region field if non-nil, zero value otherwise.

### GetRegionOk

`func (o *AutoScalingGroup) GetRegionOk() (*RegionName, bool)`

GetRegionOk returns a tuple with the Region field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRegion

`func (o *AutoScalingGroup) SetRegion(v RegionName)`

SetRegion sets Region field to given value.


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


### SetEndsAtNil

`func (o *AutoScalingGroup) SetEndsAtNil(b bool)`

 SetEndsAtNil sets the value for EndsAt to be an explicit nil

### UnsetEndsAt
`func (o *AutoScalingGroup) UnsetEndsAt()`

UnsetEndsAt ensures that no value is present for EndsAt, not even an explicit nil
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


### SetWarmupTimeNil

`func (o *AutoScalingGroup) SetWarmupTimeNil(b bool)`

 SetWarmupTimeNil sets the value for WarmupTime to be an explicit nil

### UnsetWarmupTime
`func (o *AutoScalingGroup) UnsetWarmupTime()`

UnsetWarmupTime ensures that no value is present for WarmupTime, not even an explicit nil
### GetCooldownTime

`func (o *AutoScalingGroup) GetCooldownTime() int32`

GetCooldownTime returns the CooldownTime field if non-nil, zero value otherwise.

### GetCooldownTimeOk

`func (o *AutoScalingGroup) GetCooldownTimeOk() (*int32, bool)`

GetCooldownTimeOk returns a tuple with the CooldownTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCooldownTime

`func (o *AutoScalingGroup) SetCooldownTime(v int32)`

SetCooldownTime sets CooldownTime field to given value.


### SetCooldownTimeNil

`func (o *AutoScalingGroup) SetCooldownTimeNil(b bool)`

 SetCooldownTimeNil sets the value for CooldownTime to be an explicit nil

### UnsetCooldownTime
`func (o *AutoScalingGroup) UnsetCooldownTime()`

UnsetCooldownTime ensures that no value is present for CooldownTime, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


