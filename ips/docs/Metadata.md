# Metadata

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**TotalCount** | **int32** | Total amount of elements in this collection | 
**Offset** | **int32** | The offset used to generate this response | [default to 0]
**Limit** | **int32** | The limit used to generate this response | [default to 5]

## Methods

### NewMetadata

`func NewMetadata(totalCount int32, offset int32, limit int32, ) *Metadata`

NewMetadata instantiates a new Metadata object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMetadataWithDefaults

`func NewMetadataWithDefaults() *Metadata`

NewMetadataWithDefaults instantiates a new Metadata object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetTotalCount

`func (o *Metadata) GetTotalCount() int32`

GetTotalCount returns the TotalCount field if non-nil, zero value otherwise.

### GetTotalCountOk

`func (o *Metadata) GetTotalCountOk() (*int32, bool)`

GetTotalCountOk returns a tuple with the TotalCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalCount

`func (o *Metadata) SetTotalCount(v int32)`

SetTotalCount sets TotalCount field to given value.


### GetOffset

`func (o *Metadata) GetOffset() int32`

GetOffset returns the Offset field if non-nil, zero value otherwise.

### GetOffsetOk

`func (o *Metadata) GetOffsetOk() (*int32, bool)`

GetOffsetOk returns a tuple with the Offset field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOffset

`func (o *Metadata) SetOffset(v int32)`

SetOffset sets Offset field to given value.


### GetLimit

`func (o *Metadata) GetLimit() int32`

GetLimit returns the Limit field if non-nil, zero value otherwise.

### GetLimitOk

`func (o *Metadata) GetLimitOk() (*int32, bool)`

GetLimitOk returns a tuple with the Limit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLimit

`func (o *Metadata) SetLimit(v int32)`

SetLimit sets Limit field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


