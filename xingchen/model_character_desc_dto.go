package xingchen

type CharacterDescDTO struct {
	Content *string  `json:"content,omitempty"`
	Usage   *Usage   `json:"usage,omitempty"`
	Context *Context `json:"context,omitempty"`
}

func (o *CharacterDescDTO) GetContent() *string {
	return o.Content
}

func (o *CharacterDescDTO) SetContent(v *string) {
	o.Content = v
}

func (o *CharacterDescDTO) GetUsage() *Usage {
	return o.Usage
}

func (o *CharacterDescDTO) SetUsage(v *Usage) {
	o.Usage = v
}

func (o *CharacterDescDTO) GetContext() *Context {
	return o.Context
}

func (o *CharacterDescDTO) SetContext(v *Context) {
	o.Context = v
}
