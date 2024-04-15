# Message

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**PostedBy** | Pointer to **string** | A string to indicating if the message was posted by the customer or the abuse agent. | [optional] 
**PostedAt** | Pointer to **string** | The timestamp when the message was posted. | [optional] 
**Body** | Pointer to **string** | The posted message. | [optional] 
**Attachment** | Pointer to [**Attachment**](Attachment.md) |  | [optional] 

## Methods

### NewMessage

`func NewMessage() *Message`

NewMessage instantiates a new Message object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMessageWithDefaults

`func NewMessageWithDefaults() *Message`

NewMessageWithDefaults instantiates a new Message object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetPostedBy

`func (o *Message) GetPostedBy() string`

GetPostedBy returns the PostedBy field if non-nil, zero value otherwise.

### GetPostedByOk

`func (o *Message) GetPostedByOk() (*string, bool)`

GetPostedByOk returns a tuple with the PostedBy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPostedBy

`func (o *Message) SetPostedBy(v string)`

SetPostedBy sets PostedBy field to given value.

### HasPostedBy

`func (o *Message) HasPostedBy() bool`

HasPostedBy returns a boolean if a field has been set.

### GetPostedAt

`func (o *Message) GetPostedAt() string`

GetPostedAt returns the PostedAt field if non-nil, zero value otherwise.

### GetPostedAtOk

`func (o *Message) GetPostedAtOk() (*string, bool)`

GetPostedAtOk returns a tuple with the PostedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPostedAt

`func (o *Message) SetPostedAt(v string)`

SetPostedAt sets PostedAt field to given value.

### HasPostedAt

`func (o *Message) HasPostedAt() bool`

HasPostedAt returns a boolean if a field has been set.

### GetBody

`func (o *Message) GetBody() string`

GetBody returns the Body field if non-nil, zero value otherwise.

### GetBodyOk

`func (o *Message) GetBodyOk() (*string, bool)`

GetBodyOk returns a tuple with the Body field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBody

`func (o *Message) SetBody(v string)`

SetBody sets Body field to given value.

### HasBody

`func (o *Message) HasBody() bool`

HasBody returns a boolean if a field has been set.

### GetAttachment

`func (o *Message) GetAttachment() Attachment`

GetAttachment returns the Attachment field if non-nil, zero value otherwise.

### GetAttachmentOk

`func (o *Message) GetAttachmentOk() (*Attachment, bool)`

GetAttachmentOk returns a tuple with the Attachment field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAttachment

`func (o *Message) SetAttachment(v Attachment)`

SetAttachment sets Attachment field to given value.

### HasAttachment

`func (o *Message) HasAttachment() bool`

HasAttachment returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


