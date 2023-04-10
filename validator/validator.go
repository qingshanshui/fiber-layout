package validator

import (
	"errors"
	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
	"reflect"
)

// CheckQueryParams GET验证
func CheckQueryParams(c *fiber.Ctx, obj interface{}) error {
	if err := c.QueryParser(obj); err != nil {
		return err
	}
	if err := validateStruct(obj); err != nil {
		return err
	}
	return nil
}

// CheckPostParams POST验证
func CheckPostParams(c *fiber.Ctx, obj interface{}) error {
	if err := c.BodyParser(obj); err != nil {
		return err
	}
	if err := validateStruct(obj); err != nil {
		return err
	}
	return nil
}

func validateStruct(obj interface{}) error {
	valid := validator.New()
	err := valid.Struct(obj)
	if err != nil {
		return errors.New(ProcessErr(obj, err))
	}
	return nil
}

func ProcessErr(u interface{}, err error) string {
	if err == nil { //如果为nil 说明校验通过
		return ""
	}
	invalid, ok := err.(*validator.InvalidValidationError) //如果是输入参数无效，则直接返回输入参数错误
	if ok {
		return "输入参数错误：" + invalid.Error()
	}
	validationErrs := err.(validator.ValidationErrors) //断言是ValidationErrors
	for _, validationErr := range validationErrs {
		fieldName := validationErr.Field() //获取是哪个字段不符合格式
		typeOf := reflect.TypeOf(u)
		// 如果是指针，获取其属性
		if typeOf.Kind() == reflect.Ptr {
			typeOf = typeOf.Elem()
		}
		field, ok := typeOf.FieldByName(fieldName) //通过反射获取filed
		if ok {
			errorInfo := field.Tag.Get("validate_error") // 获取field对应的validate_error tag值
			if errorInfo != "" {
				return fieldName + ":" + errorInfo // 返回错误
			} else {
				return "缺失validate_error"
			}
		} else {
			return "系统错误：获取不到反射体"
		}
	}
	return ""
}
