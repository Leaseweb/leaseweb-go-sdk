# GetMarketAppListResult

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**MarketApps** | Pointer to [**[]MarketApp**](MarketApp.md) |  | [optional] 
**Metadata** | Pointer to [**Metadata**](Metadata.md) |  | [optional] 

## Methods

### NewGetMarketAppListResult

`func NewGetMarketAppListResult() *GetMarketAppListResult`

NewGetMarketAppListResult instantiates a new GetMarketAppListResult object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetMarketAppListResultWithDefaults

`func NewGetMarketAppListResultWithDefaults() *GetMarketAppListResult`

NewGetMarketAppListResultWithDefaults instantiates a new GetMarketAppListResult object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetMarketApps

`func (o *GetMarketAppListResult) GetMarketApps() []MarketApp`

GetMarketApps returns the MarketApps field if non-nil, zero value otherwise.

### GetMarketAppsOk

`func (o *GetMarketAppListResult) GetMarketAppsOk() (*[]MarketApp, bool)`

GetMarketAppsOk returns a tuple with the MarketApps field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMarketApps

`func (o *GetMarketAppListResult) SetMarketApps(v []MarketApp)`

SetMarketApps sets MarketApps field to given value.

### HasMarketApps

`func (o *GetMarketAppListResult) HasMarketApps() bool`

HasMarketApps returns a boolean if a field has been set.

### GetMetadata

`func (o *GetMarketAppListResult) GetMetadata() Metadata`

GetMetadata returns the Metadata field if non-nil, zero value otherwise.

### GetMetadataOk

`func (o *GetMarketAppListResult) GetMetadataOk() (*Metadata, bool)`

GetMetadataOk returns a tuple with the Metadata field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetadata

`func (o *GetMarketAppListResult) SetMetadata(v Metadata)`

SetMetadata sets Metadata field to given value.

### HasMetadata

`func (o *GetMarketAppListResult) HasMetadata() bool`

HasMetadata returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


