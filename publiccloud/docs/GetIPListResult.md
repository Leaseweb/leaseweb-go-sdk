# GetIPListResult

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Ips** | Pointer to [**[]IpDetails**](IpDetails.md) |  | [optional] 
**Metadata** | Pointer to [**Metadata**](Metadata.md) |  | [optional] 

## Methods

### NewGetIPListResult

`func NewGetIPListResult() *GetIPListResult`

NewGetIPListResult instantiates a new GetIPListResult object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetIPListResultWithDefaults

`func NewGetIPListResultWithDefaults() *GetIPListResult`

NewGetIPListResultWithDefaults instantiates a new GetIPListResult object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetIps

`func (o *GetIPListResult) GetIps() []IpDetails`

GetIps returns the Ips field if non-nil, zero value otherwise.

### GetIpsOk

`func (o *GetIPListResult) GetIpsOk() (*[]IpDetails, bool)`

GetIpsOk returns a tuple with the Ips field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIps

`func (o *GetIPListResult) SetIps(v []IpDetails)`

SetIps sets Ips field to given value.

### HasIps

`func (o *GetIPListResult) HasIps() bool`

HasIps returns a boolean if a field has been set.

### GetMetadata

`func (o *GetIPListResult) GetMetadata() Metadata`

GetMetadata returns the Metadata field if non-nil, zero value otherwise.

### GetMetadataOk

`func (o *GetIPListResult) GetMetadataOk() (*Metadata, bool)`

GetMetadataOk returns a tuple with the Metadata field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetadata

`func (o *GetIPListResult) SetMetadata(v Metadata)`

SetMetadata sets Metadata field to given value.

### HasMetadata

`func (o *GetIPListResult) HasMetadata() bool`

HasMetadata returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


