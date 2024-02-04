package xingchen

type FunctionCall struct {
	Thought     *string       `json:"thought,omitempty"`
	ApiCallList []ApiCallInfo `json:"apiCallList,omitempty"`
}

func (o *FunctionCall) GetThought() *string {
	return o.Thought
}

func (o *FunctionCall) SetThought(v *string) {
	o.Thought = v
}

func (o *FunctionCall) GetApiCallList() []ApiCallInfo {
	return o.ApiCallList
}

func (o *FunctionCall) SetApiCallList(v []ApiCallInfo) {
	o.ApiCallList = v
}
