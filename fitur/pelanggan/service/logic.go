package service

import (
	"errors"
	"fmt"
	"log"
	"mime/multipart"
	"strings"
	"toko/fitur/pelanggan"
	"toko/helper"

	"toko/middlewares"
	"toko/scripts"
	"toko/validasi"

	"github.com/go-playground/validator/v10"
)

type pelangganCase struct {
	qry pelanggan.PelangganData
	vld *validator.Validate
	hps helper.Uploads
}

func NewService(ud pelanggan.PelangganData, vld *validator.Validate, hps helper.Uploads) pelanggan.PelangganService {
	return &pelangganCase{
		qry: ud,
		vld: vld,
		hps: hps,
	}
}

// Login implements user.UserService
func (Uc *pelangganCase) Login(email string, password string) (string, pelanggan.PelangganEntites, error) {

	errEmail := Uc.vld.Var(email, "required,email")
	if errEmail != nil {
		log.Println("validation error", errEmail)
		msg := validasi.ValidationErrorHandle(errEmail)
		return "", pelanggan.PelangganEntites{}, errors.New(msg)
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
		return "", pelanggan.PelangganEntites{}, errors.New(msg)
	}
	errPw := Uc.vld.Var(password, "required,min=5,required")
	if errPw != nil {
		log.Println("validation error", errPw)
		msg := validasi.ValidationErrorHandle(errPw)
		return "", pelanggan.PelangganEntites{}, errors.New(msg)
	}
	if err := scripts.CheckPassword(res.Password, password); err != nil {
		log.Println("login compare", err.Error())
		return "", pelanggan.PelangganEntites{}, errors.New("password tidak sesuai" + res.Password)
	}

	//Token expires after 1 hour
	token, _ := middlewares.CreateToken(int(res.ID), res.Role)

	return token, res, nil

}

// Register implements user.UserService
func (Uc *pelangganCase) Register(newUser pelanggan.PelangganEntites) (pelanggan.PelangganEntites, error) {
	valerr := Uc.vld.Struct(&newUser)
	if valerr != nil {
		log.Println("validation error", valerr)
		msg := validasi.ValidationErrorHandle(valerr)
		return pelanggan.PelangganEntites{}, errors.New(msg)
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
		return pelanggan.PelangganEntites{}, errors.New(msg2)
	}

	return res, nil
}

// Profile implements user.UserService
func (Uc *pelangganCase) Profile(id int) (pelanggan.PelangganEntites, error) {
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
		return pelanggan.PelangganEntites{}, errors.New(msg)
	}
	return res, nil
}

// UpdateUser implements user.UserService
func (Uc *pelangganCase) UpdateUser(id int, fileData *multipart.FileHeader, Updata pelanggan.PelangganEntites) (pelanggan.PelangganEntites, error) {
	if id <= 0 {
		log.Println("User belum terdaftar")
	}
	if fileData != nil {
		secureURL, err2 := Uc.hps.Upload(fileData)
		if err2 != nil {
			log.Println(err2)
			fmt.Print(err2)
			var msg string
			if strings.Contains(err2.Error(), "bad request") {
				msg = err2.Error()
			} else {
				msg = "failed to upload image, server error"
			}
			return pelanggan.PelangganEntites{}, errors.New(msg)
		}
		Updata.Image_url = secureURL
		fmt.Print("update data image", Updata.Image_url)
	}
	email := Updata.Email
	errEmail := Uc.vld.Var(email, "required,email")
	if errEmail != nil {
		log.Println("validation error", errEmail)
		msg := validasi.ValidationErrorHandle(errEmail)
		return pelanggan.PelangganEntites{}, errors.New(msg)
	}
	name := Updata.Nama
	errName := Uc.vld.Var(name, "required,min=5,required")
	if errName != nil {
		log.Println("validation error", errName)
		msg := validasi.ValidationErrorHandle(errName)
		return pelanggan.PelangganEntites{}, errors.New(msg)
	}
	pw := Updata.Password
	if pw != "" {
		errPw := Uc.vld.Var(pw, "required,min=5,required")
		if errPw != nil {
			log.Println("validation error", errPw)
			msg := validasi.ValidationErrorHandle(errPw)
			return pelanggan.PelangganEntites{}, errors.New(msg)
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
		return pelanggan.PelangganEntites{}, errors.New(msg2)
	}

	return res, nil
}

// DeleteUser implements user.UserService
func (Uc *pelangganCase) DeleteUser(id int) (pelanggan.PelangganEntites, error) {
	panic("unimplemented")
}
