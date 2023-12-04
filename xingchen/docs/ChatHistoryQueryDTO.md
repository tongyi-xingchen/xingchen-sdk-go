# ChatHistoryQueryDTO

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**GatewayAddContent** | Pointer to [**GatewayIssuedParams**](GatewayIssuedParams.md) |  | [optional] 
**Where** | Pointer to [**ChatHistoryQueryWhere**](ChatHistoryQueryWhere.md) |  | [optional] 
**OrderBy** | Pointer to **[]string** | 排序条件，支持的排序字段为： - gmtCreate 排序优先级按列表顺序，默认降序，如果期望升序，需要在字段后指定，如：orderBy: [\\\&quot;gmtCreate desc\\\&quot;] （gmtCreate 按降序）  | [optional] 
**PageNum** | Pointer to **int32** | 页码，默认为 0 | [optional] 
**PageSize** | Pointer to **int32** | 页条目数，1～50，默认为 10 | [optional] 

## Methods

### NewChatHistoryQueryDTO

`func NewChatHistoryQueryDTO() *ChatHistoryQueryDTO`

NewChatHistoryQueryDTO instantiates a new ChatHistoryQueryDTO object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewChatHistoryQueryDTOWithDefaults

`func NewChatHistoryQueryDTOWithDefaults() *ChatHistoryQueryDTO`

NewChatHistoryQueryDTOWithDefaults instantiates a new ChatHistoryQueryDTO object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetGatewayAddContent

`func (o *ChatHistoryQueryDTO) GetGatewayAddContent() GatewayIssuedParams`

GetGatewayAddContent returns the GatewayAddContent field if non-nil, zero value otherwise.

### GetGatewayAddContentOk

`func (o *ChatHistoryQueryDTO) GetGatewayAddContentOk() (*GatewayIssuedParams, bool)`

GetGatewayAddContentOk returns a tuple with the GatewayAddContent field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGatewayAddContent

`func (o *ChatHistoryQueryDTO) SetGatewayAddContent(v GatewayIssuedParams)`

SetGatewayAddContent sets GatewayAddContent field to given value.

### HasGatewayAddContent

`func (o *ChatHistoryQueryDTO) HasGatewayAddContent() bool`

HasGatewayAddContent returns a boolean if a field has been set.

### GetWhere

`func (o *ChatHistoryQueryDTO) GetWhere() ChatHistoryQueryWhere`

GetWhere returns the Where field if non-nil, zero value otherwise.

### GetWhereOk

`func (o *ChatHistoryQueryDTO) GetWhereOk() (*ChatHistoryQueryWhere, bool)`

GetWhereOk returns a tuple with the Where field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWhere

`func (o *ChatHistoryQueryDTO) SetWhere(v ChatHistoryQueryWhere)`

SetWhere sets Where field to given value.

### HasWhere

`func (o *ChatHistoryQueryDTO) HasWhere() bool`

HasWhere returns a boolean if a field has been set.

### GetOrderBy

`func (o *ChatHistoryQueryDTO) GetOrderBy() []string`

GetOrderBy returns the OrderBy field if non-nil, zero value otherwise.

### GetOrderByOk

`func (o *ChatHistoryQueryDTO) GetOrderByOk() (*[]string, bool)`

GetOrderByOk returns a tuple with the OrderBy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrderBy

`func (o *ChatHistoryQueryDTO) SetOrderBy(v []string)`

SetOrderBy sets OrderBy field to given value.

### HasOrderBy

`func (o *ChatHistoryQueryDTO) HasOrderBy() bool`

HasOrderBy returns a boolean if a field has been set.

### GetPageNum

`func (o *ChatHistoryQueryDTO) GetPageNum() int32`

GetPageNum returns the PageNum field if non-nil, zero value otherwise.

### GetPageNumOk

`func (o *ChatHistoryQueryDTO) GetPageNumOk() (*int32, bool)`

GetPageNumOk returns a tuple with the PageNum field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPageNum

`func (o *ChatHistoryQueryDTO) SetPageNum(v int32)`

SetPageNum sets PageNum field to given value.

### HasPageNum

`func (o *ChatHistoryQueryDTO) HasPageNum() bool`

HasPageNum returns a boolean if a field has been set.

### GetPageSize

`func (o *ChatHistoryQueryDTO) GetPageSize() int32`

GetPageSize returns the PageSize field if non-nil, zero value otherwise.

### GetPageSizeOk

`func (o *ChatHistoryQueryDTO) GetPageSizeOk() (*int32, bool)`

GetPageSizeOk returns a tuple with the PageSize field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPageSize

`func (o *ChatHistoryQueryDTO) SetPageSize(v int32)`

SetPageSize sets PageSize field to given value.

### HasPageSize

`func (o *ChatHistoryQueryDTO) HasPageSize() bool`

HasPageSize returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


