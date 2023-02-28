package handler

import (
	"fmt"
	"log"
	"net/http"
	"toko/fitur/pelanggan"
	"toko/helper"
	"toko/middlewares"

	"github.com/labstack/echo/v4"
)

type PelangganHandler struct {
	PelangganServices pelanggan.PelangganService
}

func (Ud *PelangganHandler) Register(c echo.Context) error {
	InputUser := PelangganRequest{}

	errbind := c.Bind(&InputUser)
	if errbind != nil {
		return c.JSON(http.StatusBadRequest, "format inputan salah")
	}
	dataCore := UserRequestToUserCore(InputUser)
	res, err := Ud.PelangganServices.Register(dataCore)
	if err != nil {
		return c.JSON(helper.PesanGagalHelper(err.Error()))
	}
	dataResp := ToRegisterResponse(res)
	return c.JSON(http.StatusCreated, helper.PesanDataBerhasilHelper("Register Berhasil", dataResp))
}

func (Ud *PelangganHandler) Login(c echo.Context) error {
	input := LoginRequest{}
	if err := c.Bind(&input); err != nil {
		return c.JSON(http.StatusBadRequest, "format inputan salah")
	}

	token, res, err := Ud.PelangganServices.Login(input.Email, input.Password)
	if err != nil {
		return c.JSON(helper.PesanGagalHelper(err.Error()))
	}
	dataResp := ToLoginRespon(res, token)
	return c.JSON(http.StatusOK, helper.PesanDataBerhasilHelper("Login berhasil", dataResp))
}
func (Ud *PelangganHandler) Profile(c echo.Context) error {
	id := middlewares.ExtractTokenUserId(c)

	res, err := Ud.PelangganServices.Profile(id)
	if err != nil {
		return c.JSON(helper.PesanGagalHelper(err.Error()))
	}
	dataResp := ToResponses(res)
	return c.JSON(http.StatusOK, helper.PesanDataBerhasilHelper("Melihat Profile Berhasil", dataResp))
}

func (Ud *PelangganHandler) Update(c echo.Context) error {
	input := UpdateRequest{}
	if err := c.Bind(&input); err != nil {
		return c.JSON(http.StatusBadRequest, "format inputan salah")
	}
	id := middlewares.ExtractTokenUserId(c)

	formHeader, err := c.FormFile("images")
	fmt.Print("image handler", formHeader.Filename)
	if err != nil {
		log.Println(err)
	}
	res, err := Ud.PelangganServices.UpdateUser(id, formHeader, UpdateRequestToUserCore(input))
	if err != nil {
		return c.JSON(helper.PesanGagalHelper(err.Error()))
	}
	dataResp := UpdateRespons(res)
	return c.JSON(http.StatusOK, helper.PesanDataBerhasilHelper("Update berhasil", dataResp))
}
