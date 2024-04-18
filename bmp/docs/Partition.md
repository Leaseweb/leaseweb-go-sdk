# Partition

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Filesystem** | Pointer to **string** | File system in which partition would be mounted | [optional] 
**Mountpoint** | Pointer to **string** | The partition mount point (eg &#x60;/home&#x60;). Mandatory for the root partition (&#x60;/&#x60;) and not intended to be used in swap partition | [optional] 
**Size** | Pointer to **string** | Size of the partition (Normally in MB, but this is OS-specific) | [optional] 

## Methods

### NewPartition

`func NewPartition() *Partition`

NewPartition instantiates a new Partition object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPartitionWithDefaults

`func NewPartitionWithDefaults() *Partition`

NewPartitionWithDefaults instantiates a new Partition object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetFilesystem

`func (o *Partition) GetFilesystem() string`

GetFilesystem returns the Filesystem field if non-nil, zero value otherwise.

### GetFilesystemOk

`func (o *Partition) GetFilesystemOk() (*string, bool)`

GetFilesystemOk returns a tuple with the Filesystem field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFilesystem

`func (o *Partition) SetFilesystem(v string)`

SetFilesystem sets Filesystem field to given value.

### HasFilesystem

`func (o *Partition) HasFilesystem() bool`

HasFilesystem returns a boolean if a field has been set.

### GetMountpoint

`func (o *Partition) GetMountpoint() string`

GetMountpoint returns the Mountpoint field if non-nil, zero value otherwise.

### GetMountpointOk

`func (o *Partition) GetMountpointOk() (*string, bool)`

GetMountpointOk returns a tuple with the Mountpoint field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMountpoint

`func (o *Partition) SetMountpoint(v string)`

SetMountpoint sets Mountpoint field to given value.

### HasMountpoint

`func (o *Partition) HasMountpoint() bool`

HasMountpoint returns a boolean if a field has been set.

### GetSize

`func (o *Partition) GetSize() string`

GetSize returns the Size field if non-nil, zero value otherwise.

### GetSizeOk

`func (o *Partition) GetSizeOk() (*string, bool)`

GetSizeOk returns a tuple with the Size field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSize

`func (o *Partition) SetSize(v string)`

SetSize sets Size field to given value.

### HasSize

`func (o *Partition) HasSize() bool`

HasSize returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


