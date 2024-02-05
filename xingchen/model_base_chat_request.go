package xingchen

type BaseChatRequest struct {
	Input       *Input           `json:"input,omitempty"`
	Model       *string          `json:"model,omitempty"`
	Parameters  *ModelParameters `json:"parameters,omitempty"`
	Servicename *string          `json:"servicename,omitempty"`
	Debug       bool             `json:"debug,omitempty"`
	Streaming   *bool            `json:"streaming,omitempty"`
}

type Input struct {
	Prompt   *string     `json:"prompt,omitempty"`
	Messages []Message   `json:"messages,omitempty"`
	Aca      interface{} `json:"aca,omitempty"`
}

type AcaChatExtParam struct {
	BotProfile     BotProfile       `json:"botProfile,omitempty"`
	Scenario       *Scenario        `json:"scenario,omitempty"`
	UserProfile    UserProfile      `json:"userProfile,omitempty"`
	SampleMessages []ChatSampleItem `json:"sampleMessages,omitempty"`
	Context        *ChatContext     `json:"context,omitempty"`
	FunctionList   []Function       `json:"functionList,omitempty"`
	FunctionChoice *FunctionChoice  `json:"functionChoice,omitempty"`
}

func chatReqParamsTOBaseChatRequest(params *ChatReqParams) *BaseChatRequest {
	return &BaseChatRequest{
		Input: &Input{
			Prompt:   PtrString("<|system|>"),
			Messages: params.Messages,
			Aca: AcaChatExtParam{
				BotProfile:     params.BotProfile,
				UserProfile:    params.UserProfile,
				SampleMessages: params.ChatSamples,
				Scenario:       params.Scenario,
				FunctionList:   params.FunctionList,
				FunctionChoice: params.FunctionChoice,
			},
		},
		Model:      params.ModelParameters.ModelName,
		Streaming:  params.Streaming,
		Parameters: params.ModelParameters,
	}
}
