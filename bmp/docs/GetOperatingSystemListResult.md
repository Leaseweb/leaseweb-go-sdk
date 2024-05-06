# GetOperatingSystemListResult

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Metadata** | Pointer to [**Metadata**](Metadata.md) |  | [optional] 
**OperatingSystems** | Pointer to [**[]OperatingSystem**](OperatingSystem.md) | A list of operating systems | [optional] 

## Methods

### NewGetOperatingSystemListResult

`func NewGetOperatingSystemListResult() *GetOperatingSystemListResult`

NewGetOperatingSystemListResult instantiates a new GetOperatingSystemListResult object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetOperatingSystemListResultWithDefaults

`func NewGetOperatingSystemListResultWithDefaults() *GetOperatingSystemListResult`

NewGetOperatingSystemListResultWithDefaults instantiates a new GetOperatingSystemListResult object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetMetadata

`func (o *GetOperatingSystemListResult) GetMetadata() Metadata`

GetMetadata returns the Metadata field if non-nil, zero value otherwise.

### GetMetadataOk

`func (o *GetOperatingSystemListResult) GetMetadataOk() (*Metadata, bool)`

GetMetadataOk returns a tuple with the Metadata field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetadata

`func (o *GetOperatingSystemListResult) SetMetadata(v Metadata)`

SetMetadata sets Metadata field to given value.

### HasMetadata

`func (o *GetOperatingSystemListResult) HasMetadata() bool`

HasMetadata returns a boolean if a field has been set.

### GetOperatingSystems

`func (o *GetOperatingSystemListResult) GetOperatingSystems() []OperatingSystem`

GetOperatingSystems returns the OperatingSystems field if non-nil, zero value otherwise.

### GetOperatingSystemsOk

`func (o *GetOperatingSystemListResult) GetOperatingSystemsOk() (*[]OperatingSystem, bool)`

GetOperatingSystemsOk returns a tuple with the OperatingSystems field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOperatingSystems

`func (o *GetOperatingSystemListResult) SetOperatingSystems(v []OperatingSystem)`

SetOperatingSystems sets OperatingSystems field to given value.

### HasOperatingSystems

`func (o *GetOperatingSystemListResult) HasOperatingSystems() bool`

HasOperatingSystems returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


