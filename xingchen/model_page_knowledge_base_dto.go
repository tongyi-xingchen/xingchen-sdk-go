package xingchen

type PageResultKnowledgeBaseDTO struct {
	List            []KnowledgeBaseDTO `json:"list,omitempty"`
	Page            *int32             `json:"page,omitempty"`
	PageSize        *int32             `json:"PageSize,omitempty"`
	Total           *int32             `json:"total,omitempty"`
	PageOffsetValue *int32             `json:"pageOffsetValue,omitempty"`
}

func (o *PageResultKnowledgeBaseDTO) GetList() []KnowledgeBaseDTO {
	return o.List
}

func (o *PageResultKnowledgeBaseDTO) SetList(v []KnowledgeBaseDTO) {
	o.List = v
}

func (o *PageResultKnowledgeBaseDTO) GetPage() *int32 {
	return o.Page
}

func (o *PageResultKnowledgeBaseDTO) SetPage(v *int32) {
	o.Page = v
}

func (o *PageResultKnowledgeBaseDTO) GetPageSize() *int32 {
	return o.PageSize
}

func (o *PageResultKnowledgeBaseDTO) SetPageSize(v *int32) {
	o.PageSize = v
}

func (o *PageResultKnowledgeBaseDTO) GetTotal() *int32 {
	return o.Total
}

func (o *PageResultKnowledgeBaseDTO) SetTotal(v *int32) {
	o.Total = v
}

func (o *PageResultKnowledgeBaseDTO) GetPageOffsetValue() *int32 {
	return o.PageOffsetValue
}

func (o *PageResultKnowledgeBaseDTO) SetPageOffsetValue(v *int32) {
	o.PageOffsetValue = v
}
