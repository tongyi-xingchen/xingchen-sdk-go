package xingchen

type KnowledgeBaseDetailUploadDTO struct {
	KnowledgeBaseId *string      `json:"knowledgeBaseId,omitempty"`
	Type            *string      `json:"type,omitempty"`
	FileInfos       []FileInfoVO `json:"fileInfos,omitempty"`
	UserProfile     *UserProfile `json:"userProfile,omitempty"`
}

func (o *KnowledgeBaseDetailUploadDTO) GetKnowledgeBaseId() *string {
	return o.KnowledgeBaseId
}

func (o *KnowledgeBaseDetailUploadDTO) SetKnowledgeBaseId(v *string) {
	o.KnowledgeBaseId = v
}

func (o *KnowledgeBaseDetailUploadDTO) GetType() *string {
	return o.Type
}

func (o *KnowledgeBaseDetailUploadDTO) SetType(v *string) {
	o.Type = v
}

func (o *KnowledgeBaseDetailUploadDTO) GetFileInfos() []FileInfoVO {
	return o.FileInfos
}

func (o *KnowledgeBaseDetailUploadDTO) SetFileInfos(v []FileInfoVO) {
	o.FileInfos = v
}

func (o *KnowledgeBaseDetailUploadDTO) GetUserProfile() *UserProfile {
	return o.UserProfile
}

func (o *KnowledgeBaseDetailUploadDTO) SetUserProfile(v *UserProfile) {
	o.UserProfile = v
}
