# GetIsoListResult

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Isos** | Pointer to [**[]Iso**](Iso.md) |  | [optional] 
**Metadata** | Pointer to [**Metadata**](Metadata.md) |  | [optional] 

## Methods

### NewGetIsoListResult

`func NewGetIsoListResult() *GetIsoListResult`

NewGetIsoListResult instantiates a new GetIsoListResult object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetIsoListResultWithDefaults

`func NewGetIsoListResultWithDefaults() *GetIsoListResult`

NewGetIsoListResultWithDefaults instantiates a new GetIsoListResult object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetIsos

`func (o *GetIsoListResult) GetIsos() []Iso`

GetIsos returns the Isos field if non-nil, zero value otherwise.

### GetIsosOk

`func (o *GetIsoListResult) GetIsosOk() (*[]Iso, bool)`

GetIsosOk returns a tuple with the Isos field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsos

`func (o *GetIsoListResult) SetIsos(v []Iso)`

SetIsos sets Isos field to given value.

### HasIsos

`func (o *GetIsoListResult) HasIsos() bool`

HasIsos returns a boolean if a field has been set.

### GetMetadata

`func (o *GetIsoListResult) GetMetadata() Metadata`

GetMetadata returns the Metadata field if non-nil, zero value otherwise.

### GetMetadataOk

`func (o *GetIsoListResult) GetMetadataOk() (*Metadata, bool)`

GetMetadataOk returns a tuple with the Metadata field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetadata

`func (o *GetIsoListResult) SetMetadata(v Metadata)`

SetMetadata sets Metadata field to given value.

### HasMetadata

`func (o *GetIsoListResult) HasMetadata() bool`

HasMetadata returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


