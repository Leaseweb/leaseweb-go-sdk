# ResolveReportResult

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Resolutions** | **[]string** | List of selected resolution ID&#39;s to explain how the report is resolved. | 
**Message** | Pointer to **string** | Message is required and only allowed if any of the IP(s) related to this report are null routed. | [optional] 

## Methods

### NewResolveReportResult

`func NewResolveReportResult(resolutions []string, ) *ResolveReportResult`

NewResolveReportResult instantiates a new ResolveReportResult object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewResolveReportResultWithDefaults

`func NewResolveReportResultWithDefaults() *ResolveReportResult`

NewResolveReportResultWithDefaults instantiates a new ResolveReportResult object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetResolutions

`func (o *ResolveReportResult) GetResolutions() []string`

GetResolutions returns the Resolutions field if non-nil, zero value otherwise.

### GetResolutionsOk

`func (o *ResolveReportResult) GetResolutionsOk() (*[]string, bool)`

GetResolutionsOk returns a tuple with the Resolutions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResolutions

`func (o *ResolveReportResult) SetResolutions(v []string)`

SetResolutions sets Resolutions field to given value.


### GetMessage

`func (o *ResolveReportResult) GetMessage() string`

GetMessage returns the Message field if non-nil, zero value otherwise.

### GetMessageOk

`func (o *ResolveReportResult) GetMessageOk() (*string, bool)`

GetMessageOk returns a tuple with the Message field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMessage

`func (o *ResolveReportResult) SetMessage(v string)`

SetMessage sets Message field to given value.

### HasMessage

`func (o *ResolveReportResult) HasMessage() bool`

HasMessage returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


