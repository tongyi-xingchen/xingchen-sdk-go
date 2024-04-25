package xingchen

type TextToImagePlugin struct {
	PlatformPlugin
	ImageStyle   *string `json:"imageStyle,omitempty"`
	PositiveDesc *string `json:"positiveDesc,omitempty"`
	NegativeDesc *string `json:"negativeDesc,omitempty"`
}

func (o *TextToImagePlugin) GetImageStyle() *string {
	return o.ImageStyle
}

func (o *TextToImagePlugin) SetImageStyle(v *string) {
	o.ImageStyle = v
}

func (o *TextToImagePlugin) GetPositiveDesc() *string {
	return o.PositiveDesc
}

func (o *TextToImagePlugin) SetPositiveDesc(v *string) {
	o.PositiveDesc = v
}

func (o *TextToImagePlugin) GetNegativeDesc() *string {
	return o.NegativeDesc
}

func (o *TextToImagePlugin) SetNegativeDesc(v *string) {
	o.NegativeDesc = v
}

func NewTextToImagePlugin(imageStyle *string, positiveDesc *string, negativeDesc *string) *TextToImagePlugin {
	plugin := TextToImagePlugin{
		PlatformPlugin: PlatformPlugin{
			Name:    PtrString("text_to_image_plugin"),
			Enabled: PtrBool(true),
		},
		ImageStyle:   imageStyle,
		PositiveDesc: positiveDesc,
		NegativeDesc: negativeDesc,
	}
	return &plugin
}
