# CharacterVersionCreateOrUpdateDTO

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
**CharacterId** | **string** | 角色标识 | 
**Version** | **int32** | 角色版本，相同角色版本号唯一 | 

## Methods

### NewCharacterVersionCreateOrUpdateDTO

`func NewCharacterVersionCreateOrUpdateDTO(name string, introduction string, basicInformation string, openingLine string, characterId string, version int32, ) *CharacterVersionCreateOrUpdateDTO`

NewCharacterVersionCreateOrUpdateDTO instantiates a new CharacterVersionCreateOrUpdateDTO object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCharacterVersionCreateOrUpdateDTOWithDefaults

`func NewCharacterVersionCreateOrUpdateDTOWithDefaults() *CharacterVersionCreateOrUpdateDTO`

NewCharacterVersionCreateOrUpdateDTOWithDefaults instantiates a new CharacterVersionCreateOrUpdateDTO object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetGatewayAddContent

`func (o *CharacterVersionCreateOrUpdateDTO) GetGatewayAddContent() GatewayIssuedParams`

GetGatewayAddContent returns the GatewayAddContent field if non-nil, zero value otherwise.

### GetGatewayAddContentOk

`func (o *CharacterVersionCreateOrUpdateDTO) GetGatewayAddContentOk() (*GatewayIssuedParams, bool)`

GetGatewayAddContentOk returns a tuple with the GatewayAddContent field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGatewayAddContent

`func (o *CharacterVersionCreateOrUpdateDTO) SetGatewayAddContent(v GatewayIssuedParams)`

SetGatewayAddContent sets GatewayAddContent field to given value.

### HasGatewayAddContent

`func (o *CharacterVersionCreateOrUpdateDTO) HasGatewayAddContent() bool`

HasGatewayAddContent returns a boolean if a field has been set.

### GetName

`func (o *CharacterVersionCreateOrUpdateDTO) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *CharacterVersionCreateOrUpdateDTO) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *CharacterVersionCreateOrUpdateDTO) SetName(v string)`

SetName sets Name field to given value.


### GetAvatar

`func (o *CharacterVersionCreateOrUpdateDTO) GetAvatar() FileInfoVO`

GetAvatar returns the Avatar field if non-nil, zero value otherwise.

### GetAvatarOk

`func (o *CharacterVersionCreateOrUpdateDTO) GetAvatarOk() (*FileInfoVO, bool)`

GetAvatarOk returns a tuple with the Avatar field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAvatar

`func (o *CharacterVersionCreateOrUpdateDTO) SetAvatar(v FileInfoVO)`

SetAvatar sets Avatar field to given value.

### HasAvatar

`func (o *CharacterVersionCreateOrUpdateDTO) HasAvatar() bool`

HasAvatar returns a boolean if a field has been set.

### GetIntroduction

`func (o *CharacterVersionCreateOrUpdateDTO) GetIntroduction() string`

GetIntroduction returns the Introduction field if non-nil, zero value otherwise.

### GetIntroductionOk

`func (o *CharacterVersionCreateOrUpdateDTO) GetIntroductionOk() (*string, bool)`

GetIntroductionOk returns a tuple with the Introduction field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIntroduction

`func (o *CharacterVersionCreateOrUpdateDTO) SetIntroduction(v string)`

SetIntroduction sets Introduction field to given value.


### GetBasicInformation

`func (o *CharacterVersionCreateOrUpdateDTO) GetBasicInformation() string`

GetBasicInformation returns the BasicInformation field if non-nil, zero value otherwise.

### GetBasicInformationOk

`func (o *CharacterVersionCreateOrUpdateDTO) GetBasicInformationOk() (*string, bool)`

GetBasicInformationOk returns a tuple with the BasicInformation field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBasicInformation

`func (o *CharacterVersionCreateOrUpdateDTO) SetBasicInformation(v string)`

SetBasicInformation sets BasicInformation field to given value.


### GetOpeningLine

`func (o *CharacterVersionCreateOrUpdateDTO) GetOpeningLine() string`

GetOpeningLine returns the OpeningLine field if non-nil, zero value otherwise.

### GetOpeningLineOk

`func (o *CharacterVersionCreateOrUpdateDTO) GetOpeningLineOk() (*string, bool)`

GetOpeningLineOk returns a tuple with the OpeningLine field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOpeningLine

`func (o *CharacterVersionCreateOrUpdateDTO) SetOpeningLine(v string)`

SetOpeningLine sets OpeningLine field to given value.


### GetTraits

`func (o *CharacterVersionCreateOrUpdateDTO) GetTraits() string`

GetTraits returns the Traits field if non-nil, zero value otherwise.

### GetTraitsOk

`func (o *CharacterVersionCreateOrUpdateDTO) GetTraitsOk() (*string, bool)`

GetTraitsOk returns a tuple with the Traits field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTraits

`func (o *CharacterVersionCreateOrUpdateDTO) SetTraits(v string)`

SetTraits sets Traits field to given value.

### HasTraits

`func (o *CharacterVersionCreateOrUpdateDTO) HasTraits() bool`

HasTraits returns a boolean if a field has been set.

### GetChatExample

`func (o *CharacterVersionCreateOrUpdateDTO) GetChatExample() string`

GetChatExample returns the ChatExample field if non-nil, zero value otherwise.

### GetChatExampleOk

`func (o *CharacterVersionCreateOrUpdateDTO) GetChatExampleOk() (*string, bool)`

GetChatExampleOk returns a tuple with the ChatExample field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetChatExample

`func (o *CharacterVersionCreateOrUpdateDTO) SetChatExample(v string)`

SetChatExample sets ChatExample field to given value.

### HasChatExample

`func (o *CharacterVersionCreateOrUpdateDTO) HasChatExample() bool`

HasChatExample returns a boolean if a field has been set.

### GetType

`func (o *CharacterVersionCreateOrUpdateDTO) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *CharacterVersionCreateOrUpdateDTO) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *CharacterVersionCreateOrUpdateDTO) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *CharacterVersionCreateOrUpdateDTO) HasType() bool`

HasType returns a boolean if a field has been set.

### GetChatObjective

`func (o *CharacterVersionCreateOrUpdateDTO) GetChatObjective() string`

GetChatObjective returns the ChatObjective field if non-nil, zero value otherwise.

### GetChatObjectiveOk

`func (o *CharacterVersionCreateOrUpdateDTO) GetChatObjectiveOk() (*string, bool)`

GetChatObjectiveOk returns a tuple with the ChatObjective field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetChatObjective

`func (o *CharacterVersionCreateOrUpdateDTO) SetChatObjective(v string)`

SetChatObjective sets ChatObjective field to given value.

### HasChatObjective

`func (o *CharacterVersionCreateOrUpdateDTO) HasChatObjective() bool`

HasChatObjective returns a boolean if a field has been set.

### GetAdvancedConfig

`func (o *CharacterVersionCreateOrUpdateDTO) GetAdvancedConfig() CharacterAdvancedConfig`

GetAdvancedConfig returns the AdvancedConfig field if non-nil, zero value otherwise.

### GetAdvancedConfigOk

`func (o *CharacterVersionCreateOrUpdateDTO) GetAdvancedConfigOk() (*CharacterAdvancedConfig, bool)`

GetAdvancedConfigOk returns a tuple with the AdvancedConfig field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAdvancedConfig

`func (o *CharacterVersionCreateOrUpdateDTO) SetAdvancedConfig(v CharacterAdvancedConfig)`

SetAdvancedConfig sets AdvancedConfig field to given value.

### HasAdvancedConfig

`func (o *CharacterVersionCreateOrUpdateDTO) HasAdvancedConfig() bool`

HasAdvancedConfig returns a boolean if a field has been set.

### GetCharacterId

`func (o *CharacterVersionCreateOrUpdateDTO) GetCharacterId() string`

GetCharacterId returns the CharacterId field if non-nil, zero value otherwise.

### GetCharacterIdOk

`func (o *CharacterVersionCreateOrUpdateDTO) GetCharacterIdOk() (*string, bool)`

GetCharacterIdOk returns a tuple with the CharacterId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCharacterId

`func (o *CharacterVersionCreateOrUpdateDTO) SetCharacterId(v string)`

SetCharacterId sets CharacterId field to given value.


### GetVersion

`func (o *CharacterVersionCreateOrUpdateDTO) GetVersion() int32`

GetVersion returns the Version field if non-nil, zero value otherwise.

### GetVersionOk

`func (o *CharacterVersionCreateOrUpdateDTO) GetVersionOk() (*int32, bool)`

GetVersionOk returns a tuple with the Version field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVersion

`func (o *CharacterVersionCreateOrUpdateDTO) SetVersion(v int32)`

SetVersion sets Version field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


