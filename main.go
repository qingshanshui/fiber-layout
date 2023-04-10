package main

import (
	_ "fiber-layout/config"
	"fiber-layout/initalize"
	"fiber-layout/models"
	"fiber-layout/routers"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/gofiber/fiber/v2/middleware/recover"
	"github.com/spf13/viper"
)

func main() {
	app := fiber.New()
	// 静态目录
	app.Static("/", "./static")
	// HTTP 请求/响应日志
	app.Use(logger.New())
	// 使用各种选项启用跨源资源共享(CORS)
	app.Use(cors.New())
	// Recover 中间件将可以堆栈链中的任何位置将 panic 恢复，并将处理集中到
	app.Use(recover.New())
	// 设置路由
	routers.SetRoute(app)
	// 初始化 数据表
	err := initalize.DB.Set("gorm:table_options", "ENGINE=InnoDB").AutoMigrate(&models.Course{})
	if err != nil {
		initalize.Log.Error("数据库表生成失败：", err)
		panic("数据库表，生成失败")
	}
	// 监听端口
	err = app.Listen(viper.GetString("App.Port"))
	if err != nil {
		initalize.Log.Error("项目启动失败：", err)
		panic("项目启动失败")
	}
}
