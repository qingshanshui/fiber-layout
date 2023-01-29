package service

import (
	"fiber-layout/models"
	v1 "fiber-layout/validator/form"
	"fmt"
)

type Default struct {
}

func NewDefaultService() *Default {
	return &Default{}
}

func (t *Default) Register(loginForm v1.Register) (string, error) {
	msg := fmt.Sprintf("✋ ---- %s", loginForm.UserName)
	return msg, nil
}

func (t *Default) Login(loginForm v1.Login) ([]models.SysUser, error) {
	list, err := models.NewSysUser().GetList()
	if err != nil {
		return nil, err
	}
	return list, nil
}
