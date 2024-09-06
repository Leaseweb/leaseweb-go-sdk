# Rack

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **string** | Rack id | [optional] 
**Capacity** | Pointer to **string** | Rack capacity | [optional] 
**Type** | Pointer to [**RackType**](RackType.md) |  | [optional] 

## Methods

### NewRack

`func NewRack() *Rack`

NewRack instantiates a new Rack object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRackWithDefaults

`func NewRackWithDefaults() *Rack`

NewRackWithDefaults instantiates a new Rack object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *Rack) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *Rack) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *Rack) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *Rack) HasId() bool`

HasId returns a boolean if a field has been set.

### GetCapacity

`func (o *Rack) GetCapacity() string`

GetCapacity returns the Capacity field if non-nil, zero value otherwise.

### GetCapacityOk

`func (o *Rack) GetCapacityOk() (*string, bool)`

GetCapacityOk returns a tuple with the Capacity field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCapacity

`func (o *Rack) SetCapacity(v string)`

SetCapacity sets Capacity field to given value.

### HasCapacity

`func (o *Rack) HasCapacity() bool`

HasCapacity returns a boolean if a field has been set.

### GetType

`func (o *Rack) GetType() RackType`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *Rack) GetTypeOk() (*RackType, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *Rack) SetType(v RackType)`

SetType sets Type field to given value.

### HasType

`func (o *Rack) HasType() bool`

HasType returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


