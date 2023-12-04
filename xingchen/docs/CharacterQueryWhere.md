# CharacterQueryWhere

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CharacterName** | Pointer to **string** | 角色名称 | [optional] 
**Scope** | Pointer to **string** | 查询范围：my:我创建的角色, public: 平台开放的角色, pre_configured: 预制角色 | [optional] 

## Methods

### NewCharacterQueryWhere

`func NewCharacterQueryWhere() *CharacterQueryWhere`

NewCharacterQueryWhere instantiates a new CharacterQueryWhere object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCharacterQueryWhereWithDefaults

`func NewCharacterQueryWhereWithDefaults() *CharacterQueryWhere`

NewCharacterQueryWhereWithDefaults instantiates a new CharacterQueryWhere object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCharacterName

`func (o *CharacterQueryWhere) GetCharacterName() string`

GetCharacterName returns the CharacterName field if non-nil, zero value otherwise.

### GetCharacterNameOk

`func (o *CharacterQueryWhere) GetCharacterNameOk() (*string, bool)`

GetCharacterNameOk returns a tuple with the CharacterName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCharacterName

`func (o *CharacterQueryWhere) SetCharacterName(v string)`

SetCharacterName sets CharacterName field to given value.

### HasCharacterName

`func (o *CharacterQueryWhere) HasCharacterName() bool`

HasCharacterName returns a boolean if a field has been set.

### GetScope

`func (o *CharacterQueryWhere) GetScope() string`

GetScope returns the Scope field if non-nil, zero value otherwise.

### GetScopeOk

`func (o *CharacterQueryWhere) GetScopeOk() (*string, bool)`

GetScopeOk returns a tuple with the Scope field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScope

`func (o *CharacterQueryWhere) SetScope(v string)`

SetScope sets Scope field to given value.

### HasScope

`func (o *CharacterQueryWhere) HasScope() bool`

HasScope returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


