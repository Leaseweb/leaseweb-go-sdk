# GetResponseCodesMetricsResult

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Metrics** | Pointer to [**ResponseCodesMetrics**](ResponseCodesMetrics.md) |  | [optional] 
**Metadata** | Pointer to [**MetricsMetadataProperties**](MetricsMetadataProperties.md) |  | [optional] 

## Methods

### NewGetResponseCodesMetricsResult

`func NewGetResponseCodesMetricsResult() *GetResponseCodesMetricsResult`

NewGetResponseCodesMetricsResult instantiates a new GetResponseCodesMetricsResult object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetResponseCodesMetricsResultWithDefaults

`func NewGetResponseCodesMetricsResultWithDefaults() *GetResponseCodesMetricsResult`

NewGetResponseCodesMetricsResultWithDefaults instantiates a new GetResponseCodesMetricsResult object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetMetrics

`func (o *GetResponseCodesMetricsResult) GetMetrics() ResponseCodesMetrics`

GetMetrics returns the Metrics field if non-nil, zero value otherwise.

### GetMetricsOk

`func (o *GetResponseCodesMetricsResult) GetMetricsOk() (*ResponseCodesMetrics, bool)`

GetMetricsOk returns a tuple with the Metrics field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetrics

`func (o *GetResponseCodesMetricsResult) SetMetrics(v ResponseCodesMetrics)`

SetMetrics sets Metrics field to given value.

### HasMetrics

`func (o *GetResponseCodesMetricsResult) HasMetrics() bool`

HasMetrics returns a boolean if a field has been set.

### GetMetadata

`func (o *GetResponseCodesMetricsResult) GetMetadata() MetricsMetadataProperties`

GetMetadata returns the Metadata field if non-nil, zero value otherwise.

### GetMetadataOk

`func (o *GetResponseCodesMetricsResult) GetMetadataOk() (*MetricsMetadataProperties, bool)`

GetMetadataOk returns a tuple with the Metadata field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetadata

`func (o *GetResponseCodesMetricsResult) SetMetadata(v MetricsMetadataProperties)`

SetMetadata sets Metadata field to given value.

### HasMetadata

`func (o *GetResponseCodesMetricsResult) HasMetadata() bool`

HasMetadata returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


