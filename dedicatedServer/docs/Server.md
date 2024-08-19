# Server

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AssetId** | Pointer to **string** | The Asset Id of the server | [optional] 
**Contract** | Pointer to [**Contract**](Contract.md) |  | [optional] 
**FeatureAvailability** | Pointer to [**FeatureAvailability**](FeatureAvailability.md) |  | [optional] 
**Id** | Pointer to **string** | Id of the server | [optional] 
**Location** | Pointer to [**Location**](Location.md) |  | [optional] 
**NetworkInterfaces** | Pointer to [**NetworkInterfaces**](NetworkInterfaces.md) |  | [optional] 
**PowerPorts** | Pointer to [**[]Port**](Port.md) | List of ports that can be used to manage power of the server | [optional] 
**PrivateNetworks** | Pointer to [**[]PrivateNetwork**](PrivateNetwork.md) | An array of private networks | [optional] 
**Rack** | Pointer to [**Rack**](Rack.md) |  | [optional] 
**SerialNumber** | Pointer to **string** | Serial number of server | [optional] 
**Specs** | Pointer to [**ServerSpecs**](ServerSpecs.md) |  | [optional] 

## Methods

### NewServer

`func NewServer() *Server`

NewServer instantiates a new Server object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewServerWithDefaults

`func NewServerWithDefaults() *Server`

NewServerWithDefaults instantiates a new Server object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAssetId

`func (o *Server) GetAssetId() string`

GetAssetId returns the AssetId field if non-nil, zero value otherwise.

### GetAssetIdOk

`func (o *Server) GetAssetIdOk() (*string, bool)`

GetAssetIdOk returns a tuple with the AssetId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAssetId

`func (o *Server) SetAssetId(v string)`

SetAssetId sets AssetId field to given value.

### HasAssetId

`func (o *Server) HasAssetId() bool`

HasAssetId returns a boolean if a field has been set.

### GetContract

`func (o *Server) GetContract() Contract`

GetContract returns the Contract field if non-nil, zero value otherwise.

### GetContractOk

`func (o *Server) GetContractOk() (*Contract, bool)`

GetContractOk returns a tuple with the Contract field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContract

`func (o *Server) SetContract(v Contract)`

SetContract sets Contract field to given value.

### HasContract

`func (o *Server) HasContract() bool`

HasContract returns a boolean if a field has been set.

### GetFeatureAvailability

`func (o *Server) GetFeatureAvailability() FeatureAvailability`

GetFeatureAvailability returns the FeatureAvailability field if non-nil, zero value otherwise.

### GetFeatureAvailabilityOk

`func (o *Server) GetFeatureAvailabilityOk() (*FeatureAvailability, bool)`

GetFeatureAvailabilityOk returns a tuple with the FeatureAvailability field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFeatureAvailability

`func (o *Server) SetFeatureAvailability(v FeatureAvailability)`

SetFeatureAvailability sets FeatureAvailability field to given value.

### HasFeatureAvailability

`func (o *Server) HasFeatureAvailability() bool`

HasFeatureAvailability returns a boolean if a field has been set.

### GetId

`func (o *Server) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *Server) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *Server) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *Server) HasId() bool`

HasId returns a boolean if a field has been set.

### GetLocation

`func (o *Server) GetLocation() Location`

GetLocation returns the Location field if non-nil, zero value otherwise.

### GetLocationOk

`func (o *Server) GetLocationOk() (*Location, bool)`

GetLocationOk returns a tuple with the Location field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLocation

`func (o *Server) SetLocation(v Location)`

SetLocation sets Location field to given value.

### HasLocation

`func (o *Server) HasLocation() bool`

HasLocation returns a boolean if a field has been set.

### GetNetworkInterfaces

`func (o *Server) GetNetworkInterfaces() NetworkInterfaces`

GetNetworkInterfaces returns the NetworkInterfaces field if non-nil, zero value otherwise.

### GetNetworkInterfacesOk

`func (o *Server) GetNetworkInterfacesOk() (*NetworkInterfaces, bool)`

GetNetworkInterfacesOk returns a tuple with the NetworkInterfaces field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNetworkInterfaces

`func (o *Server) SetNetworkInterfaces(v NetworkInterfaces)`

SetNetworkInterfaces sets NetworkInterfaces field to given value.

### HasNetworkInterfaces

`func (o *Server) HasNetworkInterfaces() bool`

HasNetworkInterfaces returns a boolean if a field has been set.

### GetPowerPorts

`func (o *Server) GetPowerPorts() []Port`

GetPowerPorts returns the PowerPorts field if non-nil, zero value otherwise.

### GetPowerPortsOk

`func (o *Server) GetPowerPortsOk() (*[]Port, bool)`

GetPowerPortsOk returns a tuple with the PowerPorts field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPowerPorts

`func (o *Server) SetPowerPorts(v []Port)`

SetPowerPorts sets PowerPorts field to given value.

### HasPowerPorts

`func (o *Server) HasPowerPorts() bool`

HasPowerPorts returns a boolean if a field has been set.

### GetPrivateNetworks

`func (o *Server) GetPrivateNetworks() []PrivateNetwork`

GetPrivateNetworks returns the PrivateNetworks field if non-nil, zero value otherwise.

### GetPrivateNetworksOk

`func (o *Server) GetPrivateNetworksOk() (*[]PrivateNetwork, bool)`

GetPrivateNetworksOk returns a tuple with the PrivateNetworks field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrivateNetworks

`func (o *Server) SetPrivateNetworks(v []PrivateNetwork)`

SetPrivateNetworks sets PrivateNetworks field to given value.

### HasPrivateNetworks

`func (o *Server) HasPrivateNetworks() bool`

HasPrivateNetworks returns a boolean if a field has been set.

### GetRack

`func (o *Server) GetRack() Rack`

GetRack returns the Rack field if non-nil, zero value otherwise.

### GetRackOk

`func (o *Server) GetRackOk() (*Rack, bool)`

GetRackOk returns a tuple with the Rack field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRack

`func (o *Server) SetRack(v Rack)`

SetRack sets Rack field to given value.

### HasRack

`func (o *Server) HasRack() bool`

HasRack returns a boolean if a field has been set.

### GetSerialNumber

`func (o *Server) GetSerialNumber() string`

GetSerialNumber returns the SerialNumber field if non-nil, zero value otherwise.

### GetSerialNumberOk

`func (o *Server) GetSerialNumberOk() (*string, bool)`

GetSerialNumberOk returns a tuple with the SerialNumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSerialNumber

`func (o *Server) SetSerialNumber(v string)`

SetSerialNumber sets SerialNumber field to given value.

### HasSerialNumber

`func (o *Server) HasSerialNumber() bool`

HasSerialNumber returns a boolean if a field has been set.

### GetSpecs

`func (o *Server) GetSpecs() ServerSpecs`

GetSpecs returns the Specs field if non-nil, zero value otherwise.

### GetSpecsOk

`func (o *Server) GetSpecsOk() (*ServerSpecs, bool)`

GetSpecsOk returns a tuple with the Specs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSpecs

`func (o *Server) SetSpecs(v ServerSpecs)`

SetSpecs sets Specs field to given value.

### HasSpecs

`func (o *Server) HasSpecs() bool`

HasSpecs returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


