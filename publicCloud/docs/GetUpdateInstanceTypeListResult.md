# GetUpdateInstanceTypeListResult

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Metadata** | Pointer to [**Metadata**](Metadata.md) |  | [optional] 
**InstanceTypes** | Pointer to [**[]UpdateInstanceType**](UpdateInstanceType.md) |  | [optional] 

## Methods

### NewGetUpdateInstanceTypeListResult

`func NewGetUpdateInstanceTypeListResult() *GetUpdateInstanceTypeListResult`

NewGetUpdateInstanceTypeListResult instantiates a new GetUpdateInstanceTypeListResult object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetUpdateInstanceTypeListResultWithDefaults

`func NewGetUpdateInstanceTypeListResultWithDefaults() *GetUpdateInstanceTypeListResult`

NewGetUpdateInstanceTypeListResultWithDefaults instantiates a new GetUpdateInstanceTypeListResult object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetMetadata

`func (o *GetUpdateInstanceTypeListResult) GetMetadata() Metadata`

GetMetadata returns the Metadata field if non-nil, zero value otherwise.

### GetMetadataOk

`func (o *GetUpdateInstanceTypeListResult) GetMetadataOk() (*Metadata, bool)`

GetMetadataOk returns a tuple with the Metadata field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetadata

`func (o *GetUpdateInstanceTypeListResult) SetMetadata(v Metadata)`

SetMetadata sets Metadata field to given value.

### HasMetadata

`func (o *GetUpdateInstanceTypeListResult) HasMetadata() bool`

HasMetadata returns a boolean if a field has been set.

### GetInstanceTypes

`func (o *GetUpdateInstanceTypeListResult) GetInstanceTypes() []UpdateInstanceType`

GetInstanceTypes returns the InstanceTypes field if non-nil, zero value otherwise.

### GetInstanceTypesOk

`func (o *GetUpdateInstanceTypeListResult) GetInstanceTypesOk() (*[]UpdateInstanceType, bool)`

GetInstanceTypesOk returns a tuple with the InstanceTypes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInstanceTypes

`func (o *GetUpdateInstanceTypeListResult) SetInstanceTypes(v []UpdateInstanceType)`

SetInstanceTypes sets InstanceTypes field to given value.

### HasInstanceTypes

`func (o *GetUpdateInstanceTypeListResult) HasInstanceTypes() bool`

HasInstanceTypes returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


