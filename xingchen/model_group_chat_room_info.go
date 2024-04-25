package xingchen

type GroupChatRoomInfo struct {
	Name        *string `json:"name,omitempty"`
	Description *string `json:"description,omitempty"`
}

func (o *GroupChatRoomInfo) GetName() *string {
	return o.Name
}

func (o *GroupChatRoomInfo) SetName(v *string) {
	o.Name = v
}

func (o *GroupChatRoomInfo) GetDescription() *string {
	return o.Description
}

func (o *GroupChatRoomInfo) SetDescription(v *string) {
	o.Description = v
}
