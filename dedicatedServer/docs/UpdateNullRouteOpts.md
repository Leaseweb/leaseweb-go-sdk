# UpdateNullRouteOpts

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Action** | **string** | The action to take | 
**IpAddress** | **string** | The IP address for which to announce or withdraw the null route | 

## Methods

### NewUpdateNullRouteOpts

`func NewUpdateNullRouteOpts(action string, ipAddress string, ) *UpdateNullRouteOpts`

NewUpdateNullRouteOpts instantiates a new UpdateNullRouteOpts object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUpdateNullRouteOptsWithDefaults

`func NewUpdateNullRouteOptsWithDefaults() *UpdateNullRouteOpts`

NewUpdateNullRouteOptsWithDefaults instantiates a new UpdateNullRouteOpts object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAction

`func (o *UpdateNullRouteOpts) GetAction() string`

GetAction returns the Action field if non-nil, zero value otherwise.

### GetActionOk

`func (o *UpdateNullRouteOpts) GetActionOk() (*string, bool)`

GetActionOk returns a tuple with the Action field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAction

`func (o *UpdateNullRouteOpts) SetAction(v string)`

SetAction sets Action field to given value.


### GetIpAddress

`func (o *UpdateNullRouteOpts) GetIpAddress() string`

GetIpAddress returns the IpAddress field if non-nil, zero value otherwise.

### GetIpAddressOk

`func (o *UpdateNullRouteOpts) GetIpAddressOk() (*string, bool)`

GetIpAddressOk returns a tuple with the IpAddress field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIpAddress

`func (o *UpdateNullRouteOpts) SetIpAddress(v string)`

SetIpAddress sets IpAddress field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


