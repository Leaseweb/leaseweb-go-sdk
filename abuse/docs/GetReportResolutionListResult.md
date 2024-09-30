# GetReportResolutionListResult

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Resolutions** | [**[]ResolutionList**](ResolutionList.md) | Possible resolutions to resolve this report with. | 
**IsMessageRequired** | **bool** | &#x60;true&#x60;, if any of the IP(s) related to this report are null routed | 

## Methods

### NewGetReportResolutionListResult

`func NewGetReportResolutionListResult(resolutions []ResolutionList, isMessageRequired bool, ) *GetReportResolutionListResult`

NewGetReportResolutionListResult instantiates a new GetReportResolutionListResult object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetReportResolutionListResultWithDefaults

`func NewGetReportResolutionListResultWithDefaults() *GetReportResolutionListResult`

NewGetReportResolutionListResultWithDefaults instantiates a new GetReportResolutionListResult object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetResolutions

`func (o *GetReportResolutionListResult) GetResolutions() []ResolutionList`

GetResolutions returns the Resolutions field if non-nil, zero value otherwise.

### GetResolutionsOk

`func (o *GetReportResolutionListResult) GetResolutionsOk() (*[]ResolutionList, bool)`

GetResolutionsOk returns a tuple with the Resolutions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResolutions

`func (o *GetReportResolutionListResult) SetResolutions(v []ResolutionList)`

SetResolutions sets Resolutions field to given value.


### GetIsMessageRequired

`func (o *GetReportResolutionListResult) GetIsMessageRequired() bool`

GetIsMessageRequired returns the IsMessageRequired field if non-nil, zero value otherwise.

### GetIsMessageRequiredOk

`func (o *GetReportResolutionListResult) GetIsMessageRequiredOk() (*bool, bool)`

GetIsMessageRequiredOk returns a tuple with the IsMessageRequired field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsMessageRequired

`func (o *GetReportResolutionListResult) SetIsMessageRequired(v bool)`

SetIsMessageRequired sets IsMessageRequired field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


