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

// checks if the CharacterPermissionConfig type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &CharacterPermissionConfig{}

// CharacterPermissionConfig 角色权限配置
type CharacterPermissionConfig struct {
	// 是否可查看角色属性(0:不公开;1:公开)
	IsPublic int32 `json:"isPublic"`
	// 允许在平台和其他用户聊天(0:不允许,1:允许)
	AllowChat int32 `json:"allowChat"`
	// 允许其他用户通过api调用(0:不允许,1:允许)
	AllowApi             int32 `json:"allowApi"`
	AdditionalProperties map[string]interface{}
}

type _CharacterPermissionConfig CharacterPermissionConfig

// NewCharacterPermissionConfig instantiates a new CharacterPermissionConfig object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCharacterPermissionConfig(isPublic int32, allowChat int32, allowApi int32) *CharacterPermissionConfig {
	this := CharacterPermissionConfig{}
	this.IsPublic = isPublic
	this.AllowChat = allowChat
	this.AllowApi = allowApi
	return &this
}

// NewCharacterPermissionConfigWithDefaults instantiates a new CharacterPermissionConfig object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCharacterPermissionConfigWithDefaults() *CharacterPermissionConfig {
	this := CharacterPermissionConfig{}
	return &this
}

// GetIsPublic returns the IsPublic field value
func (o *CharacterPermissionConfig) GetIsPublic() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.IsPublic
}

// GetIsPublicOk returns a tuple with the IsPublic field value
// and a boolean to check if the value has been set.
func (o *CharacterPermissionConfig) GetIsPublicOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.IsPublic, true
}

// SetIsPublic sets field value
func (o *CharacterPermissionConfig) SetIsPublic(v int32) {
	o.IsPublic = v
}

// GetAllowChat returns the AllowChat field value
func (o *CharacterPermissionConfig) GetAllowChat() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.AllowChat
}

// GetAllowChatOk returns a tuple with the AllowChat field value
// and a boolean to check if the value has been set.
func (o *CharacterPermissionConfig) GetAllowChatOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.AllowChat, true
}

// SetAllowChat sets field value
func (o *CharacterPermissionConfig) SetAllowChat(v int32) {
	o.AllowChat = v
}

// GetAllowApi returns the AllowApi field value
func (o *CharacterPermissionConfig) GetAllowApi() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.AllowApi
}

// GetAllowApiOk returns a tuple with the AllowApi field value
// and a boolean to check if the value has been set.
func (o *CharacterPermissionConfig) GetAllowApiOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.AllowApi, true
}

// SetAllowApi sets field value
func (o *CharacterPermissionConfig) SetAllowApi(v int32) {
	o.AllowApi = v
}

func (o CharacterPermissionConfig) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o CharacterPermissionConfig) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["isPublic"] = o.IsPublic
	toSerialize["allowChat"] = o.AllowChat
	toSerialize["allowApi"] = o.AllowApi

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *CharacterPermissionConfig) UnmarshalJSON(bytes []byte) (err error) {
	varCharacterPermissionConfig := _CharacterPermissionConfig{}

	err = json.Unmarshal(bytes, &varCharacterPermissionConfig)

	if err != nil {
		return err
	}

	*o = CharacterPermissionConfig(varCharacterPermissionConfig)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "isPublic")
		delete(additionalProperties, "allowChat")
		delete(additionalProperties, "allowApi")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableCharacterPermissionConfig struct {
	value *CharacterPermissionConfig
	isSet bool
}

func (v NullableCharacterPermissionConfig) Get() *CharacterPermissionConfig {
	return v.value
}

func (v *NullableCharacterPermissionConfig) Set(val *CharacterPermissionConfig) {
	v.value = val
	v.isSet = true
}

func (v NullableCharacterPermissionConfig) IsSet() bool {
	return v.isSet
}

func (v *NullableCharacterPermissionConfig) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCharacterPermissionConfig(val *CharacterPermissionConfig) *NullableCharacterPermissionConfig {
	return &NullableCharacterPermissionConfig{value: val, isSet: true}
}

func (v NullableCharacterPermissionConfig) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCharacterPermissionConfig) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}