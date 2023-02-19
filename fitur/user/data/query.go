package data

import (
	"errors"
	"log"
	"strings"
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
func (Ud *userData) Register(newUser user.UserEntites) (user.UserEntites, error) {
	userGorm := FromUserCore(newUser)

	tx := Ud.db.Create(&userGorm) // proses insert data

	if tx.Error != nil {
		log.Println("register query error", tx.Error.Error())
		msg := ""
		if strings.Contains(tx.Error.Error(), "Duplicate") {
			msg = "data is duplicated"
		} else {
			msg = "server error"
		}
		return user.UserEntites{}, errors.New(msg)
	}
	return newUser, nil
}
