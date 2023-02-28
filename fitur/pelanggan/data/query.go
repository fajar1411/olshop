package data

import (
	"errors"
	"log"
	"strings"

	pelanggan "toko/fitur/pelanggan"

	"gorm.io/gorm"
)

type pelangganData struct {
	db *gorm.DB
}

// Profile implements user.UserData

func NewPelanggan(db *gorm.DB) pelanggan.PelangganData {
	return &pelangganData{
		db: db,
	}
}

// Login implements user.UserData
func (ud *pelangganData) Login(email string) (pelanggan.PelangganEntites, error) {
	res := Pelanggan{}

	if err := ud.db.Where("email = ?", email).First(&res).Error; err != nil {
		log.Println("login query error", err.Error())
		return pelanggan.PelangganEntites{}, errors.New("data not found")
	}

	return ToCore(res), nil
}

// Register implements user.UserData
func (Ud *pelangganData) Register(newUser pelanggan.PelangganEntites) (pelanggan.PelangganEntites, error) {
	userGorm := FromEntities(newUser)
	userGorm.Role = "Pelanggan"

	tx := Ud.db.Create(&userGorm) // proses insert data

	if tx.Error != nil {
		log.Println("register query error", tx.Error.Error())
		msg := ""
		if strings.Contains(tx.Error.Error(), "Duplicate") {
			msg = "data is duplicated"
		} else {
			msg = "server error"
		}
		return pelanggan.PelangganEntites{}, errors.New(msg)
	}
	return newUser, nil
}
func (Ud *pelangganData) Profile(id int) (pelanggan.PelangganEntites, error) {
	var users Pelanggan

	if err := Ud.db.Where("id = ?", id).First(&users).Error; err != nil {
		log.Println("Get By ID query error", err.Error())
		return pelanggan.PelangganEntites{}, err
	}
	gorms := users.ModelsToCore()
	return gorms, nil
}

// UpdateUser implements user.UserData
func (Ud *pelangganData) UpdateUser(id int, Updata pelanggan.PelangganEntites) (pelanggan.PelangganEntites, error) {
	var users Pelanggan
	datacore := FromEntities(Updata)
	qry := Ud.db.Model(&users).Where("id = ?", id).Updates(&datacore)

	affrows := qry.RowsAffected
	if affrows == 0 {
		log.Println("no rows affected")
		return pelanggan.PelangganEntites{}, errors.New("no data updated")
	}
	err := qry.Error
	if err != nil {
		log.Println("update user query error", err.Error())
		return pelanggan.PelangganEntites{}, err
	}

	return ToCore(datacore), nil
}

// DeleteUser implements user.UserData
func (Ud *pelangganData) DeleteUser(id int) (pelanggan.PelangganEntites, error) {
	panic("unimplemented")
}
