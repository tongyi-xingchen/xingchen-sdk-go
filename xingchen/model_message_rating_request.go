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

// checks if the MessageRatingRequest type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &MessageRatingRequest{}

// MessageRatingRequest struct for MessageRatingRequest
type MessageRatingRequest struct {
	GatewayAddContent *GatewayIssuedParams `json:"gatewayAddContent,omitempty"`
	// 待评分消息ID
	MessageId string `json:"messageId"`
	// 评分
	Rating int32 `json:"rating"`
	// 差评的具体信息
	Info                 *string `json:"info,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _MessageRatingRequest MessageRatingRequest

// NewMessageRatingRequest instantiates a new MessageRatingRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewMessageRatingRequest(messageId string, rating int32) *MessageRatingRequest {
	this := MessageRatingRequest{}
	this.MessageId = messageId
	this.Rating = rating
	return &this
}

// NewMessageRatingRequestWithDefaults instantiates a new MessageRatingRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewMessageRatingRequestWithDefaults() *MessageRatingRequest {
	this := MessageRatingRequest{}
	return &this
}

// GetGatewayAddContent returns the GatewayAddContent field value if set, zero value otherwise.
func (o *MessageRatingRequest) GetGatewayAddContent() GatewayIssuedParams {
	if o == nil || IsNil(o.GatewayAddContent) {
		var ret GatewayIssuedParams
		return ret
	}
	return *o.GatewayAddContent
}

// GetGatewayAddContentOk returns a tuple with the GatewayAddContent field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MessageRatingRequest) GetGatewayAddContentOk() (*GatewayIssuedParams, bool) {
	if o == nil || IsNil(o.GatewayAddContent) {
		return nil, false
	}
	return o.GatewayAddContent, true
}

// HasGatewayAddContent returns a boolean if a field has been set.
func (o *MessageRatingRequest) HasGatewayAddContent() bool {
	if o != nil && !IsNil(o.GatewayAddContent) {
		return true
	}

	return false
}

// SetGatewayAddContent gets a reference to the given GatewayIssuedParams and assigns it to the GatewayAddContent field.
func (o *MessageRatingRequest) SetGatewayAddContent(v GatewayIssuedParams) {
	o.GatewayAddContent = &v
}

// GetMessageId returns the MessageId field value
func (o *MessageRatingRequest) GetMessageId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.MessageId
}

// GetMessageIdOk returns a tuple with the MessageId field value
// and a boolean to check if the value has been set.
func (o *MessageRatingRequest) GetMessageIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.MessageId, true
}

// SetMessageId sets field value
func (o *MessageRatingRequest) SetMessageId(v string) {
	o.MessageId = v
}

// GetRating returns the Rating field value
func (o *MessageRatingRequest) GetRating() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.Rating
}

// GetRatingOk returns a tuple with the Rating field value
// and a boolean to check if the value has been set.
func (o *MessageRatingRequest) GetRatingOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Rating, true
}

// SetRating sets field value
func (o *MessageRatingRequest) SetRating(v int32) {
	o.Rating = v
}

// GetInfo returns the Info field value if set, zero value otherwise.
func (o *MessageRatingRequest) GetInfo() string {
	if o == nil || IsNil(o.Info) {
		var ret string
		return ret
	}
	return *o.Info
}

// GetInfoOk returns a tuple with the Info field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MessageRatingRequest) GetInfoOk() (*string, bool) {
	if o == nil || IsNil(o.Info) {
		return nil, false
	}
	return o.Info, true
}

// HasInfo returns a boolean if a field has been set.
func (o *MessageRatingRequest) HasInfo() bool {
	if o != nil && !IsNil(o.Info) {
		return true
	}

	return false
}

// SetInfo gets a reference to the given string and assigns it to the Info field.
func (o *MessageRatingRequest) SetInfo(v string) {
	o.Info = &v
}

func (o MessageRatingRequest) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o MessageRatingRequest) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.GatewayAddContent) {
		toSerialize["gatewayAddContent"] = o.GatewayAddContent
	}
	toSerialize["messageId"] = o.MessageId
	toSerialize["rating"] = o.Rating
	if !IsNil(o.Info) {
		toSerialize["info"] = o.Info
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *MessageRatingRequest) UnmarshalJSON(bytes []byte) (err error) {
	varMessageRatingRequest := _MessageRatingRequest{}

	err = json.Unmarshal(bytes, &varMessageRatingRequest)

	if err != nil {
		return err
	}

	*o = MessageRatingRequest(varMessageRatingRequest)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "gatewayAddContent")
		delete(additionalProperties, "messageId")
		delete(additionalProperties, "rating")
		delete(additionalProperties, "info")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableMessageRatingRequest struct {
	value *MessageRatingRequest
	isSet bool
}

func (v NullableMessageRatingRequest) Get() *MessageRatingRequest {
	return v.value
}

func (v *NullableMessageRatingRequest) Set(val *MessageRatingRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableMessageRatingRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableMessageRatingRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableMessageRatingRequest(val *MessageRatingRequest) *NullableMessageRatingRequest {
	return &NullableMessageRatingRequest{value: val, isSet: true}
}

func (v NullableMessageRatingRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableMessageRatingRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}