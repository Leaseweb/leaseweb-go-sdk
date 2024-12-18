# BaseResourceRecordSetList

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**InfoMessage** | Pointer to **string** | Optional additional information | [optional] 
**Links** | Pointer to [**Links**](Links.md) |  | [optional] 

## Methods

### NewBaseResourceRecordSetList

`func NewBaseResourceRecordSetList() *BaseResourceRecordSetList`

NewBaseResourceRecordSetList instantiates a new BaseResourceRecordSetList object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewBaseResourceRecordSetListWithDefaults

`func NewBaseResourceRecordSetListWithDefaults() *BaseResourceRecordSetList`

NewBaseResourceRecordSetListWithDefaults instantiates a new BaseResourceRecordSetList object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetInfoMessage

`func (o *BaseResourceRecordSetList) GetInfoMessage() string`

GetInfoMessage returns the InfoMessage field if non-nil, zero value otherwise.

### GetInfoMessageOk

`func (o *BaseResourceRecordSetList) GetInfoMessageOk() (*string, bool)`

GetInfoMessageOk returns a tuple with the InfoMessage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInfoMessage

`func (o *BaseResourceRecordSetList) SetInfoMessage(v string)`

SetInfoMessage sets InfoMessage field to given value.

### HasInfoMessage

`func (o *BaseResourceRecordSetList) HasInfoMessage() bool`

HasInfoMessage returns a boolean if a field has been set.

### GetLinks

`func (o *BaseResourceRecordSetList) GetLinks() Links`

GetLinks returns the Links field if non-nil, zero value otherwise.

### GetLinksOk

`func (o *BaseResourceRecordSetList) GetLinksOk() (*Links, bool)`

GetLinksOk returns a tuple with the Links field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLinks

`func (o *BaseResourceRecordSetList) SetLinks(v Links)`

SetLinks sets Links field to given value.

### HasLinks

`func (o *BaseResourceRecordSetList) HasLinks() bool`

HasLinks returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


