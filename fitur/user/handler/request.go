package handler

import "toko/fitur/user"

type LoginRequest struct {
	Email    string `json:"email" form:"email"`
	Password string `json:"password" form:"password"`
}

type UserRequest struct {
	Nama     string `json:"name" form:"name"`
	Email    string `json:"email" form:"email"`
	Password string `json:"password" form:"password"`
}

func UserRequestToUserCore(data UserRequest) user.UserEntites {
	return user.UserEntites{
		Nama:     data.Nama,
		Password: data.Password,
		Email:    data.Email,
	}
}
