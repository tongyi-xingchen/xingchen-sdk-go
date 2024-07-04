package xingchen

type KnowledgeBaseCreateDTO struct {
	Name        *string      `json:"name,omitempty"`
	Description *string      `json:"description,omitempty"`
	UserProfile *UserProfile `json:"userProfile,omitempty"`
}

func (o *KnowledgeBaseCreateDTO) GetName() *string {
	return o.Name
}

func (o *KnowledgeBaseCreateDTO) SetName(v *string) {
	o.Name = v
}

func (o *KnowledgeBaseCreateDTO) GetDescription() *string {
	return o.Description
}

func (o *KnowledgeBaseCreateDTO) SetDescription(v *string) {
	o.Description = v
}

func (o *KnowledgeBaseCreateDTO) GetUserProfile() *UserProfile {
	return o.UserProfile
}

func (o *KnowledgeBaseCreateDTO) SetUserProfile(v *UserProfile) {
	o.UserProfile = v
}
