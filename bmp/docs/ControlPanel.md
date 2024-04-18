# ControlPanel

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **string** | The unique ID of this control panel | [optional] 
**Name** | Pointer to **string** | A human readable name describing the control panel | [optional] 

## Methods

### NewControlPanel

`func NewControlPanel() *ControlPanel`

NewControlPanel instantiates a new ControlPanel object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewControlPanelWithDefaults

`func NewControlPanelWithDefaults() *ControlPanel`

NewControlPanelWithDefaults instantiates a new ControlPanel object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *ControlPanel) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *ControlPanel) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *ControlPanel) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *ControlPanel) HasId() bool`

HasId returns a boolean if a field has been set.

### GetName

`func (o *ControlPanel) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *ControlPanel) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *ControlPanel) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *ControlPanel) HasName() bool`

HasName returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


