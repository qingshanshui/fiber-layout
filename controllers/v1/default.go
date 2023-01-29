package v1

import (
	"fiber-layout/controllers"
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

func (t *DefaultController) Register(c *fiber.Ctx) error {
	// 初始化参数结构体
	loginForm := form.Register{}
	// 绑定参数并使用验证器验证参数
	if err := validator.CheckQueryParams(c, &loginForm); err != nil {
		return err
	}
	// 实际业务调用
	api, err := service.NewDefaultService().Register(loginForm)
	if err != nil {
		return c.JSON(t.Fail(err))
	}
	return c.JSON(t.Ok(api)) // => ✋ register
}

func (t *DefaultController) Login(c *fiber.Ctx) error {
	// 初始化参数结构体
	loginForm := form.Login{}
	// 绑定参数并使用验证器验证参数
	if err := validator.CheckPostParams(c, &loginForm); err != nil {
		return err
	}
	// 实际业务调用
	api, err := service.NewDefaultService().Login(loginForm)
	if err != nil {
		return c.JSON(t.Fail(err, 309))
	}
	return c.JSON(t.Ok(api)) // => ✋ Login
}
