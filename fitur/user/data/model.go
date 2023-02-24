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

func FromEntities(dataCore user.UserEntites) User { //fungsi yang mengambil data dari entities usercore dan merubah data ke user gorm(model.go)
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
func (dataModel *User) ModelsToCore() user.UserEntites { //fungsi yang mengambil data dari  user gorm(model.go)  dan merubah data ke entities usercore
	return user.UserEntites{
		Nama:     dataModel.Nama,
		Email:    dataModel.Email, //mapping data core ke data gorm model
		Password: dataModel.Password,
	}
}
