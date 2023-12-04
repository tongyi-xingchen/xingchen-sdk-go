# CharacterDTO

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
**CharacterId** | Pointer to **string** | 角色唯一ID | [optional] 
**Version** | Pointer to **int32** | 角色版本 | [optional] 
**MajorVersion** | Pointer to **int32** | 主版本 | [optional] 
**GmtCreate** | Pointer to **int64** | 创建时间 | [optional] 
**GmtModified** | Pointer to **int64** | 修改时间 | [optional] 
**UserId** | Pointer to **string** | 角色创建用户ID | [optional] 
**ModifiedBy** | Pointer to **string** | 角色最后修改人ID | [optional] 
**Tags** | Pointer to **[]string** | 角色标签 | [optional] 
**Bubbles** | Pointer to **[]string** | 角色气泡 | [optional] 
**BackgroundImgUrl** | Pointer to **string** | 背景图片 | [optional] 
**PermConfig** | Pointer to [**CharacterPermissionConfig**](CharacterPermissionConfig.md) |  | [optional] 
**Versions** | Pointer to **[]int32** | 版本列表 | [optional] 
**Manageable** | Pointer to **bool** | 当前用户是否可管理 | [optional] 
**Creator** | Pointer to **string** | 创建者 | [optional] 
**IsPreConfigured** | Pointer to **bool** | 是否官方创建 | [optional] 
**ExtractMemory** | Pointer to **bool** | 是否抽取记忆 | [optional] 
**ExtractMoment** | Pointer to **bool** | 是否抽取瞬间 | [optional] 

## Methods

### NewCharacterDTO

`func NewCharacterDTO(name string, introduction string, basicInformation string, openingLine string, ) *CharacterDTO`

NewCharacterDTO instantiates a new CharacterDTO object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCharacterDTOWithDefaults

`func NewCharacterDTOWithDefaults() *CharacterDTO`

NewCharacterDTOWithDefaults instantiates a new CharacterDTO object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetGatewayAddContent

`func (o *CharacterDTO) GetGatewayAddContent() GatewayIssuedParams`

GetGatewayAddContent returns the GatewayAddContent field if non-nil, zero value otherwise.

### GetGatewayAddContentOk

`func (o *CharacterDTO) GetGatewayAddContentOk() (*GatewayIssuedParams, bool)`

GetGatewayAddContentOk returns a tuple with the GatewayAddContent field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGatewayAddContent

`func (o *CharacterDTO) SetGatewayAddContent(v GatewayIssuedParams)`

SetGatewayAddContent sets GatewayAddContent field to given value.

### HasGatewayAddContent

`func (o *CharacterDTO) HasGatewayAddContent() bool`

HasGatewayAddContent returns a boolean if a field has been set.

### GetName

`func (o *CharacterDTO) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *CharacterDTO) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *CharacterDTO) SetName(v string)`

SetName sets Name field to given value.


### GetAvatar

`func (o *CharacterDTO) GetAvatar() FileInfoVO`

GetAvatar returns the Avatar field if non-nil, zero value otherwise.

### GetAvatarOk

`func (o *CharacterDTO) GetAvatarOk() (*FileInfoVO, bool)`

GetAvatarOk returns a tuple with the Avatar field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAvatar

`func (o *CharacterDTO) SetAvatar(v FileInfoVO)`

SetAvatar sets Avatar field to given value.

### HasAvatar

`func (o *CharacterDTO) HasAvatar() bool`

HasAvatar returns a boolean if a field has been set.

### GetIntroduction

`func (o *CharacterDTO) GetIntroduction() string`

GetIntroduction returns the Introduction field if non-nil, zero value otherwise.

### GetIntroductionOk

`func (o *CharacterDTO) GetIntroductionOk() (*string, bool)`

GetIntroductionOk returns a tuple with the Introduction field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIntroduction

`func (o *CharacterDTO) SetIntroduction(v string)`

SetIntroduction sets Introduction field to given value.


### GetBasicInformation

`func (o *CharacterDTO) GetBasicInformation() string`

GetBasicInformation returns the BasicInformation field if non-nil, zero value otherwise.

### GetBasicInformationOk

`func (o *CharacterDTO) GetBasicInformationOk() (*string, bool)`

GetBasicInformationOk returns a tuple with the BasicInformation field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBasicInformation

`func (o *CharacterDTO) SetBasicInformation(v string)`

SetBasicInformation sets BasicInformation field to given value.


### GetOpeningLine

`func (o *CharacterDTO) GetOpeningLine() string`

GetOpeningLine returns the OpeningLine field if non-nil, zero value otherwise.

### GetOpeningLineOk

`func (o *CharacterDTO) GetOpeningLineOk() (*string, bool)`

GetOpeningLineOk returns a tuple with the OpeningLine field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOpeningLine

`func (o *CharacterDTO) SetOpeningLine(v string)`

SetOpeningLine sets OpeningLine field to given value.


### GetTraits

`func (o *CharacterDTO) GetTraits() string`

GetTraits returns the Traits field if non-nil, zero value otherwise.

### GetTraitsOk

`func (o *CharacterDTO) GetTraitsOk() (*string, bool)`

GetTraitsOk returns a tuple with the Traits field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTraits

`func (o *CharacterDTO) SetTraits(v string)`

SetTraits sets Traits field to given value.

### HasTraits

`func (o *CharacterDTO) HasTraits() bool`

HasTraits returns a boolean if a field has been set.

### GetChatExample

`func (o *CharacterDTO) GetChatExample() string`

GetChatExample returns the ChatExample field if non-nil, zero value otherwise.

### GetChatExampleOk

`func (o *CharacterDTO) GetChatExampleOk() (*string, bool)`

GetChatExampleOk returns a tuple with the ChatExample field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetChatExample

`func (o *CharacterDTO) SetChatExample(v string)`

SetChatExample sets ChatExample field to given value.

### HasChatExample

`func (o *CharacterDTO) HasChatExample() bool`

HasChatExample returns a boolean if a field has been set.

### GetType

`func (o *CharacterDTO) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *CharacterDTO) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *CharacterDTO) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *CharacterDTO) HasType() bool`

HasType returns a boolean if a field has been set.

### GetChatObjective

`func (o *CharacterDTO) GetChatObjective() string`

GetChatObjective returns the ChatObjective field if non-nil, zero value otherwise.

### GetChatObjectiveOk

`func (o *CharacterDTO) GetChatObjectiveOk() (*string, bool)`

GetChatObjectiveOk returns a tuple with the ChatObjective field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetChatObjective

`func (o *CharacterDTO) SetChatObjective(v string)`

SetChatObjective sets ChatObjective field to given value.

### HasChatObjective

`func (o *CharacterDTO) HasChatObjective() bool`

HasChatObjective returns a boolean if a field has been set.

### GetAdvancedConfig

`func (o *CharacterDTO) GetAdvancedConfig() CharacterAdvancedConfig`

GetAdvancedConfig returns the AdvancedConfig field if non-nil, zero value otherwise.

### GetAdvancedConfigOk

`func (o *CharacterDTO) GetAdvancedConfigOk() (*CharacterAdvancedConfig, bool)`

GetAdvancedConfigOk returns a tuple with the AdvancedConfig field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAdvancedConfig

`func (o *CharacterDTO) SetAdvancedConfig(v CharacterAdvancedConfig)`

SetAdvancedConfig sets AdvancedConfig field to given value.

### HasAdvancedConfig

`func (o *CharacterDTO) HasAdvancedConfig() bool`

HasAdvancedConfig returns a boolean if a field has been set.

### GetCharacterId

`func (o *CharacterDTO) GetCharacterId() string`

GetCharacterId returns the CharacterId field if non-nil, zero value otherwise.

### GetCharacterIdOk

`func (o *CharacterDTO) GetCharacterIdOk() (*string, bool)`

GetCharacterIdOk returns a tuple with the CharacterId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCharacterId

`func (o *CharacterDTO) SetCharacterId(v string)`

SetCharacterId sets CharacterId field to given value.

### HasCharacterId

`func (o *CharacterDTO) HasCharacterId() bool`

HasCharacterId returns a boolean if a field has been set.

### GetVersion

`func (o *CharacterDTO) GetVersion() int32`

GetVersion returns the Version field if non-nil, zero value otherwise.

### GetVersionOk

`func (o *CharacterDTO) GetVersionOk() (*int32, bool)`

GetVersionOk returns a tuple with the Version field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVersion

`func (o *CharacterDTO) SetVersion(v int32)`

SetVersion sets Version field to given value.

### HasVersion

`func (o *CharacterDTO) HasVersion() bool`

HasVersion returns a boolean if a field has been set.

### GetMajorVersion

`func (o *CharacterDTO) GetMajorVersion() int32`

GetMajorVersion returns the MajorVersion field if non-nil, zero value otherwise.

### GetMajorVersionOk

`func (o *CharacterDTO) GetMajorVersionOk() (*int32, bool)`

GetMajorVersionOk returns a tuple with the MajorVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMajorVersion

`func (o *CharacterDTO) SetMajorVersion(v int32)`

SetMajorVersion sets MajorVersion field to given value.

### HasMajorVersion

`func (o *CharacterDTO) HasMajorVersion() bool`

HasMajorVersion returns a boolean if a field has been set.

### GetGmtCreate

`func (o *CharacterDTO) GetGmtCreate() int64`

GetGmtCreate returns the GmtCreate field if non-nil, zero value otherwise.

### GetGmtCreateOk

`func (o *CharacterDTO) GetGmtCreateOk() (*int64, bool)`

GetGmtCreateOk returns a tuple with the GmtCreate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGmtCreate

`func (o *CharacterDTO) SetGmtCreate(v int64)`

SetGmtCreate sets GmtCreate field to given value.

### HasGmtCreate

`func (o *CharacterDTO) HasGmtCreate() bool`

HasGmtCreate returns a boolean if a field has been set.

### GetGmtModified

`func (o *CharacterDTO) GetGmtModified() int64`

GetGmtModified returns the GmtModified field if non-nil, zero value otherwise.

### GetGmtModifiedOk

`func (o *CharacterDTO) GetGmtModifiedOk() (*int64, bool)`

GetGmtModifiedOk returns a tuple with the GmtModified field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGmtModified

`func (o *CharacterDTO) SetGmtModified(v int64)`

SetGmtModified sets GmtModified field to given value.

### HasGmtModified

`func (o *CharacterDTO) HasGmtModified() bool`

HasGmtModified returns a boolean if a field has been set.

### GetUserId

`func (o *CharacterDTO) GetUserId() string`

GetUserId returns the UserId field if non-nil, zero value otherwise.

### GetUserIdOk

`func (o *CharacterDTO) GetUserIdOk() (*string, bool)`

GetUserIdOk returns a tuple with the UserId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserId

`func (o *CharacterDTO) SetUserId(v string)`

SetUserId sets UserId field to given value.

### HasUserId

`func (o *CharacterDTO) HasUserId() bool`

HasUserId returns a boolean if a field has been set.

### GetModifiedBy

`func (o *CharacterDTO) GetModifiedBy() string`

GetModifiedBy returns the ModifiedBy field if non-nil, zero value otherwise.

### GetModifiedByOk

`func (o *CharacterDTO) GetModifiedByOk() (*string, bool)`

GetModifiedByOk returns a tuple with the ModifiedBy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetModifiedBy

`func (o *CharacterDTO) SetModifiedBy(v string)`

SetModifiedBy sets ModifiedBy field to given value.

### HasModifiedBy

`func (o *CharacterDTO) HasModifiedBy() bool`

HasModifiedBy returns a boolean if a field has been set.

### GetTags

`func (o *CharacterDTO) GetTags() []string`

GetTags returns the Tags field if non-nil, zero value otherwise.

### GetTagsOk

`func (o *CharacterDTO) GetTagsOk() (*[]string, bool)`

GetTagsOk returns a tuple with the Tags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTags

`func (o *CharacterDTO) SetTags(v []string)`

SetTags sets Tags field to given value.

### HasTags

`func (o *CharacterDTO) HasTags() bool`

HasTags returns a boolean if a field has been set.

### GetBubbles

`func (o *CharacterDTO) GetBubbles() []string`

GetBubbles returns the Bubbles field if non-nil, zero value otherwise.

### GetBubblesOk

`func (o *CharacterDTO) GetBubblesOk() (*[]string, bool)`

GetBubblesOk returns a tuple with the Bubbles field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBubbles

`func (o *CharacterDTO) SetBubbles(v []string)`

SetBubbles sets Bubbles field to given value.

### HasBubbles

`func (o *CharacterDTO) HasBubbles() bool`

HasBubbles returns a boolean if a field has been set.

### GetBackgroundImgUrl

`func (o *CharacterDTO) GetBackgroundImgUrl() string`

GetBackgroundImgUrl returns the BackgroundImgUrl field if non-nil, zero value otherwise.

### GetBackgroundImgUrlOk

`func (o *CharacterDTO) GetBackgroundImgUrlOk() (*string, bool)`

GetBackgroundImgUrlOk returns a tuple with the BackgroundImgUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBackgroundImgUrl

`func (o *CharacterDTO) SetBackgroundImgUrl(v string)`

SetBackgroundImgUrl sets BackgroundImgUrl field to given value.

### HasBackgroundImgUrl

`func (o *CharacterDTO) HasBackgroundImgUrl() bool`

HasBackgroundImgUrl returns a boolean if a field has been set.

### GetPermConfig

`func (o *CharacterDTO) GetPermConfig() CharacterPermissionConfig`

GetPermConfig returns the PermConfig field if non-nil, zero value otherwise.

### GetPermConfigOk

`func (o *CharacterDTO) GetPermConfigOk() (*CharacterPermissionConfig, bool)`

GetPermConfigOk returns a tuple with the PermConfig field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPermConfig

`func (o *CharacterDTO) SetPermConfig(v CharacterPermissionConfig)`

SetPermConfig sets PermConfig field to given value.

### HasPermConfig

`func (o *CharacterDTO) HasPermConfig() bool`

HasPermConfig returns a boolean if a field has been set.

### GetVersions

`func (o *CharacterDTO) GetVersions() []int32`

GetVersions returns the Versions field if non-nil, zero value otherwise.

### GetVersionsOk

`func (o *CharacterDTO) GetVersionsOk() (*[]int32, bool)`

GetVersionsOk returns a tuple with the Versions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVersions

`func (o *CharacterDTO) SetVersions(v []int32)`

SetVersions sets Versions field to given value.

### HasVersions

`func (o *CharacterDTO) HasVersions() bool`

HasVersions returns a boolean if a field has been set.

### GetManageable

`func (o *CharacterDTO) GetManageable() bool`

GetManageable returns the Manageable field if non-nil, zero value otherwise.

### GetManageableOk

`func (o *CharacterDTO) GetManageableOk() (*bool, bool)`

GetManageableOk returns a tuple with the Manageable field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetManageable

`func (o *CharacterDTO) SetManageable(v bool)`

SetManageable sets Manageable field to given value.

### HasManageable

`func (o *CharacterDTO) HasManageable() bool`

HasManageable returns a boolean if a field has been set.

### GetCreator

`func (o *CharacterDTO) GetCreator() string`

GetCreator returns the Creator field if non-nil, zero value otherwise.

### GetCreatorOk

`func (o *CharacterDTO) GetCreatorOk() (*string, bool)`

GetCreatorOk returns a tuple with the Creator field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreator

`func (o *CharacterDTO) SetCreator(v string)`

SetCreator sets Creator field to given value.

### HasCreator

`func (o *CharacterDTO) HasCreator() bool`

HasCreator returns a boolean if a field has been set.

### GetIsPreConfigured

`func (o *CharacterDTO) GetIsPreConfigured() bool`

GetIsPreConfigured returns the IsPreConfigured field if non-nil, zero value otherwise.

### GetIsPreConfiguredOk

`func (o *CharacterDTO) GetIsPreConfiguredOk() (*bool, bool)`

GetIsPreConfiguredOk returns a tuple with the IsPreConfigured field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsPreConfigured

`func (o *CharacterDTO) SetIsPreConfigured(v bool)`

SetIsPreConfigured sets IsPreConfigured field to given value.

### HasIsPreConfigured

`func (o *CharacterDTO) HasIsPreConfigured() bool`

HasIsPreConfigured returns a boolean if a field has been set.

### GetExtractMemory

`func (o *CharacterDTO) GetExtractMemory() bool`

GetExtractMemory returns the ExtractMemory field if non-nil, zero value otherwise.

### GetExtractMemoryOk

`func (o *CharacterDTO) GetExtractMemoryOk() (*bool, bool)`

GetExtractMemoryOk returns a tuple with the ExtractMemory field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExtractMemory

`func (o *CharacterDTO) SetExtractMemory(v bool)`

SetExtractMemory sets ExtractMemory field to given value.

### HasExtractMemory

`func (o *CharacterDTO) HasExtractMemory() bool`

HasExtractMemory returns a boolean if a field has been set.

### GetExtractMoment

`func (o *CharacterDTO) GetExtractMoment() bool`

GetExtractMoment returns the ExtractMoment field if non-nil, zero value otherwise.

### GetExtractMomentOk

`func (o *CharacterDTO) GetExtractMomentOk() (*bool, bool)`

GetExtractMomentOk returns a tuple with the ExtractMoment field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExtractMoment

`func (o *CharacterDTO) SetExtractMoment(v bool)`

SetExtractMoment sets ExtractMoment field to given value.

### HasExtractMoment

`func (o *CharacterDTO) HasExtractMoment() bool`

HasExtractMoment returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


