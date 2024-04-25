package xingchen

type DocConvertDTO struct {
	Id     *int64  `json:"id,omitempty"`
	Status *string `json:"status,omitempty"`
	Url    *string `json:"url,omitempty"`
}

func (o *DocConvertDTO) GetId() *int64 {
	return o.Id
}

func (o *DocConvertDTO) SetId(v *int64) {
	o.Id = v
}

func (o *DocConvertDTO) GetStatus() *string {
	return o.Status
}

func (o *DocConvertDTO) SetStatus(v *string) {
	o.Status = v
}

func (o *DocConvertDTO) GetUrl() *string {
	return o.Url
}

func (o *DocConvertDTO) SetUrl(v *string) {
	o.Url = v
}
