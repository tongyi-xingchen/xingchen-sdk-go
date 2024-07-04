package xingchen

type KnowledgeBaseDetailDeleteDTO struct {
	Name            *string      `json:"name,omitempty"`
	KnowledgeBaseId *string      `json:"knowledgeBaseId,omitempty"`
	UserProfile     *UserProfile `json:"userProfile,omitempty"`
}

func (o *KnowledgeBaseDetailDeleteDTO) GetName() *string {
	return o.Name
}

func (o *KnowledgeBaseDetailDeleteDTO) SetName(v *string) {
	o.Name = v
}

func (o *KnowledgeBaseDetailDeleteDTO) GetKnowledgeBaseId() *string {
	return o.KnowledgeBaseId
}

func (o *KnowledgeBaseDetailDeleteDTO) SetKnowledgeBaseId(v *string) {
	o.KnowledgeBaseId = v
}

func (o *KnowledgeBaseDetailDeleteDTO) GetUserProfile() *UserProfile {
	return o.UserProfile
}

func (o *KnowledgeBaseDetailDeleteDTO) SetUserProfile(v *UserProfile) {
	o.UserProfile = v
}
