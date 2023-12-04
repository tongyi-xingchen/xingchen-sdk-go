# ChatHistoryQueryWhere

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CharacterId** | Pointer to **string** | 角色ID | [optional] 
**ChatRoomId** | Pointer to **int64** | 聊天室ID | [optional] 
**SessionId** | Pointer to **string** | 会话ID | [optional] 
**StartTime** | Pointer to **int64** | 开始时间 | [optional] 
**EndTime** | Pointer to **int64** | 结束时间 | [optional] 
**MessageIds** | Pointer to **[]string** | 消息ID列表 | [optional] 
**BizUserId** | Pointer to **string** | 对话用户，给PASS接口用 | [optional] 

## Methods

### NewChatHistoryQueryWhere

`func NewChatHistoryQueryWhere() *ChatHistoryQueryWhere`

NewChatHistoryQueryWhere instantiates a new ChatHistoryQueryWhere object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewChatHistoryQueryWhereWithDefaults

`func NewChatHistoryQueryWhereWithDefaults() *ChatHistoryQueryWhere`

NewChatHistoryQueryWhereWithDefaults instantiates a new ChatHistoryQueryWhere object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCharacterId

`func (o *ChatHistoryQueryWhere) GetCharacterId() string`

GetCharacterId returns the CharacterId field if non-nil, zero value otherwise.

### GetCharacterIdOk

`func (o *ChatHistoryQueryWhere) GetCharacterIdOk() (*string, bool)`

GetCharacterIdOk returns a tuple with the CharacterId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCharacterId

`func (o *ChatHistoryQueryWhere) SetCharacterId(v string)`

SetCharacterId sets CharacterId field to given value.

### HasCharacterId

`func (o *ChatHistoryQueryWhere) HasCharacterId() bool`

HasCharacterId returns a boolean if a field has been set.

### GetChatRoomId

`func (o *ChatHistoryQueryWhere) GetChatRoomId() int64`

GetChatRoomId returns the ChatRoomId field if non-nil, zero value otherwise.

### GetChatRoomIdOk

`func (o *ChatHistoryQueryWhere) GetChatRoomIdOk() (*int64, bool)`

GetChatRoomIdOk returns a tuple with the ChatRoomId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetChatRoomId

`func (o *ChatHistoryQueryWhere) SetChatRoomId(v int64)`

SetChatRoomId sets ChatRoomId field to given value.

### HasChatRoomId

`func (o *ChatHistoryQueryWhere) HasChatRoomId() bool`

HasChatRoomId returns a boolean if a field has been set.

### GetSessionId

`func (o *ChatHistoryQueryWhere) GetSessionId() string`

GetSessionId returns the SessionId field if non-nil, zero value otherwise.

### GetSessionIdOk

`func (o *ChatHistoryQueryWhere) GetSessionIdOk() (*string, bool)`

GetSessionIdOk returns a tuple with the SessionId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSessionId

`func (o *ChatHistoryQueryWhere) SetSessionId(v string)`

SetSessionId sets SessionId field to given value.

### HasSessionId

`func (o *ChatHistoryQueryWhere) HasSessionId() bool`

HasSessionId returns a boolean if a field has been set.

### GetStartTime

`func (o *ChatHistoryQueryWhere) GetStartTime() int64`

GetStartTime returns the StartTime field if non-nil, zero value otherwise.

### GetStartTimeOk

`func (o *ChatHistoryQueryWhere) GetStartTimeOk() (*int64, bool)`

GetStartTimeOk returns a tuple with the StartTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStartTime

`func (o *ChatHistoryQueryWhere) SetStartTime(v int64)`

SetStartTime sets StartTime field to given value.

### HasStartTime

`func (o *ChatHistoryQueryWhere) HasStartTime() bool`

HasStartTime returns a boolean if a field has been set.

### GetEndTime

`func (o *ChatHistoryQueryWhere) GetEndTime() int64`

GetEndTime returns the EndTime field if non-nil, zero value otherwise.

### GetEndTimeOk

`func (o *ChatHistoryQueryWhere) GetEndTimeOk() (*int64, bool)`

GetEndTimeOk returns a tuple with the EndTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEndTime

`func (o *ChatHistoryQueryWhere) SetEndTime(v int64)`

SetEndTime sets EndTime field to given value.

### HasEndTime

`func (o *ChatHistoryQueryWhere) HasEndTime() bool`

HasEndTime returns a boolean if a field has been set.

### GetMessageIds

`func (o *ChatHistoryQueryWhere) GetMessageIds() []string`

GetMessageIds returns the MessageIds field if non-nil, zero value otherwise.

### GetMessageIdsOk

`func (o *ChatHistoryQueryWhere) GetMessageIdsOk() (*[]string, bool)`

GetMessageIdsOk returns a tuple with the MessageIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMessageIds

`func (o *ChatHistoryQueryWhere) SetMessageIds(v []string)`

SetMessageIds sets MessageIds field to given value.

### HasMessageIds

`func (o *ChatHistoryQueryWhere) HasMessageIds() bool`

HasMessageIds returns a boolean if a field has been set.

### GetBizUserId

`func (o *ChatHistoryQueryWhere) GetBizUserId() string`

GetBizUserId returns the BizUserId field if non-nil, zero value otherwise.

### GetBizUserIdOk

`func (o *ChatHistoryQueryWhere) GetBizUserIdOk() (*string, bool)`

GetBizUserIdOk returns a tuple with the BizUserId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBizUserId

`func (o *ChatHistoryQueryWhere) SetBizUserId(v string)`

SetBizUserId sets BizUserId field to given value.

### HasBizUserId

`func (o *ChatHistoryQueryWhere) HasBizUserId() bool`

HasBizUserId returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


