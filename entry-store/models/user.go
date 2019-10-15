package models

import (
	"github.com/jinzhu/gorm"
)

// User : struct
type User struct {
	gorm.Model
	FullName string
	UserName string
	Email    string
	Password string
}
