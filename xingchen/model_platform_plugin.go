package xingchen

type PlatformPlugin struct {
	Enabled *bool   `json:"enabled,omitempty"`
	Name    *string `json:"name,omitempty"`
}

func (o *PlatformPlugin) GetEnabled() *bool {
	return o.Enabled
}

func (o *PlatformPlugin) SetEnabled(v *bool) {
	o.Enabled = v
}

func (o *PlatformPlugin) GetName() *string {
	return o.Name
}

func (o *PlatformPlugin) SetName(v *string) {
	o.Name = v
}
