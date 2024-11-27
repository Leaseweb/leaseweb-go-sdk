# GetServerBandwidthNotificationSettingListResult

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Metadata** | Pointer to [**Metadata**](Metadata.md) |  | [optional] 
**BandwidthNotificationSettings** | Pointer to [**[]BandwidthNotificationSetting**](BandwidthNotificationSetting.md) | An array of Bandwidth Notification Settings | [optional] 

## Methods

### NewGetServerBandwidthNotificationSettingListResult

`func NewGetServerBandwidthNotificationSettingListResult() *GetServerBandwidthNotificationSettingListResult`

NewGetServerBandwidthNotificationSettingListResult instantiates a new GetServerBandwidthNotificationSettingListResult object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetServerBandwidthNotificationSettingListResultWithDefaults

`func NewGetServerBandwidthNotificationSettingListResultWithDefaults() *GetServerBandwidthNotificationSettingListResult`

NewGetServerBandwidthNotificationSettingListResultWithDefaults instantiates a new GetServerBandwidthNotificationSettingListResult object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetMetadata

`func (o *GetServerBandwidthNotificationSettingListResult) GetMetadata() Metadata`

GetMetadata returns the Metadata field if non-nil, zero value otherwise.

### GetMetadataOk

`func (o *GetServerBandwidthNotificationSettingListResult) GetMetadataOk() (*Metadata, bool)`

GetMetadataOk returns a tuple with the Metadata field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetadata

`func (o *GetServerBandwidthNotificationSettingListResult) SetMetadata(v Metadata)`

SetMetadata sets Metadata field to given value.

### HasMetadata

`func (o *GetServerBandwidthNotificationSettingListResult) HasMetadata() bool`

HasMetadata returns a boolean if a field has been set.

### GetBandwidthNotificationSettings

`func (o *GetServerBandwidthNotificationSettingListResult) GetBandwidthNotificationSettings() []BandwidthNotificationSetting`

GetBandwidthNotificationSettings returns the BandwidthNotificationSettings field if non-nil, zero value otherwise.

### GetBandwidthNotificationSettingsOk

`func (o *GetServerBandwidthNotificationSettingListResult) GetBandwidthNotificationSettingsOk() (*[]BandwidthNotificationSetting, bool)`

GetBandwidthNotificationSettingsOk returns a tuple with the BandwidthNotificationSettings field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBandwidthNotificationSettings

`func (o *GetServerBandwidthNotificationSettingListResult) SetBandwidthNotificationSettings(v []BandwidthNotificationSetting)`

SetBandwidthNotificationSettings sets BandwidthNotificationSettings field to given value.

### HasBandwidthNotificationSettings

`func (o *GetServerBandwidthNotificationSettingListResult) HasBandwidthNotificationSettings() bool`

HasBandwidthNotificationSettings returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


