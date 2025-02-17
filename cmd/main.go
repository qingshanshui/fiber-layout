package main

import (
	"fmt"
	"os"
	"os/signal"
	"syscall"

	"github.com/gofiber/fiber/v3"
	"wat.ink/layout/fiber/internal/router"
	"wat.ink/layout/fiber/pkg/config"
	"wat.ink/layout/fiber/pkg/database"
	"wat.ink/layout/fiber/pkg/email"
	"wat.ink/layout/fiber/pkg/logger"
	"wat.ink/layout/fiber/pkg/rabbitmq"
	"wat.ink/layout/fiber/pkg/redis"
)

func main() {
	// 初始化配置
	config.Init()

	// 初始化邮件配置
	email.Init()

	// 初始化数据库连接
	database.Init()

	// 执行数据库迁移
	database.AutoMigrate()

	// 填充初始数据（仅在开发环境）
	if config.Conf.Debug {
		database.Seed()
	}

	// 初始化 Redis
	redis.Init()
	defer redis.Close()

	// 初始化 RabbitMQ
	rabbitmq.Init()
	defer rabbitmq.Close()

	// 创建 Fiber 应用
	app := fiber.New(fiber.Config{
		AppName:      config.Conf.App.Name,
		ErrorHandler: customErrorHandler,
	})

	// 设置路由
	router.SetupRoutes(app)

	// 启动服务器
	go func() {
		addr := fmt.Sprintf(":%d", config.Conf.App.Port)
		if err := app.Listen(addr); err != nil {
			logger.Fatal("Failed to start server", "error", err)
		}
	}()

	// 优雅关闭
	quit := make(chan os.Signal, 1)
	signal.Notify(quit, os.Interrupt, syscall.SIGTERM)
	<-quit

	logger.Info("Shutting down server...")
	if err := app.Shutdown(); err != nil {
		logger.Error("Server forced to shutdown", "error", err)
	}
}

func customErrorHandler(c fiber.Ctx, err error) error {
	// 记录错误日志
	logger.Error("Request error",
		"path", c.Path(),
		"method", c.Method(),
		"error", err,
	)

	// 默认错误响应
	code := fiber.StatusInternalServerError
	message := "Internal Server Error"

	// 根据错误类型设置不同的状态码和消息
	if e, ok := err.(*fiber.Error); ok {
		code = e.Code
		message = e.Message
	}

	// 返回 JSON 响应
	return c.Status(code).JSON(fiber.Map{
		"code":    code,
		"message": message,
	})
}
