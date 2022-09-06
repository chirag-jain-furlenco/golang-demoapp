package dbmodel

import (
	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	Name        string `json:"name"`
	Email       string `json:"email"`
	Description string `json:"description"`
}
