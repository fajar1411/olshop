package service

import (
	"toko/fitur/user"

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
	panic("unimplemented")
}
