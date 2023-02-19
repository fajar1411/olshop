package service

import (
	"errors"
	"strings"
	"toko/fitur/user"
	"toko/token"
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
	panic("unimplemented")
}

// Register implements user.UserService
func (Uc *userCase) Register(newUser user.UserEntites) (user.UserEntites, error) {
	valerr := validasi.Validasi(validasi.ToRegister(newUser))
	if valerr != nil {
		return user.UserEntites{}, valerr
	}

	hash := token.Bcript(newUser.Password)
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
