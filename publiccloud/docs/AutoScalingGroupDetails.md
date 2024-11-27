# AutoScalingGroupDetails

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
**TargetGroups** | [**[]TargetGroup**](TargetGroup.md) |  | 

## Methods

### NewAutoScalingGroupDetails

`func NewAutoScalingGroupDetails(id string, type_ AutoScalingGroupType, state AutoScalingGroupState, desiredAmount NullableInt32, region RegionName, reference string, createdAt time.Time, updatedAt time.Time, startsAt NullableTime, endsAt NullableTime, minimumAmount NullableInt32, maximumAmount NullableInt32, cpuThreshold NullableInt32, warmupTime NullableInt32, cooldownTime NullableInt32, targetGroups []TargetGroup, ) *AutoScalingGroupDetails`

NewAutoScalingGroupDetails instantiates a new AutoScalingGroupDetails object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAutoScalingGroupDetailsWithDefaults

`func NewAutoScalingGroupDetailsWithDefaults() *AutoScalingGroupDetails`

NewAutoScalingGroupDetailsWithDefaults instantiates a new AutoScalingGroupDetails object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *AutoScalingGroupDetails) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *AutoScalingGroupDetails) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *AutoScalingGroupDetails) SetId(v string)`

SetId sets Id field to given value.


### GetType

`func (o *AutoScalingGroupDetails) GetType() AutoScalingGroupType`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *AutoScalingGroupDetails) GetTypeOk() (*AutoScalingGroupType, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *AutoScalingGroupDetails) SetType(v AutoScalingGroupType)`

SetType sets Type field to given value.


### GetState

`func (o *AutoScalingGroupDetails) GetState() AutoScalingGroupState`

GetState returns the State field if non-nil, zero value otherwise.

### GetStateOk

`func (o *AutoScalingGroupDetails) GetStateOk() (*AutoScalingGroupState, bool)`

GetStateOk returns a tuple with the State field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetState

`func (o *AutoScalingGroupDetails) SetState(v AutoScalingGroupState)`

SetState sets State field to given value.


### GetDesiredAmount

`func (o *AutoScalingGroupDetails) GetDesiredAmount() int32`

GetDesiredAmount returns the DesiredAmount field if non-nil, zero value otherwise.

### GetDesiredAmountOk

`func (o *AutoScalingGroupDetails) GetDesiredAmountOk() (*int32, bool)`

GetDesiredAmountOk returns a tuple with the DesiredAmount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDesiredAmount

`func (o *AutoScalingGroupDetails) SetDesiredAmount(v int32)`

SetDesiredAmount sets DesiredAmount field to given value.


### SetDesiredAmountNil

`func (o *AutoScalingGroupDetails) SetDesiredAmountNil(b bool)`

 SetDesiredAmountNil sets the value for DesiredAmount to be an explicit nil

### UnsetDesiredAmount
`func (o *AutoScalingGroupDetails) UnsetDesiredAmount()`

UnsetDesiredAmount ensures that no value is present for DesiredAmount, not even an explicit nil
### GetRegion

`func (o *AutoScalingGroupDetails) GetRegion() RegionName`

GetRegion returns the Region field if non-nil, zero value otherwise.

### GetRegionOk

`func (o *AutoScalingGroupDetails) GetRegionOk() (*RegionName, bool)`

GetRegionOk returns a tuple with the Region field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRegion

`func (o *AutoScalingGroupDetails) SetRegion(v RegionName)`

SetRegion sets Region field to given value.


### GetReference

`func (o *AutoScalingGroupDetails) GetReference() string`

GetReference returns the Reference field if non-nil, zero value otherwise.

### GetReferenceOk

`func (o *AutoScalingGroupDetails) GetReferenceOk() (*string, bool)`

GetReferenceOk returns a tuple with the Reference field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReference

`func (o *AutoScalingGroupDetails) SetReference(v string)`

SetReference sets Reference field to given value.


### GetCreatedAt

`func (o *AutoScalingGroupDetails) GetCreatedAt() time.Time`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *AutoScalingGroupDetails) GetCreatedAtOk() (*time.Time, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *AutoScalingGroupDetails) SetCreatedAt(v time.Time)`

SetCreatedAt sets CreatedAt field to given value.


### GetUpdatedAt

`func (o *AutoScalingGroupDetails) GetUpdatedAt() time.Time`

GetUpdatedAt returns the UpdatedAt field if non-nil, zero value otherwise.

### GetUpdatedAtOk

`func (o *AutoScalingGroupDetails) GetUpdatedAtOk() (*time.Time, bool)`

GetUpdatedAtOk returns a tuple with the UpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedAt

`func (o *AutoScalingGroupDetails) SetUpdatedAt(v time.Time)`

SetUpdatedAt sets UpdatedAt field to given value.


### GetStartsAt

`func (o *AutoScalingGroupDetails) GetStartsAt() time.Time`

GetStartsAt returns the StartsAt field if non-nil, zero value otherwise.

### GetStartsAtOk

`func (o *AutoScalingGroupDetails) GetStartsAtOk() (*time.Time, bool)`

GetStartsAtOk returns a tuple with the StartsAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStartsAt

`func (o *AutoScalingGroupDetails) SetStartsAt(v time.Time)`

SetStartsAt sets StartsAt field to given value.


### SetStartsAtNil

`func (o *AutoScalingGroupDetails) SetStartsAtNil(b bool)`

 SetStartsAtNil sets the value for StartsAt to be an explicit nil

### UnsetStartsAt
`func (o *AutoScalingGroupDetails) UnsetStartsAt()`

UnsetStartsAt ensures that no value is present for StartsAt, not even an explicit nil
### GetEndsAt

`func (o *AutoScalingGroupDetails) GetEndsAt() time.Time`

GetEndsAt returns the EndsAt field if non-nil, zero value otherwise.

### GetEndsAtOk

`func (o *AutoScalingGroupDetails) GetEndsAtOk() (*time.Time, bool)`

GetEndsAtOk returns a tuple with the EndsAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEndsAt

`func (o *AutoScalingGroupDetails) SetEndsAt(v time.Time)`

SetEndsAt sets EndsAt field to given value.


### SetEndsAtNil

`func (o *AutoScalingGroupDetails) SetEndsAtNil(b bool)`

 SetEndsAtNil sets the value for EndsAt to be an explicit nil

### UnsetEndsAt
`func (o *AutoScalingGroupDetails) UnsetEndsAt()`

UnsetEndsAt ensures that no value is present for EndsAt, not even an explicit nil
### GetMinimumAmount

`func (o *AutoScalingGroupDetails) GetMinimumAmount() int32`

GetMinimumAmount returns the MinimumAmount field if non-nil, zero value otherwise.

### GetMinimumAmountOk

`func (o *AutoScalingGroupDetails) GetMinimumAmountOk() (*int32, bool)`

GetMinimumAmountOk returns a tuple with the MinimumAmount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMinimumAmount

`func (o *AutoScalingGroupDetails) SetMinimumAmount(v int32)`

SetMinimumAmount sets MinimumAmount field to given value.


### SetMinimumAmountNil

`func (o *AutoScalingGroupDetails) SetMinimumAmountNil(b bool)`

 SetMinimumAmountNil sets the value for MinimumAmount to be an explicit nil

### UnsetMinimumAmount
`func (o *AutoScalingGroupDetails) UnsetMinimumAmount()`

UnsetMinimumAmount ensures that no value is present for MinimumAmount, not even an explicit nil
### GetMaximumAmount

`func (o *AutoScalingGroupDetails) GetMaximumAmount() int32`

GetMaximumAmount returns the MaximumAmount field if non-nil, zero value otherwise.

### GetMaximumAmountOk

`func (o *AutoScalingGroupDetails) GetMaximumAmountOk() (*int32, bool)`

GetMaximumAmountOk returns a tuple with the MaximumAmount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaximumAmount

`func (o *AutoScalingGroupDetails) SetMaximumAmount(v int32)`

SetMaximumAmount sets MaximumAmount field to given value.


### SetMaximumAmountNil

`func (o *AutoScalingGroupDetails) SetMaximumAmountNil(b bool)`

 SetMaximumAmountNil sets the value for MaximumAmount to be an explicit nil

### UnsetMaximumAmount
`func (o *AutoScalingGroupDetails) UnsetMaximumAmount()`

UnsetMaximumAmount ensures that no value is present for MaximumAmount, not even an explicit nil
### GetCpuThreshold

`func (o *AutoScalingGroupDetails) GetCpuThreshold() int32`

GetCpuThreshold returns the CpuThreshold field if non-nil, zero value otherwise.

### GetCpuThresholdOk

`func (o *AutoScalingGroupDetails) GetCpuThresholdOk() (*int32, bool)`

GetCpuThresholdOk returns a tuple with the CpuThreshold field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCpuThreshold

`func (o *AutoScalingGroupDetails) SetCpuThreshold(v int32)`

SetCpuThreshold sets CpuThreshold field to given value.


### SetCpuThresholdNil

`func (o *AutoScalingGroupDetails) SetCpuThresholdNil(b bool)`

 SetCpuThresholdNil sets the value for CpuThreshold to be an explicit nil

### UnsetCpuThreshold
`func (o *AutoScalingGroupDetails) UnsetCpuThreshold()`

UnsetCpuThreshold ensures that no value is present for CpuThreshold, not even an explicit nil
### GetWarmupTime

`func (o *AutoScalingGroupDetails) GetWarmupTime() int32`

GetWarmupTime returns the WarmupTime field if non-nil, zero value otherwise.

### GetWarmupTimeOk

`func (o *AutoScalingGroupDetails) GetWarmupTimeOk() (*int32, bool)`

GetWarmupTimeOk returns a tuple with the WarmupTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWarmupTime

`func (o *AutoScalingGroupDetails) SetWarmupTime(v int32)`

SetWarmupTime sets WarmupTime field to given value.


### SetWarmupTimeNil

`func (o *AutoScalingGroupDetails) SetWarmupTimeNil(b bool)`

 SetWarmupTimeNil sets the value for WarmupTime to be an explicit nil

### UnsetWarmupTime
`func (o *AutoScalingGroupDetails) UnsetWarmupTime()`

UnsetWarmupTime ensures that no value is present for WarmupTime, not even an explicit nil
### GetCooldownTime

`func (o *AutoScalingGroupDetails) GetCooldownTime() int32`

GetCooldownTime returns the CooldownTime field if non-nil, zero value otherwise.

### GetCooldownTimeOk

`func (o *AutoScalingGroupDetails) GetCooldownTimeOk() (*int32, bool)`

GetCooldownTimeOk returns a tuple with the CooldownTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCooldownTime

`func (o *AutoScalingGroupDetails) SetCooldownTime(v int32)`

SetCooldownTime sets CooldownTime field to given value.


### SetCooldownTimeNil

`func (o *AutoScalingGroupDetails) SetCooldownTimeNil(b bool)`

 SetCooldownTimeNil sets the value for CooldownTime to be an explicit nil

### UnsetCooldownTime
`func (o *AutoScalingGroupDetails) UnsetCooldownTime()`

UnsetCooldownTime ensures that no value is present for CooldownTime, not even an explicit nil
### GetTargetGroups

`func (o *AutoScalingGroupDetails) GetTargetGroups() []TargetGroup`

GetTargetGroups returns the TargetGroups field if non-nil, zero value otherwise.

### GetTargetGroupsOk

`func (o *AutoScalingGroupDetails) GetTargetGroupsOk() (*[]TargetGroup, bool)`

GetTargetGroupsOk returns a tuple with the TargetGroups field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTargetGroups

`func (o *AutoScalingGroupDetails) SetTargetGroups(v []TargetGroup)`

SetTargetGroups sets TargetGroups field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


