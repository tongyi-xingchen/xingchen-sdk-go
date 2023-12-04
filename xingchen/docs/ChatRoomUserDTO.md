# ChatRoomUserDTO

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**UserId** | Pointer to **string** | 用户ID，若为普通用户，则为userId,;若为角色，则为characterId | [optional] 
**BizUserId** | Pointer to **string** | 若为ISV用户，则为ISV用户ID；若为角色，则为空 | [optional] 
**UserName** | Pointer to **string** | 用户名称，若为普通用户，则为用户名称；若为角色，则为 characterName | [optional] 
**UserType** | Pointer to **string** | 用户类型，可以是 user（用户）或character（角色） | [optional] 

## Methods

### NewChatRoomUserDTO

`func NewChatRoomUserDTO() *ChatRoomUserDTO`

NewChatRoomUserDTO instantiates a new ChatRoomUserDTO object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewChatRoomUserDTOWithDefaults

`func NewChatRoomUserDTOWithDefaults() *ChatRoomUserDTO`

NewChatRoomUserDTOWithDefaults instantiates a new ChatRoomUserDTO object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetUserId

`func (o *ChatRoomUserDTO) GetUserId() string`

GetUserId returns the UserId field if non-nil, zero value otherwise.

### GetUserIdOk

`func (o *ChatRoomUserDTO) GetUserIdOk() (*string, bool)`

GetUserIdOk returns a tuple with the UserId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserId

`func (o *ChatRoomUserDTO) SetUserId(v string)`

SetUserId sets UserId field to given value.

### HasUserId

`func (o *ChatRoomUserDTO) HasUserId() bool`

HasUserId returns a boolean if a field has been set.

### GetBizUserId

`func (o *ChatRoomUserDTO) GetBizUserId() string`

GetBizUserId returns the BizUserId field if non-nil, zero value otherwise.

### GetBizUserIdOk

`func (o *ChatRoomUserDTO) GetBizUserIdOk() (*string, bool)`

GetBizUserIdOk returns a tuple with the BizUserId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBizUserId

`func (o *ChatRoomUserDTO) SetBizUserId(v string)`

SetBizUserId sets BizUserId field to given value.

### HasBizUserId

`func (o *ChatRoomUserDTO) HasBizUserId() bool`

HasBizUserId returns a boolean if a field has been set.

### GetUserName

`func (o *ChatRoomUserDTO) GetUserName() string`

GetUserName returns the UserName field if non-nil, zero value otherwise.

### GetUserNameOk

`func (o *ChatRoomUserDTO) GetUserNameOk() (*string, bool)`

GetUserNameOk returns a tuple with the UserName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserName

`func (o *ChatRoomUserDTO) SetUserName(v string)`

SetUserName sets UserName field to given value.

### HasUserName

`func (o *ChatRoomUserDTO) HasUserName() bool`

HasUserName returns a boolean if a field has been set.

### GetUserType

`func (o *ChatRoomUserDTO) GetUserType() string`

GetUserType returns the UserType field if non-nil, zero value otherwise.

### GetUserTypeOk

`func (o *ChatRoomUserDTO) GetUserTypeOk() (*string, bool)`

GetUserTypeOk returns a tuple with the UserType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserType

`func (o *ChatRoomUserDTO) SetUserType(v string)`

SetUserType sets UserType field to given value.

### HasUserType

`func (o *ChatRoomUserDTO) HasUserType() bool`

HasUserType returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


