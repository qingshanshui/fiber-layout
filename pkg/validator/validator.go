package validator

import (
	"NextEraAbyss/FiberForge/pkg/errors"
	"NextEraAbyss/FiberForge/pkg/logger"
	"strings"

	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v3"
)

var validate = validator.New()

// ValidateStruct 验证结构体
func ValidateStruct(payload interface{}) error {
	err := validate.Struct(payload)
	if err != nil {
		// 处理验证错误
		if _, ok := err.(*validator.InvalidValidationError); ok {
			logger.Error("Invalid validation error", "error", err)
			return errors.ErrInvalidParams
		}

		validationErrors := err.(validator.ValidationErrors)
		messages := make([]string, 0)
		for _, e := range validationErrors {
			messages = append(messages, formatValidationError(e))
		}

		errMsg := strings.Join(messages, "; ")
		logger.Warn("Validation failed", "errors", errMsg)
		return errors.New(errMsg)
	}
	return nil
}

// ParseAndValidate 解析请求体并验证
func ParseAndValidate(c fiber.Ctx, payload interface{}) error {
	if err := c.Bind().JSON(payload); err != nil {
		logger.Warn("Failed to parse request body",
			"error", err,
			"path", c.Path(),
		)
		return errors.ErrInvalidParams
	}
	return ValidateStruct(payload)
}

// ParseQueryAndValidate 解析查询参数并验证
func ParseQueryAndValidate(c fiber.Ctx, payload interface{}) error {
	if err := c.Bind().Query(payload); err != nil {
		logger.Warn("Failed to parse query parameters",
			"error", err,
			"path", c.Path(),
		)
		return errors.ErrInvalidParams
	}
	return ValidateStruct(payload)
}

// formatValidationError 格式化验证错误信息
func formatValidationError(e validator.FieldError) string {
	switch e.Tag() {
	case "required":
		return e.Field() + "是必需的"
	case "email":
		return e.Field() + "必须是有效的电子邮件地址"
	case "min":
		return e.Field() + "长度不能小于" + e.Param()
	case "max":
		return e.Field() + "长度不能大于" + e.Param()
	default:
		return e.Field() + "验证失败：" + e.Tag()
	}
}
