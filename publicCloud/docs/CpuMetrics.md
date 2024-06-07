# CpuMetrics

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Values** | Pointer to [**[]CpuMetricsValue**](CpuMetricsValue.md) |  | [optional] 
**Unit** | Pointer to **string** |  | [optional] 

## Methods

### NewCpuMetrics

`func NewCpuMetrics() *CpuMetrics`

NewCpuMetrics instantiates a new CpuMetrics object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCpuMetricsWithDefaults

`func NewCpuMetricsWithDefaults() *CpuMetrics`

NewCpuMetricsWithDefaults instantiates a new CpuMetrics object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetValues

`func (o *CpuMetrics) GetValues() []CpuMetricsValue`

GetValues returns the Values field if non-nil, zero value otherwise.

### GetValuesOk

`func (o *CpuMetrics) GetValuesOk() (*[]CpuMetricsValue, bool)`

GetValuesOk returns a tuple with the Values field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValues

`func (o *CpuMetrics) SetValues(v []CpuMetricsValue)`

SetValues sets Values field to given value.

### HasValues

`func (o *CpuMetrics) HasValues() bool`

HasValues returns a boolean if a field has been set.

### GetUnit

`func (o *CpuMetrics) GetUnit() string`

GetUnit returns the Unit field if non-nil, zero value otherwise.

### GetUnitOk

`func (o *CpuMetrics) GetUnitOk() (*string, bool)`

GetUnitOk returns a tuple with the Unit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUnit

`func (o *CpuMetrics) SetUnit(v string)`

SetUnit sets Unit field to given value.

### HasUnit

`func (o *CpuMetrics) HasUnit() bool`

HasUnit returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


