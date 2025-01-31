# MetricsMetadataProperties

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**From** | Pointer to **time.Time** |  | [optional] 
**To** | Pointer to **time.Time** |  | [optional] 
**Granularity** | Pointer to [**MetricsMetadataPropertiesGranularity**](MetricsMetadataPropertiesGranularity.md) |  | [optional] 
**Aggregation** | Pointer to **string** | Defined by the query | [optional] 

## Methods

### NewMetricsMetadataProperties

`func NewMetricsMetadataProperties() *MetricsMetadataProperties`

NewMetricsMetadataProperties instantiates a new MetricsMetadataProperties object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMetricsMetadataPropertiesWithDefaults

`func NewMetricsMetadataPropertiesWithDefaults() *MetricsMetadataProperties`

NewMetricsMetadataPropertiesWithDefaults instantiates a new MetricsMetadataProperties object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetFrom

`func (o *MetricsMetadataProperties) GetFrom() time.Time`

GetFrom returns the From field if non-nil, zero value otherwise.

### GetFromOk

`func (o *MetricsMetadataProperties) GetFromOk() (*time.Time, bool)`

GetFromOk returns a tuple with the From field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFrom

`func (o *MetricsMetadataProperties) SetFrom(v time.Time)`

SetFrom sets From field to given value.

### HasFrom

`func (o *MetricsMetadataProperties) HasFrom() bool`

HasFrom returns a boolean if a field has been set.

### GetTo

`func (o *MetricsMetadataProperties) GetTo() time.Time`

GetTo returns the To field if non-nil, zero value otherwise.

### GetToOk

`func (o *MetricsMetadataProperties) GetToOk() (*time.Time, bool)`

GetToOk returns a tuple with the To field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTo

`func (o *MetricsMetadataProperties) SetTo(v time.Time)`

SetTo sets To field to given value.

### HasTo

`func (o *MetricsMetadataProperties) HasTo() bool`

HasTo returns a boolean if a field has been set.

### GetGranularity

`func (o *MetricsMetadataProperties) GetGranularity() MetricsMetadataPropertiesGranularity`

GetGranularity returns the Granularity field if non-nil, zero value otherwise.

### GetGranularityOk

`func (o *MetricsMetadataProperties) GetGranularityOk() (*MetricsMetadataPropertiesGranularity, bool)`

GetGranularityOk returns a tuple with the Granularity field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGranularity

`func (o *MetricsMetadataProperties) SetGranularity(v MetricsMetadataPropertiesGranularity)`

SetGranularity sets Granularity field to given value.

### HasGranularity

`func (o *MetricsMetadataProperties) HasGranularity() bool`

HasGranularity returns a boolean if a field has been set.

### GetAggregation

`func (o *MetricsMetadataProperties) GetAggregation() string`

GetAggregation returns the Aggregation field if non-nil, zero value otherwise.

### GetAggregationOk

`func (o *MetricsMetadataProperties) GetAggregationOk() (*string, bool)`

GetAggregationOk returns a tuple with the Aggregation field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAggregation

`func (o *MetricsMetadataProperties) SetAggregation(v string)`

SetAggregation sets Aggregation field to given value.

### HasAggregation

`func (o *MetricsMetadataProperties) HasAggregation() bool`

HasAggregation returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


