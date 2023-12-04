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

// checks if the BotProfile type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &BotProfile{}

// BotProfile struct for BotProfile
type BotProfile struct {
	CharacterId          *string `json:"characterId,omitempty"`
	Version              *int32  `json:"version,omitempty"`
	Name                 *string `json:"name,omitempty"`
	Content              *string `json:"content,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _BotProfile BotProfile

// NewBotProfile instantiates a new BotProfile object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewBotProfile() *BotProfile {
	this := BotProfile{}
	return &this
}

// NewBotProfileWithDefaults instantiates a new BotProfile object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewBotProfileWithDefaults() *BotProfile {
	this := BotProfile{}
	return &this
}

// GetCharacterId returns the CharacterId field value if set, zero value otherwise.
func (o *BotProfile) GetCharacterId() string {
	if o == nil || IsNil(o.CharacterId) {
		var ret string
		return ret
	}
	return *o.CharacterId
}

// GetCharacterIdOk returns a tuple with the CharacterId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BotProfile) GetCharacterIdOk() (*string, bool) {
	if o == nil || IsNil(o.CharacterId) {
		return nil, false
	}
	return o.CharacterId, true
}

// HasCharacterId returns a boolean if a field has been set.
func (o *BotProfile) HasCharacterId() bool {
	if o != nil && !IsNil(o.CharacterId) {
		return true
	}

	return false
}

// SetCharacterId gets a reference to the given string and assigns it to the CharacterId field.
func (o *BotProfile) SetCharacterId(v string) {
	o.CharacterId = &v
}

// GetVersion returns the Version field value if set, zero value otherwise.
func (o *BotProfile) GetVersion() int32 {
	if o == nil || IsNil(o.Version) {
		var ret int32
		return ret
	}
	return *o.Version
}

// GetVersionOk returns a tuple with the Version field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BotProfile) GetVersionOk() (*int32, bool) {
	if o == nil || IsNil(o.Version) {
		return nil, false
	}
	return o.Version, true
}

// HasVersion returns a boolean if a field has been set.
func (o *BotProfile) HasVersion() bool {
	if o != nil && !IsNil(o.Version) {
		return true
	}

	return false
}

// SetVersion gets a reference to the given int32 and assigns it to the Version field.
func (o *BotProfile) SetVersion(v int32) {
	o.Version = &v
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *BotProfile) GetName() string {
	if o == nil || IsNil(o.Name) {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BotProfile) GetNameOk() (*string, bool) {
	if o == nil || IsNil(o.Name) {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *BotProfile) HasName() bool {
	if o != nil && !IsNil(o.Name) {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *BotProfile) SetName(v string) {
	o.Name = &v
}

// GetContent returns the Content field value if set, zero value otherwise.
func (o *BotProfile) GetContent() string {
	if o == nil || IsNil(o.Content) {
		var ret string
		return ret
	}
	return *o.Content
}

// GetContentOk returns a tuple with the Content field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BotProfile) GetContentOk() (*string, bool) {
	if o == nil || IsNil(o.Content) {
		return nil, false
	}
	return o.Content, true
}

// HasContent returns a boolean if a field has been set.
func (o *BotProfile) HasContent() bool {
	if o != nil && !IsNil(o.Content) {
		return true
	}

	return false
}

// SetContent gets a reference to the given string and assigns it to the Content field.
func (o *BotProfile) SetContent(v string) {
	o.Content = &v
}

func (o BotProfile) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o BotProfile) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.CharacterId) {
		toSerialize["characterId"] = o.CharacterId
	}
	if !IsNil(o.Version) {
		toSerialize["version"] = o.Version
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

func (o *BotProfile) UnmarshalJSON(bytes []byte) (err error) {
	varBotProfile := _BotProfile{}

	err = json.Unmarshal(bytes, &varBotProfile)

	if err != nil {
		return err
	}

	*o = BotProfile(varBotProfile)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "characterId")
		delete(additionalProperties, "version")
		delete(additionalProperties, "name")
		delete(additionalProperties, "content")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableBotProfile struct {
	value *BotProfile
	isSet bool
}

func (v NullableBotProfile) Get() *BotProfile {
	return v.value
}

func (v *NullableBotProfile) Set(val *BotProfile) {
	v.value = val
	v.isSet = true
}

func (v NullableBotProfile) IsSet() bool {
	return v.isSet
}

func (v *NullableBotProfile) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableBotProfile(val *BotProfile) *NullableBotProfile {
	return &NullableBotProfile{value: val, isSet: true}
}

func (v NullableBotProfile) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableBotProfile) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}