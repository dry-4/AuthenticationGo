package models

import (
	"time"

	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	UserName  string    `json:"username" gorm:"unique; not null"`
	Name      string    `json:"name"`
	Email     string    `json:"email" gorm:"unique; not null"`
	Password  string    `json:"password" gorm:"not null"`
	IsActive  bool      `json:"is_active" gorm:"default:true"`
	LastLogin time.Time `json:"last_login" gorm:"default:NULL"`
}
