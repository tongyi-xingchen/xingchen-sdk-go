package xingchen

type GroupChatExtParams struct {
	BotProfiles  []BotProfile           `json:"botProfiles,omitempty"`
	ReplySetting *GroupChatReplySetting `json:"replySetting,omitempty"`
	UserProfile  *UserProfile           `json:"userProfile,omitempty"`
	GroupInfo    *GroupChatRoomInfo     `json:"groupInfo,omitempty"`
}

func (o *GroupChatExtParams) GetBotProfiles() []BotProfile {
	return o.BotProfiles
}

func (o *GroupChatExtParams) SetBotProfiles(v []BotProfile) {
	o.BotProfiles = v
}

func (o *GroupChatExtParams) GetReplySetting() *GroupChatReplySetting {
	return o.ReplySetting
}

func (o *GroupChatExtParams) SetReplySetting(v *GroupChatReplySetting) {
	o.ReplySetting = v
}

func (o *GroupChatExtParams) GetUserProfile() *UserProfile {
	return o.UserProfile
}

func (o *GroupChatExtParams) SetUserProfile(v *UserProfile) {
	o.UserProfile = v
}

func (o *GroupChatExtParams) GetGroupInfo() *GroupChatRoomInfo {
	return o.GroupInfo
}

func (o *GroupChatExtParams) SetGroupInfo(v *GroupChatRoomInfo) {
	o.GroupInfo = v
}
