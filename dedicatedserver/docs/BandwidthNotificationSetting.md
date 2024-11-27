# BandwidthNotificationSetting

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Actions** | Pointer to [**[]Actions**](Actions.md) | An array of notification setting actions | [optional] 
**Frequency** | **string** | Frequency | 
**Id** | **string** | Identifier of the notification setting | 
**LastCheckedAt** | Pointer to **NullableString** | Date timestamp when the system last checked the server for threshold limit | [optional] 
**Threshold** | **string** | Threshold Value | 
**ThresholdExceededAt** | Pointer to **NullableString** | Date timestamp when the threshold exceeded the limit | [optional] 
**Unit** | **string** | Unit | 

## Methods

### NewBandwidthNotificationSetting

`func NewBandwidthNotificationSetting(frequency string, id string, threshold string, unit string, ) *BandwidthNotificationSetting`

NewBandwidthNotificationSetting instantiates a new BandwidthNotificationSetting object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewBandwidthNotificationSettingWithDefaults

`func NewBandwidthNotificationSettingWithDefaults() *BandwidthNotificationSetting`

NewBandwidthNotificationSettingWithDefaults instantiates a new BandwidthNotificationSetting object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetActions

`func (o *BandwidthNotificationSetting) GetActions() []Actions`

GetActions returns the Actions field if non-nil, zero value otherwise.

### GetActionsOk

`func (o *BandwidthNotificationSetting) GetActionsOk() (*[]Actions, bool)`

GetActionsOk returns a tuple with the Actions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetActions

`func (o *BandwidthNotificationSetting) SetActions(v []Actions)`

SetActions sets Actions field to given value.

### HasActions

`func (o *BandwidthNotificationSetting) HasActions() bool`

HasActions returns a boolean if a field has been set.

### GetFrequency

`func (o *BandwidthNotificationSetting) GetFrequency() string`

GetFrequency returns the Frequency field if non-nil, zero value otherwise.

### GetFrequencyOk

`func (o *BandwidthNotificationSetting) GetFrequencyOk() (*string, bool)`

GetFrequencyOk returns a tuple with the Frequency field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFrequency

`func (o *BandwidthNotificationSetting) SetFrequency(v string)`

SetFrequency sets Frequency field to given value.


### GetId

`func (o *BandwidthNotificationSetting) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *BandwidthNotificationSetting) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *BandwidthNotificationSetting) SetId(v string)`

SetId sets Id field to given value.


### GetLastCheckedAt

`func (o *BandwidthNotificationSetting) GetLastCheckedAt() string`

GetLastCheckedAt returns the LastCheckedAt field if non-nil, zero value otherwise.

### GetLastCheckedAtOk

`func (o *BandwidthNotificationSetting) GetLastCheckedAtOk() (*string, bool)`

GetLastCheckedAtOk returns a tuple with the LastCheckedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastCheckedAt

`func (o *BandwidthNotificationSetting) SetLastCheckedAt(v string)`

SetLastCheckedAt sets LastCheckedAt field to given value.

### HasLastCheckedAt

`func (o *BandwidthNotificationSetting) HasLastCheckedAt() bool`

HasLastCheckedAt returns a boolean if a field has been set.

### SetLastCheckedAtNil

`func (o *BandwidthNotificationSetting) SetLastCheckedAtNil(b bool)`

 SetLastCheckedAtNil sets the value for LastCheckedAt to be an explicit nil

### UnsetLastCheckedAt
`func (o *BandwidthNotificationSetting) UnsetLastCheckedAt()`

UnsetLastCheckedAt ensures that no value is present for LastCheckedAt, not even an explicit nil
### GetThreshold

`func (o *BandwidthNotificationSetting) GetThreshold() string`

GetThreshold returns the Threshold field if non-nil, zero value otherwise.

### GetThresholdOk

`func (o *BandwidthNotificationSetting) GetThresholdOk() (*string, bool)`

GetThresholdOk returns a tuple with the Threshold field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetThreshold

`func (o *BandwidthNotificationSetting) SetThreshold(v string)`

SetThreshold sets Threshold field to given value.


### GetThresholdExceededAt

`func (o *BandwidthNotificationSetting) GetThresholdExceededAt() string`

GetThresholdExceededAt returns the ThresholdExceededAt field if non-nil, zero value otherwise.

### GetThresholdExceededAtOk

`func (o *BandwidthNotificationSetting) GetThresholdExceededAtOk() (*string, bool)`

GetThresholdExceededAtOk returns a tuple with the ThresholdExceededAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetThresholdExceededAt

`func (o *BandwidthNotificationSetting) SetThresholdExceededAt(v string)`

SetThresholdExceededAt sets ThresholdExceededAt field to given value.

### HasThresholdExceededAt

`func (o *BandwidthNotificationSetting) HasThresholdExceededAt() bool`

HasThresholdExceededAt returns a boolean if a field has been set.

### SetThresholdExceededAtNil

`func (o *BandwidthNotificationSetting) SetThresholdExceededAtNil(b bool)`

 SetThresholdExceededAtNil sets the value for ThresholdExceededAt to be an explicit nil

### UnsetThresholdExceededAt
`func (o *BandwidthNotificationSetting) UnsetThresholdExceededAt()`

UnsetThresholdExceededAt ensures that no value is present for ThresholdExceededAt, not even an explicit nil
### GetUnit

`func (o *BandwidthNotificationSetting) GetUnit() string`

GetUnit returns the Unit field if non-nil, zero value otherwise.

### GetUnitOk

`func (o *BandwidthNotificationSetting) GetUnitOk() (*string, bool)`

GetUnitOk returns a tuple with the Unit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUnit

`func (o *BandwidthNotificationSetting) SetUnit(v string)`

SetUnit sets Unit field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


