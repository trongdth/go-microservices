package models

import (
	"github.com/jinzhu/gorm"
)

// User : struct
type User struct {
	gorm.Model
	FullName string
	Email    string
	Password string
}
