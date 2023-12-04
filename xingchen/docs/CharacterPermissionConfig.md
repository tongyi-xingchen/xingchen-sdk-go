# CharacterPermissionConfig

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**IsPublic** | **int32** | 是否可查看角色属性(0:不公开;1:公开) | 
**AllowChat** | **int32** | 允许在平台和其他用户聊天(0:不允许,1:允许) | 
**AllowApi** | **int32** | 允许其他用户通过api调用(0:不允许,1:允许) | 

## Methods

### NewCharacterPermissionConfig

`func NewCharacterPermissionConfig(isPublic int32, allowChat int32, allowApi int32, ) *CharacterPermissionConfig`

NewCharacterPermissionConfig instantiates a new CharacterPermissionConfig object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCharacterPermissionConfigWithDefaults

`func NewCharacterPermissionConfigWithDefaults() *CharacterPermissionConfig`

NewCharacterPermissionConfigWithDefaults instantiates a new CharacterPermissionConfig object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetIsPublic

`func (o *CharacterPermissionConfig) GetIsPublic() int32`

GetIsPublic returns the IsPublic field if non-nil, zero value otherwise.

### GetIsPublicOk

`func (o *CharacterPermissionConfig) GetIsPublicOk() (*int32, bool)`

GetIsPublicOk returns a tuple with the IsPublic field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsPublic

`func (o *CharacterPermissionConfig) SetIsPublic(v int32)`

SetIsPublic sets IsPublic field to given value.


### GetAllowChat

`func (o *CharacterPermissionConfig) GetAllowChat() int32`

GetAllowChat returns the AllowChat field if non-nil, zero value otherwise.

### GetAllowChatOk

`func (o *CharacterPermissionConfig) GetAllowChatOk() (*int32, bool)`

GetAllowChatOk returns a tuple with the AllowChat field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAllowChat

`func (o *CharacterPermissionConfig) SetAllowChat(v int32)`

SetAllowChat sets AllowChat field to given value.


### GetAllowApi

`func (o *CharacterPermissionConfig) GetAllowApi() int32`

GetAllowApi returns the AllowApi field if non-nil, zero value otherwise.

### GetAllowApiOk

`func (o *CharacterPermissionConfig) GetAllowApiOk() (*int32, bool)`

GetAllowApiOk returns a tuple with the AllowApi field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAllowApi

`func (o *CharacterPermissionConfig) SetAllowApi(v int32)`

SetAllowApi sets AllowApi field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


