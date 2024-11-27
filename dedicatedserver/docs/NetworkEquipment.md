# NetworkEquipment

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Contract** | Pointer to [**Contract**](Contract.md) |  | [optional] 
**FeatureAvailability** | Pointer to [**FeatureAvailability**](FeatureAvailability.md) |  | [optional] 
**Id** | Pointer to **string** | Id of the network equipment | [optional] 
**Location** | Pointer to [**Location**](Location.md) |  | [optional] 
**Name** | Pointer to **string** | The name of the network equipment | [optional] 
**Type** | Pointer to **string** | The network equipment type | [optional] 
**NetworkInterfaces** | Pointer to [**NetworkInterfaces**](NetworkInterfaces.md) |  | [optional] 
**PowerPorts** | Pointer to [**[]Port**](Port.md) | List of ports that can be used to manage power of the network equipment | [optional] 
**Rack** | Pointer to [**Rack**](Rack.md) |  | [optional] 
**SerialNumber** | Pointer to **string** | Serial number of network equipment | [optional] 
**Specs** | Pointer to [**NetworkEquipmentSpecs**](NetworkEquipmentSpecs.md) |  | [optional] 

## Methods

### NewNetworkEquipment

`func NewNetworkEquipment() *NetworkEquipment`

NewNetworkEquipment instantiates a new NetworkEquipment object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewNetworkEquipmentWithDefaults

`func NewNetworkEquipmentWithDefaults() *NetworkEquipment`

NewNetworkEquipmentWithDefaults instantiates a new NetworkEquipment object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetContract

`func (o *NetworkEquipment) GetContract() Contract`

GetContract returns the Contract field if non-nil, zero value otherwise.

### GetContractOk

`func (o *NetworkEquipment) GetContractOk() (*Contract, bool)`

GetContractOk returns a tuple with the Contract field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContract

`func (o *NetworkEquipment) SetContract(v Contract)`

SetContract sets Contract field to given value.

### HasContract

`func (o *NetworkEquipment) HasContract() bool`

HasContract returns a boolean if a field has been set.

### GetFeatureAvailability

`func (o *NetworkEquipment) GetFeatureAvailability() FeatureAvailability`

GetFeatureAvailability returns the FeatureAvailability field if non-nil, zero value otherwise.

### GetFeatureAvailabilityOk

`func (o *NetworkEquipment) GetFeatureAvailabilityOk() (*FeatureAvailability, bool)`

GetFeatureAvailabilityOk returns a tuple with the FeatureAvailability field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFeatureAvailability

`func (o *NetworkEquipment) SetFeatureAvailability(v FeatureAvailability)`

SetFeatureAvailability sets FeatureAvailability field to given value.

### HasFeatureAvailability

`func (o *NetworkEquipment) HasFeatureAvailability() bool`

HasFeatureAvailability returns a boolean if a field has been set.

### GetId

`func (o *NetworkEquipment) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *NetworkEquipment) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *NetworkEquipment) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *NetworkEquipment) HasId() bool`

HasId returns a boolean if a field has been set.

### GetLocation

`func (o *NetworkEquipment) GetLocation() Location`

GetLocation returns the Location field if non-nil, zero value otherwise.

### GetLocationOk

`func (o *NetworkEquipment) GetLocationOk() (*Location, bool)`

GetLocationOk returns a tuple with the Location field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLocation

`func (o *NetworkEquipment) SetLocation(v Location)`

SetLocation sets Location field to given value.

### HasLocation

`func (o *NetworkEquipment) HasLocation() bool`

HasLocation returns a boolean if a field has been set.

### GetName

`func (o *NetworkEquipment) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *NetworkEquipment) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *NetworkEquipment) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *NetworkEquipment) HasName() bool`

HasName returns a boolean if a field has been set.

### GetType

`func (o *NetworkEquipment) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *NetworkEquipment) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *NetworkEquipment) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *NetworkEquipment) HasType() bool`

HasType returns a boolean if a field has been set.

### GetNetworkInterfaces

`func (o *NetworkEquipment) GetNetworkInterfaces() NetworkInterfaces`

GetNetworkInterfaces returns the NetworkInterfaces field if non-nil, zero value otherwise.

### GetNetworkInterfacesOk

`func (o *NetworkEquipment) GetNetworkInterfacesOk() (*NetworkInterfaces, bool)`

GetNetworkInterfacesOk returns a tuple with the NetworkInterfaces field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNetworkInterfaces

`func (o *NetworkEquipment) SetNetworkInterfaces(v NetworkInterfaces)`

SetNetworkInterfaces sets NetworkInterfaces field to given value.

### HasNetworkInterfaces

`func (o *NetworkEquipment) HasNetworkInterfaces() bool`

HasNetworkInterfaces returns a boolean if a field has been set.

### GetPowerPorts

`func (o *NetworkEquipment) GetPowerPorts() []Port`

GetPowerPorts returns the PowerPorts field if non-nil, zero value otherwise.

### GetPowerPortsOk

`func (o *NetworkEquipment) GetPowerPortsOk() (*[]Port, bool)`

GetPowerPortsOk returns a tuple with the PowerPorts field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPowerPorts

`func (o *NetworkEquipment) SetPowerPorts(v []Port)`

SetPowerPorts sets PowerPorts field to given value.

### HasPowerPorts

`func (o *NetworkEquipment) HasPowerPorts() bool`

HasPowerPorts returns a boolean if a field has been set.

### GetRack

`func (o *NetworkEquipment) GetRack() Rack`

GetRack returns the Rack field if non-nil, zero value otherwise.

### GetRackOk

`func (o *NetworkEquipment) GetRackOk() (*Rack, bool)`

GetRackOk returns a tuple with the Rack field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRack

`func (o *NetworkEquipment) SetRack(v Rack)`

SetRack sets Rack field to given value.

### HasRack

`func (o *NetworkEquipment) HasRack() bool`

HasRack returns a boolean if a field has been set.

### GetSerialNumber

`func (o *NetworkEquipment) GetSerialNumber() string`

GetSerialNumber returns the SerialNumber field if non-nil, zero value otherwise.

### GetSerialNumberOk

`func (o *NetworkEquipment) GetSerialNumberOk() (*string, bool)`

GetSerialNumberOk returns a tuple with the SerialNumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSerialNumber

`func (o *NetworkEquipment) SetSerialNumber(v string)`

SetSerialNumber sets SerialNumber field to given value.

### HasSerialNumber

`func (o *NetworkEquipment) HasSerialNumber() bool`

HasSerialNumber returns a boolean if a field has been set.

### GetSpecs

`func (o *NetworkEquipment) GetSpecs() NetworkEquipmentSpecs`

GetSpecs returns the Specs field if non-nil, zero value otherwise.

### GetSpecsOk

`func (o *NetworkEquipment) GetSpecsOk() (*NetworkEquipmentSpecs, bool)`

GetSpecsOk returns a tuple with the Specs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSpecs

`func (o *NetworkEquipment) SetSpecs(v NetworkEquipmentSpecs)`

SetSpecs sets Specs field to given value.

### HasSpecs

`func (o *NetworkEquipment) HasSpecs() bool`

HasSpecs returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


