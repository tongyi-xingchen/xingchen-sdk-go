package xingchen

type TextToImageResultDTO struct {
	TaskId       *string  `json:"taskId,omitempty"`
	FileUrls     []string `json:"fileUrls,omitempty"`
	Status       *string  `json:"status,omitempty"`
	ErrorMessage *string  `json:"errorMessage,omitempty"`
}

func (o *TextToImageResultDTO) GetTaskId() *string {
	return o.TaskId
}

func (o *TextToImageResultDTO) SetTaskId(v *string) {
	o.TaskId = v
}

func (o *TextToImageResultDTO) GetFileUrls() []string {
	return o.FileUrls
}

func (o *TextToImageResultDTO) SetFileUrls(v []string) {
	o.FileUrls = v
}

func (o *TextToImageResultDTO) GetStatus() *string {
	return o.Status
}

func (o *TextToImageResultDTO) SetStatus(v *string) {
	o.Status = v
}

func (o *TextToImageResultDTO) GetErrorMessage() *string {
	return o.ErrorMessage
}

func (o *TextToImageResultDTO) SetErrorMessage(v *string) {
	o.ErrorMessage = v
}
