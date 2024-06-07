# ReinstallInstanceOpts

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**OperatingSystemId** | **string** | The Operating System ID | 
**MarketAppId** | Pointer to **string** | The Market App to be installed | [optional] 

## Methods

### NewReinstallInstanceOpts

`func NewReinstallInstanceOpts(operatingSystemId string, ) *ReinstallInstanceOpts`

NewReinstallInstanceOpts instantiates a new ReinstallInstanceOpts object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewReinstallInstanceOptsWithDefaults

`func NewReinstallInstanceOptsWithDefaults() *ReinstallInstanceOpts`

NewReinstallInstanceOptsWithDefaults instantiates a new ReinstallInstanceOpts object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetOperatingSystemId

`func (o *ReinstallInstanceOpts) GetOperatingSystemId() string`

GetOperatingSystemId returns the OperatingSystemId field if non-nil, zero value otherwise.

### GetOperatingSystemIdOk

`func (o *ReinstallInstanceOpts) GetOperatingSystemIdOk() (*string, bool)`

GetOperatingSystemIdOk returns a tuple with the OperatingSystemId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOperatingSystemId

`func (o *ReinstallInstanceOpts) SetOperatingSystemId(v string)`

SetOperatingSystemId sets OperatingSystemId field to given value.


### GetMarketAppId

`func (o *ReinstallInstanceOpts) GetMarketAppId() string`

GetMarketAppId returns the MarketAppId field if non-nil, zero value otherwise.

### GetMarketAppIdOk

`func (o *ReinstallInstanceOpts) GetMarketAppIdOk() (*string, bool)`

GetMarketAppIdOk returns a tuple with the MarketAppId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMarketAppId

`func (o *ReinstallInstanceOpts) SetMarketAppId(v string)`

SetMarketAppId sets MarketAppId field to given value.

### HasMarketAppId

`func (o *ReinstallInstanceOpts) HasMarketAppId() bool`

HasMarketAppId returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


