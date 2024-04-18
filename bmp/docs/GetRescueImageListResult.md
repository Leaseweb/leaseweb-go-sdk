# GetRescueImageListResult

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Metadata** | Pointer to [**Metadata**](Metadata.md) |  | [optional] 
**RescueImages** | Pointer to [**[]RescueImage**](RescueImage.md) | A list of operating systems | [optional] 

## Methods

### NewGetRescueImageListResult

`func NewGetRescueImageListResult() *GetRescueImageListResult`

NewGetRescueImageListResult instantiates a new GetRescueImageListResult object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetRescueImageListResultWithDefaults

`func NewGetRescueImageListResultWithDefaults() *GetRescueImageListResult`

NewGetRescueImageListResultWithDefaults instantiates a new GetRescueImageListResult object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetMetadata

`func (o *GetRescueImageListResult) GetMetadata() Metadata`

GetMetadata returns the Metadata field if non-nil, zero value otherwise.

### GetMetadataOk

`func (o *GetRescueImageListResult) GetMetadataOk() (*Metadata, bool)`

GetMetadataOk returns a tuple with the Metadata field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetadata

`func (o *GetRescueImageListResult) SetMetadata(v Metadata)`

SetMetadata sets Metadata field to given value.

### HasMetadata

`func (o *GetRescueImageListResult) HasMetadata() bool`

HasMetadata returns a boolean if a field has been set.

### GetRescueImages

`func (o *GetRescueImageListResult) GetRescueImages() []RescueImage`

GetRescueImages returns the RescueImages field if non-nil, zero value otherwise.

### GetRescueImagesOk

`func (o *GetRescueImageListResult) GetRescueImagesOk() (*[]RescueImage, bool)`

GetRescueImagesOk returns a tuple with the RescueImages field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRescueImages

`func (o *GetRescueImageListResult) SetRescueImages(v []RescueImage)`

SetRescueImages sets RescueImages field to given value.

### HasRescueImages

`func (o *GetRescueImageListResult) HasRescueImages() bool`

HasRescueImages returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


