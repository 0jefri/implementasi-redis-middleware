package model

import "gorm.io/gorm"

type User struct {
	gorm.Model
	Username string `gorm:"type:varchar(255);not null" json:"username" binding:"required,alphanum"`
	Password string `gorm:"type:varchar(255);not null" json:"password,omitempty" binding:"required,alphanum,min=6"`
}
