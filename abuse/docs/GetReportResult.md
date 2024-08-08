# GetReportResult

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **string** | The ID of the abuse report. | [optional] 
**Subject** | Pointer to **string** | The subject of the report. | [optional] 
**Status** | Pointer to **string** | The current status of the report. | [optional] 
**AbuseType** | Pointer to **string** | The abuse type of the report. | [optional] 
**Reopened** | Pointer to **bool** | When the report status is open, this indicates if the report was reopened. | [optional] 
**ReportedAt** | Pointer to **string** | The date and time the report was reported at. | [optional] 
**UpdatedAt** | Pointer to **string** | When the report was updated last. | [optional] 
**Notifier** | Pointer to **string** | The email address of the notifier who reported the abuse. | [optional] 
**CustomerId** | Pointer to **string** | The customer ID of your account. | [optional] 
**LegalEntityId** | Pointer to **string** | The legal entity ID of the customer account. | [optional] 
**Body** | Pointer to **string** | The report body content. | [optional] 
**Deadline** | Pointer to **string** | The Deadline before when the report needs to be resolved. | [optional] 
**DetectedIpAddresses** | Pointer to **[]string** | The IP-addresses detected in the report body. | [optional] 
**DetectedDomainNames** | Pointer to [**[]DetectedDomainNameList**](DetectedDomainNameList.md) | The domains detected in the report body. | [optional] 
**Attachments** | Pointer to [**[]Attachment**](Attachment.md) | All the files attached to the report. | [optional] 
**TotalMessagesCount** | Pointer to **int32** | Total amount of messages in the report. | [optional] 
**LatestMessages** | Pointer to [**[]Message**](Message.md) | Array of the last 5 messages in the report. | [optional] 

## Methods

### NewGetReportResult

`func NewGetReportResult() *GetReportResult`

NewGetReportResult instantiates a new GetReportResult object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetReportResultWithDefaults

`func NewGetReportResultWithDefaults() *GetReportResult`

NewGetReportResultWithDefaults instantiates a new GetReportResult object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *GetReportResult) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *GetReportResult) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *GetReportResult) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *GetReportResult) HasId() bool`

HasId returns a boolean if a field has been set.

### GetSubject

`func (o *GetReportResult) GetSubject() string`

GetSubject returns the Subject field if non-nil, zero value otherwise.

### GetSubjectOk

`func (o *GetReportResult) GetSubjectOk() (*string, bool)`

GetSubjectOk returns a tuple with the Subject field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubject

`func (o *GetReportResult) SetSubject(v string)`

SetSubject sets Subject field to given value.

### HasSubject

`func (o *GetReportResult) HasSubject() bool`

HasSubject returns a boolean if a field has been set.

### GetStatus

`func (o *GetReportResult) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *GetReportResult) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *GetReportResult) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *GetReportResult) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetAbuseType

`func (o *GetReportResult) GetAbuseType() string`

GetAbuseType returns the AbuseType field if non-nil, zero value otherwise.

### GetAbuseTypeOk

`func (o *GetReportResult) GetAbuseTypeOk() (*string, bool)`

GetAbuseTypeOk returns a tuple with the AbuseType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAbuseType

`func (o *GetReportResult) SetAbuseType(v string)`

SetAbuseType sets AbuseType field to given value.

### HasAbuseType

`func (o *GetReportResult) HasAbuseType() bool`

HasAbuseType returns a boolean if a field has been set.

### GetReopened

`func (o *GetReportResult) GetReopened() bool`

GetReopened returns the Reopened field if non-nil, zero value otherwise.

### GetReopenedOk

`func (o *GetReportResult) GetReopenedOk() (*bool, bool)`

GetReopenedOk returns a tuple with the Reopened field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReopened

`func (o *GetReportResult) SetReopened(v bool)`

SetReopened sets Reopened field to given value.

### HasReopened

`func (o *GetReportResult) HasReopened() bool`

HasReopened returns a boolean if a field has been set.

### GetReportedAt

`func (o *GetReportResult) GetReportedAt() string`

GetReportedAt returns the ReportedAt field if non-nil, zero value otherwise.

### GetReportedAtOk

`func (o *GetReportResult) GetReportedAtOk() (*string, bool)`

GetReportedAtOk returns a tuple with the ReportedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReportedAt

`func (o *GetReportResult) SetReportedAt(v string)`

SetReportedAt sets ReportedAt field to given value.

### HasReportedAt

`func (o *GetReportResult) HasReportedAt() bool`

HasReportedAt returns a boolean if a field has been set.

### GetUpdatedAt

`func (o *GetReportResult) GetUpdatedAt() string`

GetUpdatedAt returns the UpdatedAt field if non-nil, zero value otherwise.

### GetUpdatedAtOk

`func (o *GetReportResult) GetUpdatedAtOk() (*string, bool)`

GetUpdatedAtOk returns a tuple with the UpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedAt

`func (o *GetReportResult) SetUpdatedAt(v string)`

SetUpdatedAt sets UpdatedAt field to given value.

### HasUpdatedAt

`func (o *GetReportResult) HasUpdatedAt() bool`

HasUpdatedAt returns a boolean if a field has been set.

### GetNotifier

`func (o *GetReportResult) GetNotifier() string`

GetNotifier returns the Notifier field if non-nil, zero value otherwise.

### GetNotifierOk

`func (o *GetReportResult) GetNotifierOk() (*string, bool)`

GetNotifierOk returns a tuple with the Notifier field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNotifier

`func (o *GetReportResult) SetNotifier(v string)`

SetNotifier sets Notifier field to given value.

### HasNotifier

`func (o *GetReportResult) HasNotifier() bool`

HasNotifier returns a boolean if a field has been set.

### GetCustomerId

`func (o *GetReportResult) GetCustomerId() string`

GetCustomerId returns the CustomerId field if non-nil, zero value otherwise.

### GetCustomerIdOk

`func (o *GetReportResult) GetCustomerIdOk() (*string, bool)`

GetCustomerIdOk returns a tuple with the CustomerId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomerId

`func (o *GetReportResult) SetCustomerId(v string)`

SetCustomerId sets CustomerId field to given value.

### HasCustomerId

`func (o *GetReportResult) HasCustomerId() bool`

HasCustomerId returns a boolean if a field has been set.

### GetLegalEntityId

`func (o *GetReportResult) GetLegalEntityId() string`

GetLegalEntityId returns the LegalEntityId field if non-nil, zero value otherwise.

### GetLegalEntityIdOk

`func (o *GetReportResult) GetLegalEntityIdOk() (*string, bool)`

GetLegalEntityIdOk returns a tuple with the LegalEntityId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLegalEntityId

`func (o *GetReportResult) SetLegalEntityId(v string)`

SetLegalEntityId sets LegalEntityId field to given value.

### HasLegalEntityId

`func (o *GetReportResult) HasLegalEntityId() bool`

HasLegalEntityId returns a boolean if a field has been set.

### GetBody

`func (o *GetReportResult) GetBody() string`

GetBody returns the Body field if non-nil, zero value otherwise.

### GetBodyOk

`func (o *GetReportResult) GetBodyOk() (*string, bool)`

GetBodyOk returns a tuple with the Body field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBody

`func (o *GetReportResult) SetBody(v string)`

SetBody sets Body field to given value.

### HasBody

`func (o *GetReportResult) HasBody() bool`

HasBody returns a boolean if a field has been set.

### GetDeadline

`func (o *GetReportResult) GetDeadline() string`

GetDeadline returns the Deadline field if non-nil, zero value otherwise.

### GetDeadlineOk

`func (o *GetReportResult) GetDeadlineOk() (*string, bool)`

GetDeadlineOk returns a tuple with the Deadline field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeadline

`func (o *GetReportResult) SetDeadline(v string)`

SetDeadline sets Deadline field to given value.

### HasDeadline

`func (o *GetReportResult) HasDeadline() bool`

HasDeadline returns a boolean if a field has been set.

### GetDetectedIpAddresses

`func (o *GetReportResult) GetDetectedIpAddresses() []string`

GetDetectedIpAddresses returns the DetectedIpAddresses field if non-nil, zero value otherwise.

### GetDetectedIpAddressesOk

`func (o *GetReportResult) GetDetectedIpAddressesOk() (*[]string, bool)`

GetDetectedIpAddressesOk returns a tuple with the DetectedIpAddresses field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDetectedIpAddresses

`func (o *GetReportResult) SetDetectedIpAddresses(v []string)`

SetDetectedIpAddresses sets DetectedIpAddresses field to given value.

### HasDetectedIpAddresses

`func (o *GetReportResult) HasDetectedIpAddresses() bool`

HasDetectedIpAddresses returns a boolean if a field has been set.

### GetDetectedDomainNames

`func (o *GetReportResult) GetDetectedDomainNames() []DetectedDomainNameList`

GetDetectedDomainNames returns the DetectedDomainNames field if non-nil, zero value otherwise.

### GetDetectedDomainNamesOk

`func (o *GetReportResult) GetDetectedDomainNamesOk() (*[]DetectedDomainNameList, bool)`

GetDetectedDomainNamesOk returns a tuple with the DetectedDomainNames field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDetectedDomainNames

`func (o *GetReportResult) SetDetectedDomainNames(v []DetectedDomainNameList)`

SetDetectedDomainNames sets DetectedDomainNames field to given value.

### HasDetectedDomainNames

`func (o *GetReportResult) HasDetectedDomainNames() bool`

HasDetectedDomainNames returns a boolean if a field has been set.

### GetAttachments

`func (o *GetReportResult) GetAttachments() []Attachment`

GetAttachments returns the Attachments field if non-nil, zero value otherwise.

### GetAttachmentsOk

`func (o *GetReportResult) GetAttachmentsOk() (*[]Attachment, bool)`

GetAttachmentsOk returns a tuple with the Attachments field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAttachments

`func (o *GetReportResult) SetAttachments(v []Attachment)`

SetAttachments sets Attachments field to given value.

### HasAttachments

`func (o *GetReportResult) HasAttachments() bool`

HasAttachments returns a boolean if a field has been set.

### GetTotalMessagesCount

`func (o *GetReportResult) GetTotalMessagesCount() int32`

GetTotalMessagesCount returns the TotalMessagesCount field if non-nil, zero value otherwise.

### GetTotalMessagesCountOk

`func (o *GetReportResult) GetTotalMessagesCountOk() (*int32, bool)`

GetTotalMessagesCountOk returns a tuple with the TotalMessagesCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalMessagesCount

`func (o *GetReportResult) SetTotalMessagesCount(v int32)`

SetTotalMessagesCount sets TotalMessagesCount field to given value.

### HasTotalMessagesCount

`func (o *GetReportResult) HasTotalMessagesCount() bool`

HasTotalMessagesCount returns a boolean if a field has been set.

### GetLatestMessages

`func (o *GetReportResult) GetLatestMessages() []Message`

GetLatestMessages returns the LatestMessages field if non-nil, zero value otherwise.

### GetLatestMessagesOk

`func (o *GetReportResult) GetLatestMessagesOk() (*[]Message, bool)`

GetLatestMessagesOk returns a tuple with the LatestMessages field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLatestMessages

`func (o *GetReportResult) SetLatestMessages(v []Message)`

SetLatestMessages sets LatestMessages field to given value.

### HasLatestMessages

`func (o *GetReportResult) HasLatestMessages() bool`

HasLatestMessages returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


