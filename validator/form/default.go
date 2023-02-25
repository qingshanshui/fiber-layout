package form

// CategoryRequest 详情
type CategoryRequest struct {
	Id string `form:"id" json:"id"  validate:"required,min=5" validate_error:"id不能为空"` //  验证规则：必填，最小长度为5
}
