package data

import (
	"toko/fitur/user"

	"gorm.io/gorm"
)

type userData struct {
	db *gorm.DB
}

func NewUser(db *gorm.DB) user.UserData {
	return &userData{
		db: db,
	}
}

// Login implements user.UserData
func (*userData) Login(password string) (user.UserEntites, error) {
	panic("unimplemented")
}

// Register implements user.UserData
func (*userData) Register(newUser user.UserEntites) (user.UserEntites, error) {
	panic("unimplemented")
}
