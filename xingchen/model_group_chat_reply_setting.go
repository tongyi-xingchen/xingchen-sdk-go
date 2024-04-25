package xingchen

type GroupChatReplySetting struct {
	BotName *string `json:"botName,omitempty"`
	Thought *string `json:"thought,omitempty"`
}

func (o *GroupChatReplySetting) GetBotName() *string {
	return o.BotName
}

func (o *GroupChatReplySetting) SetBotName(v *string) {
	o.BotName = v
}

func (o *GroupChatReplySetting) GetThought() *string {
	return o.Thought
}

func (o *GroupChatReplySetting) SetThought(v *string) {
	o.Thought = v
}
