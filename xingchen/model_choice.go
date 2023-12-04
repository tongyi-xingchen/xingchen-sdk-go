package xingchen

type Choice struct {
	Messages   []Message `json:"messages,omitempty"`
	StopReason *string   `json:"stopReason,omitempty"`
}

func (o *Choice) GetMessages() []Message {
	return o.Messages
}

func (o *Choice) SetMessages(v []Message) {
	o.Messages = v
}

func (o *Choice) GetStopReason() string {
	return *o.StopReason
}

func (o *Choice) SetStopReason(v string) {
	o.StopReason = &v
}
