package xingchen

type KnowledgeBaseDeleteDTO struct {
	KnowledgeBaseId *string      `json:"knowledgeBaseId,omitempty"`
	UserProfile     *UserProfile `json:"userProfile,omitempty"`
}

func (o *KnowledgeBaseDeleteDTO) GetKnowledgeBaseId() *string {
	return o.KnowledgeBaseId
}

func (o *KnowledgeBaseDeleteDTO) SetKnowledgeBaseId(v *string) {
	o.KnowledgeBaseId = v
}

func (o *KnowledgeBaseDeleteDTO) GetUserProfile() *UserProfile {
	return o.UserProfile
}

func (o *KnowledgeBaseDeleteDTO) SetUserProfile(v *UserProfile) {
	o.UserProfile = v
}
