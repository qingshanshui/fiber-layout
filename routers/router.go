package routers

import (
	v1 "fiber-layout/controllers/v1"
	"github.com/gofiber/fiber/v2"
)

func SetRoute(app *fiber.App) {
	main := v1.NewDefaultController()
	group := app.Group("/v1")
	group.Get("/list", main.List)         // 列表
	group.Get("/category", main.Category) // 详情
}
