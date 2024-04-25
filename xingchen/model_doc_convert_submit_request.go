package xingchen

type DocConvertSubmitRequest struct {
	Type      *string `json:"type,omitempty"`
	SourceUrl *string `json:"sourceUrl,omitempty"`
}

func (o *DocConvertSubmitRequest) GetType() *string {
	return o.Type
}

func (o *DocConvertSubmitRequest) SetType(v *string) {
	o.Type = v
}

func (o *DocConvertSubmitRequest) GetSourceUrl() *string {
	return o.SourceUrl
}

func (o *DocConvertSubmitRequest) SetSourceUrl(v *string) {
	o.SourceUrl = v
}

func (o *DocConvertSubmitRequest) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Type) {
		toSerialize["type"] = o.Type
	}
	if !IsNil(o.SourceUrl) {
		toSerialize["sourceUrl"] = o.SourceUrl
	}

	return toSerialize, nil
}
