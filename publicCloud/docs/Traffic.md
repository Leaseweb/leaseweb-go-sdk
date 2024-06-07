# Traffic

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Unit** | Pointer to **string** |  | [optional] 
**Values** | Pointer to [**Values**](Values.md) |  | [optional] 

## Methods

### NewTraffic

`func NewTraffic() *Traffic`

NewTraffic instantiates a new Traffic object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTrafficWithDefaults

`func NewTrafficWithDefaults() *Traffic`

NewTrafficWithDefaults instantiates a new Traffic object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetUnit

`func (o *Traffic) GetUnit() string`

GetUnit returns the Unit field if non-nil, zero value otherwise.

### GetUnitOk

`func (o *Traffic) GetUnitOk() (*string, bool)`

GetUnitOk returns a tuple with the Unit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUnit

`func (o *Traffic) SetUnit(v string)`

SetUnit sets Unit field to given value.

### HasUnit

`func (o *Traffic) HasUnit() bool`

HasUnit returns a boolean if a field has been set.

### GetValues

`func (o *Traffic) GetValues() Values`

GetValues returns the Values field if non-nil, zero value otherwise.

### GetValuesOk

`func (o *Traffic) GetValuesOk() (*Values, bool)`

GetValuesOk returns a tuple with the Values field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValues

`func (o *Traffic) SetValues(v Values)`

SetValues sets Values field to given value.

### HasValues

`func (o *Traffic) HasValues() bool`

HasValues returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


