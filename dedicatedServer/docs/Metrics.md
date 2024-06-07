# Metrics

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Metadata** | Pointer to [**MetricsMetadata**](MetricsMetadata.md) |  | [optional] 
**Metrics** | Pointer to [**MetricValues**](MetricValues.md) |  | [optional] 

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

### GetMetadata

`func (o *Metrics) GetMetadata() MetricsMetadata`

GetMetadata returns the Metadata field if non-nil, zero value otherwise.

### GetMetadataOk

`func (o *Metrics) GetMetadataOk() (*MetricsMetadata, bool)`

GetMetadataOk returns a tuple with the Metadata field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetadata

`func (o *Metrics) SetMetadata(v MetricsMetadata)`

SetMetadata sets Metadata field to given value.

### HasMetadata

`func (o *Metrics) HasMetadata() bool`

HasMetadata returns a boolean if a field has been set.

### GetMetrics

`func (o *Metrics) GetMetrics() MetricValues`

GetMetrics returns the Metrics field if non-nil, zero value otherwise.

### GetMetricsOk

`func (o *Metrics) GetMetricsOk() (*MetricValues, bool)`

GetMetricsOk returns a tuple with the Metrics field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetrics

`func (o *Metrics) SetMetrics(v MetricValues)`

SetMetrics sets Metrics field to given value.

### HasMetrics

`func (o *Metrics) HasMetrics() bool`

HasMetrics returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


