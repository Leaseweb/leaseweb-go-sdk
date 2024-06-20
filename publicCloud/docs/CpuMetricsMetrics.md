# CpuMetricsMetrics

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Values** | Pointer to [**[]CpuMetricsValue**](CpuMetricsValue.md) |  | [optional] 
**Unit** | Pointer to **string** |  | [optional] 

## Methods

### NewCpuMetricsMetrics

`func NewCpuMetricsMetrics() *CpuMetricsMetrics`

NewCpuMetricsMetrics instantiates a new CpuMetricsMetrics object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCpuMetricsMetricsWithDefaults

`func NewCpuMetricsMetricsWithDefaults() *CpuMetricsMetrics`

NewCpuMetricsMetricsWithDefaults instantiates a new CpuMetricsMetrics object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetValues

`func (o *CpuMetricsMetrics) GetValues() []CpuMetricsValue`

GetValues returns the Values field if non-nil, zero value otherwise.

### GetValuesOk

`func (o *CpuMetricsMetrics) GetValuesOk() (*[]CpuMetricsValue, bool)`

GetValuesOk returns a tuple with the Values field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValues

`func (o *CpuMetricsMetrics) SetValues(v []CpuMetricsValue)`

SetValues sets Values field to given value.

### HasValues

`func (o *CpuMetricsMetrics) HasValues() bool`

HasValues returns a boolean if a field has been set.

### GetUnit

`func (o *CpuMetricsMetrics) GetUnit() string`

GetUnit returns the Unit field if non-nil, zero value otherwise.

### GetUnitOk

`func (o *CpuMetricsMetrics) GetUnitOk() (*string, bool)`

GetUnitOk returns a tuple with the Unit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUnit

`func (o *CpuMetricsMetrics) SetUnit(v string)`

SetUnit sets Unit field to given value.

### HasUnit

`func (o *CpuMetricsMetrics) HasUnit() bool`

HasUnit returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


