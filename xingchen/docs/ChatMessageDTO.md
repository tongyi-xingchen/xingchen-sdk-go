# ChatMessageDTO

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Content** | Pointer to **string** | 消息内容 | [optional] 
**MessageType** | Pointer to **string** | 消息类型，opening_remarks: 开场白，sys_greetings: 系统问候，user: 用户, character: 角色 | [optional] 
**MessageId** | Pointer to **string** | 消息ID | [optional] 
**MessageIssuer** | Pointer to [**ChatRoomUserDTO**](ChatRoomUserDTO.md) |  | [optional] 
**SessionId** | Pointer to **string** | 消息对话session_id | [optional] 
**ChatId** | Pointer to **string** | 一次对话ID | [optional] 
**GmtCreate** | Pointer to **int64** | 消息创建时间 | [optional] 
**IsGreeting** | Pointer to **bool** | 是否为开场白 | [optional] 
**Rating** | Pointer to **float64** | 消息评分 | [optional] 
**RegenerateTimes** | Pointer to **int32** | 重新生成次数 | [optional] 

## Methods

### NewChatMessageDTO

`func NewChatMessageDTO() *ChatMessageDTO`

NewChatMessageDTO instantiates a new ChatMessageDTO object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewChatMessageDTOWithDefaults

`func NewChatMessageDTOWithDefaults() *ChatMessageDTO`

NewChatMessageDTOWithDefaults instantiates a new ChatMessageDTO object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetContent

`func (o *ChatMessageDTO) GetContent() string`

GetContent returns the Content field if non-nil, zero value otherwise.

### GetContentOk

`func (o *ChatMessageDTO) GetContentOk() (*string, bool)`

GetContentOk returns a tuple with the Content field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContent

`func (o *ChatMessageDTO) SetContent(v string)`

SetContent sets Content field to given value.

### HasContent

`func (o *ChatMessageDTO) HasContent() bool`

HasContent returns a boolean if a field has been set.

### GetMessageType

`func (o *ChatMessageDTO) GetMessageType() string`

GetMessageType returns the MessageType field if non-nil, zero value otherwise.

### GetMessageTypeOk

`func (o *ChatMessageDTO) GetMessageTypeOk() (*string, bool)`

GetMessageTypeOk returns a tuple with the MessageType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMessageType

`func (o *ChatMessageDTO) SetMessageType(v string)`

SetMessageType sets MessageType field to given value.

### HasMessageType

`func (o *ChatMessageDTO) HasMessageType() bool`

HasMessageType returns a boolean if a field has been set.

### GetMessageId

`func (o *ChatMessageDTO) GetMessageId() string`

GetMessageId returns the MessageId field if non-nil, zero value otherwise.

### GetMessageIdOk

`func (o *ChatMessageDTO) GetMessageIdOk() (*string, bool)`

GetMessageIdOk returns a tuple with the MessageId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMessageId

`func (o *ChatMessageDTO) SetMessageId(v string)`

SetMessageId sets MessageId field to given value.

### HasMessageId

`func (o *ChatMessageDTO) HasMessageId() bool`

HasMessageId returns a boolean if a field has been set.

### GetMessageIssuer

`func (o *ChatMessageDTO) GetMessageIssuer() ChatRoomUserDTO`

GetMessageIssuer returns the MessageIssuer field if non-nil, zero value otherwise.

### GetMessageIssuerOk

`func (o *ChatMessageDTO) GetMessageIssuerOk() (*ChatRoomUserDTO, bool)`

GetMessageIssuerOk returns a tuple with the MessageIssuer field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMessageIssuer

`func (o *ChatMessageDTO) SetMessageIssuer(v ChatRoomUserDTO)`

SetMessageIssuer sets MessageIssuer field to given value.

### HasMessageIssuer

`func (o *ChatMessageDTO) HasMessageIssuer() bool`

HasMessageIssuer returns a boolean if a field has been set.

### GetSessionId

`func (o *ChatMessageDTO) GetSessionId() string`

GetSessionId returns the SessionId field if non-nil, zero value otherwise.

### GetSessionIdOk

`func (o *ChatMessageDTO) GetSessionIdOk() (*string, bool)`

GetSessionIdOk returns a tuple with the SessionId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSessionId

`func (o *ChatMessageDTO) SetSessionId(v string)`

SetSessionId sets SessionId field to given value.

### HasSessionId

`func (o *ChatMessageDTO) HasSessionId() bool`

HasSessionId returns a boolean if a field has been set.

### GetChatId

`func (o *ChatMessageDTO) GetChatId() string`

GetChatId returns the ChatId field if non-nil, zero value otherwise.

### GetChatIdOk

`func (o *ChatMessageDTO) GetChatIdOk() (*string, bool)`

GetChatIdOk returns a tuple with the ChatId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetChatId

`func (o *ChatMessageDTO) SetChatId(v string)`

SetChatId sets ChatId field to given value.

### HasChatId

`func (o *ChatMessageDTO) HasChatId() bool`

HasChatId returns a boolean if a field has been set.

### GetGmtCreate

`func (o *ChatMessageDTO) GetGmtCreate() int64`

GetGmtCreate returns the GmtCreate field if non-nil, zero value otherwise.

### GetGmtCreateOk

`func (o *ChatMessageDTO) GetGmtCreateOk() (*int64, bool)`

GetGmtCreateOk returns a tuple with the GmtCreate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGmtCreate

`func (o *ChatMessageDTO) SetGmtCreate(v int64)`

SetGmtCreate sets GmtCreate field to given value.

### HasGmtCreate

`func (o *ChatMessageDTO) HasGmtCreate() bool`

HasGmtCreate returns a boolean if a field has been set.

### GetIsGreeting

`func (o *ChatMessageDTO) GetIsGreeting() bool`

GetIsGreeting returns the IsGreeting field if non-nil, zero value otherwise.

### GetIsGreetingOk

`func (o *ChatMessageDTO) GetIsGreetingOk() (*bool, bool)`

GetIsGreetingOk returns a tuple with the IsGreeting field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsGreeting

`func (o *ChatMessageDTO) SetIsGreeting(v bool)`

SetIsGreeting sets IsGreeting field to given value.

### HasIsGreeting

`func (o *ChatMessageDTO) HasIsGreeting() bool`

HasIsGreeting returns a boolean if a field has been set.

### GetRating

`func (o *ChatMessageDTO) GetRating() float64`

GetRating returns the Rating field if non-nil, zero value otherwise.

### GetRatingOk

`func (o *ChatMessageDTO) GetRatingOk() (*float64, bool)`

GetRatingOk returns a tuple with the Rating field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRating

`func (o *ChatMessageDTO) SetRating(v float64)`

SetRating sets Rating field to given value.

### HasRating

`func (o *ChatMessageDTO) HasRating() bool`

HasRating returns a boolean if a field has been set.

### GetRegenerateTimes

`func (o *ChatMessageDTO) GetRegenerateTimes() int32`

GetRegenerateTimes returns the RegenerateTimes field if non-nil, zero value otherwise.

### GetRegenerateTimesOk

`func (o *ChatMessageDTO) GetRegenerateTimesOk() (*int32, bool)`

GetRegenerateTimesOk returns a tuple with the RegenerateTimes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRegenerateTimes

`func (o *ChatMessageDTO) SetRegenerateTimes(v int32)`

SetRegenerateTimes sets RegenerateTimes field to given value.

### HasRegenerateTimes

`func (o *ChatMessageDTO) HasRegenerateTimes() bool`

HasRegenerateTimes returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


