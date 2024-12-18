# ResourceRecordSet

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | **string** | Name of the resource record set | 
**Type** | [**ResourceRecordSetType**](ResourceRecordSetType.md) |  | 
**Content** | **[]string** | Array of resource record set content entries | 
**Ttl** | [**Ttl**](Ttl.md) |  | 

## Methods

### NewResourceRecordSet

`func NewResourceRecordSet(name string, type_ ResourceRecordSetType, content []string, ttl Ttl, ) *ResourceRecordSet`

NewResourceRecordSet instantiates a new ResourceRecordSet object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewResourceRecordSetWithDefaults

`func NewResourceRecordSetWithDefaults() *ResourceRecordSet`

NewResourceRecordSetWithDefaults instantiates a new ResourceRecordSet object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *ResourceRecordSet) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *ResourceRecordSet) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *ResourceRecordSet) SetName(v string)`

SetName sets Name field to given value.


### GetType

`func (o *ResourceRecordSet) GetType() ResourceRecordSetType`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *ResourceRecordSet) GetTypeOk() (*ResourceRecordSetType, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *ResourceRecordSet) SetType(v ResourceRecordSetType)`

SetType sets Type field to given value.


### GetContent

`func (o *ResourceRecordSet) GetContent() []string`

GetContent returns the Content field if non-nil, zero value otherwise.

### GetContentOk

`func (o *ResourceRecordSet) GetContentOk() (*[]string, bool)`

GetContentOk returns a tuple with the Content field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContent

`func (o *ResourceRecordSet) SetContent(v []string)`

SetContent sets Content field to given value.


### GetTtl

`func (o *ResourceRecordSet) GetTtl() Ttl`

GetTtl returns the Ttl field if non-nil, zero value otherwise.

### GetTtlOk

`func (o *ResourceRecordSet) GetTtlOk() (*Ttl, bool)`

GetTtlOk returns a tuple with the Ttl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTtl

`func (o *ResourceRecordSet) SetTtl(v Ttl)`

SetTtl sets Ttl field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


