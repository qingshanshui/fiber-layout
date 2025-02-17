package response

import "time"

// UserResponse 用户响应结构
type UserResponse struct {
	ID        uint      `json:"id"`
	Username  string    `json:"username"`
	Email     string    `json:"email"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

// UserListResponse 用户列表响应结构
type UserListResponse struct {
	Total int64          `json:"total"`
	Items []UserResponse `json:"items"`
} 