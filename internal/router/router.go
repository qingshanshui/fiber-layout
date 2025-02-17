package router

import (
	"github.com/gofiber/fiber/v3"
	v1 "wat.ink/layout/fiber/api/v1"
	"wat.ink/layout/fiber/internal/controller"
	"wat.ink/layout/fiber/internal/middleware"
)

func SetupRoutes(app *fiber.App) {
	// 全局中间件
	app.Use(middleware.LoggerMiddleware())

	// API v1
	v1Router := app.Group(v1.APIPrefix)

	// 用户路由
	userController := controller.NewUserController()
	users := v1Router.Group(v1.UserPrefix)
	{
		users.Get("/", userController.List) // 获取用户列表
		// users.Use(middleware.AuthMiddleware())      // 以下路由需要认证
		users.Post("/", userController.Create)      // 创建用户
		users.Get("/:id", userController.Get)       // 获取用户详情
		users.Put("/:id", userController.Update)    // 更新用户
		users.Delete("/:id", userController.Delete) // 删除用户
	}
}
