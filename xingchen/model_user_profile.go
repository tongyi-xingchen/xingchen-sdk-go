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

// checks if the UserProfile type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &UserProfile{}

// UserProfile struct for UserProfile
type UserProfile struct {
	UserId               string  `json:"userId"`
	UserName             *string `json:"userName,omitempty"`
	BasicInfo            *string `json:"basicInfo,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _UserProfile UserProfile

// NewUserProfile instantiates a new UserProfile object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewUserProfile(userId string) *UserProfile {
	this := UserProfile{}
	this.UserId = userId
	return &this
}

// NewUserProfileWithDefaults instantiates a new UserProfile object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewUserProfileWithDefaults() *UserProfile {
	this := UserProfile{}
	return &this
}

// GetUserId returns the UserId field value
func (o *UserProfile) GetUserId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.UserId
}

// GetUserIdOk returns a tuple with the UserId field value
// and a boolean to check if the value has been set.
func (o *UserProfile) GetUserIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.UserId, true
}

// SetUserId sets field value
func (o *UserProfile) SetUserId(v string) {
	o.UserId = v
}

// GetUserName returns the UserName field value if set, zero value otherwise.
func (o *UserProfile) GetUserName() string {
	if o == nil || IsNil(o.UserName) {
		var ret string
		return ret
	}
	return *o.UserName
}

// GetUserNameOk returns a tuple with the UserName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UserProfile) GetUserNameOk() (*string, bool) {
	if o == nil || IsNil(o.UserName) {
		return nil, false
	}
	return o.UserName, true
}

// HasUserName returns a boolean if a field has been set.
func (o *UserProfile) HasUserName() bool {
	if o != nil && !IsNil(o.UserName) {
		return true
	}

	return false
}

// SetUserName gets a reference to the given string and assigns it to the UserName field.
func (o *UserProfile) SetUserName(v string) {
	o.UserName = &v
}

// GetBasicInfo returns the BasicInfo field value if set, zero value otherwise.
func (o *UserProfile) GetBasicInfo() string {
	if o == nil || IsNil(o.BasicInfo) {
		var ret string
		return ret
	}
	return *o.BasicInfo
}

// GetBasicInfoOk returns a tuple with the BasicInfo field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UserProfile) GetBasicInfoOk() (*string, bool) {
	if o == nil || IsNil(o.BasicInfo) {
		return nil, false
	}
	return o.BasicInfo, true
}

// HasBasicInfo returns a boolean if a field has been set.
func (o *UserProfile) HasBasicInfo() bool {
	if o != nil && !IsNil(o.BasicInfo) {
		return true
	}

	return false
}

// SetBasicInfo gets a reference to the given string and assigns it to the BasicInfo field.
func (o *UserProfile) SetBasicInfo(v string) {
	o.BasicInfo = &v
}

func (o UserProfile) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o UserProfile) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["userId"] = o.UserId
	if !IsNil(o.UserName) {
		toSerialize["userName"] = o.UserName
	}
	if !IsNil(o.BasicInfo) {
		toSerialize["basicInfo"] = o.BasicInfo
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *UserProfile) UnmarshalJSON(bytes []byte) (err error) {
	varUserProfile := _UserProfile{}

	err = json.Unmarshal(bytes, &varUserProfile)

	if err != nil {
		return err
	}

	*o = UserProfile(varUserProfile)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "userId")
		delete(additionalProperties, "userName")
		delete(additionalProperties, "basicInfo")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableUserProfile struct {
	value *UserProfile
	isSet bool
}

func (v NullableUserProfile) Get() *UserProfile {
	return v.value
}

func (v *NullableUserProfile) Set(val *UserProfile) {
	v.value = val
	v.isSet = true
}

func (v NullableUserProfile) IsSet() bool {
	return v.isSet
}

func (v *NullableUserProfile) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableUserProfile(val *UserProfile) *NullableUserProfile {
	return &NullableUserProfile{value: val, isSet: true}
}

func (v NullableUserProfile) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableUserProfile) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
