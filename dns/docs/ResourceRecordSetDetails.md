# ResourceRecordSetDetails

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | **string** | Name of the resource record set | 
**Type** | [**ResourceRecordSetType**](ResourceRecordSetType.md) |  | 
**Content** | **[]string** | Array of resource record set content entries | 
**Ttl** | [**Ttl**](Ttl.md) |  | 
**Editable** | **bool** | May the set be edited | 
**Links** | Pointer to [**LdLinks**](LdLinks.md) |  | [optional] 

## Methods

### NewResourceRecordSetDetails

`func NewResourceRecordSetDetails(name string, type_ ResourceRecordSetType, content []string, ttl Ttl, editable bool, ) *ResourceRecordSetDetails`

NewResourceRecordSetDetails instantiates a new ResourceRecordSetDetails object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewResourceRecordSetDetailsWithDefaults

`func NewResourceRecordSetDetailsWithDefaults() *ResourceRecordSetDetails`

NewResourceRecordSetDetailsWithDefaults instantiates a new ResourceRecordSetDetails object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *ResourceRecordSetDetails) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *ResourceRecordSetDetails) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *ResourceRecordSetDetails) SetName(v string)`

SetName sets Name field to given value.


### GetType

`func (o *ResourceRecordSetDetails) GetType() ResourceRecordSetType`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *ResourceRecordSetDetails) GetTypeOk() (*ResourceRecordSetType, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *ResourceRecordSetDetails) SetType(v ResourceRecordSetType)`

SetType sets Type field to given value.


### GetContent

`func (o *ResourceRecordSetDetails) GetContent() []string`

GetContent returns the Content field if non-nil, zero value otherwise.

### GetContentOk

`func (o *ResourceRecordSetDetails) GetContentOk() (*[]string, bool)`

GetContentOk returns a tuple with the Content field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContent

`func (o *ResourceRecordSetDetails) SetContent(v []string)`

SetContent sets Content field to given value.


### GetTtl

`func (o *ResourceRecordSetDetails) GetTtl() Ttl`

GetTtl returns the Ttl field if non-nil, zero value otherwise.

### GetTtlOk

`func (o *ResourceRecordSetDetails) GetTtlOk() (*Ttl, bool)`

GetTtlOk returns a tuple with the Ttl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTtl

`func (o *ResourceRecordSetDetails) SetTtl(v Ttl)`

SetTtl sets Ttl field to given value.


### GetEditable

`func (o *ResourceRecordSetDetails) GetEditable() bool`

GetEditable returns the Editable field if non-nil, zero value otherwise.

### GetEditableOk

`func (o *ResourceRecordSetDetails) GetEditableOk() (*bool, bool)`

GetEditableOk returns a tuple with the Editable field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEditable

`func (o *ResourceRecordSetDetails) SetEditable(v bool)`

SetEditable sets Editable field to given value.


### GetLinks

`func (o *ResourceRecordSetDetails) GetLinks() LdLinks`

GetLinks returns the Links field if non-nil, zero value otherwise.

### GetLinksOk

`func (o *ResourceRecordSetDetails) GetLinksOk() (*LdLinks, bool)`

GetLinksOk returns a tuple with the Links field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLinks

`func (o *ResourceRecordSetDetails) SetLinks(v LdLinks)`

SetLinks sets Links field to given value.

### HasLinks

`func (o *ResourceRecordSetDetails) HasLinks() bool`

HasLinks returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


