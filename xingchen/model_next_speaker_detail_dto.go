package xingchen

type NextSpeakerDetailDTO struct {
	Name    *string  `json:"name,omitempty"`
	Thought *string  `json:"thought,omitempty"`
	Usage   *Usage   `json:"usage,omitempty"`
	Context *Context `json:"context,omitempty"`
}

func (o *NextSpeakerDetailDTO) GetName() *string {
	return o.Name
}

func (o *NextSpeakerDetailDTO) SetName(v *string) {
	o.Name = v
}

func (o *NextSpeakerDetailDTO) GetThought() *string {
	return o.Thought
}

func (o *NextSpeakerDetailDTO) SetThought(v *string) {
	o.Thought = v
}

func (o *NextSpeakerDetailDTO) GetUsage() *Usage {
	return o.Usage
}

func (o *NextSpeakerDetailDTO) SetUsage(v *Usage) {
	o.Usage = v
}

func (o *NextSpeakerDetailDTO) GetContext() *Context {
	return o.Context
}

func (o *NextSpeakerDetailDTO) SetContext(v *Context) {
	o.Context = v
}
