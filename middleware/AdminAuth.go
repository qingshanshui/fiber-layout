package middleware

import (
	"fiber-layout/pkg/utils"
	"github.com/gofiber/fiber/v2"
	"github.com/spf13/viper"
)

func AdminAuth(ctx *fiber.Ctx) error {
	cookie := ctx.Cookies("token")
	if cookie == "" {
		return ctx.JSON(fiber.Map{
			"code": 0,
			"data": "权限不足，禁止操作",
			"msg":  "操作失败",
		})
	} else {
		_, err := utils.ParseToken(cookie, viper.GetString("Jwt.Secret"))
		if err != nil {
			return ctx.JSON(fiber.Map{
				"code": 0,
				"data": "权限不足，禁止操作",
				"msg":  "操作失败",
			})
		} else {
			return ctx.Next()
		}
	}
}
