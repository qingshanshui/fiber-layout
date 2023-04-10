package form

// CategoryRequest 详情
type CategoryRequest struct {
	Id string `form:"id" json:"id"  ` //  验证规则：必填，最小长度为5
}
