# LdResourceRecordSet

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | **string** | Name of the resource record set | 
**Type** | [**LdResourceRecordSetType**](LdResourceRecordSetType.md) |  | 
**GeoContent** | [**GeoContent**](GeoContent.md) |  | 
**Ttl** | [**Ttl**](Ttl.md) |  | 

## Methods

### NewLdResourceRecordSet

`func NewLdResourceRecordSet(name string, type_ LdResourceRecordSetType, geoContent GeoContent, ttl Ttl, ) *LdResourceRecordSet`

NewLdResourceRecordSet instantiates a new LdResourceRecordSet object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewLdResourceRecordSetWithDefaults

`func NewLdResourceRecordSetWithDefaults() *LdResourceRecordSet`

NewLdResourceRecordSetWithDefaults instantiates a new LdResourceRecordSet object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *LdResourceRecordSet) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *LdResourceRecordSet) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *LdResourceRecordSet) SetName(v string)`

SetName sets Name field to given value.


### GetType

`func (o *LdResourceRecordSet) GetType() LdResourceRecordSetType`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *LdResourceRecordSet) GetTypeOk() (*LdResourceRecordSetType, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *LdResourceRecordSet) SetType(v LdResourceRecordSetType)`

SetType sets Type field to given value.


### GetGeoContent

`func (o *LdResourceRecordSet) GetGeoContent() GeoContent`

GetGeoContent returns the GeoContent field if non-nil, zero value otherwise.

### GetGeoContentOk

`func (o *LdResourceRecordSet) GetGeoContentOk() (*GeoContent, bool)`

GetGeoContentOk returns a tuple with the GeoContent field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGeoContent

`func (o *LdResourceRecordSet) SetGeoContent(v GeoContent)`

SetGeoContent sets GeoContent field to given value.


### GetTtl

`func (o *LdResourceRecordSet) GetTtl() Ttl`

GetTtl returns the Ttl field if non-nil, zero value otherwise.

### GetTtlOk

`func (o *LdResourceRecordSet) GetTtlOk() (*Ttl, bool)`

GetTtlOk returns a tuple with the Ttl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTtl

`func (o *LdResourceRecordSet) SetTtl(v Ttl)`

SetTtl sets Ttl field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


