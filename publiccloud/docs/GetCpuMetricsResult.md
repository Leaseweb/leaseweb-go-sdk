# GetCpuMetricsResult

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Metrics** | Pointer to [**CpuMetrics**](CpuMetrics.md) |  | [optional] 
**Metadata** | Pointer to [**CpuMetricsMetadata**](CpuMetricsMetadata.md) |  | [optional] 

## Methods

### NewGetCpuMetricsResult

`func NewGetCpuMetricsResult() *GetCpuMetricsResult`

NewGetCpuMetricsResult instantiates a new GetCpuMetricsResult object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetCpuMetricsResultWithDefaults

`func NewGetCpuMetricsResultWithDefaults() *GetCpuMetricsResult`

NewGetCpuMetricsResultWithDefaults instantiates a new GetCpuMetricsResult object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetMetrics

`func (o *GetCpuMetricsResult) GetMetrics() CpuMetrics`

GetMetrics returns the Metrics field if non-nil, zero value otherwise.

### GetMetricsOk

`func (o *GetCpuMetricsResult) GetMetricsOk() (*CpuMetrics, bool)`

GetMetricsOk returns a tuple with the Metrics field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetrics

`func (o *GetCpuMetricsResult) SetMetrics(v CpuMetrics)`

SetMetrics sets Metrics field to given value.

### HasMetrics

`func (o *GetCpuMetricsResult) HasMetrics() bool`

HasMetrics returns a boolean if a field has been set.

### GetMetadata

`func (o *GetCpuMetricsResult) GetMetadata() CpuMetricsMetadata`

GetMetadata returns the Metadata field if non-nil, zero value otherwise.

### GetMetadataOk

`func (o *GetCpuMetricsResult) GetMetadataOk() (*CpuMetricsMetadata, bool)`

GetMetadataOk returns a tuple with the Metadata field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetadata

`func (o *GetCpuMetricsResult) SetMetadata(v CpuMetricsMetadata)`

SetMetadata sets Metadata field to given value.

### HasMetadata

`func (o *GetCpuMetricsResult) HasMetadata() bool`

HasMetadata returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


