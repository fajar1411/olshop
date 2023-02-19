package data

import (
	"toko/fitur/user"

	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	Nama     string
	Password string
	Email    string
}

func FromUserCore(dataCore user.UserEntites) User { //fungsi yang mengambil data dari entities usercore dan merubah data ke user gorm(model.go)
	userGorm := User{

		Email:    dataCore.Email,
		Password: dataCore.Password,
		Nama:     dataCore.Nama,
	}
	return userGorm
}
func ToCore(data User) user.UserEntites {
	return user.UserEntites{
		ID:       data.ID,
		Nama:     data.Nama,
		Email:    data.Email,
		Password: data.Password,
	}
}
