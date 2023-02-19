package service

import (
	"errors"
	"log"
	"strings"
	"toko/fitur/user"
	"toko/middlewares"
	"toko/scripts"
	"toko/validasi"

	"github.com/go-playground/validator/v10"
)

type userCase struct {
	qry user.UserData
	vld *validator.Validate
}

func NewService(ud user.UserData) user.UserService {
	return &userCase{
		qry: ud,
		vld: validator.New(),
	}
}

// Login implements user.UserService
func (Uc *userCase) Login(email string, password string) (string, user.UserEntites, error) {
	res, err := Uc.qry.Login(email)
	if err != nil {
		log.Println("query login error", err.Error())
		msg := ""
		if strings.Contains(err.Error(), "not found") || strings.Contains(err.Error(), "no rows") {
			msg = "email belum terdaftar"
		} else {
			msg = "terdapat masalah pada server"
		}
		return "", user.UserEntites{}, errors.New(msg)
	}

	if err := scripts.CheckPassword(res.Password, password); err != nil {
		log.Println("login compare", err.Error())
		return "", user.UserEntites{}, errors.New("password tidak sesuai" + res.Password)
	}

	//Token expires after 1 hour
	token, _ := middlewares.CreateToken(int(res.ID))

	return token, res, nil

}

// Register implements user.UserService
func (Uc *userCase) Register(newUser user.UserEntites) (user.UserEntites, error) {
	valerr := validasi.Validasi(validasi.ToRegister(newUser))
	if valerr != nil {
		return user.UserEntites{}, valerr
	}

	hash := scripts.Bcript(newUser.Password)
	newUser.Password = string(hash)
	res, err := Uc.qry.Register(newUser)
	if err != nil {
		msg := ""
		if strings.Contains(err.Error(), "duplicated") {
			msg = "email sudah terdaftar"
		} else {
			msg = "terdapat masalah pada server"
		}
		return user.UserEntites{}, errors.New(msg)
	}

	return res, nil
}

// Profile implements user.UserService
func (Uc *userCase) Profile(id int) (user.UserEntites, error) {

	res, err := Uc.qry.Profile(id)
	if err != nil {
		log.Println(err)
		msg := ""
		if strings.Contains(err.Error(), "not found") {
			msg = "user tidak ditemukan harap login lagi"
		} else {
			msg = "terdapat masalah pada server"
		}
		return user.UserEntites{}, errors.New(msg)
	}
	return res, nil
}
