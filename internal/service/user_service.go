package service

import (
	"wat.ink/layout/fiber/api/v1/request"
	"wat.ink/layout/fiber/api/v1/response"
	"wat.ink/layout/fiber/internal/repository"
	"wat.ink/layout/fiber/pkg/utils"
	"wat.ink/layout/fiber/pkg/errors"
	"wat.ink/layout/fiber/pkg/logger"
)

type UserService struct {
	repo         *repository.UserRepository
	emailService *EmailService
}

func NewUserService() *UserService {
	return &UserService{
		repo:         repository.NewUserRepository(),
		emailService: NewEmailService(),
	}
}

func (s *UserService) List(req request.UserListRequest) (*response.UserListResponse, error) {
	offset, limit := utils.Paginate(req.Page, req.PageSize)
	
	users, total, err := s.repo.List(offset, limit, req.Keyword)
	if err != nil {
		return nil, err
	}

	items := make([]response.UserResponse, len(users))
	for i, user := range users {
		items[i] = response.UserResponse{
			ID:        user.ID,
			Username:  user.Username,
			Email:     user.Email,
			CreatedAt: user.CreatedAt,
			UpdatedAt: user.UpdatedAt,
		}
	}

	return &response.UserListResponse{
		Total: total,
		Items: items,
	}, nil
}

// Create 创建用户
func (s *UserService) Create(req request.UserCreateRequest) (*response.UserResponse, error) {
	salt := utils.GenerateSalt(6)
	hashedPassword := utils.GeneratePassword(req.Password, salt)

	user, err := s.repo.Create(repository.CreateUserParams{
		Username: req.Username,
		Email:    req.Email,
		Password: hashedPassword,
		Salt:     salt,
	})
	if err != nil {
		return nil, err
	}

	// 发送欢迎邮件
	err = s.emailService.SendEmail(
		req.Email,
		"欢迎加入",
		"感谢您注册我们的服务！",
	)
	if err != nil {
		logger.Warn("Failed to send welcome email", 
			"email", req.Email,
			"error", err,
		)
		// 不要因为发送邮件失败而影响用户注册
	}

	return &response.UserResponse{
		ID:        user.ID,
		Username:  user.Username,
		Email:     user.Email,
		CreatedAt: user.CreatedAt,
		UpdatedAt: user.UpdatedAt,
	}, nil
}

// Get 获取用户详情
func (s *UserService) Get(id uint) (*response.UserResponse, error) {
	user, err := s.repo.GetByID(id)
	if err != nil {
		return nil, err
	}
	if user == nil {
		return nil, errors.ErrNotFound
	}

	return &response.UserResponse{
		ID:        user.ID,
		Username:  user.Username,
		Email:     user.Email,
		CreatedAt: user.CreatedAt,
		UpdatedAt: user.UpdatedAt,
	}, nil
}

// Update 更新用户
func (s *UserService) Update(id uint, req request.UserUpdateRequest) (*response.UserResponse, error) {
	params := repository.UpdateUserParams{}
	if req.Username != "" {
		params.Username = &req.Username
	}
	if req.Email != "" {
		params.Email = &req.Email
	}
	if req.Password != "" {
		salt := utils.GenerateSalt(6)
		hashedPassword := utils.GeneratePassword(req.Password, salt)
		params.Password = &hashedPassword
		params.Salt = &salt
	}

	user, err := s.repo.Update(id, params)
	if err != nil {
		return nil, err
	}
	if user == nil {
		return nil, errors.ErrNotFound
	}

	return &response.UserResponse{
		ID:        user.ID,
		Username:  user.Username,
		Email:     user.Email,
		CreatedAt: user.CreatedAt,
		UpdatedAt: user.UpdatedAt,
	}, nil
}

// Delete 删除用户
func (s *UserService) Delete(id uint) error {
	return s.repo.Delete(id)
} 