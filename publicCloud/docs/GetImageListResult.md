# GetImageListResult

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Images** | Pointer to [**[]ImageDetails**](ImageDetails.md) |  | [optional] 
**Metadata** | Pointer to [**Metadata**](Metadata.md) |  | [optional] 

## Methods

### NewGetImageListResult

`func NewGetImageListResult() *GetImageListResult`

NewGetImageListResult instantiates a new GetImageListResult object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetImageListResultWithDefaults

`func NewGetImageListResultWithDefaults() *GetImageListResult`

NewGetImageListResultWithDefaults instantiates a new GetImageListResult object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetImages

`func (o *GetImageListResult) GetImages() []ImageDetails`

GetImages returns the Images field if non-nil, zero value otherwise.

### GetImagesOk

`func (o *GetImageListResult) GetImagesOk() (*[]ImageDetails, bool)`

GetImagesOk returns a tuple with the Images field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetImages

`func (o *GetImageListResult) SetImages(v []ImageDetails)`

SetImages sets Images field to given value.

### HasImages

`func (o *GetImageListResult) HasImages() bool`

HasImages returns a boolean if a field has been set.

### GetMetadata

`func (o *GetImageListResult) GetMetadata() Metadata`

GetMetadata returns the Metadata field if non-nil, zero value otherwise.

### GetMetadataOk

`func (o *GetImageListResult) GetMetadataOk() (*Metadata, bool)`

GetMetadataOk returns a tuple with the Metadata field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetadata

`func (o *GetImageListResult) SetMetadata(v Metadata)`

SetMetadata sets Metadata field to given value.

### HasMetadata

`func (o *GetImageListResult) HasMetadata() bool`

HasMetadata returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


