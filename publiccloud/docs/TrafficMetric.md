# TrafficMetric

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Values** | Pointer to [**[]TrafficMetricValue**](TrafficMetricValue.md) |  | [optional] 
**Unit** | Pointer to [**MetricsUnit**](MetricsUnit.md) |  | [optional] 

## Methods

### NewTrafficMetric

`func NewTrafficMetric() *TrafficMetric`

NewTrafficMetric instantiates a new TrafficMetric object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTrafficMetricWithDefaults

`func NewTrafficMetricWithDefaults() *TrafficMetric`

NewTrafficMetricWithDefaults instantiates a new TrafficMetric object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetValues

`func (o *TrafficMetric) GetValues() []TrafficMetricValue`

GetValues returns the Values field if non-nil, zero value otherwise.

### GetValuesOk

`func (o *TrafficMetric) GetValuesOk() (*[]TrafficMetricValue, bool)`

GetValuesOk returns a tuple with the Values field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValues

`func (o *TrafficMetric) SetValues(v []TrafficMetricValue)`

SetValues sets Values field to given value.

### HasValues

`func (o *TrafficMetric) HasValues() bool`

HasValues returns a boolean if a field has been set.

### GetUnit

`func (o *TrafficMetric) GetUnit() MetricsUnit`

GetUnit returns the Unit field if non-nil, zero value otherwise.

### GetUnitOk

`func (o *TrafficMetric) GetUnitOk() (*MetricsUnit, bool)`

GetUnitOk returns a tuple with the Unit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUnit

`func (o *TrafficMetric) SetUnit(v MetricsUnit)`

SetUnit sets Unit field to given value.

### HasUnit

`func (o *TrafficMetric) HasUnit() bool`

HasUnit returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


