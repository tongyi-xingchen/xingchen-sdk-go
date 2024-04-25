package xingchen

type ResultDTOPollTextToImageResultDTO struct {
	RequestId       *string               `json:"requestId,omitempty"`
	Code            *int32                `json:"code,omitempty"`
	ErrorMessage    *string               `json:"errorMessage,omitempty"`
	ErrorMessageKey *string               `json:"errorMessageKey,omitempty"`
	Data            *TextToImageResultDTO `json:"data,omitempty"`
	Success         *bool                 `json:"success,omitempty"`
}

func (o *ResultDTOPollTextToImageResultDTO) GetRequestId() *string {
	return o.RequestId
}

func (o *ResultDTOPollTextToImageResultDTO) SetRequestId(v *string) {
	o.RequestId = v
}

func (o *ResultDTOPollTextToImageResultDTO) GetCode() *int32 {
	return o.Code
}

func (o *ResultDTOPollTextToImageResultDTO) SetCode(v *int32) {
	o.Code = v
}

func (o *ResultDTOPollTextToImageResultDTO) GetErrorMessage() *string {
	return o.ErrorMessage
}

func (o *ResultDTOPollTextToImageResultDTO) SetErrorMessage(v *string) {
	o.ErrorMessage = v
}

func (o *ResultDTOPollTextToImageResultDTO) GetErrorMessageKey() *string {
	return o.ErrorMessageKey
}

func (o *ResultDTOPollTextToImageResultDTO) SetErrorMessageKey(v *string) {
	o.ErrorMessageKey = v
}

func (o *ResultDTOPollTextToImageResultDTO) GetData() *TextToImageResultDTO {
	return o.Data
}

func (o *ResultDTOPollTextToImageResultDTO) SetData(v *TextToImageResultDTO) {
	o.Data = v
}

func (o *ResultDTOPollTextToImageResultDTO) GetSuccess() *bool {
	return o.Success
}

func (o *ResultDTOPollTextToImageResultDTO) SetSuccess(v *bool) {
	o.Success = v
}
