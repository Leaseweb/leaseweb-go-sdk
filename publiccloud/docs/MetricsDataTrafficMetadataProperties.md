# MetricsDataTrafficMetadataProperties

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**From** | Pointer to **time.Time** |  | [optional] 
**To** | Pointer to **time.Time** |  | [optional] 
**Granularity** | Pointer to [**DataTrafficMetricsGranularity**](DataTrafficMetricsGranularity.md) |  | [optional] 
**Aggregation** | Pointer to **string** | Defined by the query | [optional] 
**Unit** | Pointer to **string** | The unit of the summary values | [optional] 

## Methods

### NewMetricsDataTrafficMetadataProperties

`func NewMetricsDataTrafficMetadataProperties() *MetricsDataTrafficMetadataProperties`

NewMetricsDataTrafficMetadataProperties instantiates a new MetricsDataTrafficMetadataProperties object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMetricsDataTrafficMetadataPropertiesWithDefaults

`func NewMetricsDataTrafficMetadataPropertiesWithDefaults() *MetricsDataTrafficMetadataProperties`

NewMetricsDataTrafficMetadataPropertiesWithDefaults instantiates a new MetricsDataTrafficMetadataProperties object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetFrom

`func (o *MetricsDataTrafficMetadataProperties) GetFrom() time.Time`

GetFrom returns the From field if non-nil, zero value otherwise.

### GetFromOk

`func (o *MetricsDataTrafficMetadataProperties) GetFromOk() (*time.Time, bool)`

GetFromOk returns a tuple with the From field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFrom

`func (o *MetricsDataTrafficMetadataProperties) SetFrom(v time.Time)`

SetFrom sets From field to given value.

### HasFrom

`func (o *MetricsDataTrafficMetadataProperties) HasFrom() bool`

HasFrom returns a boolean if a field has been set.

### GetTo

`func (o *MetricsDataTrafficMetadataProperties) GetTo() time.Time`

GetTo returns the To field if non-nil, zero value otherwise.

### GetToOk

`func (o *MetricsDataTrafficMetadataProperties) GetToOk() (*time.Time, bool)`

GetToOk returns a tuple with the To field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTo

`func (o *MetricsDataTrafficMetadataProperties) SetTo(v time.Time)`

SetTo sets To field to given value.

### HasTo

`func (o *MetricsDataTrafficMetadataProperties) HasTo() bool`

HasTo returns a boolean if a field has been set.

### GetGranularity

`func (o *MetricsDataTrafficMetadataProperties) GetGranularity() DataTrafficMetricsGranularity`

GetGranularity returns the Granularity field if non-nil, zero value otherwise.

### GetGranularityOk

`func (o *MetricsDataTrafficMetadataProperties) GetGranularityOk() (*DataTrafficMetricsGranularity, bool)`

GetGranularityOk returns a tuple with the Granularity field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGranularity

`func (o *MetricsDataTrafficMetadataProperties) SetGranularity(v DataTrafficMetricsGranularity)`

SetGranularity sets Granularity field to given value.

### HasGranularity

`func (o *MetricsDataTrafficMetadataProperties) HasGranularity() bool`

HasGranularity returns a boolean if a field has been set.

### GetAggregation

`func (o *MetricsDataTrafficMetadataProperties) GetAggregation() string`

GetAggregation returns the Aggregation field if non-nil, zero value otherwise.

### GetAggregationOk

`func (o *MetricsDataTrafficMetadataProperties) GetAggregationOk() (*string, bool)`

GetAggregationOk returns a tuple with the Aggregation field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAggregation

`func (o *MetricsDataTrafficMetadataProperties) SetAggregation(v string)`

SetAggregation sets Aggregation field to given value.

### HasAggregation

`func (o *MetricsDataTrafficMetadataProperties) HasAggregation() bool`

HasAggregation returns a boolean if a field has been set.

### GetUnit

`func (o *MetricsDataTrafficMetadataProperties) GetUnit() string`

GetUnit returns the Unit field if non-nil, zero value otherwise.

### GetUnitOk

`func (o *MetricsDataTrafficMetadataProperties) GetUnitOk() (*string, bool)`

GetUnitOk returns a tuple with the Unit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUnit

`func (o *MetricsDataTrafficMetadataProperties) SetUnit(v string)`

SetUnit sets Unit field to given value.

### HasUnit

`func (o *MetricsDataTrafficMetadataProperties) HasUnit() bool`

HasUnit returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


