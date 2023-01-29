package controllers

import "github.com/gofiber/fiber/v2"

// 返回的 code 码值
const (
	ERROR   = 0
	SUCCESS = 1000
)

// Base 公共方法
type Base struct{}

// Ok 成功
func (r *Base) Ok(data interface{}) fiber.Map {
	return fiber.Map{
		"code": SUCCESS,
		"data": data,
		"msg":  "操作成功",
	}
}

// Fail 失败
func (r *Base) Fail(data interface{}, code ...int) fiber.Map {
	return fiber.Map{
		"code": If(code),
		"data": data,
		"msg":  "操作失败",
	}
}

func If(code []int) int {
	if code != nil {
		return code[0]
	}
	return ERROR
}
