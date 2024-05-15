# Defaults

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Device** | Pointer to **string** | Device name | [optional] 
**Partitions** | Pointer to [**[]Partition**](Partition.md) |  | [optional] 

## Methods

### NewDefaults

`func NewDefaults() *Defaults`

NewDefaults instantiates a new Defaults object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDefaultsWithDefaults

`func NewDefaultsWithDefaults() *Defaults`

NewDefaultsWithDefaults instantiates a new Defaults object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDevice

`func (o *Defaults) GetDevice() string`

GetDevice returns the Device field if non-nil, zero value otherwise.

### GetDeviceOk

`func (o *Defaults) GetDeviceOk() (*string, bool)`

GetDeviceOk returns a tuple with the Device field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDevice

`func (o *Defaults) SetDevice(v string)`

SetDevice sets Device field to given value.

### HasDevice

`func (o *Defaults) HasDevice() bool`

HasDevice returns a boolean if a field has been set.

### GetPartitions

`func (o *Defaults) GetPartitions() []Partition`

GetPartitions returns the Partitions field if non-nil, zero value otherwise.

### GetPartitionsOk

`func (o *Defaults) GetPartitionsOk() (*[]Partition, bool)`

GetPartitionsOk returns a tuple with the Partitions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPartitions

`func (o *Defaults) SetPartitions(v []Partition)`

SetPartitions sets Partitions field to given value.

### HasPartitions

`func (o *Defaults) HasPartitions() bool`

HasPartitions returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


