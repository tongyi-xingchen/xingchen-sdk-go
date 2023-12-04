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

// checks if the ResultDTOPageResultChatMessageDTO type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ResultDTOPageResultChatMessageDTO{}

// ResultDTOPageResultChatMessageDTO struct for ResultDTOPageResultChatMessageDTO
type ResultDTOPageResultChatMessageDTO struct {
	RequestId            *string                   `json:"requestId,omitempty"`
	Code                 *int32                    `json:"code,omitempty"`
	ErrorMessage         *string                   `json:"errorMessage,omitempty"`
	ErrorMessageKey      *string                   `json:"errorMessageKey,omitempty"`
	Data                 *PageResultChatMessageDTO `json:"data,omitempty"`
	Success              *bool                     `json:"success,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _ResultDTOPageResultChatMessageDTO ResultDTOPageResultChatMessageDTO

// NewResultDTOPageResultChatMessageDTO instantiates a new ResultDTOPageResultChatMessageDTO object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewResultDTOPageResultChatMessageDTO() *ResultDTOPageResultChatMessageDTO {
	this := ResultDTOPageResultChatMessageDTO{}
	return &this
}

// NewResultDTOPageResultChatMessageDTOWithDefaults instantiates a new ResultDTOPageResultChatMessageDTO object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewResultDTOPageResultChatMessageDTOWithDefaults() *ResultDTOPageResultChatMessageDTO {
	this := ResultDTOPageResultChatMessageDTO{}
	return &this
}

// GetRequestId returns the RequestId field value if set, zero value otherwise.
func (o *ResultDTOPageResultChatMessageDTO) GetRequestId() string {
	if o == nil || IsNil(o.RequestId) {
		var ret string
		return ret
	}
	return *o.RequestId
}

// GetRequestIdOk returns a tuple with the RequestId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ResultDTOPageResultChatMessageDTO) GetRequestIdOk() (*string, bool) {
	if o == nil || IsNil(o.RequestId) {
		return nil, false
	}
	return o.RequestId, true
}

// HasRequestId returns a boolean if a field has been set.
func (o *ResultDTOPageResultChatMessageDTO) HasRequestId() bool {
	if o != nil && !IsNil(o.RequestId) {
		return true
	}

	return false
}

// SetRequestId gets a reference to the given string and assigns it to the RequestId field.
func (o *ResultDTOPageResultChatMessageDTO) SetRequestId(v string) {
	o.RequestId = &v
}

// GetCode returns the Code field value if set, zero value otherwise.
func (o *ResultDTOPageResultChatMessageDTO) GetCode() int32 {
	if o == nil || IsNil(o.Code) {
		var ret int32
		return ret
	}
	return *o.Code
}

// GetCodeOk returns a tuple with the Code field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ResultDTOPageResultChatMessageDTO) GetCodeOk() (*int32, bool) {
	if o == nil || IsNil(o.Code) {
		return nil, false
	}
	return o.Code, true
}

// HasCode returns a boolean if a field has been set.
func (o *ResultDTOPageResultChatMessageDTO) HasCode() bool {
	if o != nil && !IsNil(o.Code) {
		return true
	}

	return false
}

// SetCode gets a reference to the given int32 and assigns it to the Code field.
func (o *ResultDTOPageResultChatMessageDTO) SetCode(v int32) {
	o.Code = &v
}

// GetErrorMessage returns the ErrorMessage field value if set, zero value otherwise.
func (o *ResultDTOPageResultChatMessageDTO) GetErrorMessage() string {
	if o == nil || IsNil(o.ErrorMessage) {
		var ret string
		return ret
	}
	return *o.ErrorMessage
}

// GetErrorMessageOk returns a tuple with the ErrorMessage field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ResultDTOPageResultChatMessageDTO) GetErrorMessageOk() (*string, bool) {
	if o == nil || IsNil(o.ErrorMessage) {
		return nil, false
	}
	return o.ErrorMessage, true
}

// HasErrorMessage returns a boolean if a field has been set.
func (o *ResultDTOPageResultChatMessageDTO) HasErrorMessage() bool {
	if o != nil && !IsNil(o.ErrorMessage) {
		return true
	}

	return false
}

// SetErrorMessage gets a reference to the given string and assigns it to the ErrorMessage field.
func (o *ResultDTOPageResultChatMessageDTO) SetErrorMessage(v string) {
	o.ErrorMessage = &v
}

// GetErrorMessageKey returns the ErrorMessageKey field value if set, zero value otherwise.
func (o *ResultDTOPageResultChatMessageDTO) GetErrorMessageKey() string {
	if o == nil || IsNil(o.ErrorMessageKey) {
		var ret string
		return ret
	}
	return *o.ErrorMessageKey
}

// GetErrorMessageKeyOk returns a tuple with the ErrorMessageKey field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ResultDTOPageResultChatMessageDTO) GetErrorMessageKeyOk() (*string, bool) {
	if o == nil || IsNil(o.ErrorMessageKey) {
		return nil, false
	}
	return o.ErrorMessageKey, true
}

// HasErrorMessageKey returns a boolean if a field has been set.
func (o *ResultDTOPageResultChatMessageDTO) HasErrorMessageKey() bool {
	if o != nil && !IsNil(o.ErrorMessageKey) {
		return true
	}

	return false
}

// SetErrorMessageKey gets a reference to the given string and assigns it to the ErrorMessageKey field.
func (o *ResultDTOPageResultChatMessageDTO) SetErrorMessageKey(v string) {
	o.ErrorMessageKey = &v
}

// GetData returns the Data field value if set, zero value otherwise.
func (o *ResultDTOPageResultChatMessageDTO) GetData() PageResultChatMessageDTO {
	if o == nil || IsNil(o.Data) {
		var ret PageResultChatMessageDTO
		return ret
	}
	return *o.Data
}

// GetDataOk returns a tuple with the Data field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ResultDTOPageResultChatMessageDTO) GetDataOk() (*PageResultChatMessageDTO, bool) {
	if o == nil || IsNil(o.Data) {
		return nil, false
	}
	return o.Data, true
}

// HasData returns a boolean if a field has been set.
func (o *ResultDTOPageResultChatMessageDTO) HasData() bool {
	if o != nil && !IsNil(o.Data) {
		return true
	}

	return false
}

// SetData gets a reference to the given PageResultChatMessageDTO and assigns it to the Data field.
func (o *ResultDTOPageResultChatMessageDTO) SetData(v PageResultChatMessageDTO) {
	o.Data = &v
}

// GetSuccess returns the Success field value if set, zero value otherwise.
func (o *ResultDTOPageResultChatMessageDTO) GetSuccess() bool {
	if o == nil || IsNil(o.Success) {
		var ret bool
		return ret
	}
	return *o.Success
}

// GetSuccessOk returns a tuple with the Success field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ResultDTOPageResultChatMessageDTO) GetSuccessOk() (*bool, bool) {
	if o == nil || IsNil(o.Success) {
		return nil, false
	}
	return o.Success, true
}

// HasSuccess returns a boolean if a field has been set.
func (o *ResultDTOPageResultChatMessageDTO) HasSuccess() bool {
	if o != nil && !IsNil(o.Success) {
		return true
	}

	return false
}

// SetSuccess gets a reference to the given bool and assigns it to the Success field.
func (o *ResultDTOPageResultChatMessageDTO) SetSuccess(v bool) {
	o.Success = &v
}

func (o ResultDTOPageResultChatMessageDTO) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ResultDTOPageResultChatMessageDTO) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.RequestId) {
		toSerialize["requestId"] = o.RequestId
	}
	if !IsNil(o.Code) {
		toSerialize["code"] = o.Code
	}
	if !IsNil(o.ErrorMessage) {
		toSerialize["errorMessage"] = o.ErrorMessage
	}
	if !IsNil(o.ErrorMessageKey) {
		toSerialize["errorMessageKey"] = o.ErrorMessageKey
	}
	if !IsNil(o.Data) {
		toSerialize["data"] = o.Data
	}
	if !IsNil(o.Success) {
		toSerialize["success"] = o.Success
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *ResultDTOPageResultChatMessageDTO) UnmarshalJSON(bytes []byte) (err error) {
	varResultDTOPageResultChatMessageDTO := _ResultDTOPageResultChatMessageDTO{}

	err = json.Unmarshal(bytes, &varResultDTOPageResultChatMessageDTO)

	if err != nil {
		return err
	}

	*o = ResultDTOPageResultChatMessageDTO(varResultDTOPageResultChatMessageDTO)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "requestId")
		delete(additionalProperties, "code")
		delete(additionalProperties, "errorMessage")
		delete(additionalProperties, "errorMessageKey")
		delete(additionalProperties, "data")
		delete(additionalProperties, "success")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableResultDTOPageResultChatMessageDTO struct {
	value *ResultDTOPageResultChatMessageDTO
	isSet bool
}

func (v NullableResultDTOPageResultChatMessageDTO) Get() *ResultDTOPageResultChatMessageDTO {
	return v.value
}

func (v *NullableResultDTOPageResultChatMessageDTO) Set(val *ResultDTOPageResultChatMessageDTO) {
	v.value = val
	v.isSet = true
}

func (v NullableResultDTOPageResultChatMessageDTO) IsSet() bool {
	return v.isSet
}

func (v *NullableResultDTOPageResultChatMessageDTO) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableResultDTOPageResultChatMessageDTO(val *ResultDTOPageResultChatMessageDTO) *NullableResultDTOPageResultChatMessageDTO {
	return &NullableResultDTOPageResultChatMessageDTO{value: val, isSet: true}
}

func (v NullableResultDTOPageResultChatMessageDTO) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableResultDTOPageResultChatMessageDTO) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
