# Ram

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Size** | Pointer to **int32** | The total RAM size of the server | [optional] 
**Unit** | Pointer to **string** | RAM type of the server | [optional] 

## Methods

### NewRam

`func NewRam() *Ram`

NewRam instantiates a new Ram object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRamWithDefaults

`func NewRamWithDefaults() *Ram`

NewRamWithDefaults instantiates a new Ram object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSize

`func (o *Ram) GetSize() int32`

GetSize returns the Size field if non-nil, zero value otherwise.

### GetSizeOk

`func (o *Ram) GetSizeOk() (*int32, bool)`

GetSizeOk returns a tuple with the Size field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSize

`func (o *Ram) SetSize(v int32)`

SetSize sets Size field to given value.

### HasSize

`func (o *Ram) HasSize() bool`

HasSize returns a boolean if a field has been set.

### GetUnit

`func (o *Ram) GetUnit() string`

GetUnit returns the Unit field if non-nil, zero value otherwise.

### GetUnitOk

`func (o *Ram) GetUnitOk() (*string, bool)`

GetUnitOk returns a tuple with the Unit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUnit

`func (o *Ram) SetUnit(v string)`

SetUnit sets Unit field to given value.

### HasUnit

`func (o *Ram) HasUnit() bool`

HasUnit returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


