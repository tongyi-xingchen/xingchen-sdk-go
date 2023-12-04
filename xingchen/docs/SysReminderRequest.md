# SysReminderRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**GatewayAddContent** | Pointer to [**GatewayIssuedParams**](GatewayIssuedParams.md) |  | [optional] 
**CharacterId** | **string** | 角色ID | 
**ChatRoomId** | Pointer to **int64** | 聊天室ID | [optional] 
**Content** | **string** | 系统提醒内容 | 
**BizUserId** | **string** | 业务系统ID | 

## Methods

### NewSysReminderRequest

`func NewSysReminderRequest(characterId string, content string, bizUserId string, ) *SysReminderRequest`

NewSysReminderRequest instantiates a new SysReminderRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSysReminderRequestWithDefaults

`func NewSysReminderRequestWithDefaults() *SysReminderRequest`

NewSysReminderRequestWithDefaults instantiates a new SysReminderRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetGatewayAddContent

`func (o *SysReminderRequest) GetGatewayAddContent() GatewayIssuedParams`

GetGatewayAddContent returns the GatewayAddContent field if non-nil, zero value otherwise.

### GetGatewayAddContentOk

`func (o *SysReminderRequest) GetGatewayAddContentOk() (*GatewayIssuedParams, bool)`

GetGatewayAddContentOk returns a tuple with the GatewayAddContent field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGatewayAddContent

`func (o *SysReminderRequest) SetGatewayAddContent(v GatewayIssuedParams)`

SetGatewayAddContent sets GatewayAddContent field to given value.

### HasGatewayAddContent

`func (o *SysReminderRequest) HasGatewayAddContent() bool`

HasGatewayAddContent returns a boolean if a field has been set.

### GetCharacterId

`func (o *SysReminderRequest) GetCharacterId() string`

GetCharacterId returns the CharacterId field if non-nil, zero value otherwise.

### GetCharacterIdOk

`func (o *SysReminderRequest) GetCharacterIdOk() (*string, bool)`

GetCharacterIdOk returns a tuple with the CharacterId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCharacterId

`func (o *SysReminderRequest) SetCharacterId(v string)`

SetCharacterId sets CharacterId field to given value.


### GetChatRoomId

`func (o *SysReminderRequest) GetChatRoomId() int64`

GetChatRoomId returns the ChatRoomId field if non-nil, zero value otherwise.

### GetChatRoomIdOk

`func (o *SysReminderRequest) GetChatRoomIdOk() (*int64, bool)`

GetChatRoomIdOk returns a tuple with the ChatRoomId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetChatRoomId

`func (o *SysReminderRequest) SetChatRoomId(v int64)`

SetChatRoomId sets ChatRoomId field to given value.

### HasChatRoomId

`func (o *SysReminderRequest) HasChatRoomId() bool`

HasChatRoomId returns a boolean if a field has been set.

### GetContent

`func (o *SysReminderRequest) GetContent() string`

GetContent returns the Content field if non-nil, zero value otherwise.

### GetContentOk

`func (o *SysReminderRequest) GetContentOk() (*string, bool)`

GetContentOk returns a tuple with the Content field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContent

`func (o *SysReminderRequest) SetContent(v string)`

SetContent sets Content field to given value.


### GetBizUserId

`func (o *SysReminderRequest) GetBizUserId() string`

GetBizUserId returns the BizUserId field if non-nil, zero value otherwise.

### GetBizUserIdOk

`func (o *SysReminderRequest) GetBizUserIdOk() (*string, bool)`

GetBizUserIdOk returns a tuple with the BizUserId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBizUserId

`func (o *SysReminderRequest) SetBizUserId(v string)`

SetBizUserId sets BizUserId field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


