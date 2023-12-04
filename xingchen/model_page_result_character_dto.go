/*
XingChen 开放接口定义

No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)

API version: v1
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package xingchen

import (
	"encoding/json"
)

// checks if the PageResultCharacterDTO type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &PageResultCharacterDTO{}

// PageResultCharacterDTO struct for PageResultCharacterDTO
type PageResultCharacterDTO struct {
	List                 []CharacterDTO `json:"list,omitempty"`
	Page                 *int32         `json:"page,omitempty"`
	PageSize             *int32         `json:"pageSize,omitempty"`
	Total                *int32         `json:"total,omitempty"`
	PageOffsetValue      *int32         `json:"pageOffsetValue,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _PageResultCharacterDTO PageResultCharacterDTO

// NewPageResultCharacterDTO instantiates a new PageResultCharacterDTO object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPageResultCharacterDTO() *PageResultCharacterDTO {
	this := PageResultCharacterDTO{}
	return &this
}

// NewPageResultCharacterDTOWithDefaults instantiates a new PageResultCharacterDTO object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPageResultCharacterDTOWithDefaults() *PageResultCharacterDTO {
	this := PageResultCharacterDTO{}
	return &this
}

// GetList returns the List field value if set, zero value otherwise.
func (o *PageResultCharacterDTO) GetList() []CharacterDTO {
	if o == nil || IsNil(o.List) {
		var ret []CharacterDTO
		return ret
	}
	return o.List
}

// GetListOk returns a tuple with the List field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PageResultCharacterDTO) GetListOk() ([]CharacterDTO, bool) {
	if o == nil || IsNil(o.List) {
		return nil, false
	}
	return o.List, true
}

// HasList returns a boolean if a field has been set.
func (o *PageResultCharacterDTO) HasList() bool {
	if o != nil && !IsNil(o.List) {
		return true
	}

	return false
}

// SetList gets a reference to the given []CharacterDTO and assigns it to the List field.
func (o *PageResultCharacterDTO) SetList(v []CharacterDTO) {
	o.List = v
}

// GetPage returns the Page field value if set, zero value otherwise.
func (o *PageResultCharacterDTO) GetPage() int32 {
	if o == nil || IsNil(o.Page) {
		var ret int32
		return ret
	}
	return *o.Page
}

// GetPageOk returns a tuple with the Page field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PageResultCharacterDTO) GetPageOk() (*int32, bool) {
	if o == nil || IsNil(o.Page) {
		return nil, false
	}
	return o.Page, true
}

// HasPage returns a boolean if a field has been set.
func (o *PageResultCharacterDTO) HasPage() bool {
	if o != nil && !IsNil(o.Page) {
		return true
	}

	return false
}

// SetPage gets a reference to the given int32 and assigns it to the Page field.
func (o *PageResultCharacterDTO) SetPage(v int32) {
	o.Page = &v
}

// GetPageSize returns the PageSize field value if set, zero value otherwise.
func (o *PageResultCharacterDTO) GetPageSize() int32 {
	if o == nil || IsNil(o.PageSize) {
		var ret int32
		return ret
	}
	return *o.PageSize
}

// GetPageSizeOk returns a tuple with the PageSize field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PageResultCharacterDTO) GetPageSizeOk() (*int32, bool) {
	if o == nil || IsNil(o.PageSize) {
		return nil, false
	}
	return o.PageSize, true
}

// HasPageSize returns a boolean if a field has been set.
func (o *PageResultCharacterDTO) HasPageSize() bool {
	if o != nil && !IsNil(o.PageSize) {
		return true
	}

	return false
}

// SetPageSize gets a reference to the given int32 and assigns it to the PageSize field.
func (o *PageResultCharacterDTO) SetPageSize(v int32) {
	o.PageSize = &v
}

// GetTotal returns the Total field value if set, zero value otherwise.
func (o *PageResultCharacterDTO) GetTotal() int32 {
	if o == nil || IsNil(o.Total) {
		var ret int32
		return ret
	}
	return *o.Total
}

// GetTotalOk returns a tuple with the Total field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PageResultCharacterDTO) GetTotalOk() (*int32, bool) {
	if o == nil || IsNil(o.Total) {
		return nil, false
	}
	return o.Total, true
}

// HasTotal returns a boolean if a field has been set.
func (o *PageResultCharacterDTO) HasTotal() bool {
	if o != nil && !IsNil(o.Total) {
		return true
	}

	return false
}

// SetTotal gets a reference to the given int32 and assigns it to the Total field.
func (o *PageResultCharacterDTO) SetTotal(v int32) {
	o.Total = &v
}

// GetPageOffsetValue returns the PageOffsetValue field value if set, zero value otherwise.
func (o *PageResultCharacterDTO) GetPageOffsetValue() int32 {
	if o == nil || IsNil(o.PageOffsetValue) {
		var ret int32
		return ret
	}
	return *o.PageOffsetValue
}

// GetPageOffsetValueOk returns a tuple with the PageOffsetValue field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PageResultCharacterDTO) GetPageOffsetValueOk() (*int32, bool) {
	if o == nil || IsNil(o.PageOffsetValue) {
		return nil, false
	}
	return o.PageOffsetValue, true
}

// HasPageOffsetValue returns a boolean if a field has been set.
func (o *PageResultCharacterDTO) HasPageOffsetValue() bool {
	if o != nil && !IsNil(o.PageOffsetValue) {
		return true
	}

	return false
}

// SetPageOffsetValue gets a reference to the given int32 and assigns it to the PageOffsetValue field.
func (o *PageResultCharacterDTO) SetPageOffsetValue(v int32) {
	o.PageOffsetValue = &v
}

func (o PageResultCharacterDTO) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o PageResultCharacterDTO) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.List) {
		toSerialize["list"] = o.List
	}
	if !IsNil(o.Page) {
		toSerialize["page"] = o.Page
	}
	if !IsNil(o.PageSize) {
		toSerialize["pageSize"] = o.PageSize
	}
	if !IsNil(o.Total) {
		toSerialize["total"] = o.Total
	}
	if !IsNil(o.PageOffsetValue) {
		toSerialize["pageOffsetValue"] = o.PageOffsetValue
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *PageResultCharacterDTO) UnmarshalJSON(bytes []byte) (err error) {
	varPageResultCharacterDTO := _PageResultCharacterDTO{}

	err = json.Unmarshal(bytes, &varPageResultCharacterDTO)

	if err != nil {
		return err
	}

	*o = PageResultCharacterDTO(varPageResultCharacterDTO)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "list")
		delete(additionalProperties, "page")
		delete(additionalProperties, "pageSize")
		delete(additionalProperties, "total")
		delete(additionalProperties, "pageOffsetValue")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullablePageResultCharacterDTO struct {
	value *PageResultCharacterDTO
	isSet bool
}

func (v NullablePageResultCharacterDTO) Get() *PageResultCharacterDTO {
	return v.value
}

func (v *NullablePageResultCharacterDTO) Set(val *PageResultCharacterDTO) {
	v.value = val
	v.isSet = true
}

func (v NullablePageResultCharacterDTO) IsSet() bool {
	return v.isSet
}

func (v *NullablePageResultCharacterDTO) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePageResultCharacterDTO(val *PageResultCharacterDTO) *NullablePageResultCharacterDTO {
	return &NullablePageResultCharacterDTO{value: val, isSet: true}
}

func (v NullablePageResultCharacterDTO) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePageResultCharacterDTO) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}