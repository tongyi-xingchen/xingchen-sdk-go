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

// checks if the AdvancedSettings type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &AdvancedSettings{}

// AdvancedSettings struct for AdvancedSettings
type AdvancedSettings struct {
	EnableWebSearch         *bool    `json:"enableWebSearch,omitempty"`
	SearchEnhancedKeyword   *string  `json:"searchEnhancedKeyword,omitempty"`
	EnableCharacterKbSearch *bool    `json:"enableCharacterKbSearch,omitempty"`
	KnowledgeBases          []string `json:"KnowledgeBases,omitempty"`
	EnableLongTermMemory    *bool    `json:"enableLongTermMemory,omitempty"`
	EnableDebug             *bool    `json:"enableDebug,omitempty"`
	AdditionalProperties    map[string]interface{}
}

type _AdvancedSettings AdvancedSettings

// NewAdvancedSettings instantiates a new AdvancedSettings object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAdvancedSettings() *AdvancedSettings {
	this := AdvancedSettings{}
	return &this
}

// NewAdvancedSettingsWithDefaults instantiates a new AdvancedSettings object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAdvancedSettingsWithDefaults() *AdvancedSettings {
	this := AdvancedSettings{}
	return &this
}

// GetEnableWebSearch returns the EnableWebSearch field value if set, zero value otherwise.
func (o *AdvancedSettings) GetEnableWebSearch() bool {
	if o == nil || IsNil(o.EnableWebSearch) {
		var ret bool
		return ret
	}
	return *o.EnableWebSearch
}

// GetEnableWebSearchOk returns a tuple with the EnableWebSearch field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AdvancedSettings) GetEnableWebSearchOk() (*bool, bool) {
	if o == nil || IsNil(o.EnableWebSearch) {
		return nil, false
	}
	return o.EnableWebSearch, true
}

// HasEnableWebSearch returns a boolean if a field has been set.
func (o *AdvancedSettings) HasEnableWebSearch() bool {
	if o != nil && !IsNil(o.EnableWebSearch) {
		return true
	}

	return false
}

// SetEnableWebSearch gets a reference to the given bool and assigns it to the EnableWebSearch field.
func (o *AdvancedSettings) SetEnableWebSearch(v bool) {
	o.EnableWebSearch = &v
}

// GetSearchEnhancedKeyword returns the SearchEnhancedKeyword field value if set, zero value otherwise.
func (o *AdvancedSettings) GetSearchEnhancedKeyword() string {
	if o == nil || IsNil(o.SearchEnhancedKeyword) {
		var ret string
		return ret
	}
	return *o.SearchEnhancedKeyword
}

// GetSearchEnhancedKeywordOk returns a tuple with the SearchEnhancedKeyword field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AdvancedSettings) GetSearchEnhancedKeywordOk() (*string, bool) {
	if o == nil || IsNil(o.SearchEnhancedKeyword) {
		return nil, false
	}
	return o.SearchEnhancedKeyword, true
}

// HasSearchEnhancedKeyword returns a boolean if a field has been set.
func (o *AdvancedSettings) HasSearchEnhancedKeyword() bool {
	if o != nil && !IsNil(o.SearchEnhancedKeyword) {
		return true
	}

	return false
}

// SetSearchEnhancedKeyword gets a reference to the given string and assigns it to the SearchEnhancedKeyword field.
func (o *AdvancedSettings) SetSearchEnhancedKeyword(v string) {
	o.SearchEnhancedKeyword = &v
}

// GetEnableCharacterKbSearch returns the EnableCharacterKbSearch field value if set, zero value otherwise.
func (o *AdvancedSettings) GetEnableCharacterKbSearch() bool {
	if o == nil || IsNil(o.EnableCharacterKbSearch) {
		var ret bool
		return ret
	}
	return *o.EnableCharacterKbSearch
}

// GetEnableCharacterKbSearchOk returns a tuple with the EnableCharacterKbSearch field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AdvancedSettings) GetEnableCharacterKbSearchOk() (*bool, bool) {
	if o == nil || IsNil(o.EnableCharacterKbSearch) {
		return nil, false
	}
	return o.EnableCharacterKbSearch, true
}

// HasEnableCharacterKbSearch returns a boolean if a field has been set.
func (o *AdvancedSettings) HasEnableCharacterKbSearch() bool {
	if o != nil && !IsNil(o.EnableCharacterKbSearch) {
		return true
	}

	return false
}

// SetEnableCharacterKbSearch gets a reference to the given bool and assigns it to the EnableCharacterKbSearch field.
func (o *AdvancedSettings) SetEnableCharacterKbSearch(v bool) {
	o.EnableCharacterKbSearch = &v
}

func (o *AdvancedSettings) GetKnowledgeBases() []string {
	if o == nil || IsNil(o.KnowledgeBases) {
		var ret []string
		return ret
	}
	return o.KnowledgeBases
}

func (o *AdvancedSettings) GetKnowledgeBasesOk() ([]string, bool) {
	if o == nil || IsNil(o.KnowledgeBases) {
		return nil, false
	}
	return o.KnowledgeBases, true
}

func (o *AdvancedSettings) HasKnowledgeBases() bool {
	if o != nil && !IsNil(o.KnowledgeBases) {
		return true
	}

	return false
}

func (o *AdvancedSettings) SetKnowledgeBases(v []string) {
	o.KnowledgeBases = v
}

// GetEnableLongTermMemory returns the EnableLongTermMemory field value if set, zero value otherwise.
func (o *AdvancedSettings) GetEnableLongTermMemory() bool {
	if o == nil || IsNil(o.EnableLongTermMemory) {
		var ret bool
		return ret
	}
	return *o.EnableLongTermMemory
}

// GetEnableLongTermMemoryOk returns a tuple with the EnableLongTermMemory field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AdvancedSettings) GetEnableLongTermMemoryOk() (*bool, bool) {
	if o == nil || IsNil(o.EnableLongTermMemory) {
		return nil, false
	}
	return o.EnableLongTermMemory, true
}

// HasEnableLongTermMemory returns a boolean if a field has been set.
func (o *AdvancedSettings) HasEnableLongTermMemory() bool {
	if o != nil && !IsNil(o.EnableLongTermMemory) {
		return true
	}

	return false
}

// SetEnableLongTermMemory gets a reference to the given bool and assigns it to the EnableLongTermMemory field.
func (o *AdvancedSettings) SetEnableLongTermMemory(v bool) {
	o.EnableLongTermMemory = &v
}

// GetEnableDebug returns the EnableDebug field value if set, zero value otherwise.
func (o *AdvancedSettings) GetEnableDebug() bool {
	if o == nil || IsNil(o.EnableDebug) {
		var ret bool
		return ret
	}
	return *o.EnableDebug
}

// GetEnableDebugOk returns a tuple with the EnableDebug field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AdvancedSettings) GetEnableDebugOk() (*bool, bool) {
	if o == nil || IsNil(o.EnableDebug) {
		return nil, false
	}
	return o.EnableDebug, true
}

// HasEnableDebug returns a boolean if a field has been set.
func (o *AdvancedSettings) HasEnableDebug() bool {
	if o != nil && !IsNil(o.EnableDebug) {
		return true
	}

	return false
}

// SetEnableDebug gets a reference to the given bool and assigns it to the EnableDebug field.
func (o *AdvancedSettings) SetEnableDebug(v bool) {
	o.EnableDebug = &v
}

func (o AdvancedSettings) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o AdvancedSettings) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.EnableWebSearch) {
		toSerialize["enableWebSearch"] = o.EnableWebSearch
	}
	if !IsNil(o.SearchEnhancedKeyword) {
		toSerialize["searchEnhancedKeyword"] = o.SearchEnhancedKeyword
	}
	if !IsNil(o.EnableCharacterKbSearch) {
		toSerialize["enableCharacterKbSearch"] = o.EnableCharacterKbSearch
	}
	if !IsNil(o.EnableLongTermMemory) {
		toSerialize["enableLongTermMemory"] = o.EnableLongTermMemory
	}
	if !IsNil(o.EnableDebug) {
		toSerialize["enableDebug"] = o.EnableDebug
	}
	if !IsNil(o.KnowledgeBases) {
		toSerialize["knowledgeBases"] = o.KnowledgeBases
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *AdvancedSettings) UnmarshalJSON(bytes []byte) (err error) {
	varAdvancedSettings := _AdvancedSettings{}

	err = json.Unmarshal(bytes, &varAdvancedSettings)

	if err != nil {
		return err
	}

	*o = AdvancedSettings(varAdvancedSettings)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "enableWebSearch")
		delete(additionalProperties, "searchEnhancedKeyword")
		delete(additionalProperties, "enableCharacterKbSearch")
		delete(additionalProperties, "enableLongTermMemory")
		delete(additionalProperties, "enableDebug")
		delete(additionalProperties, "knowledgeBases")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableAdvancedSettings struct {
	value *AdvancedSettings
	isSet bool
}

func (v NullableAdvancedSettings) Get() *AdvancedSettings {
	return v.value
}

func (v *NullableAdvancedSettings) Set(val *AdvancedSettings) {
	v.value = val
	v.isSet = true
}

func (v NullableAdvancedSettings) IsSet() bool {
	return v.isSet
}

func (v *NullableAdvancedSettings) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAdvancedSettings(val *AdvancedSettings) *NullableAdvancedSettings {
	return &NullableAdvancedSettings{value: val, isSet: true}
}

func (v NullableAdvancedSettings) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAdvancedSettings) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
