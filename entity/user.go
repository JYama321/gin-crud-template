package entity

import (
	"github.com/jinzhu/gorm"
)

type User struct {
	gorm.Model
	UserName string
	Email string
	PasswordDigest string
}