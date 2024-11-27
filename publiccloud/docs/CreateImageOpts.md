# CreateImageOpts

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | **string** | The name of the custom image to be created. | 
**InstanceId** | **string** | The id of the instance from which the custom image will be created. | 

## Methods

### NewCreateImageOpts

`func NewCreateImageOpts(name string, instanceId string, ) *CreateImageOpts`

NewCreateImageOpts instantiates a new CreateImageOpts object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCreateImageOptsWithDefaults

`func NewCreateImageOptsWithDefaults() *CreateImageOpts`

NewCreateImageOptsWithDefaults instantiates a new CreateImageOpts object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *CreateImageOpts) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *CreateImageOpts) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *CreateImageOpts) SetName(v string)`

SetName sets Name field to given value.


### GetInstanceId

`func (o *CreateImageOpts) GetInstanceId() string`

GetInstanceId returns the InstanceId field if non-nil, zero value otherwise.

### GetInstanceIdOk

`func (o *CreateImageOpts) GetInstanceIdOk() (*string, bool)`

GetInstanceIdOk returns a tuple with the InstanceId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInstanceId

`func (o *CreateImageOpts) SetInstanceId(v string)`

SetInstanceId sets InstanceId field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


