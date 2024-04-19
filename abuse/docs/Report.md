# Report

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **string** | The ID of the report. | [optional] 
**Subject** | Pointer to **string** | The subject of the  report. | [optional] 
**Status** | Pointer to **string** | The current status of the report. | [optional] 
**ReportedAt** | Pointer to **string** | The date and time the report was reported at. | [optional] 
**UpdatedAt** | Pointer to **string** | The date and time the report was last updated. | [optional] 
**Notifier** | Pointer to **string** | Due to compliance, &#39;REDACTED_FOR_PRIVACY&#39; is the default value. | [optional] [default to "REDACTED_FOR_PRIVACY"]
**CustomerId** | Pointer to **string** | The customer ID of your account. | [optional] 
**LegalEntityId** | Pointer to **string** | The legal entity ID of the customer account. | [optional] 
**Deadline** | Pointer to **string** | The deadline before when the report needs to be resolved. | [optional] 

## Methods

### NewReport

`func NewReport() *Report`

NewReport instantiates a new Report object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewReportWithDefaults

`func NewReportWithDefaults() *Report`

NewReportWithDefaults instantiates a new Report object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *Report) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *Report) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *Report) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *Report) HasId() bool`

HasId returns a boolean if a field has been set.

### GetSubject

`func (o *Report) GetSubject() string`

GetSubject returns the Subject field if non-nil, zero value otherwise.

### GetSubjectOk

`func (o *Report) GetSubjectOk() (*string, bool)`

GetSubjectOk returns a tuple with the Subject field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubject

`func (o *Report) SetSubject(v string)`

SetSubject sets Subject field to given value.

### HasSubject

`func (o *Report) HasSubject() bool`

HasSubject returns a boolean if a field has been set.

### GetStatus

`func (o *Report) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *Report) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *Report) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *Report) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetReportedAt

`func (o *Report) GetReportedAt() string`

GetReportedAt returns the ReportedAt field if non-nil, zero value otherwise.

### GetReportedAtOk

`func (o *Report) GetReportedAtOk() (*string, bool)`

GetReportedAtOk returns a tuple with the ReportedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReportedAt

`func (o *Report) SetReportedAt(v string)`

SetReportedAt sets ReportedAt field to given value.

### HasReportedAt

`func (o *Report) HasReportedAt() bool`

HasReportedAt returns a boolean if a field has been set.

### GetUpdatedAt

`func (o *Report) GetUpdatedAt() string`

GetUpdatedAt returns the UpdatedAt field if non-nil, zero value otherwise.

### GetUpdatedAtOk

`func (o *Report) GetUpdatedAtOk() (*string, bool)`

GetUpdatedAtOk returns a tuple with the UpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedAt

`func (o *Report) SetUpdatedAt(v string)`

SetUpdatedAt sets UpdatedAt field to given value.

### HasUpdatedAt

`func (o *Report) HasUpdatedAt() bool`

HasUpdatedAt returns a boolean if a field has been set.

### GetNotifier

`func (o *Report) GetNotifier() string`

GetNotifier returns the Notifier field if non-nil, zero value otherwise.

### GetNotifierOk

`func (o *Report) GetNotifierOk() (*string, bool)`

GetNotifierOk returns a tuple with the Notifier field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNotifier

`func (o *Report) SetNotifier(v string)`

SetNotifier sets Notifier field to given value.

### HasNotifier

`func (o *Report) HasNotifier() bool`

HasNotifier returns a boolean if a field has been set.

### GetCustomerId

`func (o *Report) GetCustomerId() string`

GetCustomerId returns the CustomerId field if non-nil, zero value otherwise.

### GetCustomerIdOk

`func (o *Report) GetCustomerIdOk() (*string, bool)`

GetCustomerIdOk returns a tuple with the CustomerId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomerId

`func (o *Report) SetCustomerId(v string)`

SetCustomerId sets CustomerId field to given value.

### HasCustomerId

`func (o *Report) HasCustomerId() bool`

HasCustomerId returns a boolean if a field has been set.

### GetLegalEntityId

`func (o *Report) GetLegalEntityId() string`

GetLegalEntityId returns the LegalEntityId field if non-nil, zero value otherwise.

### GetLegalEntityIdOk

`func (o *Report) GetLegalEntityIdOk() (*string, bool)`

GetLegalEntityIdOk returns a tuple with the LegalEntityId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLegalEntityId

`func (o *Report) SetLegalEntityId(v string)`

SetLegalEntityId sets LegalEntityId field to given value.

### HasLegalEntityId

`func (o *Report) HasLegalEntityId() bool`

HasLegalEntityId returns a boolean if a field has been set.

### GetDeadline

`func (o *Report) GetDeadline() string`

GetDeadline returns the Deadline field if non-nil, zero value otherwise.

### GetDeadlineOk

`func (o *Report) GetDeadlineOk() (*string, bool)`

GetDeadlineOk returns a tuple with the Deadline field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeadline

`func (o *Report) SetDeadline(v string)`

SetDeadline sets Deadline field to given value.

### HasDeadline

`func (o *Report) HasDeadline() bool`

HasDeadline returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


