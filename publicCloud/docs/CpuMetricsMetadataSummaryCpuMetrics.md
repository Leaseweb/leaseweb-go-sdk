# CpuMetricsMetadataSummaryCpuMetrics

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Average** | Pointer to **string** | Average CPU based on the amount of grouped data points, in percentage | [optional] 
**Expected** | Pointer to **string** | Expected CPU given the average times the amount of days between the &#x60;from&#x60; and &#x60;to&#x60; dates, in percentage | [optional] 
**Peak** | Pointer to [**CpuMetricsMetadataSummaryPeak**](CpuMetricsMetadataSummaryPeak.md) |  | [optional] 

## Methods

### NewCpuMetricsMetadataSummaryCpuMetrics

`func NewCpuMetricsMetadataSummaryCpuMetrics() *CpuMetricsMetadataSummaryCpuMetrics`

NewCpuMetricsMetadataSummaryCpuMetrics instantiates a new CpuMetricsMetadataSummaryCpuMetrics object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCpuMetricsMetadataSummaryCpuMetricsWithDefaults

`func NewCpuMetricsMetadataSummaryCpuMetricsWithDefaults() *CpuMetricsMetadataSummaryCpuMetrics`

NewCpuMetricsMetadataSummaryCpuMetricsWithDefaults instantiates a new CpuMetricsMetadataSummaryCpuMetrics object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAverage

`func (o *CpuMetricsMetadataSummaryCpuMetrics) GetAverage() string`

GetAverage returns the Average field if non-nil, zero value otherwise.

### GetAverageOk

`func (o *CpuMetricsMetadataSummaryCpuMetrics) GetAverageOk() (*string, bool)`

GetAverageOk returns a tuple with the Average field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAverage

`func (o *CpuMetricsMetadataSummaryCpuMetrics) SetAverage(v string)`

SetAverage sets Average field to given value.

### HasAverage

`func (o *CpuMetricsMetadataSummaryCpuMetrics) HasAverage() bool`

HasAverage returns a boolean if a field has been set.

### GetExpected

`func (o *CpuMetricsMetadataSummaryCpuMetrics) GetExpected() string`

GetExpected returns the Expected field if non-nil, zero value otherwise.

### GetExpectedOk

`func (o *CpuMetricsMetadataSummaryCpuMetrics) GetExpectedOk() (*string, bool)`

GetExpectedOk returns a tuple with the Expected field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExpected

`func (o *CpuMetricsMetadataSummaryCpuMetrics) SetExpected(v string)`

SetExpected sets Expected field to given value.

### HasExpected

`func (o *CpuMetricsMetadataSummaryCpuMetrics) HasExpected() bool`

HasExpected returns a boolean if a field has been set.

### GetPeak

`func (o *CpuMetricsMetadataSummaryCpuMetrics) GetPeak() CpuMetricsMetadataSummaryPeak`

GetPeak returns the Peak field if non-nil, zero value otherwise.

### GetPeakOk

`func (o *CpuMetricsMetadataSummaryCpuMetrics) GetPeakOk() (*CpuMetricsMetadataSummaryPeak, bool)`

GetPeakOk returns a tuple with the Peak field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPeak

`func (o *CpuMetricsMetadataSummaryCpuMetrics) SetPeak(v CpuMetricsMetadataSummaryPeak)`

SetPeak sets Peak field to given value.

### HasPeak

`func (o *CpuMetricsMetadataSummaryCpuMetrics) HasPeak() bool`

HasPeak returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


