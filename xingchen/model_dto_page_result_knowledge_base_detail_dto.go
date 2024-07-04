package xingchen

type ResultDTOPageResultKnowledgeBaseDetailDTO struct {
	RequestId       *string                           `json:"requestId,omitempty"`
	Code            *int32                            `json:"code,omitempty"`
	ErrorMessage    *string                           `json:"errorMessage,omitempty"`
	ErrorMessageKey *string                           `json:"errorMessageKey,omitempty"`
	Data            *PageResultKnowledgeBaseDetailDTO `json:"data,omitempty"`
	Success         *bool                             `json:"success,omitempty"`
}

func (o *ResultDTOPageResultKnowledgeBaseDetailDTO) GetRequestId() *string {
	return o.RequestId
}

func (o *ResultDTOPageResultKnowledgeBaseDetailDTO) SetRequestId(v *string) {
	o.RequestId = v
}

func (o *ResultDTOPageResultKnowledgeBaseDetailDTO) GetCode() *int32 {
	return o.Code
}

func (o *ResultDTOPageResultKnowledgeBaseDetailDTO) SetCode(v *int32) {
	o.Code = v
}

func (o *ResultDTOPageResultKnowledgeBaseDetailDTO) GetErrorMessage() *string {
	return o.ErrorMessage
}

func (o *ResultDTOPageResultKnowledgeBaseDetailDTO) SetErrorMessage(v *string) {
	o.ErrorMessage = v
}

func (o *ResultDTOPageResultKnowledgeBaseDetailDTO) GetErrorMessageKey() *string {
	return o.ErrorMessageKey
}

func (o *ResultDTOPageResultKnowledgeBaseDetailDTO) SetErrorMessageKey(v *string) {
	o.ErrorMessageKey = v
}

func (o *ResultDTOPageResultKnowledgeBaseDetailDTO) GetData() *PageResultKnowledgeBaseDetailDTO {
	return o.Data
}

func (o *ResultDTOPageResultKnowledgeBaseDetailDTO) SetData(v *PageResultKnowledgeBaseDetailDTO) {
	o.Data = v
}

func (o *ResultDTOPageResultKnowledgeBaseDetailDTO) GetSuccess() *bool {
	return o.Success
}

func (o *ResultDTOPageResultKnowledgeBaseDetailDTO) SetSuccess(v *bool) {
	o.Success = v
}
