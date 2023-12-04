package xingchen

type Usage struct {
	OutputTokens int32 `json:"outputTokens,omitempty"`
	InputTokens  int32 `json:"inputTokens,omitempty"`
	UserTokens   int32 `json:"userTokens,omitempty"`
}

func (o *Usage) GetOutputTokens() int32 {
	return o.OutputTokens
}

func (o *Usage) SetOutputTokens(v int32) {
	o.OutputTokens = v
}

func (o *Usage) GetInputTokens() int32 {
	return o.InputTokens
}

func (o *Usage) SetInputTokens(v int32) {
	o.InputTokens = v
}

func (o *Usage) GetUserTokens() int32 {
	return o.UserTokens
}

func (o *Usage) SetUserTokens(v int32) {
	o.UserTokens = v
}
