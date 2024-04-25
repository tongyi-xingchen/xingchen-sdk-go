package xingchen

type ChatContext struct {
	UseChatHistory bool    `json:"useChatHistory"`
	IsRegenerate   bool    `json:"isRegenerate"`
	QueryId        *string `json:"queryId,omitempty"`
	AnswerId       *string `json:"answerId,omitempty"`
	ResultCount    *int32  `json:"resultCount,omitempty"`
}

func (o *ChatContext) GetUseChatHistory() bool {
	return o.UseChatHistory
}

func (o *ChatContext) SetUseChatHistory(v bool) {
	o.UseChatHistory = v
}

func (o *ChatContext) GetIsRegenerate() bool {
	return o.IsRegenerate
}

func (o *ChatContext) SetIsRegenerate(v bool) {
	o.IsRegenerate = v
}

func (o *ChatContext) GetQueryId() *string {
	return o.QueryId
}

func (o *ChatContext) SetQueryId(v *string) {
	o.QueryId = v
}

func (o *ChatContext) GetAnswerId() *string {
	return o.AnswerId
}

func (o *ChatContext) SetAnswerId(v *string) {
	o.AnswerId = v
}

func (o *ChatContext) GetResultCount() *int32 {
	return o.ResultCount
}

func (o *ChatContext) SetResultCount(v *int32) {
	o.ResultCount = v
}
