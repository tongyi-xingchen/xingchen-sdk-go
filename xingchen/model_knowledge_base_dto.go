package xingchen

type KnowledgeBaseDTO struct {
	Id              *int32  `json:"id,omitempty"`
	GmtCreate       *int64  `json:"gmtCreate,omitempty"`
	GmtModified     *int64  `json:"gmtModified,omitempty"`
	Name            *string `json:"name,omitempty"`
	KnowledgeBaseId *string `json:"KnowledgeBaseId,omitempty"`
	Description     *string `json:"description,omitempty"`
	UserId          *string `json:"userId,omitempty"`
	BizUserId       *string `json:"BizUserId,omitempty"`
	Count           *int32  `json:"count,omitempty"`
}

func (o *KnowledgeBaseDTO) GetId() *int32 {
	return o.Id
}

func (o *KnowledgeBaseDTO) SetId(v *int32) {
	o.Id = v
}

func (o *KnowledgeBaseDTO) GetGmtCreate() *int64 {
	return o.GmtCreate
}

func (o *KnowledgeBaseDTO) SetGmtCreate(v *int64) {
	o.GmtCreate = v
}

func (o *KnowledgeBaseDTO) GetGmtModified() *int64 {
	return o.GmtModified
}

func (o *KnowledgeBaseDTO) SetGmtModified(v *int64) {
	o.GmtModified = v
}

func (o *KnowledgeBaseDTO) GetName() *string {
	return o.Name
}

func (o *KnowledgeBaseDTO) SetName(v *string) {
	o.Name = v
}

func (o *KnowledgeBaseDTO) GetKnowledgeBaseId() *string {
	return o.KnowledgeBaseId
}

func (o *KnowledgeBaseDTO) SetKnowledgeBaseId(v *string) {
	o.KnowledgeBaseId = v
}

func (o *KnowledgeBaseDTO) GetDescription() *string {
	return o.Description
}

func (o *KnowledgeBaseDTO) SetDescription(v *string) {
	o.Description = v
}

func (o *KnowledgeBaseDTO) GetUserId() *string {
	return o.UserId
}

func (o *KnowledgeBaseDTO) SetUserId(v *string) {
	o.UserId = v
}

func (o *KnowledgeBaseDTO) GetBizUserId() *string {
	return o.BizUserId
}

func (o *KnowledgeBaseDTO) SetBizUserId(v *string) {
	o.BizUserId = v
}

func (o *KnowledgeBaseDTO) GetCount() *int32 {
	return o.Count
}

func (o *KnowledgeBaseDTO) SetCount(v *int32) {
	o.Count = v
}
