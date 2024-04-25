package xingchen

type StopChatRequest struct {
	SessionId *string `json:"sessionId,omitempty"`
	RequestId *string `json:"requestId,omitempty"`
	AnswerId  *string `json:"answerId,omitempty"`
	Position  *int32  `json:"position,omitempty"`
	Content   *string `json:"content,omitempty"`
}

func (o *StopChatRequest) GetSessionId() *string {
	return o.SessionId
}

func (o *StopChatRequest) SetSessionId(v *string) {
	o.SessionId = v
}

func (o *StopChatRequest) GetRequestId() *string {
	return o.RequestId
}

func (o *StopChatRequest) SetRequestId(v *string) {
	o.RequestId = v
}

func (o *StopChatRequest) GetAnswerId() *string {
	return o.AnswerId
}

func (o *StopChatRequest) SetAnswerId(v *string) {
	o.AnswerId = v
}

func (o *StopChatRequest) GetPosition() *int32 {
	return o.Position
}

func (o *StopChatRequest) SetPosition(v *int32) {
	o.Position = v
}

func (o *StopChatRequest) GetContent() *string {
	return o.Content
}

func (o *StopChatRequest) SetContent(v *string) {
	o.Content = v
}
