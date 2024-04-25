package xingchen

type KVMemoryConfig struct {
	Enabled    *bool   `json:"enabled,omitempty"`
	MemoryText *string `json:"memoryText,omitempty"`
}

func (o *KVMemoryConfig) GetEnabled() *bool {
	return o.Enabled
}

func (o *KVMemoryConfig) SetEnabled(v *bool) {
	o.Enabled = v
}

func (o *KVMemoryConfig) GetMemoryText() *string {
	return o.MemoryText
}

func (o *KVMemoryConfig) SetMemoryText(v *string) {
	o.MemoryText = v
}
