# CpuMetricsValue

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Value** | Pointer to **float32** | CPU usage | [optional] 
**Timestamp** | Pointer to **time.Time** |  | [optional] 

## Methods

### NewCpuMetricsValue

`func NewCpuMetricsValue() *CpuMetricsValue`

NewCpuMetricsValue instantiates a new CpuMetricsValue object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCpuMetricsValueWithDefaults

`func NewCpuMetricsValueWithDefaults() *CpuMetricsValue`

NewCpuMetricsValueWithDefaults instantiates a new CpuMetricsValue object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetValue

`func (o *CpuMetricsValue) GetValue() float32`

GetValue returns the Value field if non-nil, zero value otherwise.

### GetValueOk

`func (o *CpuMetricsValue) GetValueOk() (*float32, bool)`

GetValueOk returns a tuple with the Value field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValue

`func (o *CpuMetricsValue) SetValue(v float32)`

SetValue sets Value field to given value.

### HasValue

`func (o *CpuMetricsValue) HasValue() bool`

HasValue returns a boolean if a field has been set.

### GetTimestamp

`func (o *CpuMetricsValue) GetTimestamp() time.Time`

GetTimestamp returns the Timestamp field if non-nil, zero value otherwise.

### GetTimestampOk

`func (o *CpuMetricsValue) GetTimestampOk() (*time.Time, bool)`

GetTimestampOk returns a tuple with the Timestamp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimestamp

`func (o *CpuMetricsValue) SetTimestamp(v time.Time)`

SetTimestamp sets Timestamp field to given value.

### HasTimestamp

`func (o *CpuMetricsValue) HasTimestamp() bool`

HasTimestamp returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


