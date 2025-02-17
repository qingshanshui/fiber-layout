package errors

import "errors"

var (
	// ErrNotFound 资源未找到
	ErrNotFound = errors.New("resource not found")
	// ErrInvalidParams 无效的参数
	ErrInvalidParams = errors.New("invalid parameters")
	// ErrUnauthorized 未授权
	ErrUnauthorized = errors.New("unauthorized")
	// ErrForbidden 禁止访问
	ErrForbidden = errors.New("forbidden")
)

// New 创建新的错误
func New(text string) error {
	return errors.New(text)
}

// Wrap 包装错误并添加消息
func Wrap(err error, message string) error {
	if err == nil {
		return nil
	}
	return errors.New(message + ": " + err.Error())
}

// Is 判断错误是否为指定类型
func Is(err, target error) bool {
	return errors.Is(err, target)
}

// As 将错误转换为指定类型
func As(err error, target interface{}) bool {
	return errors.As(err, target)
}
