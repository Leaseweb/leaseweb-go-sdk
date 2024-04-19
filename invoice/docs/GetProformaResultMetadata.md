# GetProformaResultMetadata

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Limit** | Pointer to **int32** | Limit the number of results returned | [optional] [default to 10]
**Offset** | Pointer to **int32** | Return results starting from the given offset | [optional] [default to 0]
**TotalCount** | Pointer to **int32** | Total amount of contract items. | [optional] 

## Methods

### NewGetProformaResultMetadata

`func NewGetProformaResultMetadata() *GetProformaResultMetadata`

NewGetProformaResultMetadata instantiates a new GetProformaResultMetadata object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetProformaResultMetadataWithDefaults

`func NewGetProformaResultMetadataWithDefaults() *GetProformaResultMetadata`

NewGetProformaResultMetadataWithDefaults instantiates a new GetProformaResultMetadata object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetLimit

`func (o *GetProformaResultMetadata) GetLimit() int32`

GetLimit returns the Limit field if non-nil, zero value otherwise.

### GetLimitOk

`func (o *GetProformaResultMetadata) GetLimitOk() (*int32, bool)`

GetLimitOk returns a tuple with the Limit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLimit

`func (o *GetProformaResultMetadata) SetLimit(v int32)`

SetLimit sets Limit field to given value.

### HasLimit

`func (o *GetProformaResultMetadata) HasLimit() bool`

HasLimit returns a boolean if a field has been set.

### GetOffset

`func (o *GetProformaResultMetadata) GetOffset() int32`

GetOffset returns the Offset field if non-nil, zero value otherwise.

### GetOffsetOk

`func (o *GetProformaResultMetadata) GetOffsetOk() (*int32, bool)`

GetOffsetOk returns a tuple with the Offset field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOffset

`func (o *GetProformaResultMetadata) SetOffset(v int32)`

SetOffset sets Offset field to given value.

### HasOffset

`func (o *GetProformaResultMetadata) HasOffset() bool`

HasOffset returns a boolean if a field has been set.

### GetTotalCount

`func (o *GetProformaResultMetadata) GetTotalCount() int32`

GetTotalCount returns the TotalCount field if non-nil, zero value otherwise.

### GetTotalCountOk

`func (o *GetProformaResultMetadata) GetTotalCountOk() (*int32, bool)`

GetTotalCountOk returns a tuple with the TotalCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalCount

`func (o *GetProformaResultMetadata) SetTotalCount(v int32)`

SetTotalCount sets TotalCount field to given value.

### HasTotalCount

`func (o *GetProformaResultMetadata) HasTotalCount() bool`

HasTotalCount returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


