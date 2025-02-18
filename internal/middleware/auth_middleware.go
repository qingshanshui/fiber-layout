package middleware

import (
	"NextEraAbyss/FiberForge/pkg/errors"
	"NextEraAbyss/FiberForge/pkg/jwt"
	"NextEraAbyss/FiberForge/pkg/logger"

	"github.com/gofiber/fiber/v3"
)

// AuthMiddleware 认证中间件
func AuthMiddleware() fiber.Handler {
	return func(c fiber.Ctx) error {
		// 提取公共日志字段
		logFields := []interface{}{
			"path", c.Path(),
		}

		token := c.Get("Authorization")
		if token == "" {
			logger.Warn("Missing authorization token", logFields...)
			return errors.ErrUnauthorized
		}

		claims, err := jwt.VerifyToken(token)
		if err != nil {
			logger.Warn("Invalid token", append(logFields,
				"error", err,
				"token", token,
			)...)
			return errors.ErrUnauthorized
		}

		// 将用户信息存储到上下文中
		c.Locals("user_id", claims.UserID)
		logger.Debug("User authenticated", append(logFields,
			"user_id", claims.UserID,
		)...)

		return c.Next()
	}
}
