# Metrics

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**DOWN_PUBLIC** | Pointer to [**TrafficMetric**](TrafficMetric.md) |  | [optional] 
**UP_PUBLIC** | Pointer to [**TrafficMetric**](TrafficMetric.md) |  | [optional] 

## Methods

### NewMetrics

`func NewMetrics() *Metrics`

NewMetrics instantiates a new Metrics object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMetricsWithDefaults

`func NewMetricsWithDefaults() *Metrics`

NewMetricsWithDefaults instantiates a new Metrics object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDOWN_PUBLIC

`func (o *Metrics) GetDOWN_PUBLIC() TrafficMetric`

GetDOWN_PUBLIC returns the DOWN_PUBLIC field if non-nil, zero value otherwise.

### GetDOWN_PUBLICOk

`func (o *Metrics) GetDOWN_PUBLICOk() (*TrafficMetric, bool)`

GetDOWN_PUBLICOk returns a tuple with the DOWN_PUBLIC field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDOWN_PUBLIC

`func (o *Metrics) SetDOWN_PUBLIC(v TrafficMetric)`

SetDOWN_PUBLIC sets DOWN_PUBLIC field to given value.

### HasDOWN_PUBLIC

`func (o *Metrics) HasDOWN_PUBLIC() bool`

HasDOWN_PUBLIC returns a boolean if a field has been set.

### GetUP_PUBLIC

`func (o *Metrics) GetUP_PUBLIC() TrafficMetric`

GetUP_PUBLIC returns the UP_PUBLIC field if non-nil, zero value otherwise.

### GetUP_PUBLICOk

`func (o *Metrics) GetUP_PUBLICOk() (*TrafficMetric, bool)`

GetUP_PUBLICOk returns a tuple with the UP_PUBLIC field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUP_PUBLIC

`func (o *Metrics) SetUP_PUBLIC(v TrafficMetric)`

SetUP_PUBLIC sets UP_PUBLIC field to given value.

### HasUP_PUBLIC

`func (o *Metrics) HasUP_PUBLIC() bool`

HasUP_PUBLIC returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


