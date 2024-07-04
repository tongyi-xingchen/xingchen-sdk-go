package xingchen

type PageResultKnowledgeBaseDetailDTO struct {
	List            []KnowledgeBaseDetailDTO `json:"list,omitempty"`
	Page            *int32                   `json:"page,omitempty"`
	PageSize        *int32                   `json:"PageSize,omitempty"`
	Total           *int32                   `json:"total,omitempty"`
	PageOffsetValue *int32                   `json:"pageOffsetValue,omitempty"`
}

func (o *PageResultKnowledgeBaseDetailDTO) GetList() []KnowledgeBaseDetailDTO {
	return o.List
}

func (o *PageResultKnowledgeBaseDetailDTO) SetList(v []KnowledgeBaseDetailDTO) {
	o.List = v
}

func (o *PageResultKnowledgeBaseDetailDTO) GetPage() *int32 {
	return o.Page
}

func (o *PageResultKnowledgeBaseDetailDTO) SetPage(v *int32) {
	o.Page = v
}

func (o *PageResultKnowledgeBaseDetailDTO) GetPageSize() *int32 {
	return o.PageSize
}

func (o *PageResultKnowledgeBaseDetailDTO) SetPageSize(v *int32) {
	o.PageSize = v
}

func (o *PageResultKnowledgeBaseDetailDTO) GetTotal() *int32 {
	return o.Total
}

func (o *PageResultKnowledgeBaseDetailDTO) SetTotal(v *int32) {
	o.Total = v
}

func (o *PageResultKnowledgeBaseDetailDTO) GetPageOffsetValue() *int32 {
	return o.PageOffsetValue
}

func (o *PageResultKnowledgeBaseDetailDTO) SetPageOffsetValue(v *int32) {
	o.PageOffsetValue = v
}
