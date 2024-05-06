# GetSnapshotListResult

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Snapshots** | Pointer to [**[]Snapshot**](Snapshot.md) |  | [optional] 
**Metadata** | Pointer to [**Metadata**](Metadata.md) |  | [optional] 

## Methods

### NewGetSnapshotListResult

`func NewGetSnapshotListResult() *GetSnapshotListResult`

NewGetSnapshotListResult instantiates a new GetSnapshotListResult object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetSnapshotListResultWithDefaults

`func NewGetSnapshotListResultWithDefaults() *GetSnapshotListResult`

NewGetSnapshotListResultWithDefaults instantiates a new GetSnapshotListResult object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSnapshots

`func (o *GetSnapshotListResult) GetSnapshots() []Snapshot`

GetSnapshots returns the Snapshots field if non-nil, zero value otherwise.

### GetSnapshotsOk

`func (o *GetSnapshotListResult) GetSnapshotsOk() (*[]Snapshot, bool)`

GetSnapshotsOk returns a tuple with the Snapshots field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSnapshots

`func (o *GetSnapshotListResult) SetSnapshots(v []Snapshot)`

SetSnapshots sets Snapshots field to given value.

### HasSnapshots

`func (o *GetSnapshotListResult) HasSnapshots() bool`

HasSnapshots returns a boolean if a field has been set.

### GetMetadata

`func (o *GetSnapshotListResult) GetMetadata() Metadata`

GetMetadata returns the Metadata field if non-nil, zero value otherwise.

### GetMetadataOk

`func (o *GetSnapshotListResult) GetMetadataOk() (*Metadata, bool)`

GetMetadataOk returns a tuple with the Metadata field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetadata

`func (o *GetSnapshotListResult) SetMetadata(v Metadata)`

SetMetadata sets Metadata field to given value.

### HasMetadata

`func (o *GetSnapshotListResult) HasMetadata() bool`

HasMetadata returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


