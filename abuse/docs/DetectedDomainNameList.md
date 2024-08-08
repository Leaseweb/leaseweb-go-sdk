# DetectedDomainNameList

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | Pointer to **string** | The domain. | [optional] 
**IpAddresses** | Pointer to **[]string** | The IP-addresses the domain resolves to. | [optional] 

## Methods

### NewDetectedDomainNameList

`func NewDetectedDomainNameList() *DetectedDomainNameList`

NewDetectedDomainNameList instantiates a new DetectedDomainNameList object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDetectedDomainNameListWithDefaults

`func NewDetectedDomainNameListWithDefaults() *DetectedDomainNameList`

NewDetectedDomainNameListWithDefaults instantiates a new DetectedDomainNameList object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *DetectedDomainNameList) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *DetectedDomainNameList) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *DetectedDomainNameList) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *DetectedDomainNameList) HasName() bool`

HasName returns a boolean if a field has been set.

### GetIpAddresses

`func (o *DetectedDomainNameList) GetIpAddresses() []string`

GetIpAddresses returns the IpAddresses field if non-nil, zero value otherwise.

### GetIpAddressesOk

`func (o *DetectedDomainNameList) GetIpAddressesOk() (*[]string, bool)`

GetIpAddressesOk returns a tuple with the IpAddresses field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIpAddresses

`func (o *DetectedDomainNameList) SetIpAddresses(v []string)`

SetIpAddresses sets IpAddresses field to given value.

### HasIpAddresses

`func (o *DetectedDomainNameList) HasIpAddresses() bool`

HasIpAddresses returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


