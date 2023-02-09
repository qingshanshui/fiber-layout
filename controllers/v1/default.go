package v1

import (
	"fiber-layout/controllers"
	"fiber-layout/initalize"
	"fiber-layout/service"
	"fiber-layout/validator"
	"fiber-layout/validator/form"

	"github.com/gofiber/fiber/v2"
)

type DefaultController struct {
	controllers.Base
}

func NewDefaultController() *DefaultController {
	return &DefaultController{}
}

func (t *DefaultController) GetList(c *fiber.Ctx) error {
	// 实际业务调用
	api, err := service.NewDefaultService().GetList()
	if err != nil {
		return c.JSON(t.Fail(err))
	}
	initalize.Log.Info(api)
	return c.JSON(t.Ok(api)) // => ✋ register
}

func (t *DefaultController) Category(c *fiber.Ctx) error {
	// 初始化参数结构体
	categoryForm := form.Category{}
	// 绑定参数并使用验证器验证参数
	if err := validator.CheckPostParams(c, &categoryForm); err != nil {
		return err
	}
	// 实际业务调用
	api, err := service.NewDefaultService().Category(categoryForm)
	if err != nil {
		return c.JSON(t.Fail(err, 309))
	}
	return c.JSON(t.Ok(api)) // => ✋ Login
}
