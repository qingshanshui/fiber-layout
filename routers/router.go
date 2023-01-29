package routers

import (
	v1 "fiber-layout/controllers/v1"
	"github.com/gofiber/fiber/v2"
)

func SetRoute(app *fiber.App) {

	main := v1.NewDefaultController()
	group := app.Group("/v1")
	// GET /register 	get
	group.Get("/register", main.Register)
	// GET /login 	json
	group.Post("/login", main.Login)

}
