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

// Profile implements user.UserData

func NewUser(db *gorm.DB) user.UserData {
	return &userData{
		db: db,
	}
}

// Login implements user.UserData
func (ud *userData) Login(email string) (user.UserEntites, error) {
	res := User{}

	if err := ud.db.Where("email = ?", email).First(&res).Error; err != nil {
		log.Println("login query error", err.Error())
		return user.UserEntites{}, errors.New("data not found")
	}

	return ToCore(res), nil
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
func (Ud *userData) Profile(id int) (user.UserEntites, error) {
	var users User

	if err := Ud.db.Where("id = ?", id).First(&users).Error; err != nil {
		log.Println("Get By ID query error", err.Error())
		return user.UserEntites{}, err
	}
	gorms := users.ModelsToCore()
	return gorms, nil
}
