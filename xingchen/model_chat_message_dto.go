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

// checks if the ChatMessageDTO type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ChatMessageDTO{}

// ChatMessageDTO 聊天消息
type ChatMessageDTO struct {
	// 消息内容
	Content *string `json:"content,omitempty"`
	// 消息类型，opening_remarks: 开场白，sys_greetings: 系统问候，user: 用户, character: 角色
	MessageType *string `json:"messageType,omitempty"`
	// 消息ID
	MessageId     *string          `json:"messageId,omitempty"`
	MessageIssuer *ChatRoomUserDTO `json:"messageIssuer,omitempty"`
	// 消息对话session_id
	SessionId *string `json:"sessionId,omitempty"`
	// 一次对话ID
	ChatId *string `json:"chatId,omitempty"`
	// 消息创建时间
	GmtCreate *int64 `json:"gmtCreate,omitempty"`
	// 是否为开场白
	IsGreeting *bool `json:"isGreeting,omitempty"`
	// 消息评分
	Rating *float64 `json:"rating,omitempty"`
	// 重新生成次数
	RegenerateTimes      *int32 `json:"regenerateTimes,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _ChatMessageDTO ChatMessageDTO

// NewChatMessageDTO instantiates a new ChatMessageDTO object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewChatMessageDTO() *ChatMessageDTO {
	this := ChatMessageDTO{}
	return &this
}

// NewChatMessageDTOWithDefaults instantiates a new ChatMessageDTO object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewChatMessageDTOWithDefaults() *ChatMessageDTO {
	this := ChatMessageDTO{}
	return &this
}

// GetContent returns the Content field value if set, zero value otherwise.
func (o *ChatMessageDTO) GetContent() string {
	if o == nil || IsNil(o.Content) {
		var ret string
		return ret
	}
	return *o.Content
}

// GetContentOk returns a tuple with the Content field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ChatMessageDTO) GetContentOk() (*string, bool) {
	if o == nil || IsNil(o.Content) {
		return nil, false
	}
	return o.Content, true
}

// HasContent returns a boolean if a field has been set.
func (o *ChatMessageDTO) HasContent() bool {
	if o != nil && !IsNil(o.Content) {
		return true
	}

	return false
}

// SetContent gets a reference to the given string and assigns it to the Content field.
func (o *ChatMessageDTO) SetContent(v string) {
	o.Content = &v
}

// GetMessageType returns the MessageType field value if set, zero value otherwise.
func (o *ChatMessageDTO) GetMessageType() string {
	if o == nil || IsNil(o.MessageType) {
		var ret string
		return ret
	}
	return *o.MessageType
}

// GetMessageTypeOk returns a tuple with the MessageType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ChatMessageDTO) GetMessageTypeOk() (*string, bool) {
	if o == nil || IsNil(o.MessageType) {
		return nil, false
	}
	return o.MessageType, true
}

// HasMessageType returns a boolean if a field has been set.
func (o *ChatMessageDTO) HasMessageType() bool {
	if o != nil && !IsNil(o.MessageType) {
		return true
	}

	return false
}

// SetMessageType gets a reference to the given string and assigns it to the MessageType field.
func (o *ChatMessageDTO) SetMessageType(v string) {
	o.MessageType = &v
}

// GetMessageId returns the MessageId field value if set, zero value otherwise.
func (o *ChatMessageDTO) GetMessageId() string {
	if o == nil || IsNil(o.MessageId) {
		var ret string
		return ret
	}
	return *o.MessageId
}

// GetMessageIdOk returns a tuple with the MessageId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ChatMessageDTO) GetMessageIdOk() (*string, bool) {
	if o == nil || IsNil(o.MessageId) {
		return nil, false
	}
	return o.MessageId, true
}

// HasMessageId returns a boolean if a field has been set.
func (o *ChatMessageDTO) HasMessageId() bool {
	if o != nil && !IsNil(o.MessageId) {
		return true
	}

	return false
}

// SetMessageId gets a reference to the given string and assigns it to the MessageId field.
func (o *ChatMessageDTO) SetMessageId(v string) {
	o.MessageId = &v
}

// GetMessageIssuer returns the MessageIssuer field value if set, zero value otherwise.
func (o *ChatMessageDTO) GetMessageIssuer() ChatRoomUserDTO {
	if o == nil || IsNil(o.MessageIssuer) {
		var ret ChatRoomUserDTO
		return ret
	}
	return *o.MessageIssuer
}

// GetMessageIssuerOk returns a tuple with the MessageIssuer field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ChatMessageDTO) GetMessageIssuerOk() (*ChatRoomUserDTO, bool) {
	if o == nil || IsNil(o.MessageIssuer) {
		return nil, false
	}
	return o.MessageIssuer, true
}

// HasMessageIssuer returns a boolean if a field has been set.
func (o *ChatMessageDTO) HasMessageIssuer() bool {
	if o != nil && !IsNil(o.MessageIssuer) {
		return true
	}

	return false
}

// SetMessageIssuer gets a reference to the given ChatRoomUserDTO and assigns it to the MessageIssuer field.
func (o *ChatMessageDTO) SetMessageIssuer(v ChatRoomUserDTO) {
	o.MessageIssuer = &v
}

// GetSessionId returns the SessionId field value if set, zero value otherwise.
func (o *ChatMessageDTO) GetSessionId() string {
	if o == nil || IsNil(o.SessionId) {
		var ret string
		return ret
	}
	return *o.SessionId
}

// GetSessionIdOk returns a tuple with the SessionId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ChatMessageDTO) GetSessionIdOk() (*string, bool) {
	if o == nil || IsNil(o.SessionId) {
		return nil, false
	}
	return o.SessionId, true
}

// HasSessionId returns a boolean if a field has been set.
func (o *ChatMessageDTO) HasSessionId() bool {
	if o != nil && !IsNil(o.SessionId) {
		return true
	}

	return false
}

// SetSessionId gets a reference to the given string and assigns it to the SessionId field.
func (o *ChatMessageDTO) SetSessionId(v string) {
	o.SessionId = &v
}

// GetChatId returns the ChatId field value if set, zero value otherwise.
func (o *ChatMessageDTO) GetChatId() string {
	if o == nil || IsNil(o.ChatId) {
		var ret string
		return ret
	}
	return *o.ChatId
}

// GetChatIdOk returns a tuple with the ChatId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ChatMessageDTO) GetChatIdOk() (*string, bool) {
	if o == nil || IsNil(o.ChatId) {
		return nil, false
	}
	return o.ChatId, true
}

// HasChatId returns a boolean if a field has been set.
func (o *ChatMessageDTO) HasChatId() bool {
	if o != nil && !IsNil(o.ChatId) {
		return true
	}

	return false
}

// SetChatId gets a reference to the given string and assigns it to the ChatId field.
func (o *ChatMessageDTO) SetChatId(v string) {
	o.ChatId = &v
}

// GetGmtCreate returns the GmtCreate field value if set, zero value otherwise.
func (o *ChatMessageDTO) GetGmtCreate() int64 {
	if o == nil || IsNil(o.GmtCreate) {
		var ret int64
		return ret
	}
	return *o.GmtCreate
}

// GetGmtCreateOk returns a tuple with the GmtCreate field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ChatMessageDTO) GetGmtCreateOk() (*int64, bool) {
	if o == nil || IsNil(o.GmtCreate) {
		return nil, false
	}
	return o.GmtCreate, true
}

// HasGmtCreate returns a boolean if a field has been set.
func (o *ChatMessageDTO) HasGmtCreate() bool {
	if o != nil && !IsNil(o.GmtCreate) {
		return true
	}

	return false
}

// SetGmtCreate gets a reference to the given int64 and assigns it to the GmtCreate field.
func (o *ChatMessageDTO) SetGmtCreate(v int64) {
	o.GmtCreate = &v
}

// GetIsGreeting returns the IsGreeting field value if set, zero value otherwise.
func (o *ChatMessageDTO) GetIsGreeting() bool {
	if o == nil || IsNil(o.IsGreeting) {
		var ret bool
		return ret
	}
	return *o.IsGreeting
}

// GetIsGreetingOk returns a tuple with the IsGreeting field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ChatMessageDTO) GetIsGreetingOk() (*bool, bool) {
	if o == nil || IsNil(o.IsGreeting) {
		return nil, false
	}
	return o.IsGreeting, true
}

// HasIsGreeting returns a boolean if a field has been set.
func (o *ChatMessageDTO) HasIsGreeting() bool {
	if o != nil && !IsNil(o.IsGreeting) {
		return true
	}

	return false
}

// SetIsGreeting gets a reference to the given bool and assigns it to the IsGreeting field.
func (o *ChatMessageDTO) SetIsGreeting(v bool) {
	o.IsGreeting = &v
}

// GetRating returns the Rating field value if set, zero value otherwise.
func (o *ChatMessageDTO) GetRating() float64 {
	if o == nil || IsNil(o.Rating) {
		var ret float64
		return ret
	}
	return *o.Rating
}

// GetRatingOk returns a tuple with the Rating field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ChatMessageDTO) GetRatingOk() (*float64, bool) {
	if o == nil || IsNil(o.Rating) {
		return nil, false
	}
	return o.Rating, true
}

// HasRating returns a boolean if a field has been set.
func (o *ChatMessageDTO) HasRating() bool {
	if o != nil && !IsNil(o.Rating) {
		return true
	}

	return false
}

// SetRating gets a reference to the given float64 and assigns it to the Rating field.
func (o *ChatMessageDTO) SetRating(v float64) {
	o.Rating = &v
}

// GetRegenerateTimes returns the RegenerateTimes field value if set, zero value otherwise.
func (o *ChatMessageDTO) GetRegenerateTimes() int32 {
	if o == nil || IsNil(o.RegenerateTimes) {
		var ret int32
		return ret
	}
	return *o.RegenerateTimes
}

// GetRegenerateTimesOk returns a tuple with the RegenerateTimes field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ChatMessageDTO) GetRegenerateTimesOk() (*int32, bool) {
	if o == nil || IsNil(o.RegenerateTimes) {
		return nil, false
	}
	return o.RegenerateTimes, true
}

// HasRegenerateTimes returns a boolean if a field has been set.
func (o *ChatMessageDTO) HasRegenerateTimes() bool {
	if o != nil && !IsNil(o.RegenerateTimes) {
		return true
	}

	return false
}

// SetRegenerateTimes gets a reference to the given int32 and assigns it to the RegenerateTimes field.
func (o *ChatMessageDTO) SetRegenerateTimes(v int32) {
	o.RegenerateTimes = &v
}

func (o ChatMessageDTO) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ChatMessageDTO) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Content) {
		toSerialize["content"] = o.Content
	}
	if !IsNil(o.MessageType) {
		toSerialize["messageType"] = o.MessageType
	}
	if !IsNil(o.MessageId) {
		toSerialize["messageId"] = o.MessageId
	}
	if !IsNil(o.MessageIssuer) {
		toSerialize["messageIssuer"] = o.MessageIssuer
	}
	if !IsNil(o.SessionId) {
		toSerialize["sessionId"] = o.SessionId
	}
	if !IsNil(o.ChatId) {
		toSerialize["chatId"] = o.ChatId
	}
	if !IsNil(o.GmtCreate) {
		toSerialize["gmtCreate"] = o.GmtCreate
	}
	if !IsNil(o.IsGreeting) {
		toSerialize["isGreeting"] = o.IsGreeting
	}
	if !IsNil(o.Rating) {
		toSerialize["rating"] = o.Rating
	}
	if !IsNil(o.RegenerateTimes) {
		toSerialize["regenerateTimes"] = o.RegenerateTimes
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *ChatMessageDTO) UnmarshalJSON(bytes []byte) (err error) {
	varChatMessageDTO := _ChatMessageDTO{}

	err = json.Unmarshal(bytes, &varChatMessageDTO)

	if err != nil {
		return err
	}

	*o = ChatMessageDTO(varChatMessageDTO)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "content")
		delete(additionalProperties, "messageType")
		delete(additionalProperties, "messageId")
		delete(additionalProperties, "messageIssuer")
		delete(additionalProperties, "sessionId")
		delete(additionalProperties, "chatId")
		delete(additionalProperties, "gmtCreate")
		delete(additionalProperties, "isGreeting")
		delete(additionalProperties, "rating")
		delete(additionalProperties, "regenerateTimes")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableChatMessageDTO struct {
	value *ChatMessageDTO
	isSet bool
}

func (v NullableChatMessageDTO) Get() *ChatMessageDTO {
	return v.value
}

func (v *NullableChatMessageDTO) Set(val *ChatMessageDTO) {
	v.value = val
	v.isSet = true
}

func (v NullableChatMessageDTO) IsSet() bool {
	return v.isSet
}

func (v *NullableChatMessageDTO) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableChatMessageDTO(val *ChatMessageDTO) *NullableChatMessageDTO {
	return &NullableChatMessageDTO{value: val, isSet: true}
}

func (v NullableChatMessageDTO) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableChatMessageDTO) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}