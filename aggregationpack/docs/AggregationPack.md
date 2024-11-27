# AggregationPack

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** | Id of the aggregation pack | 
**CustomerId** | **string** |  | 
**SalesOrgId** | **string** |  | 
**ContractStartDate** | **time.Time** |  | 
**BillingType** | **string** |  | 
**ConnectivityType** | **string** |  | 
**ContractTerm** | **int32** |  | 
**DataTrafficCommit** | **int32** |  | 
**DataTrafficCommitUnit** | [**DataTrafficCommitUnit**](DataTrafficCommitUnit.md) |  | 
**NetworkPerformanceType** | [**NetworkPerformanceType**](NetworkPerformanceType.md) |  | 
**AggregationType** | [**AggregationType**](AggregationType.md) |  | 

## Methods

### NewAggregationPack

`func NewAggregationPack(id string, customerId string, salesOrgId string, contractStartDate time.Time, billingType string, connectivityType string, contractTerm int32, dataTrafficCommit int32, dataTrafficCommitUnit DataTrafficCommitUnit, networkPerformanceType NetworkPerformanceType, aggregationType AggregationType, ) *AggregationPack`

NewAggregationPack instantiates a new AggregationPack object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAggregationPackWithDefaults

`func NewAggregationPackWithDefaults() *AggregationPack`

NewAggregationPackWithDefaults instantiates a new AggregationPack object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *AggregationPack) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *AggregationPack) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *AggregationPack) SetId(v string)`

SetId sets Id field to given value.


### GetCustomerId

`func (o *AggregationPack) GetCustomerId() string`

GetCustomerId returns the CustomerId field if non-nil, zero value otherwise.

### GetCustomerIdOk

`func (o *AggregationPack) GetCustomerIdOk() (*string, bool)`

GetCustomerIdOk returns a tuple with the CustomerId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomerId

`func (o *AggregationPack) SetCustomerId(v string)`

SetCustomerId sets CustomerId field to given value.


### GetSalesOrgId

`func (o *AggregationPack) GetSalesOrgId() string`

GetSalesOrgId returns the SalesOrgId field if non-nil, zero value otherwise.

### GetSalesOrgIdOk

`func (o *AggregationPack) GetSalesOrgIdOk() (*string, bool)`

GetSalesOrgIdOk returns a tuple with the SalesOrgId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSalesOrgId

`func (o *AggregationPack) SetSalesOrgId(v string)`

SetSalesOrgId sets SalesOrgId field to given value.


### GetContractStartDate

`func (o *AggregationPack) GetContractStartDate() time.Time`

GetContractStartDate returns the ContractStartDate field if non-nil, zero value otherwise.

### GetContractStartDateOk

`func (o *AggregationPack) GetContractStartDateOk() (*time.Time, bool)`

GetContractStartDateOk returns a tuple with the ContractStartDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContractStartDate

`func (o *AggregationPack) SetContractStartDate(v time.Time)`

SetContractStartDate sets ContractStartDate field to given value.


### GetBillingType

`func (o *AggregationPack) GetBillingType() string`

GetBillingType returns the BillingType field if non-nil, zero value otherwise.

### GetBillingTypeOk

`func (o *AggregationPack) GetBillingTypeOk() (*string, bool)`

GetBillingTypeOk returns a tuple with the BillingType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBillingType

`func (o *AggregationPack) SetBillingType(v string)`

SetBillingType sets BillingType field to given value.


### GetConnectivityType

`func (o *AggregationPack) GetConnectivityType() string`

GetConnectivityType returns the ConnectivityType field if non-nil, zero value otherwise.

### GetConnectivityTypeOk

`func (o *AggregationPack) GetConnectivityTypeOk() (*string, bool)`

GetConnectivityTypeOk returns a tuple with the ConnectivityType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConnectivityType

`func (o *AggregationPack) SetConnectivityType(v string)`

SetConnectivityType sets ConnectivityType field to given value.


### GetContractTerm

`func (o *AggregationPack) GetContractTerm() int32`

GetContractTerm returns the ContractTerm field if non-nil, zero value otherwise.

### GetContractTermOk

`func (o *AggregationPack) GetContractTermOk() (*int32, bool)`

GetContractTermOk returns a tuple with the ContractTerm field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContractTerm

`func (o *AggregationPack) SetContractTerm(v int32)`

SetContractTerm sets ContractTerm field to given value.


### GetDataTrafficCommit

`func (o *AggregationPack) GetDataTrafficCommit() int32`

GetDataTrafficCommit returns the DataTrafficCommit field if non-nil, zero value otherwise.

### GetDataTrafficCommitOk

`func (o *AggregationPack) GetDataTrafficCommitOk() (*int32, bool)`

GetDataTrafficCommitOk returns a tuple with the DataTrafficCommit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDataTrafficCommit

`func (o *AggregationPack) SetDataTrafficCommit(v int32)`

SetDataTrafficCommit sets DataTrafficCommit field to given value.


### GetDataTrafficCommitUnit

`func (o *AggregationPack) GetDataTrafficCommitUnit() DataTrafficCommitUnit`

GetDataTrafficCommitUnit returns the DataTrafficCommitUnit field if non-nil, zero value otherwise.

### GetDataTrafficCommitUnitOk

`func (o *AggregationPack) GetDataTrafficCommitUnitOk() (*DataTrafficCommitUnit, bool)`

GetDataTrafficCommitUnitOk returns a tuple with the DataTrafficCommitUnit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDataTrafficCommitUnit

`func (o *AggregationPack) SetDataTrafficCommitUnit(v DataTrafficCommitUnit)`

SetDataTrafficCommitUnit sets DataTrafficCommitUnit field to given value.


### GetNetworkPerformanceType

`func (o *AggregationPack) GetNetworkPerformanceType() NetworkPerformanceType`

GetNetworkPerformanceType returns the NetworkPerformanceType field if non-nil, zero value otherwise.

### GetNetworkPerformanceTypeOk

`func (o *AggregationPack) GetNetworkPerformanceTypeOk() (*NetworkPerformanceType, bool)`

GetNetworkPerformanceTypeOk returns a tuple with the NetworkPerformanceType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNetworkPerformanceType

`func (o *AggregationPack) SetNetworkPerformanceType(v NetworkPerformanceType)`

SetNetworkPerformanceType sets NetworkPerformanceType field to given value.


### GetAggregationType

`func (o *AggregationPack) GetAggregationType() AggregationType`

GetAggregationType returns the AggregationType field if non-nil, zero value otherwise.

### GetAggregationTypeOk

`func (o *AggregationPack) GetAggregationTypeOk() (*AggregationType, bool)`

GetAggregationTypeOk returns a tuple with the AggregationType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAggregationType

`func (o *AggregationPack) SetAggregationType(v AggregationType)`

SetAggregationType sets AggregationType field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


