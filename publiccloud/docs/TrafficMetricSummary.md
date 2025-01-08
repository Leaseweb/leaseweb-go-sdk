# TrafficMetricSummary

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Average** | Pointer to **float32** | Average incoming traffic based on the amount of grouped data points, in bytes | [optional] 
**Expected** | Pointer to **float32** | Expected incoming traffic given the average times the amount of days between the &#x60;from&#x60; and &#x60;to&#x60; dates, in bytes | [optional] 
**Total** | Pointer to **float32** | Total incoming traffic, in bytes | [optional] 
**Peak** | Pointer to [**Peak**](Peak.md) |  | [optional] 

## Methods

### NewTrafficMetricSummary

`func NewTrafficMetricSummary() *TrafficMetricSummary`

NewTrafficMetricSummary instantiates a new TrafficMetricSummary object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTrafficMetricSummaryWithDefaults

`func NewTrafficMetricSummaryWithDefaults() *TrafficMetricSummary`

NewTrafficMetricSummaryWithDefaults instantiates a new TrafficMetricSummary object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAverage

`func (o *TrafficMetricSummary) GetAverage() float32`

GetAverage returns the Average field if non-nil, zero value otherwise.

### GetAverageOk

`func (o *TrafficMetricSummary) GetAverageOk() (*float32, bool)`

GetAverageOk returns a tuple with the Average field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAverage

`func (o *TrafficMetricSummary) SetAverage(v float32)`

SetAverage sets Average field to given value.

### HasAverage

`func (o *TrafficMetricSummary) HasAverage() bool`

HasAverage returns a boolean if a field has been set.

### GetExpected

`func (o *TrafficMetricSummary) GetExpected() float32`

GetExpected returns the Expected field if non-nil, zero value otherwise.

### GetExpectedOk

`func (o *TrafficMetricSummary) GetExpectedOk() (*float32, bool)`

GetExpectedOk returns a tuple with the Expected field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExpected

`func (o *TrafficMetricSummary) SetExpected(v float32)`

SetExpected sets Expected field to given value.

### HasExpected

`func (o *TrafficMetricSummary) HasExpected() bool`

HasExpected returns a boolean if a field has been set.

### GetTotal

`func (o *TrafficMetricSummary) GetTotal() float32`

GetTotal returns the Total field if non-nil, zero value otherwise.

### GetTotalOk

`func (o *TrafficMetricSummary) GetTotalOk() (*float32, bool)`

GetTotalOk returns a tuple with the Total field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotal

`func (o *TrafficMetricSummary) SetTotal(v float32)`

SetTotal sets Total field to given value.

### HasTotal

`func (o *TrafficMetricSummary) HasTotal() bool`

HasTotal returns a boolean if a field has been set.

### GetPeak

`func (o *TrafficMetricSummary) GetPeak() Peak`

GetPeak returns the Peak field if non-nil, zero value otherwise.

### GetPeakOk

`func (o *TrafficMetricSummary) GetPeakOk() (*Peak, bool)`

GetPeakOk returns a tuple with the Peak field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPeak

`func (o *TrafficMetricSummary) SetPeak(v Peak)`

SetPeak sets Peak field to given value.

### HasPeak

`func (o *TrafficMetricSummary) HasPeak() bool`

HasPeak returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


