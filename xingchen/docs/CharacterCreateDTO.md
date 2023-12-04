# CharacterCreateDTO

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
**PermConfig** | Pointer to [**CharacterPermissionConfig**](CharacterPermissionConfig.md) |  | [optional] 

## Methods

### NewCharacterCreateDTO

`func NewCharacterCreateDTO(name string, introduction string, basicInformation string, openingLine string, ) *CharacterCreateDTO`

NewCharacterCreateDTO instantiates a new CharacterCreateDTO object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCharacterCreateDTOWithDefaults

`func NewCharacterCreateDTOWithDefaults() *CharacterCreateDTO`

NewCharacterCreateDTOWithDefaults instantiates a new CharacterCreateDTO object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetGatewayAddContent

`func (o *CharacterCreateDTO) GetGatewayAddContent() GatewayIssuedParams`

GetGatewayAddContent returns the GatewayAddContent field if non-nil, zero value otherwise.

### GetGatewayAddContentOk

`func (o *CharacterCreateDTO) GetGatewayAddContentOk() (*GatewayIssuedParams, bool)`

GetGatewayAddContentOk returns a tuple with the GatewayAddContent field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGatewayAddContent

`func (o *CharacterCreateDTO) SetGatewayAddContent(v GatewayIssuedParams)`

SetGatewayAddContent sets GatewayAddContent field to given value.

### HasGatewayAddContent

`func (o *CharacterCreateDTO) HasGatewayAddContent() bool`

HasGatewayAddContent returns a boolean if a field has been set.

### GetName

`func (o *CharacterCreateDTO) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *CharacterCreateDTO) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *CharacterCreateDTO) SetName(v string)`

SetName sets Name field to given value.


### GetAvatar

`func (o *CharacterCreateDTO) GetAvatar() FileInfoVO`

GetAvatar returns the Avatar field if non-nil, zero value otherwise.

### GetAvatarOk

`func (o *CharacterCreateDTO) GetAvatarOk() (*FileInfoVO, bool)`

GetAvatarOk returns a tuple with the Avatar field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAvatar

`func (o *CharacterCreateDTO) SetAvatar(v FileInfoVO)`

SetAvatar sets Avatar field to given value.

### HasAvatar

`func (o *CharacterCreateDTO) HasAvatar() bool`

HasAvatar returns a boolean if a field has been set.

### GetIntroduction

`func (o *CharacterCreateDTO) GetIntroduction() string`

GetIntroduction returns the Introduction field if non-nil, zero value otherwise.

### GetIntroductionOk

`func (o *CharacterCreateDTO) GetIntroductionOk() (*string, bool)`

GetIntroductionOk returns a tuple with the Introduction field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIntroduction

`func (o *CharacterCreateDTO) SetIntroduction(v string)`

SetIntroduction sets Introduction field to given value.


### GetBasicInformation

`func (o *CharacterCreateDTO) GetBasicInformation() string`

GetBasicInformation returns the BasicInformation field if non-nil, zero value otherwise.

### GetBasicInformationOk

`func (o *CharacterCreateDTO) GetBasicInformationOk() (*string, bool)`

GetBasicInformationOk returns a tuple with the BasicInformation field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBasicInformation

`func (o *CharacterCreateDTO) SetBasicInformation(v string)`

SetBasicInformation sets BasicInformation field to given value.


### GetOpeningLine

`func (o *CharacterCreateDTO) GetOpeningLine() string`

GetOpeningLine returns the OpeningLine field if non-nil, zero value otherwise.

### GetOpeningLineOk

`func (o *CharacterCreateDTO) GetOpeningLineOk() (*string, bool)`

GetOpeningLineOk returns a tuple with the OpeningLine field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOpeningLine

`func (o *CharacterCreateDTO) SetOpeningLine(v string)`

SetOpeningLine sets OpeningLine field to given value.


### GetTraits

`func (o *CharacterCreateDTO) GetTraits() string`

GetTraits returns the Traits field if non-nil, zero value otherwise.

### GetTraitsOk

`func (o *CharacterCreateDTO) GetTraitsOk() (*string, bool)`

GetTraitsOk returns a tuple with the Traits field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTraits

`func (o *CharacterCreateDTO) SetTraits(v string)`

SetTraits sets Traits field to given value.

### HasTraits

`func (o *CharacterCreateDTO) HasTraits() bool`

HasTraits returns a boolean if a field has been set.

### GetChatExample

`func (o *CharacterCreateDTO) GetChatExample() string`

GetChatExample returns the ChatExample field if non-nil, zero value otherwise.

### GetChatExampleOk

`func (o *CharacterCreateDTO) GetChatExampleOk() (*string, bool)`

GetChatExampleOk returns a tuple with the ChatExample field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetChatExample

`func (o *CharacterCreateDTO) SetChatExample(v string)`

SetChatExample sets ChatExample field to given value.

### HasChatExample

`func (o *CharacterCreateDTO) HasChatExample() bool`

HasChatExample returns a boolean if a field has been set.

### GetType

`func (o *CharacterCreateDTO) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *CharacterCreateDTO) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *CharacterCreateDTO) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *CharacterCreateDTO) HasType() bool`

HasType returns a boolean if a field has been set.

### GetChatObjective

`func (o *CharacterCreateDTO) GetChatObjective() string`

GetChatObjective returns the ChatObjective field if non-nil, zero value otherwise.

### GetChatObjectiveOk

`func (o *CharacterCreateDTO) GetChatObjectiveOk() (*string, bool)`

GetChatObjectiveOk returns a tuple with the ChatObjective field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetChatObjective

`func (o *CharacterCreateDTO) SetChatObjective(v string)`

SetChatObjective sets ChatObjective field to given value.

### HasChatObjective

`func (o *CharacterCreateDTO) HasChatObjective() bool`

HasChatObjective returns a boolean if a field has been set.

### GetAdvancedConfig

`func (o *CharacterCreateDTO) GetAdvancedConfig() CharacterAdvancedConfig`

GetAdvancedConfig returns the AdvancedConfig field if non-nil, zero value otherwise.

### GetAdvancedConfigOk

`func (o *CharacterCreateDTO) GetAdvancedConfigOk() (*CharacterAdvancedConfig, bool)`

GetAdvancedConfigOk returns a tuple with the AdvancedConfig field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAdvancedConfig

`func (o *CharacterCreateDTO) SetAdvancedConfig(v CharacterAdvancedConfig)`

SetAdvancedConfig sets AdvancedConfig field to given value.

### HasAdvancedConfig

`func (o *CharacterCreateDTO) HasAdvancedConfig() bool`

HasAdvancedConfig returns a boolean if a field has been set.

### GetPermConfig

`func (o *CharacterCreateDTO) GetPermConfig() CharacterPermissionConfig`

GetPermConfig returns the PermConfig field if non-nil, zero value otherwise.

### GetPermConfigOk

`func (o *CharacterCreateDTO) GetPermConfigOk() (*CharacterPermissionConfig, bool)`

GetPermConfigOk returns a tuple with the PermConfig field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPermConfig

`func (o *CharacterCreateDTO) SetPermConfig(v CharacterPermissionConfig)`

SetPermConfig sets PermConfig field to given value.

### HasPermConfig

`func (o *CharacterCreateDTO) HasPermConfig() bool`

HasPermConfig returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


