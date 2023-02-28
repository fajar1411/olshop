package handler

import (
	"mime/multipart"
	"toko/fitur/pelanggan"
)

type LoginRequest struct {
	Email    string `json:"email" form:"email"`
	Password string `json:"password" form:"password"`
}
type UpdateRequest struct {
	Nama       string `json:"name" form:"name"`
	Email      string `json:"email" form:"email"`
	Password   string `json:"password" form:"password"`
	Image_url  string `json:"image_url" form:"image_url"`
	FileHeader multipart.FileHeader
}
type PelangganRequest struct {
	Nama      string `json:"name" form:"name"`
	Email     string `json:"email" form:"email"`
	Password  string `json:"password" form:"password"`
	Image_url string `json:"images" form:"images"`
}

func UpdateRequestToUserCore(data UpdateRequest) pelanggan.PelangganEntites {
	return pelanggan.PelangganEntites{
		Nama:      data.Nama,
		Password:  data.Password,
		Email:     data.Email,
		Image_url: data.Image_url,
	}
}

func UserRequestToUserCore(data PelangganRequest) pelanggan.PelangganEntites {
	return pelanggan.PelangganEntites{
		Nama:     data.Nama,
		Password: data.Password,
		Email:    data.Email,
	}
}
