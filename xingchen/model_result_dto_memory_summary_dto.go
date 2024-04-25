package xingchen

type ResultDTOMemorySummaryDTO struct {
	RequestId       *string           `json:"requestId,omitempty"`
	Code            *int32            `json:"code,omitempty"`
	ErrorMessage    *string           `json:"errorMessage,omitempty"`
	ErrorMessageKey *string           `json:"errorMessageKey,omitempty"`
	Data            *MemorySummaryDTO `json:"data,omitempty"`
	Success         *bool             `json:"success,omitempty"`
}

func (o *ResultDTOMemorySummaryDTO) GetRequestId() *string {
	return o.RequestId
}

func (o *ResultDTOMemorySummaryDTO) SetRequestId(v *string) {
	o.RequestId = v
}

func (o *ResultDTOMemorySummaryDTO) GetCode() *int32 {
	return o.Code
}

func (o *ResultDTOMemorySummaryDTO) SetCode(v *int32) {
	o.Code = v
}

func (o *ResultDTOMemorySummaryDTO) GetErrorMessage() *string {
	return o.ErrorMessage
}

func (o *ResultDTOMemorySummaryDTO) SetErrorMessage(v *string) {
	o.ErrorMessage = v
}

func (o *ResultDTOMemorySummaryDTO) GetErrorMessageKey() *string {
	return o.ErrorMessageKey
}

func (o *ResultDTOMemorySummaryDTO) SetErrorMessageKey(v *string) {
	o.ErrorMessageKey = v
}

func (o *ResultDTOMemorySummaryDTO) GetData() *MemorySummaryDTO {
	return o.Data
}

func (o *ResultDTOMemorySummaryDTO) SetData(v *MemorySummaryDTO) {
	o.Data = v
}

func (o *ResultDTOMemorySummaryDTO) GetSuccess() *bool {
	return o.Success
}

func (o *ResultDTOMemorySummaryDTO) SetSuccess(v *bool) {
	o.Success = v
}
