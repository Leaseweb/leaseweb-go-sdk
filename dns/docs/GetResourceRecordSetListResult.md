# GetResourceRecordSetListResult

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**InfoMessage** | Pointer to **string** | Optional additional information | [optional] 
**Links** | Pointer to [**Links**](Links.md) |  | [optional] 
**ResourceRecordSets** | Pointer to [**[]ResourceRecordSetDetails**](ResourceRecordSetDetails.md) | Array of resource record sets | [optional] 

## Methods

### NewGetResourceRecordSetListResult

`func NewGetResourceRecordSetListResult() *GetResourceRecordSetListResult`

NewGetResourceRecordSetListResult instantiates a new GetResourceRecordSetListResult object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetResourceRecordSetListResultWithDefaults

`func NewGetResourceRecordSetListResultWithDefaults() *GetResourceRecordSetListResult`

NewGetResourceRecordSetListResultWithDefaults instantiates a new GetResourceRecordSetListResult object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetInfoMessage

`func (o *GetResourceRecordSetListResult) GetInfoMessage() string`

GetInfoMessage returns the InfoMessage field if non-nil, zero value otherwise.

### GetInfoMessageOk

`func (o *GetResourceRecordSetListResult) GetInfoMessageOk() (*string, bool)`

GetInfoMessageOk returns a tuple with the InfoMessage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInfoMessage

`func (o *GetResourceRecordSetListResult) SetInfoMessage(v string)`

SetInfoMessage sets InfoMessage field to given value.

### HasInfoMessage

`func (o *GetResourceRecordSetListResult) HasInfoMessage() bool`

HasInfoMessage returns a boolean if a field has been set.

### GetLinks

`func (o *GetResourceRecordSetListResult) GetLinks() Links`

GetLinks returns the Links field if non-nil, zero value otherwise.

### GetLinksOk

`func (o *GetResourceRecordSetListResult) GetLinksOk() (*Links, bool)`

GetLinksOk returns a tuple with the Links field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLinks

`func (o *GetResourceRecordSetListResult) SetLinks(v Links)`

SetLinks sets Links field to given value.

### HasLinks

`func (o *GetResourceRecordSetListResult) HasLinks() bool`

HasLinks returns a boolean if a field has been set.

### GetResourceRecordSets

`func (o *GetResourceRecordSetListResult) GetResourceRecordSets() []ResourceRecordSetDetails`

GetResourceRecordSets returns the ResourceRecordSets field if non-nil, zero value otherwise.

### GetResourceRecordSetsOk

`func (o *GetResourceRecordSetListResult) GetResourceRecordSetsOk() (*[]ResourceRecordSetDetails, bool)`

GetResourceRecordSetsOk returns a tuple with the ResourceRecordSets field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResourceRecordSets

`func (o *GetResourceRecordSetListResult) SetResourceRecordSets(v []ResourceRecordSetDetails)`

SetResourceRecordSets sets ResourceRecordSets field to given value.

### HasResourceRecordSets

`func (o *GetResourceRecordSetListResult) HasResourceRecordSets() bool`

HasResourceRecordSets returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


