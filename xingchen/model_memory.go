package xingchen

type Memory struct {
	Summaries []string       `json:"summaries,omitempty"`
	Originals []Message      `json:"originals,omitempty"`
	Tags      []MemorySchema `json:"tags,omitempty"`
}

func (o *Memory) GetSummaries() []string {
	return o.Summaries
}

func (o *Memory) SetSummaries(v []string) {
	o.Summaries = v
}

func (o *Memory) GetOriginals() []Message {
	return o.Originals
}

func (o *Memory) SetOriginals(v []Message) {
	o.Originals = v
}

func (o *Memory) GetTags() []MemorySchema {
	return o.Tags
}

func (o *Memory) SetTags(v []MemorySchema) {
	o.Tags = v
}
