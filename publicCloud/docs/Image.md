# Image

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

## Methods

### NewImage

`func NewImage(id string, name string, version string, family string, flavour string, ) *Image`

NewImage instantiates a new Image object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewImageWithDefaults

`func NewImageWithDefaults() *Image`

NewImageWithDefaults instantiates a new Image object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *Image) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *Image) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *Image) SetId(v string)`

SetId sets Id field to given value.


### GetName

`func (o *Image) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *Image) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *Image) SetName(v string)`

SetName sets Name field to given value.


### GetVersion

`func (o *Image) GetVersion() string`

GetVersion returns the Version field if non-nil, zero value otherwise.

### GetVersionOk

`func (o *Image) GetVersionOk() (*string, bool)`

GetVersionOk returns a tuple with the Version field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVersion

`func (o *Image) SetVersion(v string)`

SetVersion sets Version field to given value.


### GetFamily

`func (o *Image) GetFamily() string`

GetFamily returns the Family field if non-nil, zero value otherwise.

### GetFamilyOk

`func (o *Image) GetFamilyOk() (*string, bool)`

GetFamilyOk returns a tuple with the Family field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFamily

`func (o *Image) SetFamily(v string)`

SetFamily sets Family field to given value.


### GetFlavour

`func (o *Image) GetFlavour() string`

GetFlavour returns the Flavour field if non-nil, zero value otherwise.

### GetFlavourOk

`func (o *Image) GetFlavourOk() (*string, bool)`

GetFlavourOk returns a tuple with the Flavour field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFlavour

`func (o *Image) SetFlavour(v string)`

SetFlavour sets Flavour field to given value.


### GetArchitecture

`func (o *Image) GetArchitecture() string`

GetArchitecture returns the Architecture field if non-nil, zero value otherwise.

### GetArchitectureOk

`func (o *Image) GetArchitectureOk() (*string, bool)`

GetArchitectureOk returns a tuple with the Architecture field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetArchitecture

`func (o *Image) SetArchitecture(v string)`

SetArchitecture sets Architecture field to given value.

### HasArchitecture

`func (o *Image) HasArchitecture() bool`

HasArchitecture returns a boolean if a field has been set.

### SetArchitectureNil

`func (o *Image) SetArchitectureNil(b bool)`

 SetArchitectureNil sets the value for Architecture to be an explicit nil

### UnsetArchitecture
`func (o *Image) UnsetArchitecture()`

UnsetArchitecture ensures that no value is present for Architecture, not even an explicit nil
### GetState

`func (o *Image) GetState() string`

GetState returns the State field if non-nil, zero value otherwise.

### GetStateOk

`func (o *Image) GetStateOk() (*string, bool)`

GetStateOk returns a tuple with the State field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetState

`func (o *Image) SetState(v string)`

SetState sets State field to given value.

### HasState

`func (o *Image) HasState() bool`

HasState returns a boolean if a field has been set.

### SetStateNil

`func (o *Image) SetStateNil(b bool)`

 SetStateNil sets the value for State to be an explicit nil

### UnsetState
`func (o *Image) UnsetState()`

UnsetState ensures that no value is present for State, not even an explicit nil
### GetStateReason

`func (o *Image) GetStateReason() string`

GetStateReason returns the StateReason field if non-nil, zero value otherwise.

### GetStateReasonOk

`func (o *Image) GetStateReasonOk() (*string, bool)`

GetStateReasonOk returns a tuple with the StateReason field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStateReason

`func (o *Image) SetStateReason(v string)`

SetStateReason sets StateReason field to given value.

### HasStateReason

`func (o *Image) HasStateReason() bool`

HasStateReason returns a boolean if a field has been set.

### SetStateReasonNil

`func (o *Image) SetStateReasonNil(b bool)`

 SetStateReasonNil sets the value for StateReason to be an explicit nil

### UnsetStateReason
`func (o *Image) UnsetStateReason()`

UnsetStateReason ensures that no value is present for StateReason, not even an explicit nil
### GetRegion

`func (o *Image) GetRegion() string`

GetRegion returns the Region field if non-nil, zero value otherwise.

### GetRegionOk

`func (o *Image) GetRegionOk() (*string, bool)`

GetRegionOk returns a tuple with the Region field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRegion

`func (o *Image) SetRegion(v string)`

SetRegion sets Region field to given value.

### HasRegion

`func (o *Image) HasRegion() bool`

HasRegion returns a boolean if a field has been set.

### SetRegionNil

`func (o *Image) SetRegionNil(b bool)`

 SetRegionNil sets the value for Region to be an explicit nil

### UnsetRegion
`func (o *Image) UnsetRegion()`

UnsetRegion ensures that no value is present for Region, not even an explicit nil
### GetCreatedAt

`func (o *Image) GetCreatedAt() time.Time`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *Image) GetCreatedAtOk() (*time.Time, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *Image) SetCreatedAt(v time.Time)`

SetCreatedAt sets CreatedAt field to given value.

### HasCreatedAt

`func (o *Image) HasCreatedAt() bool`

HasCreatedAt returns a boolean if a field has been set.

### SetCreatedAtNil

`func (o *Image) SetCreatedAtNil(b bool)`

 SetCreatedAtNil sets the value for CreatedAt to be an explicit nil

### UnsetCreatedAt
`func (o *Image) UnsetCreatedAt()`

UnsetCreatedAt ensures that no value is present for CreatedAt, not even an explicit nil
### GetUpdatedAt

`func (o *Image) GetUpdatedAt() time.Time`

GetUpdatedAt returns the UpdatedAt field if non-nil, zero value otherwise.

### GetUpdatedAtOk

`func (o *Image) GetUpdatedAtOk() (*time.Time, bool)`

GetUpdatedAtOk returns a tuple with the UpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedAt

`func (o *Image) SetUpdatedAt(v time.Time)`

SetUpdatedAt sets UpdatedAt field to given value.

### HasUpdatedAt

`func (o *Image) HasUpdatedAt() bool`

HasUpdatedAt returns a boolean if a field has been set.

### SetUpdatedAtNil

`func (o *Image) SetUpdatedAtNil(b bool)`

 SetUpdatedAtNil sets the value for UpdatedAt to be an explicit nil

### UnsetUpdatedAt
`func (o *Image) UnsetUpdatedAt()`

UnsetUpdatedAt ensures that no value is present for UpdatedAt, not even an explicit nil
### GetCustom

`func (o *Image) GetCustom() bool`

GetCustom returns the Custom field if non-nil, zero value otherwise.

### GetCustomOk

`func (o *Image) GetCustomOk() (*bool, bool)`

GetCustomOk returns a tuple with the Custom field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustom

`func (o *Image) SetCustom(v bool)`

SetCustom sets Custom field to given value.

### HasCustom

`func (o *Image) HasCustom() bool`

HasCustom returns a boolean if a field has been set.

### SetCustomNil

`func (o *Image) SetCustomNil(b bool)`

 SetCustomNil sets the value for Custom to be an explicit nil

### UnsetCustom
`func (o *Image) UnsetCustom()`

UnsetCustom ensures that no value is present for Custom, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


