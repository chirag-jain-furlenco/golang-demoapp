package dbmodel

import "gorm.io/gorm"

type Task struct {
	gorm.Model
	UserId      uint
	Name        string `gorm:"index;default: ''" json:"name"`
	Description string `gorm:"index;default: ''" json:"description"`
}
