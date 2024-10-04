# GetDataTransferredPerSecondMetricsResult

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Metrics** | Pointer to [**DataTransferredMetrics**](DataTransferredMetrics.md) |  | [optional] 
**Metadata** | Pointer to [**MetricsMetadataProperties**](MetricsMetadataProperties.md) |  | [optional] 

## Methods

### NewGetDataTransferredPerSecondMetricsResult

`func NewGetDataTransferredPerSecondMetricsResult() *GetDataTransferredPerSecondMetricsResult`

NewGetDataTransferredPerSecondMetricsResult instantiates a new GetDataTransferredPerSecondMetricsResult object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetDataTransferredPerSecondMetricsResultWithDefaults

`func NewGetDataTransferredPerSecondMetricsResultWithDefaults() *GetDataTransferredPerSecondMetricsResult`

NewGetDataTransferredPerSecondMetricsResultWithDefaults instantiates a new GetDataTransferredPerSecondMetricsResult object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetMetrics

`func (o *GetDataTransferredPerSecondMetricsResult) GetMetrics() DataTransferredMetrics`

GetMetrics returns the Metrics field if non-nil, zero value otherwise.

### GetMetricsOk

`func (o *GetDataTransferredPerSecondMetricsResult) GetMetricsOk() (*DataTransferredMetrics, bool)`

GetMetricsOk returns a tuple with the Metrics field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetrics

`func (o *GetDataTransferredPerSecondMetricsResult) SetMetrics(v DataTransferredMetrics)`

SetMetrics sets Metrics field to given value.

### HasMetrics

`func (o *GetDataTransferredPerSecondMetricsResult) HasMetrics() bool`

HasMetrics returns a boolean if a field has been set.

### GetMetadata

`func (o *GetDataTransferredPerSecondMetricsResult) GetMetadata() MetricsMetadataProperties`

GetMetadata returns the Metadata field if non-nil, zero value otherwise.

### GetMetadataOk

`func (o *GetDataTransferredPerSecondMetricsResult) GetMetadataOk() (*MetricsMetadataProperties, bool)`

GetMetadataOk returns a tuple with the Metadata field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetadata

`func (o *GetDataTransferredPerSecondMetricsResult) SetMetadata(v MetricsMetadataProperties)`

SetMetadata sets Metadata field to given value.

### HasMetadata

`func (o *GetDataTransferredPerSecondMetricsResult) HasMetadata() bool`

HasMetadata returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


