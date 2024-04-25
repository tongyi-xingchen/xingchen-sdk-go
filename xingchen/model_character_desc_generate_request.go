package xingchen

type CharacterDescGenerateRequest struct {
	Type            *string          `json:"type,omitempty"`
	FileUrl         *string          `json:"fileUrl,omitempty"`
	Text            *string          `json:"text,omitempty"`
	FileName        *string          `json:"fileName,omitempty"`
	ModelParameters *ModelParameters `json:"modelParameters,omitempty"`
}

func (o *CharacterDescGenerateRequest) GetType() *string {
	return o.Type
}

func (o *CharacterDescGenerateRequest) SetType(v *string) {
	o.Type = v
}

func (o *CharacterDescGenerateRequest) GetFileUrl() *string {
	return o.FileUrl
}

func (o *CharacterDescGenerateRequest) SetFileUrl(v *string) {
	o.FileUrl = v
}

func (o *CharacterDescGenerateRequest) GetText() *string {
	return o.Text
}

func (o *CharacterDescGenerateRequest) SetText(v *string) {
	o.Text = v
}

func (o *CharacterDescGenerateRequest) GetFileName() *string {
	return o.FileName
}

func (o *CharacterDescGenerateRequest) SetFileName(v *string) {
	o.FileName = v
}

func (o *CharacterDescGenerateRequest) GetModelParameters() *ModelParameters {
	return o.ModelParameters
}

func (o *CharacterDescGenerateRequest) SetModelParameters(v *ModelParameters) {
	o.ModelParameters = v
}

func (o *CharacterDescGenerateRequest) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Type) {
		toSerialize["type"] = o.Type
	}
	if !IsNil(o.FileUrl) {
		toSerialize["fileUrl"] = o.FileUrl
	}
	if !IsNil(o.Text) {
		toSerialize["text"] = o.Text
	}
	if !IsNil(o.FileName) {
		toSerialize["fileName"] = o.FileName
	}
	if !IsNil(o.ModelParameters) {
		toSerialize["modelParameters"] = o.ModelParameters
	}

	return toSerialize, nil
}
