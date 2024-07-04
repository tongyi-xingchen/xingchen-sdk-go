package xingchen

type KnowledgeBaseDetailQueryDTO struct {
	KnowledgeBaseId *string      `json:"knowledgeBaseId,omitempty"`
	PageNum         *int32       `json:"pageNum,omitempty"`
	PageSize        *int32       `json:"pageSize,omitempty"`
	UserProfile     *UserProfile `json:"userProfile,omitempty"`
}

func (o *KnowledgeBaseDetailQueryDTO) GetKnowledgeBaseId() *string {
	return o.KnowledgeBaseId
}

func (o *KnowledgeBaseDetailQueryDTO) SetKnowledgeBaseId(v *string) {
	o.KnowledgeBaseId = v
}

func (o *KnowledgeBaseDetailQueryDTO) GetPageNum() *int32 {
	return o.PageNum
}

func (o *KnowledgeBaseDetailQueryDTO) SetPageNum(v *int32) {
	o.PageNum = v
}

func (o *KnowledgeBaseDetailQueryDTO) GetPageSize() *int32 {
	return o.PageSize
}

func (o *KnowledgeBaseDetailQueryDTO) SetPageSize(v *int32) {
	o.PageSize = v
}

func (o *KnowledgeBaseDetailQueryDTO) GetUserProfile() *UserProfile {
	return o.UserProfile
}

func (o *KnowledgeBaseDetailQueryDTO) SetUserProfile(v *UserProfile) {
	o.UserProfile = v
}
