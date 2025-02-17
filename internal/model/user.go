package model

import (
	"gorm.io/gorm"
	"time"
)

type User struct {
	ID        uint           `gorm:"primarykey" json:"id"`
	Username  string         `gorm:"size:32;uniqueIndex;not null" json:"username"`
	Email     string         `gorm:"size:128;uniqueIndex" json:"email"`
	Password  string         `gorm:"size:128;not null" json:"-"`
	Salt      string         `gorm:"size:6;not null" json:"-"`
	CreatedAt time.Time      `json:"created_at"`
	UpdatedAt time.Time      `json:"updated_at"`
	DeletedAt gorm.DeletedAt `gorm:"index" json:"-"`
} 