# TrafficMetricSummary

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Average** | Pointer to **string** | Average incoming traffic based on the amount of grouped data points, in human-readable format (KB, MB, GB) | [optional] 
**Expected** | Pointer to **string** | Expected incoming traffic given the average times the amount of days between the &#x60;from&#x60; and &#x60;to&#x60; dates, in human-readable format (KB, MB, GB) | [optional] 
**Total** | Pointer to **string** | Total incoming traffic, in human-readable format (KB, MB, GB) | [optional] 
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

`func (o *TrafficMetricSummary) GetAverage() string`

GetAverage returns the Average field if non-nil, zero value otherwise.

### GetAverageOk

`func (o *TrafficMetricSummary) GetAverageOk() (*string, bool)`

GetAverageOk returns a tuple with the Average field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAverage

`func (o *TrafficMetricSummary) SetAverage(v string)`

SetAverage sets Average field to given value.

### HasAverage

`func (o *TrafficMetricSummary) HasAverage() bool`

HasAverage returns a boolean if a field has been set.

### GetExpected

`func (o *TrafficMetricSummary) GetExpected() string`

GetExpected returns the Expected field if non-nil, zero value otherwise.

### GetExpectedOk

`func (o *TrafficMetricSummary) GetExpectedOk() (*string, bool)`

GetExpectedOk returns a tuple with the Expected field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExpected

`func (o *TrafficMetricSummary) SetExpected(v string)`

SetExpected sets Expected field to given value.

### HasExpected

`func (o *TrafficMetricSummary) HasExpected() bool`

HasExpected returns a boolean if a field has been set.

### GetTotal

`func (o *TrafficMetricSummary) GetTotal() string`

GetTotal returns the Total field if non-nil, zero value otherwise.

### GetTotalOk

`func (o *TrafficMetricSummary) GetTotalOk() (*string, bool)`

GetTotalOk returns a tuple with the Total field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotal

`func (o *TrafficMetricSummary) SetTotal(v string)`

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


