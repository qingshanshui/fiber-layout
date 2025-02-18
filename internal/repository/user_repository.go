package repository

import (
	"NextEraAbyss/FiberForge/internal/model"
	"NextEraAbyss/FiberForge/pkg/database"
)

type UserRepository struct{}

func NewUserRepository() *UserRepository {
	return &UserRepository{}
}

func (r *UserRepository) List(offset, limit int, keyword string) ([]model.User, int64, error) {
	var users []model.User
	var total int64

	query := database.DB.Model(&model.User{})

	if keyword != "" {
		query = query.Where("username LIKE ? OR email LIKE ?", "%"+keyword+"%", "%"+keyword+"%")
	}

	if err := query.Count(&total).Error; err != nil {
		return nil, 0, err
	}

	if err := query.Offset(offset).Limit(limit).Find(&users).Error; err != nil {
		return nil, 0, err
	}

	return users, total, nil
}

// CreateUserParams 创建用户参数
type CreateUserParams struct {
	Username string
	Email    string
	Password string
	Salt     string
}

// UpdateUserParams 更新用户参数
type UpdateUserParams struct {
	Username *string
	Email    *string
	Password *string
	Salt     *string
}

// Create 创建用户
func (r *UserRepository) Create(params CreateUserParams) (*model.User, error) {
	user := &model.User{
		Username: params.Username,
		Email:    params.Email,
		Password: params.Password,
		Salt:     params.Salt,
	}

	if err := database.DB.Create(user).Error; err != nil {
		return nil, err
	}
	return user, nil
}

// GetByID 通过ID获取用户
func (r *UserRepository) GetByID(id uint) (*model.User, error) {
	var user model.User
	if err := database.DB.First(&user, id).Error; err != nil {
		if err == database.ErrRecordNotFound {
			return nil, nil
		}
		return nil, err
	}
	return &user, nil
}

// Update 更新用户
func (r *UserRepository) Update(id uint, params UpdateUserParams) (*model.User, error) {
	user, err := r.GetByID(id)
	if err != nil {
		return nil, err
	}
	if user == nil {
		return nil, nil
	}

	updates := make(map[string]interface{})
	if params.Username != nil {
		updates["username"] = *params.Username
	}
	if params.Email != nil {
		updates["email"] = *params.Email
	}
	if params.Password != nil {
		updates["password"] = *params.Password
	}
	if params.Salt != nil {
		updates["salt"] = *params.Salt
	}

	if err := database.DB.Model(user).Updates(updates).Error; err != nil {
		return nil, err
	}
	return user, nil
}

// Delete 删除用户
func (r *UserRepository) Delete(id uint) error {
	return database.DB.Delete(&model.User{}, id).Error
}
