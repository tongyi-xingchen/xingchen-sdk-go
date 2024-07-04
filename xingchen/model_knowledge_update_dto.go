package xingchen

type KnowledgeBaseUpdateDTO struct {
	KnowledgeBaseId *string      `json:"knowledgeBaseId,omitempty"`
	Name            *string      `json:"name,omitempty"`
	Description     *string      `json:"description,omitempty"`
	UserProfile     *UserProfile `json:"userProfile,omitempty"`
}

func (o *KnowledgeBaseUpdateDTO) GetKnowledgeBaseId() *string {
	return o.KnowledgeBaseId
}

func (o *KnowledgeBaseUpdateDTO) SetKnowledgeBaseId(v *string) {
	o.KnowledgeBaseId = v
}

func (o *KnowledgeBaseUpdateDTO) GetName() *string {
	return o.Name
}

func (o *KnowledgeBaseUpdateDTO) SetName(v *string) {
	o.Name = v
}

func (o *KnowledgeBaseUpdateDTO) GetDescription() *string {
	return o.Description
}

func (o *KnowledgeBaseUpdateDTO) SetDescription(v *string) {
	o.Description = v
}

func (o *KnowledgeBaseUpdateDTO) GetUserProfile() *UserProfile {
	return o.UserProfile
}

func (o *KnowledgeBaseUpdateDTO) SetUserProfile(v *UserProfile) {
	o.UserProfile = v
}
