package handler

import (
	"net/http"
	"toko/fitur/user"
	"toko/helper"

	"github.com/labstack/echo/v4"
)

type UserHandler struct {
	UserServices user.UserService
}

func (Ud *UserHandler) Register(c echo.Context) error {
	InputUser := UserRequest{}

	errbind := c.Bind(&InputUser)
	if errbind != nil {
		return c.JSON(http.StatusBadRequest, "format inputan salah")
	}
	dataCore := UserRequestToUserCore(InputUser)
	res, err := Ud.UserServices.Register(dataCore)
	if err != nil {
		return c.JSON(helper.PesanGagalHelper(err.Error()))
	}
	dataResp := ToResponse(res)
	return c.JSON(http.StatusCreated, helper.PesanDataBerhasilHelper("Register Berhasil", dataResp))
}

func (Ud *UserHandler) Login(c echo.Context) error {
	panic("")
}
