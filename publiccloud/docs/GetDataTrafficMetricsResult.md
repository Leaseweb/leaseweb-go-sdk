# GetDataTrafficMetricsResult

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Metrics** | Pointer to [**DataTrafficMetrics**](DataTrafficMetrics.md) |  | [optional] 
**Metadata** | Pointer to [**TrafficMetricsMetaData**](TrafficMetricsMetaData.md) |  | [optional] 

## Methods

### NewGetDataTrafficMetricsResult

`func NewGetDataTrafficMetricsResult() *GetDataTrafficMetricsResult`

NewGetDataTrafficMetricsResult instantiates a new GetDataTrafficMetricsResult object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetDataTrafficMetricsResultWithDefaults

`func NewGetDataTrafficMetricsResultWithDefaults() *GetDataTrafficMetricsResult`

NewGetDataTrafficMetricsResultWithDefaults instantiates a new GetDataTrafficMetricsResult object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetMetrics

`func (o *GetDataTrafficMetricsResult) GetMetrics() DataTrafficMetrics`

GetMetrics returns the Metrics field if non-nil, zero value otherwise.

### GetMetricsOk

`func (o *GetDataTrafficMetricsResult) GetMetricsOk() (*DataTrafficMetrics, bool)`

GetMetricsOk returns a tuple with the Metrics field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetrics

`func (o *GetDataTrafficMetricsResult) SetMetrics(v DataTrafficMetrics)`

SetMetrics sets Metrics field to given value.

### HasMetrics

`func (o *GetDataTrafficMetricsResult) HasMetrics() bool`

HasMetrics returns a boolean if a field has been set.

### GetMetadata

`func (o *GetDataTrafficMetricsResult) GetMetadata() TrafficMetricsMetaData`

GetMetadata returns the Metadata field if non-nil, zero value otherwise.

### GetMetadataOk

`func (o *GetDataTrafficMetricsResult) GetMetadataOk() (*TrafficMetricsMetaData, bool)`

GetMetadataOk returns a tuple with the Metadata field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetadata

`func (o *GetDataTrafficMetricsResult) SetMetadata(v TrafficMetricsMetaData)`

SetMetadata sets Metadata field to given value.

### HasMetadata

`func (o *GetDataTrafficMetricsResult) HasMetadata() bool`

HasMetadata returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


