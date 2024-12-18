# UpdateResourceRecordSetOpts

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Content** | **[]string** | Array of resource record set content entries | 
**Ttl** | [**Ttl**](Ttl.md) |  | 

## Methods

### NewUpdateResourceRecordSetOpts

`func NewUpdateResourceRecordSetOpts(content []string, ttl Ttl, ) *UpdateResourceRecordSetOpts`

NewUpdateResourceRecordSetOpts instantiates a new UpdateResourceRecordSetOpts object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUpdateResourceRecordSetOptsWithDefaults

`func NewUpdateResourceRecordSetOptsWithDefaults() *UpdateResourceRecordSetOpts`

NewUpdateResourceRecordSetOptsWithDefaults instantiates a new UpdateResourceRecordSetOpts object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetContent

`func (o *UpdateResourceRecordSetOpts) GetContent() []string`

GetContent returns the Content field if non-nil, zero value otherwise.

### GetContentOk

`func (o *UpdateResourceRecordSetOpts) GetContentOk() (*[]string, bool)`

GetContentOk returns a tuple with the Content field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContent

`func (o *UpdateResourceRecordSetOpts) SetContent(v []string)`

SetContent sets Content field to given value.


### GetTtl

`func (o *UpdateResourceRecordSetOpts) GetTtl() Ttl`

GetTtl returns the Ttl field if non-nil, zero value otherwise.

### GetTtlOk

`func (o *UpdateResourceRecordSetOpts) GetTtlOk() (*Ttl, bool)`

GetTtlOk returns a tuple with the Ttl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTtl

`func (o *UpdateResourceRecordSetOpts) SetTtl(v Ttl)`

SetTtl sets Ttl field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


