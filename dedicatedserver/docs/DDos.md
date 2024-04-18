# DDos

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**DetectionProfile** | Pointer to **string** | The applied detection profile | [optional] 
**ProtectionType** | Pointer to **string** | The type of DDoS protection | [optional] 

## Methods

### NewDDos

`func NewDDos() *DDos`

NewDDos instantiates a new DDos object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDDosWithDefaults

`func NewDDosWithDefaults() *DDos`

NewDDosWithDefaults instantiates a new DDos object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDetectionProfile

`func (o *DDos) GetDetectionProfile() string`

GetDetectionProfile returns the DetectionProfile field if non-nil, zero value otherwise.

### GetDetectionProfileOk

`func (o *DDos) GetDetectionProfileOk() (*string, bool)`

GetDetectionProfileOk returns a tuple with the DetectionProfile field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDetectionProfile

`func (o *DDos) SetDetectionProfile(v string)`

SetDetectionProfile sets DetectionProfile field to given value.

### HasDetectionProfile

`func (o *DDos) HasDetectionProfile() bool`

HasDetectionProfile returns a boolean if a field has been set.

### GetProtectionType

`func (o *DDos) GetProtectionType() string`

GetProtectionType returns the ProtectionType field if non-nil, zero value otherwise.

### GetProtectionTypeOk

`func (o *DDos) GetProtectionTypeOk() (*string, bool)`

GetProtectionTypeOk returns a tuple with the ProtectionType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProtectionType

`func (o *DDos) SetProtectionType(v string)`

SetProtectionType sets ProtectionType field to given value.

### HasProtectionType

`func (o *DDos) HasProtectionType() bool`

HasProtectionType returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


