package xingchen

type ResultDTOKnowledgeBaseDTO struct {
	RequestId       *string           `json:"requestId,omitempty"`
	Code            *int32            `json:"code,omitempty"`
	ErrorMessage    *string           `json:"errorMessage,omitempty"`
	ErrorMessageKey *string           `json:"errorMessageKey,omitempty"`
	Data            *KnowledgeBaseDTO `json:"data,omitempty"`
	Success         *bool             `json:"success,omitempty"`
}

func (o *ResultDTOKnowledgeBaseDTO) GetRequestId() *string {
	return o.RequestId
}

func (o *ResultDTOKnowledgeBaseDTO) SetRequestId(v *string) {
	o.RequestId = v
}

func (o *ResultDTOKnowledgeBaseDTO) GetCode() *int32 {
	return o.Code
}

func (o *ResultDTOKnowledgeBaseDTO) SetCode(v *int32) {
	o.Code = v
}

func (o *ResultDTOKnowledgeBaseDTO) GetErrorMessage() *string {
	return o.ErrorMessage
}

func (o *ResultDTOKnowledgeBaseDTO) SetErrorMessage(v *string) {
	o.ErrorMessage = v
}

func (o *ResultDTOKnowledgeBaseDTO) GetErrorMessageKey() *string {
	return o.ErrorMessageKey
}

func (o *ResultDTOKnowledgeBaseDTO) SetErrorMessageKey(v *string) {
	o.ErrorMessageKey = v
}

func (o *ResultDTOKnowledgeBaseDTO) GetData() *KnowledgeBaseDTO {
	return o.Data
}

func (o *ResultDTOKnowledgeBaseDTO) SetData(v *KnowledgeBaseDTO) {
	o.Data = v
}

func (o *ResultDTOKnowledgeBaseDTO) GetSuccess() *bool {
	return o.Success
}

func (o *ResultDTOKnowledgeBaseDTO) SetSuccess(v *bool) {
	o.Success = v
}
