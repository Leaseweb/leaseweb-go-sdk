# GetOperatingSystemResult

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Architecture** | Pointer to **string** | The architecture of the operating system | [optional] 
**Configurable** | Pointer to **bool** | If the default options are configurable or not | [optional] 
**Defaults** | Pointer to [**Defaults**](Defaults.md) |  | [optional] 
**Family** | Pointer to **string** | The operating system family | [optional] 
**Id** | Pointer to **string** | The operating system ID | [optional] 
**Name** | Pointer to **string** | A human readable name for the operating system | [optional] 
**Type** | Pointer to **string** | The type of operating system | [optional] 
**Version** | Pointer to **string** | The version of the operating system | [optional] 
**Features** | Pointer to **[]string** | Operating system features | [optional] 
**SupportedFileSystems** | Pointer to **[]string** | Operating system supported file systems | [optional] 
**SupportedBootDevices** | Pointer to **[]string** | Operating system supported boot devices | [optional] 

## Methods

### NewGetOperatingSystemResult

`func NewGetOperatingSystemResult() *GetOperatingSystemResult`

NewGetOperatingSystemResult instantiates a new GetOperatingSystemResult object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetOperatingSystemResultWithDefaults

`func NewGetOperatingSystemResultWithDefaults() *GetOperatingSystemResult`

NewGetOperatingSystemResultWithDefaults instantiates a new GetOperatingSystemResult object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetArchitecture

`func (o *GetOperatingSystemResult) GetArchitecture() string`

GetArchitecture returns the Architecture field if non-nil, zero value otherwise.

### GetArchitectureOk

`func (o *GetOperatingSystemResult) GetArchitectureOk() (*string, bool)`

GetArchitectureOk returns a tuple with the Architecture field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetArchitecture

`func (o *GetOperatingSystemResult) SetArchitecture(v string)`

SetArchitecture sets Architecture field to given value.

### HasArchitecture

`func (o *GetOperatingSystemResult) HasArchitecture() bool`

HasArchitecture returns a boolean if a field has been set.

### GetConfigurable

`func (o *GetOperatingSystemResult) GetConfigurable() bool`

GetConfigurable returns the Configurable field if non-nil, zero value otherwise.

### GetConfigurableOk

`func (o *GetOperatingSystemResult) GetConfigurableOk() (*bool, bool)`

GetConfigurableOk returns a tuple with the Configurable field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConfigurable

`func (o *GetOperatingSystemResult) SetConfigurable(v bool)`

SetConfigurable sets Configurable field to given value.

### HasConfigurable

`func (o *GetOperatingSystemResult) HasConfigurable() bool`

HasConfigurable returns a boolean if a field has been set.

### GetDefaults

`func (o *GetOperatingSystemResult) GetDefaults() Defaults`

GetDefaults returns the Defaults field if non-nil, zero value otherwise.

### GetDefaultsOk

`func (o *GetOperatingSystemResult) GetDefaultsOk() (*Defaults, bool)`

GetDefaultsOk returns a tuple with the Defaults field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDefaults

`func (o *GetOperatingSystemResult) SetDefaults(v Defaults)`

SetDefaults sets Defaults field to given value.

### HasDefaults

`func (o *GetOperatingSystemResult) HasDefaults() bool`

HasDefaults returns a boolean if a field has been set.

### GetFamily

`func (o *GetOperatingSystemResult) GetFamily() string`

GetFamily returns the Family field if non-nil, zero value otherwise.

### GetFamilyOk

`func (o *GetOperatingSystemResult) GetFamilyOk() (*string, bool)`

GetFamilyOk returns a tuple with the Family field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFamily

`func (o *GetOperatingSystemResult) SetFamily(v string)`

SetFamily sets Family field to given value.

### HasFamily

`func (o *GetOperatingSystemResult) HasFamily() bool`

HasFamily returns a boolean if a field has been set.

### GetId

`func (o *GetOperatingSystemResult) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *GetOperatingSystemResult) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *GetOperatingSystemResult) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *GetOperatingSystemResult) HasId() bool`

HasId returns a boolean if a field has been set.

### GetName

`func (o *GetOperatingSystemResult) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *GetOperatingSystemResult) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *GetOperatingSystemResult) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *GetOperatingSystemResult) HasName() bool`

HasName returns a boolean if a field has been set.

### GetType

`func (o *GetOperatingSystemResult) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *GetOperatingSystemResult) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *GetOperatingSystemResult) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *GetOperatingSystemResult) HasType() bool`

HasType returns a boolean if a field has been set.

### GetVersion

`func (o *GetOperatingSystemResult) GetVersion() string`

GetVersion returns the Version field if non-nil, zero value otherwise.

### GetVersionOk

`func (o *GetOperatingSystemResult) GetVersionOk() (*string, bool)`

GetVersionOk returns a tuple with the Version field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVersion

`func (o *GetOperatingSystemResult) SetVersion(v string)`

SetVersion sets Version field to given value.

### HasVersion

`func (o *GetOperatingSystemResult) HasVersion() bool`

HasVersion returns a boolean if a field has been set.

### GetFeatures

`func (o *GetOperatingSystemResult) GetFeatures() []string`

GetFeatures returns the Features field if non-nil, zero value otherwise.

### GetFeaturesOk

`func (o *GetOperatingSystemResult) GetFeaturesOk() (*[]string, bool)`

GetFeaturesOk returns a tuple with the Features field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFeatures

`func (o *GetOperatingSystemResult) SetFeatures(v []string)`

SetFeatures sets Features field to given value.

### HasFeatures

`func (o *GetOperatingSystemResult) HasFeatures() bool`

HasFeatures returns a boolean if a field has been set.

### GetSupportedFileSystems

`func (o *GetOperatingSystemResult) GetSupportedFileSystems() []string`

GetSupportedFileSystems returns the SupportedFileSystems field if non-nil, zero value otherwise.

### GetSupportedFileSystemsOk

`func (o *GetOperatingSystemResult) GetSupportedFileSystemsOk() (*[]string, bool)`

GetSupportedFileSystemsOk returns a tuple with the SupportedFileSystems field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSupportedFileSystems

`func (o *GetOperatingSystemResult) SetSupportedFileSystems(v []string)`

SetSupportedFileSystems sets SupportedFileSystems field to given value.

### HasSupportedFileSystems

`func (o *GetOperatingSystemResult) HasSupportedFileSystems() bool`

HasSupportedFileSystems returns a boolean if a field has been set.

### GetSupportedBootDevices

`func (o *GetOperatingSystemResult) GetSupportedBootDevices() []string`

GetSupportedBootDevices returns the SupportedBootDevices field if non-nil, zero value otherwise.

### GetSupportedBootDevicesOk

`func (o *GetOperatingSystemResult) GetSupportedBootDevicesOk() (*[]string, bool)`

GetSupportedBootDevicesOk returns a tuple with the SupportedBootDevices field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSupportedBootDevices

`func (o *GetOperatingSystemResult) SetSupportedBootDevices(v []string)`

SetSupportedBootDevices sets SupportedBootDevices field to given value.

### HasSupportedBootDevices

`func (o *GetOperatingSystemResult) HasSupportedBootDevices() bool`

HasSupportedBootDevices returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


