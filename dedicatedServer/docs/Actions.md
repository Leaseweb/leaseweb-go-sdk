# Actions

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**LastTriggeredAt** | Pointer to **NullableString** | Date timestamp when the last notification email triggered | [optional] 
**Type** | Pointer to **string** | Type of action | [optional] 

## Methods

### NewActions

`func NewActions() *Actions`

NewActions instantiates a new Actions object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewActionsWithDefaults

`func NewActionsWithDefaults() *Actions`

NewActionsWithDefaults instantiates a new Actions object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetLastTriggeredAt

`func (o *Actions) GetLastTriggeredAt() string`

GetLastTriggeredAt returns the LastTriggeredAt field if non-nil, zero value otherwise.

### GetLastTriggeredAtOk

`func (o *Actions) GetLastTriggeredAtOk() (*string, bool)`

GetLastTriggeredAtOk returns a tuple with the LastTriggeredAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastTriggeredAt

`func (o *Actions) SetLastTriggeredAt(v string)`

SetLastTriggeredAt sets LastTriggeredAt field to given value.

### HasLastTriggeredAt

`func (o *Actions) HasLastTriggeredAt() bool`

HasLastTriggeredAt returns a boolean if a field has been set.

### SetLastTriggeredAtNil

`func (o *Actions) SetLastTriggeredAtNil(b bool)`

 SetLastTriggeredAtNil sets the value for LastTriggeredAt to be an explicit nil

### UnsetLastTriggeredAt
`func (o *Actions) UnsetLastTriggeredAt()`

UnsetLastTriggeredAt ensures that no value is present for LastTriggeredAt, not even an explicit nil
### GetType

`func (o *Actions) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *Actions) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *Actions) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *Actions) HasType() bool`

HasType returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


