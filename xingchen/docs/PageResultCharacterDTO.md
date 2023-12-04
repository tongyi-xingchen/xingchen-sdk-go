# PageResultCharacterDTO

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**List** | Pointer to [**[]CharacterDTO**](CharacterDTO.md) |  | [optional] 
**Page** | Pointer to **int32** |  | [optional] 
**PageSize** | Pointer to **int32** |  | [optional] 
**Total** | Pointer to **int32** |  | [optional] 
**PageOffsetValue** | Pointer to **int32** |  | [optional] 

## Methods

### NewPageResultCharacterDTO

`func NewPageResultCharacterDTO() *PageResultCharacterDTO`

NewPageResultCharacterDTO instantiates a new PageResultCharacterDTO object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPageResultCharacterDTOWithDefaults

`func NewPageResultCharacterDTOWithDefaults() *PageResultCharacterDTO`

NewPageResultCharacterDTOWithDefaults instantiates a new PageResultCharacterDTO object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetList

`func (o *PageResultCharacterDTO) GetList() []CharacterDTO`

GetList returns the List field if non-nil, zero value otherwise.

### GetListOk

`func (o *PageResultCharacterDTO) GetListOk() (*[]CharacterDTO, bool)`

GetListOk returns a tuple with the List field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetList

`func (o *PageResultCharacterDTO) SetList(v []CharacterDTO)`

SetList sets List field to given value.

### HasList

`func (o *PageResultCharacterDTO) HasList() bool`

HasList returns a boolean if a field has been set.

### GetPage

`func (o *PageResultCharacterDTO) GetPage() int32`

GetPage returns the Page field if non-nil, zero value otherwise.

### GetPageOk

`func (o *PageResultCharacterDTO) GetPageOk() (*int32, bool)`

GetPageOk returns a tuple with the Page field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPage

`func (o *PageResultCharacterDTO) SetPage(v int32)`

SetPage sets Page field to given value.

### HasPage

`func (o *PageResultCharacterDTO) HasPage() bool`

HasPage returns a boolean if a field has been set.

### GetPageSize

`func (o *PageResultCharacterDTO) GetPageSize() int32`

GetPageSize returns the PageSize field if non-nil, zero value otherwise.

### GetPageSizeOk

`func (o *PageResultCharacterDTO) GetPageSizeOk() (*int32, bool)`

GetPageSizeOk returns a tuple with the PageSize field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPageSize

`func (o *PageResultCharacterDTO) SetPageSize(v int32)`

SetPageSize sets PageSize field to given value.

### HasPageSize

`func (o *PageResultCharacterDTO) HasPageSize() bool`

HasPageSize returns a boolean if a field has been set.

### GetTotal

`func (o *PageResultCharacterDTO) GetTotal() int32`

GetTotal returns the Total field if non-nil, zero value otherwise.

### GetTotalOk

`func (o *PageResultCharacterDTO) GetTotalOk() (*int32, bool)`

GetTotalOk returns a tuple with the Total field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotal

`func (o *PageResultCharacterDTO) SetTotal(v int32)`

SetTotal sets Total field to given value.

### HasTotal

`func (o *PageResultCharacterDTO) HasTotal() bool`

HasTotal returns a boolean if a field has been set.

### GetPageOffsetValue

`func (o *PageResultCharacterDTO) GetPageOffsetValue() int32`

GetPageOffsetValue returns the PageOffsetValue field if non-nil, zero value otherwise.

### GetPageOffsetValueOk

`func (o *PageResultCharacterDTO) GetPageOffsetValueOk() (*int32, bool)`

GetPageOffsetValueOk returns a tuple with the PageOffsetValue field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPageOffsetValue

`func (o *PageResultCharacterDTO) SetPageOffsetValue(v int32)`

SetPageOffsetValue sets PageOffsetValue field to given value.

### HasPageOffsetValue

`func (o *PageResultCharacterDTO) HasPageOffsetValue() bool`

HasPageOffsetValue returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


