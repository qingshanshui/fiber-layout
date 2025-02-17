package request

// UserListRequest 用户列表请求参数
type UserListRequest struct {
	Page     int    `query:"page" validate:"min=1"`
	PageSize int    `query:"page_size" validate:"min=1,max=100"`
	Keyword  string `query:"keyword"`
}

// UserCreateRequest 创建用户请求参数
type UserCreateRequest struct {
	Username string `json:"username" validate:"required,min=3,max=32"`
	Email    string `json:"email" validate:"required,email"`
	Password string `json:"password" validate:"required,min=6"`
}

// UserUpdateRequest 更新用户请求参数
type UserUpdateRequest struct {
	Username string `json:"username" validate:"omitempty,min=3,max=32"`
	Email    string `json:"email" validate:"omitempty,email"`
	Password string `json:"password" validate:"omitempty,min=6"`
} 