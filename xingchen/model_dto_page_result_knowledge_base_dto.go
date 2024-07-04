package xingchen

type ResultDTOPageResultKnowledgeBaseDTO struct {
	RequestId       *string                     `json:"requestId,omitempty"`
	Code            *int32                      `json:"code,omitempty"`
	ErrorMessage    *string                     `json:"errorMessage,omitempty"`
	ErrorMessageKey *string                     `json:"errorMessageKey,omitempty"`
	Data            *PageResultKnowledgeBaseDTO `json:"data,omitempty"`
	Success         *bool                       `json:"success,omitempty"`
}

func (o *ResultDTOPageResultKnowledgeBaseDTO) GetRequestId() *string {
	return o.RequestId
}

func (o *ResultDTOPageResultKnowledgeBaseDTO) SetRequestId(v *string) {
	o.RequestId = v
}

func (o *ResultDTOPageResultKnowledgeBaseDTO) GetCode() *int32 {
	return o.Code
}

func (o *ResultDTOPageResultKnowledgeBaseDTO) SetCode(v *int32) {
	o.Code = v
}

func (o *ResultDTOPageResultKnowledgeBaseDTO) GetErrorMessage() *string {
	return o.ErrorMessage
}

func (o *ResultDTOPageResultKnowledgeBaseDTO) SetErrorMessage(v *string) {
	o.ErrorMessage = v
}

func (o *ResultDTOPageResultKnowledgeBaseDTO) GetErrorMessageKey() *string {
	return o.ErrorMessageKey
}

func (o *ResultDTOPageResultKnowledgeBaseDTO) SetErrorMessageKey(v *string) {
	o.ErrorMessageKey = v
}

func (o *ResultDTOPageResultKnowledgeBaseDTO) GetData() *PageResultKnowledgeBaseDTO {
	return o.Data
}

func (o *ResultDTOPageResultKnowledgeBaseDTO) SetData(v *PageResultKnowledgeBaseDTO) {
	o.Data = v
}

func (o *ResultDTOPageResultKnowledgeBaseDTO) GetSuccess() *bool {
	return o.Success
}

func (o *ResultDTOPageResultKnowledgeBaseDTO) SetSuccess(v *bool) {
	o.Success = v
}
