package xingchen

type MemorySummary struct {
	Content *string `json:"content,omitempty"`
}

func (o *MemorySummary) GetContent() *string {
	return o.Content
}

func (o *MemorySummary) SetContent(v *string) {
	o.Content = v
}
