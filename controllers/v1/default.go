package v1

import (
	"fiber-layout/controllers"
	"fiber-layout/initalize"
	"fiber-layout/pkg/utils"
	"fiber-layout/service"
	"fiber-layout/validator"
	"fiber-layout/validator/form"
	"github.com/gofiber/fiber/v2"
	"github.com/spf13/viper"
	"time"
)

type DefaultController struct {
	controllers.Base
}

func NewDefaultController() *DefaultController {
	return &DefaultController{}
}

// List 列表
func (t *DefaultController) List(c *fiber.Ctx) error {

	// 实际业务调用
	api, err := service.NewDefaultService().List()
	if err != nil {
		initalize.Log.Info(err)
		return c.Status(500).JSON(t.Fail(err))
	}
	// 设置cookie
	cookie := new(fiber.Cookie)
	cookie.Name = "token"
	token, err := utils.CreateToken("123456", viper.GetString("Jwt.Secret"))
	if err != nil {
		return err
	}
	cookie.Value = token
	cookie.Expires = time.Now().Add(24 * time.Hour)
	c.Cookie(cookie)
	return c.JSON(t.Ok(api))
}

// Category 详情
func (t *DefaultController) Category(c *fiber.Ctx) error {
	// 初始化参数结构体
	categoryForm := form.CategoryRequest{}
	// 绑定参数并使用验证器验证参数
	if err := validator.CheckQueryParams(c, &categoryForm); err != nil {
		initalize.Log.Info(err)
		return err
	}
	// 实际业务调用
	api, err := service.NewDefaultService().Category(categoryForm)
	if err != nil {
		initalize.Log.Info(err)
		return c.Status(500).JSON(t.Fail(err, 309))
	}
	return c.JSON(t.Ok(api))
}
