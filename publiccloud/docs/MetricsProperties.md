# MetricsProperties

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Values** | Pointer to [**[]MetricsValues**](MetricsValues.md) |  | [optional] 
**Unit** | Pointer to [**MetricsUnit**](MetricsUnit.md) |  | [optional] 

## Methods

### NewMetricsProperties

`func NewMetricsProperties() *MetricsProperties`

NewMetricsProperties instantiates a new MetricsProperties object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMetricsPropertiesWithDefaults

`func NewMetricsPropertiesWithDefaults() *MetricsProperties`

NewMetricsPropertiesWithDefaults instantiates a new MetricsProperties object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetValues

`func (o *MetricsProperties) GetValues() []MetricsValues`

GetValues returns the Values field if non-nil, zero value otherwise.

### GetValuesOk

`func (o *MetricsProperties) GetValuesOk() (*[]MetricsValues, bool)`

GetValuesOk returns a tuple with the Values field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValues

`func (o *MetricsProperties) SetValues(v []MetricsValues)`

SetValues sets Values field to given value.

### HasValues

`func (o *MetricsProperties) HasValues() bool`

HasValues returns a boolean if a field has been set.

### GetUnit

`func (o *MetricsProperties) GetUnit() MetricsUnit`

GetUnit returns the Unit field if non-nil, zero value otherwise.

### GetUnitOk

`func (o *MetricsProperties) GetUnitOk() (*MetricsUnit, bool)`

GetUnitOk returns a tuple with the Unit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUnit

`func (o *MetricsProperties) SetUnit(v MetricsUnit)`

SetUnit sets Unit field to given value.

### HasUnit

`func (o *MetricsProperties) HasUnit() bool`

HasUnit returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


