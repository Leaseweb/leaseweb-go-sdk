# GetResourceRecordSetListLdResult

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**InfoMessage** | Pointer to **string** | Optional additional information | [optional] 
**Links** | Pointer to [**Links**](Links.md) |  | [optional] 
**ResourceRecordSets** | Pointer to [**[]LdResourceRecordSetDetails**](LdResourceRecordSetDetails.md) | Array of resource record sets | [optional] 

## Methods

### NewGetResourceRecordSetListLdResult

`func NewGetResourceRecordSetListLdResult() *GetResourceRecordSetListLdResult`

NewGetResourceRecordSetListLdResult instantiates a new GetResourceRecordSetListLdResult object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetResourceRecordSetListLdResultWithDefaults

`func NewGetResourceRecordSetListLdResultWithDefaults() *GetResourceRecordSetListLdResult`

NewGetResourceRecordSetListLdResultWithDefaults instantiates a new GetResourceRecordSetListLdResult object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetInfoMessage

`func (o *GetResourceRecordSetListLdResult) GetInfoMessage() string`

GetInfoMessage returns the InfoMessage field if non-nil, zero value otherwise.

### GetInfoMessageOk

`func (o *GetResourceRecordSetListLdResult) GetInfoMessageOk() (*string, bool)`

GetInfoMessageOk returns a tuple with the InfoMessage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInfoMessage

`func (o *GetResourceRecordSetListLdResult) SetInfoMessage(v string)`

SetInfoMessage sets InfoMessage field to given value.

### HasInfoMessage

`func (o *GetResourceRecordSetListLdResult) HasInfoMessage() bool`

HasInfoMessage returns a boolean if a field has been set.

### GetLinks

`func (o *GetResourceRecordSetListLdResult) GetLinks() Links`

GetLinks returns the Links field if non-nil, zero value otherwise.

### GetLinksOk

`func (o *GetResourceRecordSetListLdResult) GetLinksOk() (*Links, bool)`

GetLinksOk returns a tuple with the Links field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLinks

`func (o *GetResourceRecordSetListLdResult) SetLinks(v Links)`

SetLinks sets Links field to given value.

### HasLinks

`func (o *GetResourceRecordSetListLdResult) HasLinks() bool`

HasLinks returns a boolean if a field has been set.

### GetResourceRecordSets

`func (o *GetResourceRecordSetListLdResult) GetResourceRecordSets() []LdResourceRecordSetDetails`

GetResourceRecordSets returns the ResourceRecordSets field if non-nil, zero value otherwise.

### GetResourceRecordSetsOk

`func (o *GetResourceRecordSetListLdResult) GetResourceRecordSetsOk() (*[]LdResourceRecordSetDetails, bool)`

GetResourceRecordSetsOk returns a tuple with the ResourceRecordSets field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResourceRecordSets

`func (o *GetResourceRecordSetListLdResult) SetResourceRecordSets(v []LdResourceRecordSetDetails)`

SetResourceRecordSets sets ResourceRecordSets field to given value.

### HasResourceRecordSets

`func (o *GetResourceRecordSetListLdResult) HasResourceRecordSets() bool`

HasResourceRecordSets returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


