package xingchen

type KnowledgeBaseQueryDTO struct {
	PageNum     *int32       `json:"PageNum,omitempty"`
	PageSize    *int32       `json:"PageSize,omitempty"`
	UserProfile *UserProfile `json:"userProfile,omitempty"`
}

func (o *KnowledgeBaseQueryDTO) GetPageNum() *int32 {
	return o.PageNum
}

func (o *KnowledgeBaseQueryDTO) SetPageNum(v *int32) {
	o.PageNum = v
}

func (o *KnowledgeBaseQueryDTO) GetPageSize() *int32 {
	return o.PageSize
}

func (o *KnowledgeBaseQueryDTO) SetPageSize(v *int32) {
	o.PageSize = v
}

func (o *KnowledgeBaseQueryDTO) GetUserProfile() *UserProfile {
	return o.UserProfile
}

func (o *KnowledgeBaseQueryDTO) SetUserProfile(v *UserProfile) {
	o.UserProfile = v
}
