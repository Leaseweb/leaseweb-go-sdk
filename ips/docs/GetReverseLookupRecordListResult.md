# GetReverseLookupRecordListResult

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ReverseLookups** | [**[]ReverseLookup**](ReverseLookup.md) |  | 
**Metadata** | [**Metadata**](Metadata.md) |  | 

## Methods

### NewGetReverseLookupRecordListResult

`func NewGetReverseLookupRecordListResult(reverseLookups []ReverseLookup, metadata Metadata, ) *GetReverseLookupRecordListResult`

NewGetReverseLookupRecordListResult instantiates a new GetReverseLookupRecordListResult object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetReverseLookupRecordListResultWithDefaults

`func NewGetReverseLookupRecordListResultWithDefaults() *GetReverseLookupRecordListResult`

NewGetReverseLookupRecordListResultWithDefaults instantiates a new GetReverseLookupRecordListResult object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetReverseLookups

`func (o *GetReverseLookupRecordListResult) GetReverseLookups() []ReverseLookup`

GetReverseLookups returns the ReverseLookups field if non-nil, zero value otherwise.

### GetReverseLookupsOk

`func (o *GetReverseLookupRecordListResult) GetReverseLookupsOk() (*[]ReverseLookup, bool)`

GetReverseLookupsOk returns a tuple with the ReverseLookups field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReverseLookups

`func (o *GetReverseLookupRecordListResult) SetReverseLookups(v []ReverseLookup)`

SetReverseLookups sets ReverseLookups field to given value.


### GetMetadata

`func (o *GetReverseLookupRecordListResult) GetMetadata() Metadata`

GetMetadata returns the Metadata field if non-nil, zero value otherwise.

### GetMetadataOk

`func (o *GetReverseLookupRecordListResult) GetMetadataOk() (*Metadata, bool)`

GetMetadataOk returns a tuple with the Metadata field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetadata

`func (o *GetReverseLookupRecordListResult) SetMetadata(v Metadata)`

SetMetadata sets Metadata field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


