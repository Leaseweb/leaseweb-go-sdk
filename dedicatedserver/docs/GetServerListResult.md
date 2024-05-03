# GetServerListResult

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Metadata** | Pointer to [**Metadata**](Metadata.md) |  | [optional] 
**Servers** | Pointer to [**[]Server**](Server.md) | An array of servers | [optional] 

## Methods

### NewGetServerListResult

`func NewGetServerListResult() *GetServerListResult`

NewGetServerListResult instantiates a new GetServerListResult object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetServerListResultWithDefaults

`func NewGetServerListResultWithDefaults() *GetServerListResult`

NewGetServerListResultWithDefaults instantiates a new GetServerListResult object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetMetadata

`func (o *GetServerListResult) GetMetadata() Metadata`

GetMetadata returns the Metadata field if non-nil, zero value otherwise.

### GetMetadataOk

`func (o *GetServerListResult) GetMetadataOk() (*Metadata, bool)`

GetMetadataOk returns a tuple with the Metadata field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetadata

`func (o *GetServerListResult) SetMetadata(v Metadata)`

SetMetadata sets Metadata field to given value.

### HasMetadata

`func (o *GetServerListResult) HasMetadata() bool`

HasMetadata returns a boolean if a field has been set.

### GetServers

`func (o *GetServerListResult) GetServers() []Server`

GetServers returns the Servers field if non-nil, zero value otherwise.

### GetServersOk

`func (o *GetServerListResult) GetServersOk() (*[]Server, bool)`

GetServersOk returns a tuple with the Servers field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServers

`func (o *GetServerListResult) SetServers(v []Server)`

SetServers sets Servers field to given value.

### HasServers

`func (o *GetServerListResult) HasServers() bool`

HasServers returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


