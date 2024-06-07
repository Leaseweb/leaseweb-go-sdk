# NetworkInterfaces

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Public** | Pointer to [**NetworkInterface**](NetworkInterface.md) |  | [optional] 
**Internal** | Pointer to [**NetworkInterface**](NetworkInterface.md) |  | [optional] 
**RemoteManagement** | Pointer to [**NetworkInterface**](NetworkInterface.md) |  | [optional] 

## Methods

### NewNetworkInterfaces

`func NewNetworkInterfaces() *NetworkInterfaces`

NewNetworkInterfaces instantiates a new NetworkInterfaces object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewNetworkInterfacesWithDefaults

`func NewNetworkInterfacesWithDefaults() *NetworkInterfaces`

NewNetworkInterfacesWithDefaults instantiates a new NetworkInterfaces object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetPublic

`func (o *NetworkInterfaces) GetPublic() NetworkInterface`

GetPublic returns the Public field if non-nil, zero value otherwise.

### GetPublicOk

`func (o *NetworkInterfaces) GetPublicOk() (*NetworkInterface, bool)`

GetPublicOk returns a tuple with the Public field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPublic

`func (o *NetworkInterfaces) SetPublic(v NetworkInterface)`

SetPublic sets Public field to given value.

### HasPublic

`func (o *NetworkInterfaces) HasPublic() bool`

HasPublic returns a boolean if a field has been set.

### GetInternal

`func (o *NetworkInterfaces) GetInternal() NetworkInterface`

GetInternal returns the Internal field if non-nil, zero value otherwise.

### GetInternalOk

`func (o *NetworkInterfaces) GetInternalOk() (*NetworkInterface, bool)`

GetInternalOk returns a tuple with the Internal field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInternal

`func (o *NetworkInterfaces) SetInternal(v NetworkInterface)`

SetInternal sets Internal field to given value.

### HasInternal

`func (o *NetworkInterfaces) HasInternal() bool`

HasInternal returns a boolean if a field has been set.

### GetRemoteManagement

`func (o *NetworkInterfaces) GetRemoteManagement() NetworkInterface`

GetRemoteManagement returns the RemoteManagement field if non-nil, zero value otherwise.

### GetRemoteManagementOk

`func (o *NetworkInterfaces) GetRemoteManagementOk() (*NetworkInterface, bool)`

GetRemoteManagementOk returns a tuple with the RemoteManagement field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRemoteManagement

`func (o *NetworkInterfaces) SetRemoteManagement(v NetworkInterface)`

SetRemoteManagement sets RemoteManagement field to given value.

### HasRemoteManagement

`func (o *NetworkInterfaces) HasRemoteManagement() bool`

HasRemoteManagement returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


