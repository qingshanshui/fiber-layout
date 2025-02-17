package middleware

import (
	"github.com/gofiber/fiber/v3"
	"time"
	"wat.ink/layout/fiber/pkg/logger"
)

// LoggerMiddleware 日志中间件
func LoggerMiddleware() fiber.Handler {
	return func(c fiber.Ctx) error {
		start := time.Now()
		path := c.Path()
		method := c.Method()

		err := c.Next()

		duration := time.Since(start)
		statusCode := c.Response().StatusCode()
		ip := c.IP()

		// 使用结构化日志
		logger.Info("HTTP Request",
			"method", method,
			"path", path,
			"status", statusCode,
			"duration", duration,
			"ip", ip,
		)

		return err
	}
} 