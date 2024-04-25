package xingchen

type ResultDTODocConvertDTO struct {
	RequestId       *string        `json:"requestId,omitempty"`
	Code            *int32         `json:"code,omitempty"`
	ErrorMessage    *string        `json:"errorMessage,omitempty"`
	ErrorMessageKey *string        `json:"errorMessageKey,omitempty"`
	Data            *DocConvertDTO `json:"data,omitempty"`
	Success         *bool          `json:"success,omitempty"`
}

func NewResultDTODocConvertDTO() *ResultDTODocConvertDTO {
	this := ResultDTODocConvertDTO{}
	return &this
}

func (o *ResultDTODocConvertDTO) GetRequestId() *string {
	return o.RequestId
}

func (o *ResultDTODocConvertDTO) SetRequestId(v *string) {
	o.RequestId = v
}

func (o *ResultDTODocConvertDTO) GetCode() *int32 {
	return o.Code
}

func (o *ResultDTODocConvertDTO) SetCode(v *int32) {
	o.Code = v
}

func (o *ResultDTODocConvertDTO) GetErrorMessage() *string {
	return o.ErrorMessage
}

func (o *ResultDTODocConvertDTO) SetErrorMessage(v *string) {
	o.ErrorMessage = v
}

func (o *ResultDTODocConvertDTO) GetErrorMessageKey() *string {
	return o.ErrorMessageKey
}

func (o *ResultDTODocConvertDTO) SetErrorMessageKey(v *string) {
	o.ErrorMessageKey = v
}

func (o *ResultDTODocConvertDTO) GetData() *DocConvertDTO {
	return o.Data
}

func (o *ResultDTODocConvertDTO) SetData(v *DocConvertDTO) {
	o.Data = v
}

func (o *ResultDTODocConvertDTO) GetSuccess() *bool {
	return o.Success
}

func (o *ResultDTODocConvertDTO) SetSuccess(v *bool) {
	o.Success = v
}
