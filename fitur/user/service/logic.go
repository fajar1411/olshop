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

func NewService(ud user.UserData, vld *validator.Validate) user.UserService {
	return &userCase{
		qry: ud,
		vld: vld,
	}
}

// Login implements user.UserService
func (Uc *userCase) Login(email string, password string) (string, user.UserEntites, error) {

	errEmail := Uc.vld.Var(email, "required,email")
	if errEmail != nil {
		log.Println("validation error", errEmail)
		msg := validasi.ValidationErrorHandle(errEmail)
		return "", user.UserEntites{}, errors.New(msg)
	}
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
	errPw := Uc.vld.Var(password, "required,min=5,required")
	if errPw != nil {
		log.Println("validation error", errPw)
		msg := validasi.ValidationErrorHandle(errPw)
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
	valerr := Uc.vld.Struct(&newUser)
	if valerr != nil {
		log.Println("validation error", valerr)
		msg := validasi.ValidationErrorHandle(valerr)
		return user.UserEntites{}, errors.New(msg)
	}

	hash := scripts.Bcript(newUser.Password)
	newUser.Password = string(hash)
	res, err := Uc.qry.Register(newUser)
	if err != nil {
		msg2 := ""
		if strings.Contains(err.Error(), "duplicated") {
			msg2 = "email sudah terdaftar"
		} else if strings.Contains(err.Error(), "empty") {
			msg2 = "username not allowed empty"
		} else {
			msg2 = "server error"
		}
		return user.UserEntites{}, errors.New(msg2)
	}

	return res, nil
}

// Profile implements user.UserService
func (Uc *userCase) Profile(id int) (user.UserEntites, error) {
	if id <= 0 {
		log.Println("User belum terdaftar")
	}
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

// UpdateUser implements user.UserService
func (Uc *userCase) UpdateUser(id int, Updata user.UserEntites) (user.UserEntites, error) {
	if id <= 0 {
		log.Println("User belum terdaftar")
	}

	email := Updata.Email
	errEmail := Uc.vld.Var(email, "required,email")
	if errEmail != nil {
		log.Println("validation error", errEmail)
		msg := validasi.ValidationErrorHandle(errEmail)
		return user.UserEntites{}, errors.New(msg)
	}
	name := Updata.Nama
	errName := Uc.vld.Var(name, "required,min=5,required")
	if errName != nil {
		log.Println("validation error", errName)
		msg := validasi.ValidationErrorHandle(errName)
		return user.UserEntites{}, errors.New(msg)
	}
	pw := Updata.Password
	if pw != "" {
		errPw := Uc.vld.Var(pw, "required,min=5,required")
		if errPw != nil {
			log.Println("validation error", errPw)
			msg := validasi.ValidationErrorHandle(errPw)
			return user.UserEntites{}, errors.New(msg)
		} else {
			hash := scripts.Bcript(Updata.Password)
			Updata.Password = string(hash)
		}

	}

	res, err := Uc.qry.UpdateUser(id, Updata)
	if err != nil {
		msg2 := ""
		if strings.Contains(err.Error(), "duplicated") {
			msg2 = "email sudah terdaftar"
		} else if strings.Contains(err.Error(), "empty") {
			msg2 = "username not allowed empty"
		} else {
			msg2 = "server error"
		}
		return user.UserEntites{}, errors.New(msg2)
	}

	return res, nil
}
