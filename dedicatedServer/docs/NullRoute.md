# NullRoute

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AutomatedUnnullingAt** | Pointer to **time.Time** | The time the null route was removed or will be removed. | [optional] 
**Comment** | Pointer to **string** | An optional comment for the reason of the null route | [optional] 
**Ip** | Pointer to **string** | The ip address that was null routed | [optional] 
**NullLevel** | Pointer to **int32** | The level of the null route | [optional] 
**NulledAt** | Pointer to **time.Time** | The time the null route was created | [optional] 
**TicketId** | Pointer to **string** | A ticket number if available | [optional] 

## Methods

### NewNullRoute

`func NewNullRoute() *NullRoute`

NewNullRoute instantiates a new NullRoute object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewNullRouteWithDefaults

`func NewNullRouteWithDefaults() *NullRoute`

NewNullRouteWithDefaults instantiates a new NullRoute object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAutomatedUnnullingAt

`func (o *NullRoute) GetAutomatedUnnullingAt() time.Time`

GetAutomatedUnnullingAt returns the AutomatedUnnullingAt field if non-nil, zero value otherwise.

### GetAutomatedUnnullingAtOk

`func (o *NullRoute) GetAutomatedUnnullingAtOk() (*time.Time, bool)`

GetAutomatedUnnullingAtOk returns a tuple with the AutomatedUnnullingAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAutomatedUnnullingAt

`func (o *NullRoute) SetAutomatedUnnullingAt(v time.Time)`

SetAutomatedUnnullingAt sets AutomatedUnnullingAt field to given value.

### HasAutomatedUnnullingAt

`func (o *NullRoute) HasAutomatedUnnullingAt() bool`

HasAutomatedUnnullingAt returns a boolean if a field has been set.

### GetComment

`func (o *NullRoute) GetComment() string`

GetComment returns the Comment field if non-nil, zero value otherwise.

### GetCommentOk

`func (o *NullRoute) GetCommentOk() (*string, bool)`

GetCommentOk returns a tuple with the Comment field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetComment

`func (o *NullRoute) SetComment(v string)`

SetComment sets Comment field to given value.

### HasComment

`func (o *NullRoute) HasComment() bool`

HasComment returns a boolean if a field has been set.

### GetIp

`func (o *NullRoute) GetIp() string`

GetIp returns the Ip field if non-nil, zero value otherwise.

### GetIpOk

`func (o *NullRoute) GetIpOk() (*string, bool)`

GetIpOk returns a tuple with the Ip field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIp

`func (o *NullRoute) SetIp(v string)`

SetIp sets Ip field to given value.

### HasIp

`func (o *NullRoute) HasIp() bool`

HasIp returns a boolean if a field has been set.

### GetNullLevel

`func (o *NullRoute) GetNullLevel() int32`

GetNullLevel returns the NullLevel field if non-nil, zero value otherwise.

### GetNullLevelOk

`func (o *NullRoute) GetNullLevelOk() (*int32, bool)`

GetNullLevelOk returns a tuple with the NullLevel field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNullLevel

`func (o *NullRoute) SetNullLevel(v int32)`

SetNullLevel sets NullLevel field to given value.

### HasNullLevel

`func (o *NullRoute) HasNullLevel() bool`

HasNullLevel returns a boolean if a field has been set.

### GetNulledAt

`func (o *NullRoute) GetNulledAt() time.Time`

GetNulledAt returns the NulledAt field if non-nil, zero value otherwise.

### GetNulledAtOk

`func (o *NullRoute) GetNulledAtOk() (*time.Time, bool)`

GetNulledAtOk returns a tuple with the NulledAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNulledAt

`func (o *NullRoute) SetNulledAt(v time.Time)`

SetNulledAt sets NulledAt field to given value.

### HasNulledAt

`func (o *NullRoute) HasNulledAt() bool`

HasNulledAt returns a boolean if a field has been set.

### GetTicketId

`func (o *NullRoute) GetTicketId() string`

GetTicketId returns the TicketId field if non-nil, zero value otherwise.

### GetTicketIdOk

`func (o *NullRoute) GetTicketIdOk() (*string, bool)`

GetTicketIdOk returns a tuple with the TicketId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTicketId

`func (o *NullRoute) SetTicketId(v string)`

SetTicketId sets TicketId field to given value.

### HasTicketId

`func (o *NullRoute) HasTicketId() bool`

HasTicketId returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


