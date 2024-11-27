# MetricsMetadata

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Aggregation** | Pointer to **string** | The aggregation type for this request | [optional] 
**From** | Pointer to **time.Time** | Start of date interval in ISO-8601 format | [optional] 
**Granularity** | Pointer to **string** | The interval for each metric | [optional] 
**To** | Pointer to **time.Time** | End of date interval in ISO-8601 format | [optional] 

## Methods

### NewMetricsMetadata

`func NewMetricsMetadata() *MetricsMetadata`

NewMetricsMetadata instantiates a new MetricsMetadata object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMetricsMetadataWithDefaults

`func NewMetricsMetadataWithDefaults() *MetricsMetadata`

NewMetricsMetadataWithDefaults instantiates a new MetricsMetadata object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAggregation

`func (o *MetricsMetadata) GetAggregation() string`

GetAggregation returns the Aggregation field if non-nil, zero value otherwise.

### GetAggregationOk

`func (o *MetricsMetadata) GetAggregationOk() (*string, bool)`

GetAggregationOk returns a tuple with the Aggregation field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAggregation

`func (o *MetricsMetadata) SetAggregation(v string)`

SetAggregation sets Aggregation field to given value.

### HasAggregation

`func (o *MetricsMetadata) HasAggregation() bool`

HasAggregation returns a boolean if a field has been set.

### GetFrom

`func (o *MetricsMetadata) GetFrom() time.Time`

GetFrom returns the From field if non-nil, zero value otherwise.

### GetFromOk

`func (o *MetricsMetadata) GetFromOk() (*time.Time, bool)`

GetFromOk returns a tuple with the From field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFrom

`func (o *MetricsMetadata) SetFrom(v time.Time)`

SetFrom sets From field to given value.

### HasFrom

`func (o *MetricsMetadata) HasFrom() bool`

HasFrom returns a boolean if a field has been set.

### GetGranularity

`func (o *MetricsMetadata) GetGranularity() string`

GetGranularity returns the Granularity field if non-nil, zero value otherwise.

### GetGranularityOk

`func (o *MetricsMetadata) GetGranularityOk() (*string, bool)`

GetGranularityOk returns a tuple with the Granularity field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGranularity

`func (o *MetricsMetadata) SetGranularity(v string)`

SetGranularity sets Granularity field to given value.

### HasGranularity

`func (o *MetricsMetadata) HasGranularity() bool`

HasGranularity returns a boolean if a field has been set.

### GetTo

`func (o *MetricsMetadata) GetTo() time.Time`

GetTo returns the To field if non-nil, zero value otherwise.

### GetToOk

`func (o *MetricsMetadata) GetToOk() (*time.Time, bool)`

GetToOk returns a tuple with the To field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTo

`func (o *MetricsMetadata) SetTo(v time.Time)`

SetTo sets To field to given value.

### HasTo

`func (o *MetricsMetadata) HasTo() bool`

HasTo returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


