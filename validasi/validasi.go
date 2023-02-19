package validasi

import (
	"errors"
	"log"
	"strings"
	user "toko/fitur/user"

	"github.com/go-playground/validator/v10"
)

var validate *validator.Validate

type RegisterValidate struct {
	Name     string `validate:"required"`
	Email    string `validate:"required,email" isuniqe:"users,email"`
	Password string
}

func ToRegister(data user.UserEntites) RegisterValidate {
	return RegisterValidate{
		Name:     data.Nama,
		Email:    data.Email,
		Password: data.Password,
	}
}

func Validasi(data interface{}) error {
	validate = validator.New()
	err := validate.Struct(data)
	if err != nil {
		log.Println(err)
		if _, ok := err.(*validator.InvalidValidationError); ok {
			log.Println(err)
		}
		msg := ""
		if strings.Contains(err.Error(), "required") {
			msg = "field required wajib diisi"
		} else if strings.Contains(err.Error(), "email") {
			msg = "format email salah"
		} else if strings.Contains(err.Error(), "duplicated") {
			msg = "email sudah terdaftar"
		}
		return errors.New(msg)
	}
	return nil
}
