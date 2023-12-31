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

// checks if the ChatSampleItem type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ChatSampleItem{}

// ChatSampleItem struct for ChatSampleItem
type ChatSampleItem struct {
	Role                 *string `json:"role,omitempty"`
	Name                 *string `json:"name,omitempty"`
	Content              *string `json:"content,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _ChatSampleItem ChatSampleItem

// NewChatSampleItem instantiates a new ChatSampleItem object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewChatSampleItem() *ChatSampleItem {
	this := ChatSampleItem{}
	return &this
}

// NewChatSampleItemWithDefaults instantiates a new ChatSampleItem object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewChatSampleItemWithDefaults() *ChatSampleItem {
	this := ChatSampleItem{}
	return &this
}

// GetRole returns the Role field value if set, zero value otherwise.
func (o *ChatSampleItem) GetRole() string {
	if o == nil || IsNil(o.Role) {
		var ret string
		return ret
	}
	return *o.Role
}

// GetRoleOk returns a tuple with the Role field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ChatSampleItem) GetRoleOk() (*string, bool) {
	if o == nil || IsNil(o.Role) {
		return nil, false
	}
	return o.Role, true
}

// HasRole returns a boolean if a field has been set.
func (o *ChatSampleItem) HasRole() bool {
	if o != nil && !IsNil(o.Role) {
		return true
	}

	return false
}

// SetRole gets a reference to the given string and assigns it to the Role field.
func (o *ChatSampleItem) SetRole(v string) {
	o.Role = &v
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *ChatSampleItem) GetName() string {
	if o == nil || IsNil(o.Name) {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ChatSampleItem) GetNameOk() (*string, bool) {
	if o == nil || IsNil(o.Name) {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *ChatSampleItem) HasName() bool {
	if o != nil && !IsNil(o.Name) {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *ChatSampleItem) SetName(v string) {
	o.Name = &v
}

// GetContent returns the Content field value if set, zero value otherwise.
func (o *ChatSampleItem) GetContent() string {
	if o == nil || IsNil(o.Content) {
		var ret string
		return ret
	}
	return *o.Content
}

// GetContentOk returns a tuple with the Content field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ChatSampleItem) GetContentOk() (*string, bool) {
	if o == nil || IsNil(o.Content) {
		return nil, false
	}
	return o.Content, true
}

// HasContent returns a boolean if a field has been set.
func (o *ChatSampleItem) HasContent() bool {
	if o != nil && !IsNil(o.Content) {
		return true
	}

	return false
}

// SetContent gets a reference to the given string and assigns it to the Content field.
func (o *ChatSampleItem) SetContent(v string) {
	o.Content = &v
}

func (o ChatSampleItem) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ChatSampleItem) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Role) {
		toSerialize["role"] = o.Role
	}
	if !IsNil(o.Name) {
		toSerialize["name"] = o.Name
	}
	if !IsNil(o.Content) {
		toSerialize["content"] = o.Content
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *ChatSampleItem) UnmarshalJSON(bytes []byte) (err error) {
	varChatSampleItem := _ChatSampleItem{}

	err = json.Unmarshal(bytes, &varChatSampleItem)

	if err != nil {
		return err
	}

	*o = ChatSampleItem(varChatSampleItem)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "role")
		delete(additionalProperties, "name")
		delete(additionalProperties, "content")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableChatSampleItem struct {
	value *ChatSampleItem
	isSet bool
}

func (v NullableChatSampleItem) Get() *ChatSampleItem {
	return v.value
}

func (v *NullableChatSampleItem) Set(val *ChatSampleItem) {
	v.value = val
	v.isSet = true
}

func (v NullableChatSampleItem) IsSet() bool {
	return v.isSet
}

func (v *NullableChatSampleItem) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableChatSampleItem(val *ChatSampleItem) *NullableChatSampleItem {
	return &NullableChatSampleItem{value: val, isSet: true}
}

func (v NullableChatSampleItem) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableChatSampleItem) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
