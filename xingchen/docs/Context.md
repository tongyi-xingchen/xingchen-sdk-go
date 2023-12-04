# Context

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ChatRoomId** | Pointer to **int64** |  | [optional] 
**SessionId** | Pointer to **string** |  | [optional] 
**ChatId** | Pointer to **string** |  | [optional] 
**MessageId** | Pointer to **string** |  | [optional] 
**ReplyMessageId** | Pointer to **string** |  | [optional] 
**EnableDataInspection** | Pointer to **bool** |  | [optional] 
**IsSave** | Pointer to **bool** |  | [optional] 
**RequestId** | Pointer to **string** |  | [optional] 
**ModelRequestId** | Pointer to **string** |  | [optional] 
**CharacterPk** | Pointer to **int64** |  | [optional] 
**ModelName** | Pointer to **string** |  | [optional] 

## Methods

### NewContext

`func NewContext() *Context`

NewContext instantiates a new Context object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewContextWithDefaults

`func NewContextWithDefaults() *Context`

NewContextWithDefaults instantiates a new Context object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetChatRoomId

`func (o *Context) GetChatRoomId() int64`

GetChatRoomId returns the ChatRoomId field if non-nil, zero value otherwise.

### GetChatRoomIdOk

`func (o *Context) GetChatRoomIdOk() (*int64, bool)`

GetChatRoomIdOk returns a tuple with the ChatRoomId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetChatRoomId

`func (o *Context) SetChatRoomId(v int64)`

SetChatRoomId sets ChatRoomId field to given value.

### HasChatRoomId

`func (o *Context) HasChatRoomId() bool`

HasChatRoomId returns a boolean if a field has been set.

### GetSessionId

`func (o *Context) GetSessionId() string`

GetSessionId returns the SessionId field if non-nil, zero value otherwise.

### GetSessionIdOk

`func (o *Context) GetSessionIdOk() (*string, bool)`

GetSessionIdOk returns a tuple with the SessionId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSessionId

`func (o *Context) SetSessionId(v string)`

SetSessionId sets SessionId field to given value.

### HasSessionId

`func (o *Context) HasSessionId() bool`

HasSessionId returns a boolean if a field has been set.

### GetChatId

`func (o *Context) GetChatId() string`

GetChatId returns the ChatId field if non-nil, zero value otherwise.

### GetChatIdOk

`func (o *Context) GetChatIdOk() (*string, bool)`

GetChatIdOk returns a tuple with the ChatId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetChatId

`func (o *Context) SetChatId(v string)`

SetChatId sets ChatId field to given value.

### HasChatId

`func (o *Context) HasChatId() bool`

HasChatId returns a boolean if a field has been set.

### GetMessageId

`func (o *Context) GetMessageId() string`

GetMessageId returns the MessageId field if non-nil, zero value otherwise.

### GetMessageIdOk

`func (o *Context) GetMessageIdOk() (*string, bool)`

GetMessageIdOk returns a tuple with the MessageId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMessageId

`func (o *Context) SetMessageId(v string)`

SetMessageId sets MessageId field to given value.

### HasMessageId

`func (o *Context) HasMessageId() bool`

HasMessageId returns a boolean if a field has been set.

### GetReplyMessageId

`func (o *Context) GetReplyMessageId() string`

GetReplyMessageId returns the ReplyMessageId field if non-nil, zero value otherwise.

### GetReplyMessageIdOk

`func (o *Context) GetReplyMessageIdOk() (*string, bool)`

GetReplyMessageIdOk returns a tuple with the ReplyMessageId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReplyMessageId

`func (o *Context) SetReplyMessageId(v string)`

SetReplyMessageId sets ReplyMessageId field to given value.

### HasReplyMessageId

`func (o *Context) HasReplyMessageId() bool`

HasReplyMessageId returns a boolean if a field has been set.

### GetEnableDataInspection

`func (o *Context) GetEnableDataInspection() bool`

GetEnableDataInspection returns the EnableDataInspection field if non-nil, zero value otherwise.

### GetEnableDataInspectionOk

`func (o *Context) GetEnableDataInspectionOk() (*bool, bool)`

GetEnableDataInspectionOk returns a tuple with the EnableDataInspection field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnableDataInspection

`func (o *Context) SetEnableDataInspection(v bool)`

SetEnableDataInspection sets EnableDataInspection field to given value.

### HasEnableDataInspection

`func (o *Context) HasEnableDataInspection() bool`

HasEnableDataInspection returns a boolean if a field has been set.

### GetIsSave

`func (o *Context) GetIsSave() bool`

GetIsSave returns the IsSave field if non-nil, zero value otherwise.

### GetIsSaveOk

`func (o *Context) GetIsSaveOk() (*bool, bool)`

GetIsSaveOk returns a tuple with the IsSave field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsSave

`func (o *Context) SetIsSave(v bool)`

SetIsSave sets IsSave field to given value.

### HasIsSave

`func (o *Context) HasIsSave() bool`

HasIsSave returns a boolean if a field has been set.

### GetRequestId

`func (o *Context) GetRequestId() string`

GetRequestId returns the RequestId field if non-nil, zero value otherwise.

### GetRequestIdOk

`func (o *Context) GetRequestIdOk() (*string, bool)`

GetRequestIdOk returns a tuple with the RequestId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRequestId

`func (o *Context) SetRequestId(v string)`

SetRequestId sets RequestId field to given value.

### HasRequestId

`func (o *Context) HasRequestId() bool`

HasRequestId returns a boolean if a field has been set.

### GetModelRequestId

`func (o *Context) GetModelRequestId() string`

GetModelRequestId returns the ModelRequestId field if non-nil, zero value otherwise.

### GetModelRequestIdOk

`func (o *Context) GetModelRequestIdOk() (*string, bool)`

GetModelRequestIdOk returns a tuple with the ModelRequestId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetModelRequestId

`func (o *Context) SetModelRequestId(v string)`

SetModelRequestId sets ModelRequestId field to given value.

### HasModelRequestId

`func (o *Context) HasModelRequestId() bool`

HasModelRequestId returns a boolean if a field has been set.

### GetCharacterPk

`func (o *Context) GetCharacterPk() int64`

GetCharacterPk returns the CharacterPk field if non-nil, zero value otherwise.

### GetCharacterPkOk

`func (o *Context) GetCharacterPkOk() (*int64, bool)`

GetCharacterPkOk returns a tuple with the CharacterPk field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCharacterPk

`func (o *Context) SetCharacterPk(v int64)`

SetCharacterPk sets CharacterPk field to given value.

### HasCharacterPk

`func (o *Context) HasCharacterPk() bool`

HasCharacterPk returns a boolean if a field has been set.

### GetModelName

`func (o *Context) GetModelName() string`

GetModelName returns the ModelName field if non-nil, zero value otherwise.

### GetModelNameOk

`func (o *Context) GetModelNameOk() (*string, bool)`

GetModelNameOk returns a tuple with the ModelName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetModelName

`func (o *Context) SetModelName(v string)`

SetModelName sets ModelName field to given value.

### HasModelName

`func (o *Context) HasModelName() bool`

HasModelName returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


