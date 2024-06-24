# StickySession

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Enabled** | **bool** | If sticky session is enabled or not | 
**MaxLifeTime** | **int32** | Time that the Load Balancer routes the requests from one requester to the same target instance | 

## Methods

### NewStickySession

`func NewStickySession(enabled bool, maxLifeTime int32, ) *StickySession`

NewStickySession instantiates a new StickySession object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewStickySessionWithDefaults

`func NewStickySessionWithDefaults() *StickySession`

NewStickySessionWithDefaults instantiates a new StickySession object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetEnabled

`func (o *StickySession) GetEnabled() bool`

GetEnabled returns the Enabled field if non-nil, zero value otherwise.

### GetEnabledOk

`func (o *StickySession) GetEnabledOk() (*bool, bool)`

GetEnabledOk returns a tuple with the Enabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnabled

`func (o *StickySession) SetEnabled(v bool)`

SetEnabled sets Enabled field to given value.


### GetMaxLifeTime

`func (o *StickySession) GetMaxLifeTime() int32`

GetMaxLifeTime returns the MaxLifeTime field if non-nil, zero value otherwise.

### GetMaxLifeTimeOk

`func (o *StickySession) GetMaxLifeTimeOk() (*int32, bool)`

GetMaxLifeTimeOk returns a tuple with the MaxLifeTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxLifeTime

`func (o *StickySession) SetMaxLifeTime(v int32)`

SetMaxLifeTime sets MaxLifeTime field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


