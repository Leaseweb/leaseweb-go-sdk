# LldpChassis

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Description** | Pointer to **string** |  | [optional] 
**MacAddress** | Pointer to **string** | Represents a MAC Address in the standard colon delimited format. Eg. &#x60;01:23:45:67:89:0A&#x60; | [optional] 
**Name** | Pointer to **string** |  | [optional] 

## Methods

### NewLldpChassis

`func NewLldpChassis() *LldpChassis`

NewLldpChassis instantiates a new LldpChassis object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewLldpChassisWithDefaults

`func NewLldpChassisWithDefaults() *LldpChassis`

NewLldpChassisWithDefaults instantiates a new LldpChassis object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDescription

`func (o *LldpChassis) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *LldpChassis) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *LldpChassis) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *LldpChassis) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetMacAddress

`func (o *LldpChassis) GetMacAddress() string`

GetMacAddress returns the MacAddress field if non-nil, zero value otherwise.

### GetMacAddressOk

`func (o *LldpChassis) GetMacAddressOk() (*string, bool)`

GetMacAddressOk returns a tuple with the MacAddress field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMacAddress

`func (o *LldpChassis) SetMacAddress(v string)`

SetMacAddress sets MacAddress field to given value.

### HasMacAddress

`func (o *LldpChassis) HasMacAddress() bool`

HasMacAddress returns a boolean if a field has been set.

### GetName

`func (o *LldpChassis) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *LldpChassis) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *LldpChassis) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *LldpChassis) HasName() bool`

HasName returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


