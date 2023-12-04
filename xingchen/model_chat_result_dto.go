package xingchen

type ChatResultDTO struct {
	RequestId *string  `json:",omitempty"`
	Stop      bool     `json:"stop,omitempty"`
	Choices   []Choice `json:"choices,omitempty"`
	Usage     *Usage   `json:"usage,omitempty"`
	Context   *Context `json:"context,omitempty"`
}

func (o *ChatResultDTO) GetChoices() []Choice {
	return o.Choices
}

func (o *ChatResultDTO) SetChoices(v []Choice) {
	o.Choices = v
}

func (o *ChatResultDTO) GetUsage() Usage {
	return *o.Usage
}

func (o *ChatResultDTO) SetUsage(v Usage) {
	o.Usage = &v
}

func (o *ChatResultDTO) GetContext() Context {
	return *o.Context
}

func (o *ChatResultDTO) SetContext(v Context) {
	o.Context = &v
}
