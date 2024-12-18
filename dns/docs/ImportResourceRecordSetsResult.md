# ImportResourceRecordSetsResult

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**InfoMessage** | Pointer to **string** | Optional additional information | [optional] 
**ResourceRecordSets** | Pointer to [**[]ResourceRecordSetDetails**](ResourceRecordSetDetails.md) | Array of resource record sets | [optional] 
**Links** | Pointer to [**Links**](Links.md) |  | [optional] 

## Methods

### NewImportResourceRecordSetsResult

`func NewImportResourceRecordSetsResult() *ImportResourceRecordSetsResult`

NewImportResourceRecordSetsResult instantiates a new ImportResourceRecordSetsResult object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewImportResourceRecordSetsResultWithDefaults

`func NewImportResourceRecordSetsResultWithDefaults() *ImportResourceRecordSetsResult`

NewImportResourceRecordSetsResultWithDefaults instantiates a new ImportResourceRecordSetsResult object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetInfoMessage

`func (o *ImportResourceRecordSetsResult) GetInfoMessage() string`

GetInfoMessage returns the InfoMessage field if non-nil, zero value otherwise.

### GetInfoMessageOk

`func (o *ImportResourceRecordSetsResult) GetInfoMessageOk() (*string, bool)`

GetInfoMessageOk returns a tuple with the InfoMessage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInfoMessage

`func (o *ImportResourceRecordSetsResult) SetInfoMessage(v string)`

SetInfoMessage sets InfoMessage field to given value.

### HasInfoMessage

`func (o *ImportResourceRecordSetsResult) HasInfoMessage() bool`

HasInfoMessage returns a boolean if a field has been set.

### GetResourceRecordSets

`func (o *ImportResourceRecordSetsResult) GetResourceRecordSets() []ResourceRecordSetDetails`

GetResourceRecordSets returns the ResourceRecordSets field if non-nil, zero value otherwise.

### GetResourceRecordSetsOk

`func (o *ImportResourceRecordSetsResult) GetResourceRecordSetsOk() (*[]ResourceRecordSetDetails, bool)`

GetResourceRecordSetsOk returns a tuple with the ResourceRecordSets field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResourceRecordSets

`func (o *ImportResourceRecordSetsResult) SetResourceRecordSets(v []ResourceRecordSetDetails)`

SetResourceRecordSets sets ResourceRecordSets field to given value.

### HasResourceRecordSets

`func (o *ImportResourceRecordSetsResult) HasResourceRecordSets() bool`

HasResourceRecordSets returns a boolean if a field has been set.

### GetLinks

`func (o *ImportResourceRecordSetsResult) GetLinks() Links`

GetLinks returns the Links field if non-nil, zero value otherwise.

### GetLinksOk

`func (o *ImportResourceRecordSetsResult) GetLinksOk() (*Links, bool)`

GetLinksOk returns a tuple with the Links field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLinks

`func (o *ImportResourceRecordSetsResult) SetLinks(v Links)`

SetLinks sets Links field to given value.

### HasLinks

`func (o *ImportResourceRecordSetsResult) HasLinks() bool`

HasLinks returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


