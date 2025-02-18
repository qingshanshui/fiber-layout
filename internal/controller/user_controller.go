package controller

import (
	"strconv"

	"NextEraAbyss/FiberForge/api/v1/request"
	"NextEraAbyss/FiberForge/internal/service"
	"NextEraAbyss/FiberForge/pkg/errors"
	"NextEraAbyss/FiberForge/pkg/response"
	"NextEraAbyss/FiberForge/pkg/validator"

	"github.com/gofiber/fiber/v3"
)

type UserController struct {
	userService *service.UserService
	response    *response.Response
}

func NewUserController() *UserController {
	return &UserController{
		userService: service.NewUserService(),
		response:    &response.Response{},
	}
}

// List 获取用户列表
func (c *UserController) List(ctx fiber.Ctx) error {
	var req request.UserListRequest
	if err := validator.ParseQueryAndValidate(ctx, &req); err != nil {
		return c.response.Error(ctx, err)
	}

	result, err := c.userService.List(req)
	if err != nil {
		return c.response.Error(ctx, err)
	}
	return c.response.Success(ctx, result)
}

// Create 创建用户
func (c *UserController) Create(ctx fiber.Ctx) error {
	var req request.UserCreateRequest
	if err := validator.ParseAndValidate(ctx, &req); err != nil {
		return c.response.Error(ctx, err)
	}

	result, err := c.userService.Create(req)
	if err != nil {
		return c.response.Error(ctx, err)
	}
	return c.response.Success(ctx, result)
}

// Get 获取用户详情
func (c *UserController) Get(ctx fiber.Ctx) error {
	idStr := ctx.Params("id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		return c.response.Error(ctx, errors.ErrInvalidParams)
	}

	result, err := c.userService.Get(uint(id))
	if err != nil {
		return c.response.Error(ctx, err)
	}
	return c.response.Success(ctx, result)
}

// Update 更新用户
func (c *UserController) Update(ctx fiber.Ctx) error {
	idStr := ctx.Params("id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		return c.response.Error(ctx, errors.ErrInvalidParams)
	}

	var req request.UserUpdateRequest
	if err := validator.ParseAndValidate(ctx, &req); err != nil {
		return c.response.Error(ctx, err)
	}

	result, err := c.userService.Update(uint(id), req)
	if err != nil {
		return c.response.Error(ctx, err)
	}
	return c.response.Success(ctx, result)
}

// Delete 删除用户
func (c *UserController) Delete(ctx fiber.Ctx) error {
	idStr := ctx.Params("id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		return c.response.Error(ctx, errors.ErrInvalidParams)
	}

	if err := c.userService.Delete(uint(id)); err != nil {
		return c.response.Error(ctx, err)
	}
	return c.response.Success(ctx, nil)
}
