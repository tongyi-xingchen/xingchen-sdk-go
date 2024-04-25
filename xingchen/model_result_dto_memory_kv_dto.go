package xingchen

type ResultRTOMemoryKVDTO struct {
	RequestId       *string      `json:"requestId,omitempty"`
	Code            *int32       `json:"code,omitempty"`
	ErrorMessage    *string      `json:"errorMessage,omitempty"`
	ErrorMessageKey *string      `json:"errorMessageKey,omitempty"`
	Data            *MemoryKVDTO `json:"data,omitempty"`
	Success         *bool        `json:"success,omitempty"`
}

func (o *ResultRTOMemoryKVDTO) GetRequestId() *string {
	return o.RequestId
}

func (o *ResultRTOMemoryKVDTO) SetRequestId(v *string) {
	o.RequestId = v
}

func (o *ResultRTOMemoryKVDTO) GetCode() *int32 {
	return o.Code
}

func (o *ResultRTOMemoryKVDTO) SetCode(v *int32) {
	o.Code = v
}

func (o *ResultRTOMemoryKVDTO) GetErrorMessage() *string {
	return o.ErrorMessage
}

func (o *ResultRTOMemoryKVDTO) SetErrorMessage(v *string) {
	o.ErrorMessage = v
}

func (o *ResultRTOMemoryKVDTO) GetErrorMessageKey() *string {
	return o.ErrorMessageKey
}

func (o *ResultRTOMemoryKVDTO) SetErrorMessageKey(v *string) {
	o.ErrorMessageKey = v
}

func (o *ResultRTOMemoryKVDTO) GetData() *MemoryKVDTO {
	return o.Data
}

func (o *ResultRTOMemoryKVDTO) SetData(v *MemoryKVDTO) {
	o.Data = v
}

func (o *ResultRTOMemoryKVDTO) GetSuccess() *bool {
	return o.Success
}

func (o *ResultRTOMemoryKVDTO) SetSuccess(v *bool) {
	o.Success = v
}
