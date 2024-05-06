# Metric

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Unit** | Pointer to **string** | The metric unit that&#39;s used | [optional] 
**Values** | Pointer to [**[]MetricValue**](MetricValue.md) |  | [optional] 

## Methods

### NewMetric

`func NewMetric() *Metric`

NewMetric instantiates a new Metric object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMetricWithDefaults

`func NewMetricWithDefaults() *Metric`

NewMetricWithDefaults instantiates a new Metric object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetUnit

`func (o *Metric) GetUnit() string`

GetUnit returns the Unit field if non-nil, zero value otherwise.

### GetUnitOk

`func (o *Metric) GetUnitOk() (*string, bool)`

GetUnitOk returns a tuple with the Unit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUnit

`func (o *Metric) SetUnit(v string)`

SetUnit sets Unit field to given value.

### HasUnit

`func (o *Metric) HasUnit() bool`

HasUnit returns a boolean if a field has been set.

### GetValues

`func (o *Metric) GetValues() []MetricValue`

GetValues returns the Values field if non-nil, zero value otherwise.

### GetValuesOk

`func (o *Metric) GetValuesOk() (*[]MetricValue, bool)`

GetValuesOk returns a tuple with the Values field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValues

`func (o *Metric) SetValues(v []MetricValue)`

SetValues sets Values field to given value.

### HasValues

`func (o *Metric) HasValues() bool`

HasValues returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


