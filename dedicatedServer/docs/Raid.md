# Raid

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Level** | Pointer to [**RaidLevel**](RaidLevel.md) |  | [optional] 
**NumberOfDisks** | Pointer to **int32** | The number of disks you want to apply RAID on. If not specified all disks are used | [optional] 
**Type** | Pointer to [**RaidType**](RaidType.md) |  | [optional] 

## Methods

### NewRaid

`func NewRaid() *Raid`

NewRaid instantiates a new Raid object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRaidWithDefaults

`func NewRaidWithDefaults() *Raid`

NewRaidWithDefaults instantiates a new Raid object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetLevel

`func (o *Raid) GetLevel() RaidLevel`

GetLevel returns the Level field if non-nil, zero value otherwise.

### GetLevelOk

`func (o *Raid) GetLevelOk() (*RaidLevel, bool)`

GetLevelOk returns a tuple with the Level field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLevel

`func (o *Raid) SetLevel(v RaidLevel)`

SetLevel sets Level field to given value.

### HasLevel

`func (o *Raid) HasLevel() bool`

HasLevel returns a boolean if a field has been set.

### GetNumberOfDisks

`func (o *Raid) GetNumberOfDisks() int32`

GetNumberOfDisks returns the NumberOfDisks field if non-nil, zero value otherwise.

### GetNumberOfDisksOk

`func (o *Raid) GetNumberOfDisksOk() (*int32, bool)`

GetNumberOfDisksOk returns a tuple with the NumberOfDisks field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNumberOfDisks

`func (o *Raid) SetNumberOfDisks(v int32)`

SetNumberOfDisks sets NumberOfDisks field to given value.

### HasNumberOfDisks

`func (o *Raid) HasNumberOfDisks() bool`

HasNumberOfDisks returns a boolean if a field has been set.

### GetType

`func (o *Raid) GetType() RaidType`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *Raid) GetTypeOk() (*RaidType, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *Raid) SetType(v RaidType)`

SetType sets Type field to given value.

### HasType

`func (o *Raid) HasType() bool`

HasType returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


