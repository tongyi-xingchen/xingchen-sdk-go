package xingchen

type KnowledgeBaseDetailUpdateDTO struct {
	KnowledgeBaseId *string      `json:"knowledgeBaseId,omitempty"`
	Name            *string      `json:"name,omitempty"`
	NewName         *string      `json:"newName,omitempty"`
	UserProfile     *UserProfile `json:"userProfile,omitempty"`
}

func (o *KnowledgeBaseDetailUpdateDTO) GetKnowledgeBaseId() *string {
	return o.KnowledgeBaseId
}

func (o *KnowledgeBaseDetailUpdateDTO) SetKnowledgeBaseId(v *string) {
	o.KnowledgeBaseId = v
}

func (o *KnowledgeBaseDetailUpdateDTO) GetName() *string {
	return o.Name
}

func (o *KnowledgeBaseDetailUpdateDTO) SetName(v *string) {
	o.Name = v
}

func (o *KnowledgeBaseDetailUpdateDTO) GetNewName() *string {
	return o.NewName
}

func (o *KnowledgeBaseDetailUpdateDTO) SetNewName(v *string) {
	o.NewName = v
}

func (o *KnowledgeBaseDetailUpdateDTO) GetUserProfile() *UserProfile {
	return o.UserProfile
}

func (o *KnowledgeBaseDetailUpdateDTO) SetUserProfile(v *UserProfile) {
	o.UserProfile = v
}
