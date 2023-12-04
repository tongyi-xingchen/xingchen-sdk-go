# CharacterUpdateDTO

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**GatewayAddContent** | Pointer to [**GatewayIssuedParams**](GatewayIssuedParams.md) |  | [optional] 
**Name** | **string** | 角色名称(最多20个字符) | 
**Avatar** | Pointer to [**FileInfoVO**](FileInfoVO.md) |  | [optional] 
**Introduction** | **string** | 角色描述 | 
**BasicInformation** | **string** | 基本信息 | 
**OpeningLine** | **string** | 开场白 | 
**Traits** | Pointer to **string** | 角色特点 | [optional] 
**ChatExample** | Pointer to **string** | 聊天示例 | [optional] 
**Type** | Pointer to **string** | 角色类型，virtual:虚拟角色，reproduction:已知角色复刻 | [optional] 
**ChatObjective** | Pointer to **string** | 对话目标 | [optional] 
**AdvancedConfig** | Pointer to [**CharacterAdvancedConfig**](CharacterAdvancedConfig.md) |  | [optional] 
**CharacterId** | **string** | 角色ID | 
**Version** | **int32** | 角色版本 | 
**PermConfig** | Pointer to [**CharacterPermissionConfig**](CharacterPermissionConfig.md) |  | [optional] 

## Methods

### NewCharacterUpdateDTO

`func NewCharacterUpdateDTO(name string, introduction string, basicInformation string, openingLine string, characterId string, version int32, ) *CharacterUpdateDTO`

NewCharacterUpdateDTO instantiates a new CharacterUpdateDTO object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCharacterUpdateDTOWithDefaults

`func NewCharacterUpdateDTOWithDefaults() *CharacterUpdateDTO`

NewCharacterUpdateDTOWithDefaults instantiates a new CharacterUpdateDTO object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetGatewayAddContent

`func (o *CharacterUpdateDTO) GetGatewayAddContent() GatewayIssuedParams`

GetGatewayAddContent returns the GatewayAddContent field if non-nil, zero value otherwise.

### GetGatewayAddContentOk

`func (o *CharacterUpdateDTO) GetGatewayAddContentOk() (*GatewayIssuedParams, bool)`

GetGatewayAddContentOk returns a tuple with the GatewayAddContent field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGatewayAddContent

`func (o *CharacterUpdateDTO) SetGatewayAddContent(v GatewayIssuedParams)`

SetGatewayAddContent sets GatewayAddContent field to given value.

### HasGatewayAddContent

`func (o *CharacterUpdateDTO) HasGatewayAddContent() bool`

HasGatewayAddContent returns a boolean if a field has been set.

### GetName

`func (o *CharacterUpdateDTO) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *CharacterUpdateDTO) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *CharacterUpdateDTO) SetName(v string)`

SetName sets Name field to given value.


### GetAvatar

`func (o *CharacterUpdateDTO) GetAvatar() FileInfoVO`

GetAvatar returns the Avatar field if non-nil, zero value otherwise.

### GetAvatarOk

`func (o *CharacterUpdateDTO) GetAvatarOk() (*FileInfoVO, bool)`

GetAvatarOk returns a tuple with the Avatar field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAvatar

`func (o *CharacterUpdateDTO) SetAvatar(v FileInfoVO)`

SetAvatar sets Avatar field to given value.

### HasAvatar

`func (o *CharacterUpdateDTO) HasAvatar() bool`

HasAvatar returns a boolean if a field has been set.

### GetIntroduction

`func (o *CharacterUpdateDTO) GetIntroduction() string`

GetIntroduction returns the Introduction field if non-nil, zero value otherwise.

### GetIntroductionOk

`func (o *CharacterUpdateDTO) GetIntroductionOk() (*string, bool)`

GetIntroductionOk returns a tuple with the Introduction field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIntroduction

`func (o *CharacterUpdateDTO) SetIntroduction(v string)`

SetIntroduction sets Introduction field to given value.


### GetBasicInformation

`func (o *CharacterUpdateDTO) GetBasicInformation() string`

GetBasicInformation returns the BasicInformation field if non-nil, zero value otherwise.

### GetBasicInformationOk

`func (o *CharacterUpdateDTO) GetBasicInformationOk() (*string, bool)`

GetBasicInformationOk returns a tuple with the BasicInformation field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBasicInformation

`func (o *CharacterUpdateDTO) SetBasicInformation(v string)`

SetBasicInformation sets BasicInformation field to given value.


### GetOpeningLine

`func (o *CharacterUpdateDTO) GetOpeningLine() string`

GetOpeningLine returns the OpeningLine field if non-nil, zero value otherwise.

### GetOpeningLineOk

`func (o *CharacterUpdateDTO) GetOpeningLineOk() (*string, bool)`

GetOpeningLineOk returns a tuple with the OpeningLine field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOpeningLine

`func (o *CharacterUpdateDTO) SetOpeningLine(v string)`

SetOpeningLine sets OpeningLine field to given value.


### GetTraits

`func (o *CharacterUpdateDTO) GetTraits() string`

GetTraits returns the Traits field if non-nil, zero value otherwise.

### GetTraitsOk

`func (o *CharacterUpdateDTO) GetTraitsOk() (*string, bool)`

GetTraitsOk returns a tuple with the Traits field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTraits

`func (o *CharacterUpdateDTO) SetTraits(v string)`

SetTraits sets Traits field to given value.

### HasTraits

`func (o *CharacterUpdateDTO) HasTraits() bool`

HasTraits returns a boolean if a field has been set.

### GetChatExample

`func (o *CharacterUpdateDTO) GetChatExample() string`

GetChatExample returns the ChatExample field if non-nil, zero value otherwise.

### GetChatExampleOk

`func (o *CharacterUpdateDTO) GetChatExampleOk() (*string, bool)`

GetChatExampleOk returns a tuple with the ChatExample field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetChatExample

`func (o *CharacterUpdateDTO) SetChatExample(v string)`

SetChatExample sets ChatExample field to given value.

### HasChatExample

`func (o *CharacterUpdateDTO) HasChatExample() bool`

HasChatExample returns a boolean if a field has been set.

### GetType

`func (o *CharacterUpdateDTO) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *CharacterUpdateDTO) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *CharacterUpdateDTO) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *CharacterUpdateDTO) HasType() bool`

HasType returns a boolean if a field has been set.

### GetChatObjective

`func (o *CharacterUpdateDTO) GetChatObjective() string`

GetChatObjective returns the ChatObjective field if non-nil, zero value otherwise.

### GetChatObjectiveOk

`func (o *CharacterUpdateDTO) GetChatObjectiveOk() (*string, bool)`

GetChatObjectiveOk returns a tuple with the ChatObjective field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetChatObjective

`func (o *CharacterUpdateDTO) SetChatObjective(v string)`

SetChatObjective sets ChatObjective field to given value.

### HasChatObjective

`func (o *CharacterUpdateDTO) HasChatObjective() bool`

HasChatObjective returns a boolean if a field has been set.

### GetAdvancedConfig

`func (o *CharacterUpdateDTO) GetAdvancedConfig() CharacterAdvancedConfig`

GetAdvancedConfig returns the AdvancedConfig field if non-nil, zero value otherwise.

### GetAdvancedConfigOk

`func (o *CharacterUpdateDTO) GetAdvancedConfigOk() (*CharacterAdvancedConfig, bool)`

GetAdvancedConfigOk returns a tuple with the AdvancedConfig field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAdvancedConfig

`func (o *CharacterUpdateDTO) SetAdvancedConfig(v CharacterAdvancedConfig)`

SetAdvancedConfig sets AdvancedConfig field to given value.

### HasAdvancedConfig

`func (o *CharacterUpdateDTO) HasAdvancedConfig() bool`

HasAdvancedConfig returns a boolean if a field has been set.

### GetCharacterId

`func (o *CharacterUpdateDTO) GetCharacterId() string`

GetCharacterId returns the CharacterId field if non-nil, zero value otherwise.

### GetCharacterIdOk

`func (o *CharacterUpdateDTO) GetCharacterIdOk() (*string, bool)`

GetCharacterIdOk returns a tuple with the CharacterId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCharacterId

`func (o *CharacterUpdateDTO) SetCharacterId(v string)`

SetCharacterId sets CharacterId field to given value.


### GetVersion

`func (o *CharacterUpdateDTO) GetVersion() int32`

GetVersion returns the Version field if non-nil, zero value otherwise.

### GetVersionOk

`func (o *CharacterUpdateDTO) GetVersionOk() (*int32, bool)`

GetVersionOk returns a tuple with the Version field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVersion

`func (o *CharacterUpdateDTO) SetVersion(v int32)`

SetVersion sets Version field to given value.


### GetPermConfig

`func (o *CharacterUpdateDTO) GetPermConfig() CharacterPermissionConfig`

GetPermConfig returns the PermConfig field if non-nil, zero value otherwise.

### GetPermConfigOk

`func (o *CharacterUpdateDTO) GetPermConfigOk() (*CharacterPermissionConfig, bool)`

GetPermConfigOk returns a tuple with the PermConfig field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPermConfig

`func (o *CharacterUpdateDTO) SetPermConfig(v CharacterPermissionConfig)`

SetPermConfig sets PermConfig field to given value.

### HasPermConfig

`func (o *CharacterUpdateDTO) HasPermConfig() bool`

HasPermConfig returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


