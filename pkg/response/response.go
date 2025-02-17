package response

import "github.com/gofiber/fiber/v3"

type Response struct{}

func (r *Response) Success(c fiber.Ctx, data interface{}) error {
	return c.JSON(fiber.Map{
		"code": 200,
		"msg":  "success",
		"data": data,
	})
}

func (r *Response) Error(c fiber.Ctx, err interface{}) error {
	return c.JSON(fiber.Map{
		"code": 500,
		"msg":  "error",
		"data": err.(error).Error(),
	})
}
