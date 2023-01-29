package form

// Login 登录
type Login struct {
	UserName string `form:"username" json:"username"  validate:"required,min=5"` //  验证规则：必填，最小长度为5
	PassWord string `form:"password" json:"password"  validate:"required,min=6"` //  验证规则：必填，最小长度为6
}

// Register 注册
type Register struct {
	UserName string `form:"username" json:"username"  validate:"required,min=5"` //  验证规则：必填，最小长度为5
	PassWord string `form:"password" json:"password"  validate:"required,min=6"` //  验证规则：必填，最小长度为6
}
