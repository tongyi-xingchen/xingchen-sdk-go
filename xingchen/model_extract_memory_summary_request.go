package xingchen

type ExtractMemorySummaryRequest struct {
	Messages        []Message        `json:"messages,omitempty"`
	KVMemoryConfigs []KVMemoryConfig `json:"kvMemoryConfigs,omitempty"`
	ModelParameters *ModelParameters `json:"modelParameters,omitempty"`
}

func (o *ExtractMemorySummaryRequest) GetMessages() []Message {
	return o.Messages
}

func (o *ExtractMemorySummaryRequest) SetMessages(v []Message) {
	o.Messages = v
}

func (o *ExtractMemorySummaryRequest) GetKVMemoryConfig() []KVMemoryConfig {
	return o.KVMemoryConfigs
}

func (o *ExtractMemorySummaryRequest) SetKVMemoryConfig(v []KVMemoryConfig) {
	o.KVMemoryConfigs = v
}

func (o *ExtractMemorySummaryRequest) GetModelParameters() *ModelParameters {
	return o.ModelParameters
}

func (o *ExtractMemorySummaryRequest) SetModelParameters(v *ModelParameters) {
	o.ModelParameters = v
}

func (o *ExtractMemorySummaryRequest) ToMap() (map[string]interface{}, error) {
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
