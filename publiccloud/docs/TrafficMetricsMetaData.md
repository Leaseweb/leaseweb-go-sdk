# TrafficMetricsMetaData

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**From** | Pointer to **time.Time** |  | [optional] 
**To** | Pointer to **time.Time** |  | [optional] 
**Granularity** | Pointer to [**MetricsMetadataPropertiesGranularity**](MetricsMetadataPropertiesGranularity.md) |  | [optional] 
**Aggregation** | Pointer to **string** | Defined by the query | [optional] 
**Summary** | Pointer to [**Summary**](Summary.md) |  | [optional] 

## Methods

### NewTrafficMetricsMetaData

`func NewTrafficMetricsMetaData() *TrafficMetricsMetaData`

NewTrafficMetricsMetaData instantiates a new TrafficMetricsMetaData object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTrafficMetricsMetaDataWithDefaults

`func NewTrafficMetricsMetaDataWithDefaults() *TrafficMetricsMetaData`

NewTrafficMetricsMetaDataWithDefaults instantiates a new TrafficMetricsMetaData object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetFrom

`func (o *TrafficMetricsMetaData) GetFrom() time.Time`

GetFrom returns the From field if non-nil, zero value otherwise.

### GetFromOk

`func (o *TrafficMetricsMetaData) GetFromOk() (*time.Time, bool)`

GetFromOk returns a tuple with the From field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFrom

`func (o *TrafficMetricsMetaData) SetFrom(v time.Time)`

SetFrom sets From field to given value.

### HasFrom

`func (o *TrafficMetricsMetaData) HasFrom() bool`

HasFrom returns a boolean if a field has been set.

### GetTo

`func (o *TrafficMetricsMetaData) GetTo() time.Time`

GetTo returns the To field if non-nil, zero value otherwise.

### GetToOk

`func (o *TrafficMetricsMetaData) GetToOk() (*time.Time, bool)`

GetToOk returns a tuple with the To field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTo

`func (o *TrafficMetricsMetaData) SetTo(v time.Time)`

SetTo sets To field to given value.

### HasTo

`func (o *TrafficMetricsMetaData) HasTo() bool`

HasTo returns a boolean if a field has been set.

### GetGranularity

`func (o *TrafficMetricsMetaData) GetGranularity() MetricsMetadataPropertiesGranularity`

GetGranularity returns the Granularity field if non-nil, zero value otherwise.

### GetGranularityOk

`func (o *TrafficMetricsMetaData) GetGranularityOk() (*MetricsMetadataPropertiesGranularity, bool)`

GetGranularityOk returns a tuple with the Granularity field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGranularity

`func (o *TrafficMetricsMetaData) SetGranularity(v MetricsMetadataPropertiesGranularity)`

SetGranularity sets Granularity field to given value.

### HasGranularity

`func (o *TrafficMetricsMetaData) HasGranularity() bool`

HasGranularity returns a boolean if a field has been set.

### GetAggregation

`func (o *TrafficMetricsMetaData) GetAggregation() string`

GetAggregation returns the Aggregation field if non-nil, zero value otherwise.

### GetAggregationOk

`func (o *TrafficMetricsMetaData) GetAggregationOk() (*string, bool)`

GetAggregationOk returns a tuple with the Aggregation field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAggregation

`func (o *TrafficMetricsMetaData) SetAggregation(v string)`

SetAggregation sets Aggregation field to given value.

### HasAggregation

`func (o *TrafficMetricsMetaData) HasAggregation() bool`

HasAggregation returns a boolean if a field has been set.

### GetSummary

`func (o *TrafficMetricsMetaData) GetSummary() Summary`

GetSummary returns the Summary field if non-nil, zero value otherwise.

### GetSummaryOk

`func (o *TrafficMetricsMetaData) GetSummaryOk() (*Summary, bool)`

GetSummaryOk returns a tuple with the Summary field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSummary

`func (o *TrafficMetricsMetaData) SetSummary(v Summary)`

SetSummary sets Summary field to given value.

### HasSummary

`func (o *TrafficMetricsMetaData) HasSummary() bool`

HasSummary returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


