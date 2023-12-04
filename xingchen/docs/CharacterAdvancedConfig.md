# CharacterAdvancedConfig

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ModelConfig** | Pointer to [**ModelConfig**](ModelConfig.md) |  | [optional] 
**RepositoryInfo** | Pointer to [**RepositoryInfo**](RepositoryInfo.md) |  | [optional] 
**IsRealInfo** | Pointer to **bool** | 是否返回真实世界信息 | [optional] 
**SearchKeyword** | Pointer to **string** | web搜索必填关键字 | [optional] 
**AllowSendImage** | Pointer to **bool** | 是否允许角色发送图片 | [optional] 
**ImageStyle** | Pointer to **string** | 角色发送图片的风格 | [optional] 
**AllowSendAsr** | Pointer to **bool** | 是否允许角色发送语音 | [optional] 
**AsrStyle** | Pointer to **string** | 角色发送语音风格 | [optional] 
**ChatDescription** | Pointer to **string** | 对话介绍 | [optional] 

## Methods

### NewCharacterAdvancedConfig

`func NewCharacterAdvancedConfig() *CharacterAdvancedConfig`

NewCharacterAdvancedConfig instantiates a new CharacterAdvancedConfig object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCharacterAdvancedConfigWithDefaults

`func NewCharacterAdvancedConfigWithDefaults() *CharacterAdvancedConfig`

NewCharacterAdvancedConfigWithDefaults instantiates a new CharacterAdvancedConfig object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetModelConfig

`func (o *CharacterAdvancedConfig) GetModelConfig() ModelConfig`

GetModelConfig returns the ModelConfig field if non-nil, zero value otherwise.

### GetModelConfigOk

`func (o *CharacterAdvancedConfig) GetModelConfigOk() (*ModelConfig, bool)`

GetModelConfigOk returns a tuple with the ModelConfig field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetModelConfig

`func (o *CharacterAdvancedConfig) SetModelConfig(v ModelConfig)`

SetModelConfig sets ModelConfig field to given value.

### HasModelConfig

`func (o *CharacterAdvancedConfig) HasModelConfig() bool`

HasModelConfig returns a boolean if a field has been set.

### GetRepositoryInfo

`func (o *CharacterAdvancedConfig) GetRepositoryInfo() RepositoryInfo`

GetRepositoryInfo returns the RepositoryInfo field if non-nil, zero value otherwise.

### GetRepositoryInfoOk

`func (o *CharacterAdvancedConfig) GetRepositoryInfoOk() (*RepositoryInfo, bool)`

GetRepositoryInfoOk returns a tuple with the RepositoryInfo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRepositoryInfo

`func (o *CharacterAdvancedConfig) SetRepositoryInfo(v RepositoryInfo)`

SetRepositoryInfo sets RepositoryInfo field to given value.

### HasRepositoryInfo

`func (o *CharacterAdvancedConfig) HasRepositoryInfo() bool`

HasRepositoryInfo returns a boolean if a field has been set.

### GetIsRealInfo

`func (o *CharacterAdvancedConfig) GetIsRealInfo() bool`

GetIsRealInfo returns the IsRealInfo field if non-nil, zero value otherwise.

### GetIsRealInfoOk

`func (o *CharacterAdvancedConfig) GetIsRealInfoOk() (*bool, bool)`

GetIsRealInfoOk returns a tuple with the IsRealInfo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsRealInfo

`func (o *CharacterAdvancedConfig) SetIsRealInfo(v bool)`

SetIsRealInfo sets IsRealInfo field to given value.

### HasIsRealInfo

`func (o *CharacterAdvancedConfig) HasIsRealInfo() bool`

HasIsRealInfo returns a boolean if a field has been set.

### GetSearchKeyword

`func (o *CharacterAdvancedConfig) GetSearchKeyword() string`

GetSearchKeyword returns the SearchKeyword field if non-nil, zero value otherwise.

### GetSearchKeywordOk

`func (o *CharacterAdvancedConfig) GetSearchKeywordOk() (*string, bool)`

GetSearchKeywordOk returns a tuple with the SearchKeyword field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSearchKeyword

`func (o *CharacterAdvancedConfig) SetSearchKeyword(v string)`

SetSearchKeyword sets SearchKeyword field to given value.

### HasSearchKeyword

`func (o *CharacterAdvancedConfig) HasSearchKeyword() bool`

HasSearchKeyword returns a boolean if a field has been set.

### GetAllowSendImage

`func (o *CharacterAdvancedConfig) GetAllowSendImage() bool`

GetAllowSendImage returns the AllowSendImage field if non-nil, zero value otherwise.

### GetAllowSendImageOk

`func (o *CharacterAdvancedConfig) GetAllowSendImageOk() (*bool, bool)`

GetAllowSendImageOk returns a tuple with the AllowSendImage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAllowSendImage

`func (o *CharacterAdvancedConfig) SetAllowSendImage(v bool)`

SetAllowSendImage sets AllowSendImage field to given value.

### HasAllowSendImage

`func (o *CharacterAdvancedConfig) HasAllowSendImage() bool`

HasAllowSendImage returns a boolean if a field has been set.

### GetImageStyle

`func (o *CharacterAdvancedConfig) GetImageStyle() string`

GetImageStyle returns the ImageStyle field if non-nil, zero value otherwise.

### GetImageStyleOk

`func (o *CharacterAdvancedConfig) GetImageStyleOk() (*string, bool)`

GetImageStyleOk returns a tuple with the ImageStyle field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetImageStyle

`func (o *CharacterAdvancedConfig) SetImageStyle(v string)`

SetImageStyle sets ImageStyle field to given value.

### HasImageStyle

`func (o *CharacterAdvancedConfig) HasImageStyle() bool`

HasImageStyle returns a boolean if a field has been set.

### GetAllowSendAsr

`func (o *CharacterAdvancedConfig) GetAllowSendAsr() bool`

GetAllowSendAsr returns the AllowSendAsr field if non-nil, zero value otherwise.

### GetAllowSendAsrOk

`func (o *CharacterAdvancedConfig) GetAllowSendAsrOk() (*bool, bool)`

GetAllowSendAsrOk returns a tuple with the AllowSendAsr field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAllowSendAsr

`func (o *CharacterAdvancedConfig) SetAllowSendAsr(v bool)`

SetAllowSendAsr sets AllowSendAsr field to given value.

### HasAllowSendAsr

`func (o *CharacterAdvancedConfig) HasAllowSendAsr() bool`

HasAllowSendAsr returns a boolean if a field has been set.

### GetAsrStyle

`func (o *CharacterAdvancedConfig) GetAsrStyle() string`

GetAsrStyle returns the AsrStyle field if non-nil, zero value otherwise.

### GetAsrStyleOk

`func (o *CharacterAdvancedConfig) GetAsrStyleOk() (*string, bool)`

GetAsrStyleOk returns a tuple with the AsrStyle field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAsrStyle

`func (o *CharacterAdvancedConfig) SetAsrStyle(v string)`

SetAsrStyle sets AsrStyle field to given value.

### HasAsrStyle

`func (o *CharacterAdvancedConfig) HasAsrStyle() bool`

HasAsrStyle returns a boolean if a field has been set.

### GetChatDescription

`func (o *CharacterAdvancedConfig) GetChatDescription() string`

GetChatDescription returns the ChatDescription field if non-nil, zero value otherwise.

### GetChatDescriptionOk

`func (o *CharacterAdvancedConfig) GetChatDescriptionOk() (*string, bool)`

GetChatDescriptionOk returns a tuple with the ChatDescription field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetChatDescription

`func (o *CharacterAdvancedConfig) SetChatDescription(v string)`

SetChatDescription sets ChatDescription field to given value.

### HasChatDescription

`func (o *CharacterAdvancedConfig) HasChatDescription() bool`

HasChatDescription returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


