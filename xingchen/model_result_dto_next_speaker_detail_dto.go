package xingchen

type ResultDTONextSpeakerDetailDTO struct {
	RequestId       *string               `json:"requestId,omitempty"`
	Code            *int32                `json:"code,omitempty"`
	ErrorMessage    *string               `json:"errorMessage,omitempty"`
	ErrorMessageKey *string               `json:"ErrorMessageKey,omitempty"`
	Data            *NextSpeakerDetailDTO `json:"data,omitempty"`
	Success         *bool                 `json:"success,omitempty"`
}

func (o *ResultDTONextSpeakerDetailDTO) GetRequestId() *string {
	return o.RequestId
}

func (o *ResultDTONextSpeakerDetailDTO) SetRequestId(v *string) {
	o.RequestId = v
}

func (o *ResultDTONextSpeakerDetailDTO) GetCode() *int32 {
	return o.Code
}

func (o *ResultDTONextSpeakerDetailDTO) SetCode(v *int32) {
	o.Code = v
}

func (o *ResultDTONextSpeakerDetailDTO) GetErrorMessage() *string {
	return o.ErrorMessage
}

func (o *ResultDTONextSpeakerDetailDTO) SetErrorMessage(v *string) {
	o.ErrorMessage = v
}

func (o *ResultDTONextSpeakerDetailDTO) GetErrorMessageKey() *string {
	return o.ErrorMessageKey
}

func (o *ResultDTONextSpeakerDetailDTO) SetErrorMessageKey(v *string) {
	o.ErrorMessageKey = v
}

func (o *ResultDTONextSpeakerDetailDTO) GetData() *NextSpeakerDetailDTO {
	return o.Data
}

func (o *ResultDTONextSpeakerDetailDTO) SetData(v *NextSpeakerDetailDTO) {
	o.Data = v
}

func (o *ResultDTONextSpeakerDetailDTO) GetSuccess() *bool {
	return o.Success
}

func (o *ResultDTONextSpeakerDetailDTO) SetSuccess(v *bool) {
	o.Success = v
}
