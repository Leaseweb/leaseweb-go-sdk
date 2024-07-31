# ImageDetails

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** | imageId can be either an Operating System or a UUID in case of a Custom Image | 
**Name** | **string** |  | 
**Version** | **string** |  | 
**Family** | **string** |  | 
**Flavour** | **string** |  | 
**Architecture** | Pointer to **NullableString** |  | [optional] 
**State** | Pointer to **NullableString** |  | [optional] 
**StateReason** | Pointer to **NullableString** | The reason in case of failure | [optional] 
**Region** | Pointer to **NullableString** | The region where the image was uploaded | [optional] 
**CreatedAt** | Pointer to **NullableTime** | Date when the image was created | [optional] 
**UpdatedAt** | Pointer to **NullableTime** | Date when the image was updated | [optional] 
**Custom** | Pointer to **NullableBool** | Standard or Custom image | [optional] 
**MarketApps** | **[]string** |  | 
**StorageTypes** | **[]string** | The supported storage types for the instance type | 

## Methods

### NewImageDetails

`func NewImageDetails(id string, name string, version string, family string, flavour string, marketApps []string, storageTypes []string, ) *ImageDetails`

NewImageDetails instantiates a new ImageDetails object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewImageDetailsWithDefaults

`func NewImageDetailsWithDefaults() *ImageDetails`

NewImageDetailsWithDefaults instantiates a new ImageDetails object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *ImageDetails) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *ImageDetails) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *ImageDetails) SetId(v string)`

SetId sets Id field to given value.


### GetName

`func (o *ImageDetails) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *ImageDetails) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *ImageDetails) SetName(v string)`

SetName sets Name field to given value.


### GetVersion

`func (o *ImageDetails) GetVersion() string`

GetVersion returns the Version field if non-nil, zero value otherwise.

### GetVersionOk

`func (o *ImageDetails) GetVersionOk() (*string, bool)`

GetVersionOk returns a tuple with the Version field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVersion

`func (o *ImageDetails) SetVersion(v string)`

SetVersion sets Version field to given value.


### GetFamily

`func (o *ImageDetails) GetFamily() string`

GetFamily returns the Family field if non-nil, zero value otherwise.

### GetFamilyOk

`func (o *ImageDetails) GetFamilyOk() (*string, bool)`

GetFamilyOk returns a tuple with the Family field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFamily

`func (o *ImageDetails) SetFamily(v string)`

SetFamily sets Family field to given value.


### GetFlavour

`func (o *ImageDetails) GetFlavour() string`

GetFlavour returns the Flavour field if non-nil, zero value otherwise.

### GetFlavourOk

`func (o *ImageDetails) GetFlavourOk() (*string, bool)`

GetFlavourOk returns a tuple with the Flavour field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFlavour

`func (o *ImageDetails) SetFlavour(v string)`

SetFlavour sets Flavour field to given value.


### GetArchitecture

`func (o *ImageDetails) GetArchitecture() string`

GetArchitecture returns the Architecture field if non-nil, zero value otherwise.

### GetArchitectureOk

`func (o *ImageDetails) GetArchitectureOk() (*string, bool)`

GetArchitectureOk returns a tuple with the Architecture field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetArchitecture

`func (o *ImageDetails) SetArchitecture(v string)`

SetArchitecture sets Architecture field to given value.

### HasArchitecture

`func (o *ImageDetails) HasArchitecture() bool`

HasArchitecture returns a boolean if a field has been set.

### SetArchitectureNil

`func (o *ImageDetails) SetArchitectureNil(b bool)`

 SetArchitectureNil sets the value for Architecture to be an explicit nil

### UnsetArchitecture
`func (o *ImageDetails) UnsetArchitecture()`

UnsetArchitecture ensures that no value is present for Architecture, not even an explicit nil
### GetState

`func (o *ImageDetails) GetState() string`

GetState returns the State field if non-nil, zero value otherwise.

### GetStateOk

`func (o *ImageDetails) GetStateOk() (*string, bool)`

GetStateOk returns a tuple with the State field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetState

`func (o *ImageDetails) SetState(v string)`

SetState sets State field to given value.

### HasState

`func (o *ImageDetails) HasState() bool`

HasState returns a boolean if a field has been set.

### SetStateNil

`func (o *ImageDetails) SetStateNil(b bool)`

 SetStateNil sets the value for State to be an explicit nil

### UnsetState
`func (o *ImageDetails) UnsetState()`

UnsetState ensures that no value is present for State, not even an explicit nil
### GetStateReason

`func (o *ImageDetails) GetStateReason() string`

GetStateReason returns the StateReason field if non-nil, zero value otherwise.

### GetStateReasonOk

`func (o *ImageDetails) GetStateReasonOk() (*string, bool)`

GetStateReasonOk returns a tuple with the StateReason field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStateReason

`func (o *ImageDetails) SetStateReason(v string)`

SetStateReason sets StateReason field to given value.

### HasStateReason

`func (o *ImageDetails) HasStateReason() bool`

HasStateReason returns a boolean if a field has been set.

### SetStateReasonNil

`func (o *ImageDetails) SetStateReasonNil(b bool)`

 SetStateReasonNil sets the value for StateReason to be an explicit nil

### UnsetStateReason
`func (o *ImageDetails) UnsetStateReason()`

UnsetStateReason ensures that no value is present for StateReason, not even an explicit nil
### GetRegion

`func (o *ImageDetails) GetRegion() string`

GetRegion returns the Region field if non-nil, zero value otherwise.

### GetRegionOk

`func (o *ImageDetails) GetRegionOk() (*string, bool)`

GetRegionOk returns a tuple with the Region field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRegion

`func (o *ImageDetails) SetRegion(v string)`

SetRegion sets Region field to given value.

### HasRegion

`func (o *ImageDetails) HasRegion() bool`

HasRegion returns a boolean if a field has been set.

### SetRegionNil

`func (o *ImageDetails) SetRegionNil(b bool)`

 SetRegionNil sets the value for Region to be an explicit nil

### UnsetRegion
`func (o *ImageDetails) UnsetRegion()`

UnsetRegion ensures that no value is present for Region, not even an explicit nil
### GetCreatedAt

`func (o *ImageDetails) GetCreatedAt() time.Time`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *ImageDetails) GetCreatedAtOk() (*time.Time, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *ImageDetails) SetCreatedAt(v time.Time)`

SetCreatedAt sets CreatedAt field to given value.

### HasCreatedAt

`func (o *ImageDetails) HasCreatedAt() bool`

HasCreatedAt returns a boolean if a field has been set.

### SetCreatedAtNil

`func (o *ImageDetails) SetCreatedAtNil(b bool)`

 SetCreatedAtNil sets the value for CreatedAt to be an explicit nil

### UnsetCreatedAt
`func (o *ImageDetails) UnsetCreatedAt()`

UnsetCreatedAt ensures that no value is present for CreatedAt, not even an explicit nil
### GetUpdatedAt

`func (o *ImageDetails) GetUpdatedAt() time.Time`

GetUpdatedAt returns the UpdatedAt field if non-nil, zero value otherwise.

### GetUpdatedAtOk

`func (o *ImageDetails) GetUpdatedAtOk() (*time.Time, bool)`

GetUpdatedAtOk returns a tuple with the UpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedAt

`func (o *ImageDetails) SetUpdatedAt(v time.Time)`

SetUpdatedAt sets UpdatedAt field to given value.

### HasUpdatedAt

`func (o *ImageDetails) HasUpdatedAt() bool`

HasUpdatedAt returns a boolean if a field has been set.

### SetUpdatedAtNil

`func (o *ImageDetails) SetUpdatedAtNil(b bool)`

 SetUpdatedAtNil sets the value for UpdatedAt to be an explicit nil

### UnsetUpdatedAt
`func (o *ImageDetails) UnsetUpdatedAt()`

UnsetUpdatedAt ensures that no value is present for UpdatedAt, not even an explicit nil
### GetCustom

`func (o *ImageDetails) GetCustom() bool`

GetCustom returns the Custom field if non-nil, zero value otherwise.

### GetCustomOk

`func (o *ImageDetails) GetCustomOk() (*bool, bool)`

GetCustomOk returns a tuple with the Custom field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustom

`func (o *ImageDetails) SetCustom(v bool)`

SetCustom sets Custom field to given value.

### HasCustom

`func (o *ImageDetails) HasCustom() bool`

HasCustom returns a boolean if a field has been set.

### SetCustomNil

`func (o *ImageDetails) SetCustomNil(b bool)`

 SetCustomNil sets the value for Custom to be an explicit nil

### UnsetCustom
`func (o *ImageDetails) UnsetCustom()`

UnsetCustom ensures that no value is present for Custom, not even an explicit nil
### GetMarketApps

`func (o *ImageDetails) GetMarketApps() []string`

GetMarketApps returns the MarketApps field if non-nil, zero value otherwise.

### GetMarketAppsOk

`func (o *ImageDetails) GetMarketAppsOk() (*[]string, bool)`

GetMarketAppsOk returns a tuple with the MarketApps field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMarketApps

`func (o *ImageDetails) SetMarketApps(v []string)`

SetMarketApps sets MarketApps field to given value.


### GetStorageTypes

`func (o *ImageDetails) GetStorageTypes() []string`

GetStorageTypes returns the StorageTypes field if non-nil, zero value otherwise.

### GetStorageTypesOk

`func (o *ImageDetails) GetStorageTypesOk() (*[]string, bool)`

GetStorageTypesOk returns a tuple with the StorageTypes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStorageTypes

`func (o *ImageDetails) SetStorageTypes(v []string)`

SetStorageTypes sets StorageTypes field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


