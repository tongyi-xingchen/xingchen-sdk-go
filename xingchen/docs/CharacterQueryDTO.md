# CharacterQueryDTO

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**GatewayAddContent** | Pointer to [**GatewayIssuedParams**](GatewayIssuedParams.md) |  | [optional] 
**Where** | Pointer to [**CharacterQueryWhere**](CharacterQueryWhere.md) |  | [optional] 
**OrderBy** | Pointer to **[]string** | 排序条件，支持的排序字段为： - modifiedTime - createTime 排序优先级按列表顺序，默认降序，如果期望升序，需要在字段后指定，如：orderBy: [\\\&quot;gmtModified desc\\\&quot;] （gmtModified 按降序）  | [optional] 
**PageNum** | Pointer to **int64** | 页码，默认为 0 | [optional] 
**PageSize** | Pointer to **int64** | 页条目数，1～50，默认为 10 | [optional] 

## Methods

### NewCharacterQueryDTO

`func NewCharacterQueryDTO() *CharacterQueryDTO`

NewCharacterQueryDTO instantiates a new CharacterQueryDTO object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCharacterQueryDTOWithDefaults

`func NewCharacterQueryDTOWithDefaults() *CharacterQueryDTO`

NewCharacterQueryDTOWithDefaults instantiates a new CharacterQueryDTO object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetGatewayAddContent

`func (o *CharacterQueryDTO) GetGatewayAddContent() GatewayIssuedParams`

GetGatewayAddContent returns the GatewayAddContent field if non-nil, zero value otherwise.

### GetGatewayAddContentOk

`func (o *CharacterQueryDTO) GetGatewayAddContentOk() (*GatewayIssuedParams, bool)`

GetGatewayAddContentOk returns a tuple with the GatewayAddContent field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGatewayAddContent

`func (o *CharacterQueryDTO) SetGatewayAddContent(v GatewayIssuedParams)`

SetGatewayAddContent sets GatewayAddContent field to given value.

### HasGatewayAddContent

`func (o *CharacterQueryDTO) HasGatewayAddContent() bool`

HasGatewayAddContent returns a boolean if a field has been set.

### GetWhere

`func (o *CharacterQueryDTO) GetWhere() CharacterQueryWhere`

GetWhere returns the Where field if non-nil, zero value otherwise.

### GetWhereOk

`func (o *CharacterQueryDTO) GetWhereOk() (*CharacterQueryWhere, bool)`

GetWhereOk returns a tuple with the Where field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWhere

`func (o *CharacterQueryDTO) SetWhere(v CharacterQueryWhere)`

SetWhere sets Where field to given value.

### HasWhere

`func (o *CharacterQueryDTO) HasWhere() bool`

HasWhere returns a boolean if a field has been set.

### GetOrderBy

`func (o *CharacterQueryDTO) GetOrderBy() []string`

GetOrderBy returns the OrderBy field if non-nil, zero value otherwise.

### GetOrderByOk

`func (o *CharacterQueryDTO) GetOrderByOk() (*[]string, bool)`

GetOrderByOk returns a tuple with the OrderBy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrderBy

`func (o *CharacterQueryDTO) SetOrderBy(v []string)`

SetOrderBy sets OrderBy field to given value.

### HasOrderBy

`func (o *CharacterQueryDTO) HasOrderBy() bool`

HasOrderBy returns a boolean if a field has been set.

### GetPageNum

`func (o *CharacterQueryDTO) GetPageNum() int64`

GetPageNum returns the PageNum field if non-nil, zero value otherwise.

### GetPageNumOk

`func (o *CharacterQueryDTO) GetPageNumOk() (*int64, bool)`

GetPageNumOk returns a tuple with the PageNum field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPageNum

`func (o *CharacterQueryDTO) SetPageNum(v int64)`

SetPageNum sets PageNum field to given value.

### HasPageNum

`func (o *CharacterQueryDTO) HasPageNum() bool`

HasPageNum returns a boolean if a field has been set.

### GetPageSize

`func (o *CharacterQueryDTO) GetPageSize() int64`

GetPageSize returns the PageSize field if non-nil, zero value otherwise.

### GetPageSizeOk

`func (o *CharacterQueryDTO) GetPageSizeOk() (*int64, bool)`

GetPageSizeOk returns a tuple with the PageSize field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPageSize

`func (o *CharacterQueryDTO) SetPageSize(v int64)`

SetPageSize sets PageSize field to given value.

### HasPageSize

`func (o *CharacterQueryDTO) HasPageSize() bool`

HasPageSize returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


