# ChatReqParams

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**GatewayAddContent** | Pointer to [**GatewayIssuedParams**](GatewayIssuedParams.md) |  | [optional] 
**BotProfile** | [**BotProfile**](BotProfile.md) |  | 
**Messages** | [**[]Message**](Message.md) |  | 
**ChatSamples** | Pointer to [**[]ChatSampleItem**](ChatSampleItem.md) |  | [optional] 
**AdvancedSettings** | Pointer to [**AdvancedSettings**](AdvancedSettings.md) |  | [optional] 
**ModelParameters** | Pointer to [**ModelParameters**](ModelParameters.md) |  | [optional] 
**UserProfile** | [**UserProfile**](UserProfile.md) |  | 
**Scenario** | Pointer to [**Scenario**](Scenario.md) |  | [optional] 
**Streaming** | Pointer to **bool** |  | [optional] 
**Context** | Pointer to [**Context**](Context.md) |  | [optional] 
**Source** | Pointer to **string** |  | [optional] 

## Methods

### NewChatReqParams

`func NewChatReqParams(botProfile BotProfile, messages []Message, userProfile UserProfile, ) *ChatReqParams`

NewChatReqParams instantiates a new ChatReqParams object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewChatReqParamsWithDefaults

`func NewChatReqParamsWithDefaults() *ChatReqParams`

NewChatReqParamsWithDefaults instantiates a new ChatReqParams object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetGatewayAddContent

`func (o *ChatReqParams) GetGatewayAddContent() GatewayIssuedParams`

GetGatewayAddContent returns the GatewayAddContent field if non-nil, zero value otherwise.

### GetGatewayAddContentOk

`func (o *ChatReqParams) GetGatewayAddContentOk() (*GatewayIssuedParams, bool)`

GetGatewayAddContentOk returns a tuple with the GatewayAddContent field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGatewayAddContent

`func (o *ChatReqParams) SetGatewayAddContent(v GatewayIssuedParams)`

SetGatewayAddContent sets GatewayAddContent field to given value.

### HasGatewayAddContent

`func (o *ChatReqParams) HasGatewayAddContent() bool`

HasGatewayAddContent returns a boolean if a field has been set.

### GetBotProfile

`func (o *ChatReqParams) GetBotProfile() BotProfile`

GetBotProfile returns the BotProfile field if non-nil, zero value otherwise.

### GetBotProfileOk

`func (o *ChatReqParams) GetBotProfileOk() (*BotProfile, bool)`

GetBotProfileOk returns a tuple with the BotProfile field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBotProfile

`func (o *ChatReqParams) SetBotProfile(v BotProfile)`

SetBotProfile sets BotProfile field to given value.


### GetMessages

`func (o *ChatReqParams) GetMessages() []Message`

GetMessages returns the Messages field if non-nil, zero value otherwise.

### GetMessagesOk

`func (o *ChatReqParams) GetMessagesOk() (*[]Message, bool)`

GetMessagesOk returns a tuple with the Messages field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMessages

`func (o *ChatReqParams) SetMessages(v []Message)`

SetMessages sets Messages field to given value.


### GetChatSamples

`func (o *ChatReqParams) GetChatSamples() []ChatSampleItem`

GetChatSamples returns the ChatSamples field if non-nil, zero value otherwise.

### GetChatSamplesOk

`func (o *ChatReqParams) GetChatSamplesOk() (*[]ChatSampleItem, bool)`

GetChatSamplesOk returns a tuple with the ChatSamples field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetChatSamples

`func (o *ChatReqParams) SetChatSamples(v []ChatSampleItem)`

SetChatSamples sets ChatSamples field to given value.

### HasChatSamples

`func (o *ChatReqParams) HasChatSamples() bool`

HasChatSamples returns a boolean if a field has been set.

### GetAdvancedSettings

`func (o *ChatReqParams) GetAdvancedSettings() AdvancedSettings`

GetAdvancedSettings returns the AdvancedSettings field if non-nil, zero value otherwise.

### GetAdvancedSettingsOk

`func (o *ChatReqParams) GetAdvancedSettingsOk() (*AdvancedSettings, bool)`

GetAdvancedSettingsOk returns a tuple with the AdvancedSettings field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAdvancedSettings

`func (o *ChatReqParams) SetAdvancedSettings(v AdvancedSettings)`

SetAdvancedSettings sets AdvancedSettings field to given value.

### HasAdvancedSettings

`func (o *ChatReqParams) HasAdvancedSettings() bool`

HasAdvancedSettings returns a boolean if a field has been set.

### GetModelParameters

`func (o *ChatReqParams) GetModelParameters() ModelParameters`

GetModelParameters returns the ModelParameters field if non-nil, zero value otherwise.

### GetModelParametersOk

`func (o *ChatReqParams) GetModelParametersOk() (*ModelParameters, bool)`

GetModelParametersOk returns a tuple with the ModelParameters field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetModelParameters

`func (o *ChatReqParams) SetModelParameters(v ModelParameters)`

SetModelParameters sets ModelParameters field to given value.

### HasModelParameters

`func (o *ChatReqParams) HasModelParameters() bool`

HasModelParameters returns a boolean if a field has been set.

### GetUserProfile

`func (o *ChatReqParams) GetUserProfile() UserProfile`

GetUserProfile returns the UserProfile field if non-nil, zero value otherwise.

### GetUserProfileOk

`func (o *ChatReqParams) GetUserProfileOk() (*UserProfile, bool)`

GetUserProfileOk returns a tuple with the UserProfile field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserProfile

`func (o *ChatReqParams) SetUserProfile(v UserProfile)`

SetUserProfile sets UserProfile field to given value.


### GetScenario

`func (o *ChatReqParams) GetScenario() Scenario`

GetScenario returns the Scenario field if non-nil, zero value otherwise.

### GetScenarioOk

`func (o *ChatReqParams) GetScenarioOk() (*Scenario, bool)`

GetScenarioOk returns a tuple with the Scenario field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScenario

`func (o *ChatReqParams) SetScenario(v Scenario)`

SetScenario sets Scenario field to given value.

### HasScenario

`func (o *ChatReqParams) HasScenario() bool`

HasScenario returns a boolean if a field has been set.

### GetStreaming

`func (o *ChatReqParams) GetStreaming() bool`

GetStreaming returns the Streaming field if non-nil, zero value otherwise.

### GetStreamingOk

`func (o *ChatReqParams) GetStreamingOk() (*bool, bool)`

GetStreamingOk returns a tuple with the Streaming field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStreaming

`func (o *ChatReqParams) SetStreaming(v bool)`

SetStreaming sets Streaming field to given value.

### HasStreaming

`func (o *ChatReqParams) HasStreaming() bool`

HasStreaming returns a boolean if a field has been set.

### GetContext

`func (o *ChatReqParams) GetContext() Context`

GetContext returns the Context field if non-nil, zero value otherwise.

### GetContextOk

`func (o *ChatReqParams) GetContextOk() (*Context, bool)`

GetContextOk returns a tuple with the Context field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContext

`func (o *ChatReqParams) SetContext(v Context)`

SetContext sets Context field to given value.

### HasContext

`func (o *ChatReqParams) HasContext() bool`

HasContext returns a boolean if a field has been set.

### GetSource

`func (o *ChatReqParams) GetSource() string`

GetSource returns the Source field if non-nil, zero value otherwise.

### GetSourceOk

`func (o *ChatReqParams) GetSourceOk() (*string, bool)`

GetSourceOk returns a tuple with the Source field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSource

`func (o *ChatReqParams) SetSource(v string)`

SetSource sets Source field to given value.

### HasSource

`func (o *ChatReqParams) HasSource() bool`

HasSource returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


