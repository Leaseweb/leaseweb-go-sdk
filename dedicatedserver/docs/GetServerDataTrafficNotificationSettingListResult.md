# GetServerDataTrafficNotificationSettingListResult

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Metadata** | Pointer to [**Metadata**](Metadata.md) |  | [optional] 
**DatatrafficNotificationSettings** | Pointer to [**[]DataTrafficNotificationSetting**](DataTrafficNotificationSetting.md) | An array of Data traffic Notification Settings | [optional] 

## Methods

### NewGetServerDataTrafficNotificationSettingListResult

`func NewGetServerDataTrafficNotificationSettingListResult() *GetServerDataTrafficNotificationSettingListResult`

NewGetServerDataTrafficNotificationSettingListResult instantiates a new GetServerDataTrafficNotificationSettingListResult object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetServerDataTrafficNotificationSettingListResultWithDefaults

`func NewGetServerDataTrafficNotificationSettingListResultWithDefaults() *GetServerDataTrafficNotificationSettingListResult`

NewGetServerDataTrafficNotificationSettingListResultWithDefaults instantiates a new GetServerDataTrafficNotificationSettingListResult object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetMetadata

`func (o *GetServerDataTrafficNotificationSettingListResult) GetMetadata() Metadata`

GetMetadata returns the Metadata field if non-nil, zero value otherwise.

### GetMetadataOk

`func (o *GetServerDataTrafficNotificationSettingListResult) GetMetadataOk() (*Metadata, bool)`

GetMetadataOk returns a tuple with the Metadata field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetadata

`func (o *GetServerDataTrafficNotificationSettingListResult) SetMetadata(v Metadata)`

SetMetadata sets Metadata field to given value.

### HasMetadata

`func (o *GetServerDataTrafficNotificationSettingListResult) HasMetadata() bool`

HasMetadata returns a boolean if a field has been set.

### GetDatatrafficNotificationSettings

`func (o *GetServerDataTrafficNotificationSettingListResult) GetDatatrafficNotificationSettings() []DataTrafficNotificationSetting`

GetDatatrafficNotificationSettings returns the DatatrafficNotificationSettings field if non-nil, zero value otherwise.

### GetDatatrafficNotificationSettingsOk

`func (o *GetServerDataTrafficNotificationSettingListResult) GetDatatrafficNotificationSettingsOk() (*[]DataTrafficNotificationSetting, bool)`

GetDatatrafficNotificationSettingsOk returns a tuple with the DatatrafficNotificationSettings field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDatatrafficNotificationSettings

`func (o *GetServerDataTrafficNotificationSettingListResult) SetDatatrafficNotificationSettings(v []DataTrafficNotificationSetting)`

SetDatatrafficNotificationSettings sets DatatrafficNotificationSettings field to given value.

### HasDatatrafficNotificationSettings

`func (o *GetServerDataTrafficNotificationSettingListResult) HasDatatrafficNotificationSettings() bool`

HasDatatrafficNotificationSettings returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


