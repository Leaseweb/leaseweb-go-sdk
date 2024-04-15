# GetReportMessageListResult

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Messages** | Pointer to [**[]Message**](Message.md) | An array of the posted messages. | [optional] 
**Metadata** | Pointer to [**Metadata**](Metadata.md) |  | [optional] 

## Methods

### NewGetReportMessageListResult

`func NewGetReportMessageListResult() *GetReportMessageListResult`

NewGetReportMessageListResult instantiates a new GetReportMessageListResult object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetReportMessageListResultWithDefaults

`func NewGetReportMessageListResultWithDefaults() *GetReportMessageListResult`

NewGetReportMessageListResultWithDefaults instantiates a new GetReportMessageListResult object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetMessages

`func (o *GetReportMessageListResult) GetMessages() []Message`

GetMessages returns the Messages field if non-nil, zero value otherwise.

### GetMessagesOk

`func (o *GetReportMessageListResult) GetMessagesOk() (*[]Message, bool)`

GetMessagesOk returns a tuple with the Messages field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMessages

`func (o *GetReportMessageListResult) SetMessages(v []Message)`

SetMessages sets Messages field to given value.

### HasMessages

`func (o *GetReportMessageListResult) HasMessages() bool`

HasMessages returns a boolean if a field has been set.

### GetMetadata

`func (o *GetReportMessageListResult) GetMetadata() Metadata`

GetMetadata returns the Metadata field if non-nil, zero value otherwise.

### GetMetadataOk

`func (o *GetReportMessageListResult) GetMetadataOk() (*Metadata, bool)`

GetMetadataOk returns a tuple with the Metadata field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetadata

`func (o *GetReportMessageListResult) SetMetadata(v Metadata)`

SetMetadata sets Metadata field to given value.

### HasMetadata

`func (o *GetReportMessageListResult) HasMetadata() bool`

HasMetadata returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


