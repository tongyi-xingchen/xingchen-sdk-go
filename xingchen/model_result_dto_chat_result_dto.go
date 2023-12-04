package xingchen

type ResultDTOChatResultDTO struct {
	RequestId       *string        `json:"requestId,omitempty"`
	Code            *int32         `json:"code,omitempty"`
	ErrorMessage    *string        `json:"message,errorMessage,omitempty"`
	ErrorMessageKey *string        `json:"errorMessageKey,omitempty"`
	Data            *ChatResultDTO `json:"data,omitempty"`
	Success         *bool          `json:"success,omitempty"`
}

func (o *ResultDTOChatResultDTO) GetRequestId() string {
	return *o.RequestId
}

func (o *ResultDTOChatResultDTO) SetRequestId(v string) {
	o.RequestId = &v
}

func (o *ResultDTOChatResultDTO) GetCode() int32 {
	return *o.Code
}

func (o *ResultDTOChatResultDTO) SetCode(v int32) {
	o.Code = &v
}

func (o *ResultDTOChatResultDTO) GetErrorMessage() string {
	return *o.ErrorMessage
}

func (o *ResultDTOChatResultDTO) SetErrorMessage(v string) {
	o.ErrorMessage = &v
}

func (o *ResultDTOChatResultDTO) GetErrorMessageKey() string {
	return *o.ErrorMessageKey
}

func (o *ResultDTOChatResultDTO) SetErrorMessageKey(v string) {
	o.ErrorMessageKey = &v
}

func (o *ResultDTOChatResultDTO) GetData() ChatResultDTO {
	return *o.Data
}

func (o *ResultDTOChatResultDTO) SetData(v ChatResultDTO) {
	o.Data = &v
}

func (o *ResultDTOChatResultDTO) GetSuccess() bool {
	return *o.Success
}

func (o *ResultDTOChatResultDTO) SetSuccess(v bool) {
	o.Success = &v
}
