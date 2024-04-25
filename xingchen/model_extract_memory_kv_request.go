package xingchen

type ExtractMemoryKVRequest struct {
	Messages        []Message        `json:"messages,omitempty"`
	KVMemoryConfigs []KVMemoryConfig `json:"kvMemoryConfigs,omitempty"`
	ModelParameters *ModelParameters `json:"modelParameters,omitempty"`
}

func (o *ExtractMemoryKVRequest) GetMessages() []Message {
	return o.Messages
}

func (o *ExtractMemoryKVRequest) SetMessages(v []Message) {
	o.Messages = v
}

func (o *ExtractMemoryKVRequest) GetKVMemoryConfig() []KVMemoryConfig {
	return o.KVMemoryConfigs
}

func (o *ExtractMemoryKVRequest) SetKVMemoryConfig(v []KVMemoryConfig) {
	o.KVMemoryConfigs = v
}

func (o *ExtractMemoryKVRequest) GetModelParameters() *ModelParameters {
	return o.ModelParameters
}

func (o *ExtractMemoryKVRequest) SetModelParameters(v *ModelParameters) {
	o.ModelParameters = v
}

func (o *ExtractMemoryKVRequest) ToMap() (map[string]interface{}, error) {
	toSeriale := map[string]interface{}{}
	if !IsNil(o.Messages) {
		toSeriale["messages"] = o.Messages
	}
	if !IsNil(o.KVMemoryConfigs) {
		toSeriale["kvMemoryConfig"] = o.KVMemoryConfigs
	}
	if !IsNil(o.ModelParameters) {
		toSeriale["modelParameters"] = o.ModelParameters
	}

	return toSeriale, nil
}
