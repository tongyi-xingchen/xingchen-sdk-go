package xingchen

type MemoryKVDTO struct {
	Schemas []MemorySchema `json:"schemas,omitempty"`
	Usage   *Usage         `json:"usage,omitempty"`
	Context *Context       `json:"context,omitempty"`
}

func (o *MemoryKVDTO) GetSchemas() []MemorySchema {
	return o.Schemas
}

func (o *MemoryKVDTO) SetSchemas(v []MemorySchema) {
	o.Schemas = v
}

func (o *MemoryKVDTO) GetUsage() *Usage {
	return o.Usage
}

func (o *MemoryKVDTO) SetUsage(v *Usage) {
	o.Usage = v
}

func (o *MemoryKVDTO) GetContext() *Context {
	return o.Context
}

func (o *MemoryKVDTO) SetContext(v *Context) {
	o.Context = v
}
