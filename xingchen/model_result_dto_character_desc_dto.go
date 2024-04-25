package xingchen

type ResultDTOCharacterDescDTO struct {
	RequestId       *string           `json:"requestId,omitempty"`
	Code            *int32            `json:"code,omitempty"`
	ErrorMessage    *string           `json:"errorMessage,omitempty"`
	ErrorMessageKey *string           `json:"errorMessageKey,omitempty"`
	Data            *CharacterDescDTO `json:"data,omitempty"`
	Success         *bool             `json:"success,omitempty"`
}

func NewResultDTOCharacterDescDTO() *ResultDTOCharacterDescDTO {
	this := ResultDTOCharacterDescDTO{}
	return &this
}

func (o *ResultDTOCharacterDescDTO) GetRequestId() *string {
	return o.RequestId
}

func (o *ResultDTOCharacterDescDTO) SetRequestId(v *string) {
	o.RequestId = v
}

func (o *ResultDTOCharacterDescDTO) GetCode() *int32 {
	return o.Code
}

func (o *ResultDTOCharacterDescDTO) SetCode(v *int32) {
	o.Code = v
}

func (o *ResultDTOCharacterDescDTO) GetErrorMessage() *string {
	return o.ErrorMessage
}

func (o *ResultDTOCharacterDescDTO) SetErrorMessage(v *string) {
	o.ErrorMessage = v
}

func (o *ResultDTOCharacterDescDTO) GetErrorMessageKey() *string {
	return o.ErrorMessageKey
}

func (o *ResultDTOCharacterDescDTO) SetErrorMessageKey(v *string) {
	o.ErrorMessageKey = v
}

func (o *ResultDTOCharacterDescDTO) GetData() *CharacterDescDTO {
	return o.Data
}

func (o *ResultDTOCharacterDescDTO) SetData(v *CharacterDescDTO) {
	o.Data = v
}

func (o *ResultDTOCharacterDescDTO) GetSuccess() *bool {
	return o.Success
}

func (o *ResultDTOCharacterDescDTO) SetSuccess(v *bool) {
	o.Success = v
}
