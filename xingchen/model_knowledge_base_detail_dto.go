package xingchen

type KnowledgeBaseDetailDTO struct {
	Id              *int32  `json:"id,omitempty"`
	GmtCreate       *int64  `json:"gmtCreate,omitempty"`
	GmtModified     *int64  `json:"gmtModified,omitempty"`
	Name            *string `json:"name,omitempty"`
	KnowledgeBaseId *string `json:"KnowledgeBaseId,omitempty"`
	Type            *string `json:"type,omitempty"`
	Status          *string `json:"status,omitempty"`
}

func (o *KnowledgeBaseDetailDTO) GetId() *int32 {
	return o.Id
}

func (o *KnowledgeBaseDetailDTO) SetId(v *int32) {
	o.Id = v
}

func (o *KnowledgeBaseDetailDTO) GetGmtCreate() *int64 {
	return o.GmtCreate
}

func (o *KnowledgeBaseDetailDTO) SetGmtCreate(v *int64) {
	o.GmtCreate = v
}

func (o *KnowledgeBaseDetailDTO) GetGmtModified() *int64 {
	return o.GmtModified
}

func (o *KnowledgeBaseDetailDTO) SetGmtModified(v *int64) {
	o.GmtModified = v
}

func (o *KnowledgeBaseDetailDTO) GetName() *string {
	return o.Name
}

func (o *KnowledgeBaseDetailDTO) SetName(v *string) {
	o.Name = v
}

func (o *KnowledgeBaseDetailDTO) GetKnowledgeBaseId() *string {
	return o.KnowledgeBaseId
}

func (o *KnowledgeBaseDetailDTO) SetKnowledgeBaseId(v *string) {
	o.KnowledgeBaseId = v
}

func (o *KnowledgeBaseDetailDTO) GetType() *string {
	return o.Type
}

func (o *KnowledgeBaseDetailDTO) SetType(v *string) {
	o.Type = v
}

func (o *KnowledgeBaseDetailDTO) GetStatus() *string {
	return o.Status
}

func (o *KnowledgeBaseDetailDTO) SetStatus(v *string) {
	o.Status = v
}
