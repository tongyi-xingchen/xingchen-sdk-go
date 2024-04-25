package xingchen

type Keyword struct {
	Value *string `json:"value,omitempty"`
}

type RejectCondition struct {
	Enabled            *bool            `json:"enabled,omitempty"`
	ConditionType      *string          `json:"conditionType,omitempty"`
	Keywords           []Keyword        `json:"keywords,omitempty"`
	SubRejectCondition *RejectCondition `json:"subRejectCondition,omitempty"`
}

func (o *RejectCondition) GetEnabled() *bool {
	return o.Enabled
}

func (o *RejectCondition) SetEnabled(v *bool) {
	o.Enabled = v
}

func (o *RejectCondition) GetConditionType() *string {
	return o.ConditionType
}

func (o *RejectCondition) SetConditionType(v *string) {
	o.ConditionType = v
}

func (o *RejectCondition) GetKeywords() []Keyword {
	return o.Keywords
}

func (o *RejectCondition) SetKeywords(v []Keyword) {
	o.Keywords = v
}

func (o *RejectCondition) GetSubRejectCondition() *RejectCondition {
	return o.SubRejectCondition
}

func (o *RejectCondition) SetSubRejectCondition(v *RejectCondition) {
	o.SubRejectCondition = v
}
