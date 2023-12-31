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

// checks if the ChatRoomUserDTO type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ChatRoomUserDTO{}

// ChatRoomUserDTO 聊天室用户
type ChatRoomUserDTO struct {
	// 用户ID，若为普通用户，则为userId,;若为角色，则为characterId
	UserId *string `json:"userId,omitempty"`
	// 若为ISV用户，则为ISV用户ID；若为角色，则为空
	BizUserId *string `json:"bizUserId,omitempty"`
	// 用户名称，若为普通用户，则为用户名称；若为角色，则为 characterName
	UserName *string `json:"userName,omitempty"`
	// 用户类型，可以是 user（用户）或character（角色）
	UserType             *string `json:"userType,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _ChatRoomUserDTO ChatRoomUserDTO

// NewChatRoomUserDTO instantiates a new ChatRoomUserDTO object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewChatRoomUserDTO() *ChatRoomUserDTO {
	this := ChatRoomUserDTO{}
	return &this
}

// NewChatRoomUserDTOWithDefaults instantiates a new ChatRoomUserDTO object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewChatRoomUserDTOWithDefaults() *ChatRoomUserDTO {
	this := ChatRoomUserDTO{}
	return &this
}

// GetUserId returns the UserId field value if set, zero value otherwise.
func (o *ChatRoomUserDTO) GetUserId() string {
	if o == nil || IsNil(o.UserId) {
		var ret string
		return ret
	}
	return *o.UserId
}

// GetUserIdOk returns a tuple with the UserId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ChatRoomUserDTO) GetUserIdOk() (*string, bool) {
	if o == nil || IsNil(o.UserId) {
		return nil, false
	}
	return o.UserId, true
}

// HasUserId returns a boolean if a field has been set.
func (o *ChatRoomUserDTO) HasUserId() bool {
	if o != nil && !IsNil(o.UserId) {
		return true
	}

	return false
}

// SetUserId gets a reference to the given string and assigns it to the UserId field.
func (o *ChatRoomUserDTO) SetUserId(v string) {
	o.UserId = &v
}

// GetBizUserId returns the BizUserId field value if set, zero value otherwise.
func (o *ChatRoomUserDTO) GetBizUserId() string {
	if o == nil || IsNil(o.BizUserId) {
		var ret string
		return ret
	}
	return *o.BizUserId
}

// GetBizUserIdOk returns a tuple with the BizUserId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ChatRoomUserDTO) GetBizUserIdOk() (*string, bool) {
	if o == nil || IsNil(o.BizUserId) {
		return nil, false
	}
	return o.BizUserId, true
}

// HasBizUserId returns a boolean if a field has been set.
func (o *ChatRoomUserDTO) HasBizUserId() bool {
	if o != nil && !IsNil(o.BizUserId) {
		return true
	}

	return false
}

// SetBizUserId gets a reference to the given string and assigns it to the BizUserId field.
func (o *ChatRoomUserDTO) SetBizUserId(v string) {
	o.BizUserId = &v
}

// GetUserName returns the UserName field value if set, zero value otherwise.
func (o *ChatRoomUserDTO) GetUserName() string {
	if o == nil || IsNil(o.UserName) {
		var ret string
		return ret
	}
	return *o.UserName
}

// GetUserNameOk returns a tuple with the UserName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ChatRoomUserDTO) GetUserNameOk() (*string, bool) {
	if o == nil || IsNil(o.UserName) {
		return nil, false
	}
	return o.UserName, true
}

// HasUserName returns a boolean if a field has been set.
func (o *ChatRoomUserDTO) HasUserName() bool {
	if o != nil && !IsNil(o.UserName) {
		return true
	}

	return false
}

// SetUserName gets a reference to the given string and assigns it to the UserName field.
func (o *ChatRoomUserDTO) SetUserName(v string) {
	o.UserName = &v
}

// GetUserType returns the UserType field value if set, zero value otherwise.
func (o *ChatRoomUserDTO) GetUserType() string {
	if o == nil || IsNil(o.UserType) {
		var ret string
		return ret
	}
	return *o.UserType
}

// GetUserTypeOk returns a tuple with the UserType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ChatRoomUserDTO) GetUserTypeOk() (*string, bool) {
	if o == nil || IsNil(o.UserType) {
		return nil, false
	}
	return o.UserType, true
}

// HasUserType returns a boolean if a field has been set.
func (o *ChatRoomUserDTO) HasUserType() bool {
	if o != nil && !IsNil(o.UserType) {
		return true
	}

	return false
}

// SetUserType gets a reference to the given string and assigns it to the UserType field.
func (o *ChatRoomUserDTO) SetUserType(v string) {
	o.UserType = &v
}

func (o ChatRoomUserDTO) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ChatRoomUserDTO) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.UserId) {
		toSerialize["userId"] = o.UserId
	}
	if !IsNil(o.BizUserId) {
		toSerialize["bizUserId"] = o.BizUserId
	}
	if !IsNil(o.UserName) {
		toSerialize["userName"] = o.UserName
	}
	if !IsNil(o.UserType) {
		toSerialize["userType"] = o.UserType
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *ChatRoomUserDTO) UnmarshalJSON(bytes []byte) (err error) {
	varChatRoomUserDTO := _ChatRoomUserDTO{}

	err = json.Unmarshal(bytes, &varChatRoomUserDTO)

	if err != nil {
		return err
	}

	*o = ChatRoomUserDTO(varChatRoomUserDTO)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "userId")
		delete(additionalProperties, "bizUserId")
		delete(additionalProperties, "userName")
		delete(additionalProperties, "userType")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableChatRoomUserDTO struct {
	value *ChatRoomUserDTO
	isSet bool
}

func (v NullableChatRoomUserDTO) Get() *ChatRoomUserDTO {
	return v.value
}

func (v *NullableChatRoomUserDTO) Set(val *ChatRoomUserDTO) {
	v.value = val
	v.isSet = true
}

func (v NullableChatRoomUserDTO) IsSet() bool {
	return v.isSet
}

func (v *NullableChatRoomUserDTO) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableChatRoomUserDTO(val *ChatRoomUserDTO) *NullableChatRoomUserDTO {
	return &NullableChatRoomUserDTO{value: val, isSet: true}
}

func (v NullableChatRoomUserDTO) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableChatRoomUserDTO) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
