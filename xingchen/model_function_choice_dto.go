package xingchen

type FunctionChoice struct {
	FunctionChoiceType *string  `json:"functionChoiceType,omitempty"`
	Names              []string `json:"names,omitempty"`
}

func (o *FunctionChoice) GetFunctionChoiceType() *string {
	return o.FunctionChoiceType
}

func (o *FunctionChoice) SetFunctionChoiceType(v *string) {
	o.FunctionChoiceType = v
}

func (o *FunctionChoice) GetNames() []string {
	return o.Names
}

func (o *FunctionChoice) SetNames(v []string) {
	o.Names = v
}
