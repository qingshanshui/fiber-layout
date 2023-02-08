package form

// Category 详情
type Category struct {
	ID string `form:"id" json:"id"  validate:"required"` //  验证规则：必填，最小长度为5
}
