# NotificationSetting

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Actions** | Pointer to [**[]Actions**](Actions.md) | An array of notification setting actions | [optional] 
**Frequency** | **string** | Frequency | 
**Id** | **string** | Identifier of the notification setting | 
**LastCheckedAt** | Pointer to **NullableString** | Date timestamp when the system last checked the server for threshold limit | [optional] 
**Threshold** | **string** | Threshold Value | 
**ThresholdExceededAt** | Pointer to **NullableString** | Date timestamp when the threshold exceeded the limit | [optional] 

## Methods

### NewNotificationSetting

`func NewNotificationSetting(frequency string, id string, threshold string, ) *NotificationSetting`

NewNotificationSetting instantiates a new NotificationSetting object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewNotificationSettingWithDefaults

`func NewNotificationSettingWithDefaults() *NotificationSetting`

NewNotificationSettingWithDefaults instantiates a new NotificationSetting object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetActions

`func (o *NotificationSetting) GetActions() []Actions`

GetActions returns the Actions field if non-nil, zero value otherwise.

### GetActionsOk

`func (o *NotificationSetting) GetActionsOk() (*[]Actions, bool)`

GetActionsOk returns a tuple with the Actions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetActions

`func (o *NotificationSetting) SetActions(v []Actions)`

SetActions sets Actions field to given value.

### HasActions

`func (o *NotificationSetting) HasActions() bool`

HasActions returns a boolean if a field has been set.

### GetFrequency

`func (o *NotificationSetting) GetFrequency() string`

GetFrequency returns the Frequency field if non-nil, zero value otherwise.

### GetFrequencyOk

`func (o *NotificationSetting) GetFrequencyOk() (*string, bool)`

GetFrequencyOk returns a tuple with the Frequency field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFrequency

`func (o *NotificationSetting) SetFrequency(v string)`

SetFrequency sets Frequency field to given value.


### GetId

`func (o *NotificationSetting) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *NotificationSetting) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *NotificationSetting) SetId(v string)`

SetId sets Id field to given value.


### GetLastCheckedAt

`func (o *NotificationSetting) GetLastCheckedAt() string`

GetLastCheckedAt returns the LastCheckedAt field if non-nil, zero value otherwise.

### GetLastCheckedAtOk

`func (o *NotificationSetting) GetLastCheckedAtOk() (*string, bool)`

GetLastCheckedAtOk returns a tuple with the LastCheckedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastCheckedAt

`func (o *NotificationSetting) SetLastCheckedAt(v string)`

SetLastCheckedAt sets LastCheckedAt field to given value.

### HasLastCheckedAt

`func (o *NotificationSetting) HasLastCheckedAt() bool`

HasLastCheckedAt returns a boolean if a field has been set.

### SetLastCheckedAtNil

`func (o *NotificationSetting) SetLastCheckedAtNil(b bool)`

 SetLastCheckedAtNil sets the value for LastCheckedAt to be an explicit nil

### UnsetLastCheckedAt
`func (o *NotificationSetting) UnsetLastCheckedAt()`

UnsetLastCheckedAt ensures that no value is present for LastCheckedAt, not even an explicit nil
### GetThreshold

`func (o *NotificationSetting) GetThreshold() string`

GetThreshold returns the Threshold field if non-nil, zero value otherwise.

### GetThresholdOk

`func (o *NotificationSetting) GetThresholdOk() (*string, bool)`

GetThresholdOk returns a tuple with the Threshold field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetThreshold

`func (o *NotificationSetting) SetThreshold(v string)`

SetThreshold sets Threshold field to given value.


### GetThresholdExceededAt

`func (o *NotificationSetting) GetThresholdExceededAt() string`

GetThresholdExceededAt returns the ThresholdExceededAt field if non-nil, zero value otherwise.

### GetThresholdExceededAtOk

`func (o *NotificationSetting) GetThresholdExceededAtOk() (*string, bool)`

GetThresholdExceededAtOk returns a tuple with the ThresholdExceededAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetThresholdExceededAt

`func (o *NotificationSetting) SetThresholdExceededAt(v string)`

SetThresholdExceededAt sets ThresholdExceededAt field to given value.

### HasThresholdExceededAt

`func (o *NotificationSetting) HasThresholdExceededAt() bool`

HasThresholdExceededAt returns a boolean if a field has been set.

### SetThresholdExceededAtNil

`func (o *NotificationSetting) SetThresholdExceededAtNil(b bool)`

 SetThresholdExceededAtNil sets the value for ThresholdExceededAt to be an explicit nil

### UnsetThresholdExceededAt
`func (o *NotificationSetting) UnsetThresholdExceededAt()`

UnsetThresholdExceededAt ensures that no value is present for ThresholdExceededAt, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


