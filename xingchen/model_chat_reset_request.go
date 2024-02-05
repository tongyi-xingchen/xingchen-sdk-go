package xingchen

import "encoding/json"

var _ MappedNullable = &ChatResetRequest{}

type ChatResetRequest struct {
	GatewayAddContent    *GatewayIssuedParams `json:"gatewayAddContent,omitempty"`
	CharacterId          *string              `json:"characterId,omitempty"`
	UserId               *string              `json:"userId,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _ChatResetRequest ChatResetRequest

func NewChatResetRequest(characterId *string, userId *string) *ChatResetRequest {
	this := ChatResetRequest{}
	this.CharacterId = characterId
	this.UserId = userId
	return &this
}

func NewChatResetRequestWithDefaults() *ChatResetRequest {
	this := ChatResetRequest{}
	return &this
}

func (o *ChatResetRequest) GetGatewayAddContent() GatewayIssuedParams {
	if o == nil || IsNil(o.GatewayAddContent) {
		var ret GatewayIssuedParams
		return ret
	}
	return *o.GatewayAddContent
}

func (o *ChatResetRequest) GetGatewayAddContentOk() (*GatewayIssuedParams, bool) {
	if o == nil || IsNil(o.GatewayAddContent) {
		return nil, false
	}
	return o.GatewayAddContent, true
}

func (o *ChatResetRequest) HasGatewayAddContent() bool {
	if o != nil && !IsNil(o.GatewayAddContent) {
		return true
	}

	return false
}

func (o *ChatResetRequest) SetGatewayAddContent(v GatewayIssuedParams) {
	o.GatewayAddContent = &v
}

func (o *ChatResetRequest) GetCharacterId() string {
	if o == nil || IsNil(o.CharacterId) {
		var ret string
		return ret
	}
	return *o.CharacterId
}

func (o *ChatResetRequest) GetCharacterIdOk() (*string, bool) {
	if o == nil || IsNil(o.CharacterId) {
		return nil, false
	}
	return o.CharacterId, true
}

func (o *ChatResetRequest) HasCharacterId() bool {
	if o != nil && !IsNil(o.CharacterId) {
		return true
	}

	return false
}

func (o *ChatResetRequest) SetCharacterId(v string) {
	o.CharacterId = &v
}

func (o *ChatResetRequest) GetUserId() string {
	if o == nil || IsNil(o.UserId) {
		var ret string
		return ret
	}
	return *o.UserId
}

func (o *ChatResetRequest) GetUserIdOk() (*string, bool) {
	if o == nil || IsNil(o.UserId) {
		return nil, false
	}
	return o.UserId, true
}

func (o *ChatResetRequest) HasUserId() bool {
	if o != nil && !IsNil(o.UserId) {
		return true
	}

	return false
}

func (o *ChatResetRequest) SetUserId(v string) {
	o.UserId = &v
}

func (o ChatResetRequest) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ChatResetRequest) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.GatewayAddContent) {
		toSerialize["gatewayAddContent"] = o.GatewayAddContent
	}
	toSerialize["characterId"] = o.CharacterId
	toSerialize["userId"] = o.UserId

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *ChatResetRequest) UnmarshalJSON(bytes []byte) (err error) {
	varchatResetRequest := _ChatResetRequest{}

	err = json.Unmarshal(bytes, &varchatResetRequest)

	if err != nil {
		return err
	}

	*o = ChatResetRequest(varchatResetRequest)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "gatewayAddContent")
		delete(additionalProperties, "characterId")
		delete(additionalProperties, "userId")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableChatResetRequest struct {
	value *ChatResetRequest
	isSet bool
}

func (v NullableChatResetRequest) Get() *ChatResetRequest {
	return v.value
}

func (v *NullableChatResetRequest) Set(val *ChatResetRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableChatResetRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableChatResetRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableChatResetRequest(val *ChatResetRequest) *NullableChatResetRequest {
	return &NullableChatResetRequest{value: val, isSet: true}
}

func (v NullableChatResetRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableChatResetRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
