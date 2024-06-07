# GetReportListResult

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Reports** | Pointer to [**[]Report**](Report.md) | An array of abuse reports. | [optional] 
**Metadata** | Pointer to [**Metadata**](Metadata.md) |  | [optional] 

## Methods

### NewGetReportListResult

`func NewGetReportListResult() *GetReportListResult`

NewGetReportListResult instantiates a new GetReportListResult object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetReportListResultWithDefaults

`func NewGetReportListResultWithDefaults() *GetReportListResult`

NewGetReportListResultWithDefaults instantiates a new GetReportListResult object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetReports

`func (o *GetReportListResult) GetReports() []Report`

GetReports returns the Reports field if non-nil, zero value otherwise.

### GetReportsOk

`func (o *GetReportListResult) GetReportsOk() (*[]Report, bool)`

GetReportsOk returns a tuple with the Reports field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReports

`func (o *GetReportListResult) SetReports(v []Report)`

SetReports sets Reports field to given value.

### HasReports

`func (o *GetReportListResult) HasReports() bool`

HasReports returns a boolean if a field has been set.

### GetMetadata

`func (o *GetReportListResult) GetMetadata() Metadata`

GetMetadata returns the Metadata field if non-nil, zero value otherwise.

### GetMetadataOk

`func (o *GetReportListResult) GetMetadataOk() (*Metadata, bool)`

GetMetadataOk returns a tuple with the Metadata field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetadata

`func (o *GetReportListResult) SetMetadata(v Metadata)`

SetMetadata sets Metadata field to given value.

### HasMetadata

`func (o *GetReportListResult) HasMetadata() bool`

HasMetadata returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


