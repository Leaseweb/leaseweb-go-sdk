# NullRoutedIP

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** | Null route ID | 
**Ip** | **string** | IP address | 
**NulledAt** | **time.Time** | Null route date in UTC in format yyyy-mm-ddThh:mm:ssZ | 
**NulledBy** | **string** | Email address of the user who created the null route or &#39;LeaseWeb&#39; if null route was created by LeaseWeb. | 
**NullLevel** | **int32** | Null route permission level. If greater than 1 then the null route can only be removed by LeaseWeb. | [default to 0]
**AutomatedUnnullingAt** | **NullableTime** | The date and time when the null route is to be automatically removed. | 
**UnnulledAt** | **NullableTime** | The date and time when the null route has been removed. If null then the null route is still active. | 
**UnnulledBy** | **NullableString** | Email address of the user who removed the null route or &#39;LeaseWeb&#39; if null route was removed by LeaseWeb. | 
**TicketId** | **NullableString** | Reference stored with the null route. | 
**Comment** | **NullableString** | Comment stored with the null route. | 
**EquipmentId** | **string** | ID of the equipment which was assigned to the IP at the time of null route creation. | 
**AssignedContract** | [**NullableAssignedContract**](AssignedContract.md) |  | 

## Methods

### NewNullRoutedIP

`func NewNullRoutedIP(id string, ip string, nulledAt time.Time, nulledBy string, nullLevel int32, automatedUnnullingAt NullableTime, unnulledAt NullableTime, unnulledBy NullableString, ticketId NullableString, comment NullableString, equipmentId string, assignedContract NullableAssignedContract, ) *NullRoutedIP`

NewNullRoutedIP instantiates a new NullRoutedIP object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewNullRoutedIPWithDefaults

`func NewNullRoutedIPWithDefaults() *NullRoutedIP`

NewNullRoutedIPWithDefaults instantiates a new NullRoutedIP object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *NullRoutedIP) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *NullRoutedIP) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *NullRoutedIP) SetId(v string)`

SetId sets Id field to given value.


### GetIp

`func (o *NullRoutedIP) GetIp() string`

GetIp returns the Ip field if non-nil, zero value otherwise.

### GetIpOk

`func (o *NullRoutedIP) GetIpOk() (*string, bool)`

GetIpOk returns a tuple with the Ip field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIp

`func (o *NullRoutedIP) SetIp(v string)`

SetIp sets Ip field to given value.


### GetNulledAt

`func (o *NullRoutedIP) GetNulledAt() time.Time`

GetNulledAt returns the NulledAt field if non-nil, zero value otherwise.

### GetNulledAtOk

`func (o *NullRoutedIP) GetNulledAtOk() (*time.Time, bool)`

GetNulledAtOk returns a tuple with the NulledAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNulledAt

`func (o *NullRoutedIP) SetNulledAt(v time.Time)`

SetNulledAt sets NulledAt field to given value.


### GetNulledBy

`func (o *NullRoutedIP) GetNulledBy() string`

GetNulledBy returns the NulledBy field if non-nil, zero value otherwise.

### GetNulledByOk

`func (o *NullRoutedIP) GetNulledByOk() (*string, bool)`

GetNulledByOk returns a tuple with the NulledBy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNulledBy

`func (o *NullRoutedIP) SetNulledBy(v string)`

SetNulledBy sets NulledBy field to given value.


### GetNullLevel

`func (o *NullRoutedIP) GetNullLevel() int32`

GetNullLevel returns the NullLevel field if non-nil, zero value otherwise.

### GetNullLevelOk

`func (o *NullRoutedIP) GetNullLevelOk() (*int32, bool)`

GetNullLevelOk returns a tuple with the NullLevel field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNullLevel

`func (o *NullRoutedIP) SetNullLevel(v int32)`

SetNullLevel sets NullLevel field to given value.


### GetAutomatedUnnullingAt

`func (o *NullRoutedIP) GetAutomatedUnnullingAt() time.Time`

GetAutomatedUnnullingAt returns the AutomatedUnnullingAt field if non-nil, zero value otherwise.

### GetAutomatedUnnullingAtOk

`func (o *NullRoutedIP) GetAutomatedUnnullingAtOk() (*time.Time, bool)`

GetAutomatedUnnullingAtOk returns a tuple with the AutomatedUnnullingAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAutomatedUnnullingAt

`func (o *NullRoutedIP) SetAutomatedUnnullingAt(v time.Time)`

SetAutomatedUnnullingAt sets AutomatedUnnullingAt field to given value.


### SetAutomatedUnnullingAtNil

`func (o *NullRoutedIP) SetAutomatedUnnullingAtNil(b bool)`

 SetAutomatedUnnullingAtNil sets the value for AutomatedUnnullingAt to be an explicit nil

### UnsetAutomatedUnnullingAt
`func (o *NullRoutedIP) UnsetAutomatedUnnullingAt()`

UnsetAutomatedUnnullingAt ensures that no value is present for AutomatedUnnullingAt, not even an explicit nil
### GetUnnulledAt

`func (o *NullRoutedIP) GetUnnulledAt() time.Time`

GetUnnulledAt returns the UnnulledAt field if non-nil, zero value otherwise.

### GetUnnulledAtOk

`func (o *NullRoutedIP) GetUnnulledAtOk() (*time.Time, bool)`

GetUnnulledAtOk returns a tuple with the UnnulledAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUnnulledAt

`func (o *NullRoutedIP) SetUnnulledAt(v time.Time)`

SetUnnulledAt sets UnnulledAt field to given value.


### SetUnnulledAtNil

`func (o *NullRoutedIP) SetUnnulledAtNil(b bool)`

 SetUnnulledAtNil sets the value for UnnulledAt to be an explicit nil

### UnsetUnnulledAt
`func (o *NullRoutedIP) UnsetUnnulledAt()`

UnsetUnnulledAt ensures that no value is present for UnnulledAt, not even an explicit nil
### GetUnnulledBy

`func (o *NullRoutedIP) GetUnnulledBy() string`

GetUnnulledBy returns the UnnulledBy field if non-nil, zero value otherwise.

### GetUnnulledByOk

`func (o *NullRoutedIP) GetUnnulledByOk() (*string, bool)`

GetUnnulledByOk returns a tuple with the UnnulledBy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUnnulledBy

`func (o *NullRoutedIP) SetUnnulledBy(v string)`

SetUnnulledBy sets UnnulledBy field to given value.


### SetUnnulledByNil

`func (o *NullRoutedIP) SetUnnulledByNil(b bool)`

 SetUnnulledByNil sets the value for UnnulledBy to be an explicit nil

### UnsetUnnulledBy
`func (o *NullRoutedIP) UnsetUnnulledBy()`

UnsetUnnulledBy ensures that no value is present for UnnulledBy, not even an explicit nil
### GetTicketId

`func (o *NullRoutedIP) GetTicketId() string`

GetTicketId returns the TicketId field if non-nil, zero value otherwise.

### GetTicketIdOk

`func (o *NullRoutedIP) GetTicketIdOk() (*string, bool)`

GetTicketIdOk returns a tuple with the TicketId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTicketId

`func (o *NullRoutedIP) SetTicketId(v string)`

SetTicketId sets TicketId field to given value.


### SetTicketIdNil

`func (o *NullRoutedIP) SetTicketIdNil(b bool)`

 SetTicketIdNil sets the value for TicketId to be an explicit nil

### UnsetTicketId
`func (o *NullRoutedIP) UnsetTicketId()`

UnsetTicketId ensures that no value is present for TicketId, not even an explicit nil
### GetComment

`func (o *NullRoutedIP) GetComment() string`

GetComment returns the Comment field if non-nil, zero value otherwise.

### GetCommentOk

`func (o *NullRoutedIP) GetCommentOk() (*string, bool)`

GetCommentOk returns a tuple with the Comment field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetComment

`func (o *NullRoutedIP) SetComment(v string)`

SetComment sets Comment field to given value.


### SetCommentNil

`func (o *NullRoutedIP) SetCommentNil(b bool)`

 SetCommentNil sets the value for Comment to be an explicit nil

### UnsetComment
`func (o *NullRoutedIP) UnsetComment()`

UnsetComment ensures that no value is present for Comment, not even an explicit nil
### GetEquipmentId

`func (o *NullRoutedIP) GetEquipmentId() string`

GetEquipmentId returns the EquipmentId field if non-nil, zero value otherwise.

### GetEquipmentIdOk

`func (o *NullRoutedIP) GetEquipmentIdOk() (*string, bool)`

GetEquipmentIdOk returns a tuple with the EquipmentId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEquipmentId

`func (o *NullRoutedIP) SetEquipmentId(v string)`

SetEquipmentId sets EquipmentId field to given value.


### GetAssignedContract

`func (o *NullRoutedIP) GetAssignedContract() AssignedContract`

GetAssignedContract returns the AssignedContract field if non-nil, zero value otherwise.

### GetAssignedContractOk

`func (o *NullRoutedIP) GetAssignedContractOk() (*AssignedContract, bool)`

GetAssignedContractOk returns a tuple with the AssignedContract field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAssignedContract

`func (o *NullRoutedIP) SetAssignedContract(v AssignedContract)`

SetAssignedContract sets AssignedContract field to given value.


### SetAssignedContractNil

`func (o *NullRoutedIP) SetAssignedContractNil(b bool)`

 SetAssignedContractNil sets the value for AssignedContract to be an explicit nil

### UnsetAssignedContract
`func (o *NullRoutedIP) UnsetAssignedContract()`

UnsetAssignedContract ensures that no value is present for AssignedContract, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


