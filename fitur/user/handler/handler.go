package handler

import (
	"net/http"
	"toko/fitur/user"
	"toko/helper"
	"toko/middlewares"

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
	dataResp := ToRegisterResponse(res)
	return c.JSON(http.StatusCreated, helper.PesanDataBerhasilHelper("Register Berhasil", dataResp))
}

func (Ud *UserHandler) Login(c echo.Context) error {
	input := LoginRequest{}
	if err := c.Bind(&input); err != nil {
		return c.JSON(http.StatusBadRequest, "format inputan salah")
	}

	token, res, err := Ud.UserServices.Login(input.Email, input.Password)
	if err != nil {
		return c.JSON(helper.PesanGagalHelper(err.Error()))
	}
	dataResp := ToLoginRespon(res, token)
	return c.JSON(http.StatusOK, helper.PesanDataBerhasilHelper("Login berhasil", dataResp))
}
func (Ud *UserHandler) Profile(c echo.Context) error {
	id := middlewares.ExtractTokenUserId(c)

	res, err := Ud.UserServices.Profile(id)
	if err != nil {
		return c.JSON(helper.PesanGagalHelper(err.Error()))
	}
	dataResp := ToResponses(res)
	return c.JSON(http.StatusOK, helper.PesanDataBerhasilHelper("Melihat Profile Berhasil", dataResp))
}

func (Ud *UserHandler) Update(c echo.Context) error {
	input := UpdateRequest{}
	if err := c.Bind(&input); err != nil {
		return c.JSON(http.StatusBadRequest, "format inputan salah")
	}
	id := middlewares.ExtractTokenUserId(c)

	res, err := Ud.UserServices.UpdateUser(id, UpdateRequestToUserCore(input))
	if err != nil {
		return c.JSON(helper.PesanGagalHelper(err.Error()))
	}
	dataResp := UpdateRespons(res)
	return c.JSON(http.StatusOK, helper.PesanDataBerhasilHelper("Update berhasil", dataResp))
}
