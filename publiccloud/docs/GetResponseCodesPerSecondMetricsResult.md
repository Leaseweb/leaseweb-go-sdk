# GetResponseCodesPerSecondMetricsResult

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Metrics** | Pointer to [**ResponseCodesMetrics**](ResponseCodesMetrics.md) |  | [optional] 
**Metadata** | Pointer to [**MetricsMetadataProperties**](MetricsMetadataProperties.md) |  | [optional] 

## Methods

### NewGetResponseCodesPerSecondMetricsResult

`func NewGetResponseCodesPerSecondMetricsResult() *GetResponseCodesPerSecondMetricsResult`

NewGetResponseCodesPerSecondMetricsResult instantiates a new GetResponseCodesPerSecondMetricsResult object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetResponseCodesPerSecondMetricsResultWithDefaults

`func NewGetResponseCodesPerSecondMetricsResultWithDefaults() *GetResponseCodesPerSecondMetricsResult`

NewGetResponseCodesPerSecondMetricsResultWithDefaults instantiates a new GetResponseCodesPerSecondMetricsResult object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetMetrics

`func (o *GetResponseCodesPerSecondMetricsResult) GetMetrics() ResponseCodesMetrics`

GetMetrics returns the Metrics field if non-nil, zero value otherwise.

### GetMetricsOk

`func (o *GetResponseCodesPerSecondMetricsResult) GetMetricsOk() (*ResponseCodesMetrics, bool)`

GetMetricsOk returns a tuple with the Metrics field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetrics

`func (o *GetResponseCodesPerSecondMetricsResult) SetMetrics(v ResponseCodesMetrics)`

SetMetrics sets Metrics field to given value.

### HasMetrics

`func (o *GetResponseCodesPerSecondMetricsResult) HasMetrics() bool`

HasMetrics returns a boolean if a field has been set.

### GetMetadata

`func (o *GetResponseCodesPerSecondMetricsResult) GetMetadata() MetricsMetadataProperties`

GetMetadata returns the Metadata field if non-nil, zero value otherwise.

### GetMetadataOk

`func (o *GetResponseCodesPerSecondMetricsResult) GetMetadataOk() (*MetricsMetadataProperties, bool)`

GetMetadataOk returns a tuple with the Metadata field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetadata

`func (o *GetResponseCodesPerSecondMetricsResult) SetMetadata(v MetricsMetadataProperties)`

SetMetadata sets Metadata field to given value.

### HasMetadata

`func (o *GetResponseCodesPerSecondMetricsResult) HasMetadata() bool`

HasMetadata returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


