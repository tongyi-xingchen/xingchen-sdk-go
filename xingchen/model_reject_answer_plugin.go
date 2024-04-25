package xingchen

type RejectAnswerPlugin struct {
	PlatformPlugin
	RejectConditions []RejectCondition `json:"rejectConditions,omitempty"`
}

func (o *RejectAnswerPlugin) GetRejectConditions() []RejectCondition {
	return o.RejectConditions
}

func (o *RejectAnswerPlugin) SetRejectConditions(v []RejectCondition) {
	o.RejectConditions = v
}

func NewRejectAnswerPlugin(rejectConditions []RejectCondition) *RejectAnswerPlugin {
	plugin := RejectAnswerPlugin{
		PlatformPlugin: PlatformPlugin{
			Name:    PtrString("reject_answer_plugin"),
			Enabled: PtrBool(true),
		},
		RejectConditions: rejectConditions,
	}
	return &plugin
}
