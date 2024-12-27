# GetNullRouteHistoryListResult

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Nullroutes** | [**[]NullRoutedIP**](NullRoutedIP.md) |  | 
**Metadata** | [**Metadata**](Metadata.md) |  | 

## Methods

### NewGetNullRouteHistoryListResult

`func NewGetNullRouteHistoryListResult(nullroutes []NullRoutedIP, metadata Metadata, ) *GetNullRouteHistoryListResult`

NewGetNullRouteHistoryListResult instantiates a new GetNullRouteHistoryListResult object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetNullRouteHistoryListResultWithDefaults

`func NewGetNullRouteHistoryListResultWithDefaults() *GetNullRouteHistoryListResult`

NewGetNullRouteHistoryListResultWithDefaults instantiates a new GetNullRouteHistoryListResult object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetNullroutes

`func (o *GetNullRouteHistoryListResult) GetNullroutes() []NullRoutedIP`

GetNullroutes returns the Nullroutes field if non-nil, zero value otherwise.

### GetNullroutesOk

`func (o *GetNullRouteHistoryListResult) GetNullroutesOk() (*[]NullRoutedIP, bool)`

GetNullroutesOk returns a tuple with the Nullroutes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNullroutes

`func (o *GetNullRouteHistoryListResult) SetNullroutes(v []NullRoutedIP)`

SetNullroutes sets Nullroutes field to given value.


### GetMetadata

`func (o *GetNullRouteHistoryListResult) GetMetadata() Metadata`

GetMetadata returns the Metadata field if non-nil, zero value otherwise.

### GetMetadataOk

`func (o *GetNullRouteHistoryListResult) GetMetadataOk() (*Metadata, bool)`

GetMetadataOk returns a tuple with the Metadata field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetadata

`func (o *GetNullRouteHistoryListResult) SetMetadata(v Metadata)`

SetMetadata sets Metadata field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


