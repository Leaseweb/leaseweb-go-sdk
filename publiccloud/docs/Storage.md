# Storage

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Local** | [**Price**](Price.md) |  | 
**Central** | [**Price**](Price.md) |  | 

## Methods

### NewStorage

`func NewStorage(local Price, central Price, ) *Storage`

NewStorage instantiates a new Storage object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewStorageWithDefaults

`func NewStorageWithDefaults() *Storage`

NewStorageWithDefaults instantiates a new Storage object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetLocal

`func (o *Storage) GetLocal() Price`

GetLocal returns the Local field if non-nil, zero value otherwise.

### GetLocalOk

`func (o *Storage) GetLocalOk() (*Price, bool)`

GetLocalOk returns a tuple with the Local field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLocal

`func (o *Storage) SetLocal(v Price)`

SetLocal sets Local field to given value.


### GetCentral

`func (o *Storage) GetCentral() Price`

GetCentral returns the Central field if non-nil, zero value otherwise.

### GetCentralOk

`func (o *Storage) GetCentralOk() (*Price, bool)`

GetCentralOk returns a tuple with the Central field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCentral

`func (o *Storage) SetCentral(v Price)`

SetCentral sets Central field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


