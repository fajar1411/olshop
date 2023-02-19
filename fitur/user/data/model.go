package data

import (
	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	Nama     string
	Password string
	Email    string
}
