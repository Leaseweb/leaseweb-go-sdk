# GetIPListResult

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Ips** | [**[]Ip**](Ip.md) |  | 
**Metadata** | [**Metadata**](Metadata.md) |  | 

## Methods

### NewGetIPListResult

`func NewGetIPListResult(ips []Ip, metadata Metadata, ) *GetIPListResult`

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

`func (o *GetIPListResult) GetIps() []Ip`

GetIps returns the Ips field if non-nil, zero value otherwise.

### GetIpsOk

`func (o *GetIPListResult) GetIpsOk() (*[]Ip, bool)`

GetIpsOk returns a tuple with the Ips field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIps

`func (o *GetIPListResult) SetIps(v []Ip)`

SetIps sets Ips field to given value.


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



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


