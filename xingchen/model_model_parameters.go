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

// checks if the ModelParameters type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ModelParameters{}

// ModelParameters struct for ModelParameters
type ModelParameters struct {
	ModelName            *string  `json:"modelName,omitempty"`
	TopP                 *float64 `json:"topP,omitempty"`
	TopK                 *int32   `json:"topK,omitempty"`
	Seed                 *int64   `json:"seed,omitempty"`
	MinLength            *int32   `json:"minLength,omitempty"`
	MaxLength            *int32   `json:"maxLength,omitempty"`
	Temperature          *float64 `json:"temperature,omitempty"`
	IncrementalOutput    *bool    `json:"incremental_output,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _ModelParameters ModelParameters

// NewModelParameters instantiates a new ModelParameters object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewModelParameters() *ModelParameters {
	this := ModelParameters{}
	return &this
}

// NewModelParametersWithDefaults instantiates a new ModelParameters object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewModelParametersWithDefaults() *ModelParameters {
	this := ModelParameters{}
	return &this
}

// GetModelName returns the ModelName field value if set, zero value otherwise.
func (o *ModelParameters) GetModelName() string {
	if o == nil || IsNil(o.ModelName) {
		var ret string
		return ret
	}
	return *o.ModelName
}

// GetModelNameOk returns a tuple with the ModelName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ModelParameters) GetModelNameOk() (*string, bool) {
	if o == nil || IsNil(o.ModelName) {
		return nil, false
	}
	return o.ModelName, true
}

// HasModelName returns a boolean if a field has been set.
func (o *ModelParameters) HasModelName() bool {
	if o != nil && !IsNil(o.ModelName) {
		return true
	}

	return false
}

// SetModelName gets a reference to the given string and assigns it to the ModelName field.
func (o *ModelParameters) SetModelName(v string) {
	o.ModelName = &v
}

// GetTopP returns the TopP field value if set, zero value otherwise.
func (o *ModelParameters) GetTopP() float64 {
	if o == nil || IsNil(o.TopP) {
		var ret float64
		return ret
	}
	return *o.TopP
}

// GetTopPOk returns a tuple with the TopP field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ModelParameters) GetTopPOk() (*float64, bool) {
	if o == nil || IsNil(o.TopP) {
		return nil, false
	}
	return o.TopP, true
}

// HasTopP returns a boolean if a field has been set.
func (o *ModelParameters) HasTopP() bool {
	if o != nil && !IsNil(o.TopP) {
		return true
	}

	return false
}

// SetTopP gets a reference to the given float64 and assigns it to the TopP field.
func (o *ModelParameters) SetTopP(v float64) {
	o.TopP = &v
}

// GetTopK returns the TopK field value if set, zero value otherwise.
func (o *ModelParameters) GetTopK() int32 {
	if o == nil || IsNil(o.TopK) {
		var ret int32
		return ret
	}
	return *o.TopK
}

// GetTopKOk returns a tuple with the TopK field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ModelParameters) GetTopKOk() (*int32, bool) {
	if o == nil || IsNil(o.TopK) {
		return nil, false
	}
	return o.TopK, true
}

// HasTopK returns a boolean if a field has been set.
func (o *ModelParameters) HasTopK() bool {
	if o != nil && !IsNil(o.TopK) {
		return true
	}

	return false
}

// SetTopK gets a reference to the given int32 and assigns it to the TopK field.
func (o *ModelParameters) SetTopK(v int32) {
	o.TopK = &v
}

// GetSeed returns the Seed field value if set, zero value otherwise.
func (o *ModelParameters) GetSeed() int64 {
	if o == nil || IsNil(o.Seed) {
		var ret int64
		return ret
	}
	return *o.Seed
}

// GetSeedOk returns a tuple with the Seed field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ModelParameters) GetSeedOk() (*int64, bool) {
	if o == nil || IsNil(o.Seed) {
		return nil, false
	}
	return o.Seed, true
}

// HasSeed returns a boolean if a field has been set.
func (o *ModelParameters) HasSeed() bool {
	if o != nil && !IsNil(o.Seed) {
		return true
	}

	return false
}

// SetSeed gets a reference to the given int64 and assigns it to the Seed field.
func (o *ModelParameters) SetSeed(v int64) {
	o.Seed = &v
}

// GetMinLength returns the MinLength field value if set, zero value otherwise.
func (o *ModelParameters) GetMinLength() int32 {
	if o == nil || IsNil(o.MinLength) {
		var ret int32
		return ret
	}
	return *o.MinLength
}

// GetMinLengthOk returns a tuple with the MinLength field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ModelParameters) GetMinLengthOk() (*int32, bool) {
	if o == nil || IsNil(o.MinLength) {
		return nil, false
	}
	return o.MinLength, true
}

// HasMinLength returns a boolean if a field has been set.
func (o *ModelParameters) HasMinLength() bool {
	if o != nil && !IsNil(o.MinLength) {
		return true
	}

	return false
}

// SetMinLength gets a reference to the given int32 and assigns it to the MinLength field.
func (o *ModelParameters) SetMinLength(v int32) {
	o.MinLength = &v
}

// GetMaxLength returns the MaxLength field value if set, zero value otherwise.
func (o *ModelParameters) GetMaxLength() int32 {
	if o == nil || IsNil(o.MaxLength) {
		var ret int32
		return ret
	}
	return *o.MaxLength
}

// GetMaxLengthOk returns a tuple with the MaxLength field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ModelParameters) GetMaxLengthOk() (*int32, bool) {
	if o == nil || IsNil(o.MaxLength) {
		return nil, false
	}
	return o.MaxLength, true
}

// HasMaxLength returns a boolean if a field has been set.
func (o *ModelParameters) HasMaxLength() bool {
	if o != nil && !IsNil(o.MaxLength) {
		return true
	}

	return false
}

// SetMaxLength gets a reference to the given int32 and assigns it to the MaxLength field.
func (o *ModelParameters) SetMaxLength(v int32) {
	o.MaxLength = &v
}

// GetTemperature returns the Temperature field value if set, zero value otherwise.
func (o *ModelParameters) GetTemperature() float64 {
	if o == nil || IsNil(o.Temperature) {
		var ret float64
		return ret
	}
	return *o.Temperature
}

// GetTemperatureOk returns a tuple with the Temperature field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ModelParameters) GetTemperatureOk() (*float64, bool) {
	if o == nil || IsNil(o.Temperature) {
		return nil, false
	}
	return o.Temperature, true
}

// HasTemperature returns a boolean if a field has been set.
func (o *ModelParameters) HasTemperature() bool {
	if o != nil && !IsNil(o.Temperature) {
		return true
	}

	return false
}

// SetTemperature gets a reference to the given float64 and assigns it to the Temperature field.
func (o *ModelParameters) SetTemperature(v float64) {
	o.Temperature = &v
}

func (o *ModelParameters) GetIncrementalOutput() bool {
	if o == nil || IsNil(o.IncrementalOutput) {
		var ret bool
		return ret
	}
	return *o.IncrementalOutput
}

func (o *ModelParameters) GetIncrementalOutputOk() (*bool, bool) {
	if o == nil || IsNil(o.IncrementalOutput) {
		return nil, false
	}
	return o.IncrementalOutput, true
}

// HasModelName returns a boolean if a field has been set.
func (o *ModelParameters) HasIncrementalOutput() bool {
	if o != nil && !IsNil(o.IncrementalOutput) {
		return true
	}

	return false
}

// SetModelName gets a reference to the given string and assigns it to the ModelName field.
func (o *ModelParameters) SetIncrementalOutput(v bool) {
	o.IncrementalOutput = &v
}

func (o ModelParameters) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ModelParameters) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.ModelName) {
		toSerialize["modelName"] = o.ModelName
	}
	if !IsNil(o.TopP) {
		toSerialize["topP"] = o.TopP
	}
	if !IsNil(o.TopK) {
		toSerialize["topK"] = o.TopK
	}
	if !IsNil(o.Seed) {
		toSerialize["seed"] = o.Seed
	}
	if !IsNil(o.MinLength) {
		toSerialize["minLength"] = o.MinLength
	}
	if !IsNil(o.MaxLength) {
		toSerialize["maxLength"] = o.MaxLength
	}
	if !IsNil(o.Temperature) {
		toSerialize["temperature"] = o.Temperature
	}
	if !IsNil(o.IncrementalOutput) {
		toSerialize["incrementalOutput"] = o.IncrementalOutput
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *ModelParameters) UnmarshalJSON(bytes []byte) (err error) {
	varModelParameters := _ModelParameters{}

	err = json.Unmarshal(bytes, &varModelParameters)

	if err != nil {
		return err
	}

	*o = ModelParameters(varModelParameters)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "modelName")
		delete(additionalProperties, "topP")
		delete(additionalProperties, "topK")
		delete(additionalProperties, "seed")
		delete(additionalProperties, "minLength")
		delete(additionalProperties, "maxLength")
		delete(additionalProperties, "temperature")
		delete(additionalProperties, "incrementalOutput")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableModelParameters struct {
	value *ModelParameters
	isSet bool
}

func (v NullableModelParameters) Get() *ModelParameters {
	return v.value
}

func (v *NullableModelParameters) Set(val *ModelParameters) {
	v.value = val
	v.isSet = true
}

func (v NullableModelParameters) IsSet() bool {
	return v.isSet
}

func (v *NullableModelParameters) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableModelParameters(val *ModelParameters) *NullableModelParameters {
	return &NullableModelParameters{value: val, isSet: true}
}

func (v NullableModelParameters) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableModelParameters) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
