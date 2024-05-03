# ControlPanelList

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Metadata** | Pointer to [**Metadata**](Metadata.md) |  | [optional] 
**ControlPanels** | Pointer to [**[]ControlPanel**](ControlPanel.md) | A list of control panels | [optional] 

## Methods

### NewControlPanelList

`func NewControlPanelList() *ControlPanelList`

NewControlPanelList instantiates a new ControlPanelList object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewControlPanelListWithDefaults

`func NewControlPanelListWithDefaults() *ControlPanelList`

NewControlPanelListWithDefaults instantiates a new ControlPanelList object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetMetadata

`func (o *ControlPanelList) GetMetadata() Metadata`

GetMetadata returns the Metadata field if non-nil, zero value otherwise.

### GetMetadataOk

`func (o *ControlPanelList) GetMetadataOk() (*Metadata, bool)`

GetMetadataOk returns a tuple with the Metadata field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetadata

`func (o *ControlPanelList) SetMetadata(v Metadata)`

SetMetadata sets Metadata field to given value.

### HasMetadata

`func (o *ControlPanelList) HasMetadata() bool`

HasMetadata returns a boolean if a field has been set.

### GetControlPanels

`func (o *ControlPanelList) GetControlPanels() []ControlPanel`

GetControlPanels returns the ControlPanels field if non-nil, zero value otherwise.

### GetControlPanelsOk

`func (o *ControlPanelList) GetControlPanelsOk() (*[]ControlPanel, bool)`

GetControlPanelsOk returns a tuple with the ControlPanels field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetControlPanels

`func (o *ControlPanelList) SetControlPanels(v []ControlPanel)`

SetControlPanels sets ControlPanels field to given value.

### HasControlPanels

`func (o *ControlPanelList) HasControlPanels() bool`

HasControlPanels returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


