# GetServerHardwareScanListResult

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Metadata** | Pointer to [**Metadata**](Metadata.md) |  | [optional] 
**HardwareScans** | Pointer to [**[]HardwareScan**](HardwareScan.md) | Simple list of Hardware Scans | [optional] 

## Methods

### NewGetServerHardwareScanListResult

`func NewGetServerHardwareScanListResult() *GetServerHardwareScanListResult`

NewGetServerHardwareScanListResult instantiates a new GetServerHardwareScanListResult object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetServerHardwareScanListResultWithDefaults

`func NewGetServerHardwareScanListResultWithDefaults() *GetServerHardwareScanListResult`

NewGetServerHardwareScanListResultWithDefaults instantiates a new GetServerHardwareScanListResult object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetMetadata

`func (o *GetServerHardwareScanListResult) GetMetadata() Metadata`

GetMetadata returns the Metadata field if non-nil, zero value otherwise.

### GetMetadataOk

`func (o *GetServerHardwareScanListResult) GetMetadataOk() (*Metadata, bool)`

GetMetadataOk returns a tuple with the Metadata field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetadata

`func (o *GetServerHardwareScanListResult) SetMetadata(v Metadata)`

SetMetadata sets Metadata field to given value.

### HasMetadata

`func (o *GetServerHardwareScanListResult) HasMetadata() bool`

HasMetadata returns a boolean if a field has been set.

### GetHardwareScans

`func (o *GetServerHardwareScanListResult) GetHardwareScans() []HardwareScan`

GetHardwareScans returns the HardwareScans field if non-nil, zero value otherwise.

### GetHardwareScansOk

`func (o *GetServerHardwareScanListResult) GetHardwareScansOk() (*[]HardwareScan, bool)`

GetHardwareScansOk returns a tuple with the HardwareScans field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHardwareScans

`func (o *GetServerHardwareScanListResult) SetHardwareScans(v []HardwareScan)`

SetHardwareScans sets HardwareScans field to given value.

### HasHardwareScans

`func (o *GetServerHardwareScanListResult) HasHardwareScans() bool`

HasHardwareScans returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


