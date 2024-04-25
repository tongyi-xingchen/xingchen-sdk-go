package xingchen

type LongTermMemory struct {
	Enabled         *bool            `json:"enabled,omitempty"`
	MemoryType      *string          `json:"memoryType,omitempty"`
	KVMemoryConfigs []KVMemoryConfig `json:"kvMemoryConfigs,omitempty"`
}

func (o *LongTermMemory) GetEnabled() *bool {
	return o.Enabled
}

func (o *LongTermMemory) SetEnabled(v *bool) {
	o.Enabled = v
}

func (o *LongTermMemory) GetMemoryType() *string {
	return o.MemoryType
}

func (o *LongTermMemory) SetMemoryType(v *string) {
	o.MemoryType = v
}

func (o *LongTermMemory) GetKVMemoryConfigs() []KVMemoryConfig {
	return o.KVMemoryConfigs
}

func (o *LongTermMemory) SetKVMemoryConfigs(v []KVMemoryConfig) {
	o.KVMemoryConfigs = v
}
