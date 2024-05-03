# Port

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AutoNegotiation** | Pointer to [**AutoNegotiation**](AutoNegotiation.md) |  | [optional] 
**Description** | Pointer to **string** |  | [optional] 

## Methods

### NewPort

`func NewPort() *Port`

NewPort instantiates a new Port object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPortWithDefaults

`func NewPortWithDefaults() *Port`

NewPortWithDefaults instantiates a new Port object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAutoNegotiation

`func (o *Port) GetAutoNegotiation() AutoNegotiation`

GetAutoNegotiation returns the AutoNegotiation field if non-nil, zero value otherwise.

### GetAutoNegotiationOk

`func (o *Port) GetAutoNegotiationOk() (*AutoNegotiation, bool)`

GetAutoNegotiationOk returns a tuple with the AutoNegotiation field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAutoNegotiation

`func (o *Port) SetAutoNegotiation(v AutoNegotiation)`

SetAutoNegotiation sets AutoNegotiation field to given value.

### HasAutoNegotiation

`func (o *Port) HasAutoNegotiation() bool`

HasAutoNegotiation returns a boolean if a field has been set.

### GetDescription

`func (o *Port) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *Port) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *Port) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *Port) HasDescription() bool`

HasDescription returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


