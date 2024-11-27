# Hdd

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **string** | Id of the hard disk drive | [optional] 
**Amount** | Pointer to **int32** | The total amount of hard disk drives | [optional] 
**Size** | Pointer to **float32** | The size number of the hard disk drive | [optional] 
**Type** | Pointer to **string** | The type of the hard disk drive | [optional] 
**Unit** | Pointer to **string** | The unit of the hard disk drive | [optional] 
**PerformanceType** | Pointer to **NullableString** | Hard disk drive performance type | [optional] 

## Methods

### NewHdd

`func NewHdd() *Hdd`

NewHdd instantiates a new Hdd object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewHddWithDefaults

`func NewHddWithDefaults() *Hdd`

NewHddWithDefaults instantiates a new Hdd object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *Hdd) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *Hdd) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *Hdd) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *Hdd) HasId() bool`

HasId returns a boolean if a field has been set.

### GetAmount

`func (o *Hdd) GetAmount() int32`

GetAmount returns the Amount field if non-nil, zero value otherwise.

### GetAmountOk

`func (o *Hdd) GetAmountOk() (*int32, bool)`

GetAmountOk returns a tuple with the Amount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAmount

`func (o *Hdd) SetAmount(v int32)`

SetAmount sets Amount field to given value.

### HasAmount

`func (o *Hdd) HasAmount() bool`

HasAmount returns a boolean if a field has been set.

### GetSize

`func (o *Hdd) GetSize() float32`

GetSize returns the Size field if non-nil, zero value otherwise.

### GetSizeOk

`func (o *Hdd) GetSizeOk() (*float32, bool)`

GetSizeOk returns a tuple with the Size field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSize

`func (o *Hdd) SetSize(v float32)`

SetSize sets Size field to given value.

### HasSize

`func (o *Hdd) HasSize() bool`

HasSize returns a boolean if a field has been set.

### GetType

`func (o *Hdd) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *Hdd) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *Hdd) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *Hdd) HasType() bool`

HasType returns a boolean if a field has been set.

### GetUnit

`func (o *Hdd) GetUnit() string`

GetUnit returns the Unit field if non-nil, zero value otherwise.

### GetUnitOk

`func (o *Hdd) GetUnitOk() (*string, bool)`

GetUnitOk returns a tuple with the Unit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUnit

`func (o *Hdd) SetUnit(v string)`

SetUnit sets Unit field to given value.

### HasUnit

`func (o *Hdd) HasUnit() bool`

HasUnit returns a boolean if a field has been set.

### GetPerformanceType

`func (o *Hdd) GetPerformanceType() string`

GetPerformanceType returns the PerformanceType field if non-nil, zero value otherwise.

### GetPerformanceTypeOk

`func (o *Hdd) GetPerformanceTypeOk() (*string, bool)`

GetPerformanceTypeOk returns a tuple with the PerformanceType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPerformanceType

`func (o *Hdd) SetPerformanceType(v string)`

SetPerformanceType sets PerformanceType field to given value.

### HasPerformanceType

`func (o *Hdd) HasPerformanceType() bool`

HasPerformanceType returns a boolean if a field has been set.

### SetPerformanceTypeNil

`func (o *Hdd) SetPerformanceTypeNil(b bool)`

 SetPerformanceTypeNil sets the value for PerformanceType to be an explicit nil

### UnsetPerformanceType
`func (o *Hdd) UnsetPerformanceType()`

UnsetPerformanceType ensures that no value is present for PerformanceType, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


