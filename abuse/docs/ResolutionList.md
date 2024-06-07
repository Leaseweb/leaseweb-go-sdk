# ResolutionList

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **string** | The resolution ID to be used when calling /resolve. | [optional] 
**Description** | Pointer to **string** | The actual text the report will be resolved with when resolving with this resolution. | [optional] 

## Methods

### NewResolutionList

`func NewResolutionList() *ResolutionList`

NewResolutionList instantiates a new ResolutionList object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewResolutionListWithDefaults

`func NewResolutionListWithDefaults() *ResolutionList`

NewResolutionListWithDefaults instantiates a new ResolutionList object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *ResolutionList) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *ResolutionList) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *ResolutionList) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *ResolutionList) HasId() bool`

HasId returns a boolean if a field has been set.

### GetDescription

`func (o *ResolutionList) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *ResolutionList) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *ResolutionList) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *ResolutionList) HasDescription() bool`

HasDescription returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


