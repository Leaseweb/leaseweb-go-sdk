# LdResourceRecordSetDetails

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | **string** | Name of the resource record set | 
**Type** | [**LdResourceRecordSetType**](LdResourceRecordSetType.md) |  | 
**GeoContent** | [**GeoContent**](GeoContent.md) |  | 
**Ttl** | [**Ttl**](Ttl.md) |  | 
**Editable** | **bool** | May the set be edited | 
**Links** | Pointer to [**LdLinks**](LdLinks.md) |  | [optional] 

## Methods

### NewLdResourceRecordSetDetails

`func NewLdResourceRecordSetDetails(name string, type_ LdResourceRecordSetType, geoContent GeoContent, ttl Ttl, editable bool, ) *LdResourceRecordSetDetails`

NewLdResourceRecordSetDetails instantiates a new LdResourceRecordSetDetails object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewLdResourceRecordSetDetailsWithDefaults

`func NewLdResourceRecordSetDetailsWithDefaults() *LdResourceRecordSetDetails`

NewLdResourceRecordSetDetailsWithDefaults instantiates a new LdResourceRecordSetDetails object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *LdResourceRecordSetDetails) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *LdResourceRecordSetDetails) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *LdResourceRecordSetDetails) SetName(v string)`

SetName sets Name field to given value.


### GetType

`func (o *LdResourceRecordSetDetails) GetType() LdResourceRecordSetType`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *LdResourceRecordSetDetails) GetTypeOk() (*LdResourceRecordSetType, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *LdResourceRecordSetDetails) SetType(v LdResourceRecordSetType)`

SetType sets Type field to given value.


### GetGeoContent

`func (o *LdResourceRecordSetDetails) GetGeoContent() GeoContent`

GetGeoContent returns the GeoContent field if non-nil, zero value otherwise.

### GetGeoContentOk

`func (o *LdResourceRecordSetDetails) GetGeoContentOk() (*GeoContent, bool)`

GetGeoContentOk returns a tuple with the GeoContent field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGeoContent

`func (o *LdResourceRecordSetDetails) SetGeoContent(v GeoContent)`

SetGeoContent sets GeoContent field to given value.


### GetTtl

`func (o *LdResourceRecordSetDetails) GetTtl() Ttl`

GetTtl returns the Ttl field if non-nil, zero value otherwise.

### GetTtlOk

`func (o *LdResourceRecordSetDetails) GetTtlOk() (*Ttl, bool)`

GetTtlOk returns a tuple with the Ttl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTtl

`func (o *LdResourceRecordSetDetails) SetTtl(v Ttl)`

SetTtl sets Ttl field to given value.


### GetEditable

`func (o *LdResourceRecordSetDetails) GetEditable() bool`

GetEditable returns the Editable field if non-nil, zero value otherwise.

### GetEditableOk

`func (o *LdResourceRecordSetDetails) GetEditableOk() (*bool, bool)`

GetEditableOk returns a tuple with the Editable field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEditable

`func (o *LdResourceRecordSetDetails) SetEditable(v bool)`

SetEditable sets Editable field to given value.


### GetLinks

`func (o *LdResourceRecordSetDetails) GetLinks() LdLinks`

GetLinks returns the Links field if non-nil, zero value otherwise.

### GetLinksOk

`func (o *LdResourceRecordSetDetails) GetLinksOk() (*LdLinks, bool)`

GetLinksOk returns a tuple with the Links field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLinks

`func (o *LdResourceRecordSetDetails) SetLinks(v LdLinks)`

SetLinks sets Links field to given value.

### HasLinks

`func (o *LdResourceRecordSetDetails) HasLinks() bool`

HasLinks returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


