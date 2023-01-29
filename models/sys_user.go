package models

import (
	"fiber-layout/initalize"
	"gorm.io/gorm"
)

type SysUser struct {
	gorm.Model
	Uid       string `json:"uid"`
	Username  string `json:"username"`
	Nickname  string `json:"nickname"`
	Password  string `json:"password"`
	Telephone string `json:"telephone"`
	Email     string `json:"email"`
	HeadImg   string `json:"headImg"`
	Sex       string `json:"sex"`
	Age       string `json:"age"`
	Birthday  string `json:"birthday"`
	Status    int    `json:"status"`
}

func NewSysUser() *SysUser {
	return &SysUser{}
}

func (t *SysUser) GetList() ([]SysUser, error) {
	var sys []SysUser
	if err := initalize.DB.Raw("select * from sys_user").Find(&sys).Error; err != nil {
		return nil, err
	}
	return sys, nil
}
