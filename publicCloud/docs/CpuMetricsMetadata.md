# CpuMetricsMetadata

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**From** | Pointer to **time.Time** |  | [optional] 
**To** | Pointer to **time.Time** |  | [optional] 
**Granularity** | Pointer to [**MetricsGranularity**](MetricsGranularity.md) |  | [optional] 
**Summary** | Pointer to [**CpuMetricsMetadataSummary**](CpuMetricsMetadataSummary.md) |  | [optional] 

## Methods

### NewCpuMetricsMetadata

`func NewCpuMetricsMetadata() *CpuMetricsMetadata`

NewCpuMetricsMetadata instantiates a new CpuMetricsMetadata object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCpuMetricsMetadataWithDefaults

`func NewCpuMetricsMetadataWithDefaults() *CpuMetricsMetadata`

NewCpuMetricsMetadataWithDefaults instantiates a new CpuMetricsMetadata object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetFrom

`func (o *CpuMetricsMetadata) GetFrom() time.Time`

GetFrom returns the From field if non-nil, zero value otherwise.

### GetFromOk

`func (o *CpuMetricsMetadata) GetFromOk() (*time.Time, bool)`

GetFromOk returns a tuple with the From field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFrom

`func (o *CpuMetricsMetadata) SetFrom(v time.Time)`

SetFrom sets From field to given value.

### HasFrom

`func (o *CpuMetricsMetadata) HasFrom() bool`

HasFrom returns a boolean if a field has been set.

### GetTo

`func (o *CpuMetricsMetadata) GetTo() time.Time`

GetTo returns the To field if non-nil, zero value otherwise.

### GetToOk

`func (o *CpuMetricsMetadata) GetToOk() (*time.Time, bool)`

GetToOk returns a tuple with the To field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTo

`func (o *CpuMetricsMetadata) SetTo(v time.Time)`

SetTo sets To field to given value.

### HasTo

`func (o *CpuMetricsMetadata) HasTo() bool`

HasTo returns a boolean if a field has been set.

### GetGranularity

`func (o *CpuMetricsMetadata) GetGranularity() MetricsGranularity`

GetGranularity returns the Granularity field if non-nil, zero value otherwise.

### GetGranularityOk

`func (o *CpuMetricsMetadata) GetGranularityOk() (*MetricsGranularity, bool)`

GetGranularityOk returns a tuple with the Granularity field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGranularity

`func (o *CpuMetricsMetadata) SetGranularity(v MetricsGranularity)`

SetGranularity sets Granularity field to given value.

### HasGranularity

`func (o *CpuMetricsMetadata) HasGranularity() bool`

HasGranularity returns a boolean if a field has been set.

### GetSummary

`func (o *CpuMetricsMetadata) GetSummary() CpuMetricsMetadataSummary`

GetSummary returns the Summary field if non-nil, zero value otherwise.

### GetSummaryOk

`func (o *CpuMetricsMetadata) GetSummaryOk() (*CpuMetricsMetadataSummary, bool)`

GetSummaryOk returns a tuple with the Summary field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSummary

`func (o *CpuMetricsMetadata) SetSummary(v CpuMetricsMetadataSummary)`

SetSummary sets Summary field to given value.

### HasSummary

`func (o *CpuMetricsMetadata) HasSummary() bool`

HasSummary returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


