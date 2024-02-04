package xingchen

type ApiCallInfo struct {
	ApiName    *string      `json:"apiName,omitempty"`
	Parameters *interface{} `json:"parameters,omitempty"`
}

func (o *ApiCallInfo) GetApiName() *string {
	return o.ApiName
}

func (o *ApiCallInfo) SetApiName(v *string) {
	o.ApiName = v
}

func (o *ApiCallInfo) GetParameters() *interface{} {
	return o.Parameters
}

func (o *ApiCallInfo) SetParameters(v *interface{}) {
	o.Parameters = v
}
