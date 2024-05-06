# GetServerDhcpReservationListResult

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Metadata** | Pointer to [**Metadata**](Metadata.md) |  | [optional] 
**Leases** | Pointer to [**[]Lease**](Lease.md) | An array of active DHCP reservations | [optional] 

## Methods

### NewGetServerDhcpReservationListResult

`func NewGetServerDhcpReservationListResult() *GetServerDhcpReservationListResult`

NewGetServerDhcpReservationListResult instantiates a new GetServerDhcpReservationListResult object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetServerDhcpReservationListResultWithDefaults

`func NewGetServerDhcpReservationListResultWithDefaults() *GetServerDhcpReservationListResult`

NewGetServerDhcpReservationListResultWithDefaults instantiates a new GetServerDhcpReservationListResult object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetMetadata

`func (o *GetServerDhcpReservationListResult) GetMetadata() Metadata`

GetMetadata returns the Metadata field if non-nil, zero value otherwise.

### GetMetadataOk

`func (o *GetServerDhcpReservationListResult) GetMetadataOk() (*Metadata, bool)`

GetMetadataOk returns a tuple with the Metadata field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetadata

`func (o *GetServerDhcpReservationListResult) SetMetadata(v Metadata)`

SetMetadata sets Metadata field to given value.

### HasMetadata

`func (o *GetServerDhcpReservationListResult) HasMetadata() bool`

HasMetadata returns a boolean if a field has been set.

### GetLeases

`func (o *GetServerDhcpReservationListResult) GetLeases() []Lease`

GetLeases returns the Leases field if non-nil, zero value otherwise.

### GetLeasesOk

`func (o *GetServerDhcpReservationListResult) GetLeasesOk() (*[]Lease, bool)`

GetLeasesOk returns a tuple with the Leases field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLeases

`func (o *GetServerDhcpReservationListResult) SetLeases(v []Lease)`

SetLeases sets Leases field to given value.

### HasLeases

`func (o *GetServerDhcpReservationListResult) HasLeases() bool`

HasLeases returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


