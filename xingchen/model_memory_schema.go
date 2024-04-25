package xingchen

type MemorySchema struct {
	Role  *string  `json:"role,omitempty"`
	Key   *string  `json:"key,omitempty"`
	Value []string `json:"value,omitempty"`
}

func (o *MemorySchema) GetRole() *string {
	return o.Role
}

func (o *MemorySchema) SetRole(v *string) {
	o.Role = v
}

func (o *MemorySchema) GetKey() *string {
	return o.Key
}

func (o *MemorySchema) SetKey(v *string) {
	o.Key = v
}

func (o *MemorySchema) GetValue() []string {
	return o.Value
}

func (o *MemorySchema) SetValue(v []string) {
	o.Value = v
}
