package xingchen

type Function struct {
	Name        *string      `json:"name,omitempty"`
	Description *string      `json:"description,omitempty"`
	Parameters  *interface{} `json:"parameters,omitempty"`
}

func (o *Function) GetName() *string {
	return o.Name
}

func (o *Function) SetName(v *string) {
	o.Name = v
}

func (o *Function) GetDescription() *string {
	return o.Description
}

func (o *Function) SetDescription(v *string) {
	o.Description = v
}

func (o *Function) GetParameters() *interface{} {
	return o.Parameters
}

func (o *Function) SetParameters(v *interface{}) {
	o.Parameters = v
}
