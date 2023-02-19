package handler

import (
	"toko/fitur/user"

	"github.com/labstack/echo/v4"
)

type UserHandler struct {
	UserServices user.UserService
}

func (Ud *UserHandler) Create(c echo.Context) error {
	panic("")
}

func (Ud *UserHandler) Login(c echo.Context) error {
	panic("")
}
