# Contract

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **string** |  | [optional] 
**CustomerId** | Pointer to **string** |  | [optional] 
**SalesOrgId** | Pointer to **string** |  | [optional] 
**DeliveryStatus** | Pointer to **string** |  | [optional] 
**Reference** | Pointer to **NullableString** |  | [optional] 
**PrivateNetworkPortSpeed** | Pointer to **NullableString** |  | [optional] 
**Subnets** | Pointer to [**[]Subnet**](Subnet.md) |  | [optional] 
**Status** | Pointer to **string** |  | [optional] 
**StartsAt** | Pointer to **time.Time** |  | [optional] 
**EndsAt** | Pointer to **NullableTime** |  | [optional] 
**Sla** | Pointer to **string** | Service Level Agreement | [optional] 
**ContractTerm** | Pointer to **int32** |  | [optional] 
**ContractType** | Pointer to **string** |  | [optional] 
**BillingCycle** | Pointer to **int32** | Applied billing cycle | [optional] 
**BillingFrequency** | Pointer to **string** | The interval for which billing will be done | [optional] 
**PricePerFrequency** | Pointer to **string** | Price per billing frequency | [optional] 
**Currency** | Pointer to **string** |  | [optional] 
**NetworkTraffic** | Pointer to [**NetworkTraffic**](NetworkTraffic.md) |  | [optional] 
**SoftwareLicenses** | Pointer to [**[]SoftwareLicense**](SoftwareLicense.md) |  | [optional] 
**ManagedServices** | Pointer to **[]string** |  | [optional] 
**AggregationPackId** | Pointer to **NullableString** |  | [optional] 
**Ipv4Quantity** | Pointer to **NullableInt32** |  | [optional] 

## Methods

### NewContract

`func NewContract() *Contract`

NewContract instantiates a new Contract object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewContractWithDefaults

`func NewContractWithDefaults() *Contract`

NewContractWithDefaults instantiates a new Contract object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *Contract) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *Contract) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *Contract) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *Contract) HasId() bool`

HasId returns a boolean if a field has been set.

### GetCustomerId

`func (o *Contract) GetCustomerId() string`

GetCustomerId returns the CustomerId field if non-nil, zero value otherwise.

### GetCustomerIdOk

`func (o *Contract) GetCustomerIdOk() (*string, bool)`

GetCustomerIdOk returns a tuple with the CustomerId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomerId

`func (o *Contract) SetCustomerId(v string)`

SetCustomerId sets CustomerId field to given value.

### HasCustomerId

`func (o *Contract) HasCustomerId() bool`

HasCustomerId returns a boolean if a field has been set.

### GetSalesOrgId

`func (o *Contract) GetSalesOrgId() string`

GetSalesOrgId returns the SalesOrgId field if non-nil, zero value otherwise.

### GetSalesOrgIdOk

`func (o *Contract) GetSalesOrgIdOk() (*string, bool)`

GetSalesOrgIdOk returns a tuple with the SalesOrgId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSalesOrgId

`func (o *Contract) SetSalesOrgId(v string)`

SetSalesOrgId sets SalesOrgId field to given value.

### HasSalesOrgId

`func (o *Contract) HasSalesOrgId() bool`

HasSalesOrgId returns a boolean if a field has been set.

### GetDeliveryStatus

`func (o *Contract) GetDeliveryStatus() string`

GetDeliveryStatus returns the DeliveryStatus field if non-nil, zero value otherwise.

### GetDeliveryStatusOk

`func (o *Contract) GetDeliveryStatusOk() (*string, bool)`

GetDeliveryStatusOk returns a tuple with the DeliveryStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeliveryStatus

`func (o *Contract) SetDeliveryStatus(v string)`

SetDeliveryStatus sets DeliveryStatus field to given value.

### HasDeliveryStatus

`func (o *Contract) HasDeliveryStatus() bool`

HasDeliveryStatus returns a boolean if a field has been set.

### GetReference

`func (o *Contract) GetReference() string`

GetReference returns the Reference field if non-nil, zero value otherwise.

### GetReferenceOk

`func (o *Contract) GetReferenceOk() (*string, bool)`

GetReferenceOk returns a tuple with the Reference field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReference

`func (o *Contract) SetReference(v string)`

SetReference sets Reference field to given value.

### HasReference

`func (o *Contract) HasReference() bool`

HasReference returns a boolean if a field has been set.

### SetReferenceNil

`func (o *Contract) SetReferenceNil(b bool)`

 SetReferenceNil sets the value for Reference to be an explicit nil

### UnsetReference
`func (o *Contract) UnsetReference()`

UnsetReference ensures that no value is present for Reference, not even an explicit nil
### GetPrivateNetworkPortSpeed

`func (o *Contract) GetPrivateNetworkPortSpeed() string`

GetPrivateNetworkPortSpeed returns the PrivateNetworkPortSpeed field if non-nil, zero value otherwise.

### GetPrivateNetworkPortSpeedOk

`func (o *Contract) GetPrivateNetworkPortSpeedOk() (*string, bool)`

GetPrivateNetworkPortSpeedOk returns a tuple with the PrivateNetworkPortSpeed field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrivateNetworkPortSpeed

`func (o *Contract) SetPrivateNetworkPortSpeed(v string)`

SetPrivateNetworkPortSpeed sets PrivateNetworkPortSpeed field to given value.

### HasPrivateNetworkPortSpeed

`func (o *Contract) HasPrivateNetworkPortSpeed() bool`

HasPrivateNetworkPortSpeed returns a boolean if a field has been set.

### SetPrivateNetworkPortSpeedNil

`func (o *Contract) SetPrivateNetworkPortSpeedNil(b bool)`

 SetPrivateNetworkPortSpeedNil sets the value for PrivateNetworkPortSpeed to be an explicit nil

### UnsetPrivateNetworkPortSpeed
`func (o *Contract) UnsetPrivateNetworkPortSpeed()`

UnsetPrivateNetworkPortSpeed ensures that no value is present for PrivateNetworkPortSpeed, not even an explicit nil
### GetSubnets

`func (o *Contract) GetSubnets() []Subnet`

GetSubnets returns the Subnets field if non-nil, zero value otherwise.

### GetSubnetsOk

`func (o *Contract) GetSubnetsOk() (*[]Subnet, bool)`

GetSubnetsOk returns a tuple with the Subnets field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubnets

`func (o *Contract) SetSubnets(v []Subnet)`

SetSubnets sets Subnets field to given value.

### HasSubnets

`func (o *Contract) HasSubnets() bool`

HasSubnets returns a boolean if a field has been set.

### GetStatus

`func (o *Contract) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *Contract) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *Contract) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *Contract) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetStartsAt

`func (o *Contract) GetStartsAt() time.Time`

GetStartsAt returns the StartsAt field if non-nil, zero value otherwise.

### GetStartsAtOk

`func (o *Contract) GetStartsAtOk() (*time.Time, bool)`

GetStartsAtOk returns a tuple with the StartsAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStartsAt

`func (o *Contract) SetStartsAt(v time.Time)`

SetStartsAt sets StartsAt field to given value.

### HasStartsAt

`func (o *Contract) HasStartsAt() bool`

HasStartsAt returns a boolean if a field has been set.

### GetEndsAt

`func (o *Contract) GetEndsAt() time.Time`

GetEndsAt returns the EndsAt field if non-nil, zero value otherwise.

### GetEndsAtOk

`func (o *Contract) GetEndsAtOk() (*time.Time, bool)`

GetEndsAtOk returns a tuple with the EndsAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEndsAt

`func (o *Contract) SetEndsAt(v time.Time)`

SetEndsAt sets EndsAt field to given value.

### HasEndsAt

`func (o *Contract) HasEndsAt() bool`

HasEndsAt returns a boolean if a field has been set.

### SetEndsAtNil

`func (o *Contract) SetEndsAtNil(b bool)`

 SetEndsAtNil sets the value for EndsAt to be an explicit nil

### UnsetEndsAt
`func (o *Contract) UnsetEndsAt()`

UnsetEndsAt ensures that no value is present for EndsAt, not even an explicit nil
### GetSla

`func (o *Contract) GetSla() string`

GetSla returns the Sla field if non-nil, zero value otherwise.

### GetSlaOk

`func (o *Contract) GetSlaOk() (*string, bool)`

GetSlaOk returns a tuple with the Sla field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSla

`func (o *Contract) SetSla(v string)`

SetSla sets Sla field to given value.

### HasSla

`func (o *Contract) HasSla() bool`

HasSla returns a boolean if a field has been set.

### GetContractTerm

`func (o *Contract) GetContractTerm() int32`

GetContractTerm returns the ContractTerm field if non-nil, zero value otherwise.

### GetContractTermOk

`func (o *Contract) GetContractTermOk() (*int32, bool)`

GetContractTermOk returns a tuple with the ContractTerm field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContractTerm

`func (o *Contract) SetContractTerm(v int32)`

SetContractTerm sets ContractTerm field to given value.

### HasContractTerm

`func (o *Contract) HasContractTerm() bool`

HasContractTerm returns a boolean if a field has been set.

### GetContractType

`func (o *Contract) GetContractType() string`

GetContractType returns the ContractType field if non-nil, zero value otherwise.

### GetContractTypeOk

`func (o *Contract) GetContractTypeOk() (*string, bool)`

GetContractTypeOk returns a tuple with the ContractType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContractType

`func (o *Contract) SetContractType(v string)`

SetContractType sets ContractType field to given value.

### HasContractType

`func (o *Contract) HasContractType() bool`

HasContractType returns a boolean if a field has been set.

### GetBillingCycle

`func (o *Contract) GetBillingCycle() int32`

GetBillingCycle returns the BillingCycle field if non-nil, zero value otherwise.

### GetBillingCycleOk

`func (o *Contract) GetBillingCycleOk() (*int32, bool)`

GetBillingCycleOk returns a tuple with the BillingCycle field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBillingCycle

`func (o *Contract) SetBillingCycle(v int32)`

SetBillingCycle sets BillingCycle field to given value.

### HasBillingCycle

`func (o *Contract) HasBillingCycle() bool`

HasBillingCycle returns a boolean if a field has been set.

### GetBillingFrequency

`func (o *Contract) GetBillingFrequency() string`

GetBillingFrequency returns the BillingFrequency field if non-nil, zero value otherwise.

### GetBillingFrequencyOk

`func (o *Contract) GetBillingFrequencyOk() (*string, bool)`

GetBillingFrequencyOk returns a tuple with the BillingFrequency field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBillingFrequency

`func (o *Contract) SetBillingFrequency(v string)`

SetBillingFrequency sets BillingFrequency field to given value.

### HasBillingFrequency

`func (o *Contract) HasBillingFrequency() bool`

HasBillingFrequency returns a boolean if a field has been set.

### GetPricePerFrequency

`func (o *Contract) GetPricePerFrequency() string`

GetPricePerFrequency returns the PricePerFrequency field if non-nil, zero value otherwise.

### GetPricePerFrequencyOk

`func (o *Contract) GetPricePerFrequencyOk() (*string, bool)`

GetPricePerFrequencyOk returns a tuple with the PricePerFrequency field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPricePerFrequency

`func (o *Contract) SetPricePerFrequency(v string)`

SetPricePerFrequency sets PricePerFrequency field to given value.

### HasPricePerFrequency

`func (o *Contract) HasPricePerFrequency() bool`

HasPricePerFrequency returns a boolean if a field has been set.

### GetCurrency

`func (o *Contract) GetCurrency() string`

GetCurrency returns the Currency field if non-nil, zero value otherwise.

### GetCurrencyOk

`func (o *Contract) GetCurrencyOk() (*string, bool)`

GetCurrencyOk returns a tuple with the Currency field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCurrency

`func (o *Contract) SetCurrency(v string)`

SetCurrency sets Currency field to given value.

### HasCurrency

`func (o *Contract) HasCurrency() bool`

HasCurrency returns a boolean if a field has been set.

### GetNetworkTraffic

`func (o *Contract) GetNetworkTraffic() NetworkTraffic`

GetNetworkTraffic returns the NetworkTraffic field if non-nil, zero value otherwise.

### GetNetworkTrafficOk

`func (o *Contract) GetNetworkTrafficOk() (*NetworkTraffic, bool)`

GetNetworkTrafficOk returns a tuple with the NetworkTraffic field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNetworkTraffic

`func (o *Contract) SetNetworkTraffic(v NetworkTraffic)`

SetNetworkTraffic sets NetworkTraffic field to given value.

### HasNetworkTraffic

`func (o *Contract) HasNetworkTraffic() bool`

HasNetworkTraffic returns a boolean if a field has been set.

### GetSoftwareLicenses

`func (o *Contract) GetSoftwareLicenses() []SoftwareLicense`

GetSoftwareLicenses returns the SoftwareLicenses field if non-nil, zero value otherwise.

### GetSoftwareLicensesOk

`func (o *Contract) GetSoftwareLicensesOk() (*[]SoftwareLicense, bool)`

GetSoftwareLicensesOk returns a tuple with the SoftwareLicenses field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSoftwareLicenses

`func (o *Contract) SetSoftwareLicenses(v []SoftwareLicense)`

SetSoftwareLicenses sets SoftwareLicenses field to given value.

### HasSoftwareLicenses

`func (o *Contract) HasSoftwareLicenses() bool`

HasSoftwareLicenses returns a boolean if a field has been set.

### GetManagedServices

`func (o *Contract) GetManagedServices() []string`

GetManagedServices returns the ManagedServices field if non-nil, zero value otherwise.

### GetManagedServicesOk

`func (o *Contract) GetManagedServicesOk() (*[]string, bool)`

GetManagedServicesOk returns a tuple with the ManagedServices field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetManagedServices

`func (o *Contract) SetManagedServices(v []string)`

SetManagedServices sets ManagedServices field to given value.

### HasManagedServices

`func (o *Contract) HasManagedServices() bool`

HasManagedServices returns a boolean if a field has been set.

### GetAggregationPackId

`func (o *Contract) GetAggregationPackId() string`

GetAggregationPackId returns the AggregationPackId field if non-nil, zero value otherwise.

### GetAggregationPackIdOk

`func (o *Contract) GetAggregationPackIdOk() (*string, bool)`

GetAggregationPackIdOk returns a tuple with the AggregationPackId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAggregationPackId

`func (o *Contract) SetAggregationPackId(v string)`

SetAggregationPackId sets AggregationPackId field to given value.

### HasAggregationPackId

`func (o *Contract) HasAggregationPackId() bool`

HasAggregationPackId returns a boolean if a field has been set.

### SetAggregationPackIdNil

`func (o *Contract) SetAggregationPackIdNil(b bool)`

 SetAggregationPackIdNil sets the value for AggregationPackId to be an explicit nil

### UnsetAggregationPackId
`func (o *Contract) UnsetAggregationPackId()`

UnsetAggregationPackId ensures that no value is present for AggregationPackId, not even an explicit nil
### GetIpv4Quantity

`func (o *Contract) GetIpv4Quantity() int32`

GetIpv4Quantity returns the Ipv4Quantity field if non-nil, zero value otherwise.

### GetIpv4QuantityOk

`func (o *Contract) GetIpv4QuantityOk() (*int32, bool)`

GetIpv4QuantityOk returns a tuple with the Ipv4Quantity field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIpv4Quantity

`func (o *Contract) SetIpv4Quantity(v int32)`

SetIpv4Quantity sets Ipv4Quantity field to given value.

### HasIpv4Quantity

`func (o *Contract) HasIpv4Quantity() bool`

HasIpv4Quantity returns a boolean if a field has been set.

### SetIpv4QuantityNil

`func (o *Contract) SetIpv4QuantityNil(b bool)`

 SetIpv4QuantityNil sets the value for Ipv4Quantity to be an explicit nil

### UnsetIpv4Quantity
`func (o *Contract) UnsetIpv4Quantity()`

UnsetIpv4Quantity ensures that no value is present for Ipv4Quantity, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


