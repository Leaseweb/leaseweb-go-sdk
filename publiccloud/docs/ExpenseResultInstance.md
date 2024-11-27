# ExpenseResultInstance

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **string** | The unique identifier for the instance. | [optional] 
**Reference** | Pointer to **string** | The reference of the instance. | [optional] 
**Resources** | Pointer to [**Resources**](Resources.md) |  | [optional] 
**Contract** | Pointer to [**Contract**](Contract.md) |  | [optional] 
**StartedAt** | Pointer to **time.Time** | Date when the instance was started | [optional] 
**EndedAt** | Pointer to **time.Time** | Date when the instance ended | [optional] 
**RootDiskSize** | Pointer to **int32** | The root disk&#39;s size in GB. Must be at least 5 GB for Linux and FreeBSD instances and 50 GB for Windows instances | [optional] 
**RootDiskStorageType** | Pointer to [**StorageType**](StorageType.md) |  | [optional] 
**BillingType** | Pointer to **string** | The billing type of the instance. PREPAID is used for monthly commited instances, POSTPAID for hourly instances. | [optional] 
**Hours** | Pointer to **int32** | The number of hours the instance has been running. | [optional] 
**From** | Pointer to **time.Time** | The start date of the billing period. | [optional] 
**To** | Pointer to **time.Time** | The end date of the billing period. | [optional] 
**Price** | Pointer to **string** | The price of the instance for the billing period. | [optional] 

## Methods

### NewExpenseResultInstance

`func NewExpenseResultInstance() *ExpenseResultInstance`

NewExpenseResultInstance instantiates a new ExpenseResultInstance object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewExpenseResultInstanceWithDefaults

`func NewExpenseResultInstanceWithDefaults() *ExpenseResultInstance`

NewExpenseResultInstanceWithDefaults instantiates a new ExpenseResultInstance object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *ExpenseResultInstance) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *ExpenseResultInstance) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *ExpenseResultInstance) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *ExpenseResultInstance) HasId() bool`

HasId returns a boolean if a field has been set.

### GetReference

`func (o *ExpenseResultInstance) GetReference() string`

GetReference returns the Reference field if non-nil, zero value otherwise.

### GetReferenceOk

`func (o *ExpenseResultInstance) GetReferenceOk() (*string, bool)`

GetReferenceOk returns a tuple with the Reference field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReference

`func (o *ExpenseResultInstance) SetReference(v string)`

SetReference sets Reference field to given value.

### HasReference

`func (o *ExpenseResultInstance) HasReference() bool`

HasReference returns a boolean if a field has been set.

### GetResources

`func (o *ExpenseResultInstance) GetResources() Resources`

GetResources returns the Resources field if non-nil, zero value otherwise.

### GetResourcesOk

`func (o *ExpenseResultInstance) GetResourcesOk() (*Resources, bool)`

GetResourcesOk returns a tuple with the Resources field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResources

`func (o *ExpenseResultInstance) SetResources(v Resources)`

SetResources sets Resources field to given value.

### HasResources

`func (o *ExpenseResultInstance) HasResources() bool`

HasResources returns a boolean if a field has been set.

### GetContract

`func (o *ExpenseResultInstance) GetContract() Contract`

GetContract returns the Contract field if non-nil, zero value otherwise.

### GetContractOk

`func (o *ExpenseResultInstance) GetContractOk() (*Contract, bool)`

GetContractOk returns a tuple with the Contract field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContract

`func (o *ExpenseResultInstance) SetContract(v Contract)`

SetContract sets Contract field to given value.

### HasContract

`func (o *ExpenseResultInstance) HasContract() bool`

HasContract returns a boolean if a field has been set.

### GetStartedAt

`func (o *ExpenseResultInstance) GetStartedAt() time.Time`

GetStartedAt returns the StartedAt field if non-nil, zero value otherwise.

### GetStartedAtOk

`func (o *ExpenseResultInstance) GetStartedAtOk() (*time.Time, bool)`

GetStartedAtOk returns a tuple with the StartedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStartedAt

`func (o *ExpenseResultInstance) SetStartedAt(v time.Time)`

SetStartedAt sets StartedAt field to given value.

### HasStartedAt

`func (o *ExpenseResultInstance) HasStartedAt() bool`

HasStartedAt returns a boolean if a field has been set.

### GetEndedAt

`func (o *ExpenseResultInstance) GetEndedAt() time.Time`

GetEndedAt returns the EndedAt field if non-nil, zero value otherwise.

### GetEndedAtOk

`func (o *ExpenseResultInstance) GetEndedAtOk() (*time.Time, bool)`

GetEndedAtOk returns a tuple with the EndedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEndedAt

`func (o *ExpenseResultInstance) SetEndedAt(v time.Time)`

SetEndedAt sets EndedAt field to given value.

### HasEndedAt

`func (o *ExpenseResultInstance) HasEndedAt() bool`

HasEndedAt returns a boolean if a field has been set.

### GetRootDiskSize

`func (o *ExpenseResultInstance) GetRootDiskSize() int32`

GetRootDiskSize returns the RootDiskSize field if non-nil, zero value otherwise.

### GetRootDiskSizeOk

`func (o *ExpenseResultInstance) GetRootDiskSizeOk() (*int32, bool)`

GetRootDiskSizeOk returns a tuple with the RootDiskSize field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRootDiskSize

`func (o *ExpenseResultInstance) SetRootDiskSize(v int32)`

SetRootDiskSize sets RootDiskSize field to given value.

### HasRootDiskSize

`func (o *ExpenseResultInstance) HasRootDiskSize() bool`

HasRootDiskSize returns a boolean if a field has been set.

### GetRootDiskStorageType

`func (o *ExpenseResultInstance) GetRootDiskStorageType() StorageType`

GetRootDiskStorageType returns the RootDiskStorageType field if non-nil, zero value otherwise.

### GetRootDiskStorageTypeOk

`func (o *ExpenseResultInstance) GetRootDiskStorageTypeOk() (*StorageType, bool)`

GetRootDiskStorageTypeOk returns a tuple with the RootDiskStorageType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRootDiskStorageType

`func (o *ExpenseResultInstance) SetRootDiskStorageType(v StorageType)`

SetRootDiskStorageType sets RootDiskStorageType field to given value.

### HasRootDiskStorageType

`func (o *ExpenseResultInstance) HasRootDiskStorageType() bool`

HasRootDiskStorageType returns a boolean if a field has been set.

### GetBillingType

`func (o *ExpenseResultInstance) GetBillingType() string`

GetBillingType returns the BillingType field if non-nil, zero value otherwise.

### GetBillingTypeOk

`func (o *ExpenseResultInstance) GetBillingTypeOk() (*string, bool)`

GetBillingTypeOk returns a tuple with the BillingType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBillingType

`func (o *ExpenseResultInstance) SetBillingType(v string)`

SetBillingType sets BillingType field to given value.

### HasBillingType

`func (o *ExpenseResultInstance) HasBillingType() bool`

HasBillingType returns a boolean if a field has been set.

### GetHours

`func (o *ExpenseResultInstance) GetHours() int32`

GetHours returns the Hours field if non-nil, zero value otherwise.

### GetHoursOk

`func (o *ExpenseResultInstance) GetHoursOk() (*int32, bool)`

GetHoursOk returns a tuple with the Hours field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHours

`func (o *ExpenseResultInstance) SetHours(v int32)`

SetHours sets Hours field to given value.

### HasHours

`func (o *ExpenseResultInstance) HasHours() bool`

HasHours returns a boolean if a field has been set.

### GetFrom

`func (o *ExpenseResultInstance) GetFrom() time.Time`

GetFrom returns the From field if non-nil, zero value otherwise.

### GetFromOk

`func (o *ExpenseResultInstance) GetFromOk() (*time.Time, bool)`

GetFromOk returns a tuple with the From field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFrom

`func (o *ExpenseResultInstance) SetFrom(v time.Time)`

SetFrom sets From field to given value.

### HasFrom

`func (o *ExpenseResultInstance) HasFrom() bool`

HasFrom returns a boolean if a field has been set.

### GetTo

`func (o *ExpenseResultInstance) GetTo() time.Time`

GetTo returns the To field if non-nil, zero value otherwise.

### GetToOk

`func (o *ExpenseResultInstance) GetToOk() (*time.Time, bool)`

GetToOk returns a tuple with the To field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTo

`func (o *ExpenseResultInstance) SetTo(v time.Time)`

SetTo sets To field to given value.

### HasTo

`func (o *ExpenseResultInstance) HasTo() bool`

HasTo returns a boolean if a field has been set.

### GetPrice

`func (o *ExpenseResultInstance) GetPrice() string`

GetPrice returns the Price field if non-nil, zero value otherwise.

### GetPriceOk

`func (o *ExpenseResultInstance) GetPriceOk() (*string, bool)`

GetPriceOk returns a tuple with the Price field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrice

`func (o *ExpenseResultInstance) SetPrice(v string)`

SetPrice sets Price field to given value.

### HasPrice

`func (o *ExpenseResultInstance) HasPrice() bool`

HasPrice returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


