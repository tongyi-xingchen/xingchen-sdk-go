package xingchen

type MemorySummaryDTO struct {
	MemoryResults []MemorySummary `json:"memoryResults,omitempty"`
	Usage         *Usage          `json:"usage,omitempty"`
	Context       *Context        `json:"context,omitempty"`
}

func (o *MemorySummaryDTO) GetMemoryResults() []MemorySummary {
	return o.MemoryResults
}

func (o *MemorySummaryDTO) SetMemoryResults(v []MemorySummary) {
	o.MemoryResults = v
}

func (o *MemorySummaryDTO) GetUsage() *Usage {
	return o.Usage
}

func (o *MemorySummaryDTO) SetUsage(v *Usage) {
	o.Usage = v
}

func (o *MemorySummaryDTO) GetContext() *Context {
	return o.Context
}

func (o *MemorySummaryDTO) SetContext(v *Context) {
	o.Context = v
}
